package main

import (
	"github.com/klarna-payments-api-v1/mcp-server/config"
	"github.com/klarna-payments-api-v1/mcp-server/models"
	tools_orders "github.com/klarna-payments-api-v1/mcp-server/tools/orders"
	tools_sessions "github.com/klarna-payments-api-v1/mcp-server/tools/sessions"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_orders.CreatePurchasetokenTool(cfg),
		tools_orders.CreateCreateorderTool(cfg),
		tools_sessions.CreateCreatecreditsessionTool(cfg),
		tools_sessions.CreateReadcreditsessionTool(cfg),
		tools_sessions.CreateUpdatecreditsessionTool(cfg),
		tools_orders.CreateCancelauthorizationTool(cfg),
	}
}
