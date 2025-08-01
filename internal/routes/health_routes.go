package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/mssola/user_agent"
	"github.com/t-saturn/auth-service-server/internal/handlers"
	"github.com/t-saturn/auth-service-server/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// RegisterHealthRoutes configura las rutas para el endpoint de salud.
func RegisterHealthRoutes(router fiber.Router, pgDB *gorm.DB, mongoDB *mongo.Database, version string, deps map[string]string) {
	health := router.Group("/health")

	// Crear el servicio de salud
	healthService := services.NewHealthService(pgDB, mongoDB, version, deps)

	// Crear el handler con el servicio inyectado
	healthHandler := handlers.NewHealthHandler(healthService)

	// Registrar la ruta GET /health
	health.Get("/", healthHandler.Health)

	// Ruta para obtener info del dispositivo
	router.Get("/device-info", func(c fiber.Ctx) error {
		// 1. Llamar a httpbun.com/get
		httpbinResp, err := http.Get("https://httpbun.com/get")
		if err != nil {
			log.Println("Error llamando a httpbin:", err)
			return c.Status(500).SendString("Error al obtener datos de httpbin")
		}
		defer httpbinResp.Body.Close()

		httpbinBody, _ := io.ReadAll(httpbinResp.Body)
		var httpbinData map[string]interface{}
		json.Unmarshal(httpbinBody, &httpbinData)

		// 2. Llamar a ipwho.is
		ipwhoResp, err := http.Get("https://ipwho.is/")
		if err != nil {
			log.Println("Error llamando a ipwho.is:", err)
			return c.Status(500).SendString("Error al obtener datos de ipwho.is")
		}
		defer ipwhoResp.Body.Close()

		ipwhoBody, _ := io.ReadAll(ipwhoResp.Body)
		var ipwhoData map[string]interface{}
		json.Unmarshal(ipwhoBody, &ipwhoData)

		// 3. Parsear el User-Agent desde headers de httpbin
		headers := httpbinData["headers"].(map[string]interface{})
		rawUA := headers["User-Agent"].(string)

		ua := user_agent.New(rawUA)
		browserName, browserVersion := ua.Browser()
		uaParsed := map[string]interface{}{
			"original": rawUA,
			"bot":      ua.Bot(),
			"mozilla":  ua.Mozilla(),
			"os":       ua.OS(),
			"platform": ua.Platform(),
			"mobile":   ua.Mobile(),
			"browser": map[string]string{
				"name":    browserName,
				"version": browserVersion,
			},
		}

		// 4. Obtener hora actual desde ipwho.is si existe
		var currentTime string
		if tz, ok := ipwhoData["timezone"].(map[string]interface{}); ok {
			if val, ok := tz["current_time"].(string); ok {
				currentTime = val
			}
		}

		// 5. Combinar todo
		result := map[string]interface{}{
			"httpbin":      httpbinData,
			"ipwhois":      ipwhoData,
			"user_agent":   uaParsed,
			"client_ip":    httpbinData["origin"],
			"requested_at": currentTime,
		}

		// 6. Mostrar en consola
		jsonPretty, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(jsonPretty))

		// 7. Retornar como respuesta JSON
		return c.Status(200).JSON(result)
	})
}
