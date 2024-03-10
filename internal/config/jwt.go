package config

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/spf13/viper"
)

func NewJWT(viper *viper.Viper) *common.JWT {
	issuer := viper.GetString("JWT_ISSUER")
	signatureKey := viper.GetString("JWT_SIGNATURE_KEY")
	return &common.JWT{
		Issuer:       issuer,
		SignatureKey: signatureKey,
	}
}
