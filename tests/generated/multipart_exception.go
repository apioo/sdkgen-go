// MultipartException automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package generated

import (
	"encoding/json"
	"fmt"

	"github.com/apioo/sdkgen-go"
)

type MultipartException struct {
	Payload  *sdkgen.Multipart
	Previous error
}

func (e *MultipartException) Error() string {
	raw, err := json.Marshal(e.Payload)
	if err != nil {
		return "could not marshal provided JSON data"
	}

	return fmt.Sprintf("The server returned an error: %s", raw)
}