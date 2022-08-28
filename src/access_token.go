package sdkgen

type AccessToken struct {
	TokenType    string
	AccessToken  string
	ExpiresIn    int
	RefreshToken string
	Scope        string
}
