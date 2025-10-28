-- Agregar foreign keys para structural_positions
ALTER TABLE structural_positions
ADD CONSTRAINT fk_structural_positions_deleted_by
FOREIGN KEY (deleted_by) REFERENCES users(id) ON DELETE SET NULL;

-- Agregar foreign keys para organic_units
ALTER TABLE organic_units
ADD CONSTRAINT fk_organic_units_parent_id
FOREIGN KEY (parent_id) REFERENCES organic_units(id) ON DELETE SET NULL;

ALTER TABLE organic_units
ADD CONSTRAINT fk_organic_units_deleted_by
FOREIGN KEY (deleted_by) REFERENCES users(id) ON DELETE SET NULL;

-- Agregar foreign keys para users
ALTER TABLE users
ADD CONSTRAINT fk_users_structural_position_id
FOREIGN KEY (structural_position_id) REFERENCES structural_positions(id) ON DELETE SET NULL;

ALTER TABLE users
ADD CONSTRAINT fk_users_organic_unit_id
FOREIGN KEY (organic_unit_id) REFERENCES organic_units(id) ON DELETE SET NULL;

ALTER TABLE users
ADD CONSTRAINT fk_users_deleted_by
FOREIGN KEY (deleted_by) REFERENCES users(id) ON DELETE SET NULL;

-- Agregar foreign keys para applications
ALTER TABLE applications
ADD CONSTRAINT fk_applications_deleted_by
FOREIGN KEY (deleted_by) REFERENCES users(id) ON DELETE SET NULL;

-- Agregar foreign keys para application_roles
ALTER TABLE application_roles
ADD CONSTRAINT fk_application_roles_application_id
FOREIGN KEY (application_id) REFERENCES applications(id) ON DELETE CASCADE;

ALTER TABLE application_roles
ADD CONSTRAINT fk_application_roles_deleted_by
FOREIGN KEY (deleted_by) REFERENCES users(id) ON DELETE SET NULL;

-- Agregar foreign keys para user_application_roles
ALTER TABLE user_application_roles
ADD CONSTRAINT fk_user_application_roles_user_id
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE user_application_roles
ADD CONSTRAINT fk_user_application_roles_application_id
FOREIGN KEY (application_id) REFERENCES applications(id) ON DELETE CASCADE;

ALTER TABLE user_application_roles
ADD CONSTRAINT fk_user_application_roles_application_role_id
FOREIGN KEY (application_role_id) REFERENCES application_roles(id) ON DELETE CASCADE;

ALTER TABLE user_application_roles
ADD CONSTRAINT fk_user_application_roles_granted_by
FOREIGN KEY (granted_by) REFERENCES users(id) ON DELETE RESTRICT;

ALTER TABLE user_application_roles
ADD CONSTRAINT fk_user_application_roles_revoked_by
FOREIGN KEY (revoked_by) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE user_application_roles
ADD CONSTRAINT fk_user_application_roles_deleted_by
FOREIGN KEY (deleted_by) REFERENCES users(id) ON DELETE SET NULL;

-- Agregar foreign keys para modules
ALTER TABLE modules
ADD CONSTRAINT fk_modules_parent_id
FOREIGN KEY (parent_id) REFERENCES modules(id) ON DELETE SET NULL;

ALTER TABLE modules
ADD CONSTRAINT fk_modules_application_id
FOREIGN KEY (application_id) REFERENCES applications(id) ON DELETE SET NULL;

ALTER TABLE modules
ADD CONSTRAINT fk_modules_deleted_by
FOREIGN KEY (deleted_by) REFERENCES users(id) ON DELETE SET NULL;

-- Agregar foreign keys para module_role_permissions
ALTER TABLE module_role_permissions
ADD CONSTRAINT fk_module_role_permissions_module_id
FOREIGN KEY (module_id) REFERENCES modules(id) ON DELETE CASCADE;

ALTER TABLE module_role_permissions
ADD CONSTRAINT fk_module_role_permissions_application_role_id
FOREIGN KEY (application_role_id) REFERENCES application_roles(id) ON DELETE CASCADE;

ALTER TABLE module_role_permissions
ADD CONSTRAINT fk_module_role_permissions_deleted_by
FOREIGN KEY (deleted_by) REFERENCES users(id) ON DELETE SET NULL;

-- Agregar foreign keys para user_module_restrictions
ALTER TABLE user_module_restrictions
ADD CONSTRAINT fk_user_module_restrictions_user_id
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE user_module_restrictions
ADD CONSTRAINT fk_user_module_restrictions_module_id
FOREIGN KEY (module_id) REFERENCES modules(id) ON DELETE CASCADE;

ALTER TABLE user_module_restrictions
ADD CONSTRAINT fk_user_module_restrictions_application_id
FOREIGN KEY (application_id) REFERENCES applications(id) ON DELETE CASCADE;

ALTER TABLE user_module_restrictions
ADD CONSTRAINT fk_user_module_restrictions_created_by
FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT;

ALTER TABLE user_module_restrictions
ADD CONSTRAINT fk_user_module_restrictions_updated_by
FOREIGN KEY (updated_by) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE user_module_restrictions
ADD CONSTRAINT fk_user_module_restrictions_deleted_by
FOREIGN KEY (deleted_by) REFERENCES users(id) ON DELETE SET NULL;

-- Agregar foreign keys para password_histories
ALTER TABLE password_histories
ADD CONSTRAINT fk_password_histories_user_id
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE password_histories
ADD CONSTRAINT fk_password_histories_changed_by
FOREIGN KEY (changed_by) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE password_histories
ADD CONSTRAINT fk_password_histories_deleted_by
FOREIGN KEY (deleted_by) REFERENCES users(id) ON DELETE SET NULL;

-- Crear índices adicionales para foreign keys para mejorar el rendimiento
CREATE INDEX idx_users_structural_position_id ON users(structural_position_id);
CREATE INDEX idx_users_organic_unit_id ON users(organic_unit_id);
CREATE INDEX idx_users_deleted_by ON users(deleted_by);
CREATE INDEX idx_applications_deleted_by ON applications(deleted_by);
CREATE INDEX idx_application_roles_application_id ON application_roles(application_id);
CREATE INDEX idx_application_roles_deleted_by ON application_roles(deleted_by);
CREATE INDEX idx_user_application_roles_user_id ON user_application_roles(user_id);
CREATE INDEX idx_user_application_roles_application_id ON user_application_roles(application_id);
CREATE INDEX idx_user_application_roles_application_role_id ON user_application_roles(application_role_id);
CREATE INDEX idx_user_application_roles_granted_by ON user_application_roles(granted_by);
CREATE INDEX idx_modules_deleted_by ON modules(deleted_by);
CREATE INDEX idx_module_role_permissions_module_id ON module_role_permissions(module_id);
CREATE INDEX idx_module_role_permissions_application_role_id ON module_role_permissions(application_role_id);
CREATE INDEX idx_user_module_restrictions_user_id ON user_module_restrictions(user_id);
CREATE INDEX idx_user_module_restrictions_module_id ON user_module_restrictions(module_id);
CREATE INDEX idx_user_module_restrictions_application_id ON user_module_restrictions(application_id);
CREATE INDEX idx_password_histories_user_id ON password_histories(user_id);
