export interface ApplicationRoleItem {
  id: string;
  name: string;
  description?: string | null;
  application_id: string;
  created_at: string;
  updated_at: string;
  is_deleted: boolean;
  deleted_at?: string | null;
  deleted_by?: string | null;
  
  // Relación expandida
  application?: {
    id: string;
    name: string;
    client_id: string;
  } | null;
  
  // Contadores
  modules_count?: number;
  users_count?: number;
}

export interface ApplicationRolesListResponse {
  data: ApplicationRoleItem[];
  total: number;
  page: number;
  page_size: number;
}

export interface ApplicationRolesStatsResponse {
  total_roles: number;
  active_roles: number;
  deleted_roles: number;
  roles_with_modules: number;
  roles_with_users: number;
}