package models

import "gorm.io/gorm"

//	type RoleModel struct {
//		gorm.Model
//		Name string `gorm:"type:text" json:"name"`
//		Slug string `gorm:"type:text" json:"slug"`
//	}
type Role struct {
	gorm.Model
	Name  string // The name of the role
	Users []User `gorm:"many2many:authority_user_roles;"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Role) TableName() string {
	return "roles"
}

// type UserRole struct {
// 	ID     uint // Unique id (it gets set automatically by the database)
// 	UserID uint // The user id
// 	RoleID uint // The role id
// }
