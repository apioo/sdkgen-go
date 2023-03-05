package sdkgen

import "net/http"

func HttpClientFactory(authenticator AuthenticatorInterface) *http.Client {
	return &http.Client{
		Transport: authenticator,
	}
}
