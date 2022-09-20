package config

import "github.com/mahiro72/go-mysql-ca/utils/helper"

// Portはサーバーのport番号を返します
func Port() string {
	return helper.GetEnvOrDefault("PORT", "8080")
}
