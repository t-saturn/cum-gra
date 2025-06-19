-- 002_add_foreign_keys.down.sql

-- Eliminar índices adicionales
DROP INDEX IF EXISTS idx_password_history_user_id;
DROP INDEX IF EXISTS idx_user_module_restrictions_application_id;
DROP INDEX IF EXISTS idx_user_module_restrictions_module_id;
DROP INDEX IF EXISTS idx_user_module_restrictions_user_id;
DROP INDEX IF EXISTS idx_module_role_permissions_application_role_id;
DROP INDEX IF EXISTS idx_module_role_permissions_module_id;
DROP INDEX IF EXISTS idx_modules_deleted_by;
DROP INDEX IF EXISTS idx_user_application_roles_granted_by;
DROP INDEX IF EXISTS idx_user_application_roles_application_role_id;
DROP INDEX IF EXISTS idx_user_application_roles_application_id;
DROP INDEX IF EXISTS idx_user_application_roles_user_id;
DROP INDEX IF EXISTS idx_application_roles_deleted_by;
DROP INDEX IF EXISTS idx_application_roles_application_id;
DROP INDEX IF EXISTS idx_applications_deleted_by;
DROP INDEX IF EXISTS idx_users_deleted_by;
DROP INDEX IF EXISTS idx_users_organic_unit_id;
DROP INDEX IF EXISTS idx_users_structural_position_id;

-- Eliminar foreign keys de password_history
ALTER TABLE password_history DROP CONSTRAINT IF EXISTS fk_password_history_deleted_by;
ALTER TABLE password_history DROP CONSTRAINT IF EXISTS fk_password_history_changed_by;
ALTER TABLE password_history DROP CONSTRAINT IF EXISTS fk_password_history_user_id;

-- Eliminar foreign keys de user_module_restrictions
ALTER TABLE user_module_restrictions DROP CONSTRAINT IF EXISTS fk_user_module_restrictions_deleted_by;
ALTER TABLE user_module_restrictions DROP CONSTRAINT IF EXISTS fk_user_module_restrictions_updated_by;
ALTER TABLE user_module_restrictions DROP CONSTRAINT IF EXISTS fk_user_module_restrictions_created_by;
ALTER TABLE user_module_restrictions DROP CONSTRAINT IF EXISTS fk_user_module_restrictions_application_id;
ALTER TABLE user_module_restrictions DROP CONSTRAINT IF EXISTS fk_user_module_restrictions_module_id;
ALTER TABLE user_module_restrictions DROP CONSTRAINT IF EXISTS fk_user_module_restrictions_user_id;

-- Eliminar foreign keys de module_role_permissions
ALTER TABLE module_role_permissions DROP CONSTRAINT IF EXISTS fk_module_role_permissions_deleted_by;
ALTER TABLE module_role_permissions DROP CONSTRAINT IF EXISTS fk_module_role_permissions_application_role_id;
ALTER TABLE module_role_permissions DROP CONSTRAINT IF EXISTS fk_module_role_permissions_module_id;

-- Eliminar foreign keys de modules
ALTER TABLE modules DROP CONSTRAINT IF EXISTS fk_modules_deleted_by;
ALTER TABLE modules DROP CONSTRAINT IF EXISTS fk_modules_application_id;
ALTER TABLE modules DROP CONSTRAINT IF EXISTS fk_modules_parent_id;

-- Eliminar foreign keys de user_application_roles
ALTER TABLE user_application_roles DROP CONSTRAINT IF EXISTS fk_user_application_roles_deleted_by;
ALTER TABLE user_application_roles DROP CONSTRAINT IF EXISTS fk_user_application_roles_revoked_by;
ALTER TABLE user_application_roles DROP CONSTRAINT IF EXISTS fk_user_application_roles_granted_by;
ALTER TABLE user_application_roles DROP CONSTRAINT IF EXISTS fk_user_application_roles_application_role_id;
ALTER TABLE user_application_roles DROP CONSTRAINT IF EXISTS fk_user_application_roles_application_id;
ALTER TABLE user_application_roles DROP CONSTRAINT IF EXISTS fk_user_application_roles_user_id;

-- Eliminar foreign keys de application_roles
ALTER TABLE application_roles DROP CONSTRAINT IF EXISTS fk_application_roles_deleted_by;
ALTER TABLE application_roles DROP CONSTRAINT IF EXISTS fk_application_roles_application_id;

-- Eliminar foreign keys de applications
ALTER TABLE applications DROP CONSTRAINT IF EXISTS fk_applications_deleted_by;

-- Eliminar foreign keys de users
ALTER TABLE users DROP CONSTRAINT IF EXISTS fk_users_deleted_by;
ALTER TABLE users DROP CONSTRAINT IF EXISTS fk_users_organic_unit_id;
ALTER TABLE users DROP CONSTRAINT IF EXISTS fk_users_structural_position_id;

-- Eliminar foreign keys de organic_units
ALTER TABLE organic_units DROP CONSTRAINT IF EXISTS fk_organic_units_deleted_by;
ALTER TABLE organic_units DROP CONSTRAINT IF EXISTS fk_organic_units_parent_id;

-- Eliminar foreign keys de structural_positions
ALTER TABLE structural_positions DROP CONSTRAINT IF EXISTS fk_structural_positions_deleted_by;
