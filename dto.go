package nwabiz

type LoginResp struct {
	Users []UserResp `json:"users"`
	Meta  MetaResp   `json:"meta"`
}

type UserResp struct {
	Token        string `json:"token"`
	ExpiresAfter string `json:"expires_after"`
}

type MetaResp struct {
	Version   string `json:"version"`
	APIStatus string `json:"api_status"`
}

type CheckContactsReq struct {
	Blocking   string   `json:"blocking"`
	Contacts   []string `json:"contacts"`
	ForceCheck bool     `json:"force_check"`
}

type ContactsResp struct {
	Input      string `json:"input"`
	Status     string `json:"status"`
	WhatsAppId string `json:"wa_id"`
}

type CheckContactsResp struct {
	Contacts []ContactsResp `json:"contacts"`
	Meta     MetaResp       `json:"meta"`
	Errors   []ErrorResp    `json:"errors"`
}

type ErrorResp struct {
	Code    int64  `json:"code"`
	Title   string `json:"title"`
	Details string `json:"details"`
}
