export interface ModuleWithAppDTO {
  id: string;
  item?: string;
  name: string;
  route?: string;
  icon?: string;
  parent_id?: string;
  application_id?: string;
  sort_order: number;
  status: string;
  created_at: string;
  updated_at: string;
  deleted_at?: string;
  deleted_by?: string;
  application_name?: string;
  application_client_id?: string;
  users_count: number;
}

export interface ModulesListResponse {
  data: ModuleWithAppDTO[];
  total: number;
  page: number;
  page_size: number;
}

export interface ModulesStatsResponse {
  total_modules: number;
  active_modules: number;
  deleted_modules: number;
  total_users: number;
}
