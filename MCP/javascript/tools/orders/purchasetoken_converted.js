/**
 * Generate a consumer token
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

export async function post_payments_v1_authorizations_authorization_token_customer_token(authorizationToken, purchase_country, purchase_currency, description, intended_use, locale, billing_address, customer) {
  try {
    const config = getConfig();
    const requestBody = {
      authorizationToken,
      purchase_country,
      purchase_currency,
      description,
      intended_use,
      locale,
      billing_address,
      customer
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

export function createPostPaymentsV1AuthorizationsAuthorizationTokenCustomerTokenTool() {
  return {
    definition: {
      name: 'post-payments-v1-authorizations-authorization-token-customer-token',
      description: 'Generate a consumer token',
      inputSchema: {
        type: 'object',
        properties: {
          authorizationToken: {
            type: 'string',
            description: ''
          },
          purchase_country: {
            type: 'string',
            description: 'Input parameter: ISO 3166 alpha-2 purchase country.'
          },
          purchase_currency: {
            type: 'string',
            description: 'Input parameter: ISO 4217 purchase currency.'
          },
          description: {
            type: 'string',
            description: 'Input parameter: Description of the purpose of the token.'
          },
          intended_use: {
            type: 'string',
            description: 'Input parameter: Intended use for the token.'
          },
          locale: {
            type: 'string',
            description: 'Input parameter: RFC 1766 customer's locale.'
          },
          billing_address: {
            type: 'object',
            description: ''
          },
          customer: {
            type: 'object',
            description: ''
          }
        },
        required: ["authorizationToken", "purchase_country", "purchase_currency", "description", "intended_use", "locale"]
      }
    },
    handler: post_payments_v1_authorizations_authorization_token_customer_token
  };
}