-- 001_create_initial_tables.up.sql

-- Crear tabla structural_positions
CREATE TABLE structural_positions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    code VARCHAR(50) UNIQUE NOT NULL,
    level INTEGER DEFAULT 2,
    description TEXT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by UUID
);

-- Crear tabla organic_units
CREATE TABLE organic_units (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) UNIQUE NOT NULL,
    acronym VARCHAR(20) UNIQUE,
    brand VARCHAR(100),
    description TEXT,
    parent_id UUID,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by UUID
);

-- Crear tabla users
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    phone VARCHAR(20),
    dni VARCHAR(8) UNIQUE NOT NULL,
    email_verified BOOLEAN DEFAULT false,
    phone_verified BOOLEAN DEFAULT false,
    two_factor_enabled BOOLEAN DEFAULT false,
    status VARCHAR(50) DEFAULT 'active',
    structural_position_id UUID,
    organic_unit_id UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by UUID
);

-- Crear tabla applications
CREATE TABLE applications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    client_id VARCHAR(255) UNIQUE NOT NULL,
    client_secret VARCHAR(255) NOT NULL,
    domain VARCHAR(255) NOT NULL,
    logo VARCHAR(255),
    description TEXT,
    callback_urls TEXT[],
    is_first_party BOOLEAN DEFAULT false,
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by UUID
);

-- Crear tabla application_roles
CREATE TABLE application_roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    application_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by UUID
);

-- Crear tabla user_application_roles
CREATE TABLE user_application_roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    application_id UUID NOT NULL,
    application_role_id UUID NOT NULL,
    granted_at TIMESTAMP DEFAULT NOW(),
    granted_by UUID NOT NULL,
    revoked_at TIMESTAMP,
    revoked_by UUID,
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by UUID
);

-- Crear tabla modules
CREATE TABLE modules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    item VARCHAR(100),
    name VARCHAR(100) NOT NULL,
    route VARCHAR(255),
    icon VARCHAR(100),
    parent_id UUID,
    application_id UUID,
    sort_order INTEGER DEFAULT 0,
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by UUID
);

-- Crear tabla module_role_permissions
CREATE TABLE module_role_permissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    module_id UUID NOT NULL,
    application_role_id UUID NOT NULL,
    permission_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by UUID
);

-- Crear tabla user_module_restrictions
CREATE TABLE user_module_restrictions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    module_id UUID NOT NULL,
    application_id UUID NOT NULL,
    restriction_type VARCHAR(50) NOT NULL,
    max_permission_level VARCHAR(50),
    reason TEXT,
    expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by UUID NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by UUID,
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by UUID
);

-- Crear tabla password_histories
CREATE TABLE password_histories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    previous_password_hash VARCHAR(255) NOT NULL,
    changed_at TIMESTAMP DEFAULT NOW(),
    changed_by UUID,
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMP,
    deleted_by UUID
);

-- Crear índices básicos para mejorar el rendimiento
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_dni ON users(dni);
CREATE INDEX idx_users_status ON users(status);
CREATE INDEX idx_users_is_deleted ON users(is_deleted);
CREATE INDEX idx_applications_client_id ON applications(client_id);
CREATE INDEX idx_applications_status ON applications(status);
CREATE INDEX idx_applications_is_deleted ON applications(is_deleted);
CREATE INDEX idx_modules_parent_id ON modules(parent_id);
CREATE INDEX idx_modules_application_id ON modules(application_id);
CREATE INDEX idx_organic_units_parent_id ON organic_units(parent_id);
