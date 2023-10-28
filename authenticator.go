package sdkgen

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type AuthenticatorInterface interface {
	http.RoundTripper
	RoundTrip(*http.Request) (*http.Response, error)
}

type AnonymousAuthenticator struct {
}

func (authenticator *AnonymousAuthenticator) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", "SDKgen Client v1.0")
	req.Header.Add("Accept", "application/json")

	return http.DefaultTransport.RoundTrip(req)
}

type HttpBasicAuthenticator struct {
	Credentials HttpBasic
}

func (authenticator *HttpBasicAuthenticator) RoundTrip(req *http.Request) (*http.Response, error) {
	var auth = base64.StdEncoding.EncodeToString([]byte(authenticator.Credentials.UserName + ":" + authenticator.Credentials.Password))
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("User-Agent", "SDKgen Client v1.0")
	req.Header.Add("Accept", "application/json")

	return http.DefaultTransport.RoundTrip(req)
}

type HttpBearerAuthenticator struct {
	Credentials HttpBearer
}

func (authenticator *HttpBearerAuthenticator) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+authenticator.Credentials.Token)
	req.Header.Add("User-Agent", "SDKgen Client v1.0")
	req.Header.Add("Accept", "application/json")

	return http.DefaultTransport.RoundTrip(req)
}

type ApiKeyAuthenticator struct {
	Credentials ApiKey
}

func (authenticator *ApiKeyAuthenticator) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add(authenticator.Credentials.Name, authenticator.Credentials.Token)
	req.Header.Add("User-Agent", "SDKgen Client v1.0")
	req.Header.Add("Accept", "application/json")

	return http.DefaultTransport.RoundTrip(req)
}

type OAuth2Authenticator struct {
	Credentials OAuth2
}

func (authenticator *OAuth2Authenticator) RoundTrip(req *http.Request) (*http.Response, error) {
	accessToken, err := authenticator.GetAccessToken(true, 60*10)
	if err == nil {
		req.Header.Add("Authorization", "Bearer "+accessToken)
	}

	req.Header.Add("User-Agent", "SDKgen Client v1.0")
	req.Header.Add("Accept", "application/json")

	return http.DefaultTransport.RoundTrip(req)
}

func (authenticator *OAuth2Authenticator) BuildRedirectUrl(redirectUrl string, scopes []string, state string) (string, error) {
	authUrl, err := url.Parse(authenticator.Credentials.AuthorizationUrl)
	if err != nil {
		return "", errors.New("could not parse authorization url")
	}

	authUrl.Query().Add("response_type", "code")
	authUrl.Query().Add("client_id", authenticator.Credentials.ClientId)

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

func (authenticator *OAuth2Authenticator) FetchAccessTokenByCode(code string) (AccessToken, error) {
	var httpClient = HttpClientFactory(&HttpBasicAuthenticator{
		Credentials: HttpBasic{
			UserName: authenticator.Credentials.ClientId,
			Password: authenticator.Credentials.ClientSecret,
		},
	})

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)

	req, err := http.NewRequest("POST", authenticator.Credentials.TokenUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return AccessToken{}, errors.New("could create request to obtain access token by code")
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		return AccessToken{}, errors.New("could not send obtain access token by code")
	}

	return authenticator.ParseTokenResponse(resp)
}

func (authenticator *OAuth2Authenticator) FetchAccessTokenByClientCredentials() (AccessToken, error) {
	var httpClient = HttpClientFactory(&HttpBasicAuthenticator{
		Credentials: HttpBasic{
			UserName: authenticator.Credentials.ClientId,
			Password: authenticator.Credentials.ClientSecret,
		},
	})

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	if len(authenticator.Credentials.Scopes) > 0 {
		data.Set("scope", strings.Join(authenticator.Credentials.Scopes, ","))
	}

	req, err := http.NewRequest("POST", authenticator.Credentials.TokenUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return AccessToken{}, errors.New("could create request to obtain access token by code")
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		return AccessToken{}, errors.New("could not send obtain access token by code")
	}

	return authenticator.ParseTokenResponse(resp)
}

func (authenticator *OAuth2Authenticator) FetchAccessTokenByRefresh(refreshToken string) (AccessToken, error) {
	var httpClient = HttpClientFactory(&HttpBasicAuthenticator{
		Credentials: HttpBasic{
			UserName: authenticator.Credentials.ClientId,
			Password: authenticator.Credentials.ClientSecret,
		},
	})

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	req, err := http.NewRequest("POST", authenticator.Credentials.TokenUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return AccessToken{}, errors.New("could create request to obtain access token by code")
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		return AccessToken{}, errors.New("could not send obtain access token by code")
	}

	return authenticator.ParseTokenResponse(resp)
}

func (authenticator *OAuth2Authenticator) GetAccessToken(automaticRefresh bool, expireThreshold int64) (string, error) {
	timestamp := time.Now().Unix()

	accessToken, err := authenticator.Credentials.TokenStore.Get()
	if err == nil || accessToken.GetExpiresInTimestamp() < timestamp {
		accessToken, err = authenticator.FetchAccessTokenByClientCredentials()
	}

	if err != nil {
		return "", errors.New("found no access Token, please obtain an access Token before making a request")
	}

	if accessToken.GetExpiresInTimestamp() > (timestamp + expireThreshold) {
		return accessToken.AccessToken, nil
	}

	if automaticRefresh && accessToken.RefreshToken != "" {
		accessToken, err = authenticator.FetchAccessTokenByRefresh(accessToken.RefreshToken)
		if err != nil {
			return "", errors.New("could not refresh access token")
		}
	}

	return accessToken.AccessToken, nil
}

func (authenticator *OAuth2Authenticator) ParseTokenResponse(resp *http.Response) (AccessToken, error) {
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

	if token.AccessToken == "" {
		return AccessToken{}, errors.New("could not obtain access Token")
	}

	err = authenticator.Credentials.TokenStore.Persist(token)
	if err != nil {
		return AccessToken{}, err
	}

	return token, nil
}
