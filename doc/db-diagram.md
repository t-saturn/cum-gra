```sql
Table users {
  id uuid [primary key]
  email varchar(255) [unique, not null]
  password_hash varchar(255) [not null]
  first_name varchar(100)
  last_name varchar(100)
  phone varchar(20)
  email_verified boolean [default: false]
  phone_verified boolean [default: false]
  two_factor_enabled boolean [default: false]
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

  // Borrado lógico
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

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]
}

Table application_roles {
  id uuid [primary key]
  name varchar(100) [not null]
  description text
  application_id uuid [ref: > applications.id]
  permissions text[] [note: 'Array of permissions for this role']
  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]
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

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]

  Note: "Relaciona un usuario con UN SOLO rol por aplicación específica"
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

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]
}

Table module_role_permissions {
  id uuid [primary key]
  module_id uuid [ref: > modules.id]
  application_role_id uuid [ref: > application_roles.id]
  permission_type enum('denied', 'read', 'write', 'admin') [not null]
  created_at timestamp [default: `now()`]

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]

  Note: "Define permisos específicos de roles sobre módulos"
}

// NUEVA TABLA: Restricciones específicas de usuario
Table user_module_restrictions {
  id uuid [primary key]
  user_id uuid [ref: > users.id, not null]
  module_id uuid [ref: > modules.id, not null]
  application_id uuid [ref: > applications.id, not null]
  restriction_type enum('block_access', 'limit_permission') [not null]
  max_permission_level enum('denied', 'read', 'write', 'admin') [note: 'Solo si restriction_type es limit_permission']
  reason text [note: 'Motivo de la restricción']
  expires_at timestamp [note: 'Fecha de expiración de la restricción']
  created_at timestamp [default: `now()`]
  created_by uuid [ref: > users.id, not null]
  updated_at timestamp [default: `now()`]
  updated_by uuid [ref: > users.id]

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]

  Note: "Restricciones específicas de módulos por usuario que sobrescriben los permisos del rol"
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

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]
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

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]
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

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]
}

Table password_history {
  id uuid [primary key]
  user_id uuid [ref: > users.id]
  previous_password_hash varchar(255) [not null]
  changed_at timestamp [default: `now()`]
  changed_by uuid [ref: > users.id]

  // Borrado lógico
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

  // Borrado lógico
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by uuid [ref: > users.id]
}
```
