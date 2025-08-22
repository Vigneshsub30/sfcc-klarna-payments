/**
 * Update an existing payment session
 */

import fs from 'fs';
import path from 'path';
import os from 'os';

function getConfig() {
  const baseURL = process.env.API_BASE_URL;
  const bearerToken = process.env.API_BEARER_TOKEN;
  
  if (!baseURL || !bearerToken) {
    const configPath = path.join(os.homedir(), '.api', 'config.json');
    try {
      const configData = JSON.parse(fs.readFileSync(configPath, 'utf8'));
      return {
        baseURL: baseURL || configData.baseURL,
        bearerToken: bearerToken || configData.bearerToken
      };
    } catch (e) {
      throw new Error('Configuration not found. Please set API_BASE_URL and API_BEARER_TOKEN environment variables or create config file at ~/.api/config.json');
    }
  }
  
  return { baseURL, bearerToken };
}

export async function post_payments_v1_sessions_session_id(session_id, authorization_token, purchase_currency, client_token, status, merchant_data, merchant_reference2, locale, design, purchase_country, acquiring_channel, expires_at, intent, merchant_reference1, order_amount, order_tax_amount, customer, billing_address, merchant_urls, options, shipping_address, attachment, order_lines, payment_method_categories, custom_payment_method_ids) {
  try {
    const config = getConfig();
    const requestBody = {
      session_id,
      authorization_token,
      purchase_currency,
      client_token,
      status,
      merchant_data,
      merchant_reference2,
      locale,
      design,
      purchase_country,
      acquiring_channel,
      expires_at,
      intent,
      merchant_reference1,
      order_amount,
      order_tax_amount,
      customer,
      billing_address,
      merchant_urls,
      options,
      shipping_address,
      attachment,
      order_lines,
      payment_method_categories,
      custom_payment_method_ids
    };
    
    const url = `${config.baseURL}/api/unknown`;
    
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${config.bearerToken}`,
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(requestBody)
    });
    
    if (!response.ok) {
      return `Failed to read response body: ${response.status} ${response.statusText}`;
    }
    
    try {
      const result = await response.json();
      return JSON.stringify(result, null, 2);
    } catch (e) {
      return await response.text();
    }
    
  } catch (error) {
    return `Failed to create request: ${error.message}`;
  }
}

export function createPostPaymentsV1SessionsSessionIdTool() {
  return {
    definition: {
      name: 'post-payments-v1-sessions-session-id',
      description: 'Update an existing payment session',
      inputSchema: {
        type: 'object',
        properties: {
          session_id: {
            type: 'string',
            description: 'session_id'
          },
          authorization_token: {
            type: 'string',
            description: 'Input parameter: Authorization token.'
          },
          purchase_currency: {
            type: 'string',
            description: 'Input parameter: The purchase currency of the order. Formatted according to ISO 4217 standard, e.g. USD, EUR, SEK, GBP, etc.'
          },
          client_token: {
            type: 'string',
            description: 'Input parameter: Token to be passed to the JS client'
          },
          status: {
            type: 'string',
            description: 'Input parameter: The current status of the session. Possible values: 'complete', 'incomplete' where 'complete' is set when the order has been placed.'
          },
          merchant_data: {
            type: 'string',
            description: 'Input parameter: Pass through field to send any information about the order to be used later for reference while retrieving the order details (max 6000 characters)'
          },
          merchant_reference2: {
            type: 'string',
            description: 'Input parameter: Used for storing merchant's internal order number or other reference. The value is available in the settlement files. (max 255 characters).'
          },
          locale: {
            type: 'string',
            description: 'Input parameter: Used to define the language and region of the customer. The locale follows the format of (RFC 1766)[https://datatracker.ietf.org/doc/rfc1766/], meaning its value consists of language-country. The following values are applicable: AT: "de-AT", "de-DE", "en-DE" BE: "be-BE", "nl-BE", "fr-BE", "en-BE" CH: "it-CH", "de-CH", "fr-CH", "en-CH" DE: "de-DE", "de-AT", "en-DE" DK: "da-DK", "en-DK" ES: "es-ES", "ca-ES", "en-ES" FI: "fi-FI", "sv-FI", "en-FI" GB: "en-GB" IT: "it-IT", "en-IT" NL: "nl-NL", "en-NL" NO: "nb-NO", "en-NO" PL: "pl-PL", "en-PL" SE: "sv-SE", "en-SE" US: "en-US".'
          },
          design: {
            type: 'string',
            description: 'Input parameter: Design package to use in the session. This can only by used if a custom design has been implemented for Klarna Payments and agreed upon in the agreement. It might have a financial impact. Delivery manager will provide the value for the parameter.'
          },
          purchase_country: {
            type: 'string',
            description: 'Input parameter: The purchase country of the customer. The billing country always overrides purchase country if the values are different. Formatted according to ISO 3166 alpha-2 standard, e.g. GB, SE, DE, US, etc.'
          },
          acquiring_channel: {
            type: 'string',
            description: 'Input parameter: The acquiring channel in which the session takes place. Ecommerce is default unless specified. Any other values should be defined in the agreement.'
          },
          expires_at: {
            type: 'string',
            description: 'Input parameter: Session expiration date'
          },
          intent: {
            type: 'string',
            description: 'Input parameter: Intent for the session. The field is designed to let partners inform Klarna of the purpose of the customer’s session.'
          },
          merchant_reference1: {
            type: 'string',
            description: 'Input parameter: Used for storing merchant's internal order number or other reference.'
          },
          order_amount: {
            type: 'number',
            description: 'Input parameter: Total amount of the order including tax and any available discounts. The value should be in non-negative minor units. Eg: 25 Euros should be 2500.'
          },
          order_tax_amount: {
            type: 'number',
            description: 'Input parameter: Total tax amount of the order. The value should be in non-negative minor units. Eg: 25 Euros should be 2500.'
          },
          customer: {
            type: 'object',
            description: ''
          },
          billing_address: {
            type: 'object',
            description: ''
          },
          merchant_urls: {
            type: 'object',
            description: ''
          },
          options: {
            type: 'object',
            description: ''
          },
          shipping_address: {
            type: 'object',
            description: ''
          },
          attachment: {
            type: 'object',
            description: ''
          },
          order_lines: {
            type: 'string',
            description: 'Input parameter: The array containing list of line items that are part of this order. Maximum of 1000 line items could be processed in a single order.'
          },
          payment_method_categories: {
            type: 'string',
            description: 'Input parameter: Available payment method categories'
          },
          custom_payment_method_ids: {
            type: 'string',
            description: 'Input parameter: Promo codes - The array could be used to define which of the configured payment options within a payment category (pay_later, pay_over_time, etc.) should be shown for this purchase. Discuss with the delivery manager to know about the promo codes that will be configured for your account. The feature could also be used to provide promotional offers to specific customers (eg: 0% financing). Please be informed that the usage of this feature can have commercial implications.'
          }
        },
        required: ["session_id"]
      }
    },
    handler: post_payments_v1_sessions_session_id
  };
}