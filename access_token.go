package sdkgen

import "time"

type AccessToken struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func (accessToken *AccessToken) GetExpiresInTimestamp() int64 {
	nowTimestamp := time.Now().Unix()

	expiresIn := accessToken.ExpiresIn
	if expiresIn < 529196400 {
		// in case the expires in is lower than 1986-10-09 we assume that the field represents the duration in seconds
		// otherwise it is probably a timestamp
		expiresIn = nowTimestamp + expiresIn
	}

	return expiresIn
}
