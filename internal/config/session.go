package config

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/spf13/viper"
)

func NewCookieAuth(viper *viper.Viper) *common.CookieAuthConfig {
	maxAge := viper.GetInt("COOKIE_AUTH_AGE")
	secretKey := viper.GetString("COOKIE_AUTH_SECRET")
	return &common.CookieAuthConfig{
		MaxAge:     (60 * maxAge),
		SecretKey:  secretKey,
		CookieName: "authentication-session",
	}
}
