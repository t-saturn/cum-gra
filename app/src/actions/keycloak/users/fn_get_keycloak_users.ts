'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';

export interface KeycloakUserSimple {
  id: string;
  username: string;
  email?: string;
  firstName?: string;
  lastName?: string;
  enabled: boolean;
  createdTimestamp?: number;
}

export async function fn_get_keycloak_users(): Promise<KeycloakUserSimple[]> {
  try {
    const users = await keycloakAdmin.execute(
      async (client) => await client.users.find({ max: 1000 }),
      'Failed to fetch Keycloak users'
    );

    return users.map((u) => ({
      id: u.id!,
      username: u.username!,
      email: u.email,
      firstName: u.firstName,
      lastName: u.lastName,
      enabled: u.enabled ?? false,
      createdTimestamp: u.createdTimestamp,
    }));
  } catch (error) {
    console.error('Error fetching Keycloak users:', error);
    throw error;
  }
}