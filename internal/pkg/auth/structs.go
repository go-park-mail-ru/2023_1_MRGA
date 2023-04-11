package auth

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRestTemp struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserRes struct {
	Email  string  `json:"email" structs:"email"`
	Name   string  `json:"name" structs:"name"`
	Age    int     `json:"age"`
	Photos []Photo `json:"photos" structs:"photos"`
}

type Photo struct {
	PhotoId uint `json:"photoId"`
	Avatar  bool `json:"avatar"`
}
