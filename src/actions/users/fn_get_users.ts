'use server';

import { UsersListResponse } from '@/types/users';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:9191';

export const fn_get_users = async (page = 1, pageSize = 20): Promise<UsersListResponse> => {
  const url = `${API_BASE_URL}/users?page=${page}&page_size=${pageSize}`;

  try {
    const res = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      // opcional: revalidate cada cierto tiempo si usas caching
      // next: { revalidate: 10 },
    });

    if (!res.ok) {
      throw new Error(`Error fetching users: ${res.statusText}`);
    }

    const data = (await res.json()) as UsersListResponse;
    return data;
  } catch (error) {
    console.error('Error al obtener usuarios:', error);
    throw error;
  }
};
