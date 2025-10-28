import type { AppMinimalDTO, RoleMinimalDTO, UserMinimalDTO } from '@/types/roles_assignments';
import type { ModuleMinimalDTO } from '@/types/roles_app';

export type UserAppAssignmentDTO = {
  app: AppMinimalDTO;
  role?: RoleMinimalDTO | null;
  modules: ModuleMinimalDTO[] | null;
  modules_restrict: ModuleMinimalDTO[] | null;
};

export type RoleRestrictDTO = {
  user: UserMinimalDTO;
  apps: UserAppAssignmentDTO[];
};

export interface RolesRestrictResponse {
  data: RoleRestrictDTO[];
  total: number;
  page: number;
  page_size: number;
}

export interface UsersRestrictionsStatsResponse {
  total_restrictions: number;
  active_restrictions: number;
  restricted_users: number;
  deleted_restrictions: number;
}
