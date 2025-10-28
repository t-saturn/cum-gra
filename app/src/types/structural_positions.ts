export interface StructuralPositionItem {
  id: string;
  name: string;
  code: string;
  level?: number | null;
  description?: string | null;
  is_active: boolean;
  created_at: string;
  updated_at: string;
  is_deleted: boolean;
  deleted_at?: string | null;
  deleted_by?: string | null;
  users_count: number;
}

export interface StructuralPositionsListResponse {
  data: StructuralPositionItem[];
  total: number;
  page: number;
  page_size: number;
}
