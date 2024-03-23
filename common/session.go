package common

type CookieAuthConfig struct {
	MaxAge     int
	SecretKey  string
	CookieName string
}
