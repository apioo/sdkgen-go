package sdkgen

import (
	"net/http"
)

type TagAbstract struct {
	HttpClient *http.Client
	Parser     *Parser
}
