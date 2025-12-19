export interface UserItem {
  id: string;
  email: string;
  dni: string;
  first_name: string;
  last_name: string;
  phone?: string | null;
  status: 'active' | 'suspended' | 'inactive';
  cod_emp_sgd?: string | null;
  structural_position_id?: string | null;
  organic_unit_id?: string | null;
  ubigeo_id?: string | null;
  created_at: string;
  updated_at: string;
  is_deleted: boolean;
  deleted_at?: string | null;
  deleted_by?: string | null;
  keycloak_id?: string | null; // Para sincronización
  
  // Relaciones expandidas
  structural_position?: {
    id: string;
    name: string;
    code: string;
    level?: number | null;
  } | null;
  
  organic_unit?: {
    id: string;
    name: string;
    acronym: string;
    parent_id?: string | null;
  } | null;
  
  ubigeo?: {
    id: string;
    ubigeo_code: string;
    department: string;
    province: string;
    district: string;
  } | null;
}

export interface UsersListResponse {
  data: UserItem[];
  total: number;
  page: number;
  page_size: number;
}

export interface UsersStatsResponse {
  total_users: number;
  active_users: number;
  suspended_users: number;
  new_users_last_month: number;
}