package services

import (
	"context"
	"net/http"
	"time"

	"github.com/t-saturn/auth-service-server/internal/dto"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// HealthService chequea el estado de servicio, bases de datos y dependencias.
type HealthService struct {
	mongoClient *mongo.Client
	pgDB        *gorm.DB
	startTime   time.Time
	version     string
	httpClient  *http.Client
	deps        map[string]string // nombre->URL de dependencia
}

// NewHealthService crea un HealthService con conexiones y configuración.
func NewHealthService(pgDB *gorm.DB, mongoDB *mongo.Database, version string, deps map[string]string) *HealthService {
	return &HealthService{
		mongoClient: mongoDB.Client(),
		pgDB:        pgDB,
		startTime:   time.Now().UTC(),
		version:     version,
		httpClient:  &http.Client{Timeout: 2 * time.Second},
		deps:        deps,
	}
}

// Check realiza el chequeo de salud.
func (s *HealthService) Check(ctx context.Context) (*dto.HealthResponseDTO, error) {
	now := time.Now().UTC()
	uptime := time.Since(s.startTime).Truncate(time.Second).String()

	// 1 MongoDB
	mStart := time.Now()
	mErr := s.mongoClient.Ping(ctx, nil)
	mRT := time.Since(mStart).Truncate(time.Millisecond).String()
	mStatus := "connected"
	if mErr != nil {
		mStatus = "down"
	}

	// 2 PostgreSQL
	sqlDB, err := s.pgDB.DB()
	if err != nil {
		return nil, err
	}
	pStart := time.Now()
	pErr := sqlDB.PingContext(ctx)
	pRT := time.Since(pStart).Truncate(time.Millisecond).String()
	pStatus := "connected"
	if pErr != nil {
		pStatus = "down"
	}

	// 3 Dependencias externas
	depHealth := make(map[string]string, len(s.deps))
	for name, url := range s.deps {
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		resp, err := s.httpClient.Do(req)
		status := "healthy"
		if err != nil || resp.StatusCode >= 400 {
			status = "down"
		}
		depHealth[name] = status
		if resp != nil {
			resp.Body.Close()
		}
	}

	// 4 Construir DTO de respuesta
	return &dto.HealthResponseDTO{
		Status:    "healthy",
		Timestamp: now,
		Version:   s.version,
		Uptime:    uptime,
		Databases: map[string]dto.DatabaseHealthDTO{
			"mongodb":    {Status: mStatus, ResponseTime: mRT},
			"postgresql": {Status: pStatus, ResponseTime: pRT},
		},
		Dependencies: depHealth,
	}, nil
}
