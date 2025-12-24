export type RestrictionType = 'block' | 'limit' | 'read_only';
export type PermissionLevel = 'read' | 'write' | 'execute' | 'delete' | 'admin';

export interface UserRestrictionItem {
  id: string;
  user_id: string;
  module_id: string;
  application_id: string;
  restriction_type: RestrictionType;
  max_permission_level?: PermissionLevel | null;
  reason?: string | null;
  expires_at?: string | null;
  created_at: string;
  created_by?: string | null;
  updated_at: string;
  updated_by?: string | null;
  is_deleted: boolean;
  deleted_at?: string | null;
  deleted_by?: string | null;

  // Relaciones expandidas
  user_email?: string;
  user_full_name?: string;
  module_name?: string;
  module_route?: string;
  application_name?: string;
  application_client_id?: string;
}

export interface UserRestrictionsListResponse {
  data: UserRestrictionItem[];
  total: number;
  page: number;
  page_size: number;
}

export interface UserRestrictionsStatsResponse {
  total_restrictions: number;
  active_restrictions: number;
  restricted_users: number;
  deleted_restrictions: number;
}

export interface CreateRestrictionInput {
  user_id: string;
  module_id: string;
  application_id: string;
  restriction_type: RestrictionType;
  max_permission_level?: PermissionLevel;
  reason?: string;
  expires_at?: string; // ISO format
}

export interface UpdateRestrictionInput {
  restriction_type?: RestrictionType;
  max_permission_level?: PermissionLevel;
  reason?: string;
  expires_at?: string | null; // null para quitar expiración
}

export interface BulkCreateRestrictionsInput {
  user_id: string;
  application_id: string;
  module_ids: string[];
  restriction_type: RestrictionType;
  max_permission_level?: PermissionLevel;
  reason?: string;
  expires_at?: string;
}

export interface BulkCreateResponse {
  created: number;
  skipped: number;
  failed: number;
  details: UserRestrictionItem[];
}