export interface UbigeoItem {
  id: string;
  ubigeo_code: string;
  inei_code: string;
  department: string;
  province: string;
  district: string;
  created_at: string;
  updated_at: string;
}

export interface UbigeosListResponse {
  data: UbigeoItem[];
  total: number;
  page: number;
  page_size: number;
}

export interface UbigeosStatsResponse {
  total_ubigeos: number;
  total_departments: number;
  total_provinces: number;
  total_districts: number;
}

export interface DepartmentItem {
  name: string;
}

export interface ProvinceItem {
  name: string;
  department: string;
}

export interface DistrictItem {
  id: string;
  name: string;
  department: string;
  province: string;
  ubigeo_code: string;
}
