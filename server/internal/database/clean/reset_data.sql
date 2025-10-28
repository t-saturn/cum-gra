-- Vaciar datos de todas las tablas (sin eliminar las estructuras)
TRUNCATE TABLE
  password_histories,
  user_module_restrictions,
  user_application_roles,
  module_role_permissions,
  users,
  application_roles,
  applications,
  structural_positions,
  organic_units,
  modules
RESTART IDENTITY CASCADE;
