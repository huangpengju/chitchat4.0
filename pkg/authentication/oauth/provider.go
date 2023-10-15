package oauth

import (
	"chitchat4.0/pkg/config"
	"golang.org/x/oauth2"
)

const (
	EmptyAuthType = "nil" // 空的授权类型
)

func IsEmptyAuthType(authType string) bool {
	return authType == "" || authType == EmptyAuthType
}

type UserInfo struct {
	ID          string
	Url         string
	AuthType    string
	Username    string
	DisplayName string
	Email       string
	AvatarUrl   string
}

type OAuthManger struct {
	conf map[string]config.OAuthConfig
}

// AuthProvider 授权提供者
type AuthProvider interface {
	GetToken(code string) (*oauth2.Token, error)
	GetUserInfo(token *oauth2.Token) (*UserInfo, error)
}
