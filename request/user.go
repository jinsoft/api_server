package request

type Login struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}