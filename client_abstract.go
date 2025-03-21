package sdkgen

import (
	"net/http"
)

type ClientAbstract struct {
	Authenticator AuthenticatorInterface
	HttpClient    *http.Client
	Parser        *Parser
}

func NewClient(baseUrl string, credentials CredentialsInterface) (*ClientAbstract, error) {
	authenticator, err := AuthenticatorFactory(credentials)
	if err != nil {
		return nil, err
	}

	return &ClientAbstract{
		Authenticator: authenticator,
		HttpClient:    HttpClientFactory(authenticator),
		Parser: &Parser{
			BaseUrl: baseUrl,
		},
	}, nil
}

func NewClientWithVersion(baseUrl string, credentials CredentialsInterface, version string) (*ClientAbstract, error) {
	authenticator, err := AuthenticatorFactory(credentials)
	if err != nil {
		return nil, err
	}

	return &ClientAbstract{
		Authenticator: authenticator,
		HttpClient:    HttpClientFactoryWithVersion(authenticator, version),
		Parser: &Parser{
			BaseUrl: baseUrl,
		},
	}, nil
}
