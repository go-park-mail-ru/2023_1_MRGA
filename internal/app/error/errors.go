package error

import "net/http"

var UserNotFound = "cant find user in db"

var ERROR = map[string]struct {
	description string
	status      int
}{
	UserNotFound: {
		description: "you dont register yet",
		status:      http.StatusBadGateway,
	},
}
