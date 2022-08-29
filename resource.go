package sdkgen

import "net/http"

type Resource struct {
	BaseUrl    string
	HttpClient *http.Client
}
