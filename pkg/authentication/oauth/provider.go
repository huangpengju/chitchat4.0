/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-10-16 08:30:33
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-29 17:08:22
 * @FilePath: \chitchat4.0\pkg\authentication\oauth\provider.go
 * @Description:
 *
 * Copyright (c) 2023 by huangpengju, All Rights Reserved.
 */
package oauth

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"chitchat4.0/pkg/config"
	"chitchat4.0/pkg/model"
	"golang.org/x/oauth2"
)

const (
	GithubAuthType = "github" // github授权类型
	WeChatAuthType = "wechat" // 微信授权类型
	EmptyAuthType  = "nil"    // 空的授权类型
)

var (
	defaultHttpClient = &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		},
		Timeout: 10 * time.Second,
	}
)

/**
 * @description: IsEmptyAuthType 检查授权类型是不是空
 * @return authType 是空时，返回true
 */
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

func (ui *UserInfo) User() *model.User {
	return &model.User{
		Name:   ui.Username,
		Email:  ui.Email,
		Avatar: ui.AvatarUrl,
		AuthInfos: []model.AuthInfo{
			{
				AuthType: ui.AuthType,
				AuthId:   ui.ID,
				Url:      ui.Url,
			},
		},
	}
}

// 授权管理
type OAuthManager struct {
	conf map[string]config.OAuthConfig // 授权配置（授权类型、客户Id、客户秘密）
}

// 创建新的授权管理
func NewOAuthManager(conf map[string]config.OAuthConfig) *OAuthManager {
	return &OAuthManager{
		conf: conf,
	}
}

/**
 * @description: GetAuthProvider 获取授权提供者
 * @param {string} authType 授权类型
 * @return 授权提供者
 */
func (m *OAuthManager) GetAuthProvider(authType string) (AuthProvider, error) {
	var provider AuthProvider    // 接口类型
	conf, ok := m.conf[authType] // 检查授权类型是否存在，存在返回结构体 OAuthConfig
	if !ok {
		return nil, fmt.Errorf("auth type %s not found in config", authType) // 在配置中找不到验证类型
	}
	switch authType {
	case GithubAuthType:
		provider = NewGithubAuth(conf.ClientId, conf.ClientSecret) // github
	case WeChatAuthType:
		provider = NewWeChatAuth(conf.ClientId, conf.ClientSecret) // wechat
	default:
		return nil, fmt.Errorf("unknown auth type: %s", authType)
	}

	return provider, nil
}

// AuthProvider 授权提供者
type AuthProvider interface {
	GetToken(code string) (*oauth2.Token, error)
	GetUserInfo(token *oauth2.Token) (*UserInfo, error)
}
