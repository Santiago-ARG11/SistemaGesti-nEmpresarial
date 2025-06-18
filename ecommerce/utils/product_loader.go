
package utils

import (
    "encoding/json"
    "os"
    "ecommerce/models"
)

// Lee productos desde un archivo JSON y los convierte en structs
func LoadProductsFromFile(path string) ([]models.Product, error) {
    file, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    var products []models.Product
    err = json.Unmarshal(file, &products)
    return products, err
}
