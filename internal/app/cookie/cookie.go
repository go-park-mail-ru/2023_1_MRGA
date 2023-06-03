package cookie

import (
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, nameCookie string, value string, expTime time.Duration) {
	http.SetCookie(w, &http.Cookie{
		Name:     nameCookie,
		Value:    value,
		Expires:  time.Now().Add(expTime),
		Path:     "/",
		// Secure:   true,
		// SameSite: http.SameSiteNoneMode,
	})
}

func GetValueCookie(r *http.Request, nameCookie string) (string, error) {
	valueCookie, err := r.Cookie(nameCookie)
	if err != nil {
		return "", err
	}
	return valueCookie.Value, err
}
