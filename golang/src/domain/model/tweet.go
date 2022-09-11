package model

type TweetId uint

type Tweet struct {
	ModelAt

	ID     TweetId `gorm:"primaryKey" json:"ID" binding:"-"`
	UserId UserId  `json:"userID"`
	Text   string  `json:"text"`

	User  *User  `json:"user" gorm:"foreignkey:UserId"`
	Reply *Reply `json:"reply"`
	Like  *Like  `json:"like"`
}
