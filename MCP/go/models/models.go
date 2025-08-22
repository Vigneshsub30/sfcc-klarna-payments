package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ErrorV2 represents the ErrorV2 schema from the OpenAPI specification
type ErrorV2 struct {
	Reason string `json:"reason,omitempty"`
	Authorized_payment_method Authorizedpaymentmethod `json:"authorized_payment_method,omitempty"`
	Correlation_id string `json:"correlation_id,omitempty"`
	Error_code string `json:"error_code,omitempty"`
	Error_messages []string `json:"error_messages,omitempty"`
	Fraud_status string `json:"fraud_status,omitempty"`
}

// Asseturls represents the Asseturls schema from the OpenAPI specification
type Asseturls struct {
	Descriptive string `json:"descriptive,omitempty"` // URL of the descriptive asset. Using this dynamic asset will make sure that any copy update of Klarna will automatically be propagated.
	Standard string `json:"standard,omitempty"` // URL of the standard asset. Using this dynamic asset will make sure that any copy update of Klarna will automatically be propagated.
}

// Session represents the Session schema from the OpenAPI specification
type Session struct {
	Expires_at string `json:"expires_at,omitempty"` // Session expiration date
	Intent string `json:"intent,omitempty"` // Intent for the session. The field is designed to let partners inform Klarna of the purpose of the customer’s session.
	Shipping_address Address `json:"shipping_address,omitempty"`
	Attachment Attachment `json:"attachment,omitempty"`
	Custom_payment_method_ids []string `json:"custom_payment_method_ids,omitempty"` // Promo codes - The array could be used to define which of the configured payment options within a payment category (pay_later, pay_over_time, etc.) should be shown for this purchase. Discuss with the delivery manager to know about the promo codes that will be configured for your account. The feature could also be used to provide promotional offers to specific customers (eg: 0% financing). Please be informed that the usage of this feature can have commercial implications.
	Merchant_reference1 string `json:"merchant_reference1,omitempty"` // Used for storing merchant's internal order number or other reference.
	Order_amount int64 `json:"order_amount,omitempty"` // Total amount of the order including tax and any available discounts. The value should be in non-negative minor units. Eg: 25 Euros should be 2500.
	Authorization_token string `json:"authorization_token,omitempty"` // Authorization token.
	Purchase_currency string `json:"purchase_currency,omitempty"` // The purchase currency of the order. Formatted according to ISO 4217 standard, e.g. USD, EUR, SEK, GBP, etc.
	Client_token string `json:"client_token,omitempty"` // Token to be passed to the JS client
	Order_lines []Orderline `json:"order_lines,omitempty"` // The array containing list of line items that are part of this order. Maximum of 1000 line items could be processed in a single order.
	Customer Customer `json:"customer,omitempty"`
	Status string `json:"status,omitempty"` // The current status of the session. Possible values: 'complete', 'incomplete' where 'complete' is set when the order has been placed.
	Payment_method_categories []Paymentmethodcategory `json:"payment_method_categories,omitempty"` // Available payment method categories
	Billing_address Address `json:"billing_address,omitempty"`
	Merchant_data string `json:"merchant_data,omitempty"` // Pass through field to send any information about the order to be used later for reference while retrieving the order details (max 6000 characters)
	Merchant_reference2 string `json:"merchant_reference2,omitempty"` // Used for storing merchant's internal order number or other reference. The value is available in the settlement files. (max 255 characters).
	Merchant_urls Merchanturls `json:"merchant_urls,omitempty"`
	Options Options `json:"options,omitempty"`
	Locale string `json:"locale,omitempty"` // Used to define the language and region of the customer. The locale follows the format of (RFC 1766)[https://datatracker.ietf.org/doc/rfc1766/], meaning its value consists of language-country. The following values are applicable: AT: "de-AT", "de-DE", "en-DE" BE: "be-BE", "nl-BE", "fr-BE", "en-BE" CH: "it-CH", "de-CH", "fr-CH", "en-CH" DE: "de-DE", "de-AT", "en-DE" DK: "da-DK", "en-DK" ES: "es-ES", "ca-ES", "en-ES" FI: "fi-FI", "sv-FI", "en-FI" GB: "en-GB" IT: "it-IT", "en-IT" NL: "nl-NL", "en-NL" NO: "nb-NO", "en-NO" PL: "pl-PL", "en-PL" SE: "sv-SE", "en-SE" US: "en-US".
	Design string `json:"design,omitempty"` // Design package to use in the session. This can only by used if a custom design has been implemented for Klarna Payments and agreed upon in the agreement. It might have a financial impact. Delivery manager will provide the value for the parameter.
	Order_tax_amount int64 `json:"order_tax_amount,omitempty"` // Total tax amount of the order. The value should be in non-negative minor units. Eg: 25 Euros should be 2500.
	Purchase_country string `json:"purchase_country,omitempty"` // The purchase country of the customer. The billing country always overrides purchase country if the values are different. Formatted according to ISO 3166 alpha-2 standard, e.g. GB, SE, DE, US, etc.
	Acquiring_channel string `json:"acquiring_channel,omitempty"` // The acquiring channel in which the session takes place. Ecommerce is default unless specified. Any other values should be defined in the agreement.
}

// Merchantsession represents the Merchantsession schema from the OpenAPI specification
type Merchantsession struct {
	Client_token string `json:"client_token"` // Client token to be passed to the JS client while initializing the JS SDK in the next step.
	Payment_method_categories []Paymentmethodcategory `json:"payment_method_categories,omitempty"` // Available payment method categories for this particular session
	Session_id string `json:"session_id"` // ID of the created session. Please use this ID to share with Klarna for identifying any issues during integration.
}

// Orderline represents the Orderline schema from the OpenAPI specification
type Orderline struct {
	Merchant_data string `json:"merchant_data,omitempty"` // Used for storing merchant's internal order number or other reference. Pass through field. (max 1024 characters)
	Product_url string `json:"product_url,omitempty"` // URL to the product in the merchant’s webshop that can be later used in communications between Klarna and the customer. (max 1024 characters)
	Unit_price int64 `json:"unit_price"` // Price for a single unit of the order line. Must be defined as minor units. Includes tax, excludes discount. (max value: 100000000)
	Quantity_unit string `json:"quantity_unit,omitempty"` // Unit used to describe the quantity, e.g. kg, pcs, etc. If defined the value has to be 1-8 characters.
	Total_amount int64 `json:"total_amount"` // Total amount of the order line. Must be defined as minor units. Includes tax and discount. Eg: 2500=25 euros Value = (quantity x unit_price) - total_discount_amount. (max value: 100000000)
	Total_tax_amount int64 `json:"total_tax_amount,omitempty"` // Total tax amount of the order line. Must be within ±1 of total_amount - total_amount 10000 / (10000 + tax_rate). Negative when type is discount.
	TypeField string `json:"type,omitempty"` // Type of the order line item. The possible values are: physical discount shipping_fee sales_tax digital gift_card store_credit surcharge
	Subscription Subscription `json:"subscription,omitempty"`
	Quantity int64 `json:"quantity"` // Quantity of the order line item. Must be a non-negative number.
	Total_discount_amount int64 `json:"total_discount_amount,omitempty"` // Non-negative minor units. Includes tax. Eg: 500=5 euros
	Image_url string `json:"image_url,omitempty"` // URL to an image that can be later embedded in communications between Klarna and the customer. (max 1024 characters). A minimum of 250x250 px resolution is recommended for the image to look good in the Klarna app, and below 50x50 px won't even show. We recommend using a good sized image (650x650 px or more), however the file size must not exceed 12MB.
	Name string `json:"name"` // Descriptive name of the order line item.
	Tax_rate int64 `json:"tax_rate,omitempty"` // Tax rate of the order line. Non-negative value. The percentage value is represented with two implicit decimals. I.e 1900 = 19%.
	Product_identifiers Productidentifiers `json:"product_identifiers,omitempty"`
	Reference string `json:"reference,omitempty"` // Client facing article number, SKU or similar. Max length is 256 characters.
}

// Authorizedpaymentmethod represents the Authorizedpaymentmethod schema from the OpenAPI specification
type Authorizedpaymentmethod struct {
	TypeField string `json:"type"`
	Number_of_days int `json:"number_of_days,omitempty"`
	Number_of_installments int `json:"number_of_installments,omitempty"`
}

// Createorderrequest represents the Createorderrequest schema from the OpenAPI specification
type Createorderrequest struct {
	Purchase_country string `json:"purchase_country"` // The purchase country of the customer. The billing country always overrides purchase country if the values are different. Formatted according to ISO 3166 alpha-2 standard, e.g. GB, SE, DE, US, etc.
	Order_amount int64 `json:"order_amount"` // Total amount of the order including tax and any available discounts. The value should be in non-negative minor units. Eg: 25 Euros should be 2500.
	Custom_payment_method_ids []string `json:"custom_payment_method_ids,omitempty"` // Promo codes - The array could be used to define which of the configured payment options within a payment category (pay_later, pay_over_time, etc.) should be shown for this purchase. Discuss with the delivery manager to know about the promo codes that will be configured for your account. The feature could also be used to provide promotional offers to specific customers (eg: 0% financing). Please be informed that the usage of this feature can have commercial implications.
	Locale string `json:"locale,omitempty"` // Used to define the language and region of the customer. The locale follows the format of (RFC 1766)[https://datatracker.ietf.org/doc/rfc1766/], meaning its value consists of language-country. The following values are applicable: AT: "de-AT", "de-DE", "en-DE" BE: "be-BE", "nl-BE", "fr-BE", "en-BE" CH: "it-CH", "de-CH", "fr-CH", "en-CH" DE: "de-DE", "de-AT", "en-DE" DK: "da-DK", "en-DK" ES: "es-ES", "ca-ES", "en-ES" FI: "fi-FI", "sv-FI", "en-FI" GB: "en-GB" IT: "it-IT", "en-IT" NL: "nl-NL", "en-NL" NO: "nb-NO", "en-NO" PL: "pl-PL", "en-PL" SE: "sv-SE", "en-SE" US: "en-US".
	Merchant_reference2 string `json:"merchant_reference2,omitempty"` // Used for storing merchant's internal order number or other reference. The value is available in the settlement files. (max 255 characters).
	Order_lines []Orderline `json:"order_lines"` // The array containing list of line items that are part of this order. Maximum of 1000 line items could be processed in a single order.
	Payment_method_categories []Paymentmethodcategory `json:"payment_method_categories,omitempty"` // Available payment method categories
	Customer Customer `json:"customer,omitempty"`
	Status string `json:"status,omitempty"` // The current status of the session. Possible values: 'complete', 'incomplete' where 'complete' is set when the order has been placed.
	Order_tax_amount int64 `json:"order_tax_amount,omitempty"` // Total tax amount of the order. The value should be in non-negative minor units. Eg: 25 Euros should be 2500.
	Billing_address Address `json:"billing_address,omitempty"`
	Purchase_currency string `json:"purchase_currency"` // The purchase currency of the order. Formatted according to ISO 4217 standard, e.g. USD, EUR, SEK, GBP, etc.
	Shipping_address Address `json:"shipping_address,omitempty"`
	Authorization_token string `json:"authorization_token,omitempty"` // Authorization token.
	Auto_capture bool `json:"auto_capture,omitempty"` // Allow merchant to trigger auto capturing.
	Merchant_data string `json:"merchant_data,omitempty"` // Pass through field to send any information about the order to be used later for reference while retrieving the order details (max 6000 characters)
	Merchant_reference1 string `json:"merchant_reference1,omitempty"` // Used for storing merchant's internal order number or other reference.
	Merchant_urls Merchanturls `json:"merchant_urls,omitempty"`
}

// Order represents the Order schema from the OpenAPI specification
type Order struct {
	Redirect_url string `json:"redirect_url,omitempty"` // URL to redirect the customer to after placing the order. This is a Klarna URL to which the merchant should redirect the customer to. Klarna will place a cookie in the customer’s browser (if redirected) and redirect the customer back to the confirmation URL provided by the merchant. This is not a mandatory step but a recommended one to improve the returning customer’s experience. It is a spontaneous step and does not harm the customer’s experience.
	Authorized_payment_method Authorizedpaymentmethod `json:"authorized_payment_method,omitempty"`
	Fraud_status string `json:"fraud_status,omitempty"` // Fraud status for the order. Either ACCEPTED or PENDING. If ACCEPTED, the order could be captured. If PENDING, please wait till you receive the notification from Klarna in the notification URL that the order has been approved. You can find additional information here.
	Order_id string `json:"order_id"` // Unique order ID of the transaction. This ID will be used for all order management processes.
}

// Sessionread represents the Sessionread schema from the OpenAPI specification
type Sessionread struct {
	Client_token string `json:"client_token,omitempty"` // Token to be passed to the JS client
	Customer Customerread `json:"customer,omitempty"`
	Attachment Attachment `json:"attachment,omitempty"`
	Options Options `json:"options,omitempty"`
	Order_lines []Orderline `json:"order_lines,omitempty"` // The array containing list of line items that are part of this order. Maximum of 1000 line items could be processed in a single order.
	Merchant_reference2 string `json:"merchant_reference2,omitempty"` // Used for storing merchant's internal order number or other reference. The value is available in the settlement files. (max 255 characters).
	Shipping_address Address `json:"shipping_address,omitempty"`
	Expires_at string `json:"expires_at,omitempty"` // Session expiration date
	Merchant_data string `json:"merchant_data,omitempty"` // Pass through field to send any information about the order to be used later for reference while retrieving the order details (max 6000 characters)
	Merchant_reference1 string `json:"merchant_reference1,omitempty"` // Used for storing merchant's internal order number or other reference.
	Order_amount int64 `json:"order_amount,omitempty"` // Total amount of the order including tax and any available discounts. The value should be in non-negative minor units. Eg: 25 Euros should be 2500.
	Payment_method_categories []Paymentmethodcategory `json:"payment_method_categories,omitempty"` // Available payment method categories
	Custom_payment_method_ids []string `json:"custom_payment_method_ids,omitempty"` // Promo codes - The array could be used to define which of the configured payment options within a payment category (pay_later, pay_over_time, etc.) should be shown for this purchase. Discuss with the delivery manager to know about the promo codes that will be configured for your account. The feature could also be used to provide promotional offers to specific customers (eg: 0% financing). Please be informed that the usage of this feature can have commercial implications.
	Purchase_country string `json:"purchase_country,omitempty"` // The purchase country of the customer. The billing country always overrides purchase country if the values are different. Formatted according to ISO 3166 alpha-2 standard, e.g. GB, SE, DE, US, etc.
	Intent string `json:"intent,omitempty"` // Intent for the session. The field is designed to let partners inform Klarna of the purpose of the customer’s session.
	Status string `json:"status,omitempty"` // The current status of the session. Possible values: 'complete', 'incomplete' where 'complete' is set when the order has been placed.
	Billing_address Address `json:"billing_address,omitempty"`
	Acquiring_channel string `json:"acquiring_channel,omitempty"` // The acquiring channel in which the session takes place. Ecommerce is default unless specified. Any other values should be defined in the agreement.
	Locale string `json:"locale,omitempty"` // Used to define the language and region of the customer. The locale follows the format of (RFC 1766)[https://datatracker.ietf.org/doc/rfc1766/], meaning its value consists of language-country. The following values are applicable: AT: "de-AT", "de-DE", "en-DE" BE: "be-BE", "nl-BE", "fr-BE", "en-BE" CH: "it-CH", "de-CH", "fr-CH", "en-CH" DE: "de-DE", "de-AT", "en-DE" DK: "da-DK", "en-DK" ES: "es-ES", "ca-ES", "en-ES" FI: "fi-FI", "sv-FI", "en-FI" GB: "en-GB" IT: "it-IT", "en-IT" NL: "nl-NL", "en-NL" NO: "nb-NO", "en-NO" PL: "pl-PL", "en-PL" SE: "sv-SE", "en-SE" US: "en-US".
	Merchant_urls Merchanturls `json:"merchant_urls,omitempty"`
	Order_tax_amount int64 `json:"order_tax_amount,omitempty"` // Total tax amount of the order. The value should be in non-negative minor units. Eg: 25 Euros should be 2500.
	Purchase_currency string `json:"purchase_currency,omitempty"` // The purchase currency of the order. Formatted according to ISO 4217 standard, e.g. USD, EUR, SEK, GBP, etc.
	Design string `json:"design,omitempty"` // Design package to use in the session. This can only by used if a custom design has been implemented for Klarna Payments and agreed upon in the agreement. It might have a financial impact. Delivery manager will provide the value for the parameter.
	Authorization_token string `json:"authorization_token,omitempty"` // Authorization token.
}

// Customerread represents the Customerread schema from the OpenAPI specification
type Customerread struct {
	Organization_registration_id string `json:"organization_registration_id,omitempty"` // Organization registration id. Only applicable for B2B customers.
	Title string `json:"title,omitempty"` // Customer’s Title. Allowed values per country: UK - "Mr", "Ms" DE - "Herr", "Frau" AT: "Herr, "Frau" CH: de-CH: "Herr, "Frau" it-CH: "Sig.", "Sig.ra" fr-CH: "M", "Mme" BE: "Dhr.", "Mevr." NL: "Dhr.", "Mevr."
	TypeField string `json:"type,omitempty"` // Type of customer in the session. If nothing is added, a B2C session will be the default. If it is a b2b-session, you should enter organization to trigger a B2B session.
	Vat_id string `json:"vat_id,omitempty"` // VAT ID. Only applicable for B2B customers.
	Date_of_birth string `json:"date_of_birth,omitempty"` // Customer’s date of birth. The format is ‘yyyy-mm-dd’
	Gender string `json:"gender,omitempty"` // Customer’s gender - ‘male’ or ‘female’
	Organization_entity_type string `json:"organization_entity_type,omitempty"` // Organization entity type. Only applicable for B2B customers.
}

// Attachment represents the Attachment schema from the OpenAPI specification
type Attachment struct {
	Body string `json:"body"` // The content of the extra merchant data should be presented as a string inside this property. The body should be an object containing any of the keys and sub-objects described below serialized to JSON. More information on that object can be found [here](https://developers.klarna.com/api/#attachment-schema).
	Content_type string `json:"content_type"` // The content type of the body. It is usually represented as "application/vnd.klarna.internal.emd-v2+json"
}

// Customer represents the Customer schema from the OpenAPI specification
type Customer struct {
	Organization_entity_type string `json:"organization_entity_type,omitempty"` // Organization entity type. Only applicable for B2B customers.
	TypeField string `json:"type,omitempty"` // Type of customer in the session. If nothing is added, a B2C session will be the default. If it is a b2b-session, you should enter organization to trigger a B2B session.
	Vat_id string `json:"vat_id,omitempty"` // VAT ID. Only applicable for B2B customers.
	Last_four_ssn string `json:"last_four_ssn,omitempty"` // Last four digits of the customer's social security number. This value is available for US customers.
	Date_of_birth string `json:"date_of_birth,omitempty"` // Customer’s date of birth. The format is ‘yyyy-mm-dd’
	Gender string `json:"gender,omitempty"` // Customer’s gender - ‘male’ or ‘female’
	National_identification_number string `json:"national_identification_number,omitempty"` // The customer's national identification number. This value is available for EU customers utilizing national identification numbers.
	Title string `json:"title,omitempty"` // Customer’s Title. Allowed values per country: UK - "Mr", "Ms" DE - "Herr", "Frau" AT: "Herr, "Frau" CH: de-CH: "Herr, "Frau" it-CH: "Sig.", "Sig.ra" fr-CH: "M", "Mme" BE: "Dhr.", "Mevr." NL: "Dhr.", "Mevr."
	Organization_registration_id string `json:"organization_registration_id,omitempty"` // Organization registration id. Only applicable for B2B customers.
}

// Customertokencreationrequest represents the Customertokencreationrequest schema from the OpenAPI specification
type Customertokencreationrequest struct {
	Locale string `json:"locale"` // RFC 1766 customer's locale.
	Purchase_country string `json:"purchase_country"` // ISO 3166 alpha-2 purchase country.
	Purchase_currency string `json:"purchase_currency"` // ISO 4217 purchase currency.
	Billing_address Address `json:"billing_address,omitempty"`
	Customer Customer `json:"customer,omitempty"`
	Description string `json:"description"` // Description of the purpose of the token.
	Intended_use string `json:"intended_use"` // Intended use for the token.
}

// Sessioncreate represents the Sessioncreate schema from the OpenAPI specification
type Sessioncreate struct {
	Options Options `json:"options,omitempty"`
	Status string `json:"status,omitempty"` // The current status of the session. Possible values: 'complete', 'incomplete' where 'complete' is set when the order has been placed.
	Billing_address Address `json:"billing_address,omitempty"`
	Custom_payment_method_ids []string `json:"custom_payment_method_ids,omitempty"` // Promo codes - The array could be used to define which of the configured payment options within a payment category (pay_later, pay_over_time, etc.) should be shown for this purchase. Discuss with the delivery manager to know about the promo codes that will be configured for your account. The feature could also be used to provide promotional offers to specific customers (eg: 0% financing). Please be informed that the usage of this feature can have commercial implications.
	Design string `json:"design,omitempty"` // Design package to use in the session. This can only by used if a custom design has been implemented for Klarna Payments and agreed upon in the agreement. It might have a financial impact. Delivery manager will provide the value for the parameter.
	Locale string `json:"locale,omitempty"` // Used to define the language and region of the customer. The locale follows the format of (RFC 1766)[https://datatracker.ietf.org/doc/rfc1766/], meaning its value consists of language-country. The following values are applicable: AT: "de-AT", "de-DE", "en-DE" BE: "be-BE", "nl-BE", "fr-BE", "en-BE" CH: "it-CH", "de-CH", "fr-CH", "en-CH" DE: "de-DE", "de-AT", "en-DE" DK: "da-DK", "en-DK" ES: "es-ES", "ca-ES", "en-ES" FI: "fi-FI", "sv-FI", "en-FI" GB: "en-GB" IT: "it-IT", "en-IT" NL: "nl-NL", "en-NL" NO: "nb-NO", "en-NO" PL: "pl-PL", "en-PL" SE: "sv-SE", "en-SE" US: "en-US". Default value is "en-US".
	Purchase_country string `json:"purchase_country"` // The purchase country of the customer. The billing country always overrides purchase country if the values are different. Formatted according to ISO 3166 alpha-2 standard, e.g. GB, SE, DE, US, etc.
	Shipping_address Address `json:"shipping_address,omitempty"`
	Authorization_token string `json:"authorization_token,omitempty"` // Authorization token.
	Merchant_data string `json:"merchant_data,omitempty"` // Pass through field to send any information about the order to be used later for reference while retrieving the order details (max 6000 characters)
	Payment_method_categories []Paymentmethodcategory `json:"payment_method_categories,omitempty"` // Available payment method categories
	Customer Customer `json:"customer,omitempty"`
	Intent string `json:"intent,omitempty"` // Intent for the session. The field is designed to let partners inform Klarna of the purpose of the customer’s session.
	Merchant_reference1 string `json:"merchant_reference1,omitempty"` // Used for storing merchant's internal order number or other reference.
	Order_tax_amount int64 `json:"order_tax_amount,omitempty"` // Total tax amount of the order. The value should be in non-negative minor units. Eg: 25 Euros should be 2500.
	Purchase_currency string `json:"purchase_currency"` // The purchase currency of the order. Formatted according to ISO 4217 standard, e.g. USD, EUR, SEK, GBP, etc.
	Merchant_reference2 string `json:"merchant_reference2,omitempty"` // Used for storing merchant's internal order number or other reference. The value is available in the settlement files. (max 255 characters).
	Order_amount int64 `json:"order_amount"` // Total amount of the order including tax and any available discounts. The value should be in non-negative minor units. Eg: 25 Euros should be 2500.
	Order_lines []Orderline `json:"order_lines"` // The array containing list of line items that are part of this order. Maximum of 1000 line items could be processed in a single order.
	Merchant_urls Merchanturls `json:"merchant_urls,omitempty"`
	Attachment Attachment `json:"attachment,omitempty"`
	Acquiring_channel string `json:"acquiring_channel,omitempty"` // The acquiring channel in which the session takes place. Ecommerce is default unless specified. Any other values should be defined in the agreement.
	Client_token string `json:"client_token,omitempty"` // Token to be passed to the JS client
	Expires_at string `json:"expires_at,omitempty"` // Session expiration date
}

// Customertokencreationresponse represents the Customertokencreationresponse schema from the OpenAPI specification
type Customertokencreationresponse struct {
	Redirect_url string `json:"redirect_url,omitempty"` // URL to redirect the customer to after placing the order. This is a Klarna URL where Klarna will place a cookie in the customer’s browser (if redirected) and redirect the customer back to the confirmation URL provided by the merchant. This is not a mandatory step but a recommended one to improve the returning customer’s experience.
	Token_id string `json:"token_id"` // Generated customer token. This token will be used to create a new order for the subscription using the Create a New order using token API.
	Billing_address Address `json:"billing_address,omitempty"`
	Customer Customerreadcreatetoken `json:"customer,omitempty"`
	Payment_method_reference string `json:"payment_method_reference,omitempty"` // Used to connect customers with payment method when it is present.
}

// Address represents the Address schema from the OpenAPI specification
type Address struct {
	Street_address string `json:"street_address,omitempty"` // Customer’s street address. Regional formatting is required, as follows: UK/US/FR: 33 Cavendish Square Rest of EU: De Ruijterkade 7
	Organization_name string `json:"organization_name,omitempty"` // Organization name (if applicable). Only applicable for B2B customers.
	Postal_code string `json:"postal_code,omitempty"` // Customer’s postal code. Validation according to Universal Postal Union addressing systems. E.g. 12345, W1G OPW.
	Country string `json:"country,omitempty"` // Customer’s country. This value overrides the purchase country if they are different. Should follow the standard of ISO 3166 alpha-2. E.g. GB, US, DE, SE.
	Given_name string `json:"given_name,omitempty"` // Customers given name in UTF-8 encoding. Cannot be only numbers, must be more than 1 character. Allowed special characters: -'’. More information can be found [in this link](https://docs.klarna.com/klarna-payments/in-depth-knowledge/customer-data-requirements/#details-needed-per-market)
	Region string `json:"region,omitempty"` // Customer’s region or state - Mandatory for US and AU market. Validations according to ISO 3166-2 format, e.g. OH, NJ, etc.
	City string `json:"city,omitempty"` // Customer’s city.
	Street_address2 string `json:"street_address2,omitempty"` // Customer’s street address. Second Line.
	Title string `json:"title,omitempty"` // Customer’s Title. Allowed values per country: UK - "Mr", "Ms" DE - "Herr", "Frau" AT: "Herr, "Frau" CH: de-CH: "Herr, "Frau" it-CH: "Sig.", "Sig.ra" fr-CH: "M", "Mme" BE: "Dhr.", "Mevr." NL: "Dhr.", "Mevr."
	Attention string `json:"attention,omitempty"` // ‘Attn.’ (if applicable). Only applicable for B2B customers.
	Email string `json:"email,omitempty"` // Customer’s email address.
	Family_name string `json:"family_name,omitempty"` // Customers family name in UTF-8 encoding. Cannot be only numbers, must be more than 1 character. Allowed special characters: -'’. More information can be found [in this link](https://docs.klarna.com/klarna-payments/in-depth-knowledge/customer-data-requirements/#details-needed-per-market)
	Phone string `json:"phone,omitempty"` // Phone number. Preferably a mobile phone number.
}

// Productidentifiers represents the Productidentifiers schema from the OpenAPI specification
type Productidentifiers struct {
	Color string `json:"color,omitempty"` // Color to be shown to the end customer (max 64 characters).
	Global_trade_item_number string `json:"global_trade_item_number,omitempty"` // The product's Global Trade Item Number (GTIN). Common types of GTIN are EAN, ISBN or UPC. Exclude dashes and spaces, where possible
	Manufacturer_part_number string `json:"manufacturer_part_number,omitempty"` // The product's Manufacturer Part Number (MPN), which - together with the brand - uniquely identifies a product. Only submit MPNs assigned by a manufacturer and use the most specific MPN possible
	Size string `json:"size,omitempty"` // Size to be shown to the end customer (max 64 characters).
	Brand string `json:"brand,omitempty"` // The product's brand name as generally recognized by consumers. If no brand is available for a product, do not supply any value.
	Category_path string `json:"category_path,omitempty"` // The product's category path as used in the merchant's webshop. Include the full and most detailed category and separate the segments with ' > '
}

// Subscription represents the Subscription schema from the OpenAPI specification
type Subscription struct {
	Interval string `json:"interval"` // The cadence unit for this.
	Interval_count int `json:"interval_count"` // The number of intervals
	Name string `json:"name"` // The name of the subscription product
}

// Customerreadcreatetoken represents the Customerreadcreatetoken schema from the OpenAPI specification
type Customerreadcreatetoken struct {
	Gender string `json:"gender,omitempty"` // Customer’s gender - ‘male’ or ‘female’
	Date_of_birth string `json:"date_of_birth,omitempty"` // Customer’s date of birth. The format is ‘yyyy-mm-dd’
}

// Options represents the Options schema from the OpenAPI specification
type Options struct {
	Color_text string `json:"color_text,omitempty"` // Color for the texts within the iFrame. Value should be a CSS hex color, e.g. "#FF9900"
	Radius_border string `json:"radius_border,omitempty"` // Radius for the border of elements within the iFrame.
	Color_border string `json:"color_border,omitempty"` // Color for the border of elements within the iFrame. Value should be a CSS hex color, e.g. "#FF9900"
	Color_border_selected string `json:"color_border_selected,omitempty"` // Color for the border of elements within the iFrame when selected by the customer. Value should be a CSS hex color, e.g. "#FF9900"
	Color_details string `json:"color_details,omitempty"` // Color for the bullet points within the iFrame. Value should be a CSS hex color, e.g. "#FF9900"
}

// Paymentmethodcategory represents the Paymentmethodcategory schema from the OpenAPI specification
type Paymentmethodcategory struct {
	Name string `json:"name,omitempty"` // Name of the payment method category. These names are dynamic depending on what payment method is in the category. Using this dynamic asset will make sure that any copy update of Klarna will automatically be propagated, or any updates of included payment methods by you.
	Asset_urls Asseturls `json:"asset_urls,omitempty"`
	Identifier string `json:"identifier,omitempty"` // ID of the payment method category to be used while loading the widget later. The possible values are:<ul><li>klarna</li><li>pay_later</li><li>pay_now</li><li>pay_over_time</li><li>direct_bank_transfer</li><li>direct_debit</li></ul>
}

// Merchanturls represents the Merchanturls schema from the OpenAPI specification
type Merchanturls struct {
	Authorization string `json:"authorization,omitempty"` // URL for receiving the authorization token when payment is completed. Used for Authorization Callback.
	Confirmation string `json:"confirmation,omitempty"` // URL of the merchant confirmation page. The consumer will be redirected back to the confirmation page if the consumer is sent to the redirect URL after placing the order. Insert {session.id} and/or {order.id} as placeholder to connect either of those IDs to the URL(max 2000 characters).
	Notification string `json:"notification,omitempty"` // URL for notifications on pending orders. Insert {session.id} and/or {order.id} as placeholder to connect either of those IDs to the URL (max 2000 characters).
	Push string `json:"push,omitempty"` // URL that will be requested when an order is completed. Should be different than checkout and confirmation URLs. Insert {session.id} and/or {order.id} as placeholder to connect either of those IDs to the URL (max 2000 characters).
}
