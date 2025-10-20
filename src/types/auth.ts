// ====== Tokens / Introspección ======

export type TokenStatus = 'active' | 'expired' | 'revoked' | 'invalid' | string;
export type TokenType = 'access' | 'refresh' | string;

export type TokenDetail = {
  valid: boolean; // true incluso si el backend dice "Token válido (expirado)"
  message: string; // "Token válido" | "Token válido (expirado)" | ...
  subject?: string; // user_id
  issued_at?: string; // ISO
  expires_at?: string; // ISO
  expires_in?: number; // segundos (si viene)
};

export type TokenEnvelope = {
  token_id: string;
  status: TokenStatus; // "active" | "expired" | "revoked" | ...
  token_type: TokenType; // "access" | "refresh"
  token_detail?: TokenDetail; // detalle opcional
};

export type GatewayTokens = {
  access_token?: TokenEnvelope;
  refresh_token?: TokenEnvelope;
};

// ====== Session / Device info ======
export type GatewayDeviceLocation = {
  Country?: string;
  CountryCode?: string;
  Region?: string;
  City?: string;
  Coordinates?: [number, number]; // [lon, lat]
  ISP?: string;
  Organization?: string;
};

export type GatewayDeviceInfo = {
  UserAgent?: string;
  IP?: string;
  DeviceID?: string;
  BrowserName?: string;
  BrowserVersion?: string;
  OS?: string;
  OSVersion?: string;
  DeviceType?: string;
  Timezone?: string;
  Language?: string;
  Location?: GatewayDeviceLocation;
};

export type GatewaySession = {
  session_id: string;
  user_id?: string; // a veces viene en la sesión
  status: 'active' | 'revoked' | 'expired' | string;
  is_active?: boolean;
  created_at?: string;
  last_activity?: string;
  expires_at?: string;
  max_refresh_at?: string;
  device_info?: GatewayDeviceInfo;
  tokens_generated?: string[]; // IDs de tokens asociados generados en la sesión
};

// ====== /auth/me (data) ======
export type GatewayMeData = {
  user_id: string;
  email?: string;
  name?: string;
  dni?: string;
  phone?: string;
  role?: string;
  status?: string; // p. ej. "active"
  session: GatewaySession;

  module_permissions?: string[]; // ["module1","module2"]
  module_restriccions?: string[]; // ["module1"]

  // Campos extra que ahora devuelve el backend (útiles para programar refresh)
  access_expires_at?: string; // ISO
  refresh_expires_at?: string; // ISO
  exp?: number; // epoch seconds (del access)
  remaining_seconds?: number; // segundos hasta exp (si viene)

  // Bloque opcional si el gateway/introsp. lo incluye en /me
  tokens?: GatewayTokens;
};

// ====== Respuestas ======
export type GatewayAPIError =
  | string
  | {
      code?: string;
      message?: string;
      details?: unknown;
    };

// /auth/me y endpoints similares que devuelven {success, data|error}
export type GatewayMeResponse = { success: true; message?: string; data: GatewayMeData } | { success: false; message?: string; error?: GatewayAPIError };

// Estado expuesto por el Provider
export type AuthStatus = 'checking' | 'authenticated' | 'unauthenticated';

// Estado interno opcional
export type SessionState = {
  me: GatewayMeData | null;
};
