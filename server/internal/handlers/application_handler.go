package handlers

import (
	"context"

	appmodels "github.com/t-saturn/central-user-manager/server/internal/models"
	"github.com/t-saturn/central-user-manager/server/internal/services"
	applicationpb "github.com/t-saturn/central-user-manager/server/pb/proto/application"
)

type ApplicationHandler struct {
	applicationpb.UnimplementedApplicationServiceServer
	Service services.ApplicationService
}

func NewApplicationHandler(service services.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{Service: service}
}

func (h *ApplicationHandler) FnCreateApplication(ctx context.Context, req *applicationpb.CreateApplicationRequest) (*applicationpb.CreateApplicationResponse, error) {
	app := &appmodels.Application{
		Name:         req.GetName(),
		ClientID:     req.GetClientId(),
		ClientSecret: req.GetClientSecret(),
		Domain:       req.GetDomain(),
		Logo:         req.GetLogo(),
		Description:  req.GetDescription(),
		CallbackURLs: req.GetCallbackUrls(),
		Scopes:       req.GetScopes(),
		IsFirstParty: req.GetIsFirstParty(),
		Status:       appmodels.ApplicationActive,
	}

	result, err := h.Service.CreateApplication(app)
	if err != nil {
		return nil, err
	}

	return &applicationpb.CreateApplicationResponse{Id: result.ID.String()}, nil
}
