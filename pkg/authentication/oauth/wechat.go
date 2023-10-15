package oauth

import (
	"net/http"

	"golang.org/x/oauth2"
)

type WeChatAuth struct {
	Client *http.Client
	Config *oauth2.Config
}

func (auth *WeChatAuth) GetToken(code string) (*oauth2.Token, error) {

}
