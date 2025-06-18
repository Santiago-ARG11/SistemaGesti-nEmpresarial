
package handlers

import (
    "html/template"
    "net/http"
    "ecommerce/utils"
)

// Muestra formulario y maneja lógica de inicio de sesión
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/login.html")

    switch r.Method {
    case "GET":
        tmpl.Execute(w, nil)
    case "POST":
        email := r.FormValue("email")
        password := r.FormValue("password")

        user, err := utils.LoginUser("data/users.json", email, password)
        if err != nil {
            tmpl.Execute(w, map[string]string{"Error": "Correo o contraseña incorrectos"})
            return
        }

        utils.CreateSession(w, user.Email)
        http.Redirect(w, r, "/catalogo", http.StatusSeeOther)
    default:
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
    }
}

// Muestra formulario y maneja lógica de registro
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/register.html")

    switch r.Method {
    case "GET":
        tmpl.Execute(w, nil)
    case "POST":
        name := r.FormValue("name")
        email := r.FormValue("email")
        password := r.FormValue("password")

        if name == "" || email == "" || password == "" {
            tmpl.Execute(w, map[string]string{"Error": "Todos los campos son obligatorios"})
            return
        }

        err := utils.RegisterUser("data/users.json", name, email, password)
        if err != nil {
            tmpl.Execute(w, map[string]string{"Error": err.Error()})
        } else {
            tmpl.Execute(w, map[string]string{"Success": "Usuario registrado correctamente. Ahora puedes iniciar sesión."})
        }
    default:
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
    }
}

// Cierra la sesión eliminando la cookie
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    utils.DestroySession(w)
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}
