'use server';

import { auth } from '@/lib/auth';
import { fn_get_organic_units } from '@/actions/units/fn_get_organic_units';
import { fn_get_positions } from '@/actions/positions/fn_get_positions';
import { fn_get_ubigeos } from '@/actions/ubigeos/fn_get_ubigeos';

export interface UserFormOptions {
  positions: Array<{ id: string; name: string; code: string }>;
  units: Array<{ id: string; name: string; acronym: string }>;
  ubigeos: Array<{ id: string; ubigeo_code: string; department: string; province: string; district: string }>;
}

export const fn_get_user_form_options = async (): Promise<UserFormOptions> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    // Obtener todas las opciones necesarias para el formulario
    const [positionsRes, unitsRes, ubigeosRes] = await Promise.all([
      fn_get_positions(1, 100, false),
      fn_get_organic_units(1, 100, false),
      fn_get_ubigeos(1, 100),
    ]);

    return {
      positions: positionsRes.data.map((p) => ({
        id: p.id,
        name: p.name,
        code: p.code,
      })),
      units: unitsRes.data.map((u) => ({
        id: u.id,
        name: u.name,
        acronym: u.acronym,
      })),
      ubigeos: ubigeosRes.data.map((ub) => ({
        id: ub.id,
        ubigeo_code: ub.ubigeo_code,
        department: ub.department,
        province: ub.province,
        district: ub.district,
      })),
    };
  } catch (err) {
    console.error('Error en fn_get_user_form_options:', err);
    throw err;
  }
};