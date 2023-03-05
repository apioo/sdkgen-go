package sdkgen

import (
	"net/http"
)

type Client struct {
	BaseUrl       string
	Authenticator *AuthenticatorInterface
	HttpClient    *http.Client
	Parser        *Parser
}

type Parser struct {
}

func HttpClientFactory(authenticator AuthenticatorInterface) *http.Client {
	return &http.Client{
		Transport: authenticator,
	}
}
