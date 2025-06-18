package utils

import (
	"encoding/json"
	"errors"

	//"fmt"
	"ecommerce/models"
	"os"
	"regexp"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Verifica si el correo tiene un formato válido
func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// Carga los usuarios desde un archivo JSON
func LoadUsersFromFile(path string) ([]models.User, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := json.Unmarshal(file, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// Guarda los usuarios en un archivo JSON
func SaveUsersToFile(path string, users []models.User) error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// Registra un nuevo usuario si el correo no está duplicado
func RegisterUser(path string, name, email, password string) error {
	if !IsValidEmail(email) {
		return errors.New("correo no tiene un formato válido")
	}

	users, err := LoadUsersFromFile(path)
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Email == email {
			return errors.New("el correo ya está registrado")
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := models.User{
		ID:       uuid.New().String(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     "cliente",
	}

	users = append(users, newUser)
	return SaveUsersToFile(path, users)
}

// Verifica las credenciales de un usuario
func LoginUser(path string, email, password string) (*models.User, error) {
	users, err := LoadUsersFromFile(path)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Email == email {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err == nil {
				return &user, nil
			}
		}
	}

	return nil, errors.New("correo o contraseña incorrectos")
}
