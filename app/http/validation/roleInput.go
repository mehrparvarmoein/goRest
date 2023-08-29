package validation

type RoleInput struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug" binding:"required"`
}
