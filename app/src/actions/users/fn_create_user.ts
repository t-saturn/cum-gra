'use server';

import { auth } from '@/lib/auth';
import type { UserItem } from '@/types/users';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface CreateUserInput {
  email: string;
  dni: string;
  password: string; // AHORA ES REQUERIDO
  first_name: string;
  last_name: string;
  phone?: string;
  status?: 'active' | 'suspended' | 'inactive';
  cod_emp_sgd?: string;
  structural_position_id?: string;
  organic_unit_id?: string;
  ubigeo_id?: string;
}

export const fn_create_user = async (input: CreateUserInput): Promise<UserItem> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    // El backend ahora se encarga de TODO (creación en Keycloak + BD)
    const res = await fetch(`${API_BASE_URL}/api/users`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify(input),
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