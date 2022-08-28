package sdkgen

import "net/http"

type ResourceAbstract struct {
	BaseUrl    string
	HttpClient http.Client
}
