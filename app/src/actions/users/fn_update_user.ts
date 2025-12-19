'use server';

import { auth } from '@/lib/auth';
import type { UserItem } from '@/types/users';
import { fn_update_keycloak_user } from '@/actions/keycloak/users/fn_update_keycloak_user';
import { fn_enable_keycloak_user } from '@/actions/keycloak/users/fn_enable_keycloak_user';
import { fn_disable_keycloak_user } from '@/actions/keycloak/users/fn_disable_keycloak_user';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface UpdateUserInput {
  email?: string;
  first_name?: string;
  last_name?: string;
  phone?: string;
  status?: 'active' | 'suspended' | 'inactive';
  cod_emp_sgd?: string;
  structural_position_id?: string | null;
  organic_unit_id?: string | null;
  ubigeo_id?: string | null;
  keycloak_id?: string | null;
}

export const fn_update_user = async (id: string, input: UpdateUserInput): Promise<UserItem> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    // 1. Actualizar en Keycloak si tiene keycloak_id
    if (input.keycloak_id) {
      try {
        // Actualizar datos básicos
        await fn_update_keycloak_user({
          keycloakId: input.keycloak_id,
          email: input.email,
          firstName: input.first_name,
          lastName: input.last_name,
        });

        // Actualizar estado (enabled/disabled)
        if (input.status === 'active') {
          await fn_enable_keycloak_user(input.keycloak_id);
        } else if (input.status === 'suspended' || input.status === 'inactive') {
          await fn_disable_keycloak_user(input.keycloak_id);
        }

        console.log('Usuario actualizado en Keycloak');
      } catch (kcError) {
        console.error('Error actualizando en Keycloak (continuando con backend):', kcError);
      }
    }

    // 2. Actualizar en el backend
    const res = await fetch(`${API_BASE_URL}/api/users/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify(input),
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al actualizar usuario: ${res.statusText}`);
    }

    const data: UserItem = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_update_user:', err);
    throw err;
  }
};