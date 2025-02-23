package sdkgen

import "net/http"

func HttpClientFactory(authenticator AuthenticatorInterface) *http.Client {
	return &http.Client{
		Transport: &DefaultTransport{
			Authenticator: authenticator,
		},
	}
}

func HttpClientFactoryWithVersion(authenticator AuthenticatorInterface, version string) *http.Client {
	return &http.Client{
		Transport: &DefaultTransport{
			Authenticator: authenticator,
			Version:       version,
		},
	}
}

type DefaultTransport struct {
	Authenticator AuthenticatorInterface
	Version       string
}

func (transport *DefaultTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if transport.Version != "" {
		req.Header.Add("User-Agent", "SDKgen/"+transport.Version)
	} else {
		req.Header.Add("User-Agent", "SDKgen")
	}
	req.Header.Add("Accept", "application/json")

	req, err := transport.Authenticator.Intercept(req)
	if err != nil {
		return nil, err
	}

	return http.DefaultTransport.RoundTrip(req)
}
