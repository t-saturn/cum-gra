export interface AdminUser {
  full_name: string;
  dni: string;
  email: string;
}

export interface ApplicationItem {
  id: string;
  name: string;
  client_id: string;
  domain: string;
  logo?: string;
  description?: string;
  status: string;
  created_at: string;
  updated_at: string;
  is_deleted: boolean;
  deleted_at?: string;
  deleted_by?: string;
  admins?: AdminUser[] | null;
  users_count: number;
}

export interface ApplicationsListResponse {
  data: ApplicationItem[];
  total: number;
  page: number;
  page_size: number;
}

export type ApplicationsStatsResponse = {
  total_applications: number;
  active_applications: number;
  total_users: number;
};
