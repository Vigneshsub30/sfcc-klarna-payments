/**
 * Cancel an existing authorization
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

export async function delete_payments_v1_authorizations_authorization_token(authorizationToken) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (authorizationToken) params.append("authorizationToken", authorizationToken);
    const queryString = params.toString();
    const finalUrl = queryString ? `${url}?${queryString}` : url;
    
    const url = `${config.baseURL}/api/unknown`;
    
    const response = await fetch(finalUrl, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${config.bearerToken}`,
        'Accept': 'application/json'
      }
    });
    
    if (!response.ok) {
      return `Failed to format JSON: ${response.status} ${response.statusText}`;
    }
    
    try {
      const result = await response.json();
      return JSON.stringify(result, null, 2);
    } catch (e) {
      return await response.text();
    }
    
  } catch (error) {
    return `Request failed: ${error.message}`;
  }
}

export function createDeletePaymentsV1AuthorizationsAuthorizationTokenTool() {
  return {
    definition: {
      name: 'delete-payments-v1-authorizations-authorization-token',
      description: 'Cancel an existing authorization',
      inputSchema: {
        type: 'object',
        properties: {
          authorizationToken: {
            type: 'string',
            description: ''
          }
        },
        required: ["authorizationToken"]
      }
    },
    handler: delete_payments_v1_authorizations_authorization_token
  };
}