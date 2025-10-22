export type BasicModule = {
  id: string;
  name: string;
  icon: string;
};

export type BasicApp = {
  id: string;
  name: string;
  client_id: string;
};

export type BasicRole = {
  id: string;
  name: string;
};

export type RolesAppData = {
  apps: Array<BasicApp>;
  roles: Array<BasicRole>;
  modules: Array<BasicModule>;
};

export interface RolesAppsResponse {
  data: RolesAppData;
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
