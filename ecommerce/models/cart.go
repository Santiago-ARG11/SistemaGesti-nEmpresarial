
package models

// CartItem representa un producto dentro del carrito de compras
type CartItem struct {
    ProductID string  `json:"product_id"`
    Name      string  `json:"name"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
    Subtotal  float64 `json:"subtotal"`
}

// Cart representa el carrito de compras completo
type Cart struct {
    UserID string     `json:"user_id"`
    Items  []CartItem `json:"items"`
    Total  float64    `json:"total"`
}
