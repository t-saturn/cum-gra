'use server';

import { fn_get_applications } from '@/actions/applications/fn_get_applications';
import { fn_get_users } from '@/actions/users/fn_get_users';

export interface RestrictionFormOptions {
  applications: Array<{ id: string; name: string }>;
  users: Array<{ id: string; full_name: string; email: string; dni: string }>;
}

export const fn_get_restriction_form_options = async (): Promise<RestrictionFormOptions> => {
  const [apps, users] = await Promise.all([
    fn_get_applications(1, 100, false),
    fn_get_users(1, 200, { status: 'active', is_deleted: false }),
  ]);
  
  return {
    applications: apps.data.map((a) => ({
      id: a.id,
      name: a.name,
    })),
    users: users.data.map((u) => ({
      id: u.id,
      full_name: `${u.first_name} ${u.last_name}`.trim(), 
      email: u.email,
      dni: u.dni,
    })),
  };
};