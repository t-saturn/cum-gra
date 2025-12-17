export interface ApplicationItem {
  id: string;
  name: string;
  client_id: string;
  client_secret?: string;
  domain: string;
  logo?: string | null;
  description?: string | null;
  status: 'active' | 'inactive' | 'development';
  created_at: string;
  updated_at: string;
  is_deleted: boolean;
  deleted_at?: string | null;
  deleted_by?: string | null;
  admins?: Array<{
    full_name: string;
    dni: string;
    email: string;
  }>;
  users_count: number;
  keycloak_id?: string | null; // ID interno de Keycloak
}

export interface ApplicationsListResponse {
  data: ApplicationItem[];
  total: number;
  page: number;
  page_size: number;
}

export interface ApplicationsStatsResponse {
  total_applications: number;
  active_applications: number;
  deleted_applications: number;
  total_users: number;
}