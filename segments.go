package mailchimpExtract

type SegmentsReponse struct {
	Segments   []Segment   `json:"segments"`
	ListId     string      `json:"list_id"`
	TotalItems int         `json:"total_items"`
	Links      interface{} `json:"_links"`
}

type Segment struct {
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	MemberCount int            `json:"member_count"`
	Type        string         `json:"type"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
	Options     SegmentOptions `json:"options"`
	ListId      string         `json:"list_id"`
	Links       interface{}    `json:"_links"`
}

type SegmentOptions struct {
	Match      string      `json:"match"`
	Conditions interface{} `json:"conditions"`
}
