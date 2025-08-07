package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
)

func TestJSONResponse(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		data := struct {
			Value string `json:"value"`
		}{Value: "ok"}
		return JSONResponse(c, http.StatusOK, true, "message", data, nil)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	var body struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    struct {
			Value string `json:"value"`
		} `json:"data"`
		Error *dto.ErrorDTO `json:"error"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if !body.Success || body.Message != "message" || body.Data.Value != "ok" || body.Error != nil {
		t.Fatalf("unexpected body: %+v", body)
	}
}

func TestJSONError(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		return JSONError(c, http.StatusBadRequest, "BAD_CODE", "bad request", "detail")
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
	var body struct {
		Success bool          `json:"success"`
		Message string        `json:"message"`
		Data    any           `json:"data"`
		Error   *dto.ErrorDTO `json:"error"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if body.Success || body.Message != "bad request" || body.Data != nil || body.Error == nil || body.Error.Code != "BAD_CODE" || body.Error.Details != "detail" {
		t.Fatalf("unexpected body: %+v", body)
	}
}
