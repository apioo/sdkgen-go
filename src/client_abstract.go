package sdkgen

import (
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type ClientAbstract struct {
	BaseUrl     string
	Credentials CredentialsInterface
	TokenStore  TokenStoreInterface
	Scopes      []string
}

func (client ClientAbstract) BuildRedirectUrl(redirectUrl string, scopes []string, state string) (string, error) {
	var credentials = client.Credentials.(AuthorizationCode)

	url, err := url.Parse(credentials.AuthorizationUrl)
	if err != nil {
		return "", errors.New("could not parse authorization url")
	}

	url.Query().Add("response_type", "code")
	url.Query().Add("client_id", credentials.ClientId)

	if redirectUrl != "" {
		url.Query().Add("redirect_uri", redirectUrl)
	}

	if len(scopes) > 0 {
		url.Query().Add("scopes", strings.Join(scopes, ","))
	}

	if state != "" {
		url.Query().Add("state", state)
	}

	return url.String(), nil
}

func (client ClientAbstract) FetchAccessTokenByCode() AccessToken {
	var credentials = client.Credentials.(AuthorizationCode)

	/*
	   httpClient = client.NewHttpClient(new HttpBasic{credentials.ClientId, credentials.ClientSecret})

	   public async fetchAccessTokenByCode(code: string): Promise<AccessToken> {
	       if (!(this.credentials instanceof AuthorizationCode)) {
	       throw new InvalidCredentialsException('The configured credentials do not support the OAuth2 authorization code flow');
	       }

	       const httpClient = await this.newHttpClient(new HttpBasic(this.credentials.clientId, this.credentials.clientSecret));

	       const response = await httpClient.post<AccessToken>(this.credentials.tokenUrl, {
	       grant_type: 'authorization_code',
	           code: code,
	       }, {
	       headers: {
	       Accept: 'application/json'
	       },
	       });

	       return this.parseTokenResponse(response);
	   }
	*/
}

func (client ClientAbstract) FetchAccessTokenByClientCredentials() {
	var credentials = client.Credentials.(ClientCredentials)

	/*
	   public async fetchAccessTokenByClientCredentials(): Promise<AccessToken> {
	       if (!(this.credentials instanceof ClientCredentials)) {
	       throw new InvalidCredentialsException('The configured credentials do not support the OAuth2 client credentials flow');
	       }

	       const httpClient = await this.newHttpClient(new HttpBasic(this.credentials.clientId, this.credentials.clientSecret));

	       let data: {grant_type: string, scope?: string} = {
	       grant_type: 'client_credentials'
	       };

	       if (this.scopes && this.scopes.length > 0) {
	           data.scope = this.scopes.join(',');
	       }

	       const response = await httpClient.post<AccessToken>(this.credentials.tokenUrl, data, {
	       headers: {
	       Accept: 'application/json'
	       },
	       });

	       return this.parseTokenResponse(response);
	   }

	*/
}

func (client ClientAbstract) FetchAccessTokenByRefresh() {
	var credentials = client.Credentials.(AuthorizationCode)

	/*
	   public async fetchAccessTokenByRefresh(refreshToken: string): Promise<AccessToken> {
	       if (!(this.credentials instanceof AuthorizationCode)) {
	       throw new InvalidCredentialsException('The configured credentials do not support the OAuth2 flow');
	       }

	       const httpClient = await this.newHttpClient(new HttpBasic(this.credentials.clientId, this.credentials.clientSecret));

	       const response = await httpClient.post<AccessToken>(this.credentials.tokenUrl, {
	       grant_type: 'refresh_token',
	           refresh_token: refreshToken,
	       }, {
	       headers: {
	       Accept: 'application/json'
	       },
	       });

	       return this.parseTokenResponse(response);
	   }

	*/
}

func (client ClientAbstract) NewHttpClient(credentials CredentialsInterface) http.Client {
	instance := http.Client{
		Transport: NewAuthorizationTransport(credentials, client),
	}
	return instance
}

func (client ClientAbstract) GetAccessToken() string {
	/*
	   private async getAccessToken(automaticRefresh: boolean = true, expireThreshold: number = ClientAbstract.EXPIRE_THRESHOLD): Promise<string> {
	       const timestamp = Math.floor(Date.now() / 1000);

	       let accessToken = this.tokenStore.get();
	       if ((!accessToken || accessToken.expires_in < timestamp) && this.credentials instanceof ClientCredentials) {
	           accessToken = await this.fetchAccessTokenByClientCredentials();
	       }

	       if (!accessToken) {
	           throw new FoundNoAccessTokenException('Found no access Token, please obtain an access Token before making a request');
	       }

	       if (accessToken.expires_in > (timestamp + expireThreshold)) {
	           return accessToken.access_token;
	       }

	       if (automaticRefresh && accessToken.refresh_token) {
	           accessToken = await this.fetchAccessTokenByRefresh(accessToken.refresh_token);
	       }

	       return accessToken.access_token;
	   }

	*/
}

func (client ClientAbstract) ParseTokenResponse() AccessToken {
	/*
	   private async parseTokenResponse(response: AxiosResponse<AccessToken>): Promise<AccessToken> {
	       if (response.status !== 200) {
	           throw new InvalidAccessTokenException('Could not obtain access Token, received a non successful status code: ' + response.status);
	       }

	       if (!response.data.access_token) {
	           throw new InvalidAccessTokenException('Could not obtain access Token');
	       }

	       this.tokenStore.persist(response.data);

	       return response.data;
	   }

	*/
}

type AuthorizationTransport struct {
	Credentials CredentialsInterface
	Client      ClientAbstract
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
		req.Header.Add("Authorization", "Bearer "+transport.Client.GetAccessToken())
	}

	return http.DefaultTransport.RoundTrip(req)
}

func NewAuthorizationTransport(credentials CredentialsInterface, client ClientAbstract) *AuthorizationTransport {
	return &AuthorizationTransport{credentials, client}
}
