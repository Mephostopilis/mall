package cookie

import (
	"net/http"
)

func SetJwt(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     mw.CookieName,
		Value:    tokenString,
		Path:     "/",
		Domain:   mw.CookieDomain,
		MaxAge:   maxage,
		Secure:   mw.SecureCookie,
		HttpOnly: mw.CookieHTTPOnly,
	}
	http.SetCookie(w, &cookie)
}
