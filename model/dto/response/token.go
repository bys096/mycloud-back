package response

type UserToken struct {
	Email        string `json:"email"`
	ExpiresIn    int64  `json:"expires_in"`
	Scope        string `json:"scope"`
	UserId       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
