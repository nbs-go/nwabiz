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
	Href    string `json:"href"`
}

type SendMessageReq struct {
	To   string    `json:"to"`
	Type string    `json:"type"`
	HSM  HSMObject `json:"hsm"`
}

type HSMObject struct {
	Namespace         string             `json:"namespace"`
	ElementName       string             `json:"element_name"`
	Language          Language           `json:"language"`
	LocalizableParams []LocalizableParam `json:"localizable_params"`
}

type Language struct {
	Policy string `json:"policy"`
	Code   string `json:"code"`
}

type LocalizableParam struct {
	Default  string    `json:"default,omitempty"`
	Currency *Currency `json:"currency,omitempty"`
	DateTime string    `json:"date_time,omitempty"`
}

type Currency struct {
	CurrencyCode string `json:"currency_code"`
	Amount1000   int64  `json:"amount_1000"`
}

type DateTimeReq struct {
	Component DateTimeComponent `json:"component"`
}

type DateTimeComponent struct {
	DayOfWeek  int64 `json:"day_of_week"`
	DayOfMonth int64 `json:"day_of_month"`
	Year       int64 `json:"year"`
	Month      int64 `json:"month"`
	Hour       int64 `json:"hour"`
	Minute     int64 `json:"minute"`
}

type SendMessageResp struct {
	Messages []MessageResp `json:"messages"`
	Meta     MetaResp      `json:"meta"`
	Errors   []ErrorResp   `json:"errors"`
}

type MessageResp struct {
	Id string `json:"id"`
}
