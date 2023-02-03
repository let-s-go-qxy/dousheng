package message

type Message struct {
	Id         int    `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_date,omitempty"`
}
