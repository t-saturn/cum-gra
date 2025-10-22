export type AppMinimalDTO = {
  id: string;
  name: string;
  client_id: string;
};

export type RoleMinimalDTO = {
  id: string;
  name: string;
};

export type ModuleMinimalDTO = {
  id: string;
  name: string;
  icon: string | null;
};

export type RoleAppModulesItemDTO = {
  role: RoleMinimalDTO;
  app: AppMinimalDTO;
  app_modules: ModuleMinimalDTO[];
  role_modules: ModuleMinimalDTO[];
};

export interface RolesAppsResponse {
  data: RoleAppModulesItemDTO[];
  total: number;
  page: number;
  page_size: number;
}

export interface RolesAppsStatsResponse {
  total_roles: number;
  active_roles: number;
  admin_roles: number;
  assigned_users: number;
}
