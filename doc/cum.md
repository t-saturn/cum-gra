```sql
Table user_sessions {
  id uuid [primary key]
  user_id uuid [ref: > users.id]
  session_token varchar(512) [unique, not null]
  refresh_token varchar(512) [unique, not null]
  device_info text
  ip_address inet
  user_agent text
  is_active boolean [default: true]
  expires_at timestamp [not null]
  created_at timestamp [default: `now()`]
  last_activity_at timestamp [default: `now()`]

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
}

Table applications {
  id uuid [primary key]
  name varchar(100) [not null]
  client_id varchar(255) [unique, not null]
  client_secret varchar(255) [not null]
  domain varchar(255) [not null]
  logo varchar(255)
  description text
  callback_urls text[] [note: 'Array of allowed callback URLs']
  scopes text[] [note: 'Array of available scopes']
  is_first_party boolean [default: false, note: 'True for internal applications']
  status enum('active', 'suspended') [default: 'active']
  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]
}

Table application_roles {
  id uuid [primary key]
  name varchar(100) [not null]
  description text
  application_id uuid [ref: > applications.id]
  permissions text[] [note: 'Array of permissions for this role']
  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]
}

Table user_application_roles {
  id uuid [primary key]
  user_id uuid [ref: > users.id]
  application_id uuid [ref: > applications.id]
  application_role_id uuid [ref: > application_roles.id]
  granted_at timestamp [default: `now()`]
  granted_by uuid [ref: > users.id]
  revoked_at timestamp
  revoked_by uuid [ref: > users.id]

  Note: "Relaciona un usuario con uno o varios roles de una aplicación específica"
}

Table modules {
  id uuid [primary key]
  item varchar(100)
  name varchar(100) [not null]
  label varchar(100)
  route varchar(255)
  icon varchar(100)
  parent_id uuid [ref: > modules.id]
  application_id uuid [ref: > applications.id]
  sort_order int [default: 0]
  is_menu_item boolean [default: true]
  status enum('active', 'inactive') [default: 'active']
  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]
}

Table module_role_permissions {
  id uuid [primary key]
  module_id uuid [ref: > modules.id]
  application_role_id uuid [ref: > application_roles.id]
  permission_type enum('denied', 'read', 'write', 'admin') [not null]
  created_at timestamp [default: `now()`]

  Note: "Define permisos específicos de roles sobre módulos"
}

Table oauth_tokens {
  id uuid [primary key]
  user_id uuid [ref: > users.id]
  application_id uuid [ref: > applications.id]
  access_token varchar(512) [unique, not null]
  refresh_token varchar(512) [unique]
  token_type varchar(50) [default: 'Bearer']
  scopes text[] [note: 'Granted scopes']
  expires_at timestamp [not null]
  created_at timestamp [default: `now()`]
  revoked_at timestamp [note: 'When token was revoked']
}

Table user_permissions {
  id uuid [primary key]
  user_id uuid [ref: > users.id]
  application_id uuid [ref: > applications.id]
  scopes text[] [note: 'Granted scopes for this app']
  granted_at timestamp [default: `now()`]
  granted_by uuid [ref: > users.id]
  revoked_at timestamp
  revoked_by uuid [ref: > users.id]
}

Table structural_positions {
  id uuid [primary key]
  name varchar(255) [not null]
  code varchar(50) [unique, not null]
  level varchar(50)
  description text
  is_active boolean [default: true]
  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]
}

Table organic_units {
  id uuid [primary key]
  name varchar(255) [not null]
  acronym varchar(20)
  brand varchar(100)
  level varchar(50)
  description text
  parent_id uuid [ref: > organic_units.id]
  is_active boolean [default: true]
  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]
}

Table password_history {
  id uuid [primary key]
  user_id uuid [ref: > users.id]
  previous_password_hash varchar(255) [not null]
  changed_at timestamp [default: `now()`]
  changed_by uuid [ref: > users.id]

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]
}

Table password_resets {
  id uuid [primary key]
  user_id uuid [ref: > users.id]
  token varchar(255) [unique, not null]
  expires_at timestamp [not null]
  used_at timestamp
  ip_address inet
  user_agent text
  created_at timestamp [default: `now()`]
}

Table two_factor_secrets {
  id uuid [primary key]
  user_id uuid [ref: > users.id]
  secret varchar(255) [not null]
  backup_codes text[] [note: 'Array of backup codes']
  created_at timestamp [default: `now()`]
  last_used_at timestamp
  recovery_codes_used int [default: 0]
}


Table audit_logs {
  id uuid [primary key]
  user_id uuid [ref: > users.id]
  application_id uuid [ref: > applications.id]
  action varchar(100) [not null, note: 'login, logout, token_refresh, permission_grant, role_assign, etc.']
  resource_type varchar(50) [note: 'user, role, module, etc.']
  resource_id uuid [note: 'ID of affected resource']
  ip_address inet
  user_agent text
  details jsonb [note: 'Additional context data']
  created_at timestamp [default: `now()`]
}

Table application_settings {
  id uuid [primary key]
  application_id uuid [ref: > applications.id]
  setting_key varchar(100) [not null]
  setting_value text
  data_type enum('string', 'number', 'boolean', 'json') [default: 'string']
  description text
  is_public boolean [default: false]
  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]

  Note: "Configuraciones específicas por aplicación"
}

Table user_preferences {
  id uuid [primary key]
  user_id uuid [ref: > users.id]
  application_id uuid [ref: > applications.id]
  preference_key varchar(100) [not null]
  preference_value text
  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]

  Note: "Preferencias del usuario por aplicación (tema, idioma, etc.)"
}


Table authentication_methods {
  id uuid [primary key]
  name varchar(50) [not null, unique] // 'email_password', 'google_oauth', 'face_id', 'microsoft_oauth', etc.
  display_name varchar(100) [not null] // 'Email y Contraseña', 'Google', 'Face ID', etc.
  provider varchar(50) [note: 'google, microsoft, apple, internal, biometric']
  is_active boolean [default: true]
  configuration jsonb [note: 'Configuración específica del método (client_id, endpoints, etc.)']
  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]
}

Table application_auth_methods {
  id uuid [primary key]
  application_id uuid [ref: > applications.id]
  authentication_method_id uuid [ref: > authentication_methods.id]
  is_enabled boolean [default: true]
  is_primary boolean [default: false] // Método principal de la app
  configuration jsonb [note: 'Config específica para esta app']
  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]
}

Table user_authentication_credentials {
  id uuid [primary key]
  user_id uuid [ref: > users.id]
  authentication_method_id uuid [ref: > authentication_methods.id]

  // Para email/password
  email varchar(255) [null]
  password_hash varchar(255) [null]

  // Para OAuth (Google, Microsoft, etc.)
  provider_user_id varchar(255) [null] // ID del usuario en el proveedor
  provider_email varchar(255) [null]
  provider_data jsonb [null] // Datos adicionales del proveedor

  // Para biométricos (Face ID, Touch ID)
  biometric_template text [null] // Template encriptado
  device_id varchar(255) [null] // ID del dispositivo registrado

  // Metadatos
  is_verified boolean [default: false]
  is_primary boolean [default: false] // Método principal del usuario
  last_used_at timestamp
  verification_token varchar(255) [null]
  verification_expires_at timestamp [null]

  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]

  Note: "Almacena las credenciales específicas de cada método de autenticación por usuario"
}

Table authentication_history {
  id uuid [primary key]
  user_id uuid [ref: > users.id]
  application_id uuid [ref: > applications.id]
  authentication_method_id uuid [ref: > authentication_methods.id]
  user_authentication_credential_id uuid [ref: > user_authentication_credentials.id]

  success boolean [not null]
  failure_reason varchar(100) [null] // 'invalid_credentials', 'biometric_failed', 'oauth_error', etc.
  ip_address inet
  user_agent text
  device_fingerprint varchar(255) [null]
  location_data jsonb [null] // Geolocalización si está disponible

  // Datos específicos del método
  method_specific_data jsonb [null] // Datos adicionales según el método

  created_at timestamp [default: `now()`]

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]
}

Table oauth_provider_tokens {
  id uuid [primary key]
  user_authentication_credential_id uuid [ref: > user_authentication_credentials.id]
  provider varchar(50) [not null] // 'google', 'microsoft', etc.
  access_token text [not null]
  refresh_token text [null]
  token_type varchar(50) [default: 'Bearer']
  scope text [null]
  expires_at timestamp [not null]
  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]
}

Table biometric_devices {
  id uuid [primary key]
  user_id uuid [ref: > users.id]
  device_id varchar(255) [not null] // Identificador único del dispositivo
  device_name varchar(100) [not null] // Nombre del dispositivo
  device_type enum('ios', 'android', 'windows', 'macos') [not null]
  biometric_type enum('face_id', 'touch_id', 'fingerprint', 'voice') [not null]
  public_key text [not null] // Clave pública para verificación
  attestation_data jsonb [null] // Datos de attestation del dispositivo

  is_active boolean [default: true]
  last_used_at timestamp
  registration_ip inet
  registration_user_agent text

  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]
}

Table users {
  id uuid [primary key]
  email varchar(255) [unique, null] // Ahora puede ser null si usa solo OAuth
  password_hash varchar(255) [null] // Null si no usa password
  first_name varchar(100)
  last_name varchar(100)
  phone varchar(20)

  // Verificaciones por método
  email_verified boolean [default: false]
  phone_verified boolean [default: false]

  // Configuración de seguridad
  two_factor_enabled boolean [default: false]
  require_biometric boolean [default: false] // Requiere biométrico como 2FA

  // Método de autenticación preferido
  preferred_auth_method_id uuid [ref: > authentication_methods.id]

  status enum('active', 'suspended', 'deleted') [default: 'active']
  structural_position_id uuid [ref: > structural_positions.id]
  organic_unit_id uuid [ref: > organic_units.id]
  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]
  last_login_at timestamp

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]
}

Table application_auth_policies {
  id uuid [primary key]
  application_id uuid [ref: > applications.id]

  // Políticas generales
  require_2fa boolean [default: false]
  allow_multiple_sessions boolean [default: true]
  session_timeout_minutes int [default: 480] // 8 horas

  // Políticas por método
  min_password_length int [default: 8]
  require_biometric_fallback boolean [default: false]
  oauth_auto_registration boolean [default: true] // Crear usuario automáticamente con OAuth

  // Políticas de seguridad
  max_failed_attempts int [default: 5]
  lockout_duration_minutes int [default: 30]
  require_device_registration boolean [default: false]

  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]
}

Table login_attempts {
  id uuid [primary key]
  email varchar(255) [null] // Puede ser null para métodos no basados en email
  user_id uuid [null, ref: > users.id] // Null si el usuario no existe
  authentication_method_id uuid [ref: > authentication_methods.id]
  application_id uuid [ref: > applications.id]

  ip_address inet [not null]
  user_agent text
  device_fingerprint varchar(255)

  success boolean [not null]
  failure_reason varchar(100) [null]
  blocked_until timestamp [null]

  // Datos específicos del intento
  method_specific_data jsonb [null]

  created_at timestamp [default: `now()`]

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]
}
```
