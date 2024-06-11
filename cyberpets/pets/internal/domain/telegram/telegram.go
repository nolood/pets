package telegram

type WebAppData struct {
	AuthDate int64  `json:"auth_date"`
	QueryID  string `json:"query_id"`
	Hash     string `json:"hash"`
	User     string `json:"user"`
	Token    string
}
