package packages

import (
	"rest_api/config"

	"github.com/harranali/authority"
)

var Auth *authority.Authority

func InitAuthority() {
	Auth = authority.New(authority.Options{
		TablesPrefix: "authority_",
		DB:           config.Database,
	})
}
