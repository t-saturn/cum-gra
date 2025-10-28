export type OrganicUnitItemDTO = {
  id: string;
  name: string;
  acronym: string;
  brand?: string;
  description?: string;
  parent_id?: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
  is_deleted: boolean;
  deleted_at?: string;
  deleted_by?: string;
  users_count: number;
};

export type OrganicUnitsListResponse = {
  data: OrganicUnitItemDTO[];
  total: number;
  page: number;
  page_size: number;
};

export type OrganicUnitsStatsResponse = {
  total_organic_units: number;
  active_organic_units: number;
  deleted_organic_units: number;
  total_employees: number;
};
