package structs

type Canvasser struct {
	Id       int64  `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
}
