package mailchimpExtract

import "encoding/json"

type ListsResponse struct {
	Lists      []List      `json:"lists"`
	TotalItems int         `json:"total_items"`
	Links      interface{} `json:"_links"`
}

type List struct {
	Id                   string           `json:"id"`
	WebId                int              `json:"web_id"`
	Name                 string           `json:"name"`
	Contact              interface{}      `json:"contact"`
	PermissionReminder   string           `json:"permission_reminder"`
	UseArchiveBar        bool             `json:"use_archive_bar"`
	CampaignDefaults     CampaignDefaults `json:"campaign_defaults"`
	NotifyOnSubscribe    string           `json:"notify_on_subscribe"`
	NotifyOnUnsubscribe  string           `json:"notify_on_unsubscribe"`
	DateCreated          string           `json:"date_created"`
	ListRating           int              `json:"list_rating"`
	EmailTypeOption      bool             `json:"email_type_option"`
	SubscribeUrlShort    string           `json:"subscribe_url_short"`
	SubscribeUrlLong     string           `json:"subscribe_url_long"`
	BeamerAddress        string           `json:"beamer_address"`
	Visibility           string           `json:"visibility"`
	DoubleOptin          bool             `json:"double_optin"`
	MarketingPermissions bool             `json:"marketing_permissions"`
	Modules              interface{}      `json:"modules"`
	Stats                ListStats        `json:"stats"`
	Links                interface{}      `json:"_links"`
}

type ListStats struct {
	MemberCount               int     `json:"member_count"`
	UnsubscribeCount          int     `json:"unsubscribe_count"`
	CleanedCount              int     `json:"cleaned_count"`
	MemberCountSinceSend      int     `json:"member_count_since_send"`
	UnsubscribeCountSinceSend int     `json:"unsubscribe_count_since_send"`
	CleanedCountSinceSend     int     `json:"cleaned_count_since_send"`
	CampaignCount             int     `json:"campaign_count"`
	CampaignLastSent          string  `json:"campaign_last_sent"`
	MergeFieldCount           int     `json:"merge_field_count"`
	AvgSubRate                float32 `json:"avg_sub_rate"`
	AvgUnsubRate              float32 `json:"avg_unsub_rate"`
	TargetSubRate             float32 `json:"target_sub_rate"`
	OpenRate                  float32 `json:"open_rate"`
	ClickRate                 float32 `json:"click_rate"`
	LastSubDate               string  `json:"last_sub_date"`
	LastUnsubDate             string  `json:"last_unsub_date"`
}

type CampaignDefaults struct {
	FromName  string `json:"from_name"`
	FromEmail string `json:"from_email"`
	Subject   string `json:"subject"`
	Language  string `json:"language"`
}

type ContactInfo struct {
	Company  string `json:"company"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
	Country  string `json:"country"`
	Phone    string `json:"phone"`
}


func getLists(apiKey, apiRoot string) (lists ListsResponse) {
	endPoint := "lists"
	response := makeReq(apiKey, apiRoot, endPoint)
	json.Unmarshal(response, &lists)
	return
}

func getList(apiKey, apiRoot, listId string) (list List) {
	endPoint := "lists/" + listId
	response := makeReq(apiKey, apiRoot, endPoint)
	json.Unmarshal(response, &list)
	return
}