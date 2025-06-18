
package utils

import (
    "net/http"
    "time"
)

const SessionCookieName = "session_user"

// Crea una cookie de sesión con el email del usuario
func CreateSession(w http.ResponseWriter, email string) {
    cookie := &http.Cookie{
        Name:     SessionCookieName,
        Value:    email,
        Path:     "/",
        HttpOnly: true,
        Expires:  time.Now().Add(1 * time.Hour),
    }
    http.SetCookie(w, cookie)
}

// Obtiene el email del usuario desde la cookie de sesión
func GetSessionEmail(r *http.Request) (string, error) {
    cookie, err := r.Cookie(SessionCookieName)
    if err != nil {
        return "", err
    }
    return cookie.Value, nil
}

// Elimina la cookie de sesión
func DestroySession(w http.ResponseWriter) {
    expired := &http.Cookie{
        Name:   SessionCookieName,
        Value:  "",
        Path:   "/",
        MaxAge: -1,
    }
    http.SetCookie(w, expired)
}
