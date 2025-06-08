```
// Integrated System: Central User Manager + Organic Units with Full Audit Trail
// Compatible with dbdiagram.io - Public Entity Standards

// ===== ORGANIC UNITS HIERARCHY =====
Table organic_units {
  unit_id bigint [pk, increment]
  name varchar(255) [not null, note: 'Full name of the organic unit']
  acronym varchar(20) [not null, note: 'Acronym/initials of the organic unit']
  brand text [note: 'URL of the unit image/logo']
  parent_id bigint [ref: > organic_units.unit_id, note: 'Parent organic unit ID (self-reference)']
  level int [not null, default: 1, note: 'Hierarchical level (1-7)']
  description text [note: 'Description of the organic unit']
  active tinyint [not null, default: 1, note: 'Active/inactive status (1/0)']
  sort_order int [default: 0, note: 'Display order within the same level']

  // Logical deletion and audit fields
  deleted_at timestamp [null, note: 'Logical deletion timestamp']
  is_deleted boolean [not null, default: false, note: 'Logical deletion flag']
  version int [not null, default: 1, note: 'Record version for change tracking']

  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]
  created_by bigint [ref: > users.user_id]
  updated_by bigint [ref: > users.user_id]
  deleted_by bigint [ref: > users.user_id, note: 'User who performed logical deletion']

  indexes {
    acronym
    parent_id
    level
    active
    is_deleted
    (parent_id, level)
    (acronym, is_deleted) [unique, note: 'Unique acronym among non-deleted records']
  }

  Note: '''
  Table to model the organizational hierarchy of an institution.
  Allows up to 7 levels of depth through self-reference.
  Implements logical deletion and version control for audit compliance.
  '''
}

// Version history for organic units
Table organic_units_history {
  history_id bigint [pk, increment]
  unit_id bigint [ref: > organic_units.unit_id]
  name varchar(255) [not null]
  acronym varchar(20) [not null]
  brand text
  parent_id bigint
  level int [not null]
  description text
  active tinyint [not null]
  sort_order int
  version int [not null]

  // Change tracking
  change_type varchar(20) [not null, note: 'INSERT, UPDATE, DELETE']
  changed_fields json [note: 'JSON array of changed field names']
  old_values json [note: 'JSON object with previous values']
  new_values json [note: 'JSON object with new values']

  created_at timestamp [default: `now()`]
  created_by bigint [ref: > users.user_id]

  indexes {
    unit_id
    version
    change_type
    created_at
  }

  Note: '''
  Complete version history of organic units changes.
  Stores all modifications for audit trail compliance.
  '''
}

// ===== CENTRAL USER MANAGER =====
Table users {
  user_id bigint [pk, increment]
  email varchar(255) [not null]
  password_hash varchar(255) [not null]
  full_name varchar(255) [not null]
  phone varchar(20)
  address text
  profile_picture varchar(500)
  organic_unit_id bigint [ref: > organic_units.unit_id, note: 'Organic unit where the user belongs']
  status smallint [not null, default: 1, note: '1=active, 0=inactive, 2=suspended']

  // Logical deletion and audit fields
  deleted_at timestamp [null, note: 'Logical deletion timestamp']
  is_deleted boolean [not null, default: false, note: 'Logical deletion flag']
  version int [not null, default: 1, note: 'Record version for change tracking']

  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]
  created_by bigint
  updated_by bigint
  deleted_by bigint [ref: > users.user_id, note: 'User who performed logical deletion']

  indexes {
    (email, is_deleted) [unique, note: 'Unique email among non-deleted users']
    status
    organic_unit_id
    is_deleted
    created_at
  }

  Note: '''
  Users table with organic unit relationship and full audit trail.
  Implements logical deletion to preserve data integrity.
  '''
}

// Version history for users
Table users_history {
  history_id bigint [pk, increment]
  user_id bigint [ref: > users.user_id]
  email varchar(255) [not null]
  full_name varchar(255) [not null]
  phone varchar(20)
  address text
  profile_picture varchar(500)
  organic_unit_id bigint
  status smallint [not null]
  version int [not null]

  // Change tracking
  change_type varchar(20) [not null, note: 'INSERT, UPDATE, DELETE']
  changed_fields json [note: 'JSON array of changed field names']
  old_values json [note: 'JSON object with previous values']
  new_values json [note: 'JSON object with new values']

  created_at timestamp [default: `now()`]
  created_by bigint [ref: > users.user_id]

  indexes {
    user_id
    version
    change_type
    created_at
  }

  Note: '''
  Complete version history of user changes.
  Essential for public entity audit requirements.
  '''
}

Table password_history {
  history_id bigint [pk, increment]
  user_id bigint [ref: > users.user_id]
  previous_password_hash varchar(255) [not null]
  changed_at timestamp [default: `now()`]

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by bigint [ref: > users.user_id]

  indexes {
    user_id
    changed_at
    is_deleted
  }
}

Table active_sessions {
  session_id varchar(36) [pk]
  user_id bigint [ref: > users.user_id]
  session_token varchar(500) [not null]
  started_at timestamp [default: `now()`]
  last_accessed_at timestamp [default: `now()`]
  expires_at timestamp [not null]
  ip_address varchar(45)
  user_agent text
  session_status smallint [default: 1, note: '1=active, 0=closed, 2=expired']

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]

  indexes {
    user_id
    (session_token, is_deleted) [unique]
    expires_at
    session_status
    is_deleted
  }
}

Table session_history {
  session_history_id bigint [pk, increment]
  session_id varchar(36) [ref: > active_sessions.session_id]
  closed_at timestamp [default: `now()`]
  close_reason varchar(255)

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by bigint [ref: > users.user_id]

  indexes {
    session_id
    closed_at
    is_deleted
  }
}

Table systems {
  system_id bigint [pk, increment]
  system_name varchar(100) [not null]
  description text
  system_code varchar(20)
  status smallint [default: 1, note: '1=active, 0=inactive']
  configuration_json json

  // Logical deletion and audit
  deleted_at timestamp [null]
  is_deleted boolean [not null, default: false]
  version int [not null, default: 1]

  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]
  created_by bigint [ref: > users.user_id]
  updated_by bigint [ref: > users.user_id]
  deleted_by bigint [ref: > users.user_id]

  indexes {
    (system_name, is_deleted) [unique]
    (system_code, is_deleted) [unique]
    status
    is_deleted
  }
}

Table systems_history {
  history_id bigint [pk, increment]
  system_id bigint [ref: > systems.system_id]
  system_name varchar(100) [not null]
  description text
  system_code varchar(20)
  status smallint
  configuration_json json
  version int [not null]

  change_type varchar(20) [not null]
  changed_fields json
  old_values json
  new_values json

  created_at timestamp [default: `now()`]
  created_by bigint [ref: > users.user_id]

  indexes {
    system_id
    version
    created_at
  }
}

Table roles {
  role_id bigint [pk, increment]
  system_id bigint [ref: > systems.system_id]
  role_name varchar(100) [not null]
  role_description text
  role_status smallint [default: 1, note: '1=active, 0=inactive']
  priority_level int [default: 0]

  // Logical deletion and audit
  deleted_at timestamp [null]
  is_deleted boolean [not null, default: false]
  version int [not null, default: 1]

  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]
  created_by bigint [ref: > users.user_id]
  updated_by bigint [ref: > users.user_id]
  deleted_by bigint [ref: > users.user_id]

  indexes {
    (system_id, role_name, is_deleted) [unique]
    system_id
    role_status
    is_deleted
  }
}

Table permissions {
  permission_id bigint [pk, increment]
  system_id bigint [ref: > systems.system_id]
  permission_name varchar(100) [not null]
  permission_description text
  permission_category varchar(50)
  action_type varchar(20) [note: 'create, read, update, delete']
  permission_status smallint [default: 1, note: '1=active, 0=inactive']

  // Logical deletion and audit
  deleted_at timestamp [null]
  is_deleted boolean [not null, default: false]
  version int [not null, default: 1]

  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]
  created_by bigint [ref: > users.user_id]
  updated_by bigint [ref: > users.user_id]
  deleted_by bigint [ref: > users.user_id]

  indexes {
    (system_id, permission_name, is_deleted) [unique]
    system_id
    permission_category
    action_type
    permission_status
    is_deleted
  }
}

Table role_permissions {
  role_permission_id bigint [pk, increment]
  role_id bigint [ref: > roles.role_id]
  permission_id bigint [ref: > permissions.permission_id]
  assigned_at timestamp [default: `now()`]
  assigned_by bigint [ref: > users.user_id]

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by bigint [ref: > users.user_id]

  indexes {
    (role_id, permission_id, is_deleted) [unique]
    role_id
    permission_id
    is_deleted
  }
}

Table user_roles {
  user_role_id bigint [pk, increment]
  user_id bigint [ref: > users.user_id]
  role_id bigint [ref: > roles.role_id]
  system_id bigint [ref: > systems.system_id]
  assigned_at timestamp [default: `now()`]
  assigned_by bigint [ref: > users.user_id]
  assignment_status smallint [default: 1, note: '1=active, 0=inactive']

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by bigint [ref: > users.user_id]

  indexes {
    (user_id, role_id, system_id, is_deleted) [unique]
    user_id
    role_id
    system_id
    assignment_status
    is_deleted
  }
}

// ===== STRUCTURAL POSITIONS =====
Table structural_positions {
  position_id bigint [pk, increment]
  organic_unit_id bigint [ref: > organic_units.unit_id]
  position_name varchar(255) [not null]
  position_code varchar(20)
  description text
  level int [note: 'Position level within the unit']
  is_head_position boolean [default: false, note: 'Indicates if this is the head position of the unit']
  active tinyint [not null, default: 1]

  // Logical deletion and audit
  deleted_at timestamp [null]
  is_deleted boolean [not null, default: false]
  version int [not null, default: 1]

  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]
  created_by bigint [ref: > users.user_id]
  updated_by bigint [ref: > users.user_id]
  deleted_by bigint [ref: > users.user_id]

  indexes {
    organic_unit_id
    (organic_unit_id, position_code, is_deleted) [unique]
    is_head_position
    active
    is_deleted
  }

  Note: '''
  Defines specific structural positions within each organic unit.
  Implements full audit trail for public entity compliance.
  '''
}

// Version history for structural positions
Table structural_positions_history {
  history_id bigint [pk, increment]
  position_id bigint [ref: > structural_positions.position_id]
  organic_unit_id bigint
  position_name varchar(255) [not null]
  position_code varchar(20)
  description text
  level int
  is_head_position boolean
  active tinyint [not null]
  version int [not null]

  // Change tracking
  change_type varchar(20) [not null]
  changed_fields json
  old_values json
  new_values json

  created_at timestamp [default: `now()`]
  created_by bigint [ref: > users.user_id]

  indexes {
    position_id
    version
    change_type
    created_at
  }

  Note: '''
  Complete version history of structural positions changes.
  Critical for organizational audit requirements.
  '''
}

Table user_structural_positions {
  user_position_id bigint [pk, increment]
  user_id bigint [ref: > users.user_id]
  position_id bigint [ref: > structural_positions.position_id]
  assignment_date date [not null]
  end_date date [note: 'NULL for active assignments']
  assignment_type varchar(50) [note: 'permanent, temporary, acting, etc.']
  status smallint [default: 1, note: '1=active, 0=inactive']

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by bigint [ref: > users.user_id]

  created_at timestamp [default: `now()`]
  created_by bigint [ref: > users.user_id]

  indexes {
    user_id
    position_id
    (user_id, position_id, assignment_date, is_deleted) [unique]
    assignment_date
    status
    is_deleted
  }
}

// ===== PERSONNEL MOVEMENTS WITH AUDIT =====
Table personnel_movements {
  movement_id bigint [pk, increment]
  user_id bigint [ref: > users.user_id]
  from_unit_id bigint [ref: > organic_units.unit_id]
  to_unit_id bigint [ref: > organic_units.unit_id]
  from_position_id bigint [ref: > structural_positions.position_id, note: 'Previous structural position']
  to_position_id bigint [ref: > structural_positions.position_id, note: 'New structural position']
  movement_type varchar(50) [note: 'transfer, promotion, assignment, etc.']
  effective_date date [not null]
  reason text
  approved_by bigint [ref: > users.user_id]

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by bigint [ref: > users.user_id]

  created_at timestamp [default: `now()`]

  indexes {
    user_id
    from_unit_id
    to_unit_id
    effective_date
    movement_type
    is_deleted
  }

  Note: '''
  Tracks personnel movements between organic units and positions.
  Essential for maintaining complete organizational audit trail.
  '''
}

// ===== REMAINING TABLES WITH LOGICAL DELETION =====
Table modules {
  module_id bigint [pk, increment]
  system_id bigint [ref: > systems.system_id]
  module_name varchar(100) [not null]
  module_path varchar(255)
  module_description text
  module_status smallint [default: 1, note: '1=active, 0=inactive']

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by bigint [ref: > users.user_id]

  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]
  created_by bigint [ref: > users.user_id]
  updated_by bigint [ref: > users.user_id]

  indexes {
    (system_id, module_name, is_deleted) [unique]
    system_id
    module_status
    is_deleted
  }
}

Table module_permissions {
  module_permission_id bigint [pk, increment]
  module_id bigint [ref: > modules.module_id]
  permission_id bigint [ref: > permissions.permission_id]
  detail_description text
  status smallint [default: 1, note: '1=active, 0=inactive']

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by bigint [ref: > users.user_id]

  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]

  indexes {
    (module_id, permission_id, is_deleted) [unique]
    module_id
    permission_id
    is_deleted
  }
}

Table audit_logs {
  log_id bigint [pk, increment]
  user_id bigint [ref: > users.user_id]
  action varchar(100) [not null]
  affected_table varchar(50)
  affected_record_id bigint
  description text
  ip_address varchar(45)
  user_agent text
  logged_at timestamp [default: `now()`]

  Note: '''
  Audit logs are NEVER deleted - permanent record required.
  Critical for public entity compliance and legal requirements.
  '''

  indexes {
    user_id
    action
    affected_table
    logged_at
  }
}

Table verification_tokens {
  token_id bigint [pk, increment]
  user_id bigint [ref: > users.user_id]
  token varchar(255) [not null]
  token_type smallint [not null, note: '1=email_verification, 2=password_reset']
  created_at timestamp [default: `now()`]
  expires_at timestamp [not null]
  token_status smallint [default: 1, note: '1=active, 2=used, 3=expired']

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]

  indexes {
    (token, is_deleted) [unique]
    user_id
    token_type
    token_status
    expires_at
    is_deleted
  }
}

Table mfa_devices {
  mfa_id bigint [pk, increment]
  user_id bigint [ref: > users.user_id]
  mfa_type smallint [not null, note: '1=SMS, 2=APP, 3=U2F, 4=EMAIL']
  secret_key varchar(255)
  mfa_status smallint [default: 1, note: '1=active, 0=inactive']
  registered_at timestamp [default: `now()`]
  last_validated_at timestamp

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by bigint [ref: > users.user_id]

  indexes {
    user_id
    mfa_type
    mfa_status
    is_deleted
  }
}

Table groups {
  group_id bigint [pk, increment]
  group_name varchar(100) [not null]
  group_description text
  system_id bigint [ref: > systems.system_id]

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by bigint [ref: > users.user_id]

  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]
  created_by bigint [ref: > users.user_id]
  updated_by bigint [ref: > users.user_id]

  indexes {
    (system_id, group_name, is_deleted) [unique]
    system_id
    is_deleted
  }
}

Table user_groups {
  user_group_id bigint [pk, increment]
  user_id bigint [ref: > users.user_id]
  group_id bigint [ref: > groups.group_id]
  assigned_at timestamp [default: `now()`]
  assigned_by bigint [ref: > users.user_id]
  status smallint [default: 1, note: '1=active, 0=inactive']

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by bigint [ref: > users.user_id]

  indexes {
    (user_id, group_id, is_deleted) [unique]
    user_id
    group_id
    is_deleted
  }
}

Table session_expiration_policies {
  policy_id bigint [pk, increment]
  system_id bigint [ref: > systems.system_id]
  expiration_minutes int [not null, default: 480]
  remember_me boolean [default: false]

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by bigint [ref: > users.user_id]

  created_at timestamp [default: `now()`]
  updated_at timestamp [default: `now()`]

  indexes {
    (system_id, is_deleted) [unique]
    is_deleted
  }
}

Table api_tokens {
  api_token_id bigint [pk, increment]
  user_id bigint [ref: > users.user_id]
  system_id bigint [ref: > systems.system_id]
  api_token varchar(255) [not null]
  scopes json
  created_at timestamp [default: `now()`]
  expires_at timestamp
  status smallint [default: 1, note: '1=active, 0=inactive']

  // Logical deletion
  is_deleted boolean [not null, default: false]
  deleted_at timestamp [null]
  deleted_by bigint [ref: > users.user_id]

  indexes {
    (api_token, is_deleted) [unique]
    user_id
    system_id
    status
    is_deleted
  }
}

// ===== SELF-REFERENCING RELATIONSHIPS =====
Ref: users.created_by > users.user_id
Ref: users.updated_by > users.user_id
```
