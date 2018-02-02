package slackposter

type Attachment struct {
	Fallback   string             `json:"fallback"`
	Color      string             `json:"color"`
	Pretext    string             `json:"pretext"`
	AuthorName string             `json:"author_name"`
	AuthorLink string             `json:"author_link"`
	AuthorIcon string             `json:"author_icon"`
	Title      string             `json:"title"`
	TitleLink  string             `json:"title_link"`
	Text       string             `json:"text"`
	Fields     []*AttachmentField `json:"fields"`
	ImageUrl   string             `json:"image_url"`
	ThumbUrl   string             `json:"thumb_url"`
	Footer     string             `json:"footer"`
	FooterIcon string             `json:"footer_icon"`
	Ts         int64              `json:"ts"`
}

type AttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}
