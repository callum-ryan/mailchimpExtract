package mailchimpExtract

import "encoding/json"

type Campaign struct {
	Id                string                  `json:"id"`
	WebId             int                     `json:"web_id"`
	ParentCampaignId  string                  `json:"parent_campaign_id"`
	Type              string                  `json:"type"`
	CreateTime        string                  `json:"create_time"`
	ArchiveUrl        string                  `json:"archive_url"`
	LongArchiveUrl    string                  `json:"long_archive_url"`
	Status            string                  `json:"status"`
	EmailsSent        int                     `json:"emails_sent"`
	SendTime          string                  `json:"send_time"`
	ContentType       string                  `json:"content_type"`
	NeedsBlockRefresh bool                    `json:"needs_block_refresh"`
	Recipients        CampaignRecipients      `json:"recipients"`
	Settings          CampaignSettings        `json:"settings"`
	VariateSettings   CampaignVariateSettings `json:"variate_settings"`
	Tracking          CampaignTracking        `json:"tracking"`
	RssOpts           CampaignRSSOpts         `json:"rss_opts"`
	AbSplitOpts       CampaignABSplitOptions  `json:"ab_split_opts"`
	SocialCard        CampaignSocialCard      `json:"social_card"`
	ReportSummary     CampaignReportSummary   `json:"report_summary"`
	DeliveryStatus    CampaignDelivery        `json:"delivery_status"`
	Links             interface{}             `json:"_links"`
}

type CampaignRecipients struct {
	ListId         string              `json:"list_id"`
	ListIsActive   bool                `json:"list_is_active"`
	ListName       string              `json:"list_name"`
	SegmentText    string              `json:"segment_text"`
	RecipientCount int                 `json:"recipient_count"`
	SegmentOpts    CampaignSegmentOpts `json:"segment_opts"`
}

type CampaignSegmentOpts struct {
	SavedSegmentId    int         `json:"saved_segment_id"`
	PrebuiltSegmentId string      `json:"prebuilt_segment_id"`
	Match             string      `json:"match"`
	Conditions        interface{} `json:"conditions"`
}

type CampaignSettings struct {
	SubjectLine     string      `json:"subject_line"`
	PreviewText     string      `json:"preview_text"`
	Title           string      `json:"title"`
	FromName        string      `json:"from_name"`
	ReplyTo         string      `json:"reply_to"`
	UseConversation bool        `json:"use_conversation"`
	ToName          string      `json:"to_name"`
	FolderId        string      `json:"folder_id"`
	Authenticate    bool        `json:"authenticate"`
	AutoFooter      bool        `json:"auto_footer"`
	InlineCss       bool        `json:"inline_css"`
	AutoTweet       bool        `json:"auto_tweet"`
	AutoFbPost      interface{} `json:"auto_fb_post"`
	FbComments      bool        `json:"fb_comments"`
	Timewarp        bool        `json:"timewarp"`
	TemplateId      int         `json:"template_id"`
	DragAndDrop     bool        `json:"drag_and_drop"`
}

type CampaignVariateSettings struct {
	WinningCombinationId string                      `json:"winning_combination_id"`
	WinningCampaignId    string                      `json:"winning_campaign_id"`
	WinnerCriteria       string                      `json:"winner_criteria"`
	WaitTime             int                         `json:"wait_time"`
	TestSize             int                         `json:"test_size"`
	SubjectLines         interface{}                 `json:"subject_lines"`
	SendTimes            interface{}                 `json:"send_times"`
	FromNames            interface{}                 `json:"from_names"`
	ReplyToAddresses     interface{}                 `json:"reply_to_addresses"`
	Contents             interface{}                 `json:"contents"`
	Combinations         CampaignVariateCombinations `json:"combinations"`
}

type CampaignVariateCombinations struct {
	Id                 string `json:"id"`
	SubjectLine        int    `json:"subject_line"`
	SendTime           int    `json:"send_time"`
	FromName           int    `json:"from_name"`
	ReplyTo            int    `json:"reply_to"`
	ContentDescription int    `json:"content_description"`
	Recipients         int    `json:"recipients"`
}

type CampaignTracking struct {
	Opens           bool               `json:"opens"`
	HtmlClicks      bool               `json:"html_clicks"`
	TextClicks      bool               `json:"text_clicks"`
	GoalTracking    bool               `json:"goal_tracking"`
	Ecomm360        bool               `json:"ecomm360"`
	GoogleAnalytics string             `json:"google_analytics"`
	Clicktale       string             `json:"clicktale"`
	Salesforce      SalesForceTracking `json:"salesforce"`
	Capsule         CapsuleTracking    `json:"capsule"`
}

type SalesForceTracking struct {
	Campaign bool `json:"campaign"`
	Notes    bool `json:"notes"`
}

type CapsuleTracking struct {
	Notes bool `json:"notes"`
}

type CampaignRSSOpts struct {
	FeedUrl         string              `json:"feed_url"`
	Frequency       string              `json:"frequency"`
	Schedule        CampaignRSSSchedule `json:"schedule"`
	LastSent        string              `json:"last_sent"`
	ConstrainRssImg bool                `json:"constrain_rss_img"`
}

type CampaignRSSSchedule struct {
	Hour            int                  `json:"hour"`
	DailySend       CampaignRSSDailySend `json:"daily_send"`
	WeeklySendDay   string               `json:"weekly_send_day"`
	MonthlySendDate float32              `json:"monthly_send_date"`
}

type CampaignRSSDailySend struct {
	Sunday    bool `json:"sunday"`
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
	Saturday  bool `json:"saturday"`
}

type CampaignABSplitOptions struct {
	SplitTest      string `json:"split_test"`
	PickWinner     string `json:"pick_winner"`
	WaitUnits      string `json:"wait_units"`
	WaitTime       int    `json:"wait_time"`
	SplitSize      int    `json:"split_size"`
	FromNameA      string `json:"from_name_a"`
	FromNameB      string `json:"from_name_b"`
	ReplyEmailA    string `json:"reply_email_a"`
	ReplyEmailB    string `json:"reply_email_b"`
	SubjectA       string `json:"subject_a"`
	SubjectB       string `json:"subject_b"`
	SendTimeA      string `json:"send_time_a"`
	SendTimeB      string `json:"send_time_b"`
	SendTimeWinner string `json:"send_time_winner"`
}

type CampaignSocialCard struct {
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

type CampaignReportSummary struct {
	Opens            int                      `json:"opens"`
	UniqueOpens      int                      `json:"unique_opens"`
	OpenRate         float32                  `json:"open_rate"`
	Clicks           int                      `json:"clicks"`
	SubscriberClicks int                      `json:"subscriber_clicks"`
	ClickRate        float32                  `json:"click_rate"`
	Ecommerce        CampaignEcommerceSummary `json:"ecommerce"`
}

type CampaignEcommerceSummary struct {
	TotalOrders  int     `json:"total_orders"`
	TotalSpent   float32 `json:"total_spent"`
	TotalRevenue float32 `json:"total_revenue"`
}

type CampaignDelivery struct {
	Enabled        bool   `json:"enabled"`
	CanCancel      bool   `json:"can_cancel"`
	Status         string `json:"status"`
	EmailsSent     int    `json:"emails_sent"`
	EmailsCanceled int    `json:"emails_canceled"`
}

type CampaignsResponse struct {
	Campaigns  []Campaign  `json:"campaigns"`
	TotalItems int         `json:"total_items"`
	Links      interface{} `json:"_links"`
}

func GetAllCampaigns(apiRoot, apiKey string) (campaigns CampaignsResponse) {
	endPoint := "campaigns?count=2000"
	response := MakeReq(apiKey, apiRoot, endPoint)
	json.Unmarshal(response, &campaigns)
	return
}

func GetCampaign(campaignId, apiRoot, apiKey string) (campaign Campaign) {
	endPoint := "campaigns/" + campaignId
	report := MakeReq(apiKey, apiRoot, endPoint)
	json.Unmarshal(report, &campaign)
	return
}
