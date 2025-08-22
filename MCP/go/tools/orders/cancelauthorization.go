package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/klarna-payments-api-v1/mcp-server/config"
	"github.com/klarna-payments-api-v1/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CancelauthorizationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		authorizationTokenVal, ok := args["authorizationToken"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: authorizationToken"), nil
		}
		authorizationToken, ok := authorizationTokenVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: authorizationToken"), nil
		}
		url := fmt.Sprintf("%s/payments/v1/authorizations/%s", cfg.BaseURL, authorizationToken)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCancelauthorizationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_payments_v1_authorizations_authorizationToken",
		mcp.WithDescription("Cancel an existing authorization"),
		mcp.WithString("authorizationToken", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CancelauthorizationHandler(cfg),
	}
}
