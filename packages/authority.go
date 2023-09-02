package packages

import (
	"rest_api/config"

	"github.com/pooriaghaedi/authority"
)

var Rbac *authority.Authority

func InitAuthority() {
	Rbac = authority.New(authority.Options{
		TablesPrefix: "authority_",
		DB:           config.Database,
	})

}

var Role *authority.Role
