package mailchimpExtract

import (
	"encoding/json"
)

// AutomationResponse is the struct returned when calling automations/automation_id
type AutomationReponse struct {
	Automations []Automation `json:"automations"`
	TotalItems  int          `json:"total_items"`
	Links       interface{}  `json:"_links"`
}
type Automation struct {
	Id              string               `json:"id"`
	CreateTime      string               `json:"create_time"`
	StartTime       string               `json:"start_time"`
	Status          string               `json:"status"`
	EmailsSent      int                  `json:"emails_sent"`
	Recipients      AutomationRecipients `json:"recipients"`
	Settings        AutomationSettings   `json:"settings"`
	Tracking        AutomationTracking   `json:"tracking"`
	TriggerSettings interface{}          `json:"trigger_settings"`
	ReportSummary   ReportSummary        `json:"report_summary"`
	Links           interface{}          `json:"_links"`
}

type AutomationRecipients struct {
	ListId       string     `json:"list_id"`
	ListIsActive bool       `json:"list_is_active"`
	ListName     string     `json:"list_name"`
	SegmentOpts  SegmentOps `json:"segment_opts"`
	StoreId      string     `json:"store_id"`
}

type SegmentOps struct {
	SavedSegmentId int         `json:"saved_segment_id"`
	Match          string      `json:"match"`
	Conditions     interface{} `json:"conditions"`
}

type AutomationSettings struct {
	Title           string `json:"title"`
	FromName        string `json:"from_name"`
	ReplyTo         string `json:"reply_to"`
	UseConversation bool   `json:"use_conversation"`
	ToName          string `json:"to_name"`
	Authenticate    bool   `json:"authenticate"`
	AutoFooter      bool   `json:"auto_footer"`
	InlineCSS       bool   `json:"inline_css"`
}

type AutomationTracking struct {
	Opens           bool        `json:"opens"`
	HtmlClicks      bool        `json:"html_clicks"`
	TextClicks      bool        `json:"text_clicks"`
	GoalTracking    bool        `json:"goal_tracking"`
	Ecomm360        bool        `json:"ecomm360"`
	GoogleAnalytics string      `json:"google_analytics"`
	Clicktale       string      `json:"clicktale"`
	Salesforce      interface{} `json:"salesforce"`
	Capsule         interface{} `json:"capsule"`
}

type ReportSummary struct {
	Opens            int     `json:"opens"`
	UniqueOpens      int     `json:"unique_opens"`
	OpenRate         float32 `json:"open_rate"`
	Clicks           int     `json:"clicks"`
	SubscriberClicks int     `json:"subscriber_clicks"`
	ClickRate        float32 `json:"click_rate"`
}

// AutomationEmailResponse is the struct returned when calling automations/automation_id/emails
type AutomationEmailResponse struct {
	Emails     []AutomationEmail `json:"emails"`
	TotalItems int               `json:"total_items"`
	Links      interface{}       `json:"_links"`
}

type AutomationEmail struct {
	Id                string                    `json:"id"`
	WebId             int                       `json:"web_id"`
	WorkflowId        string                    `json:"workflow_id"`
	Position          int                       `json:"position"`
	Delay             AutomationDelay           `json:"delay"`
	CreateTime        string                    `json:"create_time"`
	StartTime         string                    `json:"start_time"`
	ArchiveUrl        string                    `json:"archive_url"`
	Status            string                    `json:"status"`
	EmailsSent        int                       `json:"emails_sent"`
	SendTime          string                    `json:"send_time"`
	ContentType       string                    `json:"content_type"`
	NeedsBlockRefresh bool                      `json:"needs_block_refresh"`
	Recipients        AutomationEmailRecipients `json:"recipients"`
	Settings          AutomationEmailSettings   `json:"settings"`
	Tracking          AutomationTracking        `json:"tracking"`
	SocialCard        interface{}               `json:"social_card"`
	TriggerSettings   interface{}               `json:"trigger_settings"`
	ReportSummary     ReportSummary             `json:"report_summary"`
	Links             interface{}               `json:"_links"`
}

type AutomationDelay struct {
	Amount            int    `json:"amount"`
	Type              string `json:"type"`
	Direction         string `json:"direction"`
	Action            string `json:"action"`
	ActionDescription string `json:"action_description"`
	FullDescription   string `json:"full_description"`
}

type AutomationEmailRecipients struct {
	ListId       string      `json:"list_id"`
	ListIsActive bool        `json:"list_is_active"`
	SegmentOpts  interface{} `json:"segment_opts"`
}

type AutomationEmailSettings struct {
	SubjectLine  string      `json:"subject_line"`
	PreviewText  string      `json:"preview_text"`
	Title        string      `json:"title"`
	FromName     string      `json:"from_name"`
	ReplyTo      string      `json:"reply_to"`
	Authenticate bool        `json:"authenticate"`
	AutoFooter   bool        `json:"auto_footer"`
	InlineCss    bool        `json:"inline_css"`
	AutoTweet    bool        `json:"auto_tweet"`
	AutoFbPost   interface{} `json:"auto_fb_post"`
	FbComments   bool        `json:"fb_comments"`
	TemplateId   int         `json:"template_id"`
	DragAndDrop  bool        `json:"drag_and_drop"`
}

func GetAllAutomations(apiRoot, apiKey string) (automations AutomationReponse) {
	endPoint := "automations?count=2000"
	response := MakeReq(apiKey, apiRoot, endPoint)
	json.Unmarshal(response, &automations)
	return
}

func GetAutomation(automationId, apiRoot, apiKey string) (automation Automation) {
	endPoint := "automations/" + automationId
	report := MakeReq(apiKey, apiRoot, endPoint)
	json.Unmarshal(report, &automation)
	return automation
}

func GetAutomationEmails(automationId, apiRoot, apiKey string) (aReport AutomationEmailResponse) {
	endPoint := "automations/" + automationId + "/emails?count=2000"
	report := MakeReq(apiKey, apiRoot, endPoint)
	json.Unmarshal(report, &aReport)
	return aReport
}
