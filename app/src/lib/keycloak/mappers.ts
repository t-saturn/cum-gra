import type { KeycloakClientRepresentation, KeycloakApplication } from '@/types/keycloak/clients';

/**
 * Convierte un cliente de Keycloak a ApplicationItem compatible
 */
export function mapKeycloakClientToApplication(client: KeycloakClientRepresentation): KeycloakApplication {
  // Extraer dominio de las URLs
  const domain = extractDomain(client.baseUrl || client.rootUrl || client.redirectUris?.[0] || '');
  
  // Determinar estado basado en si está habilitado
  const status = client.enabled ? 'active' : 'inactive';
  
  // Keycloak no guarda created_at en el cliente, usar fecha actual o desde attributes
  const created_at = client.attributes?.created_at || new Date().toISOString();

  return {
    id: client.id || '',
    name: client.name || client.clientId || 'Sin nombre',
    client_id: client.clientId || '',
    domain,
    description: client.description,
    status,
    created_at,
    protocol: client.protocol || 'openid-connect',
    redirect_uris: client.redirectUris || [],
    enabled: client.enabled ?? true,
  };
}

/**
 * Extrae el dominio de una URL
 */
function extractDomain(url: string): string {
  if (!url) return '—';
  
  try {
    const urlObj = new URL(url);
    return urlObj.hostname;
  } catch {
    return url;
  }
}

/**
 * Calcula estadísticas de los clientes
 */
export function calculateClientStats(clients: KeycloakClientRepresentation[]) {
  return {
    total_clients: clients.length,
    enabled_clients: clients.filter(c => c.enabled).length,
    oauth_clients: clients.filter(c => c.protocol === 'openid-connect').length,
    saml_clients: clients.filter(c => c.protocol === 'saml').length,
  };
}