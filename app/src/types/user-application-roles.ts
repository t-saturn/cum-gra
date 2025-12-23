export interface UserApplicationRoleItem {
  id: string;
  user_id: string;
  application_id: string;
  application_role_id: string;
  granted_at: string;
  granted_by?: string | null;
  revoked_at?: string | null;
  revoked_by?: string | null;
  revoke_reason?: string | null;
  is_deleted: boolean;
  deleted_at?: string | null;
  deleted_by?: string | null;
  created_at: string;
  updated_at: string;
  
  // Relaciones expandidas
  user_email?: string | null;
  user_full_name?: string | null;
  application_name?: string | null;
  application_client_id?: string | null;
  role_name?: string | null;
  granted_by_email?: string | null;
  revoked_by_email?: string | null;
}

export interface UserApplicationRolesListResponse {
  data: UserApplicationRoleItem[];
  total: number;
  page: number;
  page_size: number;
}

export interface UserApplicationRolesStatsResponse {
  total_assignments: number;
  active_assignments: number;
  revoked_assignments: number;
  deleted_assignments: number;
  users_with_roles: number;
}

export interface BulkAssignResponse {
  created: number;
  skipped: number;
  failed: number;
  details: UserApplicationRoleItem[];
}