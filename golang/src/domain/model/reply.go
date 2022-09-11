package model

type ReplyId uint

type Reply struct {
	ModelAt

	ID      ReplyId `gorm:"primaryKey" json:"ID" binding:"-"`
	TweetId TweetId `json:"tweetID"`
	UserId  UserId  `json:"userID"`
	Text    string  `json:"text"`

	Tweet *Tweet `json:"tweet" gorm:"foreignkey:TweetId"`
}
