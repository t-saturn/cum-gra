import { query } from '@/lib/database';

export type AppRole = { id: string; name: string };

export async function getUserRoleForCurrentApp(userId: string): Promise<AppRole | null> {
  const clientId = process.env.APP_CLIENT_ID;
  if (!clientId) {
    console.error('APP_CLIENT_ID no configurado');
    return null;
  }

  const { rows } = await query<{ id: string; name: string; is_deleted: boolean }>(
    `
    SELECT ar.id, ar.name, ar.is_deleted
    FROM application_roles ar
    JOIN applications a ON a.id = ar.application_id
    JOIN user_application_roles uar ON uar.application_role_id = ar.id
    WHERE uar.user_id = $1
      AND a.client_id = $2
    LIMIT 1
    `,
    [userId, clientId],
  );

  return rows[0] ?? null;
}
