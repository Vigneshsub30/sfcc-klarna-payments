package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/klarna-payments-api-v1/mcp-server/config"
	"github.com/klarna-payments-api-v1/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func PurchasetokenHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody models.Customertokencreationrequest
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/payments/v1/authorizations/%s/customer-token", cfg.BaseURL, authorizationToken)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.Customertokencreationresponse
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

func CreatePurchasetokenTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_payments_v1_authorizations_authorizationToken_customer-token",
		mcp.WithDescription("Generate a consumer token"),
		mcp.WithString("authorizationToken", mcp.Required(), mcp.Description("")),
		mcp.WithString("purchase_country", mcp.Required(), mcp.Description("Input parameter: ISO 3166 alpha-2 purchase country.")),
		mcp.WithString("purchase_currency", mcp.Required(), mcp.Description("Input parameter: ISO 4217 purchase currency.")),
		mcp.WithObject("billing_address", mcp.Description("")),
		mcp.WithObject("customer", mcp.Description("")),
		mcp.WithString("description", mcp.Required(), mcp.Description("Input parameter: Description of the purpose of the token.")),
		mcp.WithString("intended_use", mcp.Required(), mcp.Description("Input parameter: Intended use for the token.")),
		mcp.WithString("locale", mcp.Required(), mcp.Description("Input parameter: RFC 1766 customer's locale.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PurchasetokenHandler(cfg),
	}
}
