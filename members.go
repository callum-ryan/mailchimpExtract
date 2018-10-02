package mailchimpExtract

import "encoding/json"

type MembersResponse struct {
	Members    []Member    `json:"members"`
	ListId     string      `json:"list_id"`
	TotalItems int         `json:"total_items"`
	Links      interface{} `json:"_links"`
}

type Member struct {
	Id                   string               `json:"id"`
	EmailAddress         string               `json:"email_address"`
	UniqueEmailId        string               `json:"unique_email_id"`
	EmailType            string               `json:"email_type"`
	Status               string               `json:"status"`
	UnsubscribeReason    string               `json:"unsubscribe_reason"`
	MergeFields          interface{}          `json:"merge_fields"`
	Interests            interface{}          `json:"interests"`
	Stats                Stats                `json:"stats"`
	IpSignup             string               `json:"ip_signup"`
	TimestampSignup      string               `json:"timestamp_signup"`
	IpOpt                string               `json:"ip_opt"`
	TimestampOpt         string               `json:"timestamp_opt"`
	MemberRating         int                  `json:"member_rating"`
	LastChanged          string               `json:"last_changed"`
	Language             string               `json:"language"`
	Vip                  bool                 `json:"vip"`
	EmailClient          string               `json:"email_client"`
	Location             Location             `json:"location"`
	MarketingPermissions MarketingPermissions `json:"marketing_permissions"`
	LastNote             Note                 `json:"last_note"`
	TagsCount            int                  `json:"tags_count"`
	Tags                 interface{}          `json:"tags"`
	ListId               string               `json:"list_id"`
	Links                interface{}          `json:"_links"`
}

type Note struct {
	NoteId    int    `json:"note_id"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	Note      string `json:"note"`
}

type MarketingPermissions struct {
	MarketingPermissionId string `json:"marketing_permission_id"`
	Text                  string `json:"text"`
	Enabled               bool   `json:"enabled"`
}

type Location struct {
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	GmtOff      int     `json:"gmtoff"`
	DstOff      int     `json:"dstoff"`
	CountryCode string  `json:"country_code"`
	Timezone    string  `json:"timezone"`
}

type Stats struct {
	AvgOpenRate  float32 `json:"avg_open_rate"`
	AvgClickRate float32 `json:"avg_click_rate"`
}

type ActivityResponse struct {
	Activity   []Activity  `json:"activity"`
	EmailId    string      `json:"email_id"`
	ListId     string      `json:"list_id"`
	TotalItems int         `json:"total_items"`
	Links      interface{} `json:"_links"`
}

type Activity struct {
	Action         string `json:"action"`
	Timestamp      string `json:"timestamp"`
	Url            string `json:"url"`
	Type           string `json:"type"`
	CampaignId     string `json:"campaign_id"`
	Title          string `json:"title"`
	ParentCampaign string `json:"parent_campaign"`
}


func GetMemberActivity(apiRoot, apiKey, listId, subscriberHash string) (activities ActivityResponse){
	endPoint := "lists/" + listId + "/members/" + subscriberHash + "/activity"
	report := MakeReq(apiKey, apiRoot, endPoint)
	json.Unmarshal(report, &activities)
	return
}

func GetMembers(apiRoot, apiKey, listId string) (members MembersResponse) {
	//GET /lists/{list_id}/members
	endPoint := "lists/" + listId + "/members"
	report := MakeReq(apiKey, apiRoot, endPoint)
	json.Unmarshal(report, &members)
	return
}