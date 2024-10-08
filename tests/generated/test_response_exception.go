// TestResponseException automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package generated

import (
	"encoding/json"
	"fmt"
)

type TestResponseException struct {
	Payload  TestResponse
	Previous error
}

func (e *TestResponseException) Error() string {
	raw, err := json.Marshal(e.Payload)
	if err != nil {
		return "could not marshal provided JSON data"
	}

	return fmt.Sprintf("The server returned an error: %s", raw)
}
