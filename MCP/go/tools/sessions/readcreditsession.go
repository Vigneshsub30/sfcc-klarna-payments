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

func ReadcreditsessionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		session_idVal, ok := args["session_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: session_id"), nil
		}
		session_id, ok := session_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: session_id"), nil
		}
		url := fmt.Sprintf("%s/payments/v1/sessions/%s", cfg.BaseURL, session_id)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.Sessionread
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

func CreateReadcreditsessionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_payments_v1_sessions_session_id",
		mcp.WithDescription("Read an existing payment session"),
		mcp.WithString("session_id", mcp.Required(), mcp.Description("session_id")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ReadcreditsessionHandler(cfg),
	}
}
