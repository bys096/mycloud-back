package config

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	googleOAuth "golang.org/x/oauth2/google"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
	"log"
	"mycloud/model/dto/response"
	"os"
)

type Google struct {
	Config *oauth2.Config
}

func NewGoogle() *Google {
	return newGoogle()
}

func newGoogle() *Google {
	google := &Google{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Endpoint:     googleOAuth.Endpoint,
			Scopes:       []string{"openid", "email", "profile"},
			RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		},
	}

	if google.Config == nil {
		panic("==== invalid key. google api ====")
	}

	return google
}

func (g *Google) GetLoginURL(state string) (clientID string) {
	return g.Config.AuthCodeURL(state)
}

func (g *Google) GenerateToken(code string) (*response.UserToken, error) {
	cxt := context.Background()
	httpClient, _ := g.Config.Exchange(cxt, code)
	if httpClient == nil {
		return &response.UserToken{}, errors.New("接続エラー")
	}
	token := &response.UserToken{
		AccessToken:  httpClient.AccessToken,
		RefreshToken: httpClient.RefreshToken,
	}
	g.GetUserInfo(httpClient, token)

	log.Println("token", token)
	return token, nil

}

func (g *Google) GetUserInfo(httpClient *oauth2.Token, token *response.UserToken) {

	cxt := context.Background()
	client := g.Config.Client(cxt, httpClient)

	service, _ := v2.NewService(cxt, option.WithHTTPClient(client))

	// 토큰 유효성 검사
	userInfo, _ := service.Tokeninfo().AccessToken(httpClient.AccessToken).Context(cxt).Do()

	token.Email = userInfo.Email
	token.ExpiresIn = userInfo.ExpiresIn
	token.Scope = userInfo.Scope

	//return token, nil
}
