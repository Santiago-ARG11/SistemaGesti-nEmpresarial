
package handlers

import (
    "html/template"
    "net/http"
    "ecommerce/utils"
)

// Muestra el catálogo si el usuario ha iniciado sesión
func CatalogHandler(w http.ResponseWriter, r *http.Request) {
    email, err := utils.GetSessionEmail(r)
    if err != nil {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    products, err := utils.LoadProductsFromFile("data/products.json")
    if err != nil {
        http.Error(w, "Error cargando productos", http.StatusInternalServerError)
        return
    }

    tmpl, _ := template.ParseFiles("templates/products.html")
    tmpl.Execute(w, struct {
        Email    string
        Products interface{}
    }{
        Email:    email,
        Products: products,
    })
}
