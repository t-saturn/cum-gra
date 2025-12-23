'use server';

import { auth } from '@/lib/auth';
import { fn_get_applications } from '@/actions/applications/fn_get_applications';
import { fn_get_users } from '@/actions/users/fn_get_users';
import { fn_get_application_roles } from '@/actions/application-roles/fn_get_application_roles';

export interface AssignmentFormOptions {
  applications: Array<{ id: string; name: string; client_id: string }>;
  users: Array<{ id: string; email: string; full_name: string }>;
  roles: Array<{ id: string; name: string; application_id: string }>;
}

export const fn_get_assignment_form_options = async (
  applicationId?: string
): Promise<AssignmentFormOptions> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    // Obtener aplicaciones y usuarios
    const [applicationsRes, usersRes] = await Promise.all([
      fn_get_applications(1, 100, false),
      fn_get_users(1, 200),
    ]);

    // Obtener roles filtrados por aplicación si se proporciona
    const rolesFilters = applicationId ? { application_id: applicationId } : {};
    const rolesRes = await fn_get_application_roles(1, 200, rolesFilters);

    return {
      applications: applicationsRes.data.map((a) => ({
        id: a.id,
        name: a.name,
        client_id: a.client_id,
      })),
      users: usersRes.data.map((u) => ({
        id: u.id,
        email: u.email,
        full_name: `${u.first_name} ${u.last_name}`,
      })),
      roles: rolesRes.data.map((r) => ({
        id: r.id,
        name: r.name,
        application_id: r.application_id,
      })),
    };
  } catch (err) {
    console.error('Error en fn_get_assignment_form_options:', err);
    throw err;
  }
};