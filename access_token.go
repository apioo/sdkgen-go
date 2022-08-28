package sdkgen

type AccessToken struct {
	TokenType    string
	AccessToken  string
	ExpiresIn    int64
	RefreshToken string
	Scope        string
}
