'use server';

import { auth } from '@/lib/auth';
import { fn_get_applications } from '@/actions/applications/fn_get_applications';
import { fn_get_modules } from './fn_get_modules';

export interface ModuleFormOptions {
  applications: Array<{ id: string; name: string; client_id: string }>;
  modules: Array<{ id: string; name: string; route: string; application_id?: string | null }>;
}

export const fn_get_module_form_options = async (applicationId?: string): Promise<ModuleFormOptions> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    // Obtener aplicaciones
    const applicationsRes = await fn_get_applications(1, 100, false);

    // Obtener módulos filtrados por aplicación si se proporciona
    const filters = applicationId ? { application_id: applicationId } : {};
    const modulesRes = await fn_get_modules(1, 200, filters);

    return {
      applications: applicationsRes.data.map((a) => ({
        id: a.id,
        name: a.name,
        client_id: a.client_id,
      })),
      modules: modulesRes.data.map((m) => ({
        id: m.id,
        name: m.name,
        route: m.route,
        application_id: m.application_id,
      })),
    };
  } catch (err) {
    console.error('Error en fn_get_module_form_options:', err);
    throw err;
  }
};