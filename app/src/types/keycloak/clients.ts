// Tipos nativos de Keycloak
export interface KeycloakClientRepresentation {
  id?: string;
  clientId?: string;
  name?: string;
  description?: string;
  rootUrl?: string;
  baseUrl?: string;
  enabled?: boolean;
  clientAuthenticatorType?: string;
  secret?: string;
  redirectUris?: string[];
  webOrigins?: string[];
  protocol?: string;
  attributes?: Record<string, string>;
  // ... otros campos que Keycloak devuelve
}

// Tipo adaptado a tu ApplicationItem
export interface KeycloakApplication {
  id: string;
  name: string;
  client_id: string;
  domain: string;
  description?: string;
  status: 'active' | 'inactive' | 'development';
  created_at: string;
  protocol: string;
  redirect_uris: string[];
  enabled: boolean;
}

export interface KeycloakApplicationsStats {
  total_clients: number;
  enabled_clients: number;
  oauth_clients: number;
  saml_clients: number;
}