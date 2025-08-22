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

func CreatecreditsessionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Sessioncreate
		
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
		url := fmt.Sprintf("%s/payments/v1/sessions", cfg.BaseURL)
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
		var result models.Merchantsession
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

func CreateCreatecreditsessionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_payments_v1_sessions",
		mcp.WithDescription("Create a new payment session"),
		mcp.WithString("status", mcp.Description("Input parameter: The current status of the session. Possible values: 'complete', 'incomplete' where 'complete' is set when the order has been placed.")),
		mcp.WithObject("billing_address", mcp.Description("")),
		mcp.WithArray("custom_payment_method_ids", mcp.Description("Input parameter: Promo codes - The array could be used to define which of the configured payment options within a payment category (pay_later, pay_over_time, etc.) should be shown for this purchase. Discuss with the delivery manager to know about the promo codes that will be configured for your account. The feature could also be used to provide promotional offers to specific customers (eg: 0% financing). Please be informed that the usage of this feature can have commercial implications. ")),
		mcp.WithString("design", mcp.Description("Input parameter: Design package to use in the session. This can only by used if a custom design has been implemented for Klarna Payments and agreed upon in the agreement. It might have a financial impact. Delivery manager will provide the value for the parameter.")),
		mcp.WithString("locale", mcp.Description("Input parameter: Used to define the language and region of the customer. The locale follows the format of (RFC 1766)[https://datatracker.ietf.org/doc/rfc1766/], meaning its value consists of language-country.\nThe following values are applicable:\n\nAT: \"de-AT\", \"de-DE\", \"en-DE\"\nBE: \"be-BE\", \"nl-BE\", \"fr-BE\", \"en-BE\"\nCH: \"it-CH\", \"de-CH\", \"fr-CH\", \"en-CH\"\nDE: \"de-DE\", \"de-AT\", \"en-DE\"\nDK: \"da-DK\", \"en-DK\"\nES: \"es-ES\", \"ca-ES\", \"en-ES\"\nFI: \"fi-FI\", \"sv-FI\", \"en-FI\"\nGB: \"en-GB\"\nIT: \"it-IT\", \"en-IT\"\nNL: \"nl-NL\", \"en-NL\"\nNO: \"nb-NO\", \"en-NO\"\nPL: \"pl-PL\", \"en-PL\"\nSE: \"sv-SE\", \"en-SE\"\nUS: \"en-US\". Default value is \"en-US\".")),
		mcp.WithString("purchase_country", mcp.Required(), mcp.Description("Input parameter: The purchase country of the customer. The billing country always overrides purchase country if the values are different. Formatted according to ISO 3166 alpha-2 standard, e.g. GB, SE, DE, US, etc.")),
		mcp.WithObject("shipping_address", mcp.Description("")),
		mcp.WithString("authorization_token", mcp.Description("Input parameter: Authorization token.")),
		mcp.WithString("merchant_data", mcp.Description("Input parameter: Pass through field to send any information about the order to be used later for reference while retrieving the order details (max 6000 characters)")),
		mcp.WithArray("payment_method_categories", mcp.Description("Input parameter: Available payment method categories")),
		mcp.WithObject("customer", mcp.Description("")),
		mcp.WithString("intent", mcp.Description("Input parameter: Intent for the session. The field is designed to let partners inform Klarna of the purpose of the customer’s session.")),
		mcp.WithString("merchant_reference1", mcp.Description("Input parameter: Used for storing merchant's internal order number or other reference.")),
		mcp.WithNumber("order_tax_amount", mcp.Description("Input parameter: Total tax amount of the order. The value should be in non-negative minor units. Eg: 25 Euros should be 2500.")),
		mcp.WithString("purchase_currency", mcp.Required(), mcp.Description("Input parameter: The purchase currency of the order. Formatted according to ISO 4217 standard, e.g. USD, EUR, SEK, GBP, etc.")),
		mcp.WithString("merchant_reference2", mcp.Description("Input parameter: Used for storing merchant's internal order number or other reference. The value is available in the settlement files. (max 255 characters).")),
		mcp.WithNumber("order_amount", mcp.Required(), mcp.Description("Input parameter: Total amount of the order including tax and any available discounts. The value should be in non-negative minor units. Eg: 25 Euros should be 2500.")),
		mcp.WithArray("order_lines", mcp.Required(), mcp.Description("Input parameter: The array containing list of line items that are part of this order. Maximum of 1000 line items could be processed in a single order.")),
		mcp.WithObject("merchant_urls", mcp.Description("")),
		mcp.WithObject("attachment", mcp.Description("")),
		mcp.WithString("acquiring_channel", mcp.Description("Input parameter: The acquiring channel in which the session takes place. Ecommerce is default unless specified. Any other values should be defined in the agreement.")),
		mcp.WithString("client_token", mcp.Description("Input parameter: Token to be passed to the JS client")),
		mcp.WithString("expires_at", mcp.Description("Input parameter: Session expiration date")),
		mcp.WithObject("options", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreatecreditsessionHandler(cfg),
	}
}
