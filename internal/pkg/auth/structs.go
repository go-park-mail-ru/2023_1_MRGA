package auth

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRestTemp struct {
	Name string `json:"name"`
}

type UserRes struct {
	Name   string `json:"name" structs:"name"`
	Age    int    `json:"age" structs:"age"`
	Avatar uint   `json:"avatarId" structs:"avatarId"`
}

type Photo struct {
	PhotoId uint `json:"photoId"`
	Avatar  bool `json:"avatar"`
}
