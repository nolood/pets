package sso

type ValidateData struct {
	Token    string
	User     string
	AuthDate int64
	QueryId  string
	Hash     string
}
