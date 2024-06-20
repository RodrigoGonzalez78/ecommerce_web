package utils

import (
	"net/http"
	"strconv"
	"time"
)

var id_user int

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{
		Name:     "id_user",             // Nombre de la cookie
		Value:    strconv.Itoa(id_user), // Valor de la cookie convertido a string
		Expires:  expiration,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	w.Write([]byte("Cookie has been set!"))
}

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		if err == http.ErrNoCookie {
			w.Write([]byte("No cookie found"))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		w.Write([]byte("Cookie value: " + cookie.Value))
	}
}
