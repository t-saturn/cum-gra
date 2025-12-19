'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';
import type UserRepresentation from '@keycloak/keycloak-admin-client/lib/defs/userRepresentation';

export interface CreateKeycloakUserInput {
  email: string;
  firstName: string;
  lastName: string;
  username: string; // Generalmente el DNI o email
  enabled?: boolean;
}

export async function fn_create_keycloak_user(input: CreateKeycloakUserInput) {
  try {
    const userConfig: UserRepresentation = {
      username: input.username,
      email: input.email,
      firstName: input.firstName,
      lastName: input.lastName,
      enabled: input.enabled ?? true,
      emailVerified: false,
      credentials: [
        {
          type: 'password',
          value: generateTemporaryPassword(), // Contraseña temporal
          temporary: true, // Forzar cambio en primer login
        },
      ],
    };

    const result = await keycloakAdmin.execute(
      async (client) => {
        return await client.users.create(userConfig);
      },
      'Failed to create Keycloak user'
    );

    return {
      success: true,
      userId: result.id,
    };
  } catch (error) {
    console.error('Error creating Keycloak user:', error);
    throw error;
  }
}

// Generar contraseña temporal segura
function generateTemporaryPassword(): string {
  const length = 12;
  const charset = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*';
  let password = '';
  for (let i = 0; i < length; i++) {
    password += charset.charAt(Math.floor(Math.random() * charset.length));
  }
  return password;
}