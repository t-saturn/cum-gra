'use server';

import { auth } from '@/lib/auth';
import type { UserItem } from '@/types/users';
import { fn_create_keycloak_user } from '@/actions/keycloak/users/fn_create_keycloak_user';
import { fn_get_keycloak_user_by_username } from '@/actions/keycloak/users/fn_get_keycloak_user_by_username';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface CreateUserInput {
  email: string;
  dni: string;
  first_name: string;
  last_name: string;
  phone?: string;
  status?: 'active' | 'suspended' | 'inactive';
  cod_emp_sgd?: string;
  structural_position_id?: string;
  organic_unit_id?: string;
  ubigeo_id?: string;
  sync_to_keycloak?: boolean;
}

export const fn_create_user = async (input: CreateUserInput): Promise<UserItem> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    let keycloakId: string | undefined;

    // 1. Crear en Keycloak si sync_to_keycloak es true
    if (input.sync_to_keycloak !== false) {
      try {
        // Verificar si ya existe en Keycloak
        const existingUser = await fn_get_keycloak_user_by_username(input.dni);
        
        if (existingUser) {
          keycloakId = existingUser.id;
          console.log('Usuario ya existe en Keycloak:', keycloakId);
        } else {
          const keycloakResult = await fn_create_keycloak_user({
            username: input.dni,
            email: input.email,
            firstName: input.first_name,
            lastName: input.last_name,
            enabled: input.status === 'active',
          });
          keycloakId = keycloakResult.userId;
          console.log('Usuario creado en Keycloak:', keycloakId);
        }
      } catch (kcError) {
        console.error('Error creando en Keycloak (continuando con backend):', kcError);
      }
    }

    // 2. Crear en el backend
    const res = await fetch(`${API_BASE_URL}/api/users`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify({
        ...input,
        keycloak_id: keycloakId,
      }),
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al crear usuario: ${res.statusText}`);
    }

    const data: UserItem = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_create_user:', err);
    throw err;
  }
};