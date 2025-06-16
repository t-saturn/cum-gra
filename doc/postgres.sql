CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "email" varchar(255) UNIQUE NOT NULL,
  "password_hash" varchar(255) NOT NULL,
  "first_name" varchar(100),
  "last_name" varchar(100),
  "phone" varchar(20),
  "email_verified" boolean DEFAULT false,
  "phone_verified" boolean DEFAULT false,
  "two_factor_enabled" boolean DEFAULT false,
  "status" enum(active,suspended,deleted) DEFAULT 'active',
  "structural_position_id" uuid,
  "organic_unit_id" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "last_login_at" timestamp,
  "is_deleted" boolean NOT NULL DEFAULT false,
  "deleted_at" timestamp,
  "deleted_by" uuid
);

CREATE TABLE "user_sessions" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid,
  "session_token" varchar(512) UNIQUE NOT NULL,
  "refresh_token" varchar(512) UNIQUE NOT NULL,
  "device_info" text,
  "ip_address" inet,
  "user_agent" text,
  "is_active" boolean DEFAULT true,
  "expires_at" timestamp NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "last_activity_at" timestamp DEFAULT (now()),
  "is_deleted" boolean NOT NULL DEFAULT false,
  "deleted_at" timestamp
);

CREATE TABLE "applications" (
  "id" uuid PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "client_id" varchar(255) UNIQUE NOT NULL,
  "client_secret" varchar(255) NOT NULL,
  "domain" varchar(255) NOT NULL,
  "logo" varchar(255),
  "description" text,
  "callback_urls" text[],
  "scopes" text[],
  "is_first_party" boolean DEFAULT false,
  "status" enum(active,suspended) DEFAULT 'active',
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "is_deleted" boolean NOT NULL DEFAULT false,
  "deleted_at" timestamp,
  "deleted_by" uuid
);

CREATE TABLE "application_roles" (
  "id" uuid PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "description" text,
  "application_id" uuid,
  "permissions" text[],
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "is_deleted" boolean NOT NULL DEFAULT false,
  "deleted_at" timestamp,
  "deleted_by" uuid
);

CREATE TABLE "user_application_roles" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid,
  "application_id" uuid,
  "application_role_id" uuid,
  "granted_at" timestamp DEFAULT (now()),
  "granted_by" uuid,
  "revoked_at" timestamp,
  "revoked_by" uuid,
  "is_deleted" boolean NOT NULL DEFAULT false,
  "deleted_at" timestamp,
  "deleted_by" uuid
);

CREATE TABLE "modules" (
  "id" uuid PRIMARY KEY,
  "item" varchar(100),
  "name" varchar(100) NOT NULL,
  "label" varchar(100),
  "route" varchar(255),
  "icon" varchar(100),
  "parent_id" uuid,
  "application_id" uuid,
  "sort_order" int DEFAULT 0,
  "is_menu_item" boolean DEFAULT true,
  "status" enum(active,inactive) DEFAULT 'active',
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "is_deleted" boolean NOT NULL DEFAULT false,
  "deleted_at" timestamp,
  "deleted_by" uuid
);

CREATE TABLE "module_role_permissions" (
  "id" uuid PRIMARY KEY,
  "module_id" uuid,
  "application_role_id" uuid,
  "permission_type" enum(denied,read,write,admin) NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "is_deleted" boolean NOT NULL DEFAULT false,
  "deleted_at" timestamp,
  "deleted_by" uuid
);

CREATE TABLE "oauth_tokens" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid,
  "application_id" uuid,
  "access_token" varchar(512) UNIQUE NOT NULL,
  "refresh_token" varchar(512) UNIQUE,
  "token_type" varchar(50) DEFAULT 'Bearer',
  "scopes" text[],
  "expires_at" timestamp NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "revoked_at" timestamp,
  "is_deleted" boolean NOT NULL DEFAULT false,
  "deleted_at" timestamp,
  "deleted_by" uuid
);

CREATE TABLE "structural_positions" (
  "id" uuid PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "code" varchar(50) UNIQUE NOT NULL,
  "level" varchar(50),
  "description" text,
  "is_active" boolean DEFAULT true,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "is_deleted" boolean NOT NULL DEFAULT false,
  "deleted_at" timestamp,
  "deleted_by" uuid
);

CREATE TABLE "organic_units" (
  "id" uuid PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "acronym" varchar(20),
  "brand" varchar(100),
  "level" varchar(50),
  "description" text,
  "parent_id" uuid,
  "is_active" boolean DEFAULT true,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "is_deleted" boolean NOT NULL DEFAULT false,
  "deleted_at" timestamp,
  "deleted_by" uuid
);

CREATE TABLE "password_history" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid,
  "previous_password_hash" varchar(255) NOT NULL,
  "changed_at" timestamp DEFAULT (now()),
  "changed_by" uuid,
  "is_deleted" boolean NOT NULL DEFAULT false,
  "deleted_at" timestamp,
  "deleted_by" uuid
);

CREATE TABLE "password_resets" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid,
  "token" varchar(255) UNIQUE NOT NULL,
  "expires_at" timestamp NOT NULL,
  "used_at" timestamp,
  "ip_address" inet,
  "user_agent" text,
  "created_at" timestamp DEFAULT (now()),
  "is_deleted" boolean NOT NULL DEFAULT false,
  "deleted_at" timestamp,
  "deleted_by" uuid
);

COMMENT ON COLUMN "applications"."callback_urls" IS 'Array of allowed callback URLs';

COMMENT ON COLUMN "applications"."scopes" IS 'Array of available scopes';

COMMENT ON COLUMN "applications"."is_first_party" IS 'True for internal applications';

COMMENT ON COLUMN "application_roles"."permissions" IS 'Array of permissions for this role';

COMMENT ON TABLE "user_application_roles" IS 'Relaciona un usuario con UN SOLO rol por aplicación específica';

COMMENT ON TABLE "module_role_permissions" IS 'Define permisos específicos de roles sobre módulos';

COMMENT ON COLUMN "oauth_tokens"."scopes" IS 'Granted scopes';

COMMENT ON COLUMN "oauth_tokens"."revoked_at" IS 'When token was revoked';

ALTER TABLE "users" ADD FOREIGN KEY ("structural_position_id") REFERENCES "structural_positions" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("organic_unit_id") REFERENCES "organic_units" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("deleted_by") REFERENCES "users" ("id");

ALTER TABLE "user_sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "applications" ADD FOREIGN KEY ("deleted_by") REFERENCES "users" ("id");

ALTER TABLE "application_roles" ADD FOREIGN KEY ("application_id") REFERENCES "applications" ("id");

ALTER TABLE "application_roles" ADD FOREIGN KEY ("deleted_by") REFERENCES "users" ("id");

ALTER TABLE "user_application_roles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_application_roles" ADD FOREIGN KEY ("application_id") REFERENCES "applications" ("id");

ALTER TABLE "user_application_roles" ADD FOREIGN KEY ("application_role_id") REFERENCES "application_roles" ("id");

ALTER TABLE "user_application_roles" ADD FOREIGN KEY ("granted_by") REFERENCES "users" ("id");

ALTER TABLE "user_application_roles" ADD FOREIGN KEY ("revoked_by") REFERENCES "users" ("id");

ALTER TABLE "user_application_roles" ADD FOREIGN KEY ("deleted_by") REFERENCES "users" ("id");

ALTER TABLE "modules" ADD FOREIGN KEY ("parent_id") REFERENCES "modules" ("id");

ALTER TABLE "modules" ADD FOREIGN KEY ("application_id") REFERENCES "applications" ("id");

ALTER TABLE "modules" ADD FOREIGN KEY ("deleted_by") REFERENCES "users" ("id");

ALTER TABLE "module_role_permissions" ADD FOREIGN KEY ("module_id") REFERENCES "modules" ("id");

ALTER TABLE "module_role_permissions" ADD FOREIGN KEY ("application_role_id") REFERENCES "application_roles" ("id");

ALTER TABLE "module_role_permissions" ADD FOREIGN KEY ("deleted_by") REFERENCES "users" ("id");

ALTER TABLE "oauth_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "oauth_tokens" ADD FOREIGN KEY ("application_id") REFERENCES "applications" ("id");

ALTER TABLE "oauth_tokens" ADD FOREIGN KEY ("deleted_by") REFERENCES "users" ("id");

ALTER TABLE "structural_positions" ADD FOREIGN KEY ("deleted_by") REFERENCES "users" ("id");

ALTER TABLE "organic_units" ADD FOREIGN KEY ("parent_id") REFERENCES "organic_units" ("id");

ALTER TABLE "organic_units" ADD FOREIGN KEY ("deleted_by") REFERENCES "users" ("id");

ALTER TABLE "password_history" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "password_history" ADD FOREIGN KEY ("changed_by") REFERENCES "users" ("id");

ALTER TABLE "password_history" ADD FOREIGN KEY ("deleted_by") REFERENCES "users" ("id");

ALTER TABLE "password_resets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "password_resets" ADD FOREIGN KEY ("deleted_by") REFERENCES "users" ("id");
