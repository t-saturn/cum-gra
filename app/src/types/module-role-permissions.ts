// Tipos de permiso disponibles
export type PermissionType = 'read' | 'write' | 'execute' | 'delete' | 'admin';

// Item de permiso de módulo-rol
export interface ModuleRolePermissionItem {
  id: string;
  module_id: string;
  application_role_id: string;
  permission_type: PermissionType;
  created_at: string;
  is_deleted: boolean;
  deleted_at: string | null;
  deleted_by: string | null;
  module_name: string;
  module_route: string;
  role_name: string;
  application_name: string;
  application_client_id: string;
}

export interface ModuleRolePermissionsListResponse {
  data: ModuleRolePermissionItem[];
  total: number;
  page: number;
  page_size: number;
}

export interface ModuleRolePermissionsStatsResponse {
  total_permissions: number;
  active_permissions: number;
  deleted_permissions: number;
  unique_modules: number;
  unique_roles: number;
  permissions_by_type: {
    read: number;
    write: number;
    execute: number;
    delete: number;
    admin: number;
  };
}

export interface CreateModuleRolePermissionInput {
  module_id: string;
  application_role_id: string;
  permission_type: PermissionType;
}

export interface UpdateModuleRolePermissionInput {
  permission_type: PermissionType;
}

export interface BulkAssignPermissionsInput {
  application_role_id: string;
  module_ids: string[];
  permission_type: PermissionType;
}

export interface BulkAssignPermissionsResponse {
  created: number;
  skipped: number;
  failed: number;
  details: ModuleRolePermissionItem[];
}

export interface ModuleRolePermissionFilters {
  module_id?: string;
  role_id?: string;
  is_deleted?: boolean;
}