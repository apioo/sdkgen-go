package sdkgen

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type Client struct {
	BaseUrl     string
	Credentials CredentialsInterface
	TokenStore  TokenStoreInterface
	Scopes      []string
}

func (client Client) BuildRedirectUrl(redirectUrl string, scopes []string, state string) (string, error) {
	var credentials = client.Credentials.(AuthorizationCode)

	authUrl, err := url.Parse(credentials.AuthorizationUrl)
	if err != nil {
		return "", errors.New("could not parse authorization url")
	}

	authUrl.Query().Add("response_type", "code")
	authUrl.Query().Add("client_id", credentials.ClientId)

	if redirectUrl != "" {
		authUrl.Query().Add("redirect_uri", redirectUrl)
	}

	if len(scopes) > 0 {
		authUrl.Query().Add("scopes", strings.Join(scopes, ","))
	}

	if state != "" {
		authUrl.Query().Add("state", state)
	}

	return authUrl.String(), nil
}

func (client Client) FetchAccessTokenByCode(code string) (AccessToken, error) {
	var credentials = client.Credentials.(AuthorizationCode)
	var httpClient = client.NewHttpClient(HttpBasic{UserName: credentials.ClientId, Password: credentials.ClientSecret})

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)

	req, err := http.NewRequest("POST", credentials.TokenUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return AccessToken{}, errors.New("could create request to obtain access token by code")
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return AccessToken{}, errors.New("could not send obtain access token by code")
	}

	return client.ParseTokenResponse(resp)
}

func (client Client) FetchAccessTokenByClientCredentials() (AccessToken, error) {
	var credentials = client.Credentials.(ClientCredentials)
	var httpClient = client.NewHttpClient(HttpBasic{UserName: credentials.ClientId, Password: credentials.ClientSecret})

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	if len(client.Scopes) > 0 {
		data.Set("scope", strings.Join(client.Scopes, ","))
	}

	req, err := http.NewRequest("POST", credentials.TokenUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return AccessToken{}, errors.New("could create request to obtain access token by code")
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return AccessToken{}, errors.New("could not send obtain access token by code")
	}

	return client.ParseTokenResponse(resp)
}

func (client Client) FetchAccessTokenByRefresh(refreshToken string) (AccessToken, error) {
	var credentials = client.Credentials.(AuthorizationCode)
	var httpClient = client.NewHttpClient(HttpBasic{UserName: credentials.ClientId, Password: credentials.ClientSecret})

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	req, err := http.NewRequest("POST", credentials.TokenUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return AccessToken{}, errors.New("could create request to obtain access token by code")
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return AccessToken{}, errors.New("could not send obtain access token by code")
	}

	return client.ParseTokenResponse(resp)
}

func (client Client) NewHttpClient(credentials CredentialsInterface) *http.Client {
	instance := &http.Client{
		Transport: NewAuthorizationTransport(credentials, client),
	}
	return instance
}

func (client Client) GetAccessToken(automaticRefresh bool, expireThreshold int64) (string, error) {
	timestamp := time.Now().Unix()

	accessToken, err := client.TokenStore.Get()
	if err == nil || accessToken.ExpiresIn < timestamp {
		accessToken, err = client.FetchAccessTokenByClientCredentials()
	}

	if err != nil {
		return "", errors.New("found no access Token, please obtain an access Token before making a request")
	}

	if accessToken.ExpiresIn > (timestamp + expireThreshold) {
		return accessToken.AccessToken, nil
	}

	if automaticRefresh && accessToken.RefreshToken != "" {
		accessToken, err = client.FetchAccessTokenByRefresh(accessToken.RefreshToken)
		if err != nil {
			return "", errors.New("could not refresh access token")
		}
	}

	return accessToken.AccessToken, nil
}

func (client Client) ParseTokenResponse(resp *http.Response) (AccessToken, error) {
	if resp.StatusCode != 200 {
		return AccessToken{}, errors.New("could not obtain access Token, received a non successful status code: " + resp.Status)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AccessToken{}, errors.New("could not read response body")
	}

	var token AccessToken
	err = json.Unmarshal(respBody, &token)
	if err != nil {
		return AccessToken{}, errors.New("could not unmarshal access Token")
	}

	if token.AccessToken != "" {
		return AccessToken{}, errors.New("could not obtain access Token")
	}

	err = client.TokenStore.Persist(token)
	if err != nil {
		return AccessToken{}, err
	}

	return token, nil
}

func (client Client) GetResource() *Resource {
	return &Resource{
		BaseUrl:    client.BaseUrl,
		HttpClient: client.NewHttpClient(client.Credentials),
	}
}

func NewClient(baseUrl string, credentials CredentialsInterface, tokenStore TokenStoreInterface, scopes []string) *Client {
	return &Client{
		BaseUrl:     baseUrl,
		Credentials: credentials,
		TokenStore:  tokenStore,
		Scopes:      scopes,
	}
}

type AuthorizationTransport struct {
	Credentials CredentialsInterface
	Client      Client
}

func (transport *AuthorizationTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", "SDKgen Client v0.1")
	req.Header.Add("Accept", "application/json")

	if reflect.TypeOf(transport.Credentials).Name() == "HttpBasic" {
		var cred = transport.Credentials.(HttpBasic)
		req.Header.Add("Authorization", "Basic "+cred.UserName+":"+cred.Password)
	} else if reflect.TypeOf(transport.Credentials).Name() == "HttpBearer" {
		var cred = transport.Credentials.(HttpBearer)
		req.Header.Add("Authorization", "Bearer "+cred.Token)
	} else if reflect.TypeOf(transport.Credentials).Name() == "ApiKey" {
		var cred = transport.Credentials.(ApiKey)
		req.Header.Add(cred.Name, cred.Token)
	} else if reflect.TypeOf(transport.Credentials).Name() == "AuthorizationCode" || reflect.TypeOf(transport.Credentials).Name() == "ClientCredentials" {
		accessToken, err := transport.Client.GetAccessToken(true, 60*10)
		if err == nil {
			req.Header.Add("Authorization", "Bearer "+accessToken)
		}
	}

	return http.DefaultTransport.RoundTrip(req)
}

func NewAuthorizationTransport(credentials CredentialsInterface, client Client) *AuthorizationTransport {
	return &AuthorizationTransport{credentials, client}
}
