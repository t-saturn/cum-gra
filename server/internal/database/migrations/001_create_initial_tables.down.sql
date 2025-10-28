-- Eliminar índices
DROP INDEX IF EXISTS idx_organic_units_parent_id;
DROP INDEX IF EXISTS idx_modules_application_id;
DROP INDEX IF EXISTS idx_modules_parent_id;
DROP INDEX IF EXISTS idx_applications_is_deleted;
DROP INDEX IF EXISTS idx_applications_status;
DROP INDEX IF EXISTS idx_applications_client_id;
DROP INDEX IF EXISTS idx_users_is_deleted;
DROP INDEX IF EXISTS idx_users_status;
DROP INDEX IF EXISTS idx_users_dni;
DROP INDEX IF EXISTS idx_users_email;

-- Eliminar tablas en orden inverso de dependencias
DROP TABLE IF EXISTS password_histories;
DROP TABLE IF EXISTS user_module_restrictions;
DROP TABLE IF EXISTS module_role_permissions;
DROP TABLE IF EXISTS modules;
DROP TABLE IF EXISTS user_application_roles;
DROP TABLE IF EXISTS application_roles;
DROP TABLE IF EXISTS applications;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS organic_units;
DROP TABLE IF EXISTS structural_positions;
