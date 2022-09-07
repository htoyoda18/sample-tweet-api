package model

type UserId uint

type User struct {
	ModelAt

	ID       UserId `json:"ID" binding:"-"`
	Email    string `gorm:"unique_index" json:"email"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Password string `json:"-" gorm:"-" sql:"-"`
}
