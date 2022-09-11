package tweet

type AddTweetReq struct {
	Text string `json:"text" validate:"required"`
}
