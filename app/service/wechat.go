package service

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

var DefaultWechat *Wechat

type Wechat struct {
	wechat          *wechat.Wechat
	config          *offConfig.Config
	officialAccount *officialaccount.OfficialAccount
}

func NewWechat(appId, appSecret, token, aesKey string) *Wechat {
	cfg := &offConfig.Config{
		AppID:          appId,
		AppSecret:      appSecret,
		Token:          token,
		EncodingAESKey: aesKey,
		Cache:          cache.NewMemory(),
	}

	var (
		wt   = wechat.NewWechat()
		offa = wt.GetOfficialAccount(cfg)
	)
	return &Wechat{
		config:          cfg,
		wechat:          wt,
		officialAccount: offa,
	}
}

func (w *Wechat) DefaultOfficialAccount() *officialaccount.OfficialAccount {
	return w.officialAccount
}
