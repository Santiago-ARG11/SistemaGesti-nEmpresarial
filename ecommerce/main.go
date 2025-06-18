
package main

import (
    "fmt"
    "net/http"
    "ecommerce/handlers"
)

func main() {
    // Configuración de rutas
    http.HandleFunc("/login", handlers.LoginHandler)
    http.HandleFunc("/register", handlers.RegisterHandler)
    http.HandleFunc("/catalogo", handlers.CatalogHandler)
    http.HandleFunc("/logout", handlers.LogoutHandler)

    fmt.Println("Servidor corriendo en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
