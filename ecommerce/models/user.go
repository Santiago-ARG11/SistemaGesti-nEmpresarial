
package models

// User representa un usuario registrado en el sistema
type User struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Role     string `json:"role"` // Puede ser "cliente" o "admin"
}
