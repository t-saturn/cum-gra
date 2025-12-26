'use server';

import { auth } from '@/lib/auth';
import { fn_get_all_organic_units } from '@/actions/units/fn_get_all_organic_units';
import { fn_get_all_positions } from '@/actions/positions/fn_get_all_positions';

export interface UserFormOptions {
  positions: Array<{ id: string; name: string; code: string }>;
  units: Array<{ id: string; name: string; acronym: string }>;
}

export const fn_get_user_form_options = async (): Promise<UserFormOptions> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const [positions, units] = await Promise.all([
      fn_get_all_positions(true),
      fn_get_all_organic_units(true),
    ]);

    return {
      positions: positions.map((p) => ({
        id: p.id,
        name: p.name,
        code: p.code,
      })),
      units: units.map((u) => ({
        id: u.id,
        name: u.name,
        acronym: u.acronym,
      })),
    };
  } catch (err) {
    console.error('Error en fn_get_user_form_options:', err);
    throw err;
  }
};