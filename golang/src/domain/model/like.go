package model

type LikeId uint

type Like struct {
	ModelAt

	ID      LikeId  `gorm:"primaryKey" json:"ID" binding:"-"`
	TweetId TweetId `json:"tweetID"`
	UserId  UserId  `json:"userID"`

	Tweet *Tweet `json:"tweet" gorm:"foreignkey:TweetId"`
	User  *User  `json:"User" gorm:"foreignkey:UserId"`
}
