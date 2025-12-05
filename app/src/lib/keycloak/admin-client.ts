// src/lib/keycloak/admin-client.ts
import KcAdminClient from '@keycloak/keycloak-admin-client';
import { keycloakConfig } from './config';
import { auth } from '@/lib/auth';

class KeycloakAdminService {
  private static instance: KeycloakAdminService;

  private constructor() {}

  public static getInstance(): KeycloakAdminService {
    if (!KeycloakAdminService.instance) {
      KeycloakAdminService.instance = new KeycloakAdminService();
    }
    return KeycloakAdminService.instance;
  }

  /**
   * Obtiene un cliente autenticado con el token del usuario actual
   */
  public async getAuthenticatedClient(): Promise<KcAdminClient> {
    const session = await auth();
    
    if (!session?.accessToken) {
      throw new Error('No authenticated session found');
    }

    const client = new KcAdminClient({
      baseUrl: keycloakConfig.baseUrl,
      realmName: keycloakConfig.realmName,
    });

    // Usar el access token del usuario autenticado
    client.setAccessToken(session.accessToken as string);

    return client;
  }

  /**
   * Método helper para ejecutar operaciones con el token del usuario
   */
  public async execute<T>(
    operation: (client: KcAdminClient) => Promise<T>,
    errorMessage = 'Keycloak operation failed'
  ): Promise<T> {
    try {
      const client = await this.getAuthenticatedClient();
      return await operation(client);
    } catch (error) {
      console.error(`${errorMessage}:`, error);
      throw error;
    }
  }
}

export const keycloakAdmin = KeycloakAdminService.getInstance();