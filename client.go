package sdkgen

import (
	"net/http"
)

type Client struct {
	Authenticator AuthenticatorInterface
	HttpClient    *http.Client
	Parser        *Parser
}

type TagAbstract struct {
	HttpClient *http.Client
	Parser     *Parser
}

func NewClient(baseUrl string, credentials CredentialsInterface) (*Client, error) {
	authenticator, err := AuthenticatorFactory(credentials)
	if err != nil {
		return nil, err
	}

	return &Client{
		Authenticator: authenticator,
		HttpClient:    HttpClientFactory(authenticator),
		Parser: &Parser{
			BaseUrl: baseUrl,
		},
	}, nil
}
