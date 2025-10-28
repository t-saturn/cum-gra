export interface SimpleStructuralPosition {
  id: string;
  name: string;
  code: string;
  level?: number;
}

export interface SimpleOrganicUnit {
  id: string;
  name: string;
  acronym: string;
  parent_id?: string;
}

export interface UserListItem {
  id: string;
  email: string;
  first_name?: string;
  last_name?: string;
  phone?: string;
  dni: string;
  status: string;
  created_at: string;
  updated_at: string;
  is_deleted: boolean;
  deleted_at?: string;
  deleted_by?: string;
  organic_unit?: SimpleOrganicUnit;
  structural_position?: SimpleStructuralPosition;
}

export interface UsersListResponse {
  data: UserListItem[];
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
