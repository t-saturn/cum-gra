export type UUID = string;

export type UserMinimalDTO = {
  id: UUID;
  first_name: string | null;
  last_name: string | null;
  email: string;
  dni: string;
};

export type AppMinimalDTO = {
  id: UUID;
  name: string;
  client_id: string;
};

export type RoleMinimalDTO = {
  id: UUID;
  name: string;
};

export type UserAppRoleDTO = {
  application: AppMinimalDTO;
  role: RoleMinimalDTO;
};

export type RoleAssignmentsDTO = {
  user: UserMinimalDTO;
  assignments: UserAppRoleDTO[];
};

export interface RolesAssignmentsResponse {
  data: RoleAssignmentsDTO[];
  total: number;
  page: number;
  page_size: number;
}

export interface UserRoleOverallStatsResponse {
  total_users: number;
  admin_users: number;
  users_with_roles: number;
  users_without_roles: number;
}
