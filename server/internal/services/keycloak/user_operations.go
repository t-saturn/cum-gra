package keycloak

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/internal/config"
	"strings"
)

type CreateKeycloakUserInput struct {
	Email     string
	FirstName string
	LastName  string
	DNI       string
	Password  string // Siempre requerido
}

type KeycloakUserResult struct {
	UserID string
}

// Crear usuario usando el token del usuario autenticado
func CreateKeycloakUser(accessToken string, input CreateKeycloakUserInput) (*KeycloakUserResult, error) {
	cfg := config.GetConfig()

	user := map[string]interface{}{
		"username":      input.DNI,
		"email":         input.Email,
		"firstName":     input.FirstName,
		"lastName":      input.LastName,
		"enabled":       true,
		"emailVerified": true, // Email verificado
		"credentials": []map[string]interface{}{
			{
				"type":      "password",
				"value":     input.Password,
				"temporary": false, // NO es temporal
			},
		},
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		return nil, fmt.Errorf("error serializando usuario: %w", err)
	}

	url := fmt.Sprintf("%s/admin/realms/%s/users", cfg.KeycloakSSOURL, cfg.KeycloakRealm)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creando request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error llamando a keycloak: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error creando usuario en keycloak: %d - %s", resp.StatusCode, string(body))
	}

	// Obtener el ID del usuario del header Location
	location := resp.Header.Get("Location")
	if location == "" {
		return nil, fmt.Errorf("keycloak no retornó Location header")
	}

	// Location viene como: .../users/{user-id}
	parts := strings.Split(location, "/")
	userID := parts[len(parts)-1]

	return &KeycloakUserResult{
		UserID: userID,
	}, nil
}

// Verificar si un usuario existe en Keycloak por email o username
func UserExistsInKeycloak(accessToken, email, dni string) (bool, string, error) {
	cfg := config.GetConfig()

	// Buscar por email
	url := fmt.Sprintf("%s/admin/realms/%s/users?email=%s&exact=true",
		cfg.KeycloakSSOURL, cfg.KeycloakRealm, email)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, "", err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var users []map[string]interface{}
	var unmarshalErr error
	if unmarshalErr = json.Unmarshal(body, &users); unmarshalErr != nil {
		return false, "", err
	}

	if len(users) > 0 {
		userID, _ := users[0]["id"].(string)
		return true, userID, nil
	}

	// Buscar por username (DNI)
	url = fmt.Sprintf("%s/admin/realms/%s/users?username=%s&exact=true",
		cfg.KeycloakSSOURL, cfg.KeycloakRealm, dni)

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return false, "", err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err = client.Do(req)
	if err != nil {
		return false, "", err
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &users); err != nil {
		return false, "", err
	}

	if len(users) > 0 {
		userID, _ := users[0]["id"].(string)
		return true, userID, nil
	}

	return false, "", nil
}