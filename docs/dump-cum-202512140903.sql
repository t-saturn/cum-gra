--
-- PostgreSQL database dump
--

\restrict kI76rdtxNdVGDNcKGXrqBO6wiyjEvYAnUoqEUtEJVWkGbKqhVnqzZsjWX3BkFwo

-- Dumped from database version 17.6
-- Dumped by pg_dump version 17.6

-- Started on 2025-12-14 09:03:00

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 219 (class 1259 OID 21878)
-- Name: application_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.application_roles (
    id uuid NOT NULL,
    name character varying(100) NOT NULL,
    description text,
    application_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    is_deleted boolean DEFAULT false NOT NULL,
    deleted_at timestamp with time zone,
    deleted_by uuid
);


ALTER TABLE public.application_roles OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 21865)
-- Name: applications; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.applications (
    id uuid NOT NULL,
    name character varying(100) NOT NULL,
    client_id character varying(255) NOT NULL,
    client_secret character varying(255) NOT NULL,
    domain character varying(255) NOT NULL,
    logo character varying(255),
    description text,
    status character varying(20) DEFAULT 'active'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    is_deleted boolean DEFAULT false NOT NULL,
    deleted_at timestamp with time zone,
    deleted_by uuid
);


ALTER TABLE public.applications OWNER TO postgres;

--
-- TOC entry 230 (class 1259 OID 22017)
-- Name: module_role_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.module_role_permissions (
    id uuid NOT NULL,
    module_id uuid NOT NULL,
    application_role_id uuid NOT NULL,
    permission_type character varying(20) NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    is_deleted boolean DEFAULT false NOT NULL,
    deleted_at timestamp with time zone,
    deleted_by uuid
);


ALTER TABLE public.module_role_permissions OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 21893)
-- Name: modules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.modules (
    id uuid NOT NULL,
    item character varying(100),
    name character varying(100) NOT NULL,
    route character varying(255),
    icon character varying(100),
    parent_id uuid,
    application_id uuid,
    sort_order bigint DEFAULT 0,
    status character varying(20) DEFAULT 'active'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    deleted_by uuid
);


ALTER TABLE public.modules OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 21915)
-- Name: organic_units; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.organic_units (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    acronym character varying(20),
    brand character varying(100),
    description text,
    parent_id bigint,
    is_active boolean DEFAULT true,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    is_deleted boolean DEFAULT false NOT NULL,
    deleted_at timestamp with time zone,
    deleted_by uuid,
    cod_dep_sgd character varying(5)
);


ALTER TABLE public.organic_units OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 21914)
-- Name: organic_units_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.organic_units_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.organic_units_id_seq OWNER TO postgres;

--
-- TOC entry 5044 (class 0 OID 0)
-- Dependencies: 221
-- Name: organic_units_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.organic_units_id_seq OWNED BY public.organic_units.id;


--
-- TOC entry 224 (class 1259 OID 21937)
-- Name: structural_positions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.structural_positions (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    code character varying(50) NOT NULL,
    level integer NOT NULL,
    description text,
    is_active boolean DEFAULT true,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    is_deleted boolean DEFAULT false NOT NULL,
    deleted_at timestamp with time zone,
    deleted_by uuid,
    cod_car_sgd character varying(4)
);


ALTER TABLE public.structural_positions OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 21936)
-- Name: structural_positions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.structural_positions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.structural_positions_id_seq OWNER TO postgres;

--
-- TOC entry 5045 (class 0 OID 0)
-- Dependencies: 223
-- Name: structural_positions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.structural_positions_id_seq OWNED BY public.structural_positions.id;


--
-- TOC entry 226 (class 1259 OID 21954)
-- Name: ubigeos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ubigeos (
    id bigint NOT NULL,
    ubigeo_code character varying(10) NOT NULL,
    inei_code character varying(10),
    department character varying(100) NOT NULL,
    province character varying(100) NOT NULL,
    district character varying(100) NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);


ALTER TABLE public.ubigeos OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 21953)
-- Name: ubigeos_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ubigeos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ubigeos_id_seq OWNER TO postgres;

--
-- TOC entry 5046 (class 0 OID 0)
-- Dependencies: 225
-- Name: ubigeos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ubigeos_id_seq OWNED BY public.ubigeos.id;


--
-- TOC entry 229 (class 1259 OID 21993)
-- Name: user_application_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_application_roles (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    application_id uuid NOT NULL,
    application_role_id uuid NOT NULL,
    granted_at timestamp with time zone DEFAULT now(),
    granted_by uuid NOT NULL,
    revoked_at timestamp with time zone,
    revoked_by uuid,
    is_deleted boolean DEFAULT false NOT NULL,
    deleted_at timestamp with time zone,
    deleted_by uuid,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);


ALTER TABLE public.user_application_roles OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 21965)
-- Name: user_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_details (
    id bigint NOT NULL,
    user_id uuid NOT NULL,
    cod_emp_sgd character varying(5),
    first_name character varying(100),
    last_name character varying(100),
    phone character varying(20),
    structural_position_id bigint,
    organic_unit_id bigint,
    ubigeo_id bigint
);


ALTER TABLE public.user_details OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 21964)
-- Name: user_details_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_details_id_seq OWNER TO postgres;

--
-- TOC entry 5047 (class 0 OID 0)
-- Dependencies: 227
-- Name: user_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_details_id_seq OWNED BY public.user_details.id;


--
-- TOC entry 231 (class 1259 OID 22034)
-- Name: user_module_restrictions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_module_restrictions (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    module_id uuid NOT NULL,
    application_id uuid NOT NULL,
    restriction_type character varying(20) NOT NULL,
    max_permission_level character varying(20),
    reason text,
    expires_at timestamp with time zone,
    created_at timestamp with time zone DEFAULT now(),
    created_by uuid NOT NULL,
    updated_at timestamp with time zone DEFAULT now(),
    updated_by uuid,
    is_deleted boolean DEFAULT false NOT NULL,
    deleted_at timestamp with time zone,
    deleted_by uuid
);


ALTER TABLE public.user_module_restrictions OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 21852)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    email character varying(255) NOT NULL,
    dni character varying(8) NOT NULL,
    status character varying(20) DEFAULT 'active'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    is_deleted boolean DEFAULT false NOT NULL,
    deleted_at timestamp with time zone,
    deleted_by uuid
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 4800 (class 2604 OID 21918)
-- Name: organic_units id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organic_units ALTER COLUMN id SET DEFAULT nextval('public.organic_units_id_seq'::regclass);


--
-- TOC entry 4805 (class 2604 OID 21940)
-- Name: structural_positions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.structural_positions ALTER COLUMN id SET DEFAULT nextval('public.structural_positions_id_seq'::regclass);


--
-- TOC entry 4810 (class 2604 OID 21957)
-- Name: ubigeos id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ubigeos ALTER COLUMN id SET DEFAULT nextval('public.ubigeos_id_seq'::regclass);


--
-- TOC entry 4813 (class 2604 OID 21968)
-- Name: user_details id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_details ALTER COLUMN id SET DEFAULT nextval('public.user_details_id_seq'::regclass);


--
-- TOC entry 5026 (class 0 OID 21878)
-- Dependencies: 219
-- Data for Name: application_roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.application_roles (id, name, description, application_id, created_at, updated_at, is_deleted, deleted_at, deleted_by) FROM stdin;
434bed03-d869-4a7b-a9af-404d3c394dd7	default	Rol por defecto para usuarios de la Intranet.	d5a6d085-424d-45c8-8cd4-48f38981af85	2025-12-05 15:55:52.27926-05	2025-12-05 15:55:52.27926-05	f	\N	\N
78908267-5474-4d26-9233-4ac1a847ff06	default	Rol por defecto para usuarios del Central User Manager.	003ce6c1-b1fe-481d-854e-076fbc0882da	2025-12-05 15:55:52.280527-05	2025-12-05 15:55:52.280527-05	f	\N	\N
ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	admin	Rol administrativo con permisos avanzados en Central User Manager.	003ce6c1-b1fe-481d-854e-076fbc0882da	2025-12-05 15:55:52.282247-05	2025-12-05 15:55:52.282247-05	f	\N	\N
d61a0a42-408c-4d0d-85ed-4e1c7003a735	default	Rol básico para acceso a la aplicación de resoluciones.	799de4e5-f480-4f20-a9ba-d3e8174e4193	2025-12-05 15:55:52.283363-05	2025-12-05 15:55:52.283363-05	f	\N	\N
d73fa769-7ab7-43e6-9541-fbd80c0e5644	designer	Rol encargado del diseño y preparación de resoluciones.	799de4e5-f480-4f20-a9ba-d3e8174e4193	2025-12-05 15:55:52.283885-05	2025-12-05 15:55:52.283885-05	f	\N	\N
b0b22ec5-73ff-4551-950c-4b5b8780abc2	approval	Rol responsable de la revisión y aprobación de resoluciones.	799de4e5-f480-4f20-a9ba-d3e8174e4193	2025-12-05 15:55:52.284407-05	2025-12-05 15:55:52.284407-05	f	\N	\N
0771975d-4776-4603-a409-692070094283	sign	Rol autorizado para la firma de resoluciones.	799de4e5-f480-4f20-a9ba-d3e8174e4193	2025-12-05 15:55:52.28546-05	2025-12-05 15:55:52.28546-05	f	\N	\N
276cdc35-8604-415a-8eca-a22eb99c4074	citizen	Rol orientado a consulta y seguimiento por parte de ciudadanos.	799de4e5-f480-4f20-a9ba-d3e8174e4193	2025-12-05 15:55:52.28546-05	2025-12-05 15:55:52.28546-05	f	\N	\N
7dd3c008-6b79-4d5a-982f-2af0bbbe292d	admin	Rol administrativo con control total sobre resoluciones.	799de4e5-f480-4f20-a9ba-d3e8174e4193	2025-12-05 15:55:52.285974-05	2025-12-05 15:55:52.285974-05	f	\N	\N
80f0567b-ed2a-4795-9b4b-55069a518964	default	Rol básico para acceso a la aplicación de certificados.	e0af8703-b3b0-4103-a0f6-64336111dc82	2025-12-05 15:55:52.286498-05	2025-12-05 15:55:52.286498-05	f	\N	\N
0c0e9a41-8e06-4798-8f83-a2e45d262807	admin	Rol administrativo para gestión de certificados.	e0af8703-b3b0-4103-a0f6-64336111dc82	2025-12-05 15:55:52.287507-05	2025-12-05 15:55:52.287507-05	f	\N	\N
ca072fe9-fd40-4553-8759-9feffa54c4c7	super-admin	Rol con control total y configuración avanzada en certificados.	e0af8703-b3b0-4103-a0f6-64336111dc82	2025-12-05 15:55:52.28843-05	2025-12-05 15:55:52.28843-05	f	\N	\N
416cac7b-8d7e-4e8a-bc65-bdeb0cb9eff1	default	Rol básico para acceso a la aplicación de viáticos.	5e18521f-8cf7-4570-b10e-26a432219cfa	2025-12-05 15:55:52.288952-05	2025-12-05 15:55:52.288952-05	f	\N	\N
058426d1-41a3-4137-83dd-aa9887051883	admin	Administrador con permisos especiales y acceso total	778ad41f-39eb-4164-adfd-147c83747224	2025-12-12 16:42:31.314517-05	2025-12-12 16:43:23.826183-05	t	2025-12-12 16:43:23.826183-05	a54e954d-0c61-4861-ab94-d49eaf516672
\.


--
-- TOC entry 5025 (class 0 OID 21865)
-- Dependencies: 218
-- Data for Name: applications; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.applications (id, name, client_id, client_secret, domain, logo, description, status, created_at, updated_at, is_deleted, deleted_at, deleted_by) FROM stdin;
d5a6d085-424d-45c8-8cd4-48f38981af85	intranet-app	intranet-app	nix-os-by-tsaturn-team-sponsored-by-otic	https://intranet.regionayacucho.gob.pe	https://files.regionayacucho.gob.pe/apps/intranet-logo.svg	intranet de la Region Ayacucho	active	2025-12-05 15:55:52.272651-05	2025-12-05 15:55:52.272651-05	f	\N	\N
003ce6c1-b1fe-481d-854e-076fbc0882da	central user manager	cum-app	nix-os-by-tsaturn-team-sponsored-by-otic	https://cum.regionayacucho.gob.pe	https://files.regionayacucho.gob.pe/apps/central-user-manager-logo.svg	Control central de usuarios	active	2025-12-05 15:55:52.273663-05	2025-12-05 15:55:52.273663-05	f	\N	\N
799de4e5-f480-4f20-a9ba-d3e8174e4193	resoluciones	res-app	nix-os-by-tsaturn-team-sponsored-by-otic	https://resoluciones.regionayacucho.gob.pe	https://files.regionayacucho.gob.pe/apps/resoluciones-logo.svg	Gestión de resoluciones administrativas y seguimiento de trámites.	active	2025-12-05 15:55:52.274737-05	2025-12-05 15:55:52.274737-05	f	\N	\N
e0af8703-b3b0-4103-a0f6-64336111dc82	certificados	cert-app	nix-os-by-tsaturn-team-sponsored-by-otic	https://certificados.regionayacucho.gob.pe	https://files.regionayacucho.gob.pe/apps/certificados-logo.svg	Gestión de certificados administrativos y seguimiento de trámites.	active	2025-12-05 15:55:52.275344-05	2025-12-05 15:55:52.275344-05	f	\N	\N
5e18521f-8cf7-4570-b10e-26a432219cfa	viaticos	viaticos-app	nix-os-by-tsaturn-team-sponsored-by-otic	https://viaticos.regionayacucho.gob.pe	https://files.regionayacucho.gob.pe/apps/viaticos-logo.svg	Gestión de viáticos administrativos y seguimiento de trámites.	active	2025-12-05 15:55:52.275344-05	2025-12-05 15:55:52.275344-05	f	\N	\N
bb08d3fc-39eb-43ac-880b-9438fcb7c71f	Mi Nueva Aplicación	mi-app-client-2024	super-secret-password-123	https://mi-app.regionayacucho.gob.pe	https://mi-app.regionayacucho.gob.pe/logo.png	Esta es una aplicación de prueba para el sistema	active	2025-12-12 10:08:00.37775-05	2025-12-12 10:08:00.37775-05	f	\N	\N
778ad41f-39eb-4164-adfd-147c83747224	Aplicación Actualizada	mi-app-client-2025	super-secret-password-123	https://mi-app.regionayacucho.gob.pe	https://mi-app.regionayacucho.gob.pe/logo.png	Descripción actualizada de la aplicación	active	2025-12-12 10:11:00.968248-05	2025-12-12 11:26:37.226491-05	f	\N	\N
\.


--
-- TOC entry 5037 (class 0 OID 22017)
-- Dependencies: 230
-- Data for Name: module_role_permissions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.module_role_permissions (id, module_id, application_role_id, permission_type, created_at, is_deleted, deleted_at, deleted_by) FROM stdin;
59d929c5-7883-4f54-b77f-890e861b826a	07f3f174-cfc8-4c08-8217-69bdfb25d8e4	78908267-5474-4d26-9233-4ac1a847ff06	access	2025-12-05 15:55:52.291341-05	f	\N	\N
ed9406d4-63ba-4c24-8d56-aeee2a196037	07f3f174-cfc8-4c08-8217-69bdfb25d8e4	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	access	2025-12-05 15:55:52.291341-05	f	\N	\N
54493b76-8a5e-4228-92c7-b10388333618	7ed7d783-e1d4-4fc3-8550-88e578cab2cf	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	access	2025-12-05 15:55:52.297986-05	f	\N	\N
3c1b4fc2-cd74-4fb0-98e5-deaf1b368c4b	35bab6b8-028d-4759-882a-cdd0bc0d16ee	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	access	2025-12-05 15:55:52.302921-05	f	\N	\N
2ba7259d-7e34-49d5-89ad-423323130218	feb64b88-5830-48a9-93ca-d1c19a7ad017	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	access	2025-12-05 15:55:52.305199-05	f	\N	\N
cddd6e4e-785c-4621-9408-755d38370208	831166d7-213c-4b5b-884c-eb05eae01527	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	access	2025-12-05 15:55:52.306385-05	f	\N	\N
dbec2653-4813-4ffc-b625-3fa82d1599cb	ff11d439-f872-4f95-8a9e-d959a6202ad8	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	access	2025-12-05 15:55:52.309821-05	f	\N	\N
57817344-4669-4f9c-82a4-1ba27f001199	5bc13897-9ad9-4164-aefa-5a6a5d6c6281	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	access	2025-12-05 15:55:52.312133-05	f	\N	\N
96d153f0-e15e-4583-ac8d-b8dacdea7fe6	03aba1b5-201c-4e03-bc98-dda000e767b5	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	access	2025-12-05 15:55:52.31568-05	f	\N	\N
1ba2be3f-54c4-4dcb-ba99-8a3b395aaab9	eb4e2021-dc7f-49c5-8deb-a05da13a9542	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	access	2025-12-05 15:55:52.31737-05	f	\N	\N
06cc7fae-e9a0-44a2-b413-d6eda626a5f7	d88b0bce-f48e-499e-bac3-0ccd96729125	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	access	2025-12-05 15:55:52.319183-05	f	\N	\N
293f6fe8-16c6-4abf-b074-fbec0b6fd3b6	b868346a-1974-40d8-94c0-10643ccf5d4e	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	access	2025-12-05 15:55:52.3208-05	f	\N	\N
8f8b9d14-3d58-40d7-aad1-8f41fbdf0160	ab1af32e-7334-464a-9de0-e3aa08d37390	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	access	2025-12-05 15:55:52.322369-05	f	\N	\N
12f6c09c-e001-4348-a961-bb440e374030	7aa2ff83-a79e-445c-917c-80d212218c00	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	access	2025-12-05 15:55:52.324374-05	f	\N	\N
8f6cf826-99a0-46f9-9884-716222b89b57	f09ad030-e782-4336-b498-e151559eefab	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	access	2025-12-05 15:55:52.325413-05	f	\N	\N
0af04367-af4b-4f68-9755-0b73aa421485	e340217d-a882-4ea4-9967-9e58c31d4c34	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	access	2025-12-05 15:55:52.328383-05	f	\N	\N
481f560f-e5bc-4b2a-acda-53c8f25ce74d	5383c350-82bb-41bd-8ffc-49b79473830c	d61a0a42-408c-4d0d-85ed-4e1c7003a735	access	2025-12-05 15:55:52.329503-05	f	\N	\N
09dd33a2-9497-4a4c-ace9-6d31c4ab6b93	5383c350-82bb-41bd-8ffc-49b79473830c	7dd3c008-6b79-4d5a-982f-2af0bbbe292d	access	2025-12-05 15:55:52.329503-05	f	\N	\N
4b995a2c-02a4-43ee-aa33-549e62f0bb65	5ce91401-ba32-40f7-aeb4-38201dcfdba7	80f0567b-ed2a-4795-9b4b-55069a518964	access	2025-12-05 15:55:52.333403-05	f	\N	\N
433130e3-c51d-4dae-bb0f-c5d250beffc6	5ce91401-ba32-40f7-aeb4-38201dcfdba7	0c0e9a41-8e06-4798-8f83-a2e45d262807	access	2025-12-05 15:55:52.333403-05	f	\N	\N
ca1246de-898e-4d03-94b5-0dc332e366f2	1629756e-5eea-4ded-9d8c-f4bf39c164df	416cac7b-8d7e-4e8a-bc65-bdeb0cb9eff1	access	2025-12-05 15:55:52.337809-05	f	\N	\N
\.


--
-- TOC entry 5027 (class 0 OID 21893)
-- Dependencies: 220
-- Data for Name: modules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.modules (id, item, name, route, icon, parent_id, application_id, sort_order, status, created_at, updated_at, deleted_at, deleted_by) FROM stdin;
07f3f174-cfc8-4c08-8217-69bdfb25d8e4	Menú	Dashboard	/dashboard	LayoutDashboard	\N	003ce6c1-b1fe-481d-854e-076fbc0882da	1	active	2025-12-05 15:55:52.291341-05	2025-12-05 15:55:52.291341-05	\N	\N
7ed7d783-e1d4-4fc3-8550-88e578cab2cf	Gestión	Usuarios (lista)	/dashboard/users	UsersRound	7ed7d783-e1d4-4fc3-8550-88e578cab2cf	003ce6c1-b1fe-481d-854e-076fbc0882da	11	active	2025-12-05 15:55:52.297986-05	2025-12-05 15:55:52.300434-05	\N	\N
35bab6b8-028d-4759-882a-cdd0bc0d16ee	Gestión	Posiciones Estructurales	/dashboard/structural-positions	CircleQuestionMark	7ed7d783-e1d4-4fc3-8550-88e578cab2cf	003ce6c1-b1fe-481d-854e-076fbc0882da	12	active	2025-12-05 15:55:52.302921-05	2025-12-05 15:55:52.302921-05	\N	\N
feb64b88-5830-48a9-93ca-d1c19a7ad017	Gestión	Unidades Orgánicas	/dashboard/organic-units	Hexagon	7ed7d783-e1d4-4fc3-8550-88e578cab2cf	003ce6c1-b1fe-481d-854e-076fbc0882da	13	active	2025-12-05 15:55:52.305199-05	2025-12-05 15:55:52.305199-05	\N	\N
831166d7-213c-4b5b-884c-eb05eae01527	Gestión	Aplicaciones (lista)	/dashboard/applications	Boxes	831166d7-213c-4b5b-884c-eb05eae01527	003ce6c1-b1fe-481d-854e-076fbc0882da	21	active	2025-12-05 15:55:52.306385-05	2025-12-05 15:55:52.307535-05	\N	\N
ff11d439-f872-4f95-8a9e-d959a6202ad8	Gestión	Módulos	/dashboard/modules	Package	831166d7-213c-4b5b-884c-eb05eae01527	003ce6c1-b1fe-481d-854e-076fbc0882da	22	active	2025-12-05 15:55:52.309821-05	2025-12-05 15:55:52.309821-05	\N	\N
5bc13897-9ad9-4164-aefa-5a6a5d6c6281	Seguridad	Seguridad	/dashboard/security	Shield	\N	003ce6c1-b1fe-481d-854e-076fbc0882da	30	active	2025-12-05 15:55:52.312133-05	2025-12-05 15:55:52.312133-05	\N	\N
03aba1b5-201c-4e03-bc98-dda000e767b5	Seguridad	Sesiones Activas	/dashboard/security/active-sessions	Activity	5bc13897-9ad9-4164-aefa-5a6a5d6c6281	003ce6c1-b1fe-481d-854e-076fbc0882da	31	active	2025-12-05 15:55:52.31568-05	2025-12-05 15:55:52.31568-05	\N	\N
eb4e2021-dc7f-49c5-8deb-a05da13a9542	Seguridad	Logs de Auditoría	/dashboard/security/audit-logs	ScrollText	5bc13897-9ad9-4164-aefa-5a6a5d6c6281	003ce6c1-b1fe-481d-854e-076fbc0882da	32	active	2025-12-05 15:55:52.31737-05	2025-12-05 15:55:52.31737-05	\N	\N
d88b0bce-f48e-499e-bac3-0ccd96729125	Seguridad	Roles y Permisos	/dashboard/security/roles-permissions	ShieldCheck	\N	003ce6c1-b1fe-481d-854e-076fbc0882da	40	active	2025-12-05 15:55:52.319183-05	2025-12-05 15:55:52.319183-05	\N	\N
b868346a-1974-40d8-94c0-10643ccf5d4e	Seguridad	Roles de Aplicación	/dashboard/security/application-roles	Shield	d88b0bce-f48e-499e-bac3-0ccd96729125	003ce6c1-b1fe-481d-854e-076fbc0882da	41	active	2025-12-05 15:55:52.3208-05	2025-12-05 15:55:52.3208-05	\N	\N
ab1af32e-7334-464a-9de0-e3aa08d37390	Seguridad	Asignación de Roles	/dashboard/security/user-roles	UserCog	d88b0bce-f48e-499e-bac3-0ccd96729125	003ce6c1-b1fe-481d-854e-076fbc0882da	42	active	2025-12-05 15:55:52.322369-05	2025-12-05 15:55:52.322369-05	\N	\N
7aa2ff83-a79e-445c-917c-80d212218c00	Seguridad	Restricción de Usuario	/dashboard/security/user-restrictions	Ban	d88b0bce-f48e-499e-bac3-0ccd96729125	003ce6c1-b1fe-481d-854e-076fbc0882da	43	active	2025-12-05 15:55:52.324374-05	2025-12-05 15:55:52.324374-05	\N	\N
f09ad030-e782-4336-b498-e151559eefab	Configuración	Ajustes	/dashboard/settings	Settings2	\N	003ce6c1-b1fe-481d-854e-076fbc0882da	50	active	2025-12-05 15:55:52.325413-05	2025-12-05 15:55:52.325413-05	\N	\N
e340217d-a882-4ea4-9967-9e58c31d4c34	Configuración	Cuenta	/dashboard/account	UserCircle	\N	003ce6c1-b1fe-481d-854e-076fbc0882da	51	active	2025-12-05 15:55:52.328383-05	2025-12-05 15:55:52.328383-05	\N	\N
5383c350-82bb-41bd-8ffc-49b79473830c	Menú	Dashboard	/dashboard	LayoutDashboard	\N	799de4e5-f480-4f20-a9ba-d3e8174e4193	1	active	2025-12-05 15:55:52.329503-05	2025-12-05 15:55:52.329503-05	\N	\N
5ce91401-ba32-40f7-aeb4-38201dcfdba7	Menú	Dashboard	/dashboard	LayoutDashboard	\N	e0af8703-b3b0-4103-a0f6-64336111dc82	1	active	2025-12-05 15:55:52.333403-05	2025-12-05 15:55:52.333403-05	\N	\N
1629756e-5eea-4ded-9d8c-f4bf39c164df	Menú	Dashboard	/dashboard	LayoutDashboard	\N	5e18521f-8cf7-4570-b10e-26a432219cfa	1	active	2025-12-05 15:55:52.337809-05	2025-12-05 15:55:52.337809-05	\N	\N
\.


--
-- TOC entry 5029 (class 0 OID 21915)
-- Dependencies: 222
-- Data for Name: organic_units; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.organic_units (id, name, acronym, brand, description, parent_id, is_active, created_at, updated_at, is_deleted, deleted_at, deleted_by, cod_dep_sgd) FROM stdin;
1	Gobernación Regional	GR	https://resources.regionayacucho.gob.pe/gr_brand.webp	Máxima autoridad ejecutiva del gobierno regional de Ayacucho.	\N	t	2025-12-05 15:55:51.139985-05	2025-12-05 15:55:51.139985-05	f	\N	\N	\N
2	Consejo Regional	CR	https://resources.regionayacucho.gob.pe/cr_brand.webp	Órgano normativo, resolutivo y fiscalizador del gobierno regional.	\N	t	2025-12-05 15:55:51.141788-05	2025-12-05 15:55:51.141788-05	f	\N	\N	\N
3	Vicegobernación Regional	VGR	https://resources.regionayacucho.gob.pe/vgr_brand.webp	Segunda autoridad ejecutiva del gobierno regional.	\N	t	2025-12-05 15:55:51.143228-05	2025-12-05 15:55:51.143228-05	f	\N	\N	\N
4	Gerencia General Regional	GG	https://resources.regionayacucho.gob.pe/gg_brand.webp	Órgano de dirección ejecutiva del gobierno regional.	\N	t	2025-12-05 15:55:51.143918-05	2025-12-05 15:55:51.143918-05	f	\N	\N	\N
5	Gerencia Regional de Desarrollo Económico	GRDE	https://resources.regionayacucho.gob.pe/grde_brand.webp	Encargada del desarrollo económico y productivo regional.	\N	t	2025-12-05 15:55:51.145166-05	2025-12-05 15:55:51.145166-05	f	\N	\N	\N
6	Gerencia Regional de Desarrollo Social	GRDS	https://resources.regionayacucho.gob.pe/grds_brand.webp	Responsable del desarrollo social y bienestar ciudadano.	\N	t	2025-12-05 15:55:51.146243-05	2025-12-05 15:55:51.146243-05	f	\N	\N	\N
7	Gerencia Regional de Infraestructura	GRI	https://resources.regionayacucho.gob.pe/gri_brand.webp	Gestiona proyectos de infraestructura regional.	\N	t	2025-12-05 15:55:51.14685-05	2025-12-05 15:55:51.14685-05	f	\N	\N	\N
8	Gerencia Regional de Planeamiento, Presupuesto y Acondicionamiento Territorial	GRPPAT	https://resources.regionayacucho.gob.pe/grppat_brand.webp	Planificación, presupuesto y ordenamiento territorial.	\N	t	2025-12-05 15:55:51.147514-05	2025-12-05 15:55:51.147514-05	f	\N	\N	\N
9	Subgerencia de Estudios de Preinversión y Estadística	SGEPE	https://resources.regionayacucho.gob.pe/sgepe_brand.webp	Estudios de preinversión y gestión estadística regional.	\N	t	2025-12-05 15:55:51.148545-05	2025-12-05 15:55:51.148545-05	f	\N	\N	\N
10	Gerencia Regional de Recursos Naturales y Gestión del Medio Ambiente	GRRNGMA	https://resources.regionayacucho.gob.pe/grrngma_brand.webp	Gestión de recursos naturales y protección ambiental.	\N	t	2025-12-05 15:55:51.149068-05	2025-12-05 15:55:51.149068-05	f	\N	\N	\N
11	Oficina de Imagen Institucional	OII	https://resources.regionayacucho.gob.pe/oii_brand.webp	Comunicación y relaciones públicas institucionales.	\N	t	2025-12-05 15:55:51.149943-05	2025-12-05 15:55:51.149943-05	f	\N	\N	\N
12	Procuraduría Pública Regional	PPR	https://resources.regionayacucho.gob.pe/ppr_brand.webp	Defensa jurídica de los intereses del Estado regional.	\N	t	2025-12-05 15:55:51.151135-05	2025-12-05 15:55:51.151135-05	f	\N	\N	\N
13	Oficina Subregional Cangallo	OSRC	https://resources.regionayacucho.gob.pe/osrc_brand.webp	Oficina descentralizada para la provincia de Cangallo.	\N	t	2025-12-05 15:55:51.151663-05	2025-12-05 15:55:51.151663-05	f	\N	\N	\N
14	Oficina Subregional Fajardo	OSRF	https://resources.regionayacucho.gob.pe/osrf_brand.webp	Oficina descentralizada para la provincia de Fajardo.	\N	t	2025-12-05 15:55:51.152992-05	2025-12-05 15:55:51.152992-05	f	\N	\N	\N
15	Oficina Subregional Huanta	OSRHTA	https://resources.regionayacucho.gob.pe/osrhta_brand.webp	Oficina descentralizada para la provincia de Huanta.	\N	t	2025-12-05 15:55:51.153566-05	2025-12-05 15:55:51.153566-05	f	\N	\N	\N
16	Oficina Subregional La Mar	OSRLM	https://resources.regionayacucho.gob.pe/osrlm_brand.webp	Oficina descentralizada para la provincia de La Mar.	\N	t	2025-12-05 15:55:51.15456-05	2025-12-05 15:55:51.15456-05	f	\N	\N	\N
17	Oficina Subregional Lucanas	OSRL	https://resources.regionayacucho.gob.pe/osrl_brand.webp	Oficina descentralizada para la provincia de Lucanas.	\N	t	2025-12-05 15:55:51.155134-05	2025-12-05 15:55:51.155134-05	f	\N	\N	\N
18	Oficina Subregional Parinacochas	OSRP	https://resources.regionayacucho.gob.pe/osrp_brand.webp	Oficina descentralizada para la provincia de Parinacochas.	\N	t	2025-12-05 15:55:51.155641-05	2025-12-05 15:55:51.155641-05	f	\N	\N	\N
19	Oficina Subregional Vilcashuamán	OSVH	https://resources.regionayacucho.gob.pe/osvh_brand.webp	Oficina descentralizada para la provincia de Vilcashuamán.	\N	t	2025-12-05 15:55:51.156255-05	2025-12-05 15:55:51.156255-05	f	\N	\N	\N
20	Unidad de Prevención y Gestión de Conflictos Sociales	UPGCS	https://resources.regionayacucho.gob.pe/upgcs_brand.webp	Prevención y mediación de conflictos sociales.	\N	t	2025-12-05 15:55:51.156778-05	2025-12-05 15:55:51.156778-05	f	\N	\N	\N
21	Secretaría General	SG	https://resources.regionayacucho.gob.pe/sg_brand.webp	Gestión documentaria y apoyo administrativo general.	\N	t	2025-12-05 15:55:51.15729-05	2025-12-05 15:55:51.15729-05	f	\N	\N	\N
22	Mesa de Partes	MP	https://resources.regionayacucho.gob.pe/mp_brand.webp	Recepción y distribución de documentos institucionales.	\N	t	2025-12-05 15:55:51.1578-05	2025-12-05 15:55:51.1578-05	f	\N	\N	\N
23	Subgerencia de Operación y Mantenimiento de Equipo Mecánico	SGOMEM	https://resources.regionayacucho.gob.pe/sgomem_brand.webp	Operación y mantenimiento de maquinaria institucional.	\N	t	2025-12-05 15:55:51.1578-05	2025-12-05 15:55:51.1578-05	f	\N	\N	\N
24	Unidad de Coordinación Lima	UCL	https://resources.regionayacucho.gob.pe/ucl_brand.webp	Representación y coordinación en la capital del país.	\N	t	2025-12-05 15:55:51.158307-05	2025-12-05 15:55:51.158307-05	f	\N	\N	\N
25	Oficina Regional de Administración	ORADM	https://resources.regionayacucho.gob.pe/oradm_brand.webp	Gestión administrativa y financiera institucional.	\N	t	2025-12-05 15:55:51.158836-05	2025-12-05 15:55:51.158836-05	f	\N	\N	\N
26	Subgerencia de Estudios Definitivos y Expedientes Técnicos	SGEDET	https://resources.regionayacucho.gob.pe/sgedet_brand.webp	Elaboración de estudios definitivos para proyectos.	\N	t	2025-12-05 15:55:51.16008-05	2025-12-05 15:55:51.16008-05	f	\N	\N	\N
27	Unidad Operativa Las Cabezadas	UOC	https://resources.regionayacucho.gob.pe/uoc_brand.webp	Operación del proyecto hidráulico Las Cabezadas.	\N	t	2025-12-05 15:55:51.160607-05	2025-12-05 15:55:51.160607-05	f	\N	\N	\N
28	Dirección Subregional Territorial VRAEM	DSTVRAEM	https://resources.regionayacucho.gob.pe/dstvraem_brand.webp	Gestión territorial en el Valle de los ríos Apurímac, Ene y Mantaro.	\N	t	2025-12-05 15:55:51.161655-05	2025-12-05 15:55:51.161655-05	f	\N	\N	\N
29	Dirección Regional de Comercio Exterior y Turismo	DIRCETUR	https://resources.regionayacucho.gob.pe/dircetur_brand.webp	Promoción del comercio exterior y desarrollo turístico.	\N	t	2025-12-05 15:55:51.162243-05	2025-12-05 15:55:51.162243-05	f	\N	\N	\N
30	Dirección Regional de Trabajo y Promoción del Empleo	DRTPE	https://resources.regionayacucho.gob.pe/drtpe_brand.webp	Promoción del empleo y relaciones laborales.	\N	t	2025-12-05 15:55:51.162757-05	2025-12-05 15:55:51.162757-05	f	\N	\N	\N
31	Dirección Regional de Vivienda, Construcción y Saneamiento	DRVCS	https://resources.regionayacucho.gob.pe/drvcs_brand.webp	Desarrollo habitacional y saneamiento regional.	\N	t	2025-12-05 15:55:51.163311-05	2025-12-05 15:55:51.163311-05	f	\N	\N	\N
32	Dirección Regional de Energía, Minas e Hidrocarburos	DREMH	https://resources.regionayacucho.gob.pe/dremh_brand.webp	Supervisión y promoción del sector energético y minero.	\N	t	2025-12-05 15:55:51.163828-05	2025-12-05 15:55:51.163828-05	f	\N	\N	\N
33	Dirección Regional de la Producción	DRP	https://resources.regionayacucho.gob.pe/drp_brand.webp	Desarrollo productivo, pesca e industria regional.	\N	t	2025-12-05 15:55:51.164342-05	2025-12-05 15:55:51.164342-05	f	\N	\N	\N
34	Oficina de Gestión de Recursos Humanos	OGRH	https://resources.regionayacucho.gob.pe/ogrh_brand.webp	Administración del talento humano institucional.	\N	t	2025-12-05 15:55:51.165579-05	2025-12-05 15:55:51.165579-05	f	\N	\N	\N
35	Oficina de Tecnologías de Información y Comunicaciones	OTIC	https://resources.regionayacucho.gob.pe/otic_brand.webp	Gestión de tecnologías de información y comunicaciones.	\N	t	2025-12-05 15:55:51.166152-05	2025-12-05 15:55:51.166152-05	f	\N	\N	\N
36	Oficina Regional de Asesoría Jurídica	ORAJ	https://resources.regionayacucho.gob.pe/oraj_brand.webp	Asesoramiento legal y jurídico institucional.	\N	t	2025-12-05 15:55:51.166772-05	2025-12-05 15:55:51.166772-05	f	\N	\N	\N
37	Oficina de Abastecimiento y Patrimonio	OAPF	https://resources.regionayacucho.gob.pe/oapf_brand.webp	Gestión de adquisiciones y control patrimonial.	\N	t	2025-12-05 15:55:51.168434-05	2025-12-05 15:55:51.168434-05	f	\N	\N	\N
38	Oficina de Contabilidad	OC	https://resources.regionayacucho.gob.pe/oc_brand.webp	Registro y control contable institucional.	\N	t	2025-12-05 15:55:51.169424-05	2025-12-05 15:55:51.169424-05	f	\N	\N	\N
39	Oficina de Tesorería	OT	https://resources.regionayacucho.gob.pe/ot_brand.webp	Gestión financiera y pagos institucionales.	\N	t	2025-12-05 15:55:51.169424-05	2025-12-05 15:55:51.169424-05	f	\N	\N	\N
40	Gerencia Regional de Gestión de Riesgo de Desastres, Seguridad y Defensa	GRGRDSD	https://resources.regionayacucho.gob.pe/grgrdsd_brand.webp	Gestión de riesgos, seguridad y defensa civil.	\N	t	2025-12-05 15:55:51.169949-05	2025-12-05 15:55:51.169949-05	f	\N	\N	\N
\.


--
-- TOC entry 5031 (class 0 OID 21937)
-- Dependencies: 224
-- Data for Name: structural_positions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.structural_positions (id, name, code, level, description, is_active, created_at, updated_at, is_deleted, deleted_at, deleted_by, cod_car_sgd) FROM stdin;
1	ABOGADO	ABO-001	2	Profesional en derecho encargado de asuntos legales.	t	2025-12-05 15:55:51.046155-05	2025-12-05 15:55:51.046155-05	f	\N	\N	\N
2	ABOGADO AREA CIVIL	ABO-003	2	Especialista en derecho civil y procedimientos civiles.	t	2025-12-05 15:55:51.049159-05	2025-12-05 15:55:51.049159-05	f	\N	\N	\N
3	ABOGADO CALIFICADOR	ABO-004	2	Encargado de calificar y evaluar documentos legales.	t	2025-12-05 15:55:51.050383-05	2025-12-05 15:55:51.050383-05	f	\N	\N	\N
4	ABOGADO DEL AREA ARBITRAJE	ABO-005	2	Especialista en procesos de arbitraje y mediación.	t	2025-12-05 15:55:51.050982-05	2025-12-05 15:55:51.050982-05	f	\N	\N	\N
5	ADMINISTRADOR	ADM-006	2	Responsable de la gestión administrativa general.	t	2025-12-05 15:55:51.052685-05	2025-12-05 15:55:51.052685-05	f	\N	\N	\N
6	ADMINISTRADOR DE OBRA	ADM-007	2	Encargado de la administración de proyectos de obra.	t	2025-12-05 15:55:51.054199-05	2025-12-05 15:55:51.054199-05	f	\N	\N	\N
7	ADMINISTRADOR DE PROYECTO Y/O META	ADM-008	2	Responsable de la administración de proyectos específicos.	t	2025-12-05 15:55:51.054735-05	2025-12-05 15:55:51.054735-05	f	\N	\N	\N
8	ADMINISTRADOR DE REDES DE LA UNIDAD DE INFORMATICA	ADM-009	2	Especialista en administración de redes informáticas.	t	2025-12-05 15:55:51.055433-05	2025-12-05 15:55:51.055433-05	f	\N	\N	\N
9	ADMINISTRADOR DEL PORTAL WEB	ADM-010	2	Encargado de la gestión y mantenimiento del portal web.	t	2025-12-05 15:55:51.056987-05	2025-12-05 15:55:51.056987-05	f	\N	\N	\N
10	ADMINISTRADORA	ADM-011	2	Responsable de la gestión administrativa general.	t	2025-12-05 15:55:51.058169-05	2025-12-05 15:55:51.058169-05	f	\N	\N	\N
11	ADMINISTRADORA DE OBRA	ADM-012	2	Encargada de la administración de proyectos de obra.	t	2025-12-05 15:55:51.058822-05	2025-12-05 15:55:51.058822-05	f	\N	\N	\N
12	ANALISTA DE COMUNICACION SOCIAL	ANA-013	2	Especialista en análisis de comunicación e imagen institucional.	t	2025-12-05 15:55:51.059863-05	2025-12-05 15:55:51.059863-05	f	\N	\N	\N
13	ANALISTA DE CUENTAS	ANA-014	2	Encargado del análisis y seguimiento de cuentas.	t	2025-12-05 15:55:51.060372-05	2025-12-05 15:55:51.060372-05	f	\N	\N	\N
14	ANALISTA EN SEGUIMIENTO Y CONTROL DE ACTIVIDADES	ANA-015	2	Responsable del monitoreo y control de actividades institucionales.	t	2025-12-05 15:55:51.061584-05	2025-12-05 15:55:51.061584-05	f	\N	\N	\N
15	ANALISTA TECNICO LEGAL	ANA-016	2	Especialista en análisis técnico de asuntos legales.	t	2025-12-05 15:55:51.062113-05	2025-12-05 15:55:51.062113-05	f	\N	\N	\N
16	APOYO	APO-017	2	Personal de apoyo en diversas funciones.	t	2025-12-05 15:55:51.062652-05	2025-12-05 15:55:51.062652-05	f	\N	\N	\N
17	APOYO ADMINISTRATIVO	APO-018	2	Asistente en labores administrativas generales.	t	2025-12-05 15:55:51.063296-05	2025-12-05 15:55:51.063296-05	f	\N	\N	\N
18	APOYO ADMINISTRATIVO O/S	APO-019	2	Apoyo administrativo para órdenes de servicio.	t	2025-12-05 15:55:51.063841-05	2025-12-05 15:55:51.063841-05	f	\N	\N	\N
19	APOYO ADQUISISCION DE O/C Y O/S	APO-020	2	Asistente en procesos de adquisición de órdenes de compra y servicio.	t	2025-12-05 15:55:51.064363-05	2025-12-05 15:55:51.064363-05	f	\N	\N	\N
20	APOYO EN CONTRATACIONES	APO-021	2	Asistente en procesos de contratación pública.	t	2025-12-05 15:55:51.065304-05	2025-12-05 15:55:51.065304-05	f	\N	\N	\N
21	APOYO LEGAL	APO-022	2	Asistente en asuntos legales e jurídicos.	t	2025-12-05 15:55:51.066277-05	2025-12-05 15:55:51.066277-05	f	\N	\N	\N
22	APOYO REMUNERACIONES	APO-023	2	Asistente en el área de remuneraciones y planillas.	t	2025-12-05 15:55:51.066831-05	2025-12-05 15:55:51.066831-05	f	\N	\N	\N
23	APOYO TECNICO	APO-024	2	Asistente técnico en diversas especialidades.	t	2025-12-05 15:55:51.067341-05	2025-12-05 15:55:51.067341-05	f	\N	\N	\N
24	APOYO TECNICO ADMINISTRATIVO	APO-025	2	Asistente técnico en labores administrativas.	t	2025-12-05 15:55:51.067876-05	2025-12-05 15:55:51.067876-05	f	\N	\N	\N
25	ASESOR	ASE-026	2	Profesional que brinda asesoría especializada.	t	2025-12-05 15:55:51.068582-05	2025-12-05 15:55:51.068582-05	f	\N	\N	\N
26	ASESOR DE GOBERNACION	ASE-027	2	Consejero especializado para la gobernación regional.	t	2025-12-05 15:55:51.069333-05	2025-12-05 15:55:51.069333-05	f	\N	\N	\N
27	ASESOR I	ASE-028	2	Asesor de primer nivel en áreas específicas.	t	2025-12-05 15:55:51.069876-05	2025-12-05 15:55:51.069876-05	f	\N	\N	\N
28	ASESOR I GERENCIA GENERAL	ASE-029	2	Asesor especializado de la gerencia general.	t	2025-12-05 15:55:51.069876-05	2025-12-05 15:55:51.069876-05	f	\N	\N	\N
29	ASESOR LEGAL	ASE-030	2	Consejero especializado en asuntos jurídicos.	t	2025-12-05 15:55:51.070387-05	2025-12-05 15:55:51.070387-05	f	\N	\N	\N
30	ASISTENTE	ASI-031	2	Personal de asistencia en diversas funciones.	t	2025-12-05 15:55:51.0709-05	2025-12-05 15:55:51.0709-05	f	\N	\N	\N
31	ASISTENTE ADMINISTRATIVO	ASI-032	2	Asistente en labores administrativas generales.	t	2025-12-05 15:55:51.071516-05	2025-12-05 15:55:51.071516-05	f	\N	\N	\N
32	ASISTENTE ADMINISTRATIVO I	ASI-033	2	Asistente administrativo de primer nivel.	t	2025-12-05 15:55:51.072028-05	2025-12-05 15:55:51.072028-05	f	\N	\N	\N
33	ASISTENTE ADMINISTRATIVO II	ASI-034	2	Asistente administrativo de segundo nivel.	t	2025-12-05 15:55:51.072599-05	2025-12-05 15:55:51.072599-05	f	\N	\N	\N
34	ASISTENTE ADMINISTRATIVO III	ASI-035	2	Asistente administrativo de tercer nivel.	t	2025-12-05 15:55:51.073134-05	2025-12-05 15:55:51.073134-05	f	\N	\N	\N
35	ASISTENTE ADMINISTRATIVO PROVIAS PCD-CG	ASI-036	2	Asistente administrativo especializado en Provias.	t	2025-12-05 15:55:51.073824-05	2025-12-05 15:55:51.073824-05	f	\N	\N	\N
36	ASISTENTE DE ARQUITECTURA	ASI-037	2	Asistente técnico en proyectos arquitectónicos.	t	2025-12-05 15:55:51.073824-05	2025-12-05 15:55:51.073824-05	f	\N	\N	\N
37	ASISTENTE DE AUDITORIA	ASI-038	2	Asistente en procesos de auditoría institucional.	t	2025-12-05 15:55:51.074338-05	2025-12-05 15:55:51.074338-05	f	\N	\N	\N
38	ASISTENTE DE DESPACHO	ASI-039	2	Asistente en labores de despacho institucional.	t	2025-12-05 15:55:51.074859-05	2025-12-05 15:55:51.074859-05	f	\N	\N	\N
39	ASISTENTE DE ESTRUCTURAS	ASI-040	2	Asistente técnico en ingeniería estructural.	t	2025-12-05 15:55:51.07544-05	2025-12-05 15:55:51.07544-05	f	\N	\N	\N
40	ASISTENTE DE GESTION DE RECURSOS HUMANOS	ASI-041	2	Asistente en la gestión del personal institucional.	t	2025-12-05 15:55:51.07598-05	2025-12-05 15:55:51.07598-05	f	\N	\N	\N
41	ASISTENTE DE INFRAESTRUCTURA	ASI-042	2	Asistente técnico en proyectos de infraestructura.	t	2025-12-05 15:55:51.077028-05	2025-12-05 15:55:51.077028-05	f	\N	\N	\N
42	ASISTENTE DE INGENIERIA	ASI-043	2	Asistente técnico en proyectos de ingeniería.	t	2025-12-05 15:55:51.077568-05	2025-12-05 15:55:51.077568-05	f	\N	\N	\N
43	ASISTENTE DE MECANICA DE SUELOS	ASI-044	2	Asistente técnico en estudios de mecánica de suelos.	t	2025-12-05 15:55:51.078659-05	2025-12-05 15:55:51.078659-05	f	\N	\N	\N
44	ASISTENTE DE RESIDENTE DE OBRA	ASI-045	2	Asistente del residente en obras de construcción.	t	2025-12-05 15:55:51.079381-05	2025-12-05 15:55:51.079381-05	f	\N	\N	\N
45	ASISTENTE DE SISTEMAS	ASI-046	2	Asistente técnico en sistemas informáticos.	t	2025-12-05 15:55:51.0799-05	2025-12-05 15:55:51.0799-05	f	\N	\N	\N
46	ASISTENTE EN INGENIERIA CIVIL	ASI-047	2	Asistente técnico en proyectos de ingeniería civil.	t	2025-12-05 15:55:51.080489-05	2025-12-05 15:55:51.080489-05	f	\N	\N	\N
47	ASISTENTE EN SERVICIOS DE INFRAESTRUCTURA II	ASI-048	2	Asistente de segundo nivel en servicios de infraestructura.	t	2025-12-05 15:55:51.081546-05	2025-12-05 15:55:51.081546-05	f	\N	\N	\N
48	ASISTENTE EN SERVICIOS DE SALUD	ASI-049	2	Asistente en la prestación de servicios de salud.	t	2025-12-05 15:55:51.082088-05	2025-12-05 15:55:51.082088-05	f	\N	\N	\N
49	ASISTENTE EN SERVICIOS SOCIALES II	ASI-050	2	Asistente de segundo nivel en servicios sociales.	t	2025-12-05 15:55:51.082806-05	2025-12-05 15:55:51.082806-05	f	\N	\N	\N
50	ASISTENTE JURIDICO	ASI-051	2	Asistente en asuntos jurídicos e legales.	t	2025-12-05 15:55:51.083315-05	2025-12-05 15:55:51.083315-05	f	\N	\N	\N
51	ASISTENTE LEGAL	ASI-052	2	Asistente en procedimientos legales.	t	2025-12-05 15:55:51.084163-05	2025-12-05 15:55:51.084163-05	f	\N	\N	\N
52	ASISTENTE LEGAL EN PROCESOS ARBITRALES	ASI-053	2	Asistente especializado en procesos de arbitraje.	t	2025-12-05 15:55:51.085519-05	2025-12-05 15:55:51.085519-05	f	\N	\N	\N
53	ASISTENTE SOCIAL	ASI-054	2	Profesional en trabajo social y asistencia social.	t	2025-12-05 15:55:51.086023-05	2025-12-05 15:55:51.086023-05	f	\N	\N	\N
54	ASISTENTE SOCIAL II	ASI-055	2	Asistente social de segundo nivel.	t	2025-12-05 15:55:51.086535-05	2025-12-05 15:55:51.086535-05	f	\N	\N	\N
55	ASISTENTE SOCIAL IV	ASI-056	2	Asistente social de cuarto nivel.	t	2025-12-05 15:55:51.087133-05	2025-12-05 15:55:51.087133-05	f	\N	\N	\N
56	ASISTENTE TECNICO	ASI-057	2	Asistente técnico en diversas especialidades.	t	2025-12-05 15:55:51.08774-05	2025-12-05 15:55:51.08774-05	f	\N	\N	\N
57	ASISTENTE TECNICO II	ASI-058	2	Asistente técnico de segundo nivel.	t	2025-12-05 15:55:51.088407-05	2025-12-05 15:55:51.088407-05	f	\N	\N	\N
58	ASISTENTE TECNICO LEGAL	ASI-059	2	Asistente técnico en asuntos legales.	t	2025-12-05 15:55:51.088925-05	2025-12-05 15:55:51.088925-05	f	\N	\N	\N
59	AUDITOR	AUD-060	2	Profesional encargado de auditorías institucionales.	t	2025-12-05 15:55:51.089432-05	2025-12-05 15:55:51.089432-05	f	\N	\N	\N
60	AUXILIAR	AUX-061	2	Personal auxiliar en diversas funciones.	t	2025-12-05 15:55:51.090515-05	2025-12-05 15:55:51.090515-05	f	\N	\N	\N
61	AUXILIAR ADMINISTRATIVO	AUX-062	2	Auxiliar en labores administrativas básicas.	t	2025-12-05 15:55:51.091031-05	2025-12-05 15:55:51.091031-05	f	\N	\N	\N
62	AUXILIAR DE SISTEMAS ADMINISTRATIVO I	AUX-063	2	Auxiliar de primer nivel en sistemas administrativos.	t	2025-12-05 15:55:51.091552-05	2025-12-05 15:55:51.091552-05	f	\N	\N	\N
63	BIOLOGO III	BIO-064	2	Biólogo de tercer nivel especializado.	t	2025-12-05 15:55:51.091552-05	2025-12-05 15:55:51.091552-05	f	\N	\N	\N
64	BROMATOLOGO NUTRICIONISTA	BRO-065	2	Especialista en bromatología y nutrición.	t	2025-12-05 15:55:51.092595-05	2025-12-05 15:55:51.092595-05	f	\N	\N	\N
65	CADISTA	CAD-066	2	Especialista en diseño asistido por computadora.	t	2025-12-05 15:55:51.093205-05	2025-12-05 15:55:51.093205-05	f	\N	\N	\N
66	CALCULO DE PENALIDAD	CAL-067	2	Encargado del cálculo de penalidades contractuales.	t	2025-12-05 15:55:51.093716-05	2025-12-05 15:55:51.093716-05	f	\N	\N	\N
67	CAMAROGRAFO	CAM-068	2	Operador de cámara para producción audiovisual.	t	2025-12-05 15:55:51.09491-05	2025-12-05 15:55:51.09491-05	f	\N	\N	\N
68	CAMAROGRAFO INSTITUCIONAL	CAM-069	2	Camarógrafo especializado en eventos institucionales.	t	2025-12-05 15:55:51.095429-05	2025-12-05 15:55:51.095429-05	f	\N	\N	\N
69	CATALOGADOR SIGA	CAT-070	2	Encargado de catalogación en el sistema SIGA.	t	2025-12-05 15:55:51.095949-05	2025-12-05 15:55:51.095949-05	f	\N	\N	\N
70	CIRUJANO DENTISTA	CIR-071	2	Profesional especializado en odontología.	t	2025-12-05 15:55:51.096495-05	2025-12-05 15:55:51.096495-05	f	\N	\N	\N
71	COMUNICADOR	COM-072	2	Profesional en comunicación social.	t	2025-12-05 15:55:51.097077-05	2025-12-05 15:55:51.097077-05	f	\N	\N	\N
72	COMUNICADOR SOCIAL	COM-073	2	Especialista en comunicación e imagen institucional.	t	2025-12-05 15:55:51.098243-05	2025-12-05 15:55:51.098243-05	f	\N	\N	\N
73	COMUNICADORA	COM-074	2	Profesional en comunicación social.	t	2025-12-05 15:55:51.098784-05	2025-12-05 15:55:51.098784-05	f	\N	\N	\N
74	CONDUCTOR	CON-075	2	Conductor de vehículos institucionales.	t	2025-12-05 15:55:51.099314-05	2025-12-05 15:55:51.099314-05	f	\N	\N	\N
75	CONDUCTOR II	CON-076	2	Conductor de segundo nivel.	t	2025-12-05 15:55:51.099854-05	2025-12-05 15:55:51.099854-05	f	\N	\N	\N
76	CONDUCTOR III	CON-077	2	Conductor de tercer nivel.	t	2025-12-05 15:55:51.100592-05	2025-12-05 15:55:51.100592-05	f	\N	\N	\N
77	CONSEJERO REGIONAL	CON-078	1	Representante elegido del consejo regional.	t	2025-12-05 15:55:51.101101-05	2025-12-05 15:55:51.101101-05	f	\N	\N	\N
78	CONSEJERO REGIONAL CANGALLO	CON-079	1	Consejero regional representante de Cangallo.	t	2025-12-05 15:55:51.101752-05	2025-12-05 15:55:51.101752-05	f	\N	\N	\N
79	CONSEJERO REGIONAL FAJARDO	CON-080	1	Consejero regional representante de Fajardo.	t	2025-12-05 15:55:51.102265-05	2025-12-05 15:55:51.102265-05	f	\N	\N	\N
80	CONSEJERO REGIONAL HUAMANGA	CON-081	1	Consejero regional representante de Huamanga.	t	2025-12-05 15:55:51.102265-05	2025-12-05 15:55:51.102265-05	f	\N	\N	\N
81	CONSEJERO REGIONAL HUANCA SANCOS	CON-082	1	Consejero regional representante de Huanca Sancos.	t	2025-12-05 15:55:51.102777-05	2025-12-05 15:55:51.102777-05	f	\N	\N	\N
82	CONSEJERO REGIONAL HUANTA	CON-083	1	Consejero regional representante de Huanta.	t	2025-12-05 15:55:51.103321-05	2025-12-05 15:55:51.103321-05	f	\N	\N	\N
83	CONSEJERO REGIONAL LA MAR	CON-084	1	Consejero regional representante de La Mar.	t	2025-12-05 15:55:51.103321-05	2025-12-05 15:55:51.103321-05	f	\N	\N	\N
84	CONSEJERO REGIONAL PARINACOCHAS	CON-085	1	Consejero regional representante de Parinacochas.	t	2025-12-05 15:55:51.10388-05	2025-12-05 15:55:51.10388-05	f	\N	\N	\N
85	CONSEJERO REGIONAL VILCASHUAMAN	CON-086	1	Consejero regional representante de Vilcashuamán.	t	2025-12-05 15:55:51.104408-05	2025-12-05 15:55:51.104408-05	f	\N	\N	\N
86	CONTADOR	CON-087	2	Profesional contable encargado de registros financieros.	t	2025-12-05 15:55:51.105239-05	2025-12-05 15:55:51.105239-05	f	\N	\N	\N
87	CONTADOR II	CON-088	2	Contador de segundo nivel.	t	2025-12-05 15:55:51.105837-05	2025-12-05 15:55:51.105837-05	f	\N	\N	\N
88	CONTADOR III	CON-089	2	Contador de tercer nivel.	t	2025-12-05 15:55:51.106351-05	2025-12-05 15:55:51.106351-05	f	\N	\N	\N
89	CONTADOR IV	CON-090	2	Contador de cuarto nivel.	t	2025-12-05 15:55:51.106351-05	2025-12-05 15:55:51.106351-05	f	\N	\N	\N
90	COORDINADOR	COO-091	1	Responsable de coordinar actividades específicas.	t	2025-12-05 15:55:51.10688-05	2025-12-05 15:55:51.10688-05	f	\N	\N	\N
91	COORDINADOR ADMINISTRATIVO PROVIAS PCD-CG	COO-092	1	Coordinador administrativo especializado en Provias.	t	2025-12-05 15:55:51.107485-05	2025-12-05 15:55:51.107485-05	f	\N	\N	\N
92	COORDINADOR DE OBRA	COO-093	1	Coordinador de proyectos de obra y construcción.	t	2025-12-05 15:55:51.108039-05	2025-12-05 15:55:51.108039-05	f	\N	\N	\N
93	COORDINADOR DE PROCOMPITE	COO-094	1	Coordinador del programa Procompite.	t	2025-12-05 15:55:51.108039-05	2025-12-05 15:55:51.108039-05	f	\N	\N	\N
94	COORDINADOR DE PROYECTOS	COO-095	1	Coordinador de proyectos institucionales.	t	2025-12-05 15:55:51.108554-05	2025-12-05 15:55:51.108554-05	f	\N	\N	\N
95	DEFENSOR DEL ASEGURADO	DEF-096	2	Defensor de los derechos de los asegurados.	t	2025-12-05 15:55:51.109072-05	2025-12-05 15:55:51.109072-05	f	\N	\N	\N
96	DIGITADOR	DIG-097	2	Encargado de digitalización de documentos.	t	2025-12-05 15:55:51.109789-05	2025-12-05 15:55:51.109789-05	f	\N	\N	\N
97	DIRCTOR ASESORIA JURIDICA	DIR-098	1	Director del área de asesoría jurídica.	t	2025-12-05 15:55:51.110475-05	2025-12-05 15:55:51.110475-05	f	\N	\N	\N
98	DIRECTOR	DIR-099	1	Responsable de la dirección de unidades organizacionales.	t	2025-12-05 15:55:51.111521-05	2025-12-05 15:55:51.111521-05	f	\N	\N	\N
99	DIRECTOR DE ABASTECIMIENTO	DIR-100	1	Director del área de abastecimiento institucional.	t	2025-12-05 15:55:51.112042-05	2025-12-05 15:55:51.112042-05	f	\N	\N	\N
100	DIRECTOR DE ADMINISTRACION	DIR-101	1	Director del área de administración general.	t	2025-12-05 15:55:51.112748-05	2025-12-05 15:55:51.112748-05	f	\N	\N	\N
101	DIRECTOR DE ARCHIVO REGIONAL	DIR-102	1	Director del archivo regional institucional.	t	2025-12-05 15:55:51.113306-05	2025-12-05 15:55:51.113306-05	f	\N	\N	\N
102	DIRECTOR DE OFICINA	DIR-103	1	Director de oficina específica.	t	2025-12-05 15:55:51.114147-05	2025-12-05 15:55:51.114147-05	f	\N	\N	\N
103	DIRECTOR DE OREI	DIR-104	1	Director de la Oficina Regional de Evaluación Independiente.	t	2025-12-05 15:55:51.114862-05	2025-12-05 15:55:51.114862-05	f	\N	\N	\N
104	DIRECTOR DE RECURSOS HUMANOS	DIR-105	1	Director del área de recursos humanos.	t	2025-12-05 15:55:51.115897-05	2025-12-05 15:55:51.115897-05	f	\N	\N	\N
105	DIRECTOR DE SALUD AMBIENTAL	DIR-106	1	Director del área de salud ambiental.	t	2025-12-05 15:55:51.115897-05	2025-12-05 15:55:51.115897-05	f	\N	\N	\N
106	DIRECTOR DE SISTEMA ADMINISTRATIVO II	DIR-107	1	Director de sistemas administrativos de segundo nivel.	t	2025-12-05 15:55:51.116657-05	2025-12-05 15:55:51.116657-05	f	\N	\N	\N
107	DIRECTOR DE TESORERIA	DIR-108	1	Director del área de tesorería institucional.	t	2025-12-05 15:55:51.117465-05	2025-12-05 15:55:51.117465-05	f	\N	\N	\N
108	DIRECTOR DEL HOSPITAL REFERENCIAL CANGALLO	DIR-109	1	Director del hospital referencial de Cangallo.	t	2025-12-05 15:55:51.11802-05	2025-12-05 15:55:51.11802-05	f	\N	\N	\N
109	DIRECTOR EJECUTIVO	DIR-110	1	Director ejecutivo institucional.	t	2025-12-05 15:55:51.11802-05	2025-12-05 15:55:51.11802-05	f	\N	\N	\N
110	DIRECTOR GENERAL	DIR-111	1	Director general de la institución.	t	2025-12-05 15:55:51.118562-05	2025-12-05 15:55:51.118562-05	f	\N	\N	\N
111	DIRECTOR SUB REGIONAL	DIR-112	1	Director de nivel sub regional.	t	2025-12-05 15:55:51.119115-05	2025-12-05 15:55:51.119115-05	f	\N	\N	\N
112	DIRECTORA	DIR-113	1	Directora de unidades organizacionales.	t	2025-12-05 15:55:51.119637-05	2025-12-05 15:55:51.119637-05	f	\N	\N	\N
113	DIRECTORA EJECUTIVA	DIR-114	1	Directora ejecutiva institucional.	t	2025-12-05 15:55:51.119637-05	2025-12-05 15:55:51.119637-05	f	\N	\N	\N
114	ECONOMISTA	ECO-115	2	Profesional especializado en ciencias económicas.	t	2025-12-05 15:55:51.120182-05	2025-12-05 15:55:51.120182-05	f	\N	\N	\N
115	ECONOMISTA II	ECO-116	2	Economista de segundo nivel.	t	2025-12-05 15:55:51.120758-05	2025-12-05 15:55:51.120758-05	f	\N	\N	\N
116	ECONOMISTA III	ECO-117	2	Economista de tercer nivel.	t	2025-12-05 15:55:51.121402-05	2025-12-05 15:55:51.121402-05	f	\N	\N	\N
117	EDITOR AUDIOVISUAL	EDI-118	2	Especialista en edición de contenido audiovisual.	t	2025-12-05 15:55:51.121988-05	2025-12-05 15:55:51.121988-05	f	\N	\N	\N
118	ELABORADOR DE O/C	ELA-119	2	Encargado de elaborar órdenes de compra.	t	2025-12-05 15:55:51.122497-05	2025-12-05 15:55:51.122497-05	f	\N	\N	\N
119	ENCARGADO DE DESPACHO PERSONAL	ENC-120	2	Encargado del despacho de personal.	t	2025-12-05 15:55:51.122497-05	2025-12-05 15:55:51.122497-05	f	\N	\N	\N
120	ENCARGADO DE MESA DE PARTES	ENC-121	2	Encargado de la mesa de partes institucional.	t	2025-12-05 15:55:51.123026-05	2025-12-05 15:55:51.123026-05	f	\N	\N	\N
121	ENCARGADO DE PERU COMPRAS	ENC-122	2	Encargado de las operaciones en Perú Compras.	t	2025-12-05 15:55:51.123541-05	2025-12-05 15:55:51.123541-05	f	\N	\N	\N
122	ENFERMERA	ENF-123	2	Profesional de enfermería.	t	2025-12-05 15:55:51.12421-05	2025-12-05 15:55:51.12421-05	f	\N	\N	\N
123	ENFERMERA/O III	ENF-124	2	Enfermero/a de tercer nivel.	t	2025-12-05 15:55:51.12421-05	2025-12-05 15:55:51.12421-05	f	\N	\N	\N
124	ESPECIALISTA ADMINISTRATIVO	ESP-125	2	Especialista en procesos administrativos.	t	2025-12-05 15:55:51.124747-05	2025-12-05 15:55:51.124747-05	f	\N	\N	\N
125	ESPECIALISTA ADMINISTRATIVO I	ESP-126	2	Especialista administrativo de primer nivel.	t	2025-12-05 15:55:51.125271-05	2025-12-05 15:55:51.125271-05	f	\N	\N	\N
126	ESPECIALISTA ADMINISTRATIVO II	ESP-127	2	Especialista administrativo de segundo nivel.	t	2025-12-05 15:55:51.125785-05	2025-12-05 15:55:51.125785-05	f	\N	\N	\N
127	ESPECIALISTA AMBIENTAL	ESP-128	2	Especialista en temas ambientales y ecológicos.	t	2025-12-05 15:55:51.126306-05	2025-12-05 15:55:51.126306-05	f	\N	\N	\N
128	ESPECIALISTA ARQUEOLOGIA	ESP-129	2	Especialista en arqueología y patrimonio cultural.	t	2025-12-05 15:55:51.127395-05	2025-12-05 15:55:51.127395-05	f	\N	\N	\N
129	ESPECIALISTA COSTOS Y VALORIZACIONES	ESP-130	2	Especialista en costos y valorizaciones de obra.	t	2025-12-05 15:55:51.127905-05	2025-12-05 15:55:51.127905-05	f	\N	\N	\N
130	ESPECIALISTA DE COSTOS Y PRESUPUESTO	ESP-131	2	Especialista en elaboración de costos y presupuestos.	t	2025-12-05 15:55:51.12843-05	2025-12-05 15:55:51.12843-05	f	\N	\N	\N
131	ESPECIALISTA DE ESTRUCTURAS	ESP-132	2	Especialista en ingeniería estructural.	t	2025-12-05 15:55:51.129097-05	2025-12-05 15:55:51.129097-05	f	\N	\N	\N
132	ESPECIALISTA DE FAUNA	ESP-133	2	Especialista en fauna y biodiversidad.	t	2025-12-05 15:55:51.129733-05	2025-12-05 15:55:51.129733-05	f	\N	\N	\N
133	ESPECIALISTA DE GENERO	ESP-134	2	Especialista en temas de género e igualdad.	t	2025-12-05 15:55:51.130259-05	2025-12-05 15:55:51.130259-05	f	\N	\N	\N
134	ESPECIALISTA EIA	ESP-135	2	Especialista en Estudios de Impacto Ambiental.	t	2025-12-05 15:55:51.130813-05	2025-12-05 15:55:51.130813-05	f	\N	\N	\N
135	ESPECIALISTA EJECUCION	ESP-136	2	Especialista en ejecución de proyectos.	t	2025-12-05 15:55:51.131871-05	2025-12-05 15:55:51.131871-05	f	\N	\N	\N
136	ESPECIALISTA EN ACR	ESP-137	2	Especialista en Áreas de Conservación Regional.	t	2025-12-05 15:55:51.132379-05	2025-12-05 15:55:51.132379-05	f	\N	\N	\N
137	ESPECIALISTA EN ALIMENTOS Y ESTANDARES DE CALIDAD	ESP-138	2	Especialista en calidad alimentaria y estándares.	t	2025-12-05 15:55:51.132959-05	2025-12-05 15:55:51.132959-05	f	\N	\N	\N
138	ESPECIALISTA EN ARCHIVO I	ESP-139	2	Especialista de primer nivel en archivo documental.	t	2025-12-05 15:55:51.133483-05	2025-12-05 15:55:51.133483-05	f	\N	\N	\N
139	ESPECIALISTA EN ARCHIVO II	ESP-140	2	Especialista de segundo nivel en archivo documental.	t	2025-12-05 15:55:51.13402-05	2025-12-05 15:55:51.13402-05	f	\N	\N	\N
140	ESPECIALISTA EN ARCHIVO III	ESP-141	2	Especialista de tercer nivel en archivo documental.	t	2025-12-05 15:55:51.134534-05	2025-12-05 15:55:51.134534-05	f	\N	\N	\N
141	ESPECIALISTA EN BIODIVERSIDAD	ESP-143	2	Especialista en biodiversidad y conservación.	t	2025-12-05 15:55:51.135058-05	2025-12-05 15:55:51.135058-05	f	\N	\N	\N
142	ESPECIALISTA EN BIENESTAR Y SOCIAL I	ESP-144	2	Especialista de primer nivel en bienestar social.	t	2025-12-05 15:55:51.135625-05	2025-12-05 15:55:51.135625-05	f	\N	\N	\N
\.


--
-- TOC entry 5033 (class 0 OID 21954)
-- Dependencies: 226
-- Data for Name: ubigeos; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ubigeos (id, ubigeo_code, inei_code, department, province, district, created_at, updated_at) FROM stdin;
1	010101	010101	AMAZONAS	CHACHAPOYAS	CHACHAPOYAS	2025-12-05 15:55:51.193402-05	2025-12-05 15:55:51.193402-05
2	010102	010102	AMAZONAS	CHACHAPOYAS	ASUNCION	2025-12-05 15:55:51.194979-05	2025-12-05 15:55:51.194979-05
3	010103	010103	AMAZONAS	CHACHAPOYAS	BALSAS	2025-12-05 15:55:51.196046-05	2025-12-05 15:55:51.196046-05
4	010104	010104	AMAZONAS	CHACHAPOYAS	CHETO	2025-12-05 15:55:51.196562-05	2025-12-05 15:55:51.196562-05
5	010105	010105	AMAZONAS	CHACHAPOYAS	CHILIQUIN	2025-12-05 15:55:51.197692-05	2025-12-05 15:55:51.197692-05
6	010106	010106	AMAZONAS	CHACHAPOYAS	CHUQUIBAMBA	2025-12-05 15:55:51.198568-05	2025-12-05 15:55:51.198568-05
7	010107	010107	AMAZONAS	CHACHAPOYAS	GRANADA	2025-12-05 15:55:51.199097-05	2025-12-05 15:55:51.199097-05
8	010108	010108	AMAZONAS	CHACHAPOYAS	HUANCAS	2025-12-05 15:55:51.199617-05	2025-12-05 15:55:51.199617-05
9	010109	010109	AMAZONAS	CHACHAPOYAS	LA JALCA	2025-12-05 15:55:51.20013-05	2025-12-05 15:55:51.20013-05
10	010110	010110	AMAZONAS	CHACHAPOYAS	LEIMEBAMBA	2025-12-05 15:55:51.200648-05	2025-12-05 15:55:51.200648-05
11	010111	010111	AMAZONAS	CHACHAPOYAS	LEVANTO	2025-12-05 15:55:51.201201-05	2025-12-05 15:55:51.201201-05
12	010112	010112	AMAZONAS	CHACHAPOYAS	MAGDALENA	2025-12-05 15:55:51.201636-05	2025-12-05 15:55:51.201636-05
13	010113	010113	AMAZONAS	CHACHAPOYAS	MARISCAL CASTILLA	2025-12-05 15:55:51.202268-05	2025-12-05 15:55:51.202268-05
14	010114	010114	AMAZONAS	CHACHAPOYAS	MOLINOPAMPA	2025-12-05 15:55:51.202268-05	2025-12-05 15:55:51.202268-05
15	010115	010115	AMAZONAS	CHACHAPOYAS	MONTEVIDEO	2025-12-05 15:55:51.202848-05	2025-12-05 15:55:51.202848-05
16	010116	010116	AMAZONAS	CHACHAPOYAS	OLLEROS	2025-12-05 15:55:51.203476-05	2025-12-05 15:55:51.203476-05
17	010117	010117	AMAZONAS	CHACHAPOYAS	QUINJALCA	2025-12-05 15:55:51.203476-05	2025-12-05 15:55:51.203476-05
18	010118	010118	AMAZONAS	CHACHAPOYAS	SAN FRANCISCO DE DAGUAS	2025-12-05 15:55:51.204019-05	2025-12-05 15:55:51.204019-05
19	010119	010119	AMAZONAS	CHACHAPOYAS	SAN ISIDRO DE MAINO	2025-12-05 15:55:51.204527-05	2025-12-05 15:55:51.204527-05
20	010120	010120	AMAZONAS	CHACHAPOYAS	SOLOCO	2025-12-05 15:55:51.204527-05	2025-12-05 15:55:51.204527-05
21	010121	010121	AMAZONAS	CHACHAPOYAS	SONCHE	2025-12-05 15:55:51.205124-05	2025-12-05 15:55:51.205124-05
22	010205	010201	AMAZONAS	BAGUA	BAGUA	2025-12-05 15:55:51.205631-05	2025-12-05 15:55:51.205631-05
23	010202	010202	AMAZONAS	BAGUA	ARAMANGO	2025-12-05 15:55:51.206175-05	2025-12-05 15:55:51.206175-05
24	010203	010203	AMAZONAS	BAGUA	COPALLIN	2025-12-05 15:55:51.206541-05	2025-12-05 15:55:51.206541-05
25	010204	010204	AMAZONAS	BAGUA	EL PARCO	2025-12-05 15:55:51.207077-05	2025-12-05 15:55:51.207077-05
26	010206	010205	AMAZONAS	BAGUA	IMAZA	2025-12-05 15:55:51.207643-05	2025-12-05 15:55:51.207643-05
27	010201	010206	AMAZONAS	BAGUA	LA PECA	2025-12-05 15:55:51.208194-05	2025-12-05 15:55:51.208194-05
28	010301	010301	AMAZONAS	BONGARA	JUMBILLA	2025-12-05 15:55:51.208194-05	2025-12-05 15:55:51.208194-05
29	010304	010302	AMAZONAS	BONGARA	CHISQUILLA	2025-12-05 15:55:51.208709-05	2025-12-05 15:55:51.208709-05
30	010305	010303	AMAZONAS	BONGARA	CHURUJA	2025-12-05 15:55:51.209264-05	2025-12-05 15:55:51.209264-05
31	010302	010304	AMAZONAS	BONGARA	COROSHA	2025-12-05 15:55:51.209886-05	2025-12-05 15:55:51.209886-05
32	010303	010305	AMAZONAS	BONGARA	CUISPES	2025-12-05 15:55:51.210578-05	2025-12-05 15:55:51.210578-05
33	010306	010306	AMAZONAS	BONGARA	FLORIDA	2025-12-05 15:55:51.211245-05	2025-12-05 15:55:51.211245-05
34	010312	010307	AMAZONAS	BONGARA	JAZAN	2025-12-05 15:55:51.211945-05	2025-12-05 15:55:51.211945-05
35	010307	010308	AMAZONAS	BONGARA	RECTA	2025-12-05 15:55:51.212594-05	2025-12-05 15:55:51.212594-05
36	010308	010309	AMAZONAS	BONGARA	SAN CARLOS	2025-12-05 15:55:51.213215-05	2025-12-05 15:55:51.213215-05
37	010309	010310	AMAZONAS	BONGARA	SHIPASBAMBA	2025-12-05 15:55:51.213215-05	2025-12-05 15:55:51.213215-05
38	010310	010311	AMAZONAS	BONGARA	VALERA	2025-12-05 15:55:51.214489-05	2025-12-05 15:55:51.214489-05
39	010311	010312	AMAZONAS	BONGARA	YAMBRASBAMBA	2025-12-05 15:55:51.215281-05	2025-12-05 15:55:51.215281-05
40	010601	010401	AMAZONAS	CONDORCANQUI	NIEVA	2025-12-05 15:55:51.215802-05	2025-12-05 15:55:51.215802-05
41	010603	010402	AMAZONAS	CONDORCANQUI	EL CENEPA	2025-12-05 15:55:51.216258-05	2025-12-05 15:55:51.216258-05
42	010602	010403	AMAZONAS	CONDORCANQUI	RIO SANTIAGO	2025-12-05 15:55:51.217162-05	2025-12-05 15:55:51.217162-05
43	010401	010501	AMAZONAS	LUYA	LAMUD	2025-12-05 15:55:51.217162-05	2025-12-05 15:55:51.217162-05
44	010402	010502	AMAZONAS	LUYA	CAMPORREDONDO	2025-12-05 15:55:51.217685-05	2025-12-05 15:55:51.217685-05
45	010403	010503	AMAZONAS	LUYA	COCABAMBA	2025-12-05 15:55:51.218197-05	2025-12-05 15:55:51.218197-05
46	010404	010504	AMAZONAS	LUYA	COLCAMAR	2025-12-05 15:55:51.218816-05	2025-12-05 15:55:51.218816-05
47	010405	010505	AMAZONAS	LUYA	CONILA	2025-12-05 15:55:51.219338-05	2025-12-05 15:55:51.219338-05
48	010406	010506	AMAZONAS	LUYA	INGUILPATA	2025-12-05 15:55:51.219338-05	2025-12-05 15:55:51.219338-05
49	010407	010507	AMAZONAS	LUYA	LONGUITA	2025-12-05 15:55:51.219853-05	2025-12-05 15:55:51.219853-05
50	010408	010508	AMAZONAS	LUYA	LONYA CHICO	2025-12-05 15:55:51.220386-05	2025-12-05 15:55:51.220386-05
51	010409	010509	AMAZONAS	LUYA	LUYA	2025-12-05 15:55:51.220903-05	2025-12-05 15:55:51.220903-05
52	010410	010510	AMAZONAS	LUYA	LUYA VIEJO	2025-12-05 15:55:51.221677-05	2025-12-05 15:55:51.221677-05
53	010411	010511	AMAZONAS	LUYA	MARIA	2025-12-05 15:55:51.221677-05	2025-12-05 15:55:51.221677-05
54	010412	010512	AMAZONAS	LUYA	OCALLI	2025-12-05 15:55:51.222199-05	2025-12-05 15:55:51.222199-05
55	010413	010513	AMAZONAS	LUYA	OCUMAL	2025-12-05 15:55:51.222705-05	2025-12-05 15:55:51.222705-05
56	010414	010514	AMAZONAS	LUYA	PISUQUIA	2025-12-05 15:55:51.225124-05	2025-12-05 15:55:51.225124-05
57	010423	010515	AMAZONAS	LUYA	PROVIDENCIA	2025-12-05 15:55:51.225627-05	2025-12-05 15:55:51.225627-05
58	010415	010516	AMAZONAS	LUYA	SAN CRISTOBAL	2025-12-05 15:55:51.226149-05	2025-12-05 15:55:51.226149-05
59	010416	010517	AMAZONAS	LUYA	SAN FRANCISCO DEL YESO	2025-12-05 15:55:51.227181-05	2025-12-05 15:55:51.227181-05
60	010417	010518	AMAZONAS	LUYA	SAN JERONIMO	2025-12-05 15:55:51.227745-05	2025-12-05 15:55:51.227745-05
61	010418	010519	AMAZONAS	LUYA	SAN JUAN DE LOPECANCHA	2025-12-05 15:55:51.2283-05	2025-12-05 15:55:51.2283-05
62	010419	010520	AMAZONAS	LUYA	SANTA CATALINA	2025-12-05 15:55:51.229053-05	2025-12-05 15:55:51.229053-05
63	010420	010521	AMAZONAS	LUYA	SANTO TOMAS	2025-12-05 15:55:51.229675-05	2025-12-05 15:55:51.229675-05
64	010421	010522	AMAZONAS	LUYA	TINGO	2025-12-05 15:55:51.230218-05	2025-12-05 15:55:51.230218-05
65	010422	010523	AMAZONAS	LUYA	TRITA	2025-12-05 15:55:51.230738-05	2025-12-05 15:55:51.230738-05
66	010501	010601	AMAZONAS	RODRIGUEZ DE MENDOZA	SAN NICOLAS	2025-12-05 15:55:51.231656-05	2025-12-05 15:55:51.231656-05
67	010503	010602	AMAZONAS	RODRIGUEZ DE MENDOZA	CHIRIMOTO	2025-12-05 15:55:51.232538-05	2025-12-05 15:55:51.232538-05
68	010502	010603	AMAZONAS	RODRIGUEZ DE MENDOZA	COCHAMAL	2025-12-05 15:55:51.233117-05	2025-12-05 15:55:51.233117-05
69	010504	010604	AMAZONAS	RODRIGUEZ DE MENDOZA	HUAMBO	2025-12-05 15:55:51.233641-05	2025-12-05 15:55:51.233641-05
70	010505	010605	AMAZONAS	RODRIGUEZ DE MENDOZA	LIMABAMBA	2025-12-05 15:55:51.234156-05	2025-12-05 15:55:51.234156-05
71	010506	010606	AMAZONAS	RODRIGUEZ DE MENDOZA	LONGAR	2025-12-05 15:55:51.234156-05	2025-12-05 15:55:51.234156-05
72	010508	010607	AMAZONAS	RODRIGUEZ DE MENDOZA	MARISCAL BENAVIDES	2025-12-05 15:55:51.234664-05	2025-12-05 15:55:51.234664-05
73	010507	010608	AMAZONAS	RODRIGUEZ DE MENDOZA	MILPUC	2025-12-05 15:55:51.23526-05	2025-12-05 15:55:51.23526-05
74	010509	010609	AMAZONAS	RODRIGUEZ DE MENDOZA	OMIA	2025-12-05 15:55:51.23526-05	2025-12-05 15:55:51.23526-05
75	010510	010610	AMAZONAS	RODRIGUEZ DE MENDOZA	SANTA ROSA	2025-12-05 15:55:51.235768-05	2025-12-05 15:55:51.235768-05
76	010511	010611	AMAZONAS	RODRIGUEZ DE MENDOZA	TOTORA	2025-12-05 15:55:51.236396-05	2025-12-05 15:55:51.236396-05
77	010512	010612	AMAZONAS	RODRIGUEZ DE MENDOZA	VISTA ALEGRE	2025-12-05 15:55:51.237027-05	2025-12-05 15:55:51.237027-05
78	010701	010701	AMAZONAS	UTCUBAMBA	BAGUA GRANDE	2025-12-05 15:55:51.237544-05	2025-12-05 15:55:51.237544-05
79	010702	010702	AMAZONAS	UTCUBAMBA	CAJARURO	2025-12-05 15:55:51.237544-05	2025-12-05 15:55:51.237544-05
80	010703	010703	AMAZONAS	UTCUBAMBA	CUMBA	2025-12-05 15:55:51.238082-05	2025-12-05 15:55:51.238082-05
81	010704	010704	AMAZONAS	UTCUBAMBA	EL MILAGRO	2025-12-05 15:55:51.238596-05	2025-12-05 15:55:51.238596-05
82	010705	010705	AMAZONAS	UTCUBAMBA	JAMALCA	2025-12-05 15:55:51.239108-05	2025-12-05 15:55:51.239108-05
83	010706	010706	AMAZONAS	UTCUBAMBA	LONYA GRANDE	2025-12-05 15:55:51.23962-05	2025-12-05 15:55:51.23962-05
84	010707	010707	AMAZONAS	UTCUBAMBA	YAMON	2025-12-05 15:55:51.240134-05	2025-12-05 15:55:51.240134-05
85	020101	020101	ANCASH	HUARAZ	HUARAZ	2025-12-05 15:55:51.240793-05	2025-12-05 15:55:51.240793-05
86	020103	020102	ANCASH	HUARAZ	COCHABAMBA	2025-12-05 15:55:51.240793-05	2025-12-05 15:55:51.240793-05
87	020104	020103	ANCASH	HUARAZ	COLCABAMBA	2025-12-05 15:55:51.241305-05	2025-12-05 15:55:51.241305-05
88	020105	020104	ANCASH	HUARAZ	HUANCHAY	2025-12-05 15:55:51.242231-05	2025-12-05 15:55:51.242231-05
89	020102	020105	ANCASH	HUARAZ	INDEPENDENCIA	2025-12-05 15:55:51.242764-05	2025-12-05 15:55:51.242764-05
90	020106	020106	ANCASH	HUARAZ	JANGAS	2025-12-05 15:55:51.243284-05	2025-12-05 15:55:51.243284-05
91	020107	020107	ANCASH	HUARAZ	LA LIBERTAD	2025-12-05 15:55:51.243978-05	2025-12-05 15:55:51.243978-05
92	020108	020108	ANCASH	HUARAZ	OLLEROS	2025-12-05 15:55:51.244492-05	2025-12-05 15:55:51.244492-05
93	020109	020109	ANCASH	HUARAZ	PAMPAS	2025-12-05 15:55:51.245562-05	2025-12-05 15:55:51.245562-05
94	020110	020110	ANCASH	HUARAZ	PARIACOTO	2025-12-05 15:55:51.245562-05	2025-12-05 15:55:51.245562-05
95	020111	020111	ANCASH	HUARAZ	PIRA	2025-12-05 15:55:51.246082-05	2025-12-05 15:55:51.246082-05
96	020112	020112	ANCASH	HUARAZ	TARICA	2025-12-05 15:55:51.246599-05	2025-12-05 15:55:51.246599-05
97	020201	020201	ANCASH	AIJA	AIJA	2025-12-05 15:55:51.247532-05	2025-12-05 15:55:51.247532-05
98	020203	020202	ANCASH	AIJA	CORIS	2025-12-05 15:55:51.248562-05	2025-12-05 15:55:51.248562-05
99	020205	020203	ANCASH	AIJA	HUACLLAN	2025-12-05 15:55:51.249078-05	2025-12-05 15:55:51.249078-05
100	020206	020204	ANCASH	AIJA	LA MERCED	2025-12-05 15:55:51.249597-05	2025-12-05 15:55:51.249597-05
101	020208	020205	ANCASH	AIJA	SUCCHA	2025-12-05 15:55:51.250125-05	2025-12-05 15:55:51.250125-05
102	021601	020301	ANCASH	ANTONIO RAYMONDI	LLAMELLIN	2025-12-05 15:55:51.250645-05	2025-12-05 15:55:51.250645-05
103	021602	020302	ANCASH	ANTONIO RAYMONDI	ACZO	2025-12-05 15:55:51.250645-05	2025-12-05 15:55:51.250645-05
104	021603	020303	ANCASH	ANTONIO RAYMONDI	CHACCHO	2025-12-05 15:55:51.251206-05	2025-12-05 15:55:51.251206-05
105	021604	020304	ANCASH	ANTONIO RAYMONDI	CHINGAS	2025-12-05 15:55:51.251751-05	2025-12-05 15:55:51.251751-05
106	021605	020305	ANCASH	ANTONIO RAYMONDI	MIRGAS	2025-12-05 15:55:51.251751-05	2025-12-05 15:55:51.251751-05
107	021606	020306	ANCASH	ANTONIO RAYMONDI	SAN JUAN DE RONTOY	2025-12-05 15:55:51.252292-05	2025-12-05 15:55:51.252292-05
108	021801	020401	ANCASH	ASUNCION	CHACAS	2025-12-05 15:55:51.252816-05	2025-12-05 15:55:51.252816-05
109	021802	020402	ANCASH	ASUNCION	ACOCHACA	2025-12-05 15:55:51.252816-05	2025-12-05 15:55:51.252816-05
110	020301	020501	ANCASH	BOLOGNESI	CHIQUIAN	2025-12-05 15:55:51.253597-05	2025-12-05 15:55:51.253597-05
111	020302	020502	ANCASH	BOLOGNESI	ABELARDO PARDO LEZAMETA	2025-12-05 15:55:51.254385-05	2025-12-05 15:55:51.254385-05
112	020321	020503	ANCASH	BOLOGNESI	ANTONIO RAYMONDI	2025-12-05 15:55:51.254385-05	2025-12-05 15:55:51.254385-05
113	020304	020504	ANCASH	BOLOGNESI	AQUIA	2025-12-05 15:55:51.254899-05	2025-12-05 15:55:51.254899-05
114	020305	020505	ANCASH	BOLOGNESI	CAJACAY	2025-12-05 15:55:51.255425-05	2025-12-05 15:55:51.255425-05
115	020322	020506	ANCASH	BOLOGNESI	CANIS	2025-12-05 15:55:51.255425-05	2025-12-05 15:55:51.255425-05
116	020323	020507	ANCASH	BOLOGNESI	COLQUIOC	2025-12-05 15:55:51.255932-05	2025-12-05 15:55:51.255932-05
117	020325	020508	ANCASH	BOLOGNESI	HUALLANCA	2025-12-05 15:55:51.256478-05	2025-12-05 15:55:51.256478-05
118	020311	020509	ANCASH	BOLOGNESI	HUASTA	2025-12-05 15:55:51.256478-05	2025-12-05 15:55:51.256478-05
119	020310	020510	ANCASH	BOLOGNESI	HUAYLLACAYAN	2025-12-05 15:55:51.257029-05	2025-12-05 15:55:51.257029-05
120	020324	020511	ANCASH	BOLOGNESI	LA PRIMAVERA	2025-12-05 15:55:51.257542-05	2025-12-05 15:55:51.257542-05
121	020313	020512	ANCASH	BOLOGNESI	MANGAS	2025-12-05 15:55:51.258197-05	2025-12-05 15:55:51.258197-05
122	020315	020513	ANCASH	BOLOGNESI	PACLLON	2025-12-05 15:55:51.258197-05	2025-12-05 15:55:51.258197-05
123	020317	020514	ANCASH	BOLOGNESI	SAN MIGUEL DE CORPANQUI	2025-12-05 15:55:51.259287-05	2025-12-05 15:55:51.259287-05
124	020320	020515	ANCASH	BOLOGNESI	TICLLOS	2025-12-05 15:55:51.259798-05	2025-12-05 15:55:51.259798-05
125	020401	020601	ANCASH	CARHUAZ	CARHUAZ	2025-12-05 15:55:51.260465-05	2025-12-05 15:55:51.260465-05
126	020402	020602	ANCASH	CARHUAZ	ACOPAMPA	2025-12-05 15:55:51.261521-05	2025-12-05 15:55:51.261521-05
127	020403	020603	ANCASH	CARHUAZ	AMASHCA	2025-12-05 15:55:51.262057-05	2025-12-05 15:55:51.262057-05
128	020404	020604	ANCASH	CARHUAZ	ANTA	2025-12-05 15:55:51.262575-05	2025-12-05 15:55:51.262575-05
129	020405	020605	ANCASH	CARHUAZ	ATAQUERO	2025-12-05 15:55:51.263096-05	2025-12-05 15:55:51.263096-05
130	020406	020606	ANCASH	CARHUAZ	MARCARA	2025-12-05 15:55:51.263617-05	2025-12-05 15:55:51.263617-05
131	020407	020607	ANCASH	CARHUAZ	PARIAHUANCA	2025-12-05 15:55:51.264135-05	2025-12-05 15:55:51.264135-05
132	020408	020608	ANCASH	CARHUAZ	SAN MIGUEL DE ACO	2025-12-05 15:55:51.264897-05	2025-12-05 15:55:51.264897-05
133	020409	020609	ANCASH	CARHUAZ	SHILLA	2025-12-05 15:55:51.265466-05	2025-12-05 15:55:51.265466-05
134	020410	020610	ANCASH	CARHUAZ	TINCO	2025-12-05 15:55:51.2665-05	2025-12-05 15:55:51.2665-05
135	020411	020611	ANCASH	CARHUAZ	YUNGAR	2025-12-05 15:55:51.2665-05	2025-12-05 15:55:51.2665-05
136	021701	020701	ANCASH	CARLOS FERMIN FITZCARRALD	SAN LUIS	2025-12-05 15:55:51.267018-05	2025-12-05 15:55:51.267018-05
137	021703	020702	ANCASH	CARLOS FERMIN FITZCARRALD	SAN NICOLAS	2025-12-05 15:55:51.267609-05	2025-12-05 15:55:51.267609-05
138	021702	020703	ANCASH	CARLOS FERMIN FITZCARRALD	YAUYA	2025-12-05 15:55:51.268078-05	2025-12-05 15:55:51.268078-05
139	020501	020801	ANCASH	CASMA	CASMA	2025-12-05 15:55:51.268509-05	2025-12-05 15:55:51.268509-05
140	020502	020802	ANCASH	CASMA	BUENA VISTA ALTA	2025-12-05 15:55:51.269134-05	2025-12-05 15:55:51.269134-05
141	020503	020803	ANCASH	CASMA	COMANDANTE NOEL	2025-12-05 15:55:51.269134-05	2025-12-05 15:55:51.269134-05
142	020505	020804	ANCASH	CASMA	YAUTAN	2025-12-05 15:55:51.269652-05	2025-12-05 15:55:51.269652-05
143	020601	020901	ANCASH	CORONGO	CORONGO	2025-12-05 15:55:51.270318-05	2025-12-05 15:55:51.270318-05
144	020602	020902	ANCASH	CORONGO	ACO	2025-12-05 15:55:51.270318-05	2025-12-05 15:55:51.270318-05
145	020603	020903	ANCASH	CORONGO	BAMBAS	2025-12-05 15:55:51.270847-05	2025-12-05 15:55:51.270847-05
146	020604	020904	ANCASH	CORONGO	CUSCA	2025-12-05 15:55:51.271419-05	2025-12-05 15:55:51.271419-05
147	020605	020905	ANCASH	CORONGO	LA PAMPA	2025-12-05 15:55:51.271419-05	2025-12-05 15:55:51.271419-05
148	020606	020906	ANCASH	CORONGO	YANAC	2025-12-05 15:55:51.271947-05	2025-12-05 15:55:51.271947-05
149	020607	020907	ANCASH	CORONGO	YUPAN	2025-12-05 15:55:51.271947-05	2025-12-05 15:55:51.271947-05
150	020801	021001	ANCASH	HUARI	HUARI	2025-12-05 15:55:51.272514-05	2025-12-05 15:55:51.272514-05
151	020816	021002	ANCASH	HUARI	ANRA	2025-12-05 15:55:51.27317-05	2025-12-05 15:55:51.27317-05
152	020802	021003	ANCASH	HUARI	CAJAY	2025-12-05 15:55:51.273668-05	2025-12-05 15:55:51.273668-05
153	020803	021004	ANCASH	HUARI	CHAVIN DE HUANTAR	2025-12-05 15:55:51.274228-05	2025-12-05 15:55:51.274228-05
154	020804	021005	ANCASH	HUARI	HUACACHI	2025-12-05 15:55:51.274228-05	2025-12-05 15:55:51.274228-05
155	020806	021006	ANCASH	HUARI	HUACCHIS	2025-12-05 15:55:51.27475-05	2025-12-05 15:55:51.27475-05
156	020805	021007	ANCASH	HUARI	HUACHIS	2025-12-05 15:55:51.275271-05	2025-12-05 15:55:51.275271-05
157	020807	021008	ANCASH	HUARI	HUANTAR	2025-12-05 15:55:51.275818-05	2025-12-05 15:55:51.275818-05
158	020808	021009	ANCASH	HUARI	MASIN	2025-12-05 15:55:51.276338-05	2025-12-05 15:55:51.276338-05
159	020809	021010	ANCASH	HUARI	PAUCAS	2025-12-05 15:55:51.276954-05	2025-12-05 15:55:51.276954-05
160	020810	021011	ANCASH	HUARI	PONTO	2025-12-05 15:55:51.277627-05	2025-12-05 15:55:51.277627-05
161	020811	021012	ANCASH	HUARI	RAHUAPAMPA	2025-12-05 15:55:51.278215-05	2025-12-05 15:55:51.278215-05
162	020812	021013	ANCASH	HUARI	RAPAYAN	2025-12-05 15:55:51.278859-05	2025-12-05 15:55:51.278859-05
163	020813	021014	ANCASH	HUARI	SAN MARCOS	2025-12-05 15:55:51.279377-05	2025-12-05 15:55:51.279377-05
164	020814	021015	ANCASH	HUARI	SAN PEDRO DE CHANA	2025-12-05 15:55:51.280062-05	2025-12-05 15:55:51.280062-05
165	020815	021016	ANCASH	HUARI	UCO	2025-12-05 15:55:51.280614-05	2025-12-05 15:55:51.280614-05
166	021901	021101	ANCASH	HUARMEY	HUARMEY	2025-12-05 15:55:51.281612-05	2025-12-05 15:55:51.281612-05
167	021902	021102	ANCASH	HUARMEY	COCHAPETI	2025-12-05 15:55:51.282419-05	2025-12-05 15:55:51.282419-05
168	021905	021103	ANCASH	HUARMEY	CULEBRAS	2025-12-05 15:55:51.282886-05	2025-12-05 15:55:51.282886-05
169	021903	021104	ANCASH	HUARMEY	HUAYAN	2025-12-05 15:55:51.283358-05	2025-12-05 15:55:51.283358-05
170	021904	021105	ANCASH	HUARMEY	MALVAS	2025-12-05 15:55:51.28388-05	2025-12-05 15:55:51.28388-05
171	020701	021201	ANCASH	HUAYLAS	CARAZ	2025-12-05 15:55:51.284404-05	2025-12-05 15:55:51.284404-05
172	020702	021202	ANCASH	HUAYLAS	HUALLANCA	2025-12-05 15:55:51.284925-05	2025-12-05 15:55:51.284925-05
173	020703	021203	ANCASH	HUAYLAS	HUATA	2025-12-05 15:55:51.284925-05	2025-12-05 15:55:51.284925-05
174	020704	021204	ANCASH	HUAYLAS	HUAYLAS	2025-12-05 15:55:51.285998-05	2025-12-05 15:55:51.285998-05
175	020705	021205	ANCASH	HUAYLAS	MATO	2025-12-05 15:55:51.285998-05	2025-12-05 15:55:51.285998-05
176	020706	021206	ANCASH	HUAYLAS	PAMPAROMAS	2025-12-05 15:55:51.286524-05	2025-12-05 15:55:51.286524-05
177	020707	021207	ANCASH	HUAYLAS	PUEBLO LIBRE	2025-12-05 15:55:51.287096-05	2025-12-05 15:55:51.287096-05
178	020708	021208	ANCASH	HUAYLAS	SANTA CRUZ	2025-12-05 15:55:51.287644-05	2025-12-05 15:55:51.287644-05
179	020710	021209	ANCASH	HUAYLAS	SANTO TORIBIO	2025-12-05 15:55:51.288262-05	2025-12-05 15:55:51.288262-05
180	020709	021210	ANCASH	HUAYLAS	YURACMARCA	2025-12-05 15:55:51.288815-05	2025-12-05 15:55:51.288815-05
181	020901	021301	ANCASH	MARISCAL LUZURIAGA	PISCOBAMBA	2025-12-05 15:55:51.289373-05	2025-12-05 15:55:51.289373-05
182	020902	021302	ANCASH	MARISCAL LUZURIAGA	CASCA	2025-12-05 15:55:51.289373-05	2025-12-05 15:55:51.289373-05
183	020908	021303	ANCASH	MARISCAL LUZURIAGA	ELEAZAR GUZMAN BARRON	2025-12-05 15:55:51.289904-05	2025-12-05 15:55:51.289904-05
184	020904	021304	ANCASH	MARISCAL LUZURIAGA	FIDEL OLIVAS ESCUDERO	2025-12-05 15:55:51.29048-05	2025-12-05 15:55:51.29048-05
185	020905	021305	ANCASH	MARISCAL LUZURIAGA	LLAMA	2025-12-05 15:55:51.291001-05	2025-12-05 15:55:51.291001-05
186	020906	021306	ANCASH	MARISCAL LUZURIAGA	LLUMPA	2025-12-05 15:55:51.291001-05	2025-12-05 15:55:51.291001-05
187	020903	021307	ANCASH	MARISCAL LUZURIAGA	LUCMA	2025-12-05 15:55:51.291533-05	2025-12-05 15:55:51.291533-05
188	020907	021308	ANCASH	MARISCAL LUZURIAGA	MUSGA	2025-12-05 15:55:51.291533-05	2025-12-05 15:55:51.291533-05
189	022007	021401	ANCASH	OCROS	OCROS	2025-12-05 15:55:51.292631-05	2025-12-05 15:55:51.292631-05
190	022001	021402	ANCASH	OCROS	ACAS	2025-12-05 15:55:51.29315-05	2025-12-05 15:55:51.29315-05
191	022002	021403	ANCASH	OCROS	CAJAMARQUILLA	2025-12-05 15:55:51.293815-05	2025-12-05 15:55:51.293815-05
192	022003	021404	ANCASH	OCROS	CARHUAPAMPA	2025-12-05 15:55:51.294959-05	2025-12-05 15:55:51.294959-05
193	022004	021405	ANCASH	OCROS	COCHAS	2025-12-05 15:55:51.295485-05	2025-12-05 15:55:51.295485-05
194	022005	021406	ANCASH	OCROS	CONGAS	2025-12-05 15:55:51.295485-05	2025-12-05 15:55:51.295485-05
195	022006	021407	ANCASH	OCROS	LLIPA	2025-12-05 15:55:51.296132-05	2025-12-05 15:55:51.296132-05
196	022008	021408	ANCASH	OCROS	SAN CRISTOBAL DE RAJAN	2025-12-05 15:55:51.29667-05	2025-12-05 15:55:51.29667-05
197	022009	021409	ANCASH	OCROS	SAN PEDRO	2025-12-05 15:55:51.297607-05	2025-12-05 15:55:51.297607-05
198	022010	021410	ANCASH	OCROS	SANTIAGO DE CHILCAS	2025-12-05 15:55:51.298258-05	2025-12-05 15:55:51.298258-05
199	021001	021501	ANCASH	PALLASCA	CABANA	2025-12-05 15:55:51.29931-05	2025-12-05 15:55:51.29931-05
200	021002	021502	ANCASH	PALLASCA	BOLOGNESI	2025-12-05 15:55:51.299828-05	2025-12-05 15:55:51.299828-05
201	021003	021503	ANCASH	PALLASCA	CONCHUCOS	2025-12-05 15:55:51.300349-05	2025-12-05 15:55:51.300349-05
202	021004	021504	ANCASH	PALLASCA	HUACASCHUQUE	2025-12-05 15:55:51.300349-05	2025-12-05 15:55:51.300349-05
203	021005	021505	ANCASH	PALLASCA	HUANDOVAL	2025-12-05 15:55:51.300876-05	2025-12-05 15:55:51.300876-05
204	021006	021506	ANCASH	PALLASCA	LACABAMBA	2025-12-05 15:55:51.301403-05	2025-12-05 15:55:51.301403-05
205	021007	021507	ANCASH	PALLASCA	LLAPO	2025-12-05 15:55:51.30194-05	2025-12-05 15:55:51.30194-05
206	021008	021508	ANCASH	PALLASCA	PALLASCA	2025-12-05 15:55:51.30194-05	2025-12-05 15:55:51.30194-05
207	021009	021509	ANCASH	PALLASCA	PAMPAS	2025-12-05 15:55:51.30246-05	2025-12-05 15:55:51.30246-05
208	021010	021510	ANCASH	PALLASCA	SANTA ROSA	2025-12-05 15:55:51.302979-05	2025-12-05 15:55:51.302979-05
209	021011	021511	ANCASH	PALLASCA	TAUCA	2025-12-05 15:55:51.303597-05	2025-12-05 15:55:51.303597-05
210	021101	021601	ANCASH	POMABAMBA	POMABAMBA	2025-12-05 15:55:51.304131-05	2025-12-05 15:55:51.304131-05
211	021102	021602	ANCASH	POMABAMBA	HUAYLLAN	2025-12-05 15:55:51.304508-05	2025-12-05 15:55:51.304508-05
212	021103	021603	ANCASH	POMABAMBA	PAROBAMBA	2025-12-05 15:55:51.305035-05	2025-12-05 15:55:51.305035-05
213	021104	021604	ANCASH	POMABAMBA	QUINUABAMBA	2025-12-05 15:55:51.305559-05	2025-12-05 15:55:51.305559-05
214	021201	021701	ANCASH	RECUAY	RECUAY	2025-12-05 15:55:51.305559-05	2025-12-05 15:55:51.305559-05
215	021210	021702	ANCASH	RECUAY	CATAC	2025-12-05 15:55:51.306089-05	2025-12-05 15:55:51.306089-05
216	110334	120434	JUNIN	JAUJA	YAUYOS	2025-12-05 15:55:51.306641-05	2025-12-05 15:55:51.306641-05
217	021202	021703	ANCASH	RECUAY	COTAPARACO	2025-12-05 15:55:51.306641-05	2025-12-05 15:55:51.306641-05
218	021203	021704	ANCASH	RECUAY	HUAYLLAPAMPA	2025-12-05 15:55:51.307195-05	2025-12-05 15:55:51.307195-05
219	021209	021705	ANCASH	RECUAY	LLACLLIN	2025-12-05 15:55:51.307792-05	2025-12-05 15:55:51.307792-05
220	021204	021706	ANCASH	RECUAY	MARCA	2025-12-05 15:55:51.307792-05	2025-12-05 15:55:51.307792-05
221	021205	021707	ANCASH	RECUAY	PAMPAS CHICO	2025-12-05 15:55:51.308348-05	2025-12-05 15:55:51.308348-05
222	021206	021708	ANCASH	RECUAY	PARARIN	2025-12-05 15:55:51.30889-05	2025-12-05 15:55:51.30889-05
223	021207	021709	ANCASH	RECUAY	TAPACOCHA	2025-12-05 15:55:51.309522-05	2025-12-05 15:55:51.309522-05
224	021208	021710	ANCASH	RECUAY	TICAPAMPA	2025-12-05 15:55:51.310119-05	2025-12-05 15:55:51.310119-05
225	021301	021801	ANCASH	SANTA	CHIMBOTE	2025-12-05 15:55:51.310821-05	2025-12-05 15:55:51.310821-05
226	021302	021802	ANCASH	SANTA	CACERES DEL PERU	2025-12-05 15:55:51.311364-05	2025-12-05 15:55:51.311364-05
227	021308	021803	ANCASH	SANTA	COISHCO	2025-12-05 15:55:51.311981-05	2025-12-05 15:55:51.311981-05
228	021303	021804	ANCASH	SANTA	MACATE	2025-12-05 15:55:51.312494-05	2025-12-05 15:55:51.312494-05
229	021304	021805	ANCASH	SANTA	MORO	2025-12-05 15:55:51.313006-05	2025-12-05 15:55:51.313006-05
230	021305	021806	ANCASH	SANTA	NEPEÑA	2025-12-05 15:55:51.313006-05	2025-12-05 15:55:51.313006-05
231	021306	021807	ANCASH	SANTA	SAMANCO	2025-12-05 15:55:51.313527-05	2025-12-05 15:55:51.313527-05
232	021307	021808	ANCASH	SANTA	SANTA	2025-12-05 15:55:51.314389-05	2025-12-05 15:55:51.314389-05
233	021309	021809	ANCASH	SANTA	NUEVO CHIMBOTE	2025-12-05 15:55:51.315565-05	2025-12-05 15:55:51.315565-05
234	021401	021901	ANCASH	SIHUAS	SIHUAS	2025-12-05 15:55:51.316731-05	2025-12-05 15:55:51.316731-05
235	021407	021902	ANCASH	SIHUAS	ACOBAMBA	2025-12-05 15:55:51.3173-05	2025-12-05 15:55:51.3173-05
236	021402	021903	ANCASH	SIHUAS	ALFONSO UGARTE	2025-12-05 15:55:51.317991-05	2025-12-05 15:55:51.317991-05
237	021408	021904	ANCASH	SIHUAS	CASHAPAMPA	2025-12-05 15:55:51.318565-05	2025-12-05 15:55:51.318565-05
238	021403	021905	ANCASH	SIHUAS	CHINGALPO	2025-12-05 15:55:51.319086-05	2025-12-05 15:55:51.319086-05
239	021404	021906	ANCASH	SIHUAS	HUAYLLABAMBA	2025-12-05 15:55:51.31963-05	2025-12-05 15:55:51.31963-05
240	021405	021907	ANCASH	SIHUAS	QUICHES	2025-12-05 15:55:51.320787-05	2025-12-05 15:55:51.320787-05
241	021409	021908	ANCASH	SIHUAS	RAGASH	2025-12-05 15:55:51.321314-05	2025-12-05 15:55:51.321314-05
242	021410	021909	ANCASH	SIHUAS	SAN JUAN	2025-12-05 15:55:51.322014-05	2025-12-05 15:55:51.322014-05
243	021406	021910	ANCASH	SIHUAS	SICSIBAMBA	2025-12-05 15:55:51.323092-05	2025-12-05 15:55:51.323092-05
244	021501	022001	ANCASH	YUNGAY	YUNGAY	2025-12-05 15:55:51.32406-05	2025-12-05 15:55:51.32406-05
245	021502	022002	ANCASH	YUNGAY	CASCAPARA	2025-12-05 15:55:51.324642-05	2025-12-05 15:55:51.324642-05
246	021503	022003	ANCASH	YUNGAY	MANCOS	2025-12-05 15:55:51.325152-05	2025-12-05 15:55:51.325152-05
247	021504	022004	ANCASH	YUNGAY	MATACOTO	2025-12-05 15:55:51.326224-05	2025-12-05 15:55:51.326224-05
248	021505	022005	ANCASH	YUNGAY	QUILLO	2025-12-05 15:55:51.327469-05	2025-12-05 15:55:51.327469-05
249	021506	022006	ANCASH	YUNGAY	RANRAHIRCA	2025-12-05 15:55:51.328269-05	2025-12-05 15:55:51.328269-05
250	021507	022007	ANCASH	YUNGAY	SHUPLUY	2025-12-05 15:55:51.329542-05	2025-12-05 15:55:51.329542-05
251	021508	022008	ANCASH	YUNGAY	YANAMA	2025-12-05 15:55:51.330909-05	2025-12-05 15:55:51.330909-05
252	030101	030101	APURIMAC	ABANCAY	ABANCAY	2025-12-05 15:55:51.331445-05	2025-12-05 15:55:51.331445-05
253	030104	030102	APURIMAC	ABANCAY	CHACOCHE	2025-12-05 15:55:51.332611-05	2025-12-05 15:55:51.332611-05
254	030102	030103	APURIMAC	ABANCAY	CIRCA	2025-12-05 15:55:51.333892-05	2025-12-05 15:55:51.333892-05
255	030103	030104	APURIMAC	ABANCAY	CURAHUASI	2025-12-05 15:55:51.334424-05	2025-12-05 15:55:51.334424-05
256	030105	030105	APURIMAC	ABANCAY	HUANIPACA	2025-12-05 15:55:51.335679-05	2025-12-05 15:55:51.335679-05
257	030106	030106	APURIMAC	ABANCAY	LAMBRAMA	2025-12-05 15:55:51.336242-05	2025-12-05 15:55:51.336242-05
258	030107	030107	APURIMAC	ABANCAY	PICHIRHUA	2025-12-05 15:55:51.337389-05	2025-12-05 15:55:51.337389-05
259	030108	030108	APURIMAC	ABANCAY	SAN PEDRO DE CACHORA	2025-12-05 15:55:51.337918-05	2025-12-05 15:55:51.337918-05
260	030109	030109	APURIMAC	ABANCAY	TAMBURCO	2025-12-05 15:55:51.339093-05	2025-12-05 15:55:51.339093-05
261	030301	030201	APURIMAC	ANDAHUAYLAS	ANDAHUAYLAS	2025-12-05 15:55:51.339988-05	2025-12-05 15:55:51.339988-05
262	030302	030202	APURIMAC	ANDAHUAYLAS	ANDARAPA	2025-12-05 15:55:51.340614-05	2025-12-05 15:55:51.340614-05
263	030303	030203	APURIMAC	ANDAHUAYLAS	CHIARA	2025-12-05 15:55:51.341748-05	2025-12-05 15:55:51.341748-05
264	030304	030204	APURIMAC	ANDAHUAYLAS	HUANCARAMA	2025-12-05 15:55:51.342288-05	2025-12-05 15:55:51.342288-05
265	030305	030205	APURIMAC	ANDAHUAYLAS	HUANCARAY	2025-12-05 15:55:51.343357-05	2025-12-05 15:55:51.343357-05
266	030317	030206	APURIMAC	ANDAHUAYLAS	HUAYANA	2025-12-05 15:55:51.343901-05	2025-12-05 15:55:51.343901-05
267	030306	030207	APURIMAC	ANDAHUAYLAS	KISHUARA	2025-12-05 15:55:51.345027-05	2025-12-05 15:55:51.345027-05
268	030307	030208	APURIMAC	ANDAHUAYLAS	PACOBAMBA	2025-12-05 15:55:51.345891-05	2025-12-05 15:55:51.345891-05
269	030313	030209	APURIMAC	ANDAHUAYLAS	PACUCHA	2025-12-05 15:55:51.346961-05	2025-12-05 15:55:51.346961-05
270	030308	030210	APURIMAC	ANDAHUAYLAS	PAMPACHIRI	2025-12-05 15:55:51.347544-05	2025-12-05 15:55:51.347544-05
271	030314	030211	APURIMAC	ANDAHUAYLAS	POMACOCHA	2025-12-05 15:55:51.348813-05	2025-12-05 15:55:51.348813-05
272	030309	030212	APURIMAC	ANDAHUAYLAS	SAN ANTONIO DE CACHI	2025-12-05 15:55:51.349823-05	2025-12-05 15:55:51.349823-05
273	030310	030213	APURIMAC	ANDAHUAYLAS	SAN JERONIMO	2025-12-05 15:55:51.351058-05	2025-12-05 15:55:51.351058-05
274	030318	030214	APURIMAC	ANDAHUAYLAS	SAN MIGUEL DE CHACCRAMPA	2025-12-05 15:55:51.351703-05	2025-12-05 15:55:51.351703-05
275	030315	030215	APURIMAC	ANDAHUAYLAS	SANTA MARIA DE CHICMO	2025-12-05 15:55:51.352969-05	2025-12-05 15:55:51.352969-05
276	030311	030216	APURIMAC	ANDAHUAYLAS	TALAVERA	2025-12-05 15:55:51.354039-05	2025-12-05 15:55:51.354039-05
277	030316	030217	APURIMAC	ANDAHUAYLAS	TUMAY HUARACA	2025-12-05 15:55:51.354561-05	2025-12-05 15:55:51.354561-05
278	030312	030218	APURIMAC	ANDAHUAYLAS	TURPO	2025-12-05 15:55:51.355641-05	2025-12-05 15:55:51.355641-05
279	030319	030219	APURIMAC	ANDAHUAYLAS	KAQUIABAMBA	2025-12-05 15:55:51.356171-05	2025-12-05 15:55:51.356171-05
280	030320	030220	APURIMAC	ANDAHUAYLAS	JOSE MARIA ARGUEDAS	2025-12-05 15:55:51.356872-05	2025-12-05 15:55:51.356872-05
281	030401	030301	APURIMAC	ANTABAMBA	ANTABAMBA	2025-12-05 15:55:51.357493-05	2025-12-05 15:55:51.357493-05
282	030402	030302	APURIMAC	ANTABAMBA	EL ORO	2025-12-05 15:55:51.358174-05	2025-12-05 15:55:51.358174-05
283	030403	030303	APURIMAC	ANTABAMBA	HUAQUIRCA	2025-12-05 15:55:51.358825-05	2025-12-05 15:55:51.358825-05
284	030404	030304	APURIMAC	ANTABAMBA	JUAN ESPINOZA MEDRANO	2025-12-05 15:55:51.359974-05	2025-12-05 15:55:51.359974-05
285	030405	030305	APURIMAC	ANTABAMBA	OROPESA	2025-12-05 15:55:51.360597-05	2025-12-05 15:55:51.360597-05
286	030406	030306	APURIMAC	ANTABAMBA	PACHACONAS	2025-12-05 15:55:51.361482-05	2025-12-05 15:55:51.361482-05
287	030407	030307	APURIMAC	ANTABAMBA	SABAINO	2025-12-05 15:55:51.362537-05	2025-12-05 15:55:51.362537-05
288	030201	030401	APURIMAC	AYMARAES	CHALHUANCA	2025-12-05 15:55:51.363121-05	2025-12-05 15:55:51.363121-05
289	030202	030402	APURIMAC	AYMARAES	CAPAYA	2025-12-05 15:55:51.364258-05	2025-12-05 15:55:51.364258-05
290	030203	030403	APURIMAC	AYMARAES	CARAYBAMBA	2025-12-05 15:55:51.365146-05	2025-12-05 15:55:51.365146-05
291	030206	030404	APURIMAC	AYMARAES	CHAPIMARCA	2025-12-05 15:55:51.36576-05	2025-12-05 15:55:51.36576-05
292	030204	030405	APURIMAC	AYMARAES	COLCABAMBA	2025-12-05 15:55:51.367159-05	2025-12-05 15:55:51.367159-05
293	030205	030406	APURIMAC	AYMARAES	COTARUSE	2025-12-05 15:55:51.367707-05	2025-12-05 15:55:51.367707-05
294	030207	030407	APURIMAC	AYMARAES	HUAYLLO	2025-12-05 15:55:51.368272-05	2025-12-05 15:55:51.368272-05
295	030217	030408	APURIMAC	AYMARAES	JUSTO APU SAHUARAURA	2025-12-05 15:55:51.369345-05	2025-12-05 15:55:51.369345-05
296	030208	030409	APURIMAC	AYMARAES	LUCRE	2025-12-05 15:55:51.369876-05	2025-12-05 15:55:51.369876-05
297	030209	030410	APURIMAC	AYMARAES	POCOHUANCA	2025-12-05 15:55:51.370386-05	2025-12-05 15:55:51.370386-05
298	030216	030411	APURIMAC	AYMARAES	SAN JUAN DE CHACÑA	2025-12-05 15:55:51.370932-05	2025-12-05 15:55:51.370932-05
299	030210	030412	APURIMAC	AYMARAES	SAÑAYCA	2025-12-05 15:55:51.371446-05	2025-12-05 15:55:51.371446-05
300	030211	030413	APURIMAC	AYMARAES	SORAYA	2025-12-05 15:55:51.372-05	2025-12-05 15:55:51.372-05
301	030212	030414	APURIMAC	AYMARAES	TAPAIRIHUA	2025-12-05 15:55:51.372527-05	2025-12-05 15:55:51.372527-05
302	030213	030415	APURIMAC	AYMARAES	TINTAY	2025-12-05 15:55:51.373182-05	2025-12-05 15:55:51.373182-05
303	030214	030416	APURIMAC	AYMARAES	TORAYA	2025-12-05 15:55:51.373893-05	2025-12-05 15:55:51.373893-05
304	030215	030417	APURIMAC	AYMARAES	YANACA	2025-12-05 15:55:51.374505-05	2025-12-05 15:55:51.374505-05
305	030501	030501	APURIMAC	COTABAMBAS	TAMBOBAMBA	2025-12-05 15:55:51.375026-05	2025-12-05 15:55:51.375026-05
306	030503	030502	APURIMAC	COTABAMBAS	COTABAMBAS	2025-12-05 15:55:51.375026-05	2025-12-05 15:55:51.375026-05
307	030502	030503	APURIMAC	COTABAMBAS	COYLLURQUI	2025-12-05 15:55:51.376085-05	2025-12-05 15:55:51.376085-05
308	030504	030504	APURIMAC	COTABAMBAS	HAQUIRA	2025-12-05 15:55:51.376633-05	2025-12-05 15:55:51.376633-05
309	030505	030505	APURIMAC	COTABAMBAS	MARA	2025-12-05 15:55:51.377371-05	2025-12-05 15:55:51.377371-05
310	030506	030506	APURIMAC	COTABAMBAS	CHALLHUAHUACHO	2025-12-05 15:55:51.378666-05	2025-12-05 15:55:51.378666-05
311	030701	030601	APURIMAC	CHINCHEROS	CHINCHEROS	2025-12-05 15:55:51.379189-05	2025-12-05 15:55:51.379189-05
312	030705	030602	APURIMAC	CHINCHEROS	ANCO-HUALLO	2025-12-05 15:55:51.380441-05	2025-12-05 15:55:51.380441-05
313	030704	030603	APURIMAC	CHINCHEROS	COCHARCAS	2025-12-05 15:55:51.380964-05	2025-12-05 15:55:51.380964-05
314	030706	030604	APURIMAC	CHINCHEROS	HUACCANA	2025-12-05 15:55:51.382217-05	2025-12-05 15:55:51.382217-05
315	030703	030605	APURIMAC	CHINCHEROS	OCOBAMBA	2025-12-05 15:55:51.383051-05	2025-12-05 15:55:51.383051-05
316	030702	030606	APURIMAC	CHINCHEROS	ONGOY	2025-12-05 15:55:51.384209-05	2025-12-05 15:55:51.384209-05
317	030707	030607	APURIMAC	CHINCHEROS	URANMARCA	2025-12-05 15:55:51.385381-05	2025-12-05 15:55:51.385381-05
318	030708	030608	APURIMAC	CHINCHEROS	RANRACANCHA	2025-12-05 15:55:51.385896-05	2025-12-05 15:55:51.385896-05
319	030709	030609	APURIMAC	CHINCHEROS	ROCCHACC	2025-12-05 15:55:51.386417-05	2025-12-05 15:55:51.386417-05
320	030710	030610	APURIMAC	CHINCHEROS	EL PORVENIR	2025-12-05 15:55:51.3877-05	2025-12-05 15:55:51.3877-05
321	030711	030611	APURIMAC	CHINCHEROS	LOS CHANKAS	2025-12-05 15:55:51.388255-05	2025-12-05 15:55:51.388255-05
322	030712	030612	APURIMAC	CHINCHEROS	AHUAYRO	2025-12-05 15:55:51.389328-05	2025-12-05 15:55:51.389328-05
323	030601	030701	APURIMAC	GRAU	CHUQUIBAMBILLA	2025-12-05 15:55:51.390402-05	2025-12-05 15:55:51.390402-05
324	030602	030702	APURIMAC	GRAU	CURPAHUASI	2025-12-05 15:55:51.390983-05	2025-12-05 15:55:51.390983-05
325	030605	030703	APURIMAC	GRAU	GAMARRA	2025-12-05 15:55:51.39153-05	2025-12-05 15:55:51.39153-05
326	030603	030704	APURIMAC	GRAU	HUAYLLATI	2025-12-05 15:55:51.393665-05	2025-12-05 15:55:51.393665-05
327	030604	030705	APURIMAC	GRAU	MAMARA	2025-12-05 15:55:51.394763-05	2025-12-05 15:55:51.394763-05
328	030606	030706	APURIMAC	GRAU	MICAELA BASTIDAS	2025-12-05 15:55:51.395886-05	2025-12-05 15:55:51.395886-05
329	030608	030707	APURIMAC	GRAU	PATAYPAMPA	2025-12-05 15:55:51.396409-05	2025-12-05 15:55:51.396409-05
330	030607	030708	APURIMAC	GRAU	PROGRESO	2025-12-05 15:55:51.397183-05	2025-12-05 15:55:51.397183-05
331	030609	030709	APURIMAC	GRAU	SAN ANTONIO	2025-12-05 15:55:51.398235-05	2025-12-05 15:55:51.398235-05
332	030613	030710	APURIMAC	GRAU	SANTA ROSA	2025-12-05 15:55:51.399592-05	2025-12-05 15:55:51.399592-05
333	030610	030711	APURIMAC	GRAU	TURPAY	2025-12-05 15:55:51.401039-05	2025-12-05 15:55:51.401039-05
334	030611	030712	APURIMAC	GRAU	VILCABAMBA	2025-12-05 15:55:51.40214-05	2025-12-05 15:55:51.40214-05
335	030612	030713	APURIMAC	GRAU	VIRUNDO	2025-12-05 15:55:51.40267-05	2025-12-05 15:55:51.40267-05
336	030614	030714	APURIMAC	GRAU	CURASCO	2025-12-05 15:55:51.403935-05	2025-12-05 15:55:51.403935-05
337	040101	040101	AREQUIPA	AREQUIPA	AREQUIPA	2025-12-05 15:55:51.40455-05	2025-12-05 15:55:51.40455-05
338	040128	040102	AREQUIPA	AREQUIPA	ALTO SELVA ALEGRE	2025-12-05 15:55:51.405796-05	2025-12-05 15:55:51.405796-05
339	040102	040103	AREQUIPA	AREQUIPA	CAYMA	2025-12-05 15:55:51.406435-05	2025-12-05 15:55:51.406435-05
340	040103	040104	AREQUIPA	AREQUIPA	CERRO COLORADO	2025-12-05 15:55:51.407548-05	2025-12-05 15:55:51.407548-05
341	040104	040105	AREQUIPA	AREQUIPA	CHARACATO	2025-12-05 15:55:51.40807-05	2025-12-05 15:55:51.40807-05
342	040105	040106	AREQUIPA	AREQUIPA	CHIGUATA	2025-12-05 15:55:51.409372-05	2025-12-05 15:55:51.409372-05
343	040127	040107	AREQUIPA	AREQUIPA	JACOBO HUNTER	2025-12-05 15:55:51.409943-05	2025-12-05 15:55:51.409943-05
344	040106	040108	AREQUIPA	AREQUIPA	LA JOYA	2025-12-05 15:55:51.411108-05	2025-12-05 15:55:51.411108-05
345	040126	040109	AREQUIPA	AREQUIPA	MARIANO MELGAR	2025-12-05 15:55:51.412169-05	2025-12-05 15:55:51.412169-05
346	040107	040110	AREQUIPA	AREQUIPA	MIRAFLORES	2025-12-05 15:55:51.413317-05	2025-12-05 15:55:51.413317-05
347	040108	040111	AREQUIPA	AREQUIPA	MOLLEBAYA	2025-12-05 15:55:51.413842-05	2025-12-05 15:55:51.413842-05
348	040109	040112	AREQUIPA	AREQUIPA	PAUCARPATA	2025-12-05 15:55:51.414883-05	2025-12-05 15:55:51.414883-05
349	040110	040113	AREQUIPA	AREQUIPA	POCSI	2025-12-05 15:55:51.416262-05	2025-12-05 15:55:51.416262-05
350	040111	040114	AREQUIPA	AREQUIPA	POLOBAYA	2025-12-05 15:55:51.416819-05	2025-12-05 15:55:51.416819-05
351	040112	040115	AREQUIPA	AREQUIPA	QUEQUEÑA	2025-12-05 15:55:51.417338-05	2025-12-05 15:55:51.417338-05
352	040113	040116	AREQUIPA	AREQUIPA	SABANDIA	2025-12-05 15:55:51.418671-05	2025-12-05 15:55:51.418671-05
353	040114	040117	AREQUIPA	AREQUIPA	SACHACA	2025-12-05 15:55:51.419333-05	2025-12-05 15:55:51.419333-05
354	040115	040118	AREQUIPA	AREQUIPA	SAN JUAN DE SIGUAS	2025-12-05 15:55:51.419956-05	2025-12-05 15:55:51.419956-05
355	040116	040119	AREQUIPA	AREQUIPA	SAN JUAN DE TARUCANI	2025-12-05 15:55:51.420995-05	2025-12-05 15:55:51.420995-05
356	040117	040120	AREQUIPA	AREQUIPA	SANTA ISABEL DE SIGUAS	2025-12-05 15:55:51.421554-05	2025-12-05 15:55:51.421554-05
357	110401	120501	JUNIN	JUNIN	JUNIN	2025-12-05 15:55:51.422589-05	2025-12-05 15:55:51.422589-05
358	040118	040121	AREQUIPA	AREQUIPA	SANTA RITA DE SIGUAS	2025-12-05 15:55:51.423586-05	2025-12-05 15:55:51.423586-05
359	040119	040122	AREQUIPA	AREQUIPA	SOCABAYA	2025-12-05 15:55:51.424192-05	2025-12-05 15:55:51.424192-05
360	040120	040123	AREQUIPA	AREQUIPA	TIABAYA	2025-12-05 15:55:51.425254-05	2025-12-05 15:55:51.425254-05
361	040121	040124	AREQUIPA	AREQUIPA	UCHUMAYO	2025-12-05 15:55:51.425771-05	2025-12-05 15:55:51.425771-05
362	040122	040125	AREQUIPA	AREQUIPA	VITOR	2025-12-05 15:55:51.42684-05	2025-12-05 15:55:51.42684-05
363	040123	040126	AREQUIPA	AREQUIPA	YANAHUARA	2025-12-05 15:55:51.427365-05	2025-12-05 15:55:51.427365-05
364	040124	040127	AREQUIPA	AREQUIPA	YARABAMBA	2025-12-05 15:55:51.428691-05	2025-12-05 15:55:51.428691-05
365	040125	040128	AREQUIPA	AREQUIPA	YURA	2025-12-05 15:55:51.42974-05	2025-12-05 15:55:51.42974-05
366	040129	040129	AREQUIPA	AREQUIPA	JOSE LUIS BUSTAMANTE Y RIVERO	2025-12-05 15:55:51.430292-05	2025-12-05 15:55:51.430292-05
367	040301	040201	AREQUIPA	CAMANA	CAMANA	2025-12-05 15:55:51.431394-05	2025-12-05 15:55:51.431394-05
368	040302	040202	AREQUIPA	CAMANA	JOSE MARIA QUIMPER	2025-12-05 15:55:51.431994-05	2025-12-05 15:55:51.431994-05
369	040303	040203	AREQUIPA	CAMANA	MARIANO NICOLAS VALCARCEL	2025-12-05 15:55:51.43318-05	2025-12-05 15:55:51.43318-05
370	040304	040204	AREQUIPA	CAMANA	MARISCAL CACERES	2025-12-05 15:55:51.434477-05	2025-12-05 15:55:51.434477-05
371	040305	040205	AREQUIPA	CAMANA	NICOLAS DE PIEROLA	2025-12-05 15:55:51.442226-05	2025-12-05 15:55:51.442226-05
372	040306	040206	AREQUIPA	CAMANA	OCOÑA	2025-12-05 15:55:51.445476-05	2025-12-05 15:55:51.445476-05
373	040307	040207	AREQUIPA	CAMANA	QUILCA	2025-12-05 15:55:51.446512-05	2025-12-05 15:55:51.446512-05
374	040308	040208	AREQUIPA	CAMANA	SAMUEL PASTOR	2025-12-05 15:55:51.447727-05	2025-12-05 15:55:51.447727-05
375	040401	040301	AREQUIPA	CARAVELI	CARAVELI	2025-12-05 15:55:51.448825-05	2025-12-05 15:55:51.448825-05
376	040402	040302	AREQUIPA	CARAVELI	ACARI	2025-12-05 15:55:51.450139-05	2025-12-05 15:55:51.450139-05
377	040403	040303	AREQUIPA	CARAVELI	ATICO	2025-12-05 15:55:51.451827-05	2025-12-05 15:55:51.451827-05
378	040404	040304	AREQUIPA	CARAVELI	ATIQUIPA	2025-12-05 15:55:51.453016-05	2025-12-05 15:55:51.453016-05
379	040405	040305	AREQUIPA	CARAVELI	BELLA UNION	2025-12-05 15:55:51.454093-05	2025-12-05 15:55:51.454093-05
380	040406	040306	AREQUIPA	CARAVELI	CAHUACHO	2025-12-05 15:55:51.455125-05	2025-12-05 15:55:51.455125-05
381	040407	040307	AREQUIPA	CARAVELI	CHALA	2025-12-05 15:55:51.455673-05	2025-12-05 15:55:51.455673-05
382	040408	040308	AREQUIPA	CARAVELI	CHAPARRA	2025-12-05 15:55:51.456745-05	2025-12-05 15:55:51.456745-05
383	040409	040309	AREQUIPA	CARAVELI	HUANUHUANU	2025-12-05 15:55:51.45727-05	2025-12-05 15:55:51.45727-05
384	040410	040310	AREQUIPA	CARAVELI	JAQUI	2025-12-05 15:55:51.457791-05	2025-12-05 15:55:51.457791-05
385	040411	040311	AREQUIPA	CARAVELI	LOMAS	2025-12-05 15:55:51.458472-05	2025-12-05 15:55:51.458472-05
386	040412	040312	AREQUIPA	CARAVELI	QUICACHA	2025-12-05 15:55:51.460017-05	2025-12-05 15:55:51.460017-05
387	040413	040313	AREQUIPA	CARAVELI	YAUCA	2025-12-05 15:55:51.461062-05	2025-12-05 15:55:51.461062-05
388	040501	040401	AREQUIPA	CASTILLA	APLAO	2025-12-05 15:55:51.462187-05	2025-12-05 15:55:51.462187-05
389	040502	040402	AREQUIPA	CASTILLA	ANDAGUA	2025-12-05 15:55:51.463242-05	2025-12-05 15:55:51.463242-05
390	040503	040403	AREQUIPA	CASTILLA	AYO	2025-12-05 15:55:51.463827-05	2025-12-05 15:55:51.463827-05
391	040504	040404	AREQUIPA	CASTILLA	CHACHAS	2025-12-05 15:55:51.464338-05	2025-12-05 15:55:51.464338-05
392	040505	040405	AREQUIPA	CASTILLA	CHILCAYMARCA	2025-12-05 15:55:51.465418-05	2025-12-05 15:55:51.465418-05
393	040506	040406	AREQUIPA	CASTILLA	CHOCO	2025-12-05 15:55:51.467803-05	2025-12-05 15:55:51.467803-05
394	040507	040407	AREQUIPA	CASTILLA	HUANCARQUI	2025-12-05 15:55:51.469233-05	2025-12-05 15:55:51.469233-05
395	040508	040408	AREQUIPA	CASTILLA	MACHAGUAY	2025-12-05 15:55:51.470278-05	2025-12-05 15:55:51.470278-05
396	040509	040409	AREQUIPA	CASTILLA	ORCOPAMPA	2025-12-05 15:55:51.471456-05	2025-12-05 15:55:51.471456-05
397	040510	040410	AREQUIPA	CASTILLA	PAMPACOLCA	2025-12-05 15:55:51.471976-05	2025-12-05 15:55:51.471976-05
398	040511	040411	AREQUIPA	CASTILLA	TIPAN	2025-12-05 15:55:51.473123-05	2025-12-05 15:55:51.473123-05
399	040513	040412	AREQUIPA	CASTILLA	UÑON	2025-12-05 15:55:51.473662-05	2025-12-05 15:55:51.473662-05
400	040512	040413	AREQUIPA	CASTILLA	URACA	2025-12-05 15:55:51.474779-05	2025-12-05 15:55:51.474779-05
401	040514	040414	AREQUIPA	CASTILLA	VIRACO	2025-12-05 15:55:51.475301-05	2025-12-05 15:55:51.475301-05
402	040201	040501	AREQUIPA	CAYLLOMA	CHIVAY	2025-12-05 15:55:51.476369-05	2025-12-05 15:55:51.476369-05
403	040202	040502	AREQUIPA	CAYLLOMA	ACHOMA	2025-12-05 15:55:51.477481-05	2025-12-05 15:55:51.477481-05
404	040203	040503	AREQUIPA	CAYLLOMA	CABANACONDE	2025-12-05 15:55:51.477992-05	2025-12-05 15:55:51.477992-05
405	040205	040504	AREQUIPA	CAYLLOMA	CALLALLI	2025-12-05 15:55:51.479068-05	2025-12-05 15:55:51.479068-05
406	040204	040505	AREQUIPA	CAYLLOMA	CAYLLOMA	2025-12-05 15:55:51.479688-05	2025-12-05 15:55:51.479688-05
407	040206	040506	AREQUIPA	CAYLLOMA	COPORAQUE	2025-12-05 15:55:51.480965-05	2025-12-05 15:55:51.480965-05
408	040207	040507	AREQUIPA	CAYLLOMA	HUAMBO	2025-12-05 15:55:51.482119-05	2025-12-05 15:55:51.482119-05
409	040208	040508	AREQUIPA	CAYLLOMA	HUANCA	2025-12-05 15:55:51.483415-05	2025-12-05 15:55:51.483415-05
410	040209	040509	AREQUIPA	CAYLLOMA	ICHUPAMPA	2025-12-05 15:55:51.483941-05	2025-12-05 15:55:51.483941-05
411	040210	040510	AREQUIPA	CAYLLOMA	LARI	2025-12-05 15:55:51.484665-05	2025-12-05 15:55:51.484665-05
412	040211	040511	AREQUIPA	CAYLLOMA	LLUTA	2025-12-05 15:55:51.485384-05	2025-12-05 15:55:51.485384-05
413	040212	040512	AREQUIPA	CAYLLOMA	MACA	2025-12-05 15:55:51.486098-05	2025-12-05 15:55:51.486098-05
414	040213	040513	AREQUIPA	CAYLLOMA	MADRIGAL	2025-12-05 15:55:51.486737-05	2025-12-05 15:55:51.486737-05
415	040214	040514	AREQUIPA	CAYLLOMA	SAN ANTONIO DE CHUCA	2025-12-05 15:55:51.487373-05	2025-12-05 15:55:51.487373-05
416	040215	040515	AREQUIPA	CAYLLOMA	SIBAYO	2025-12-05 15:55:51.488046-05	2025-12-05 15:55:51.488046-05
417	040216	040516	AREQUIPA	CAYLLOMA	TAPAY	2025-12-05 15:55:51.4891-05	2025-12-05 15:55:51.4891-05
418	040217	040517	AREQUIPA	CAYLLOMA	TISCO	2025-12-05 15:55:51.489628-05	2025-12-05 15:55:51.489628-05
419	040218	040518	AREQUIPA	CAYLLOMA	TUTI	2025-12-05 15:55:51.490804-05	2025-12-05 15:55:51.490804-05
420	040219	040519	AREQUIPA	CAYLLOMA	YANQUE	2025-12-05 15:55:51.491324-05	2025-12-05 15:55:51.491324-05
421	040220	040520	AREQUIPA	CAYLLOMA	MAJES	2025-12-05 15:55:51.492361-05	2025-12-05 15:55:51.492361-05
422	040601	040601	AREQUIPA	CONDESUYOS	CHUQUIBAMBA	2025-12-05 15:55:51.49296-05	2025-12-05 15:55:51.49296-05
423	040602	040602	AREQUIPA	CONDESUYOS	ANDARAY	2025-12-05 15:55:51.493487-05	2025-12-05 15:55:51.493487-05
424	040603	040603	AREQUIPA	CONDESUYOS	CAYARANI	2025-12-05 15:55:51.494821-05	2025-12-05 15:55:51.494821-05
425	040604	040604	AREQUIPA	CONDESUYOS	CHICHAS	2025-12-05 15:55:51.495348-05	2025-12-05 15:55:51.495348-05
426	040605	040605	AREQUIPA	CONDESUYOS	IRAY	2025-12-05 15:55:51.495875-05	2025-12-05 15:55:51.495875-05
427	040608	040606	AREQUIPA	CONDESUYOS	RIO GRANDE	2025-12-05 15:55:51.495875-05	2025-12-05 15:55:51.495875-05
428	040606	040607	AREQUIPA	CONDESUYOS	SALAMANCA	2025-12-05 15:55:51.496401-05	2025-12-05 15:55:51.496401-05
429	040607	040608	AREQUIPA	CONDESUYOS	YANAQUIHUA	2025-12-05 15:55:51.497468-05	2025-12-05 15:55:51.497468-05
430	040701	040701	AREQUIPA	ISLAY	MOLLENDO	2025-12-05 15:55:51.498095-05	2025-12-05 15:55:51.498095-05
431	040702	040702	AREQUIPA	ISLAY	COCACHACRA	2025-12-05 15:55:51.498622-05	2025-12-05 15:55:51.498622-05
432	040703	040703	AREQUIPA	ISLAY	DEAN VALDIVIA	2025-12-05 15:55:51.499135-05	2025-12-05 15:55:51.499135-05
433	040704	040704	AREQUIPA	ISLAY	ISLAY	2025-12-05 15:55:51.499734-05	2025-12-05 15:55:51.499734-05
434	040705	040705	AREQUIPA	ISLAY	MEJIA	2025-12-05 15:55:51.500344-05	2025-12-05 15:55:51.500344-05
435	040706	040706	AREQUIPA	ISLAY	PUNTA DE BOMBON	2025-12-05 15:55:51.500861-05	2025-12-05 15:55:51.500861-05
436	040801	040801	AREQUIPA	LA UNION	COTAHUASI	2025-12-05 15:55:51.501373-05	2025-12-05 15:55:51.501373-05
437	040802	040802	AREQUIPA	LA UNION	ALCA	2025-12-05 15:55:51.50191-05	2025-12-05 15:55:51.50191-05
438	040803	040803	AREQUIPA	LA UNION	CHARCANA	2025-12-05 15:55:51.50191-05	2025-12-05 15:55:51.50191-05
439	040804	040804	AREQUIPA	LA UNION	HUAYNACOTAS	2025-12-05 15:55:51.502943-05	2025-12-05 15:55:51.502943-05
440	040805	040805	AREQUIPA	LA UNION	PAMPAMARCA	2025-12-05 15:55:51.503478-05	2025-12-05 15:55:51.503478-05
441	040806	040806	AREQUIPA	LA UNION	PUYCA	2025-12-05 15:55:51.504324-05	2025-12-05 15:55:51.504324-05
442	040807	040807	AREQUIPA	LA UNION	QUECHUALLA	2025-12-05 15:55:51.505438-05	2025-12-05 15:55:51.505438-05
443	040808	040808	AREQUIPA	LA UNION	SAYLA	2025-12-05 15:55:51.506074-05	2025-12-05 15:55:51.506074-05
444	040809	040809	AREQUIPA	LA UNION	TAURIA	2025-12-05 15:55:51.507145-05	2025-12-05 15:55:51.507145-05
445	040810	040810	AREQUIPA	LA UNION	TOMEPAMPA	2025-12-05 15:55:51.507731-05	2025-12-05 15:55:51.507731-05
446	040811	040811	AREQUIPA	LA UNION	TORO	2025-12-05 15:55:51.509312-05	2025-12-05 15:55:51.509312-05
447	050101	050101	AYACUCHO	HUAMANGA	AYACUCHO	2025-12-05 15:55:51.510008-05	2025-12-05 15:55:51.510008-05
448	050111	050102	AYACUCHO	HUAMANGA	ACOCRO	2025-12-05 15:55:51.511323-05	2025-12-05 15:55:51.511323-05
449	050102	050103	AYACUCHO	HUAMANGA	ACOS VINCHOS	2025-12-05 15:55:51.5129-05	2025-12-05 15:55:51.5129-05
450	050103	050104	AYACUCHO	HUAMANGA	CARMEN ALTO	2025-12-05 15:55:51.514491-05	2025-12-05 15:55:51.514491-05
451	050104	050105	AYACUCHO	HUAMANGA	CHIARA	2025-12-05 15:55:51.515676-05	2025-12-05 15:55:51.515676-05
452	050113	050106	AYACUCHO	HUAMANGA	OCROS	2025-12-05 15:55:51.516208-05	2025-12-05 15:55:51.516208-05
453	050114	050107	AYACUCHO	HUAMANGA	PACAYCASA	2025-12-05 15:55:51.517508-05	2025-12-05 15:55:51.517508-05
454	050105	050108	AYACUCHO	HUAMANGA	QUINUA	2025-12-05 15:55:51.518108-05	2025-12-05 15:55:51.518108-05
455	050106	050109	AYACUCHO	HUAMANGA	SAN JOSE DE TICLLAS	2025-12-05 15:55:51.519311-05	2025-12-05 15:55:51.519311-05
456	050107	050110	AYACUCHO	HUAMANGA	SAN JUAN BAUTISTA	2025-12-05 15:55:51.519844-05	2025-12-05 15:55:51.519844-05
457	050108	050111	AYACUCHO	HUAMANGA	SANTIAGO DE PISCHA	2025-12-05 15:55:51.52057-05	2025-12-05 15:55:51.52057-05
458	050112	050112	AYACUCHO	HUAMANGA	SOCOS	2025-12-05 15:55:51.520991-05	2025-12-05 15:55:51.520991-05
459	050110	050113	AYACUCHO	HUAMANGA	TAMBILLO	2025-12-05 15:55:51.52205-05	2025-12-05 15:55:51.52205-05
460	050109	050114	AYACUCHO	HUAMANGA	VINCHOS	2025-12-05 15:55:51.52205-05	2025-12-05 15:55:51.52205-05
461	050115	050115	AYACUCHO	HUAMANGA	JESUS NAZARENO	2025-12-05 15:55:51.522574-05	2025-12-05 15:55:51.522574-05
462	050116	050116	AYACUCHO	HUAMANGA	ANDRES AVELINO CACERES DORREGA	2025-12-05 15:55:51.523176-05	2025-12-05 15:55:51.523176-05
463	050201	050201	AYACUCHO	CANGALLO	CANGALLO	2025-12-05 15:55:51.523744-05	2025-12-05 15:55:51.523744-05
464	050204	050202	AYACUCHO	CANGALLO	CHUSCHI	2025-12-05 15:55:51.523744-05	2025-12-05 15:55:51.523744-05
465	050206	050203	AYACUCHO	CANGALLO	LOS MOROCHUCOS	2025-12-05 15:55:51.524261-05	2025-12-05 15:55:51.524261-05
466	050211	050204	AYACUCHO	CANGALLO	MARIA PARADO DE BELLIDO	2025-12-05 15:55:51.524827-05	2025-12-05 15:55:51.524827-05
467	050207	050205	AYACUCHO	CANGALLO	PARAS	2025-12-05 15:55:51.525364-05	2025-12-05 15:55:51.525364-05
468	050208	050206	AYACUCHO	CANGALLO	TOTOS	2025-12-05 15:55:51.525708-05	2025-12-05 15:55:51.525708-05
469	050801	050301	AYACUCHO	HUANCA SANCOS	SANCOS	2025-12-05 15:55:51.526779-05	2025-12-05 15:55:51.526779-05
470	050804	050302	AYACUCHO	HUANCA SANCOS	CARAPO	2025-12-05 15:55:51.527298-05	2025-12-05 15:55:51.527298-05
471	050802	050303	AYACUCHO	HUANCA SANCOS	SACSAMARCA	2025-12-05 15:55:51.527858-05	2025-12-05 15:55:51.527858-05
472	050803	050304	AYACUCHO	HUANCA SANCOS	SANTIAGO DE LUCANAMARCA	2025-12-05 15:55:51.528548-05	2025-12-05 15:55:51.528548-05
473	050301	050401	AYACUCHO	HUANTA	HUANTA	2025-12-05 15:55:51.52916-05	2025-12-05 15:55:51.52916-05
474	050302	050402	AYACUCHO	HUANTA	AYAHUANCO	2025-12-05 15:55:51.529835-05	2025-12-05 15:55:51.529835-05
475	050303	050403	AYACUCHO	HUANTA	HUAMANGUILLA	2025-12-05 15:55:51.530352-05	2025-12-05 15:55:51.530352-05
476	050304	050404	AYACUCHO	HUANTA	IGUAIN	2025-12-05 15:55:51.530889-05	2025-12-05 15:55:51.530889-05
477	050305	050405	AYACUCHO	HUANTA	LURICOCHA	2025-12-05 15:55:51.531837-05	2025-12-05 15:55:51.531837-05
478	050307	050406	AYACUCHO	HUANTA	SANTILLANA	2025-12-05 15:55:51.532756-05	2025-12-05 15:55:51.532756-05
479	050308	050407	AYACUCHO	HUANTA	SIVIA	2025-12-05 15:55:51.533849-05	2025-12-05 15:55:51.533849-05
480	050309	050408	AYACUCHO	HUANTA	LLOCHEGUA	2025-12-05 15:55:51.534377-05	2025-12-05 15:55:51.534377-05
481	050310	050409	AYACUCHO	HUANTA	CANAYRE	2025-12-05 15:55:51.534898-05	2025-12-05 15:55:51.534898-05
482	050311	050410	AYACUCHO	HUANTA	UCHURACCAY	2025-12-05 15:55:51.535447-05	2025-12-05 15:55:51.535447-05
483	050312	050411	AYACUCHO	HUANTA	PUCACOLPA	2025-12-05 15:55:51.5361-05	2025-12-05 15:55:51.5361-05
484	050313	050412	AYACUCHO	HUANTA	CHACA	2025-12-05 15:55:51.536616-05	2025-12-05 15:55:51.536616-05
485	050314	050413	AYACUCHO	HUANTA	PUTIS	2025-12-05 15:55:51.537153-05	2025-12-05 15:55:51.537153-05
486	050401	050501	AYACUCHO	LA MAR	SAN MIGUEL	2025-12-05 15:55:51.537666-05	2025-12-05 15:55:51.537666-05
487	050402	050502	AYACUCHO	LA MAR	ANCO	2025-12-05 15:55:51.53819-05	2025-12-05 15:55:51.53819-05
488	050403	050503	AYACUCHO	LA MAR	AYNA	2025-12-05 15:55:51.538991-05	2025-12-05 15:55:51.538991-05
489	050404	050504	AYACUCHO	LA MAR	CHILCAS	2025-12-05 15:55:51.539522-05	2025-12-05 15:55:51.539522-05
490	050405	050505	AYACUCHO	LA MAR	CHUNGUI	2025-12-05 15:55:51.540047-05	2025-12-05 15:55:51.540047-05
491	050407	050506	AYACUCHO	LA MAR	LUIS CARRANZA	2025-12-05 15:55:51.540047-05	2025-12-05 15:55:51.540047-05
492	050408	050507	AYACUCHO	LA MAR	SANTA ROSA	2025-12-05 15:55:51.540596-05	2025-12-05 15:55:51.540596-05
493	050406	050508	AYACUCHO	LA MAR	TAMBO	2025-12-05 15:55:51.54112-05	2025-12-05 15:55:51.54112-05
494	050409	050509	AYACUCHO	LA MAR	SAMUGARI	2025-12-05 15:55:51.541662-05	2025-12-05 15:55:51.541662-05
495	050410	050510	AYACUCHO	LA MAR	ANCHIHUAY	2025-12-05 15:55:51.542699-05	2025-12-05 15:55:51.542699-05
496	050411	050511	AYACUCHO	LA MAR	ORONCCOY	2025-12-05 15:55:51.543264-05	2025-12-05 15:55:51.543264-05
497	050412	050512	AYACUCHO	LA MAR	UNION PROGRESO	2025-12-05 15:55:51.544457-05	2025-12-05 15:55:51.544457-05
498	050415	050513	AYACUCHO	LA MAR	RIO MAGDALENA	2025-12-05 15:55:51.545075-05	2025-12-05 15:55:51.545075-05
499	140502	150802	LIMA	HUAURA	AMBAR	2025-12-05 15:55:51.546233-05	2025-12-05 15:55:51.546233-05
500	050414	050514	AYACUCHO	LA MAR	NINABAMBA	2025-12-05 15:55:51.546755-05	2025-12-05 15:55:51.546755-05
501	050413	050515	AYACUCHO	LA MAR	PATIBAMBA	2025-12-05 15:55:51.547265-05	2025-12-05 15:55:51.547265-05
502	050501	050601	AYACUCHO	LUCANAS	PUQUIO	2025-12-05 15:55:51.548288-05	2025-12-05 15:55:51.548288-05
503	050502	050602	AYACUCHO	LUCANAS	AUCARA	2025-12-05 15:55:51.549541-05	2025-12-05 15:55:51.549541-05
504	050503	050603	AYACUCHO	LUCANAS	CABANA	2025-12-05 15:55:51.550082-05	2025-12-05 15:55:51.550082-05
505	050504	050604	AYACUCHO	LUCANAS	CARMEN SALCEDO	2025-12-05 15:55:51.550727-05	2025-12-05 15:55:51.550727-05
506	050506	050605	AYACUCHO	LUCANAS	CHAVIÑA	2025-12-05 15:55:51.551837-05	2025-12-05 15:55:51.551837-05
507	050508	050606	AYACUCHO	LUCANAS	CHIPAO	2025-12-05 15:55:51.552348-05	2025-12-05 15:55:51.552348-05
508	050510	050607	AYACUCHO	LUCANAS	HUAC-HUAS	2025-12-05 15:55:51.552348-05	2025-12-05 15:55:51.552348-05
509	050511	050608	AYACUCHO	LUCANAS	LARAMATE	2025-12-05 15:55:51.55297-05	2025-12-05 15:55:51.55297-05
510	050512	050609	AYACUCHO	LUCANAS	LEONCIO PRADO	2025-12-05 15:55:51.55369-05	2025-12-05 15:55:51.55369-05
511	050514	050610	AYACUCHO	LUCANAS	LLAUTA	2025-12-05 15:55:51.554262-05	2025-12-05 15:55:51.554262-05
512	050513	050611	AYACUCHO	LUCANAS	LUCANAS	2025-12-05 15:55:51.554787-05	2025-12-05 15:55:51.554787-05
513	050516	050612	AYACUCHO	LUCANAS	OCAÑA	2025-12-05 15:55:51.555294-05	2025-12-05 15:55:51.555294-05
514	050517	050613	AYACUCHO	LUCANAS	OTOCA	2025-12-05 15:55:51.555835-05	2025-12-05 15:55:51.555835-05
515	050529	050614	AYACUCHO	LUCANAS	SAISA	2025-12-05 15:55:51.555835-05	2025-12-05 15:55:51.555835-05
516	050532	050615	AYACUCHO	LUCANAS	SAN CRISTOBAL	2025-12-05 15:55:51.556353-05	2025-12-05 15:55:51.556353-05
517	050521	050616	AYACUCHO	LUCANAS	SAN JUAN	2025-12-05 15:55:51.556939-05	2025-12-05 15:55:51.556939-05
518	050522	050617	AYACUCHO	LUCANAS	SAN PEDRO	2025-12-05 15:55:51.557466-05	2025-12-05 15:55:51.557466-05
519	050531	050618	AYACUCHO	LUCANAS	SAN PEDRO DE PALCO	2025-12-05 15:55:51.55805-05	2025-12-05 15:55:51.55805-05
520	050520	050619	AYACUCHO	LUCANAS	SANCOS	2025-12-05 15:55:51.559082-05	2025-12-05 15:55:51.559082-05
521	050524	050620	AYACUCHO	LUCANAS	SANTA ANA DE HUAYCAHUACHO	2025-12-05 15:55:51.559597-05	2025-12-05 15:55:51.559597-05
522	050525	050621	AYACUCHO	LUCANAS	SANTA LUCIA	2025-12-05 15:55:51.560631-05	2025-12-05 15:55:51.560631-05
523	050601	050701	AYACUCHO	PARINACOCHAS	CORACORA	2025-12-05 15:55:51.561159-05	2025-12-05 15:55:51.561159-05
524	050605	050702	AYACUCHO	PARINACOCHAS	CHUMPI	2025-12-05 15:55:51.561772-05	2025-12-05 15:55:51.561772-05
525	050604	050703	AYACUCHO	PARINACOCHAS	CORONEL CASTAÑEDA	2025-12-05 15:55:51.562378-05	2025-12-05 15:55:51.562378-05
526	050608	050704	AYACUCHO	PARINACOCHAS	PACAPAUSA	2025-12-05 15:55:51.563016-05	2025-12-05 15:55:51.563016-05
527	050611	050705	AYACUCHO	PARINACOCHAS	PULLO	2025-12-05 15:55:51.563703-05	2025-12-05 15:55:51.563703-05
528	050612	050706	AYACUCHO	PARINACOCHAS	PUYUSCA	2025-12-05 15:55:51.564343-05	2025-12-05 15:55:51.564343-05
529	050615	050707	AYACUCHO	PARINACOCHAS	SAN FRANCISCO DE RAVACAYCO	2025-12-05 15:55:51.565057-05	2025-12-05 15:55:51.565057-05
530	050616	050708	AYACUCHO	PARINACOCHAS	UPAHUACHO	2025-12-05 15:55:51.565655-05	2025-12-05 15:55:51.565655-05
531	051001	050801	AYACUCHO	PAUCAR DEL SARA SARA	PAUSA	2025-12-05 15:55:51.566222-05	2025-12-05 15:55:51.566222-05
532	051002	050802	AYACUCHO	PAUCAR DEL SARA SARA	COLTA	2025-12-05 15:55:51.566744-05	2025-12-05 15:55:51.566744-05
533	051003	050803	AYACUCHO	PAUCAR DEL SARA SARA	CORCULLA	2025-12-05 15:55:51.56769-05	2025-12-05 15:55:51.56769-05
534	051004	050804	AYACUCHO	PAUCAR DEL SARA SARA	LAMPA	2025-12-05 15:55:51.568636-05	2025-12-05 15:55:51.568636-05
535	051005	050805	AYACUCHO	PAUCAR DEL SARA SARA	MARCABAMBA	2025-12-05 15:55:51.569148-05	2025-12-05 15:55:51.569148-05
536	051006	050806	AYACUCHO	PAUCAR DEL SARA SARA	OYOLO	2025-12-05 15:55:51.569789-05	2025-12-05 15:55:51.569789-05
537	051007	050807	AYACUCHO	PAUCAR DEL SARA SARA	PARARCA	2025-12-05 15:55:51.57034-05	2025-12-05 15:55:51.57034-05
538	051008	050808	AYACUCHO	PAUCAR DEL SARA SARA	SAN JAVIER DE ALPABAMBA	2025-12-05 15:55:51.570852-05	2025-12-05 15:55:51.570852-05
539	051009	050809	AYACUCHO	PAUCAR DEL SARA SARA	SAN JOSE DE USHUA	2025-12-05 15:55:51.570852-05	2025-12-05 15:55:51.570852-05
540	051010	050810	AYACUCHO	PAUCAR DEL SARA SARA	SARA SARA	2025-12-05 15:55:51.571403-05	2025-12-05 15:55:51.571403-05
541	051101	050901	AYACUCHO	SUCRE	QUEROBAMBA	2025-12-05 15:55:51.571942-05	2025-12-05 15:55:51.571942-05
542	051102	050902	AYACUCHO	SUCRE	BELEN	2025-12-05 15:55:51.572459-05	2025-12-05 15:55:51.572459-05
543	051103	050903	AYACUCHO	SUCRE	CHALCOS	2025-12-05 15:55:51.572459-05	2025-12-05 15:55:51.572459-05
544	051110	050904	AYACUCHO	SUCRE	CHILCAYOC	2025-12-05 15:55:51.573452-05	2025-12-05 15:55:51.573452-05
545	051109	050905	AYACUCHO	SUCRE	HUACAÑA	2025-12-05 15:55:51.573452-05	2025-12-05 15:55:51.573452-05
546	051111	050906	AYACUCHO	SUCRE	MORCOLLA	2025-12-05 15:55:51.574136-05	2025-12-05 15:55:51.574136-05
547	051105	050907	AYACUCHO	SUCRE	PAICO	2025-12-05 15:55:51.574662-05	2025-12-05 15:55:51.574662-05
548	051107	050908	AYACUCHO	SUCRE	SAN PEDRO DE LARCAY	2025-12-05 15:55:51.574662-05	2025-12-05 15:55:51.574662-05
549	051104	050909	AYACUCHO	SUCRE	SAN SALVADOR DE QUIJE	2025-12-05 15:55:51.575202-05	2025-12-05 15:55:51.575202-05
550	051106	050910	AYACUCHO	SUCRE	SANTIAGO DE PAUCARAY	2025-12-05 15:55:51.575763-05	2025-12-05 15:55:51.575763-05
551	051108	050911	AYACUCHO	SUCRE	SORAS	2025-12-05 15:55:51.576802-05	2025-12-05 15:55:51.576802-05
552	050701	051001	AYACUCHO	VICTOR FAJARDO	HUANCAPI	2025-12-05 15:55:51.577313-05	2025-12-05 15:55:51.577313-05
553	050702	051002	AYACUCHO	VICTOR FAJARDO	ALCAMENCA	2025-12-05 15:55:51.577837-05	2025-12-05 15:55:51.577837-05
554	050703	051003	AYACUCHO	VICTOR FAJARDO	APONGO	2025-12-05 15:55:51.579009-05	2025-12-05 15:55:51.579009-05
555	050715	051004	AYACUCHO	VICTOR FAJARDO	ASQUIPATA	2025-12-05 15:55:51.579528-05	2025-12-05 15:55:51.579528-05
556	050704	051005	AYACUCHO	VICTOR FAJARDO	CANARIA	2025-12-05 15:55:51.580043-05	2025-12-05 15:55:51.580043-05
557	050706	051006	AYACUCHO	VICTOR FAJARDO	CAYARA	2025-12-05 15:55:51.580563-05	2025-12-05 15:55:51.580563-05
558	050707	051007	AYACUCHO	VICTOR FAJARDO	COLCA	2025-12-05 15:55:51.581396-05	2025-12-05 15:55:51.581396-05
559	050709	051008	AYACUCHO	VICTOR FAJARDO	HUAMANQUIQUIA	2025-12-05 15:55:51.582048-05	2025-12-05 15:55:51.582048-05
560	050710	051009	AYACUCHO	VICTOR FAJARDO	HUANCARAYLLA	2025-12-05 15:55:51.582559-05	2025-12-05 15:55:51.582559-05
561	050708	051010	AYACUCHO	VICTOR FAJARDO	HUAYA	2025-12-05 15:55:51.5831-05	2025-12-05 15:55:51.5831-05
562	050713	051011	AYACUCHO	VICTOR FAJARDO	SARHUA	2025-12-05 15:55:51.583702-05	2025-12-05 15:55:51.583702-05
563	050714	051012	AYACUCHO	VICTOR FAJARDO	VILCANCHOS	2025-12-05 15:55:51.584435-05	2025-12-05 15:55:51.584435-05
564	050901	051101	AYACUCHO	VILCAS HUAMAN	VILCAS HUAMAN	2025-12-05 15:55:51.584435-05	2025-12-05 15:55:51.584435-05
565	050903	051102	AYACUCHO	VILCAS HUAMAN	ACCOMARCA	2025-12-05 15:55:51.584946-05	2025-12-05 15:55:51.584946-05
566	050904	051103	AYACUCHO	VILCAS HUAMAN	CARHUANCA	2025-12-05 15:55:51.585454-05	2025-12-05 15:55:51.585454-05
567	050905	051104	AYACUCHO	VILCAS HUAMAN	CONCEPCION	2025-12-05 15:55:51.585454-05	2025-12-05 15:55:51.585454-05
568	050906	051105	AYACUCHO	VILCAS HUAMAN	HUAMBALPA	2025-12-05 15:55:51.586097-05	2025-12-05 15:55:51.586097-05
569	050908	051106	AYACUCHO	VILCAS HUAMAN	INDEPENDENCIA	2025-12-05 15:55:51.586615-05	2025-12-05 15:55:51.586615-05
570	050907	051107	AYACUCHO	VILCAS HUAMAN	SAURAMA	2025-12-05 15:55:51.587213-05	2025-12-05 15:55:51.587213-05
571	050902	051108	AYACUCHO	VILCAS HUAMAN	VISCHONGO	2025-12-05 15:55:51.587746-05	2025-12-05 15:55:51.587746-05
572	060101	060101	CAJAMARCA	CAJAMARCA	CAJAMARCA	2025-12-05 15:55:51.587746-05	2025-12-05 15:55:51.587746-05
573	060102	060102	CAJAMARCA	CAJAMARCA	ASUNCION	2025-12-05 15:55:51.588867-05	2025-12-05 15:55:51.588867-05
574	060104	060103	CAJAMARCA	CAJAMARCA	CHETILLA	2025-12-05 15:55:51.588867-05	2025-12-05 15:55:51.588867-05
575	060103	060104	CAJAMARCA	CAJAMARCA	COSPAN	2025-12-05 15:55:51.589412-05	2025-12-05 15:55:51.589412-05
576	060105	060105	CAJAMARCA	CAJAMARCA	ENCAÑADA	2025-12-05 15:55:51.589956-05	2025-12-05 15:55:51.589956-05
577	060106	060106	CAJAMARCA	CAJAMARCA	JESUS	2025-12-05 15:55:51.590466-05	2025-12-05 15:55:51.590466-05
578	060108	060107	CAJAMARCA	CAJAMARCA	LLACANORA	2025-12-05 15:55:51.590466-05	2025-12-05 15:55:51.590466-05
579	060107	060108	CAJAMARCA	CAJAMARCA	LOS BAÑOS DEL INCA	2025-12-05 15:55:51.591003-05	2025-12-05 15:55:51.591003-05
580	060109	060109	CAJAMARCA	CAJAMARCA	MAGDALENA	2025-12-05 15:55:51.591625-05	2025-12-05 15:55:51.591625-05
581	060110	060110	CAJAMARCA	CAJAMARCA	MATARA	2025-12-05 15:55:51.591625-05	2025-12-05 15:55:51.591625-05
582	060111	060111	CAJAMARCA	CAJAMARCA	NAMORA	2025-12-05 15:55:51.592158-05	2025-12-05 15:55:51.592158-05
583	060112	060112	CAJAMARCA	CAJAMARCA	SAN JUAN	2025-12-05 15:55:51.593184-05	2025-12-05 15:55:51.593184-05
584	060201	060201	CAJAMARCA	CAJABAMBA	CAJABAMBA	2025-12-05 15:55:51.593696-05	2025-12-05 15:55:51.593696-05
585	060202	060202	CAJAMARCA	CAJABAMBA	CACHACHI	2025-12-05 15:55:51.594859-05	2025-12-05 15:55:51.594859-05
586	060203	060203	CAJAMARCA	CAJABAMBA	CONDEBAMBA	2025-12-05 15:55:51.595438-05	2025-12-05 15:55:51.595438-05
587	060205	060204	CAJAMARCA	CAJABAMBA	SITACOCHA	2025-12-05 15:55:51.595952-05	2025-12-05 15:55:51.595952-05
588	060301	060301	CAJAMARCA	CELENDIN	CELENDIN	2025-12-05 15:55:51.596479-05	2025-12-05 15:55:51.596479-05
589	060303	060302	CAJAMARCA	CELENDIN	CHUMUCH	2025-12-05 15:55:51.597027-05	2025-12-05 15:55:51.597027-05
590	060302	060303	CAJAMARCA	CELENDIN	CORTEGANA	2025-12-05 15:55:51.597861-05	2025-12-05 15:55:51.597861-05
591	060304	060304	CAJAMARCA	CELENDIN	HUASMIN	2025-12-05 15:55:51.598986-05	2025-12-05 15:55:51.598986-05
592	060305	060305	CAJAMARCA	CELENDIN	JORGE CHAVEZ	2025-12-05 15:55:51.599446-05	2025-12-05 15:55:51.599446-05
593	060306	060306	CAJAMARCA	CELENDIN	JOSE GALVEZ	2025-12-05 15:55:51.600531-05	2025-12-05 15:55:51.600531-05
594	060307	060307	CAJAMARCA	CELENDIN	MIGUEL IGLESIAS	2025-12-05 15:55:51.600531-05	2025-12-05 15:55:51.600531-05
595	060308	060308	CAJAMARCA	CELENDIN	OXAMARCA	2025-12-05 15:55:51.601108-05	2025-12-05 15:55:51.601108-05
596	060309	060309	CAJAMARCA	CELENDIN	SOROCHUCO	2025-12-05 15:55:51.601618-05	2025-12-05 15:55:51.601618-05
597	060310	060310	CAJAMARCA	CELENDIN	SUCRE	2025-12-05 15:55:51.601618-05	2025-12-05 15:55:51.601618-05
598	060311	060311	CAJAMARCA	CELENDIN	UTCO	2025-12-05 15:55:51.602254-05	2025-12-05 15:55:51.602254-05
599	060312	060312	CAJAMARCA	CELENDIN	LA LIBERTAD DE PALLAN	2025-12-05 15:55:51.602787-05	2025-12-05 15:55:51.602787-05
600	060601	060401	CAJAMARCA	CHOTA	CHOTA	2025-12-05 15:55:51.602787-05	2025-12-05 15:55:51.602787-05
601	060602	060402	CAJAMARCA	CHOTA	ANGUIA	2025-12-05 15:55:51.603352-05	2025-12-05 15:55:51.603352-05
602	060605	060403	CAJAMARCA	CHOTA	CHADIN	2025-12-05 15:55:51.603873-05	2025-12-05 15:55:51.603873-05
603	060606	060404	CAJAMARCA	CHOTA	CHIGUIRIP	2025-12-05 15:55:51.604719-05	2025-12-05 15:55:51.604719-05
604	060607	060405	CAJAMARCA	CHOTA	CHIMBAN	2025-12-05 15:55:51.605237-05	2025-12-05 15:55:51.605237-05
605	060618	060406	CAJAMARCA	CHOTA	CHOROPAMPA	2025-12-05 15:55:51.605237-05	2025-12-05 15:55:51.605237-05
606	060603	060407	CAJAMARCA	CHOTA	COCHABAMBA	2025-12-05 15:55:51.605749-05	2025-12-05 15:55:51.605749-05
607	060604	060408	CAJAMARCA	CHOTA	CONCHAN	2025-12-05 15:55:51.606352-05	2025-12-05 15:55:51.606352-05
608	060608	060409	CAJAMARCA	CHOTA	HUAMBOS	2025-12-05 15:55:51.606352-05	2025-12-05 15:55:51.606352-05
609	060609	060410	CAJAMARCA	CHOTA	LAJAS	2025-12-05 15:55:51.606929-05	2025-12-05 15:55:51.606929-05
610	060610	060411	CAJAMARCA	CHOTA	LLAMA	2025-12-05 15:55:51.606929-05	2025-12-05 15:55:51.606929-05
611	060611	060412	CAJAMARCA	CHOTA	MIRACOSTA	2025-12-05 15:55:51.607526-05	2025-12-05 15:55:51.607526-05
612	060612	060413	CAJAMARCA	CHOTA	PACCHA	2025-12-05 15:55:51.60805-05	2025-12-05 15:55:51.60805-05
613	060613	060414	CAJAMARCA	CHOTA	PION	2025-12-05 15:55:51.60805-05	2025-12-05 15:55:51.60805-05
614	060614	060415	CAJAMARCA	CHOTA	QUEROCOTO	2025-12-05 15:55:51.609134-05	2025-12-05 15:55:51.609134-05
615	060617	060416	CAJAMARCA	CHOTA	SAN JUAN DE LICUPIS	2025-12-05 15:55:51.609647-05	2025-12-05 15:55:51.609647-05
616	060615	060417	CAJAMARCA	CHOTA	TACABAMBA	2025-12-05 15:55:51.610694-05	2025-12-05 15:55:51.610694-05
617	060616	060418	CAJAMARCA	CHOTA	TOCMOCHE	2025-12-05 15:55:51.611216-05	2025-12-05 15:55:51.611216-05
618	060619	060419	CAJAMARCA	CHOTA	CHALAMARCA	2025-12-05 15:55:51.611741-05	2025-12-05 15:55:51.611741-05
619	060401	060501	CAJAMARCA	CONTUMAZA	CONTUMAZA	2025-12-05 15:55:51.612262-05	2025-12-05 15:55:51.612262-05
620	060403	060502	CAJAMARCA	CONTUMAZA	CHILETE	2025-12-05 15:55:51.612768-05	2025-12-05 15:55:51.612768-05
621	060406	060503	CAJAMARCA	CONTUMAZA	CUPISNIQUE	2025-12-05 15:55:51.613845-05	2025-12-05 15:55:51.613845-05
622	060404	060504	CAJAMARCA	CONTUMAZA	GUZMANGO	2025-12-05 15:55:51.614581-05	2025-12-05 15:55:51.614581-05
623	060405	060505	CAJAMARCA	CONTUMAZA	SAN BENITO	2025-12-05 15:55:51.615572-05	2025-12-05 15:55:51.615572-05
624	060409	060506	CAJAMARCA	CONTUMAZA	SANTA CRUZ DE TOLEDO	2025-12-05 15:55:51.616154-05	2025-12-05 15:55:51.616154-05
625	060407	060507	CAJAMARCA	CONTUMAZA	TANTARICA	2025-12-05 15:55:51.616675-05	2025-12-05 15:55:51.616675-05
626	060408	060508	CAJAMARCA	CONTUMAZA	YONAN	2025-12-05 15:55:51.617183-05	2025-12-05 15:55:51.617183-05
627	060501	060601	CAJAMARCA	CUTERVO	CUTERVO	2025-12-05 15:55:51.617698-05	2025-12-05 15:55:51.617698-05
628	060502	060602	CAJAMARCA	CUTERVO	CALLAYUC	2025-12-05 15:55:51.618248-05	2025-12-05 15:55:51.618248-05
629	060504	060603	CAJAMARCA	CUTERVO	CHOROS	2025-12-05 15:55:51.618248-05	2025-12-05 15:55:51.618248-05
630	060503	060604	CAJAMARCA	CUTERVO	CUJILLO	2025-12-05 15:55:51.618769-05	2025-12-05 15:55:51.618769-05
631	060505	060605	CAJAMARCA	CUTERVO	LA RAMADA	2025-12-05 15:55:51.619292-05	2025-12-05 15:55:51.619292-05
632	060506	060606	CAJAMARCA	CUTERVO	PIMPINGOS	2025-12-05 15:55:51.619845-05	2025-12-05 15:55:51.619845-05
633	060507	060607	CAJAMARCA	CUTERVO	QUEROCOTILLO	2025-12-05 15:55:51.620429-05	2025-12-05 15:55:51.620429-05
634	060508	060608	CAJAMARCA	CUTERVO	SAN ANDRES DE CUTERVO	2025-12-05 15:55:51.621469-05	2025-12-05 15:55:51.621469-05
635	060509	060609	CAJAMARCA	CUTERVO	SAN JUAN DE CUTERVO	2025-12-05 15:55:51.622983-05	2025-12-05 15:55:51.622983-05
636	060510	060610	CAJAMARCA	CUTERVO	SAN LUIS DE LUCMA	2025-12-05 15:55:51.624584-05	2025-12-05 15:55:51.624584-05
637	060511	060611	CAJAMARCA	CUTERVO	SANTA CRUZ	2025-12-05 15:55:51.625091-05	2025-12-05 15:55:51.625091-05
638	060512	060612	CAJAMARCA	CUTERVO	SANTO DOMINGO DE LA CAPILLA	2025-12-05 15:55:51.625605-05	2025-12-05 15:55:51.625605-05
639	060513	060613	CAJAMARCA	CUTERVO	SANTO TOMAS	2025-12-05 15:55:51.626661-05	2025-12-05 15:55:51.626661-05
640	060514	060614	CAJAMARCA	CUTERVO	SOCOTA	2025-12-05 15:55:51.62723-05	2025-12-05 15:55:51.62723-05
641	060515	060615	CAJAMARCA	CUTERVO	TORIBIO CASANOVA	2025-12-05 15:55:51.627784-05	2025-12-05 15:55:51.627784-05
642	060701	060701	CAJAMARCA	HUALGAYOC	BAMBAMARCA	2025-12-05 15:55:51.628818-05	2025-12-05 15:55:51.628818-05
643	060702	060702	CAJAMARCA	HUALGAYOC	CHUGUR	2025-12-05 15:55:51.629331-05	2025-12-05 15:55:51.629331-05
644	060703	060703	CAJAMARCA	HUALGAYOC	HUALGAYOC	2025-12-05 15:55:51.629842-05	2025-12-05 15:55:51.629842-05
645	060801	060801	CAJAMARCA	JAEN	JAEN	2025-12-05 15:55:51.63036-05	2025-12-05 15:55:51.63036-05
646	060802	060802	CAJAMARCA	JAEN	BELLAVISTA	2025-12-05 15:55:51.631379-05	2025-12-05 15:55:51.631379-05
647	060804	060803	CAJAMARCA	JAEN	CHONTALI	2025-12-05 15:55:51.631903-05	2025-12-05 15:55:51.631903-05
648	060803	060804	CAJAMARCA	JAEN	COLASAY	2025-12-05 15:55:51.633023-05	2025-12-05 15:55:51.633023-05
649	060812	060805	CAJAMARCA	JAEN	HUABAL	2025-12-05 15:55:51.63414-05	2025-12-05 15:55:51.63414-05
650	060811	060806	CAJAMARCA	JAEN	LAS PIRIAS	2025-12-05 15:55:51.63466-05	2025-12-05 15:55:51.63466-05
651	060805	060807	CAJAMARCA	JAEN	POMAHUACA	2025-12-05 15:55:51.635174-05	2025-12-05 15:55:51.635174-05
652	060806	060808	CAJAMARCA	JAEN	PUCARA	2025-12-05 15:55:51.635784-05	2025-12-05 15:55:51.635784-05
653	060807	060809	CAJAMARCA	JAEN	SALLIQUE	2025-12-05 15:55:51.635784-05	2025-12-05 15:55:51.635784-05
654	060808	060810	CAJAMARCA	JAEN	SAN FELIPE	2025-12-05 15:55:51.636296-05	2025-12-05 15:55:51.636296-05
655	060809	060811	CAJAMARCA	JAEN	SAN JOSE DEL ALTO	2025-12-05 15:55:51.636837-05	2025-12-05 15:55:51.636837-05
656	060810	060812	CAJAMARCA	JAEN	SANTA ROSA	2025-12-05 15:55:51.637403-05	2025-12-05 15:55:51.637403-05
657	061101	060901	CAJAMARCA	SAN IGNACIO	SAN IGNACIO	2025-12-05 15:55:51.637941-05	2025-12-05 15:55:51.637941-05
658	061102	060902	CAJAMARCA	SAN IGNACIO	CHIRINOS	2025-12-05 15:55:51.63845-05	2025-12-05 15:55:51.63845-05
659	061103	060903	CAJAMARCA	SAN IGNACIO	HUARANGO	2025-12-05 15:55:51.638964-05	2025-12-05 15:55:51.638964-05
660	061105	060904	CAJAMARCA	SAN IGNACIO	LA COIPA	2025-12-05 15:55:51.639563-05	2025-12-05 15:55:51.639563-05
661	061104	060905	CAJAMARCA	SAN IGNACIO	NAMBALLE	2025-12-05 15:55:51.639563-05	2025-12-05 15:55:51.639563-05
662	061106	060906	CAJAMARCA	SAN IGNACIO	SAN JOSE DE LOURDES	2025-12-05 15:55:51.640137-05	2025-12-05 15:55:51.640137-05
663	061107	060907	CAJAMARCA	SAN IGNACIO	TABACONAS	2025-12-05 15:55:51.640722-05	2025-12-05 15:55:51.640722-05
664	061201	061001	CAJAMARCA	SAN MARCOS	PEDRO GALVEZ	2025-12-05 15:55:51.640722-05	2025-12-05 15:55:51.640722-05
665	061207	061002	CAJAMARCA	SAN MARCOS	CHANCAY	2025-12-05 15:55:51.641264-05	2025-12-05 15:55:51.641264-05
666	061205	061003	CAJAMARCA	SAN MARCOS	EDUARDO VILLANUEVA	2025-12-05 15:55:51.641264-05	2025-12-05 15:55:51.641264-05
667	061203	061004	CAJAMARCA	SAN MARCOS	GREGORIO PITA	2025-12-05 15:55:51.641875-05	2025-12-05 15:55:51.641875-05
668	061202	061005	CAJAMARCA	SAN MARCOS	ICHOCAN	2025-12-05 15:55:51.642433-05	2025-12-05 15:55:51.642433-05
669	061204	061006	CAJAMARCA	SAN MARCOS	JOSE MANUEL QUIROZ	2025-12-05 15:55:51.643626-05	2025-12-05 15:55:51.643626-05
670	061206	061007	CAJAMARCA	SAN MARCOS	JOSE SABOGAL	2025-12-05 15:55:51.644165-05	2025-12-05 15:55:51.644165-05
671	061001	061101	CAJAMARCA	SAN MIGUEL	SAN MIGUEL	2025-12-05 15:55:51.644695-05	2025-12-05 15:55:51.644695-05
672	061013	061102	CAJAMARCA	SAN MIGUEL	BOLIVAR	2025-12-05 15:55:51.645237-05	2025-12-05 15:55:51.645237-05
673	061002	061103	CAJAMARCA	SAN MIGUEL	CALQUIS	2025-12-05 15:55:51.645831-05	2025-12-05 15:55:51.645831-05
674	061012	061104	CAJAMARCA	SAN MIGUEL	CATILLUC	2025-12-05 15:55:51.646461-05	2025-12-05 15:55:51.646461-05
675	061009	061105	CAJAMARCA	SAN MIGUEL	EL PRADO	2025-12-05 15:55:51.646988-05	2025-12-05 15:55:51.646988-05
676	061003	061106	CAJAMARCA	SAN MIGUEL	LA FLORIDA	2025-12-05 15:55:51.647711-05	2025-12-05 15:55:51.647711-05
677	061004	061107	CAJAMARCA	SAN MIGUEL	LLAPA	2025-12-05 15:55:51.648414-05	2025-12-05 15:55:51.648414-05
678	061005	061108	CAJAMARCA	SAN MIGUEL	NANCHOC	2025-12-05 15:55:51.649073-05	2025-12-05 15:55:51.649073-05
679	061006	061109	CAJAMARCA	SAN MIGUEL	NIEPOS	2025-12-05 15:55:51.649901-05	2025-12-05 15:55:51.649901-05
680	061007	061110	CAJAMARCA	SAN MIGUEL	SAN GREGORIO	2025-12-05 15:55:51.65055-05	2025-12-05 15:55:51.65055-05
681	061008	061111	CAJAMARCA	SAN MIGUEL	SAN SILVESTRE DE COCHAN	2025-12-05 15:55:51.651075-05	2025-12-05 15:55:51.651075-05
682	061011	061112	CAJAMARCA	SAN MIGUEL	TONGOD	2025-12-05 15:55:51.651075-05	2025-12-05 15:55:51.651075-05
683	061010	061113	CAJAMARCA	SAN MIGUEL	UNION AGUA BLANCA	2025-12-05 15:55:51.6516-05	2025-12-05 15:55:51.6516-05
684	061301	061201	CAJAMARCA	SAN PABLO	SAN PABLO	2025-12-05 15:55:51.652154-05	2025-12-05 15:55:51.652154-05
685	061302	061202	CAJAMARCA	SAN PABLO	SAN BERNARDINO	2025-12-05 15:55:51.652689-05	2025-12-05 15:55:51.652689-05
686	061303	061203	CAJAMARCA	SAN PABLO	SAN LUIS	2025-12-05 15:55:51.652689-05	2025-12-05 15:55:51.652689-05
687	061304	061204	CAJAMARCA	SAN PABLO	TUMBADEN	2025-12-05 15:55:51.653214-05	2025-12-05 15:55:51.653214-05
688	060901	061301	CAJAMARCA	SANTA CRUZ	SANTA CRUZ	2025-12-05 15:55:51.653822-05	2025-12-05 15:55:51.653822-05
689	060910	061302	CAJAMARCA	SANTA CRUZ	ANDABAMBA	2025-12-05 15:55:51.654376-05	2025-12-05 15:55:51.654376-05
690	060902	061303	CAJAMARCA	SANTA CRUZ	CATACHE	2025-12-05 15:55:51.655008-05	2025-12-05 15:55:51.655008-05
691	060903	061304	CAJAMARCA	SANTA CRUZ	CHANCAYBAÑOS	2025-12-05 15:55:51.655008-05	2025-12-05 15:55:51.655008-05
692	060904	061305	CAJAMARCA	SANTA CRUZ	LA ESPERANZA	2025-12-05 15:55:51.655008-05	2025-12-05 15:55:51.655008-05
693	060905	061306	CAJAMARCA	SANTA CRUZ	NINABAMBA	2025-12-05 15:55:51.655796-05	2025-12-05 15:55:51.655796-05
694	060906	061307	CAJAMARCA	SANTA CRUZ	PULAN	2025-12-05 15:55:51.656366-05	2025-12-05 15:55:51.656366-05
695	060911	061308	CAJAMARCA	SANTA CRUZ	SAUCEPAMPA	2025-12-05 15:55:51.656889-05	2025-12-05 15:55:51.656889-05
696	060907	061309	CAJAMARCA	SANTA CRUZ	SEXI	2025-12-05 15:55:51.656889-05	2025-12-05 15:55:51.656889-05
697	060908	061310	CAJAMARCA	SANTA CRUZ	UTICYACU	2025-12-05 15:55:51.657411-05	2025-12-05 15:55:51.657411-05
698	060909	061311	CAJAMARCA	SANTA CRUZ	YAUYUCAN	2025-12-05 15:55:51.657933-05	2025-12-05 15:55:51.657933-05
699	240101	070101	CALLAO	CALLAO	CALLAO	2025-12-05 15:55:51.657933-05	2025-12-05 15:55:51.657933-05
700	240102	070102	CALLAO	CALLAO	BELLAVISTA	2025-12-05 15:55:51.658489-05	2025-12-05 15:55:51.658489-05
701	240104	070103	CALLAO	CALLAO	CARMEN DE LA LEGUA REYNOSO	2025-12-05 15:55:51.659004-05	2025-12-05 15:55:51.659004-05
702	240105	070104	CALLAO	CALLAO	LA PERLA	2025-12-05 15:55:51.66011-05	2025-12-05 15:55:51.66011-05
703	240103	070105	CALLAO	CALLAO	LA PUNTA	2025-12-05 15:55:51.660733-05	2025-12-05 15:55:51.660733-05
704	240106	070106	CALLAO	CALLAO	VENTANILLA	2025-12-05 15:55:51.661263-05	2025-12-05 15:55:51.661263-05
705	240107	070107	CALLAO	CALLAO	MI PERU	2025-12-05 15:55:51.66186-05	2025-12-05 15:55:51.66186-05
706	070101	080101	CUSCO	CUSCO	CUSCO	2025-12-05 15:55:51.662476-05	2025-12-05 15:55:51.662476-05
707	070102	080102	CUSCO	CUSCO	CCORCA	2025-12-05 15:55:51.663002-05	2025-12-05 15:55:51.663002-05
708	070103	080103	CUSCO	CUSCO	POROY	2025-12-05 15:55:51.663002-05	2025-12-05 15:55:51.663002-05
709	070104	080104	CUSCO	CUSCO	SAN JERONIMO	2025-12-05 15:55:51.664017-05	2025-12-05 15:55:51.664017-05
710	070105	080105	CUSCO	CUSCO	SAN SEBASTIAN	2025-12-05 15:55:51.664548-05	2025-12-05 15:55:51.664548-05
711	070106	080106	CUSCO	CUSCO	SANTIAGO	2025-12-05 15:55:51.665318-05	2025-12-05 15:55:51.665318-05
712	070107	080107	CUSCO	CUSCO	SAYLLA	2025-12-05 15:55:51.665876-05	2025-12-05 15:55:51.665876-05
713	070108	080108	CUSCO	CUSCO	WANCHAQ	2025-12-05 15:55:51.666491-05	2025-12-05 15:55:51.666491-05
714	070201	080201	CUSCO	ACOMAYO	ACOMAYO	2025-12-05 15:55:51.667023-05	2025-12-05 15:55:51.667023-05
715	070202	080202	CUSCO	ACOMAYO	ACOPIA	2025-12-05 15:55:51.667023-05	2025-12-05 15:55:51.667023-05
716	070203	080203	CUSCO	ACOMAYO	ACOS	2025-12-05 15:55:51.667564-05	2025-12-05 15:55:51.667564-05
717	070207	080204	CUSCO	ACOMAYO	MOSOC LLACTA	2025-12-05 15:55:51.668124-05	2025-12-05 15:55:51.668124-05
718	070204	080205	CUSCO	ACOMAYO	POMACANCHI	2025-12-05 15:55:51.668653-05	2025-12-05 15:55:51.668653-05
719	070205	080206	CUSCO	ACOMAYO	RONDOCAN	2025-12-05 15:55:51.668653-05	2025-12-05 15:55:51.668653-05
720	070206	080207	CUSCO	ACOMAYO	SANGARARA	2025-12-05 15:55:51.669178-05	2025-12-05 15:55:51.669178-05
721	070301	080301	CUSCO	ANTA	ANTA	2025-12-05 15:55:51.669178-05	2025-12-05 15:55:51.669178-05
722	070309	080302	CUSCO	ANTA	ANCAHUASI	2025-12-05 15:55:51.670394-05	2025-12-05 15:55:51.670394-05
723	070308	080303	CUSCO	ANTA	CACHIMAYO	2025-12-05 15:55:51.670757-05	2025-12-05 15:55:51.670757-05
724	070302	080304	CUSCO	ANTA	CHINCHAYPUJIO	2025-12-05 15:55:51.671467-05	2025-12-05 15:55:51.671467-05
725	070303	080305	CUSCO	ANTA	HUAROCONDO	2025-12-05 15:55:51.671467-05	2025-12-05 15:55:51.671467-05
726	070304	080306	CUSCO	ANTA	LIMATAMBO	2025-12-05 15:55:51.671988-05	2025-12-05 15:55:51.671988-05
727	070305	080307	CUSCO	ANTA	MOLLEPATA	2025-12-05 15:55:51.672524-05	2025-12-05 15:55:51.672524-05
728	070306	080308	CUSCO	ANTA	PUCYURA	2025-12-05 15:55:51.673043-05	2025-12-05 15:55:51.673043-05
729	070307	080309	CUSCO	ANTA	ZURITE	2025-12-05 15:55:51.673043-05	2025-12-05 15:55:51.673043-05
730	070401	080401	CUSCO	CALCA	CALCA	2025-12-05 15:55:51.67366-05	2025-12-05 15:55:51.67366-05
731	070402	080402	CUSCO	CALCA	COYA	2025-12-05 15:55:51.67366-05	2025-12-05 15:55:51.67366-05
732	070403	080403	CUSCO	CALCA	LAMAY	2025-12-05 15:55:51.674222-05	2025-12-05 15:55:51.674222-05
733	070404	080404	CUSCO	CALCA	LARES	2025-12-05 15:55:51.674776-05	2025-12-05 15:55:51.674776-05
734	070405	080405	CUSCO	CALCA	PISAC	2025-12-05 15:55:51.675094-05	2025-12-05 15:55:51.675094-05
735	070406	080406	CUSCO	CALCA	SAN SALVADOR	2025-12-05 15:55:51.675615-05	2025-12-05 15:55:51.675615-05
736	070407	080407	CUSCO	CALCA	TARAY	2025-12-05 15:55:51.676242-05	2025-12-05 15:55:51.676242-05
737	070408	080408	CUSCO	CALCA	YANATILE	2025-12-05 15:55:51.676887-05	2025-12-05 15:55:51.676887-05
738	070501	080501	CUSCO	CANAS	YANAOCA	2025-12-05 15:55:51.677967-05	2025-12-05 15:55:51.677967-05
739	070502	080502	CUSCO	CANAS	CHECCA	2025-12-05 15:55:51.678487-05	2025-12-05 15:55:51.678487-05
740	070503	080503	CUSCO	CANAS	KUNTURKANKI	2025-12-05 15:55:51.679015-05	2025-12-05 15:55:51.679015-05
741	070504	080504	CUSCO	CANAS	LANGUI	2025-12-05 15:55:51.679598-05	2025-12-05 15:55:51.679598-05
742	070505	080505	CUSCO	CANAS	LAYO	2025-12-05 15:55:51.680119-05	2025-12-05 15:55:51.680119-05
743	070506	080506	CUSCO	CANAS	PAMPAMARCA	2025-12-05 15:55:51.68065-05	2025-12-05 15:55:51.68065-05
744	070507	080507	CUSCO	CANAS	QUEHUE	2025-12-05 15:55:51.681792-05	2025-12-05 15:55:51.681792-05
745	070508	080508	CUSCO	CANAS	TUPAC AMARU	2025-12-05 15:55:51.682321-05	2025-12-05 15:55:51.682321-05
746	070601	080601	CUSCO	CANCHIS	SICUANI	2025-12-05 15:55:51.68284-05	2025-12-05 15:55:51.68284-05
747	070603	080602	CUSCO	CANCHIS	CHECACUPE	2025-12-05 15:55:51.683392-05	2025-12-05 15:55:51.683392-05
748	070602	080603	CUSCO	CANCHIS	COMBAPATA	2025-12-05 15:55:51.683392-05	2025-12-05 15:55:51.683392-05
749	070604	080604	CUSCO	CANCHIS	MARANGANI	2025-12-05 15:55:51.683966-05	2025-12-05 15:55:51.683966-05
750	070605	080605	CUSCO	CANCHIS	PITUMARCA	2025-12-05 15:55:51.684526-05	2025-12-05 15:55:51.684526-05
751	070606	080606	CUSCO	CANCHIS	SAN PABLO	2025-12-05 15:55:51.684887-05	2025-12-05 15:55:51.684887-05
752	070607	080607	CUSCO	CANCHIS	SAN PEDRO	2025-12-05 15:55:51.685316-05	2025-12-05 15:55:51.685316-05
753	070608	080608	CUSCO	CANCHIS	TINTA	2025-12-05 15:55:51.685846-05	2025-12-05 15:55:51.685846-05
754	070701	080701	CUSCO	CHUMBIVILCAS	SANTO TOMAS	2025-12-05 15:55:51.68638-05	2025-12-05 15:55:51.68638-05
755	070702	080702	CUSCO	CHUMBIVILCAS	CAPACMARCA	2025-12-05 15:55:51.68695-05	2025-12-05 15:55:51.68695-05
756	070704	080703	CUSCO	CHUMBIVILCAS	CHAMACA	2025-12-05 15:55:51.687486-05	2025-12-05 15:55:51.687486-05
757	070703	080704	CUSCO	CHUMBIVILCAS	COLQUEMARCA	2025-12-05 15:55:51.687486-05	2025-12-05 15:55:51.687486-05
758	070705	080705	CUSCO	CHUMBIVILCAS	LIVITACA	2025-12-05 15:55:51.688025-05	2025-12-05 15:55:51.688025-05
759	070706	080706	CUSCO	CHUMBIVILCAS	LLUSCO	2025-12-05 15:55:51.688025-05	2025-12-05 15:55:51.688025-05
760	070707	080707	CUSCO	CHUMBIVILCAS	QUIÑOTA	2025-12-05 15:55:51.688664-05	2025-12-05 15:55:51.688664-05
761	070708	080708	CUSCO	CHUMBIVILCAS	VELILLE	2025-12-05 15:55:51.689177-05	2025-12-05 15:55:51.689177-05
762	070801	080801	CUSCO	ESPINAR	ESPINAR	2025-12-05 15:55:51.689177-05	2025-12-05 15:55:51.689177-05
763	070802	080802	CUSCO	ESPINAR	CONDOROMA	2025-12-05 15:55:51.690364-05	2025-12-05 15:55:51.690364-05
764	070803	080803	CUSCO	ESPINAR	COPORAQUE	2025-12-05 15:55:51.690684-05	2025-12-05 15:55:51.690684-05
765	070804	080804	CUSCO	ESPINAR	OCORURO	2025-12-05 15:55:51.690684-05	2025-12-05 15:55:51.690684-05
766	070805	080805	CUSCO	ESPINAR	PALLPATA	2025-12-05 15:55:51.691308-05	2025-12-05 15:55:51.691308-05
767	070806	080806	CUSCO	ESPINAR	PICHIGUA	2025-12-05 15:55:51.691832-05	2025-12-05 15:55:51.691832-05
768	070807	080807	CUSCO	ESPINAR	SUYCKUTAMBO	2025-12-05 15:55:51.692423-05	2025-12-05 15:55:51.692423-05
769	070808	080808	CUSCO	ESPINAR	ALTO PICHIGUA	2025-12-05 15:55:51.692943-05	2025-12-05 15:55:51.692943-05
770	070901	080901	CUSCO	LA CONVENCION	SANTA ANA	2025-12-05 15:55:51.693474-05	2025-12-05 15:55:51.693474-05
771	070902	080902	CUSCO	LA CONVENCION	ECHARATE	2025-12-05 15:55:51.694209-05	2025-12-05 15:55:51.694209-05
772	070903	080903	CUSCO	LA CONVENCION	HUAYOPATA	2025-12-05 15:55:51.694909-05	2025-12-05 15:55:51.694909-05
773	070904	080904	CUSCO	LA CONVENCION	MARANURA	2025-12-05 15:55:51.695617-05	2025-12-05 15:55:51.695617-05
774	070905	080905	CUSCO	LA CONVENCION	OCOBAMBA	2025-12-05 15:55:51.696235-05	2025-12-05 15:55:51.696235-05
775	070908	080906	CUSCO	LA CONVENCION	QUELLOUNO	2025-12-05 15:55:51.697252-05	2025-12-05 15:55:51.697252-05
776	070909	080907	CUSCO	LA CONVENCION	KIMBIRI	2025-12-05 15:55:51.698433-05	2025-12-05 15:55:51.698433-05
777	070906	080908	CUSCO	LA CONVENCION	SANTA TERESA	2025-12-05 15:55:51.699059-05	2025-12-05 15:55:51.699059-05
778	070907	080909	CUSCO	LA CONVENCION	VILCABAMBA	2025-12-05 15:55:51.699059-05	2025-12-05 15:55:51.699059-05
779	070910	080910	CUSCO	LA CONVENCION	PICHARI	2025-12-05 15:55:51.699637-05	2025-12-05 15:55:51.699637-05
780	070911	080911	CUSCO	LA CONVENCION	INKAWASI	2025-12-05 15:55:51.700189-05	2025-12-05 15:55:51.700189-05
781	070912	080912	CUSCO	LA CONVENCION	VILLA VIRGEN	2025-12-05 15:55:51.700721-05	2025-12-05 15:55:51.700721-05
782	070913	080913	CUSCO	LA CONVENCION	VILLA KINTIARINA	2025-12-05 15:55:51.701263-05	2025-12-05 15:55:51.701263-05
783	070915	080914	CUSCO	LA CONVENCION	MEGANTONI	2025-12-05 15:55:51.701828-05	2025-12-05 15:55:51.701828-05
784	070916	080915	CUSCO	LA CONVENCION	KUMPIRUSHIATO	2025-12-05 15:55:51.701828-05	2025-12-05 15:55:51.701828-05
785	070917	080916	CUSCO	LA CONVENCION	CIELO PUNCO	2025-12-05 15:55:51.702345-05	2025-12-05 15:55:51.702345-05
786	070918	080917	CUSCO	LA CONVENCION	MANITEA	2025-12-05 15:55:51.702905-05	2025-12-05 15:55:51.702905-05
787	070919	080918	CUSCO	LA CONVENCION	UNION ASHÁNINKA	2025-12-05 15:55:51.703441-05	2025-12-05 15:55:51.703441-05
788	071001	081001	CUSCO	PARURO	PARURO	2025-12-05 15:55:51.703987-05	2025-12-05 15:55:51.703987-05
789	071002	081002	CUSCO	PARURO	ACCHA	2025-12-05 15:55:51.703987-05	2025-12-05 15:55:51.703987-05
790	071003	081003	CUSCO	PARURO	CCAPI	2025-12-05 15:55:51.704502-05	2025-12-05 15:55:51.704502-05
791	071004	081004	CUSCO	PARURO	COLCHA	2025-12-05 15:55:51.705048-05	2025-12-05 15:55:51.705048-05
792	071005	081005	CUSCO	PARURO	HUANOQUITE	2025-12-05 15:55:51.705048-05	2025-12-05 15:55:51.705048-05
793	071006	081006	CUSCO	PARURO	OMACHA	2025-12-05 15:55:51.705665-05	2025-12-05 15:55:51.705665-05
794	071008	081007	CUSCO	PARURO	PACCARITAMBO	2025-12-05 15:55:51.706288-05	2025-12-05 15:55:51.706288-05
795	071009	081008	CUSCO	PARURO	PILLPINTO	2025-12-05 15:55:51.706806-05	2025-12-05 15:55:51.706806-05
796	071007	081009	CUSCO	PARURO	YAURISQUE	2025-12-05 15:55:51.707328-05	2025-12-05 15:55:51.707328-05
797	071101	081101	CUSCO	PAUCARTAMBO	PAUCARTAMBO	2025-12-05 15:55:51.707855-05	2025-12-05 15:55:51.707855-05
798	071102	081102	CUSCO	PAUCARTAMBO	CAICAY	2025-12-05 15:55:51.707855-05	2025-12-05 15:55:51.707855-05
799	071104	081103	CUSCO	PAUCARTAMBO	CHALLABAMBA	2025-12-05 15:55:51.708432-05	2025-12-05 15:55:51.708432-05
800	071103	081104	CUSCO	PAUCARTAMBO	COLQUEPATA	2025-12-05 15:55:51.709491-05	2025-12-05 15:55:51.709491-05
801	071106	081105	CUSCO	PAUCARTAMBO	HUANCARANI	2025-12-05 15:55:51.710019-05	2025-12-05 15:55:51.710019-05
802	071105	081106	CUSCO	PAUCARTAMBO	KOSÑIPATA	2025-12-05 15:55:51.710566-05	2025-12-05 15:55:51.710566-05
803	071201	081201	CUSCO	QUISPICANCHI	URCOS	2025-12-05 15:55:51.711093-05	2025-12-05 15:55:51.711093-05
804	071202	081202	CUSCO	QUISPICANCHI	ANDAHUAYLILLAS	2025-12-05 15:55:51.712212-05	2025-12-05 15:55:51.712212-05
805	071203	081203	CUSCO	QUISPICANCHI	CAMANTI	2025-12-05 15:55:51.712733-05	2025-12-05 15:55:51.712733-05
806	071204	081204	CUSCO	QUISPICANCHI	CCARHUAYO	2025-12-05 15:55:51.712733-05	2025-12-05 15:55:51.712733-05
807	071205	081205	CUSCO	QUISPICANCHI	CCATCA	2025-12-05 15:55:51.713778-05	2025-12-05 15:55:51.713778-05
808	071206	081206	CUSCO	QUISPICANCHI	CUSIPATA	2025-12-05 15:55:51.714553-05	2025-12-05 15:55:51.714553-05
809	071207	081207	CUSCO	QUISPICANCHI	HUARO	2025-12-05 15:55:51.714903-05	2025-12-05 15:55:51.714903-05
810	071208	081208	CUSCO	QUISPICANCHI	LUCRE	2025-12-05 15:55:51.715543-05	2025-12-05 15:55:51.715543-05
811	071209	081209	CUSCO	QUISPICANCHI	MARCAPATA	2025-12-05 15:55:51.716067-05	2025-12-05 15:55:51.716067-05
812	071210	081210	CUSCO	QUISPICANCHI	OCONGATE	2025-12-05 15:55:51.716851-05	2025-12-05 15:55:51.716851-05
813	071211	081211	CUSCO	QUISPICANCHI	OROPESA	2025-12-05 15:55:51.717484-05	2025-12-05 15:55:51.717484-05
814	071212	081212	CUSCO	QUISPICANCHI	QUIQUIJANA	2025-12-05 15:55:51.717484-05	2025-12-05 15:55:51.717484-05
815	071301	081301	CUSCO	URUBAMBA	URUBAMBA	2025-12-05 15:55:51.718052-05	2025-12-05 15:55:51.718052-05
816	071302	081302	CUSCO	URUBAMBA	CHINCHERO	2025-12-05 15:55:51.718584-05	2025-12-05 15:55:51.718584-05
817	071303	081303	CUSCO	URUBAMBA	HUAYLLABAMBA	2025-12-05 15:55:51.719107-05	2025-12-05 15:55:51.719107-05
818	071304	081304	CUSCO	URUBAMBA	MACHUPICCHU	2025-12-05 15:55:51.719107-05	2025-12-05 15:55:51.719107-05
819	071305	081305	CUSCO	URUBAMBA	MARAS	2025-12-05 15:55:51.719622-05	2025-12-05 15:55:51.719622-05
820	071306	081306	CUSCO	URUBAMBA	OLLANTAYTAMBO	2025-12-05 15:55:51.719622-05	2025-12-05 15:55:51.719622-05
821	071307	081307	CUSCO	URUBAMBA	YUCAY	2025-12-05 15:55:51.720242-05	2025-12-05 15:55:51.720242-05
822	080101	090101	HUANCAVELICA	HUANCAVELICA	HUANCAVELICA	2025-12-05 15:55:51.720779-05	2025-12-05 15:55:51.720779-05
823	080102	090102	HUANCAVELICA	HUANCAVELICA	ACOBAMBILLA	2025-12-05 15:55:51.721343-05	2025-12-05 15:55:51.721343-05
824	080103	090103	HUANCAVELICA	HUANCAVELICA	ACORIA	2025-12-05 15:55:51.721922-05	2025-12-05 15:55:51.721922-05
825	080104	090104	HUANCAVELICA	HUANCAVELICA	CONAYCA	2025-12-05 15:55:51.722437-05	2025-12-05 15:55:51.722437-05
826	080105	090105	HUANCAVELICA	HUANCAVELICA	CUENCA	2025-12-05 15:55:51.723019-05	2025-12-05 15:55:51.723019-05
827	080106	090106	HUANCAVELICA	HUANCAVELICA	HUACHOCOLPA	2025-12-05 15:55:51.723019-05	2025-12-05 15:55:51.723019-05
828	080108	090107	HUANCAVELICA	HUANCAVELICA	HUAYLLAHUARA	2025-12-05 15:55:51.723587-05	2025-12-05 15:55:51.723587-05
829	080109	090108	HUANCAVELICA	HUANCAVELICA	IZCUCHACA	2025-12-05 15:55:51.724154-05	2025-12-05 15:55:51.724154-05
830	080110	090109	HUANCAVELICA	HUANCAVELICA	LARIA	2025-12-05 15:55:51.724154-05	2025-12-05 15:55:51.724154-05
831	080111	090110	HUANCAVELICA	HUANCAVELICA	MANTA	2025-12-05 15:55:51.724678-05	2025-12-05 15:55:51.724678-05
832	080112	090111	HUANCAVELICA	HUANCAVELICA	MARISCAL CACERES	2025-12-05 15:55:51.725246-05	2025-12-05 15:55:51.725246-05
833	080113	090112	HUANCAVELICA	HUANCAVELICA	MOYA	2025-12-05 15:55:51.725791-05	2025-12-05 15:55:51.725791-05
834	080114	090113	HUANCAVELICA	HUANCAVELICA	NUEVO OCCORO	2025-12-05 15:55:51.726309-05	2025-12-05 15:55:51.726309-05
835	080115	090114	HUANCAVELICA	HUANCAVELICA	PALCA	2025-12-05 15:55:51.72686-05	2025-12-05 15:55:51.72686-05
836	080116	090115	HUANCAVELICA	HUANCAVELICA	PILCHACA	2025-12-05 15:55:51.727942-05	2025-12-05 15:55:51.727942-05
837	080117	090116	HUANCAVELICA	HUANCAVELICA	VILCA	2025-12-05 15:55:51.728493-05	2025-12-05 15:55:51.728493-05
838	080118	090117	HUANCAVELICA	HUANCAVELICA	YAULI	2025-12-05 15:55:51.729134-05	2025-12-05 15:55:51.729134-05
839	080119	090118	HUANCAVELICA	HUANCAVELICA	ASCENSION	2025-12-05 15:55:51.729708-05	2025-12-05 15:55:51.729708-05
840	080120	090119	HUANCAVELICA	HUANCAVELICA	HUANDO	2025-12-05 15:55:51.730234-05	2025-12-05 15:55:51.730234-05
841	080201	090201	HUANCAVELICA	ACOBAMBA	ACOBAMBA	2025-12-05 15:55:51.731443-05	2025-12-05 15:55:51.731443-05
842	080203	090202	HUANCAVELICA	ACOBAMBA	ANDABAMBA	2025-12-05 15:55:51.732271-05	2025-12-05 15:55:51.732271-05
843	080202	090203	HUANCAVELICA	ACOBAMBA	ANTA	2025-12-05 15:55:51.732847-05	2025-12-05 15:55:51.732847-05
844	080204	090204	HUANCAVELICA	ACOBAMBA	CAJA	2025-12-05 15:55:51.733365-05	2025-12-05 15:55:51.733365-05
845	080205	090205	HUANCAVELICA	ACOBAMBA	MARCAS	2025-12-05 15:55:51.733885-05	2025-12-05 15:55:51.733885-05
846	080206	090206	HUANCAVELICA	ACOBAMBA	PAUCARA	2025-12-05 15:55:51.733885-05	2025-12-05 15:55:51.733885-05
847	080207	090207	HUANCAVELICA	ACOBAMBA	POMACOCHA	2025-12-05 15:55:51.734463-05	2025-12-05 15:55:51.734463-05
848	080208	090208	HUANCAVELICA	ACOBAMBA	ROSARIO	2025-12-05 15:55:51.735019-05	2025-12-05 15:55:51.735019-05
849	080301	090301	HUANCAVELICA	ANGARAES	LIRCAY	2025-12-05 15:55:51.735019-05	2025-12-05 15:55:51.735019-05
850	080302	090302	HUANCAVELICA	ANGARAES	ANCHONGA	2025-12-05 15:55:51.735589-05	2025-12-05 15:55:51.735589-05
851	080303	090303	HUANCAVELICA	ANGARAES	CALLANMARCA	2025-12-05 15:55:51.73611-05	2025-12-05 15:55:51.73611-05
852	080312	090304	HUANCAVELICA	ANGARAES	CCOCHACCASA	2025-12-05 15:55:51.73611-05	2025-12-05 15:55:51.73611-05
853	080305	090305	HUANCAVELICA	ANGARAES	CHINCHO	2025-12-05 15:55:51.737297-05	2025-12-05 15:55:51.737297-05
854	080304	090306	HUANCAVELICA	ANGARAES	CONGALLA	2025-12-05 15:55:51.737297-05	2025-12-05 15:55:51.737297-05
855	080307	090307	HUANCAVELICA	ANGARAES	HUANCA-HUANCA	2025-12-05 15:55:51.738139-05	2025-12-05 15:55:51.738139-05
856	080306	090308	HUANCAVELICA	ANGARAES	HUAYLLAY GRANDE	2025-12-05 15:55:51.738706-05	2025-12-05 15:55:51.738706-05
857	080308	090309	HUANCAVELICA	ANGARAES	JULCAMARCA	2025-12-05 15:55:51.738706-05	2025-12-05 15:55:51.738706-05
858	080309	090310	HUANCAVELICA	ANGARAES	SAN ANTONIO DE ANTAPARCO	2025-12-05 15:55:51.739285-05	2025-12-05 15:55:51.739285-05
859	080310	090311	HUANCAVELICA	ANGARAES	SANTO TOMAS DE PATA	2025-12-05 15:55:51.739285-05	2025-12-05 15:55:51.739285-05
860	080311	090312	HUANCAVELICA	ANGARAES	SECCLLA	2025-12-05 15:55:51.739843-05	2025-12-05 15:55:51.739843-05
861	080401	090401	HUANCAVELICA	CASTROVIRREYNA	CASTROVIRREYNA	2025-12-05 15:55:51.740367-05	2025-12-05 15:55:51.740367-05
862	080402	090402	HUANCAVELICA	CASTROVIRREYNA	ARMA	2025-12-05 15:55:51.740367-05	2025-12-05 15:55:51.740367-05
863	080403	090403	HUANCAVELICA	CASTROVIRREYNA	AURAHUA	2025-12-05 15:55:51.740971-05	2025-12-05 15:55:51.740971-05
864	080405	090404	HUANCAVELICA	CASTROVIRREYNA	CAPILLAS	2025-12-05 15:55:51.740971-05	2025-12-05 15:55:51.740971-05
865	080408	090405	HUANCAVELICA	CASTROVIRREYNA	CHUPAMARCA	2025-12-05 15:55:51.741525-05	2025-12-05 15:55:51.741525-05
866	080406	090406	HUANCAVELICA	CASTROVIRREYNA	COCAS	2025-12-05 15:55:51.742398-05	2025-12-05 15:55:51.742398-05
867	080409	090407	HUANCAVELICA	CASTROVIRREYNA	HUACHOS	2025-12-05 15:55:51.743023-05	2025-12-05 15:55:51.743023-05
868	080410	090408	HUANCAVELICA	CASTROVIRREYNA	HUAMATAMBO	2025-12-05 15:55:51.744084-05	2025-12-05 15:55:51.744084-05
869	080414	090409	HUANCAVELICA	CASTROVIRREYNA	MOLLEPAMPA	2025-12-05 15:55:51.744656-05	2025-12-05 15:55:51.744656-05
870	080422	090410	HUANCAVELICA	CASTROVIRREYNA	SAN JUAN	2025-12-05 15:55:51.745184-05	2025-12-05 15:55:51.745184-05
871	080429	090411	HUANCAVELICA	CASTROVIRREYNA	SANTA ANA	2025-12-05 15:55:51.745785-05	2025-12-05 15:55:51.745785-05
872	080427	090412	HUANCAVELICA	CASTROVIRREYNA	TANTARA	2025-12-05 15:55:51.746308-05	2025-12-05 15:55:51.746308-05
873	080428	090413	HUANCAVELICA	CASTROVIRREYNA	TICRAPO	2025-12-05 15:55:51.746843-05	2025-12-05 15:55:51.746843-05
874	080701	090501	HUANCAVELICA	CHURCAMPA	CHURCAMPA	2025-12-05 15:55:51.747785-05	2025-12-05 15:55:51.747785-05
875	080702	090502	HUANCAVELICA	CHURCAMPA	ANCO	2025-12-05 15:55:51.748895-05	2025-12-05 15:55:51.748895-05
876	080703	090503	HUANCAVELICA	CHURCAMPA	CHINCHIHUASI	2025-12-05 15:55:51.749401-05	2025-12-05 15:55:51.749401-05
877	080704	090504	HUANCAVELICA	CHURCAMPA	EL CARMEN	2025-12-05 15:55:51.749914-05	2025-12-05 15:55:51.749914-05
878	080705	090505	HUANCAVELICA	CHURCAMPA	LA MERCED	2025-12-05 15:55:51.750431-05	2025-12-05 15:55:51.750431-05
879	080706	090506	HUANCAVELICA	CHURCAMPA	LOCROJA	2025-12-05 15:55:51.750949-05	2025-12-05 15:55:51.750949-05
880	080707	090507	HUANCAVELICA	CHURCAMPA	PAUCARBAMBA	2025-12-05 15:55:51.750949-05	2025-12-05 15:55:51.750949-05
881	080708	090508	HUANCAVELICA	CHURCAMPA	SAN MIGUEL DE MAYOCC	2025-12-05 15:55:51.751473-05	2025-12-05 15:55:51.751473-05
882	080709	090509	HUANCAVELICA	CHURCAMPA	SAN PEDRO DE CORIS	2025-12-05 15:55:51.752036-05	2025-12-05 15:55:51.752036-05
883	080710	090510	HUANCAVELICA	CHURCAMPA	PACHAMARCA	2025-12-05 15:55:51.752525-05	2025-12-05 15:55:51.752525-05
884	080711	090511	HUANCAVELICA	CHURCAMPA	COSME	2025-12-05 15:55:51.752827-05	2025-12-05 15:55:51.752827-05
885	080604	090601	HUANCAVELICA	HUAYTARA	HUAYTARA	2025-12-05 15:55:51.75351-05	2025-12-05 15:55:51.75351-05
886	080601	090602	HUANCAVELICA	HUAYTARA	AYAVI	2025-12-05 15:55:51.754289-05	2025-12-05 15:55:51.754289-05
887	080602	090603	HUANCAVELICA	HUAYTARA	CORDOVA	2025-12-05 15:55:51.754939-05	2025-12-05 15:55:51.754939-05
888	080603	090604	HUANCAVELICA	HUAYTARA	HUAYACUNDO ARMA	2025-12-05 15:55:51.754939-05	2025-12-05 15:55:51.754939-05
889	080605	090605	HUANCAVELICA	HUAYTARA	LARAMARCA	2025-12-05 15:55:51.755461-05	2025-12-05 15:55:51.755461-05
890	080606	090606	HUANCAVELICA	HUAYTARA	OCOYO	2025-12-05 15:55:51.755974-05	2025-12-05 15:55:51.755974-05
891	080607	090607	HUANCAVELICA	HUAYTARA	PILPICHACA	2025-12-05 15:55:51.756631-05	2025-12-05 15:55:51.756631-05
892	080608	090608	HUANCAVELICA	HUAYTARA	QUERCO	2025-12-05 15:55:51.756631-05	2025-12-05 15:55:51.756631-05
893	080609	090609	HUANCAVELICA	HUAYTARA	QUITO-ARMA	2025-12-05 15:55:51.757148-05	2025-12-05 15:55:51.757148-05
894	080610	090610	HUANCAVELICA	HUAYTARA	SAN ANTONIO DE CUSICANCHA	2025-12-05 15:55:51.757667-05	2025-12-05 15:55:51.757667-05
895	080611	090611	HUANCAVELICA	HUAYTARA	SAN FRANCISCO DE SANGAYAICO	2025-12-05 15:55:51.758105-05	2025-12-05 15:55:51.758105-05
896	080612	090612	HUANCAVELICA	HUAYTARA	SAN ISIDRO	2025-12-05 15:55:51.758441-05	2025-12-05 15:55:51.758441-05
897	080613	090613	HUANCAVELICA	HUAYTARA	SANTIAGO DE CHOCORVOS	2025-12-05 15:55:51.758965-05	2025-12-05 15:55:51.758965-05
898	080614	090614	HUANCAVELICA	HUAYTARA	SANTIAGO DE QUIRAHUARA	2025-12-05 15:55:51.760009-05	2025-12-05 15:55:51.760009-05
899	080615	090615	HUANCAVELICA	HUAYTARA	SANTO DOMINGO DE CAPILLAS	2025-12-05 15:55:51.760537-05	2025-12-05 15:55:51.760537-05
900	080616	090616	HUANCAVELICA	HUAYTARA	TAMBO	2025-12-05 15:55:51.761233-05	2025-12-05 15:55:51.761233-05
901	080501	090701	HUANCAVELICA	TAYACAJA	PAMPAS	2025-12-05 15:55:51.761915-05	2025-12-05 15:55:51.761915-05
902	080502	090702	HUANCAVELICA	TAYACAJA	ACOSTAMBO	2025-12-05 15:55:51.762594-05	2025-12-05 15:55:51.762594-05
903	080503	090703	HUANCAVELICA	TAYACAJA	ACRAQUIA	2025-12-05 15:55:51.76312-05	2025-12-05 15:55:51.76312-05
904	080504	090704	HUANCAVELICA	TAYACAJA	AHUAYCHA	2025-12-05 15:55:51.763993-05	2025-12-05 15:55:51.763993-05
905	080506	090705	HUANCAVELICA	TAYACAJA	COLCABAMBA	2025-12-05 15:55:51.764586-05	2025-12-05 15:55:51.764586-05
906	080509	090706	HUANCAVELICA	TAYACAJA	DANIEL HERNANDEZ	2025-12-05 15:55:51.765252-05	2025-12-05 15:55:51.765252-05
907	080511	090707	HUANCAVELICA	TAYACAJA	HUACHOCOLPA	2025-12-05 15:55:51.765772-05	2025-12-05 15:55:51.765772-05
908	080512	090709	HUANCAVELICA	TAYACAJA	HUARIBAMBA	2025-12-05 15:55:51.766365-05	2025-12-05 15:55:51.766365-05
909	080515	090710	HUANCAVELICA	TAYACAJA	ÑAHUIMPUQUIO	2025-12-05 15:55:51.766965-05	2025-12-05 15:55:51.766965-05
910	080517	090711	HUANCAVELICA	TAYACAJA	PAZOS	2025-12-05 15:55:51.766965-05	2025-12-05 15:55:51.766965-05
911	080518	090713	HUANCAVELICA	TAYACAJA	QUISHUAR	2025-12-05 15:55:51.767514-05	2025-12-05 15:55:51.767514-05
912	080519	090714	HUANCAVELICA	TAYACAJA	SALCABAMBA	2025-12-05 15:55:51.768076-05	2025-12-05 15:55:51.768076-05
913	080526	090715	HUANCAVELICA	TAYACAJA	SALCAHUASI	2025-12-05 15:55:51.768526-05	2025-12-05 15:55:51.768526-05
914	080520	090716	HUANCAVELICA	TAYACAJA	SAN MARCOS DE ROCCHAC	2025-12-05 15:55:51.769048-05	2025-12-05 15:55:51.769048-05
915	080523	090717	HUANCAVELICA	TAYACAJA	SURCUBAMBA	2025-12-05 15:55:51.769573-05	2025-12-05 15:55:51.769573-05
916	080525	090718	HUANCAVELICA	TAYACAJA	TINTAY PUNCU	2025-12-05 15:55:51.770232-05	2025-12-05 15:55:51.770232-05
917	080528	090719	HUANCAVELICA	TAYACAJA	QUICHUAS	2025-12-05 15:55:51.770855-05	2025-12-05 15:55:51.770855-05
918	080529	090720	HUANCAVELICA	TAYACAJA	ANDAYMARCA	2025-12-05 15:55:51.771408-05	2025-12-05 15:55:51.771408-05
919	080530	090721	HUANCAVELICA	TAYACAJA	ROBLE	2025-12-05 15:55:51.771925-05	2025-12-05 15:55:51.771925-05
920	080531	090722	HUANCAVELICA	TAYACAJA	PICHOS	2025-12-05 15:55:51.772499-05	2025-12-05 15:55:51.772499-05
921	080532	090723	HUANCAVELICA	TAYACAJA	SANTIAGO DE TUCUMA	2025-12-05 15:55:51.772499-05	2025-12-05 15:55:51.772499-05
922	080533	090724	HUANCAVELICA	TAYACAJA	LAMBRAS	2025-12-05 15:55:51.773548-05	2025-12-05 15:55:51.773548-05
923	080534	090725	HUANCAVELICA	TAYACAJA	COCHABAMBA	2025-12-05 15:55:51.773883-05	2025-12-05 15:55:51.773883-05
924	090101	100101	HUANUCO	HUANUCO	HUANUCO	2025-12-05 15:55:51.773883-05	2025-12-05 15:55:51.773883-05
925	090110	100102	HUANUCO	HUANUCO	AMARILIS	2025-12-05 15:55:51.77475-05	2025-12-05 15:55:51.77475-05
926	090102	100103	HUANUCO	HUANUCO	CHINCHAO	2025-12-05 15:55:51.77475-05	2025-12-05 15:55:51.77475-05
927	090103	100104	HUANUCO	HUANUCO	CHURUBAMBA	2025-12-05 15:55:51.775322-05	2025-12-05 15:55:51.775322-05
928	090104	100105	HUANUCO	HUANUCO	MARGOS	2025-12-05 15:55:51.775886-05	2025-12-05 15:55:51.775886-05
929	090105	100106	HUANUCO	HUANUCO	QUISQUI	2025-12-05 15:55:51.776589-05	2025-12-05 15:55:51.776589-05
930	090106	100107	HUANUCO	HUANUCO	SAN FRANCISCO DE CAYRAN	2025-12-05 15:55:51.77723-05	2025-12-05 15:55:51.77723-05
931	090107	100108	HUANUCO	HUANUCO	SAN PEDRO DE CHAULAN	2025-12-05 15:55:51.778318-05	2025-12-05 15:55:51.778318-05
932	090108	100109	HUANUCO	HUANUCO	SANTA MARIA DEL VALLE	2025-12-05 15:55:51.778859-05	2025-12-05 15:55:51.778859-05
933	090109	100110	HUANUCO	HUANUCO	YARUMAYO	2025-12-05 15:55:51.779479-05	2025-12-05 15:55:51.779479-05
934	090111	100111	HUANUCO	HUANUCO	PILLCO MARCA	2025-12-05 15:55:51.779996-05	2025-12-05 15:55:51.779996-05
935	090112	100112	HUANUCO	HUANUCO	YACUS	2025-12-05 15:55:51.780611-05	2025-12-05 15:55:51.780611-05
936	090113	100113	HUANUCO	HUANUCO	SAN PABLO DE PILLAO	2025-12-05 15:55:51.781126-05	2025-12-05 15:55:51.781126-05
937	090201	100201	HUANUCO	AMBO	AMBO	2025-12-05 15:55:51.782046-05	2025-12-05 15:55:51.782046-05
938	090202	100202	HUANUCO	AMBO	CAYNA	2025-12-05 15:55:51.782436-05	2025-12-05 15:55:51.782436-05
939	090203	100203	HUANUCO	AMBO	COLPAS	2025-12-05 15:55:51.783811-05	2025-12-05 15:55:51.783811-05
940	090204	100204	HUANUCO	AMBO	CONCHAMARCA	2025-12-05 15:55:51.784368-05	2025-12-05 15:55:51.784368-05
941	090205	100205	HUANUCO	AMBO	HUACAR	2025-12-05 15:55:51.784368-05	2025-12-05 15:55:51.784368-05
942	090206	100206	HUANUCO	AMBO	SAN FRANCISCO	2025-12-05 15:55:51.784951-05	2025-12-05 15:55:51.784951-05
943	090207	100207	HUANUCO	AMBO	SAN RAFAEL	2025-12-05 15:55:51.785578-05	2025-12-05 15:55:51.785578-05
944	090208	100208	HUANUCO	AMBO	TOMAY KICHWA	2025-12-05 15:55:51.785578-05	2025-12-05 15:55:51.785578-05
945	090301	100301	HUANUCO	DOS DE MAYO	LA UNION	2025-12-05 15:55:51.786144-05	2025-12-05 15:55:51.786144-05
946	090307	100307	HUANUCO	DOS DE MAYO	CHUQUIS	2025-12-05 15:55:51.786697-05	2025-12-05 15:55:51.786697-05
947	090312	100311	HUANUCO	DOS DE MAYO	MARIAS	2025-12-05 15:55:51.786697-05	2025-12-05 15:55:51.786697-05
948	090314	100313	HUANUCO	DOS DE MAYO	PACHAS	2025-12-05 15:55:51.787218-05	2025-12-05 15:55:51.787218-05
949	090316	100316	HUANUCO	DOS DE MAYO	QUIVILLA	2025-12-05 15:55:51.787769-05	2025-12-05 15:55:51.787769-05
950	090317	100317	HUANUCO	DOS DE MAYO	RIPAN	2025-12-05 15:55:51.788183-05	2025-12-05 15:55:51.788183-05
951	090321	100321	HUANUCO	DOS DE MAYO	SHUNQUI	2025-12-05 15:55:51.788673-05	2025-12-05 15:55:51.788673-05
952	090322	100322	HUANUCO	DOS DE MAYO	SILLAPATA	2025-12-05 15:55:51.788673-05	2025-12-05 15:55:51.788673-05
953	090323	100323	HUANUCO	DOS DE MAYO	YANAS	2025-12-05 15:55:51.789407-05	2025-12-05 15:55:51.789407-05
954	090901	100401	HUANUCO	HUACAYBAMBA	HUACAYBAMBA	2025-12-05 15:55:51.789933-05	2025-12-05 15:55:51.789933-05
955	090903	100402	HUANUCO	HUACAYBAMBA	CANCHABAMBA	2025-12-05 15:55:51.790455-05	2025-12-05 15:55:51.790455-05
956	090904	100403	HUANUCO	HUACAYBAMBA	COCHABAMBA	2025-12-05 15:55:51.790455-05	2025-12-05 15:55:51.790455-05
957	090902	100404	HUANUCO	HUACAYBAMBA	PINRA	2025-12-05 15:55:51.791034-05	2025-12-05 15:55:51.791034-05
958	090401	100501	HUANUCO	HUAMALIES	LLATA	2025-12-05 15:55:51.791034-05	2025-12-05 15:55:51.791034-05
959	090402	100502	HUANUCO	HUAMALIES	ARANCAY	2025-12-05 15:55:51.791584-05	2025-12-05 15:55:51.791584-05
960	090403	100503	HUANUCO	HUAMALIES	CHAVIN DE PARIARCA	2025-12-05 15:55:51.792122-05	2025-12-05 15:55:51.792122-05
961	090404	100504	HUANUCO	HUAMALIES	JACAS GRANDE	2025-12-05 15:55:51.792633-05	2025-12-05 15:55:51.792633-05
962	090405	100505	HUANUCO	HUAMALIES	JIRCAN	2025-12-05 15:55:51.793184-05	2025-12-05 15:55:51.793184-05
963	090406	100506	HUANUCO	HUAMALIES	MIRAFLORES	2025-12-05 15:55:51.794419-05	2025-12-05 15:55:51.794419-05
964	090407	100507	HUANUCO	HUAMALIES	MONZON	2025-12-05 15:55:51.794945-05	2025-12-05 15:55:51.794945-05
965	090408	100508	HUANUCO	HUAMALIES	PUNCHAO	2025-12-05 15:55:51.795547-05	2025-12-05 15:55:51.795547-05
966	090409	100509	HUANUCO	HUAMALIES	PUÑOS	2025-12-05 15:55:51.796129-05	2025-12-05 15:55:51.796129-05
967	090410	100510	HUANUCO	HUAMALIES	SINGA	2025-12-05 15:55:51.796129-05	2025-12-05 15:55:51.796129-05
968	090411	100511	HUANUCO	HUAMALIES	TANTAMAYO	2025-12-05 15:55:51.797191-05	2025-12-05 15:55:51.797191-05
969	090601	100601	HUANUCO	LEONCIO PRADO	RUPA-RUPA	2025-12-05 15:55:51.798047-05	2025-12-05 15:55:51.798047-05
970	090602	100602	HUANUCO	LEONCIO PRADO	DANIEL ALOMIAS ROBLES	2025-12-05 15:55:51.799428-05	2025-12-05 15:55:51.799428-05
971	090603	100603	HUANUCO	LEONCIO PRADO	HERMILIO VALDIZAN	2025-12-05 15:55:51.799921-05	2025-12-05 15:55:51.799921-05
972	090606	100604	HUANUCO	LEONCIO PRADO	JOSE CRESPO Y CASTILLO	2025-12-05 15:55:51.800447-05	2025-12-05 15:55:51.800447-05
973	090604	100605	HUANUCO	LEONCIO PRADO	LUYANDO	2025-12-05 15:55:51.801003-05	2025-12-05 15:55:51.801003-05
974	090605	100606	HUANUCO	LEONCIO PRADO	MARIANO DAMASO BERAUN	2025-12-05 15:55:51.801003-05	2025-12-05 15:55:51.801003-05
975	090607	100607	HUANUCO	LEONCIO PRADO	PUCAYACU	2025-12-05 15:55:51.801527-05	2025-12-05 15:55:51.801527-05
976	090608	100608	HUANUCO	LEONCIO PRADO	CASTILLO GRANDE	2025-12-05 15:55:51.802069-05	2025-12-05 15:55:51.802069-05
977	090609	100609	HUANUCO	LEONCIO PRADO	PUEBLO NUEVO	2025-12-05 15:55:51.802069-05	2025-12-05 15:55:51.802069-05
978	090610	100610	HUANUCO	LEONCIO PRADO	SANTO DOMINGO DE ANDA	2025-12-05 15:55:51.802584-05	2025-12-05 15:55:51.802584-05
979	090501	100701	HUANUCO	MARAÑON	HUACRACHUCO	2025-12-05 15:55:51.803097-05	2025-12-05 15:55:51.803097-05
980	090502	100702	HUANUCO	MARAÑON	CHOLON	2025-12-05 15:55:51.803626-05	2025-12-05 15:55:51.803626-05
981	090505	100703	HUANUCO	MARAÑON	SAN BUENAVENTURA	2025-12-05 15:55:51.803626-05	2025-12-05 15:55:51.803626-05
982	090506	100704	HUANUCO	MARAÑON	LA MORADA	2025-12-05 15:55:51.804134-05	2025-12-05 15:55:51.804134-05
983	090507	100705	HUANUCO	MARAÑON	SANTA ROSA DE ALTO YANAJANCA	2025-12-05 15:55:51.804656-05	2025-12-05 15:55:51.804656-05
984	090701	100801	HUANUCO	PACHITEA	PANAO	2025-12-05 15:55:51.805248-05	2025-12-05 15:55:51.805248-05
985	090702	100802	HUANUCO	PACHITEA	CHAGLLA	2025-12-05 15:55:51.805248-05	2025-12-05 15:55:51.805248-05
986	090704	100803	HUANUCO	PACHITEA	MOLINO	2025-12-05 15:55:51.80599-05	2025-12-05 15:55:51.80599-05
987	090706	100804	HUANUCO	PACHITEA	UMARI	2025-12-05 15:55:51.806535-05	2025-12-05 15:55:51.806535-05
988	090802	100901	HUANUCO	PUERTO INCA	PUERTO INCA	2025-12-05 15:55:51.806535-05	2025-12-05 15:55:51.806535-05
989	090803	100902	HUANUCO	PUERTO INCA	CODO DEL POZUZO	2025-12-05 15:55:51.807155-05	2025-12-05 15:55:51.807155-05
990	090801	100903	HUANUCO	PUERTO INCA	HONORIA	2025-12-05 15:55:51.807155-05	2025-12-05 15:55:51.807155-05
991	090804	100904	HUANUCO	PUERTO INCA	TOURNAVISTA	2025-12-05 15:55:51.80772-05	2025-12-05 15:55:51.80772-05
992	090805	100905	HUANUCO	PUERTO INCA	YUYAPICHIS	2025-12-05 15:55:51.808239-05	2025-12-05 15:55:51.808239-05
993	091001	101001	HUANUCO	LAURICOCHA	JESUS	2025-12-05 15:55:51.808239-05	2025-12-05 15:55:51.808239-05
994	091002	101002	HUANUCO	LAURICOCHA	BAÑOS	2025-12-05 15:55:51.808755-05	2025-12-05 15:55:51.808755-05
995	091007	101003	HUANUCO	LAURICOCHA	JIVIA	2025-12-05 15:55:51.809431-05	2025-12-05 15:55:51.809431-05
996	091004	101004	HUANUCO	LAURICOCHA	QUEROPALCA	2025-12-05 15:55:51.810063-05	2025-12-05 15:55:51.810063-05
997	091006	101005	HUANUCO	LAURICOCHA	RONDOS	2025-12-05 15:55:51.810745-05	2025-12-05 15:55:51.810745-05
998	091003	101006	HUANUCO	LAURICOCHA	SAN FRANCISCO DE ASIS	2025-12-05 15:55:51.811862-05	2025-12-05 15:55:51.811862-05
999	091005	101007	HUANUCO	LAURICOCHA	SAN MIGUEL DE CAURI	2025-12-05 15:55:51.812396-05	2025-12-05 15:55:51.812396-05
1000	091101	101101	HUANUCO	YAROWILCA	CHAVINILLO	2025-12-05 15:55:51.812396-05	2025-12-05 15:55:51.812396-05
1001	091103	101102	HUANUCO	YAROWILCA	CAHUAC	2025-12-05 15:55:51.812928-05	2025-12-05 15:55:51.812928-05
1002	091104	101103	HUANUCO	YAROWILCA	CHACABAMBA	2025-12-05 15:55:51.813449-05	2025-12-05 15:55:51.813449-05
1003	091102	101104	HUANUCO	YAROWILCA	APARICIO POMARES	2025-12-05 15:55:51.8145-05	2025-12-05 15:55:51.8145-05
1004	091105	101105	HUANUCO	YAROWILCA	JACAS CHICO	2025-12-05 15:55:51.815862-05	2025-12-05 15:55:51.815862-05
1005	091106	101106	HUANUCO	YAROWILCA	OBAS	2025-12-05 15:55:51.816394-05	2025-12-05 15:55:51.816394-05
1006	091107	101107	HUANUCO	YAROWILCA	PAMPAMARCA	2025-12-05 15:55:51.817096-05	2025-12-05 15:55:51.817096-05
1007	091108	101108	HUANUCO	YAROWILCA	CHORAS	2025-12-05 15:55:51.817758-05	2025-12-05 15:55:51.817758-05
1008	100101	110101	ICA	ICA	ICA	2025-12-05 15:55:51.81828-05	2025-12-05 15:55:51.81828-05
1009	100102	110102	ICA	ICA	LA TINGUIÑA	2025-12-05 15:55:51.81828-05	2025-12-05 15:55:51.81828-05
1010	100103	110103	ICA	ICA	LOS AQUIJES	2025-12-05 15:55:51.818823-05	2025-12-05 15:55:51.818823-05
1011	100114	110104	ICA	ICA	OCUCAJE	2025-12-05 15:55:51.819349-05	2025-12-05 15:55:51.819349-05
1012	100113	110105	ICA	ICA	PACHACUTEC	2025-12-05 15:55:51.819928-05	2025-12-05 15:55:51.819928-05
1013	100104	110106	ICA	ICA	PARCONA	2025-12-05 15:55:51.820527-05	2025-12-05 15:55:51.820527-05
1014	100105	110107	ICA	ICA	PUEBLO NUEVO	2025-12-05 15:55:51.82114-05	2025-12-05 15:55:51.82114-05
1015	100106	110108	ICA	ICA	SALAS	2025-12-05 15:55:51.821903-05	2025-12-05 15:55:51.821903-05
1016	100107	110109	ICA	ICA	SAN JOSE DE LOS MOLINOS	2025-12-05 15:55:51.821903-05	2025-12-05 15:55:51.821903-05
1017	100108	110110	ICA	ICA	SAN JUAN BAUTISTA	2025-12-05 15:55:51.822442-05	2025-12-05 15:55:51.822442-05
1018	100109	110111	ICA	ICA	SANTIAGO	2025-12-05 15:55:51.822956-05	2025-12-05 15:55:51.822956-05
1019	100110	110112	ICA	ICA	SUBTANJALLA	2025-12-05 15:55:51.823481-05	2025-12-05 15:55:51.823481-05
1020	100112	110113	ICA	ICA	TATE	2025-12-05 15:55:51.823998-05	2025-12-05 15:55:51.823998-05
1021	100111	110114	ICA	ICA	YAUCA DEL ROSARIO	2025-12-05 15:55:51.824519-05	2025-12-05 15:55:51.824519-05
1022	100201	110201	ICA	CHINCHA	CHINCHA ALTA	2025-12-05 15:55:51.825043-05	2025-12-05 15:55:51.825043-05
1023	100209	110202	ICA	CHINCHA	ALTO LARAN	2025-12-05 15:55:51.826168-05	2025-12-05 15:55:51.826168-05
1024	100202	110203	ICA	CHINCHA	CHAVIN	2025-12-05 15:55:51.826715-05	2025-12-05 15:55:51.826715-05
1025	100203	110204	ICA	CHINCHA	CHINCHA BAJA	2025-12-05 15:55:51.827796-05	2025-12-05 15:55:51.827796-05
1026	100204	110205	ICA	CHINCHA	EL CARMEN	2025-12-05 15:55:51.828372-05	2025-12-05 15:55:51.828372-05
1027	100205	110206	ICA	CHINCHA	GROCIO PRADO	2025-12-05 15:55:51.8289-05	2025-12-05 15:55:51.8289-05
1028	100210	110207	ICA	CHINCHA	PUEBLO NUEVO	2025-12-05 15:55:51.829573-05	2025-12-05 15:55:51.829573-05
1029	100211	110208	ICA	CHINCHA	SAN JUAN DE YANAC	2025-12-05 15:55:51.830183-05	2025-12-05 15:55:51.830183-05
1030	100206	110209	ICA	CHINCHA	SAN PEDRO DE HUACARPANA	2025-12-05 15:55:51.830874-05	2025-12-05 15:55:51.830874-05
1031	100207	110210	ICA	CHINCHA	SUNAMPE	2025-12-05 15:55:51.832212-05	2025-12-05 15:55:51.832212-05
1032	100208	110211	ICA	CHINCHA	TAMBO DE MORA	2025-12-05 15:55:51.833372-05	2025-12-05 15:55:51.833372-05
1033	100301	110301	ICA	NAZCA	NAZCA	2025-12-05 15:55:51.834034-05	2025-12-05 15:55:51.834034-05
1034	100302	110302	ICA	NAZCA	CHANGUILLO	2025-12-05 15:55:51.834588-05	2025-12-05 15:55:51.834588-05
1035	100303	110303	ICA	NAZCA	EL INGENIO	2025-12-05 15:55:51.835107-05	2025-12-05 15:55:51.835107-05
1036	100304	110304	ICA	NAZCA	MARCONA	2025-12-05 15:55:51.835634-05	2025-12-05 15:55:51.835634-05
1037	100305	110305	ICA	NAZCA	VISTA ALEGRE	2025-12-05 15:55:51.836158-05	2025-12-05 15:55:51.836158-05
1038	100501	110401	ICA	PALPA	PALPA	2025-12-05 15:55:51.836686-05	2025-12-05 15:55:51.836686-05
1039	100502	110402	ICA	PALPA	LLIPATA	2025-12-05 15:55:51.836686-05	2025-12-05 15:55:51.836686-05
1040	100503	110403	ICA	PALPA	RIO GRANDE	2025-12-05 15:55:51.83721-05	2025-12-05 15:55:51.83721-05
1041	100504	110404	ICA	PALPA	SANTA CRUZ	2025-12-05 15:55:51.837732-05	2025-12-05 15:55:51.837732-05
1042	100505	110405	ICA	PALPA	TIBILLO	2025-12-05 15:55:51.838282-05	2025-12-05 15:55:51.838282-05
1043	100401	110501	ICA	PISCO	PISCO	2025-12-05 15:55:51.838825-05	2025-12-05 15:55:51.838825-05
1044	100402	110502	ICA	PISCO	HUANCANO	2025-12-05 15:55:51.839851-05	2025-12-05 15:55:51.839851-05
1045	100403	110503	ICA	PISCO	HUMAY	2025-12-05 15:55:51.840368-05	2025-12-05 15:55:51.840368-05
1046	100404	110504	ICA	PISCO	INDEPENDENCIA	2025-12-05 15:55:51.840964-05	2025-12-05 15:55:51.840964-05
1047	100405	110505	ICA	PISCO	PARACAS	2025-12-05 15:55:51.841479-05	2025-12-05 15:55:51.841479-05
1048	100406	110506	ICA	PISCO	SAN ANDRES	2025-12-05 15:55:51.842009-05	2025-12-05 15:55:51.842009-05
1049	100407	110507	ICA	PISCO	SAN CLEMENTE	2025-12-05 15:55:51.842531-05	2025-12-05 15:55:51.842531-05
1050	100408	110508	ICA	PISCO	TUPAC AMARU INCA	2025-12-05 15:55:51.843061-05	2025-12-05 15:55:51.843061-05
1051	110101	120101	JUNIN	HUANCAYO	HUANCAYO	2025-12-05 15:55:51.843616-05	2025-12-05 15:55:51.843616-05
1052	110103	120104	JUNIN	HUANCAYO	CARHUACALLANGA	2025-12-05 15:55:51.844268-05	2025-12-05 15:55:51.844268-05
1053	110106	120105	JUNIN	HUANCAYO	CHACAPAMPA	2025-12-05 15:55:51.845334-05	2025-12-05 15:55:51.845334-05
1054	110107	120106	JUNIN	HUANCAYO	CHICCHE	2025-12-05 15:55:51.845859-05	2025-12-05 15:55:51.845859-05
1055	110108	120107	JUNIN	HUANCAYO	CHILCA	2025-12-05 15:55:51.846382-05	2025-12-05 15:55:51.846382-05
1056	110109	120108	JUNIN	HUANCAYO	CHONGOS ALTO	2025-12-05 15:55:51.847487-05	2025-12-05 15:55:51.847487-05
1057	110112	120111	JUNIN	HUANCAYO	CHUPURO	2025-12-05 15:55:51.848004-05	2025-12-05 15:55:51.848004-05
1058	110104	120112	JUNIN	HUANCAYO	COLCA	2025-12-05 15:55:51.848935-05	2025-12-05 15:55:51.848935-05
1059	110105	120113	JUNIN	HUANCAYO	CULLHUAS	2025-12-05 15:55:51.849613-05	2025-12-05 15:55:51.849613-05
1060	110113	120114	JUNIN	HUANCAYO	EL TAMBO	2025-12-05 15:55:51.850179-05	2025-12-05 15:55:51.850179-05
1061	110114	120116	JUNIN	HUANCAYO	HUACRAPUQUIO	2025-12-05 15:55:51.850867-05	2025-12-05 15:55:51.850867-05
1062	110116	120117	JUNIN	HUANCAYO	HUALHUAS	2025-12-05 15:55:51.8514-05	2025-12-05 15:55:51.8514-05
1063	110118	120119	JUNIN	HUANCAYO	HUANCAN	2025-12-05 15:55:51.851938-05	2025-12-05 15:55:51.851938-05
1064	110119	120120	JUNIN	HUANCAYO	HUASICANCHA	2025-12-05 15:55:51.851938-05	2025-12-05 15:55:51.851938-05
1065	110120	120121	JUNIN	HUANCAYO	HUAYUCACHI	2025-12-05 15:55:51.852527-05	2025-12-05 15:55:51.852527-05
1066	110121	120122	JUNIN	HUANCAYO	INGENIO	2025-12-05 15:55:51.853053-05	2025-12-05 15:55:51.853053-05
1067	110122	120124	JUNIN	HUANCAYO	PARIAHUANCA	2025-12-05 15:55:51.853053-05	2025-12-05 15:55:51.853053-05
1068	110123	120125	JUNIN	HUANCAYO	PILCOMAYO	2025-12-05 15:55:51.853573-05	2025-12-05 15:55:51.853573-05
1069	110124	120126	JUNIN	HUANCAYO	PUCARA	2025-12-05 15:55:51.854116-05	2025-12-05 15:55:51.854116-05
1070	110125	120127	JUNIN	HUANCAYO	QUICHUAY	2025-12-05 15:55:51.854653-05	2025-12-05 15:55:51.854653-05
1071	110126	120128	JUNIN	HUANCAYO	QUILCAS	2025-12-05 15:55:51.855233-05	2025-12-05 15:55:51.855233-05
1072	110127	120129	JUNIN	HUANCAYO	SAN AGUSTIN	2025-12-05 15:55:51.855233-05	2025-12-05 15:55:51.855233-05
1073	110128	120130	JUNIN	HUANCAYO	SAN JERONIMO DE TUNAN	2025-12-05 15:55:51.855755-05	2025-12-05 15:55:51.855755-05
1074	110132	120132	JUNIN	HUANCAYO	SAÑO	2025-12-05 15:55:51.856268-05	2025-12-05 15:55:51.856268-05
1075	110133	120133	JUNIN	HUANCAYO	SAPALLANGA	2025-12-05 15:55:51.857029-05	2025-12-05 15:55:51.857029-05
1076	110134	120134	JUNIN	HUANCAYO	SICAYA	2025-12-05 15:55:51.857029-05	2025-12-05 15:55:51.857029-05
1077	110131	120135	JUNIN	HUANCAYO	SANTO DOMINGO DE ACOBAMBA	2025-12-05 15:55:51.857555-05	2025-12-05 15:55:51.857555-05
1078	110136	120136	JUNIN	HUANCAYO	VIQUES	2025-12-05 15:55:51.858077-05	2025-12-05 15:55:51.858077-05
1079	110201	120201	JUNIN	CONCEPCION	CONCEPCION	2025-12-05 15:55:51.858077-05	2025-12-05 15:55:51.858077-05
1080	110202	120202	JUNIN	CONCEPCION	ACO	2025-12-05 15:55:51.858597-05	2025-12-05 15:55:51.858597-05
1081	110203	120203	JUNIN	CONCEPCION	ANDAMARCA	2025-12-05 15:55:51.859141-05	2025-12-05 15:55:51.859141-05
1082	110206	120204	JUNIN	CONCEPCION	CHAMBARA	2025-12-05 15:55:51.859658-05	2025-12-05 15:55:51.859658-05
1083	110205	120205	JUNIN	CONCEPCION	COCHAS	2025-12-05 15:55:51.860382-05	2025-12-05 15:55:51.860382-05
1084	110204	120206	JUNIN	CONCEPCION	COMAS	2025-12-05 15:55:51.861433-05	2025-12-05 15:55:51.861433-05
1085	110207	120207	JUNIN	CONCEPCION	HEROINAS TOLEDO	2025-12-05 15:55:51.861968-05	2025-12-05 15:55:51.861968-05
1086	110208	120208	JUNIN	CONCEPCION	MANZANARES	2025-12-05 15:55:51.86258-05	2025-12-05 15:55:51.86258-05
1087	110209	120209	JUNIN	CONCEPCION	MARISCAL CASTILLA	2025-12-05 15:55:51.863097-05	2025-12-05 15:55:51.863097-05
1088	110210	120210	JUNIN	CONCEPCION	MATAHUASI	2025-12-05 15:55:51.863618-05	2025-12-05 15:55:51.863618-05
1089	110211	120211	JUNIN	CONCEPCION	MITO	2025-12-05 15:55:51.864472-05	2025-12-05 15:55:51.864472-05
1090	110212	120212	JUNIN	CONCEPCION	NUEVE DE JULIO	2025-12-05 15:55:51.865183-05	2025-12-05 15:55:51.865183-05
1091	110213	120213	JUNIN	CONCEPCION	ORCOTUNA	2025-12-05 15:55:51.865746-05	2025-12-05 15:55:51.865746-05
1092	110215	120214	JUNIN	CONCEPCION	SAN JOSE DE QUERO	2025-12-05 15:55:51.866343-05	2025-12-05 15:55:51.866343-05
1093	110214	120215	JUNIN	CONCEPCION	SANTA ROSA DE OCOPA	2025-12-05 15:55:51.867225-05	2025-12-05 15:55:51.867225-05
1094	110801	120301	JUNIN	CHANCHAMAYO	CHANCHAMAYO	2025-12-05 15:55:51.867225-05	2025-12-05 15:55:51.867225-05
1095	110806	120302	JUNIN	CHANCHAMAYO	PERENE	2025-12-05 15:55:51.867748-05	2025-12-05 15:55:51.867748-05
1096	110805	120303	JUNIN	CHANCHAMAYO	PICHANAQUI	2025-12-05 15:55:51.868288-05	2025-12-05 15:55:51.868288-05
1097	110804	120304	JUNIN	CHANCHAMAYO	SAN LUIS DE SHUARO	2025-12-05 15:55:51.868805-05	2025-12-05 15:55:51.868805-05
1098	110802	120305	JUNIN	CHANCHAMAYO	SAN RAMON	2025-12-05 15:55:51.868805-05	2025-12-05 15:55:51.868805-05
1099	110803	120306	JUNIN	CHANCHAMAYO	VITOC	2025-12-05 15:55:51.86933-05	2025-12-05 15:55:51.86933-05
1100	110301	120401	JUNIN	JAUJA	JAUJA	2025-12-05 15:55:51.869904-05	2025-12-05 15:55:51.869904-05
1101	110302	120402	JUNIN	JAUJA	ACOLLA	2025-12-05 15:55:51.870539-05	2025-12-05 15:55:51.870539-05
1102	110303	120403	JUNIN	JAUJA	APATA	2025-12-05 15:55:51.870539-05	2025-12-05 15:55:51.870539-05
1103	110304	120404	JUNIN	JAUJA	ATAURA	2025-12-05 15:55:51.871077-05	2025-12-05 15:55:51.871077-05
1104	110305	120405	JUNIN	JAUJA	CANCHAYLLO	2025-12-05 15:55:51.871587-05	2025-12-05 15:55:51.871587-05
1105	110331	120406	JUNIN	JAUJA	CURICACA	2025-12-05 15:55:51.872112-05	2025-12-05 15:55:51.872112-05
1106	110306	120407	JUNIN	JAUJA	EL MANTARO	2025-12-05 15:55:51.872112-05	2025-12-05 15:55:51.872112-05
1107	110307	120408	JUNIN	JAUJA	HUAMALI	2025-12-05 15:55:51.872823-05	2025-12-05 15:55:51.872823-05
1108	110308	120409	JUNIN	JAUJA	HUARIPAMPA	2025-12-05 15:55:51.873388-05	2025-12-05 15:55:51.873388-05
1109	110309	120410	JUNIN	JAUJA	HUERTAS	2025-12-05 15:55:51.873388-05	2025-12-05 15:55:51.873388-05
1110	110310	120411	JUNIN	JAUJA	JANJAILLO	2025-12-05 15:55:51.873937-05	2025-12-05 15:55:51.873937-05
1111	110311	120412	JUNIN	JAUJA	JULCAN	2025-12-05 15:55:51.873937-05	2025-12-05 15:55:51.873937-05
1112	110312	120413	JUNIN	JAUJA	LEONOR ORDOÑEZ	2025-12-05 15:55:51.874501-05	2025-12-05 15:55:51.874501-05
1113	110313	120414	JUNIN	JAUJA	LLOCLLAPAMPA	2025-12-05 15:55:51.875043-05	2025-12-05 15:55:51.875043-05
1114	110314	120415	JUNIN	JAUJA	MARCO	2025-12-05 15:55:51.875043-05	2025-12-05 15:55:51.875043-05
1115	110315	120416	JUNIN	JAUJA	MASMA	2025-12-05 15:55:51.876095-05	2025-12-05 15:55:51.876095-05
1116	110332	120417	JUNIN	JAUJA	MASMA CHICCHE	2025-12-05 15:55:51.876623-05	2025-12-05 15:55:51.876623-05
1117	110316	120418	JUNIN	JAUJA	MOLINOS	2025-12-05 15:55:51.877146-05	2025-12-05 15:55:51.877146-05
1118	110317	120419	JUNIN	JAUJA	MONOBAMBA	2025-12-05 15:55:51.877726-05	2025-12-05 15:55:51.877726-05
1119	110318	120420	JUNIN	JAUJA	MUQUI	2025-12-05 15:55:51.878276-05	2025-12-05 15:55:51.878276-05
1120	110319	120421	JUNIN	JAUJA	MUQUIYAUYO	2025-12-05 15:55:51.878839-05	2025-12-05 15:55:51.878839-05
1121	110320	120422	JUNIN	JAUJA	PACA	2025-12-05 15:55:51.879363-05	2025-12-05 15:55:51.879363-05
1122	110321	120423	JUNIN	JAUJA	PACCHA	2025-12-05 15:55:51.879921-05	2025-12-05 15:55:51.879921-05
1123	110322	120424	JUNIN	JAUJA	PANCAN	2025-12-05 15:55:51.881013-05	2025-12-05 15:55:51.881013-05
1124	110323	120425	JUNIN	JAUJA	PARCO	2025-12-05 15:55:51.882008-05	2025-12-05 15:55:51.882008-05
1125	110324	120426	JUNIN	JAUJA	POMACANCHA	2025-12-05 15:55:51.882538-05	2025-12-05 15:55:51.882538-05
1126	110325	120427	JUNIN	JAUJA	RICRAN	2025-12-05 15:55:51.883192-05	2025-12-05 15:55:51.883192-05
1127	110326	120428	JUNIN	JAUJA	SAN LORENZO	2025-12-05 15:55:51.883821-05	2025-12-05 15:55:51.883821-05
1128	110327	120429	JUNIN	JAUJA	SAN PEDRO DE CHUNAN	2025-12-05 15:55:51.884424-05	2025-12-05 15:55:51.884424-05
1129	110333	120430	JUNIN	JAUJA	SAUSA	2025-12-05 15:55:51.884939-05	2025-12-05 15:55:51.884939-05
1130	110328	120431	JUNIN	JAUJA	SINCOS	2025-12-05 15:55:51.884939-05	2025-12-05 15:55:51.884939-05
1131	110329	120432	JUNIN	JAUJA	TUNAN MARCA	2025-12-05 15:55:51.885475-05	2025-12-05 15:55:51.885475-05
1132	110330	120433	JUNIN	JAUJA	YAULI	2025-12-05 15:55:51.886001-05	2025-12-05 15:55:51.886001-05
1133	110402	120502	JUNIN	JUNIN	CARHUAMAYO	2025-12-05 15:55:51.886537-05	2025-12-05 15:55:51.886537-05
1134	110403	120503	JUNIN	JUNIN	ONDORES	2025-12-05 15:55:51.887073-05	2025-12-05 15:55:51.887073-05
1135	110404	120504	JUNIN	JUNIN	ULCUMAYO	2025-12-05 15:55:51.887595-05	2025-12-05 15:55:51.887595-05
1136	110701	120601	JUNIN	SATIPO	SATIPO	2025-12-05 15:55:51.887595-05	2025-12-05 15:55:51.887595-05
1137	110702	120602	JUNIN	SATIPO	COVIRIALI	2025-12-05 15:55:51.888542-05	2025-12-05 15:55:51.888542-05
1138	110703	120603	JUNIN	SATIPO	LLAYLLA	2025-12-05 15:55:51.889065-05	2025-12-05 15:55:51.889065-05
1139	110704	120604	JUNIN	SATIPO	MAZAMARI	2025-12-05 15:55:51.889582-05	2025-12-05 15:55:51.889582-05
1140	110705	120605	JUNIN	SATIPO	PAMPA HERMOSA	2025-12-05 15:55:51.889582-05	2025-12-05 15:55:51.889582-05
1141	110706	120606	JUNIN	SATIPO	PANGOA	2025-12-05 15:55:51.890631-05	2025-12-05 15:55:51.890631-05
1142	110707	120607	JUNIN	SATIPO	RIO NEGRO	2025-12-05 15:55:51.890631-05	2025-12-05 15:55:51.890631-05
1143	110708	120608	JUNIN	SATIPO	RIO TAMBO	2025-12-05 15:55:51.891189-05	2025-12-05 15:55:51.891189-05
1144	110709	120609	JUNIN	SATIPO	VIZCATAN DEL ENE	2025-12-05 15:55:51.891717-05	2025-12-05 15:55:51.891717-05
1145	110501	120701	JUNIN	TARMA	TARMA	2025-12-05 15:55:51.891717-05	2025-12-05 15:55:51.891717-05
1146	110502	120702	JUNIN	TARMA	ACOBAMBA	2025-12-05 15:55:51.89224-05	2025-12-05 15:55:51.89224-05
1147	110503	120703	JUNIN	TARMA	HUARICOLCA	2025-12-05 15:55:51.893292-05	2025-12-05 15:55:51.893292-05
1148	110504	120704	JUNIN	TARMA	HUASAHUASI	2025-12-05 15:55:51.893847-05	2025-12-05 15:55:51.893847-05
1149	110505	120705	JUNIN	TARMA	LA UNION	2025-12-05 15:55:51.894383-05	2025-12-05 15:55:51.894383-05
1150	110506	120706	JUNIN	TARMA	PALCA	2025-12-05 15:55:51.895019-05	2025-12-05 15:55:51.895019-05
1151	110507	120707	JUNIN	TARMA	PALCAMAYO	2025-12-05 15:55:51.895789-05	2025-12-05 15:55:51.895789-05
1152	110508	120708	JUNIN	TARMA	SAN PEDRO DE CAJAS	2025-12-05 15:55:51.896403-05	2025-12-05 15:55:51.896403-05
1153	110509	120709	JUNIN	TARMA	TAPO	2025-12-05 15:55:51.896922-05	2025-12-05 15:55:51.896922-05
1154	110601	120801	JUNIN	YAULI	LA OROYA	2025-12-05 15:55:51.898283-05	2025-12-05 15:55:51.898283-05
1155	110602	120802	JUNIN	YAULI	CHACAPALPA	2025-12-05 15:55:51.898791-05	2025-12-05 15:55:51.898791-05
1156	110603	120803	JUNIN	YAULI	HUAY-HUAY	2025-12-05 15:55:51.899394-05	2025-12-05 15:55:51.899394-05
1157	110604	120804	JUNIN	YAULI	MARCAPOMACOCHA	2025-12-05 15:55:51.899921-05	2025-12-05 15:55:51.899921-05
1158	110605	120805	JUNIN	YAULI	MOROCOCHA	2025-12-05 15:55:51.900382-05	2025-12-05 15:55:51.900382-05
1159	110606	120806	JUNIN	YAULI	PACCHA	2025-12-05 15:55:51.900872-05	2025-12-05 15:55:51.900872-05
1160	110607	120807	JUNIN	YAULI	SANTA BARBARA DE CARHUACAYAN	2025-12-05 15:55:51.900872-05	2025-12-05 15:55:51.900872-05
1161	110610	120808	JUNIN	YAULI	SANTA ROSA DE SACCO	2025-12-05 15:55:51.901658-05	2025-12-05 15:55:51.901658-05
1162	110608	120809	JUNIN	YAULI	SUITUCANCHA	2025-12-05 15:55:51.902188-05	2025-12-05 15:55:51.902188-05
1163	110609	120810	JUNIN	YAULI	YAULI	2025-12-05 15:55:51.902188-05	2025-12-05 15:55:51.902188-05
1164	110901	120901	JUNIN	CHUPACA	CHUPACA	2025-12-05 15:55:51.902188-05	2025-12-05 15:55:51.902188-05
1165	110902	120902	JUNIN	CHUPACA	AHUAC	2025-12-05 15:55:51.902939-05	2025-12-05 15:55:51.902939-05
1166	110903	120903	JUNIN	CHUPACA	CHONGOS BAJO	2025-12-05 15:55:51.903476-05	2025-12-05 15:55:51.903476-05
1167	110904	120904	JUNIN	CHUPACA	HUACHAC	2025-12-05 15:55:51.903476-05	2025-12-05 15:55:51.903476-05
1168	110905	120905	JUNIN	CHUPACA	HUAMANCACA CHICO	2025-12-05 15:55:51.904094-05	2025-12-05 15:55:51.904094-05
1169	110906	120906	JUNIN	CHUPACA	SAN JUAN DE YSCOS	2025-12-05 15:55:51.904649-05	2025-12-05 15:55:51.904649-05
1170	110907	120907	JUNIN	CHUPACA	SAN JUAN DE JARPA	2025-12-05 15:55:51.90526-05	2025-12-05 15:55:51.90526-05
1171	110908	120908	JUNIN	CHUPACA	TRES DE DICIEMBRE	2025-12-05 15:55:51.905859-05	2025-12-05 15:55:51.905859-05
1172	110909	120909	JUNIN	CHUPACA	YANACANCHA	2025-12-05 15:55:51.905859-05	2025-12-05 15:55:51.905859-05
1173	120101	130101	LA LIBERTAD	TRUJILLO	TRUJILLO	2025-12-05 15:55:51.906388-05	2025-12-05 15:55:51.906388-05
1174	120110	130102	LA LIBERTAD	TRUJILLO	EL PORVENIR	2025-12-05 15:55:51.906929-05	2025-12-05 15:55:51.906929-05
1175	120112	130103	LA LIBERTAD	TRUJILLO	FLORENCIA DE MORA	2025-12-05 15:55:51.906929-05	2025-12-05 15:55:51.906929-05
1176	120102	130104	LA LIBERTAD	TRUJILLO	HUANCHACO	2025-12-05 15:55:51.907523-05	2025-12-05 15:55:51.907523-05
1177	120111	130105	LA LIBERTAD	TRUJILLO	LA ESPERANZA	2025-12-05 15:55:51.907523-05	2025-12-05 15:55:51.907523-05
1178	120103	130106	LA LIBERTAD	TRUJILLO	LAREDO	2025-12-05 15:55:51.90808-05	2025-12-05 15:55:51.90808-05
1179	120104	130107	LA LIBERTAD	TRUJILLO	MOCHE	2025-12-05 15:55:51.908647-05	2025-12-05 15:55:51.908647-05
1180	120109	130108	LA LIBERTAD	TRUJILLO	POROTO	2025-12-05 15:55:51.908647-05	2025-12-05 15:55:51.908647-05
1181	120105	130109	LA LIBERTAD	TRUJILLO	SALAVERRY	2025-12-05 15:55:51.909707-05	2025-12-05 15:55:51.909707-05
1182	120106	130110	LA LIBERTAD	TRUJILLO	SIMBAL	2025-12-05 15:55:51.91023-05	2025-12-05 15:55:51.91023-05
1183	120107	130111	LA LIBERTAD	TRUJILLO	VICTOR LARCO HERRERA	2025-12-05 15:55:51.911402-05	2025-12-05 15:55:51.911402-05
1184	120801	130201	LA LIBERTAD	ASCOPE	ASCOPE	2025-12-05 15:55:51.911993-05	2025-12-05 15:55:51.911993-05
1185	120802	130202	LA LIBERTAD	ASCOPE	CHICAMA	2025-12-05 15:55:51.91251-05	2025-12-05 15:55:51.91251-05
1186	120803	130203	LA LIBERTAD	ASCOPE	CHOCOPE	2025-12-05 15:55:51.913692-05	2025-12-05 15:55:51.913692-05
1187	120805	130204	LA LIBERTAD	ASCOPE	MAGDALENA DE CAO	2025-12-05 15:55:51.914937-05	2025-12-05 15:55:51.914937-05
1188	120806	130205	LA LIBERTAD	ASCOPE	PAIJAN	2025-12-05 15:55:51.915694-05	2025-12-05 15:55:51.915694-05
1189	120807	130206	LA LIBERTAD	ASCOPE	RAZURI	2025-12-05 15:55:51.916217-05	2025-12-05 15:55:51.916217-05
1190	120804	130207	LA LIBERTAD	ASCOPE	SANTIAGO DE CAO	2025-12-05 15:55:51.916744-05	2025-12-05 15:55:51.916744-05
1191	120808	130208	LA LIBERTAD	ASCOPE	CASA GRANDE	2025-12-05 15:55:51.917284-05	2025-12-05 15:55:51.917284-05
1192	120201	130301	LA LIBERTAD	BOLIVAR	BOLIVAR	2025-12-05 15:55:51.917876-05	2025-12-05 15:55:51.917876-05
1193	120202	130302	LA LIBERTAD	BOLIVAR	BAMBAMARCA	2025-12-05 15:55:51.917876-05	2025-12-05 15:55:51.917876-05
1194	120203	130303	LA LIBERTAD	BOLIVAR	CONDORMARCA	2025-12-05 15:55:51.918395-05	2025-12-05 15:55:51.918395-05
1195	120204	130304	LA LIBERTAD	BOLIVAR	LONGOTEA	2025-12-05 15:55:51.918918-05	2025-12-05 15:55:51.918918-05
1196	120206	130305	LA LIBERTAD	BOLIVAR	UCHUMARCA	2025-12-05 15:55:51.919464-05	2025-12-05 15:55:51.919464-05
1197	120205	130306	LA LIBERTAD	BOLIVAR	UCUNCHA	2025-12-05 15:55:51.919464-05	2025-12-05 15:55:51.919464-05
1198	120901	130401	LA LIBERTAD	CHEPEN	CHEPEN	2025-12-05 15:55:51.920636-05	2025-12-05 15:55:51.920636-05
1199	120902	130402	LA LIBERTAD	CHEPEN	PACANGA	2025-12-05 15:55:51.921162-05	2025-12-05 15:55:51.921162-05
1200	120903	130403	LA LIBERTAD	CHEPEN	PUEBLO NUEVO	2025-12-05 15:55:51.921802-05	2025-12-05 15:55:51.921802-05
1201	121001	130501	LA LIBERTAD	JULCAN	JULCAN	2025-12-05 15:55:51.921802-05	2025-12-05 15:55:51.921802-05
1202	121003	130502	LA LIBERTAD	JULCAN	CALAMARCA	2025-12-05 15:55:51.922331-05	2025-12-05 15:55:51.922331-05
1203	121002	130503	LA LIBERTAD	JULCAN	CARABAMBA	2025-12-05 15:55:51.922861-05	2025-12-05 15:55:51.922861-05
1204	121004	130504	LA LIBERTAD	JULCAN	HUASO	2025-12-05 15:55:51.922861-05	2025-12-05 15:55:51.922861-05
1205	120401	130601	LA LIBERTAD	OTUZCO	OTUZCO	2025-12-05 15:55:51.923451-05	2025-12-05 15:55:51.923451-05
1206	120402	130602	LA LIBERTAD	OTUZCO	AGALLPAMPA	2025-12-05 15:55:51.923975-05	2025-12-05 15:55:51.923975-05
1207	120403	130604	LA LIBERTAD	OTUZCO	CHARAT	2025-12-05 15:55:51.924506-05	2025-12-05 15:55:51.924506-05
1208	120404	130605	LA LIBERTAD	OTUZCO	HUARANCHAL	2025-12-05 15:55:51.924506-05	2025-12-05 15:55:51.924506-05
1209	120405	130606	LA LIBERTAD	OTUZCO	LA CUESTA	2025-12-05 15:55:51.92503-05	2025-12-05 15:55:51.92503-05
1210	120413	130608	LA LIBERTAD	OTUZCO	MACHE	2025-12-05 15:55:51.92557-05	2025-12-05 15:55:51.92557-05
1211	120408	130610	LA LIBERTAD	OTUZCO	PARANDAY	2025-12-05 15:55:51.926082-05	2025-12-05 15:55:51.926082-05
1212	120409	130611	LA LIBERTAD	OTUZCO	SALPO	2025-12-05 15:55:51.926786-05	2025-12-05 15:55:51.926786-05
1213	120410	130613	LA LIBERTAD	OTUZCO	SINSICAP	2025-12-05 15:55:51.927502-05	2025-12-05 15:55:51.927502-05
1214	120411	130614	LA LIBERTAD	OTUZCO	USQUIL	2025-12-05 15:55:51.928139-05	2025-12-05 15:55:51.928139-05
1215	120501	130701	LA LIBERTAD	PACASMAYO	SAN PEDRO DE LLOC	2025-12-05 15:55:51.929056-05	2025-12-05 15:55:51.929056-05
1216	120503	130702	LA LIBERTAD	PACASMAYO	GUADALUPE	2025-12-05 15:55:51.92956-05	2025-12-05 15:55:51.92956-05
1217	120504	130703	LA LIBERTAD	PACASMAYO	JEQUETEPEQUE	2025-12-05 15:55:51.930105-05	2025-12-05 15:55:51.930105-05
1218	120506	130704	LA LIBERTAD	PACASMAYO	PACASMAYO	2025-12-05 15:55:51.930716-05	2025-12-05 15:55:51.930716-05
1219	120508	130705	LA LIBERTAD	PACASMAYO	SAN JOSE	2025-12-05 15:55:51.931229-05	2025-12-05 15:55:51.931229-05
1220	120601	130801	LA LIBERTAD	PATAZ	TAYABAMBA	2025-12-05 15:55:51.932478-05	2025-12-05 15:55:51.932478-05
1221	120602	130802	LA LIBERTAD	PATAZ	BULDIBUYO	2025-12-05 15:55:51.93316-05	2025-12-05 15:55:51.93316-05
1222	120603	130803	LA LIBERTAD	PATAZ	CHILLIA	2025-12-05 15:55:51.933774-05	2025-12-05 15:55:51.933774-05
1223	120605	130804	LA LIBERTAD	PATAZ	HUANCASPATA	2025-12-05 15:55:51.934324-05	2025-12-05 15:55:51.934324-05
1224	120604	130805	LA LIBERTAD	PATAZ	HUAYLILLAS	2025-12-05 15:55:51.934324-05	2025-12-05 15:55:51.934324-05
1225	120606	130806	LA LIBERTAD	PATAZ	HUAYO	2025-12-05 15:55:51.934883-05	2025-12-05 15:55:51.934883-05
1226	120607	130807	LA LIBERTAD	PATAZ	ONGON	2025-12-05 15:55:51.935405-05	2025-12-05 15:55:51.935405-05
1227	120608	130808	LA LIBERTAD	PATAZ	PARCOY	2025-12-05 15:55:51.935405-05	2025-12-05 15:55:51.935405-05
1228	120609	130809	LA LIBERTAD	PATAZ	PATAZ	2025-12-05 15:55:51.935966-05	2025-12-05 15:55:51.935966-05
1229	120610	130810	LA LIBERTAD	PATAZ	PIAS	2025-12-05 15:55:51.9367-05	2025-12-05 15:55:51.9367-05
1230	120613	130811	LA LIBERTAD	PATAZ	SANTIAGO DE CHALLAS	2025-12-05 15:55:51.937306-05	2025-12-05 15:55:51.937306-05
1231	120611	130812	LA LIBERTAD	PATAZ	TAURIJA	2025-12-05 15:55:51.937833-05	2025-12-05 15:55:51.937833-05
1232	120612	130813	LA LIBERTAD	PATAZ	URPAY	2025-12-05 15:55:51.937833-05	2025-12-05 15:55:51.937833-05
1233	120301	130901	LA LIBERTAD	SANCHEZ CARRION	HUAMACHUCO	2025-12-05 15:55:51.938373-05	2025-12-05 15:55:51.938373-05
1234	120304	130902	LA LIBERTAD	SANCHEZ CARRION	CHUGAY	2025-12-05 15:55:51.938373-05	2025-12-05 15:55:51.938373-05
1235	120302	130903	LA LIBERTAD	SANCHEZ CARRION	COCHORCO	2025-12-05 15:55:51.938965-05	2025-12-05 15:55:51.938965-05
1236	120303	130904	LA LIBERTAD	SANCHEZ CARRION	CURGOS	2025-12-05 15:55:51.939502-05	2025-12-05 15:55:51.939502-05
1237	120305	130905	LA LIBERTAD	SANCHEZ CARRION	MARCABAL	2025-12-05 15:55:51.939502-05	2025-12-05 15:55:51.939502-05
1238	120306	130906	LA LIBERTAD	SANCHEZ CARRION	SANAGORAN	2025-12-05 15:55:51.940042-05	2025-12-05 15:55:51.940042-05
1239	120307	130907	LA LIBERTAD	SANCHEZ CARRION	SARIN	2025-12-05 15:55:51.940597-05	2025-12-05 15:55:51.940597-05
1240	120308	130908	LA LIBERTAD	SANCHEZ CARRION	SARTIMBAMBA	2025-12-05 15:55:51.941121-05	2025-12-05 15:55:51.941121-05
1241	120701	131001	LA LIBERTAD	SANTIAGO DE CHUCO	SANTIAGO DE CHUCO	2025-12-05 15:55:51.941601-05	2025-12-05 15:55:51.941601-05
1242	120708	131002	LA LIBERTAD	SANTIAGO DE CHUCO	ANGASMARCA	2025-12-05 15:55:51.942125-05	2025-12-05 15:55:51.942125-05
1243	120702	131003	LA LIBERTAD	SANTIAGO DE CHUCO	CACHICADAN	2025-12-05 15:55:51.942656-05	2025-12-05 15:55:51.942656-05
1244	120703	131004	LA LIBERTAD	SANTIAGO DE CHUCO	MOLLEBAMBA	2025-12-05 15:55:51.943282-05	2025-12-05 15:55:51.943282-05
1245	120704	131005	LA LIBERTAD	SANTIAGO DE CHUCO	MOLLEPATA	2025-12-05 15:55:51.943945-05	2025-12-05 15:55:51.943945-05
1246	120705	131006	LA LIBERTAD	SANTIAGO DE CHUCO	QUIRUVILCA	2025-12-05 15:55:51.944608-05	2025-12-05 15:55:51.944608-05
1247	120706	131007	LA LIBERTAD	SANTIAGO DE CHUCO	SANTA CRUZ DE CHUCA	2025-12-05 15:55:51.945214-05	2025-12-05 15:55:51.945214-05
1248	120707	131008	LA LIBERTAD	SANTIAGO DE CHUCO	SITABAMBA	2025-12-05 15:55:51.945732-05	2025-12-05 15:55:51.945732-05
1249	121101	131101	LA LIBERTAD	GRAN CHIMU	CASCAS	2025-12-05 15:55:51.946356-05	2025-12-05 15:55:51.946356-05
1250	121102	131102	LA LIBERTAD	GRAN CHIMU	LUCMA	2025-12-05 15:55:51.947332-05	2025-12-05 15:55:51.947332-05
1251	121103	131103	LA LIBERTAD	GRAN CHIMU	MARMOT	2025-12-05 15:55:51.947879-05	2025-12-05 15:55:51.947879-05
1252	121104	131104	LA LIBERTAD	GRAN CHIMU	SAYAPULLO	2025-12-05 15:55:51.948542-05	2025-12-05 15:55:51.948542-05
1253	121201	131201	LA LIBERTAD	VIRU	VIRU	2025-12-05 15:55:51.949123-05	2025-12-05 15:55:51.949123-05
1254	121202	131202	LA LIBERTAD	VIRU	CHAO	2025-12-05 15:55:51.94967-05	2025-12-05 15:55:51.94967-05
1255	121203	131203	LA LIBERTAD	VIRU	GUADALUPITO	2025-12-05 15:55:51.950213-05	2025-12-05 15:55:51.950213-05
1256	130101	140101	LAMBAYEQUE	CHICLAYO	CHICLAYO	2025-12-05 15:55:51.950213-05	2025-12-05 15:55:51.950213-05
1257	130102	140102	LAMBAYEQUE	CHICLAYO	CHONGOYAPE	2025-12-05 15:55:51.951048-05	2025-12-05 15:55:51.951048-05
1258	130103	140103	LAMBAYEQUE	CHICLAYO	ETEN	2025-12-05 15:55:51.95149-05	2025-12-05 15:55:51.95149-05
1259	130104	140104	LAMBAYEQUE	CHICLAYO	ETEN PUERTO	2025-12-05 15:55:51.952044-05	2025-12-05 15:55:51.952044-05
1260	130112	140105	LAMBAYEQUE	CHICLAYO	JOSE LEONARDO ORTIZ	2025-12-05 15:55:51.952044-05	2025-12-05 15:55:51.952044-05
1261	130115	140106	LAMBAYEQUE	CHICLAYO	LA VICTORIA	2025-12-05 15:55:51.952562-05	2025-12-05 15:55:51.952562-05
1262	130105	140107	LAMBAYEQUE	CHICLAYO	LAGUNAS	2025-12-05 15:55:51.953098-05	2025-12-05 15:55:51.953098-05
1263	130106	140108	LAMBAYEQUE	CHICLAYO	MONSEFU	2025-12-05 15:55:51.953659-05	2025-12-05 15:55:51.953659-05
1264	130107	140109	LAMBAYEQUE	CHICLAYO	NUEVA ARICA	2025-12-05 15:55:51.954271-05	2025-12-05 15:55:51.954271-05
1265	130108	140110	LAMBAYEQUE	CHICLAYO	OYOTUN	2025-12-05 15:55:51.95483-05	2025-12-05 15:55:51.95483-05
1266	130109	140111	LAMBAYEQUE	CHICLAYO	PICSI	2025-12-05 15:55:51.955361-05	2025-12-05 15:55:51.955361-05
1267	130110	140112	LAMBAYEQUE	CHICLAYO	PIMENTEL	2025-12-05 15:55:51.955361-05	2025-12-05 15:55:51.955361-05
1268	130111	140113	LAMBAYEQUE	CHICLAYO	REQUE	2025-12-05 15:55:51.956131-05	2025-12-05 15:55:51.956131-05
1269	130113	140114	LAMBAYEQUE	CHICLAYO	SANTA ROSA	2025-12-05 15:55:51.956818-05	2025-12-05 15:55:51.956818-05
1270	130114	140115	LAMBAYEQUE	CHICLAYO	SAÑA	2025-12-05 15:55:51.957325-05	2025-12-05 15:55:51.957325-05
1271	130116	140116	LAMBAYEQUE	CHICLAYO	CAYALTI	2025-12-05 15:55:51.957849-05	2025-12-05 15:55:51.957849-05
1272	130117	140117	LAMBAYEQUE	CHICLAYO	PATAPO	2025-12-05 15:55:51.958372-05	2025-12-05 15:55:51.958372-05
1273	130118	140118	LAMBAYEQUE	CHICLAYO	POMALCA	2025-12-05 15:55:51.958372-05	2025-12-05 15:55:51.958372-05
1274	130119	140119	LAMBAYEQUE	CHICLAYO	PUCALA	2025-12-05 15:55:51.959411-05	2025-12-05 15:55:51.959411-05
1275	130120	140120	LAMBAYEQUE	CHICLAYO	TUMAN	2025-12-05 15:55:51.960022-05	2025-12-05 15:55:51.960022-05
1276	130201	140201	LAMBAYEQUE	FERREÑAFE	FERREÑAFE	2025-12-05 15:55:51.960551-05	2025-12-05 15:55:51.960551-05
1277	130203	140202	LAMBAYEQUE	FERREÑAFE	CAÑARIS	2025-12-05 15:55:51.961171-05	2025-12-05 15:55:51.961171-05
1278	130202	140203	LAMBAYEQUE	FERREÑAFE	INCAHUASI	2025-12-05 15:55:51.962211-05	2025-12-05 15:55:51.962211-05
1279	130206	140204	LAMBAYEQUE	FERREÑAFE	MANUEL ANTONIO MESONES MURO	2025-12-05 15:55:51.962737-05	2025-12-05 15:55:51.962737-05
1280	130204	140205	LAMBAYEQUE	FERREÑAFE	PITIPO	2025-12-05 15:55:51.962737-05	2025-12-05 15:55:51.962737-05
1281	130205	140206	LAMBAYEQUE	FERREÑAFE	PUEBLO NUEVO	2025-12-05 15:55:51.963258-05	2025-12-05 15:55:51.963258-05
1282	130301	140301	LAMBAYEQUE	LAMBAYEQUE	LAMBAYEQUE	2025-12-05 15:55:51.964139-05	2025-12-05 15:55:51.964139-05
1283	130302	140302	LAMBAYEQUE	LAMBAYEQUE	CHOCHOPE	2025-12-05 15:55:51.965295-05	2025-12-05 15:55:51.965295-05
1284	130303	140303	LAMBAYEQUE	LAMBAYEQUE	ILLIMO	2025-12-05 15:55:51.965889-05	2025-12-05 15:55:51.965889-05
1285	130304	140304	LAMBAYEQUE	LAMBAYEQUE	JAYANCA	2025-12-05 15:55:51.965889-05	2025-12-05 15:55:51.965889-05
1286	130305	140305	LAMBAYEQUE	LAMBAYEQUE	MOCHUMI	2025-12-05 15:55:51.966475-05	2025-12-05 15:55:51.966475-05
1287	130306	140306	LAMBAYEQUE	LAMBAYEQUE	MORROPE	2025-12-05 15:55:51.96707-05	2025-12-05 15:55:51.96707-05
1288	130307	140307	LAMBAYEQUE	LAMBAYEQUE	MOTUPE	2025-12-05 15:55:51.967671-05	2025-12-05 15:55:51.967671-05
1289	130308	140308	LAMBAYEQUE	LAMBAYEQUE	OLMOS	2025-12-05 15:55:51.967972-05	2025-12-05 15:55:51.967972-05
1290	130309	140309	LAMBAYEQUE	LAMBAYEQUE	PACORA	2025-12-05 15:55:51.968611-05	2025-12-05 15:55:51.968611-05
1291	130310	140310	LAMBAYEQUE	LAMBAYEQUE	SALAS	2025-12-05 15:55:51.968611-05	2025-12-05 15:55:51.968611-05
1292	130311	140311	LAMBAYEQUE	LAMBAYEQUE	SAN JOSE	2025-12-05 15:55:51.969183-05	2025-12-05 15:55:51.969183-05
1293	130312	140312	LAMBAYEQUE	LAMBAYEQUE	TUCUME	2025-12-05 15:55:51.96971-05	2025-12-05 15:55:51.96971-05
1294	140101	150101	LIMA	LIMA	LIMA	2025-12-05 15:55:51.96971-05	2025-12-05 15:55:51.96971-05
1295	140102	150102	LIMA	LIMA	ANCON	2025-12-05 15:55:51.970234-05	2025-12-05 15:55:51.970234-05
1296	140103	150103	LIMA	LIMA	ATE	2025-12-05 15:55:51.970799-05	2025-12-05 15:55:51.970799-05
1297	140125	150104	LIMA	LIMA	BARRANCO	2025-12-05 15:55:51.971325-05	2025-12-05 15:55:51.971325-05
1298	140104	150105	LIMA	LIMA	BREÑA	2025-12-05 15:55:51.971844-05	2025-12-05 15:55:51.971844-05
1299	140105	150106	LIMA	LIMA	CARABAYLLO	2025-12-05 15:55:51.972421-05	2025-12-05 15:55:51.972421-05
1300	140107	150107	LIMA	LIMA	CHACLACAYO	2025-12-05 15:55:51.973143-05	2025-12-05 15:55:51.973143-05
1301	140108	150108	LIMA	LIMA	CHORRILLOS	2025-12-05 15:55:51.973686-05	2025-12-05 15:55:51.973686-05
1302	140139	150109	LIMA	LIMA	CIENEGUILLA	2025-12-05 15:55:51.973686-05	2025-12-05 15:55:51.973686-05
1303	140106	150110	LIMA	LIMA	COMAS	2025-12-05 15:55:51.97421-05	2025-12-05 15:55:51.97421-05
1304	140135	150111	LIMA	LIMA	EL AGUSTINO	2025-12-05 15:55:51.974734-05	2025-12-05 15:55:51.974734-05
1305	140134	150112	LIMA	LIMA	INDEPENDENCIA	2025-12-05 15:55:51.974734-05	2025-12-05 15:55:51.974734-05
1306	140133	150113	LIMA	LIMA	JESUS MARIA	2025-12-05 15:55:51.975276-05	2025-12-05 15:55:51.975276-05
1307	140110	150114	LIMA	LIMA	LA MOLINA	2025-12-05 15:55:51.975793-05	2025-12-05 15:55:51.975793-05
1308	140109	150115	LIMA	LIMA	LA VICTORIA	2025-12-05 15:55:51.976378-05	2025-12-05 15:55:51.976378-05
1309	140111	150116	LIMA	LIMA	LINCE	2025-12-05 15:55:51.97702-05	2025-12-05 15:55:51.97702-05
1310	140142	150117	LIMA	LIMA	LOS OLIVOS	2025-12-05 15:55:51.97813-05	2025-12-05 15:55:51.97813-05
1311	140112	150118	LIMA	LIMA	LURIGANCHO	2025-12-05 15:55:51.978755-05	2025-12-05 15:55:51.978755-05
1312	140113	150119	LIMA	LIMA	LURIN	2025-12-05 15:55:51.979301-05	2025-12-05 15:55:51.979301-05
1313	140114	150120	LIMA	LIMA	MAGDALENA DEL MAR	2025-12-05 15:55:51.979301-05	2025-12-05 15:55:51.979301-05
1314	140117	150121	LIMA	LIMA	PUEBLO LIBRE	2025-12-05 15:55:51.980362-05	2025-12-05 15:55:51.980362-05
1315	140115	150122	LIMA	LIMA	MIRAFLORES	2025-12-05 15:55:51.980833-05	2025-12-05 15:55:51.980833-05
1316	140116	150123	LIMA	LIMA	PACHACAMAC	2025-12-05 15:55:51.981605-05	2025-12-05 15:55:51.981605-05
1317	140118	150124	LIMA	LIMA	PUCUSANA	2025-12-05 15:55:51.982457-05	2025-12-05 15:55:51.982457-05
1318	140119	150125	LIMA	LIMA	PUENTE PIEDRA	2025-12-05 15:55:51.982867-05	2025-12-05 15:55:51.982867-05
1319	140120	150126	LIMA	LIMA	PUNTA HERMOSA	2025-12-05 15:55:51.983393-05	2025-12-05 15:55:51.983393-05
1320	140121	150127	LIMA	LIMA	PUNTA NEGRA	2025-12-05 15:55:51.983912-05	2025-12-05 15:55:51.983912-05
1321	140122	150128	LIMA	LIMA	RIMAC	2025-12-05 15:55:51.983912-05	2025-12-05 15:55:51.983912-05
1322	140123	150129	LIMA	LIMA	SAN BARTOLO	2025-12-05 15:55:51.984528-05	2025-12-05 15:55:51.984528-05
1323	140140	150130	LIMA	LIMA	SAN BORJA	2025-12-05 15:55:51.985049-05	2025-12-05 15:55:51.985049-05
1324	140124	150131	LIMA	LIMA	SAN ISIDRO	2025-12-05 15:55:51.985049-05	2025-12-05 15:55:51.985049-05
1325	140137	150132	LIMA	LIMA	SAN JUAN DE LURIGANCHO	2025-12-05 15:55:51.985632-05	2025-12-05 15:55:51.985632-05
1326	140136	150133	LIMA	LIMA	SAN JUAN DE MIRAFLORES	2025-12-05 15:55:51.986172-05	2025-12-05 15:55:51.986172-05
1327	140138	150134	LIMA	LIMA	SAN LUIS	2025-12-05 15:55:51.986172-05	2025-12-05 15:55:51.986172-05
1328	140126	150135	LIMA	LIMA	SAN MARTIN DE PORRES	2025-12-05 15:55:51.987183-05	2025-12-05 15:55:51.987183-05
1329	140127	150136	LIMA	LIMA	SAN MIGUEL	2025-12-05 15:55:51.987979-05	2025-12-05 15:55:51.987979-05
1330	140143	150137	LIMA	LIMA	SANTA ANITA	2025-12-05 15:55:51.987979-05	2025-12-05 15:55:51.987979-05
1331	140128	150138	LIMA	LIMA	SANTA MARIA DEL MAR	2025-12-05 15:55:51.98851-05	2025-12-05 15:55:51.98851-05
1332	140129	150139	LIMA	LIMA	SANTA ROSA	2025-12-05 15:55:51.989024-05	2025-12-05 15:55:51.989024-05
1333	140130	150140	LIMA	LIMA	SANTIAGO DE SURCO	2025-12-05 15:55:51.989539-05	2025-12-05 15:55:51.989539-05
1334	140131	150141	LIMA	LIMA	SURQUILLO	2025-12-05 15:55:51.989539-05	2025-12-05 15:55:51.989539-05
1335	140141	150142	LIMA	LIMA	VILLA EL SALVADOR	2025-12-05 15:55:51.990067-05	2025-12-05 15:55:51.990067-05
1336	140132	150143	LIMA	LIMA	VILLA MARIA DEL TRIUNFO	2025-12-05 15:55:51.990586-05	2025-12-05 15:55:51.990586-05
1337	140144	150144	LIMA	LIMA	SANTA MARIA DE HUACHIPA	2025-12-05 15:55:51.991118-05	2025-12-05 15:55:51.991118-05
1338	140901	150201	LIMA	BARRANCA	BARRANCA	2025-12-05 15:55:51.991118-05	2025-12-05 15:55:51.991118-05
1339	140902	150202	LIMA	BARRANCA	PARAMONGA	2025-12-05 15:55:51.991636-05	2025-12-05 15:55:51.991636-05
1340	140903	150203	LIMA	BARRANCA	PATIVILCA	2025-12-05 15:55:51.992075-05	2025-12-05 15:55:51.992075-05
1341	140904	150204	LIMA	BARRANCA	SUPE	2025-12-05 15:55:51.992859-05	2025-12-05 15:55:51.992859-05
1342	140905	150205	LIMA	BARRANCA	SUPE PUERTO	2025-12-05 15:55:51.994055-05	2025-12-05 15:55:51.994055-05
1343	140201	150301	LIMA	CAJATAMBO	CAJATAMBO	2025-12-05 15:55:51.994656-05	2025-12-05 15:55:51.994656-05
1344	140205	150302	LIMA	CAJATAMBO	COPA	2025-12-05 15:55:51.995178-05	2025-12-05 15:55:51.995178-05
1345	140206	150303	LIMA	CAJATAMBO	GORGOR	2025-12-05 15:55:51.995709-05	2025-12-05 15:55:51.995709-05
1346	140207	150304	LIMA	CAJATAMBO	HUANCAPON	2025-12-05 15:55:51.996228-05	2025-12-05 15:55:51.996228-05
1347	140208	150305	LIMA	CAJATAMBO	MANAS	2025-12-05 15:55:51.996751-05	2025-12-05 15:55:51.996751-05
1348	140301	150401	LIMA	CANTA	CANTA	2025-12-05 15:55:51.997278-05	2025-12-05 15:55:51.997278-05
1349	140302	150402	LIMA	CANTA	ARAHUAY	2025-12-05 15:55:51.998208-05	2025-12-05 15:55:51.998208-05
1350	140303	150403	LIMA	CANTA	HUAMANTANGA	2025-12-05 15:55:51.99931-05	2025-12-05 15:55:51.99931-05
1351	140304	150404	LIMA	CANTA	HUAROS	2025-12-05 15:55:51.999827-05	2025-12-05 15:55:51.999827-05
1352	140305	150405	LIMA	CANTA	LACHAQUI	2025-12-05 15:55:52.000361-05	2025-12-05 15:55:52.000361-05
1353	140306	150406	LIMA	CANTA	SAN BUENAVENTURA	2025-12-05 15:55:52.000361-05	2025-12-05 15:55:52.000361-05
1354	140307	150407	LIMA	CANTA	SANTA ROSA DE QUIVES	2025-12-05 15:55:52.001011-05	2025-12-05 15:55:52.001011-05
1355	140401	150501	LIMA	CAÑETE	SAN VICENTE DE CAÑETE	2025-12-05 15:55:52.001552-05	2025-12-05 15:55:52.001552-05
1356	140416	150502	LIMA	CAÑETE	ASIA	2025-12-05 15:55:52.002122-05	2025-12-05 15:55:52.002122-05
1357	140402	150503	LIMA	CAÑETE	CALANGO	2025-12-05 15:55:52.002742-05	2025-12-05 15:55:52.002742-05
1358	140403	150504	LIMA	CAÑETE	CERRO AZUL	2025-12-05 15:55:52.002742-05	2025-12-05 15:55:52.002742-05
1359	140405	150505	LIMA	CAÑETE	CHILCA	2025-12-05 15:55:52.003318-05	2025-12-05 15:55:52.003318-05
1360	140404	150506	LIMA	CAÑETE	COAYLLO	2025-12-05 15:55:52.003931-05	2025-12-05 15:55:52.003931-05
1361	140406	150507	LIMA	CAÑETE	IMPERIAL	2025-12-05 15:55:52.003931-05	2025-12-05 15:55:52.003931-05
1362	140407	150508	LIMA	CAÑETE	LUNAHUANA	2025-12-05 15:55:52.004495-05	2025-12-05 15:55:52.004495-05
1363	140408	150509	LIMA	CAÑETE	MALA	2025-12-05 15:55:52.005044-05	2025-12-05 15:55:52.005044-05
1364	140409	150510	LIMA	CAÑETE	NUEVO IMPERIAL	2025-12-05 15:55:52.005044-05	2025-12-05 15:55:52.005044-05
1365	140410	150511	LIMA	CAÑETE	PACARAN	2025-12-05 15:55:52.005562-05	2025-12-05 15:55:52.005562-05
1366	140411	150512	LIMA	CAÑETE	QUILMANA	2025-12-05 15:55:52.006075-05	2025-12-05 15:55:52.006075-05
1367	140412	150513	LIMA	CAÑETE	SAN ANTONIO	2025-12-05 15:55:52.006075-05	2025-12-05 15:55:52.006075-05
1368	140413	150514	LIMA	CAÑETE	SAN LUIS	2025-12-05 15:55:52.006598-05	2025-12-05 15:55:52.006598-05
1369	140414	150515	LIMA	CAÑETE	SANTA CRUZ DE FLORES	2025-12-05 15:55:52.006984-05	2025-12-05 15:55:52.006984-05
1370	140415	150516	LIMA	CAÑETE	ZUÑIGA	2025-12-05 15:55:52.007513-05	2025-12-05 15:55:52.007513-05
1371	140801	150601	LIMA	HUARAL	HUARAL	2025-12-05 15:55:52.00808-05	2025-12-05 15:55:52.00808-05
1372	140802	150602	LIMA	HUARAL	ATAVILLOS ALTO	2025-12-05 15:55:52.008664-05	2025-12-05 15:55:52.008664-05
1373	140803	150603	LIMA	HUARAL	ATAVILLOS BAJO	2025-12-05 15:55:52.009075-05	2025-12-05 15:55:52.009075-05
1374	140804	150604	LIMA	HUARAL	AUCALLAMA	2025-12-05 15:55:52.009694-05	2025-12-05 15:55:52.009694-05
1375	140805	150605	LIMA	HUARAL	CHANCAY	2025-12-05 15:55:52.010786-05	2025-12-05 15:55:52.010786-05
1376	140806	150606	LIMA	HUARAL	IHUARI	2025-12-05 15:55:52.011356-05	2025-12-05 15:55:52.011356-05
1377	140807	150607	LIMA	HUARAL	LAMPIAN	2025-12-05 15:55:52.011877-05	2025-12-05 15:55:52.011877-05
1378	140808	150608	LIMA	HUARAL	PACARAOS	2025-12-05 15:55:52.012412-05	2025-12-05 15:55:52.012412-05
1379	140809	150609	LIMA	HUARAL	SAN MIGUEL DE ACOS	2025-12-05 15:55:52.0129-05	2025-12-05 15:55:52.0129-05
1380	140811	150610	LIMA	HUARAL	SANTA CRUZ DE ANDAMARCA	2025-12-05 15:55:52.013443-05	2025-12-05 15:55:52.013443-05
1381	140812	150611	LIMA	HUARAL	SUMBILCA	2025-12-05 15:55:52.013443-05	2025-12-05 15:55:52.013443-05
1382	140810	150612	LIMA	HUARAL	VEINTISIETE DE NOVIEMBRE	2025-12-05 15:55:52.01403-05	2025-12-05 15:55:52.01403-05
1383	140601	150701	LIMA	HUAROCHIRI	MATUCANA	2025-12-05 15:55:52.014552-05	2025-12-05 15:55:52.014552-05
1384	140602	150702	LIMA	HUAROCHIRI	ANTIOQUIA	2025-12-05 15:55:52.014552-05	2025-12-05 15:55:52.014552-05
1385	140603	150703	LIMA	HUAROCHIRI	CALLAHUANCA	2025-12-05 15:55:52.015062-05	2025-12-05 15:55:52.015062-05
1386	140604	150704	LIMA	HUAROCHIRI	CARAMPOMA	2025-12-05 15:55:52.01562-05	2025-12-05 15:55:52.01562-05
1387	140607	150705	LIMA	HUAROCHIRI	CHICLA	2025-12-05 15:55:52.01562-05	2025-12-05 15:55:52.01562-05
1388	140606	150706	LIMA	HUAROCHIRI	CUENCA	2025-12-05 15:55:52.016181-05	2025-12-05 15:55:52.016181-05
1389	140630	150707	LIMA	HUAROCHIRI	HUACHUPAMPA	2025-12-05 15:55:52.016701-05	2025-12-05 15:55:52.016701-05
1390	140608	150708	LIMA	HUAROCHIRI	HUANZA	2025-12-05 15:55:52.016701-05	2025-12-05 15:55:52.016701-05
1391	140609	150709	LIMA	HUAROCHIRI	HUAROCHIRI	2025-12-05 15:55:52.017235-05	2025-12-05 15:55:52.017235-05
1392	140610	150710	LIMA	HUAROCHIRI	LAHUAYTAMBO	2025-12-05 15:55:52.017754-05	2025-12-05 15:55:52.017754-05
1393	140611	150711	LIMA	HUAROCHIRI	LANGA	2025-12-05 15:55:52.018539-05	2025-12-05 15:55:52.018539-05
1394	140631	150712	LIMA	HUAROCHIRI	LARAOS	2025-12-05 15:55:52.018895-05	2025-12-05 15:55:52.018895-05
1395	140612	150713	LIMA	HUAROCHIRI	MARIATANA	2025-12-05 15:55:52.019451-05	2025-12-05 15:55:52.019451-05
1396	140613	150714	LIMA	HUAROCHIRI	RICARDO PALMA	2025-12-05 15:55:52.020021-05	2025-12-05 15:55:52.020021-05
1397	140614	150715	LIMA	HUAROCHIRI	SAN ANDRES DE TUPICOCHA	2025-12-05 15:55:52.020021-05	2025-12-05 15:55:52.020021-05
1398	140615	150716	LIMA	HUAROCHIRI	SAN ANTONIO	2025-12-05 15:55:52.020589-05	2025-12-05 15:55:52.020589-05
1399	140616	150717	LIMA	HUAROCHIRI	SAN BARTOLOME	2025-12-05 15:55:52.021141-05	2025-12-05 15:55:52.021141-05
1400	140617	150718	LIMA	HUAROCHIRI	SAN DAMIAN	2025-12-05 15:55:52.021141-05	2025-12-05 15:55:52.021141-05
1401	140632	150719	LIMA	HUAROCHIRI	SAN JUAN DE IRIS	2025-12-05 15:55:52.02166-05	2025-12-05 15:55:52.02166-05
1402	140619	150720	LIMA	HUAROCHIRI	SAN JUAN DE TANTARANCHE	2025-12-05 15:55:52.022175-05	2025-12-05 15:55:52.022175-05
1403	140620	150721	LIMA	HUAROCHIRI	SAN LORENZO DE QUINTI	2025-12-05 15:55:52.022175-05	2025-12-05 15:55:52.022175-05
1404	140621	150722	LIMA	HUAROCHIRI	SAN MATEO	2025-12-05 15:55:52.022695-05	2025-12-05 15:55:52.022695-05
1405	140622	150723	LIMA	HUAROCHIRI	SAN MATEO DE OTAO	2025-12-05 15:55:52.023275-05	2025-12-05 15:55:52.023275-05
1406	140605	150724	LIMA	HUAROCHIRI	SAN PEDRO DE CASTA	2025-12-05 15:55:52.023939-05	2025-12-05 15:55:52.023939-05
1407	140623	150725	LIMA	HUAROCHIRI	SAN PEDRO DE HUANCAYRE	2025-12-05 15:55:52.024495-05	2025-12-05 15:55:52.024495-05
1408	140618	150726	LIMA	HUAROCHIRI	SANGALLAYA	2025-12-05 15:55:52.025016-05	2025-12-05 15:55:52.025016-05
1409	140624	150727	LIMA	HUAROCHIRI	SANTA CRUZ DE COCACHACRA	2025-12-05 15:55:52.025016-05	2025-12-05 15:55:52.025016-05
1410	140625	150728	LIMA	HUAROCHIRI	SANTA EULALIA	2025-12-05 15:55:52.025566-05	2025-12-05 15:55:52.025566-05
1411	140626	150729	LIMA	HUAROCHIRI	SANTIAGO DE ANCHUCAYA	2025-12-05 15:55:52.026228-05	2025-12-05 15:55:52.026228-05
1412	140627	150730	LIMA	HUAROCHIRI	SANTIAGO DE TUNA	2025-12-05 15:55:52.026917-05	2025-12-05 15:55:52.026917-05
1413	140628	150731	LIMA	HUAROCHIRI	SANTO DOMINGO DE LOS OLLEROS	2025-12-05 15:55:52.027977-05	2025-12-05 15:55:52.027977-05
1414	140629	150732	LIMA	HUAROCHIRI	SURCO	2025-12-05 15:55:52.028553-05	2025-12-05 15:55:52.028553-05
1415	140501	150801	LIMA	HUAURA	HUACHO	2025-12-05 15:55:52.029097-05	2025-12-05 15:55:52.029097-05
1416	140504	150803	LIMA	HUAURA	CALETA DE CARQUIN	2025-12-05 15:55:52.029624-05	2025-12-05 15:55:52.029624-05
1417	140505	150804	LIMA	HUAURA	CHECRAS	2025-12-05 15:55:52.03068-05	2025-12-05 15:55:52.03068-05
1418	140506	150805	LIMA	HUAURA	HUALMAY	2025-12-05 15:55:52.031248-05	2025-12-05 15:55:52.031248-05
1419	140507	150806	LIMA	HUAURA	HUAURA	2025-12-05 15:55:52.03183-05	2025-12-05 15:55:52.03183-05
1420	140508	150807	LIMA	HUAURA	LEONCIO PRADO	2025-12-05 15:55:52.032451-05	2025-12-05 15:55:52.032451-05
1421	140509	150808	LIMA	HUAURA	PACCHO	2025-12-05 15:55:52.032971-05	2025-12-05 15:55:52.032971-05
1422	140511	150809	LIMA	HUAURA	SANTA LEONOR	2025-12-05 15:55:52.033495-05	2025-12-05 15:55:52.033495-05
1423	140512	150810	LIMA	HUAURA	SANTA MARIA	2025-12-05 15:55:52.033918-05	2025-12-05 15:55:52.033918-05
1424	140513	150811	LIMA	HUAURA	SAYAN	2025-12-05 15:55:52.03446-05	2025-12-05 15:55:52.03446-05
1425	140516	150812	LIMA	HUAURA	VEGUETA	2025-12-05 15:55:52.035119-05	2025-12-05 15:55:52.035119-05
1426	141001	150901	LIMA	OYON	OYON	2025-12-05 15:55:52.035119-05	2025-12-05 15:55:52.035119-05
1427	141004	150902	LIMA	OYON	ANDAJES	2025-12-05 15:55:52.035649-05	2025-12-05 15:55:52.035649-05
1428	141003	150903	LIMA	OYON	CAUJUL	2025-12-05 15:55:52.036185-05	2025-12-05 15:55:52.036185-05
1429	141006	150904	LIMA	OYON	COCHAMARCA	2025-12-05 15:55:52.036185-05	2025-12-05 15:55:52.036185-05
1430	141002	150905	LIMA	OYON	NAVAN	2025-12-05 15:55:52.036745-05	2025-12-05 15:55:52.036745-05
1431	141005	150906	LIMA	OYON	PACHANGARA	2025-12-05 15:55:52.037284-05	2025-12-05 15:55:52.037284-05
1432	140701	151001	LIMA	YAUYOS	YAUYOS	2025-12-05 15:55:52.037284-05	2025-12-05 15:55:52.037284-05
1433	140702	151002	LIMA	YAUYOS	ALIS	2025-12-05 15:55:52.037885-05	2025-12-05 15:55:52.037885-05
1434	140703	151003	LIMA	YAUYOS	AYAUCA	2025-12-05 15:55:52.038448-05	2025-12-05 15:55:52.038448-05
1435	140704	151004	LIMA	YAUYOS	AYAVIRI	2025-12-05 15:55:52.038448-05	2025-12-05 15:55:52.038448-05
1436	140705	151005	LIMA	YAUYOS	AZANGARO	2025-12-05 15:55:52.039329-05	2025-12-05 15:55:52.039329-05
1437	140706	151006	LIMA	YAUYOS	CACRA	2025-12-05 15:55:52.039867-05	2025-12-05 15:55:52.039867-05
1438	140707	151007	LIMA	YAUYOS	CARANIA	2025-12-05 15:55:52.039867-05	2025-12-05 15:55:52.039867-05
1439	140733	151008	LIMA	YAUYOS	CATAHUASI	2025-12-05 15:55:52.040398-05	2025-12-05 15:55:52.040398-05
1440	140710	151009	LIMA	YAUYOS	CHOCOS	2025-12-05 15:55:52.040915-05	2025-12-05 15:55:52.040915-05
1441	140708	151010	LIMA	YAUYOS	COCHAS	2025-12-05 15:55:52.040915-05	2025-12-05 15:55:52.040915-05
1442	140709	151011	LIMA	YAUYOS	COLONIA	2025-12-05 15:55:52.041527-05	2025-12-05 15:55:52.041527-05
1443	140730	151012	LIMA	YAUYOS	HONGOS	2025-12-05 15:55:52.041527-05	2025-12-05 15:55:52.041527-05
1444	140711	151013	LIMA	YAUYOS	HUAMPARA	2025-12-05 15:55:52.042612-05	2025-12-05 15:55:52.042612-05
1445	140712	151014	LIMA	YAUYOS	HUANCAYA	2025-12-05 15:55:52.04315-05	2025-12-05 15:55:52.04315-05
1446	140713	151015	LIMA	YAUYOS	HUANGASCAR	2025-12-05 15:55:52.043686-05	2025-12-05 15:55:52.043686-05
1447	140714	151016	LIMA	YAUYOS	HUANTAN	2025-12-05 15:55:52.044278-05	2025-12-05 15:55:52.044278-05
1448	140715	151017	LIMA	YAUYOS	HUAÑEC	2025-12-05 15:55:52.045005-05	2025-12-05 15:55:52.045005-05
1449	140716	151018	LIMA	YAUYOS	LARAOS	2025-12-05 15:55:52.045688-05	2025-12-05 15:55:52.045688-05
1450	140717	151019	LIMA	YAUYOS	LINCHA	2025-12-05 15:55:52.046315-05	2025-12-05 15:55:52.046315-05
1451	140731	151020	LIMA	YAUYOS	MADEAN	2025-12-05 15:55:52.046848-05	2025-12-05 15:55:52.046848-05
1452	140718	151021	LIMA	YAUYOS	MIRAFLORES	2025-12-05 15:55:52.047377-05	2025-12-05 15:55:52.047377-05
1453	140719	151022	LIMA	YAUYOS	OMAS	2025-12-05 15:55:52.048364-05	2025-12-05 15:55:52.048364-05
1454	140732	151023	LIMA	YAUYOS	PUTINZA	2025-12-05 15:55:52.049047-05	2025-12-05 15:55:52.049047-05
1455	140720	151024	LIMA	YAUYOS	QUINCHES	2025-12-05 15:55:52.049571-05	2025-12-05 15:55:52.049571-05
1456	140721	151025	LIMA	YAUYOS	QUINOCAY	2025-12-05 15:55:52.050144-05	2025-12-05 15:55:52.050144-05
1457	140722	151026	LIMA	YAUYOS	SAN JOAQUIN	2025-12-05 15:55:52.050703-05	2025-12-05 15:55:52.050703-05
1458	140723	151027	LIMA	YAUYOS	SAN PEDRO DE PILAS	2025-12-05 15:55:52.051219-05	2025-12-05 15:55:52.051219-05
1459	140724	151028	LIMA	YAUYOS	TANTA	2025-12-05 15:55:52.051801-05	2025-12-05 15:55:52.051801-05
1460	140725	151029	LIMA	YAUYOS	TAURIPAMPA	2025-12-05 15:55:52.051801-05	2025-12-05 15:55:52.051801-05
1461	140727	151030	LIMA	YAUYOS	TOMAS	2025-12-05 15:55:52.052316-05	2025-12-05 15:55:52.052316-05
1462	140726	151031	LIMA	YAUYOS	TUPE	2025-12-05 15:55:52.052842-05	2025-12-05 15:55:52.052842-05
1463	140728	151032	LIMA	YAUYOS	VIÑAC	2025-12-05 15:55:52.053928-05	2025-12-05 15:55:52.053928-05
1464	140729	151033	LIMA	YAUYOS	VITIS	2025-12-05 15:55:52.053928-05	2025-12-05 15:55:52.053928-05
1465	150101	160101	LORETO	MAYNAS	IQUITOS	2025-12-05 15:55:52.054446-05	2025-12-05 15:55:52.054446-05
1466	150102	160102	LORETO	MAYNAS	ALTO NANAY	2025-12-05 15:55:52.05497-05	2025-12-05 15:55:52.05497-05
1467	150103	160103	LORETO	MAYNAS	FERNANDO LORES	2025-12-05 15:55:52.05497-05	2025-12-05 15:55:52.05497-05
1468	150110	160104	LORETO	MAYNAS	INDIANA	2025-12-05 15:55:52.055621-05	2025-12-05 15:55:52.055621-05
1469	150104	160105	LORETO	MAYNAS	LAS AMAZONAS	2025-12-05 15:55:52.056133-05	2025-12-05 15:55:52.056133-05
1470	150105	160106	LORETO	MAYNAS	MAZAN	2025-12-05 15:55:52.056538-05	2025-12-05 15:55:52.056538-05
1471	150106	160107	LORETO	MAYNAS	NAPO	2025-12-05 15:55:52.057071-05	2025-12-05 15:55:52.057071-05
1472	150111	160108	LORETO	MAYNAS	PUNCHANA	2025-12-05 15:55:52.057071-05	2025-12-05 15:55:52.057071-05
1473	150107	160109	LORETO	MAYNAS	PUTUMAYO	2025-12-05 15:55:52.057615-05	2025-12-05 15:55:52.057615-05
1474	150108	160110	LORETO	MAYNAS	TORRES CAUSANA	2025-12-05 15:55:52.058133-05	2025-12-05 15:55:52.058133-05
1475	150112	160112	LORETO	MAYNAS	BELEN	2025-12-05 15:55:52.058133-05	2025-12-05 15:55:52.058133-05
1476	150113	160113	LORETO	MAYNAS	SAN JUAN BAUTISTA	2025-12-05 15:55:52.05871-05	2025-12-05 15:55:52.05871-05
1477	150114	160114	LORETO	MAYNAS	TENIENTE MANUEL CLAVERO	2025-12-05 15:55:52.059744-05	2025-12-05 15:55:52.059744-05
1478	150201	160201	LORETO	ALTO AMAZONAS	YURIMAGUAS	2025-12-05 15:55:52.060359-05	2025-12-05 15:55:52.060359-05
1479	150202	160202	LORETO	ALTO AMAZONAS	BALSAPUERTO	2025-12-05 15:55:52.06145-05	2025-12-05 15:55:52.06145-05
1480	150205	160205	LORETO	ALTO AMAZONAS	JEBEROS	2025-12-05 15:55:52.06198-05	2025-12-05 15:55:52.06198-05
1481	150206	160206	LORETO	ALTO AMAZONAS	LAGUNAS	2025-12-05 15:55:52.062675-05	2025-12-05 15:55:52.062675-05
1482	150210	160210	LORETO	ALTO AMAZONAS	SANTA CRUZ	2025-12-05 15:55:52.063317-05	2025-12-05 15:55:52.063317-05
1483	150211	160211	LORETO	ALTO AMAZONAS	TENIENTE CESAR LOPEZ ROJAS	2025-12-05 15:55:52.064718-05	2025-12-05 15:55:52.064718-05
1484	150301	160301	LORETO	LORETO	NAUTA	2025-12-05 15:55:52.065379-05	2025-12-05 15:55:52.065379-05
1485	150302	160302	LORETO	LORETO	PARINARI	2025-12-05 15:55:52.065915-05	2025-12-05 15:55:52.065915-05
1486	150303	160303	LORETO	LORETO	TIGRE	2025-12-05 15:55:52.066432-05	2025-12-05 15:55:52.066432-05
1487	150305	160304	LORETO	LORETO	TROMPETEROS	2025-12-05 15:55:52.067466-05	2025-12-05 15:55:52.067466-05
1488	150304	160305	LORETO	LORETO	URARINAS	2025-12-05 15:55:52.067466-05	2025-12-05 15:55:52.067466-05
1489	150601	160401	LORETO	MARISCAL RAMON CASTILLA	RAMON CASTILLA	2025-12-05 15:55:52.067979-05	2025-12-05 15:55:52.067979-05
1490	150602	160402	LORETO	MARISCAL RAMON CASTILLA	PEBAS	2025-12-05 15:55:52.068497-05	2025-12-05 15:55:52.068497-05
1491	150603	160403	LORETO	MARISCAL RAMON CASTILLA	YAVARI	2025-12-05 15:55:52.069028-05	2025-12-05 15:55:52.069028-05
1492	150604	160404	LORETO	MARISCAL RAMON CASTILLA	SAN PABLO	2025-12-05 15:55:52.069028-05	2025-12-05 15:55:52.069028-05
1493	150401	160501	LORETO	REQUENA	REQUENA	2025-12-05 15:55:52.069548-05	2025-12-05 15:55:52.069548-05
1494	150402	160502	LORETO	REQUENA	ALTO TAPICHE	2025-12-05 15:55:52.07009-05	2025-12-05 15:55:52.07009-05
1495	150403	160503	LORETO	REQUENA	CAPELO	2025-12-05 15:55:52.070619-05	2025-12-05 15:55:52.070619-05
1496	150404	160504	LORETO	REQUENA	EMILIO SAN MARTIN	2025-12-05 15:55:52.071136-05	2025-12-05 15:55:52.071136-05
1497	150405	160505	LORETO	REQUENA	MAQUIA	2025-12-05 15:55:52.071136-05	2025-12-05 15:55:52.071136-05
1498	150406	160506	LORETO	REQUENA	PUINAHUA	2025-12-05 15:55:52.071652-05	2025-12-05 15:55:52.071652-05
1499	150407	160507	LORETO	REQUENA	SAQUENA	2025-12-05 15:55:52.072472-05	2025-12-05 15:55:52.072472-05
1500	150408	160508	LORETO	REQUENA	SOPLIN	2025-12-05 15:55:52.072995-05	2025-12-05 15:55:52.072995-05
1501	150409	160509	LORETO	REQUENA	TAPICHE	2025-12-05 15:55:52.072995-05	2025-12-05 15:55:52.072995-05
1502	150410	160510	LORETO	REQUENA	JENARO HERRERA	2025-12-05 15:55:52.073518-05	2025-12-05 15:55:52.073518-05
1503	150411	160511	LORETO	REQUENA	YAQUERANA	2025-12-05 15:55:52.074049-05	2025-12-05 15:55:52.074049-05
1504	150501	160601	LORETO	UCAYALI	CONTAMANA	2025-12-05 15:55:52.074049-05	2025-12-05 15:55:52.074049-05
1505	150506	160602	LORETO	UCAYALI	INAHUAYA	2025-12-05 15:55:52.074566-05	2025-12-05 15:55:52.074566-05
1506	150503	160603	LORETO	UCAYALI	PADRE MARQUEZ	2025-12-05 15:55:52.075125-05	2025-12-05 15:55:52.075125-05
1507	150504	160604	LORETO	UCAYALI	PAMPA HERMOSA	2025-12-05 15:55:52.075676-05	2025-12-05 15:55:52.075676-05
1508	150505	160605	LORETO	UCAYALI	SARAYACU	2025-12-05 15:55:52.076321-05	2025-12-05 15:55:52.076321-05
1509	150502	160606	LORETO	UCAYALI	VARGAS GUERRA	2025-12-05 15:55:52.077517-05	2025-12-05 15:55:52.077517-05
1510	150701	160701	LORETO	DATEM DEL MARAÑON	BARRANCA	2025-12-05 15:55:52.078133-05	2025-12-05 15:55:52.078133-05
1511	150703	160702	LORETO	DATEM DEL MARAÑON	CAHUAPANAS	2025-12-05 15:55:52.078709-05	2025-12-05 15:55:52.078709-05
1512	150704	160703	LORETO	DATEM DEL MARAÑON	MANSERICHE	2025-12-05 15:55:52.079311-05	2025-12-05 15:55:52.079311-05
1513	150705	160704	LORETO	DATEM DEL MARAÑON	MORONA	2025-12-05 15:55:52.079905-05	2025-12-05 15:55:52.079905-05
1514	150706	160705	LORETO	DATEM DEL MARAÑON	PASTAZA	2025-12-05 15:55:52.080901-05	2025-12-05 15:55:52.080901-05
1515	150702	160706	LORETO	DATEM DEL MARAÑON	ANDOAS	2025-12-05 15:55:52.081901-05	2025-12-05 15:55:52.081901-05
1516	150901	160801	LORETO	PUTUMAYO	PUTUMAYO	2025-12-05 15:55:52.082347-05	2025-12-05 15:55:52.082347-05
1517	150902	160802	LORETO	PUTUMAYO	ROSA PANDURO	2025-12-05 15:55:52.082734-05	2025-12-05 15:55:52.082734-05
1518	150903	160803	LORETO	PUTUMAYO	TENIENTE MANUEL CLAVERO	2025-12-05 15:55:52.083265-05	2025-12-05 15:55:52.083265-05
1519	150904	160804	LORETO	PUTUMAYO	YAGUAS	2025-12-05 15:55:52.083782-05	2025-12-05 15:55:52.083782-05
1520	160101	170101	MADRE DE DIOS	TAMBOPATA	TAMBOPATA	2025-12-05 15:55:52.084351-05	2025-12-05 15:55:52.084351-05
1521	160102	170102	MADRE DE DIOS	TAMBOPATA	INAMBARI	2025-12-05 15:55:52.084887-05	2025-12-05 15:55:52.084887-05
1522	160103	170103	MADRE DE DIOS	TAMBOPATA	LAS PIEDRAS	2025-12-05 15:55:52.084887-05	2025-12-05 15:55:52.084887-05
1523	160104	170104	MADRE DE DIOS	TAMBOPATA	LABERINTO	2025-12-05 15:55:52.085488-05	2025-12-05 15:55:52.085488-05
1524	160201	170201	MADRE DE DIOS	MANU	MANU	2025-12-05 15:55:52.085488-05	2025-12-05 15:55:52.085488-05
1525	160202	170202	MADRE DE DIOS	MANU	FITZCARRALD	2025-12-05 15:55:52.086043-05	2025-12-05 15:55:52.086043-05
1526	160203	170203	MADRE DE DIOS	MANU	MADRE DE DIOS	2025-12-05 15:55:52.086585-05	2025-12-05 15:55:52.086585-05
1527	160204	170204	MADRE DE DIOS	MANU	HUEPETUHE	2025-12-05 15:55:52.086585-05	2025-12-05 15:55:52.086585-05
1528	160301	170301	MADRE DE DIOS	TAHUAMANU	IÑAPARI	2025-12-05 15:55:52.087103-05	2025-12-05 15:55:52.087103-05
1529	160302	170302	MADRE DE DIOS	TAHUAMANU	IBERIA	2025-12-05 15:55:52.087647-05	2025-12-05 15:55:52.087647-05
1530	160303	170303	MADRE DE DIOS	TAHUAMANU	TAHUAMANU	2025-12-05 15:55:52.088249-05	2025-12-05 15:55:52.088249-05
1531	170101	180101	MOQUEGUA	MARISCAL NIETO	MOQUEGUA	2025-12-05 15:55:52.088928-05	2025-12-05 15:55:52.088928-05
1532	170102	180102	MOQUEGUA	MARISCAL NIETO	CARUMAS	2025-12-05 15:55:52.089442-05	2025-12-05 15:55:52.089442-05
1533	170103	180103	MOQUEGUA	MARISCAL NIETO	CUCHUMBAYA	2025-12-05 15:55:52.089442-05	2025-12-05 15:55:52.089442-05
1534	170106	180104	MOQUEGUA	MARISCAL NIETO	SAMEGUA	2025-12-05 15:55:52.089974-05	2025-12-05 15:55:52.089974-05
1535	170104	180105	MOQUEGUA	MARISCAL NIETO	SAN CRISTOBAL	2025-12-05 15:55:52.090518-05	2025-12-05 15:55:52.090518-05
1536	170105	180106	MOQUEGUA	MARISCAL NIETO	TORATA	2025-12-05 15:55:52.090518-05	2025-12-05 15:55:52.090518-05
1537	170201	180201	MOQUEGUA	GENERAL SANCHEZ CERRO	OMATE	2025-12-05 15:55:52.091039-05	2025-12-05 15:55:52.091039-05
1538	170203	180202	MOQUEGUA	GENERAL SANCHEZ CERRO	CHOJATA	2025-12-05 15:55:52.09156-05	2025-12-05 15:55:52.09156-05
1539	170202	180203	MOQUEGUA	GENERAL SANCHEZ CERRO	COALAQUE	2025-12-05 15:55:52.09156-05	2025-12-05 15:55:52.09156-05
1540	170204	180204	MOQUEGUA	GENERAL SANCHEZ CERRO	ICHUÑA	2025-12-05 15:55:52.092083-05	2025-12-05 15:55:52.092083-05
1541	170205	180205	MOQUEGUA	GENERAL SANCHEZ CERRO	LA CAPILLA	2025-12-05 15:55:52.092656-05	2025-12-05 15:55:52.092656-05
1542	170206	180206	MOQUEGUA	GENERAL SANCHEZ CERRO	LLOQUE	2025-12-05 15:55:52.093711-05	2025-12-05 15:55:52.093711-05
1543	170207	180207	MOQUEGUA	GENERAL SANCHEZ CERRO	MATALAQUE	2025-12-05 15:55:52.094237-05	2025-12-05 15:55:52.094237-05
1544	170208	180208	MOQUEGUA	GENERAL SANCHEZ CERRO	PUQUINA	2025-12-05 15:55:52.094977-05	2025-12-05 15:55:52.094977-05
1545	170209	180209	MOQUEGUA	GENERAL SANCHEZ CERRO	QUINISTAQUILLAS	2025-12-05 15:55:52.095639-05	2025-12-05 15:55:52.095639-05
1546	170210	180210	MOQUEGUA	GENERAL SANCHEZ CERRO	UBINAS	2025-12-05 15:55:52.096298-05	2025-12-05 15:55:52.096298-05
1547	170211	180211	MOQUEGUA	GENERAL SANCHEZ CERRO	YUNGA	2025-12-05 15:55:52.096812-05	2025-12-05 15:55:52.096812-05
1548	170301	180301	MOQUEGUA	ILO	ILO	2025-12-05 15:55:52.097336-05	2025-12-05 15:55:52.097336-05
1549	170302	180302	MOQUEGUA	ILO	EL ALGARROBAL	2025-12-05 15:55:52.098697-05	2025-12-05 15:55:52.098697-05
1550	170303	180303	MOQUEGUA	ILO	PACOCHA	2025-12-05 15:55:52.09931-05	2025-12-05 15:55:52.09931-05
1551	180101	190101	PASCO	PASCO	CHAUPIMARCA	2025-12-05 15:55:52.09985-05	2025-12-05 15:55:52.09985-05
1552	180103	190102	PASCO	PASCO	HUACHON	2025-12-05 15:55:52.100382-05	2025-12-05 15:55:52.100382-05
1553	180104	190103	PASCO	PASCO	HUARIACA	2025-12-05 15:55:52.100899-05	2025-12-05 15:55:52.100899-05
1554	180105	190104	PASCO	PASCO	HUAYLLAY	2025-12-05 15:55:52.100899-05	2025-12-05 15:55:52.100899-05
1555	180106	190105	PASCO	PASCO	NINACACA	2025-12-05 15:55:52.101422-05	2025-12-05 15:55:52.101422-05
1556	180107	190106	PASCO	PASCO	PALLANCHACRA	2025-12-05 15:55:52.10194-05	2025-12-05 15:55:52.10194-05
1557	180108	190107	PASCO	PASCO	PAUCARTAMBO	2025-12-05 15:55:52.10194-05	2025-12-05 15:55:52.10194-05
1558	180109	190108	PASCO	PASCO	SAN FRANCISCO DE ASIS DE YARUS	2025-12-05 15:55:52.102463-05	2025-12-05 15:55:52.102463-05
1559	180110	190109	PASCO	PASCO	SIMON BOLIVAR	2025-12-05 15:55:52.103071-05	2025-12-05 15:55:52.103071-05
1560	180111	190110	PASCO	PASCO	TICLACAYAN	2025-12-05 15:55:52.103597-05	2025-12-05 15:55:52.103597-05
1561	180112	190111	PASCO	PASCO	TINYAHUARCO	2025-12-05 15:55:52.103597-05	2025-12-05 15:55:52.103597-05
1562	180113	190112	PASCO	PASCO	VICCO	2025-12-05 15:55:52.104545-05	2025-12-05 15:55:52.104545-05
1563	180114	190113	PASCO	PASCO	YANACANCHA	2025-12-05 15:55:52.105159-05	2025-12-05 15:55:52.105159-05
1564	180201	190201	PASCO	DANIEL ALCIDES CARRION	YANAHUANCA	2025-12-05 15:55:52.105159-05	2025-12-05 15:55:52.105159-05
1565	180202	190202	PASCO	DANIEL ALCIDES CARRION	CHACAYAN	2025-12-05 15:55:52.105693-05	2025-12-05 15:55:52.105693-05
1566	180203	190203	PASCO	DANIEL ALCIDES CARRION	GOYLLARISQUIZGA	2025-12-05 15:55:52.10621-05	2025-12-05 15:55:52.10621-05
1567	180204	190204	PASCO	DANIEL ALCIDES CARRION	PAUCAR	2025-12-05 15:55:52.10621-05	2025-12-05 15:55:52.10621-05
1568	180205	190205	PASCO	DANIEL ALCIDES CARRION	SAN PEDRO DE PILLAO	2025-12-05 15:55:52.106738-05	2025-12-05 15:55:52.106738-05
1569	180206	190206	PASCO	DANIEL ALCIDES CARRION	SANTA ANA DE TUSI	2025-12-05 15:55:52.107289-05	2025-12-05 15:55:52.107289-05
1570	180207	190207	PASCO	DANIEL ALCIDES CARRION	TAPUC	2025-12-05 15:55:52.107289-05	2025-12-05 15:55:52.107289-05
1571	180208	190208	PASCO	DANIEL ALCIDES CARRION	VILCABAMBA	2025-12-05 15:55:52.107847-05	2025-12-05 15:55:52.107847-05
1572	180301	190301	PASCO	OXAPAMPA	OXAPAMPA	2025-12-05 15:55:52.108395-05	2025-12-05 15:55:52.108395-05
1573	180302	190302	PASCO	OXAPAMPA	CHONTABAMBA	2025-12-05 15:55:52.108395-05	2025-12-05 15:55:52.108395-05
1574	180303	190303	PASCO	OXAPAMPA	HUANCABAMBA	2025-12-05 15:55:52.109527-05	2025-12-05 15:55:52.109527-05
1575	180307	190304	PASCO	OXAPAMPA	PALCAZU	2025-12-05 15:55:52.110057-05	2025-12-05 15:55:52.110057-05
1576	180306	190305	PASCO	OXAPAMPA	POZUZO	2025-12-05 15:55:52.111274-05	2025-12-05 15:55:52.111274-05
1577	180304	190306	PASCO	OXAPAMPA	PUERTO BERMUDEZ	2025-12-05 15:55:52.111817-05	2025-12-05 15:55:52.111817-05
1578	180305	190307	PASCO	OXAPAMPA	VILLA RICA	2025-12-05 15:55:52.112343-05	2025-12-05 15:55:52.112343-05
1579	180308	190308	PASCO	OXAPAMPA	CONSTITUCION	2025-12-05 15:55:52.112343-05	2025-12-05 15:55:52.112343-05
1580	190101	200101	PIURA	PIURA	PIURA	2025-12-05 15:55:52.112866-05	2025-12-05 15:55:52.112866-05
1581	190103	200104	PIURA	PIURA	CASTILLA	2025-12-05 15:55:52.113453-05	2025-12-05 15:55:52.113453-05
1582	190104	200105	PIURA	PIURA	CATACAOS	2025-12-05 15:55:52.114298-05	2025-12-05 15:55:52.114298-05
1583	190113	200107	PIURA	PIURA	CURA MORI	2025-12-05 15:55:52.114829-05	2025-12-05 15:55:52.114829-05
1584	190114	200108	PIURA	PIURA	EL TALLAN	2025-12-05 15:55:52.115888-05	2025-12-05 15:55:52.115888-05
1585	190105	200109	PIURA	PIURA	LA ARENA	2025-12-05 15:55:52.116503-05	2025-12-05 15:55:52.116503-05
1586	190106	200110	PIURA	PIURA	LA UNION	2025-12-05 15:55:52.11705-05	2025-12-05 15:55:52.11705-05
1587	190107	200111	PIURA	PIURA	LAS LOMAS	2025-12-05 15:55:52.117576-05	2025-12-05 15:55:52.117576-05
1588	190109	200114	PIURA	PIURA	TAMBO GRANDE	2025-12-05 15:55:52.117576-05	2025-12-05 15:55:52.117576-05
1589	190115	200115	PIURA	PIURA	VEINTISEIS DE OCTUBRE	2025-12-05 15:55:52.118099-05	2025-12-05 15:55:52.118099-05
1590	190201	200201	PIURA	AYABACA	AYABACA	2025-12-05 15:55:52.118628-05	2025-12-05 15:55:52.118628-05
1591	190202	200202	PIURA	AYABACA	FRIAS	2025-12-05 15:55:52.119283-05	2025-12-05 15:55:52.119283-05
1592	190209	200203	PIURA	AYABACA	JILILI	2025-12-05 15:55:52.119283-05	2025-12-05 15:55:52.119283-05
1593	190203	200204	PIURA	AYABACA	LAGUNAS	2025-12-05 15:55:52.119878-05	2025-12-05 15:55:52.119878-05
1594	190204	200205	PIURA	AYABACA	MONTERO	2025-12-05 15:55:52.120404-05	2025-12-05 15:55:52.120404-05
1595	190205	200206	PIURA	AYABACA	PACAIPAMPA	2025-12-05 15:55:52.120404-05	2025-12-05 15:55:52.120404-05
1596	190210	200207	PIURA	AYABACA	PAIMAS	2025-12-05 15:55:52.121004-05	2025-12-05 15:55:52.121004-05
1597	190206	200208	PIURA	AYABACA	SAPILLICA	2025-12-05 15:55:52.121534-05	2025-12-05 15:55:52.121534-05
1598	190207	200209	PIURA	AYABACA	SICCHEZ	2025-12-05 15:55:52.122284-05	2025-12-05 15:55:52.122284-05
1599	190208	200210	PIURA	AYABACA	SUYO	2025-12-05 15:55:52.122284-05	2025-12-05 15:55:52.122284-05
1600	190301	200301	PIURA	HUANCABAMBA	HUANCABAMBA	2025-12-05 15:55:52.122798-05	2025-12-05 15:55:52.122798-05
1601	190302	200302	PIURA	HUANCABAMBA	CANCHAQUE	2025-12-05 15:55:52.123335-05	2025-12-05 15:55:52.123335-05
1602	190306	200303	PIURA	HUANCABAMBA	EL CARMEN DE LA FRONTERA	2025-12-05 15:55:52.123335-05	2025-12-05 15:55:52.123335-05
1603	190303	200304	PIURA	HUANCABAMBA	HUARMACA	2025-12-05 15:55:52.124407-05	2025-12-05 15:55:52.124407-05
1604	190308	200305	PIURA	HUANCABAMBA	LALAQUIZ	2025-12-05 15:55:52.124925-05	2025-12-05 15:55:52.124925-05
1605	190307	200306	PIURA	HUANCABAMBA	SAN MIGUEL DE EL FAIQUE	2025-12-05 15:55:52.124925-05	2025-12-05 15:55:52.124925-05
1606	190304	200307	PIURA	HUANCABAMBA	SONDOR	2025-12-05 15:55:52.125445-05	2025-12-05 15:55:52.125445-05
1607	190305	200308	PIURA	HUANCABAMBA	SONDORILLO	2025-12-05 15:55:52.126492-05	2025-12-05 15:55:52.126492-05
1608	190401	200401	PIURA	MORROPON	CHULUCANAS	2025-12-05 15:55:52.127012-05	2025-12-05 15:55:52.127012-05
1609	190402	200402	PIURA	MORROPON	BUENOS AIRES	2025-12-05 15:55:52.127543-05	2025-12-05 15:55:52.127543-05
1610	190403	200403	PIURA	MORROPON	CHALACO	2025-12-05 15:55:52.128069-05	2025-12-05 15:55:52.128069-05
1611	190408	200404	PIURA	MORROPON	LA MATANZA	2025-12-05 15:55:52.12859-05	2025-12-05 15:55:52.12859-05
1612	190404	200405	PIURA	MORROPON	MORROPON	2025-12-05 15:55:52.129635-05	2025-12-05 15:55:52.129635-05
1613	190405	200406	PIURA	MORROPON	SALITRAL	2025-12-05 15:55:52.129635-05	2025-12-05 15:55:52.129635-05
1614	190410	200407	PIURA	MORROPON	SAN JUAN DE BIGOTE	2025-12-05 15:55:52.130154-05	2025-12-05 15:55:52.130154-05
1615	190406	200408	PIURA	MORROPON	SANTA CATALINA DE MOSSA	2025-12-05 15:55:52.13072-05	2025-12-05 15:55:52.13072-05
1616	190407	200409	PIURA	MORROPON	SANTO DOMINGO	2025-12-05 15:55:52.132371-05	2025-12-05 15:55:52.132371-05
1617	190409	200410	PIURA	MORROPON	YAMANGO	2025-12-05 15:55:52.132888-05	2025-12-05 15:55:52.132888-05
1618	190501	200501	PIURA	PAITA	PAITA	2025-12-05 15:55:52.133412-05	2025-12-05 15:55:52.133412-05
1619	190502	200502	PIURA	PAITA	AMOTAPE	2025-12-05 15:55:52.134025-05	2025-12-05 15:55:52.134025-05
1620	190503	200503	PIURA	PAITA	ARENAL	2025-12-05 15:55:52.134025-05	2025-12-05 15:55:52.134025-05
1621	190505	200504	PIURA	PAITA	COLAN	2025-12-05 15:55:52.134549-05	2025-12-05 15:55:52.134549-05
1622	190504	200505	PIURA	PAITA	LA HUACA	2025-12-05 15:55:52.135117-05	2025-12-05 15:55:52.135117-05
1623	190506	200506	PIURA	PAITA	TAMARINDO	2025-12-05 15:55:52.135117-05	2025-12-05 15:55:52.135117-05
1624	190507	200507	PIURA	PAITA	VICHAYAL	2025-12-05 15:55:52.135681-05	2025-12-05 15:55:52.135681-05
1625	190601	200601	PIURA	SULLANA	SULLANA	2025-12-05 15:55:52.13624-05	2025-12-05 15:55:52.13624-05
1626	190602	200602	PIURA	SULLANA	BELLAVISTA	2025-12-05 15:55:52.13624-05	2025-12-05 15:55:52.13624-05
1627	190608	200603	PIURA	SULLANA	IGNACIO ESCUDERO	2025-12-05 15:55:52.136763-05	2025-12-05 15:55:52.136763-05
1628	190603	200604	PIURA	SULLANA	LANCONES	2025-12-05 15:55:52.137341-05	2025-12-05 15:55:52.137341-05
1629	190604	200605	PIURA	SULLANA	MARCAVELICA	2025-12-05 15:55:52.137978-05	2025-12-05 15:55:52.137978-05
1630	190605	200606	PIURA	SULLANA	MIGUEL CHECA	2025-12-05 15:55:52.137978-05	2025-12-05 15:55:52.137978-05
1631	190606	200607	PIURA	SULLANA	QUERECOTILLO	2025-12-05 15:55:52.138528-05	2025-12-05 15:55:52.138528-05
1632	190607	200608	PIURA	SULLANA	SALITRAL	2025-12-05 15:55:52.139039-05	2025-12-05 15:55:52.139039-05
1633	190701	200701	PIURA	TALARA	PARIÑAS	2025-12-05 15:55:52.139039-05	2025-12-05 15:55:52.139039-05
1634	190702	200702	PIURA	TALARA	EL ALTO	2025-12-05 15:55:52.139578-05	2025-12-05 15:55:52.139578-05
1635	190703	200703	PIURA	TALARA	LA BREA	2025-12-05 15:55:52.140091-05	2025-12-05 15:55:52.140091-05
1636	190704	200704	PIURA	TALARA	LOBITOS	2025-12-05 15:55:52.140091-05	2025-12-05 15:55:52.140091-05
1637	190706	200705	PIURA	TALARA	LOS ORGANOS	2025-12-05 15:55:52.14061-05	2025-12-05 15:55:52.14061-05
1638	190705	200706	PIURA	TALARA	MANCORA	2025-12-05 15:55:52.14121-05	2025-12-05 15:55:52.14121-05
1639	190801	200801	PIURA	SECHURA	SECHURA	2025-12-05 15:55:52.14121-05	2025-12-05 15:55:52.14121-05
1640	190804	200802	PIURA	SECHURA	BELLAVISTA DE LA UNION	2025-12-05 15:55:52.141741-05	2025-12-05 15:55:52.141741-05
1641	190803	200803	PIURA	SECHURA	BERNAL	2025-12-05 15:55:52.142401-05	2025-12-05 15:55:52.142401-05
1642	190805	200804	PIURA	SECHURA	CRISTO NOS VALGA	2025-12-05 15:55:52.143-05	2025-12-05 15:55:52.143-05
1643	190802	200805	PIURA	SECHURA	VICE	2025-12-05 15:55:52.143755-05	2025-12-05 15:55:52.143755-05
1644	190806	200806	PIURA	SECHURA	RINCONADA LLICUAR	2025-12-05 15:55:52.144275-05	2025-12-05 15:55:52.144275-05
1645	200101	210101	PUNO	PUNO	PUNO	2025-12-05 15:55:52.145335-05	2025-12-05 15:55:52.145335-05
1646	200102	210102	PUNO	PUNO	ACORA	2025-12-05 15:55:52.145951-05	2025-12-05 15:55:52.145951-05
1647	200115	210103	PUNO	PUNO	AMANTANI	2025-12-05 15:55:52.146477-05	2025-12-05 15:55:52.146477-05
1648	200103	210104	PUNO	PUNO	ATUNCOLLA	2025-12-05 15:55:52.146477-05	2025-12-05 15:55:52.146477-05
1649	200104	210105	PUNO	PUNO	CAPACHICA	2025-12-05 15:55:52.14761-05	2025-12-05 15:55:52.14761-05
1650	200106	210106	PUNO	PUNO	CHUCUITO	2025-12-05 15:55:52.148316-05	2025-12-05 15:55:52.148316-05
1651	200105	210107	PUNO	PUNO	COATA	2025-12-05 15:55:52.148833-05	2025-12-05 15:55:52.148833-05
1652	200107	210108	PUNO	PUNO	HUATA	2025-12-05 15:55:52.1494-05	2025-12-05 15:55:52.1494-05
1653	200108	210109	PUNO	PUNO	MAÑAZO	2025-12-05 15:55:52.149987-05	2025-12-05 15:55:52.149987-05
1654	200109	210110	PUNO	PUNO	PAUCARCOLLA	2025-12-05 15:55:52.150588-05	2025-12-05 15:55:52.150588-05
1655	200110	210111	PUNO	PUNO	PICHACANI	2025-12-05 15:55:52.150588-05	2025-12-05 15:55:52.150588-05
1656	200114	210112	PUNO	PUNO	PLATERIA	2025-12-05 15:55:52.151112-05	2025-12-05 15:55:52.151112-05
1657	200111	210113	PUNO	PUNO	SAN ANTONIO	2025-12-05 15:55:52.15163-05	2025-12-05 15:55:52.15163-05
1658	200112	210114	PUNO	PUNO	TIQUILLACA	2025-12-05 15:55:52.152448-05	2025-12-05 15:55:52.152448-05
1659	200113	210115	PUNO	PUNO	VILQUE	2025-12-05 15:55:52.152448-05	2025-12-05 15:55:52.152448-05
1660	200201	210201	PUNO	AZANGARO	AZANGARO	2025-12-05 15:55:52.153031-05	2025-12-05 15:55:52.153031-05
1661	200202	210202	PUNO	AZANGARO	ACHAYA	2025-12-05 15:55:52.153586-05	2025-12-05 15:55:52.153586-05
1662	200203	210203	PUNO	AZANGARO	ARAPA	2025-12-05 15:55:52.153586-05	2025-12-05 15:55:52.153586-05
1663	200204	210204	PUNO	AZANGARO	ASILLO	2025-12-05 15:55:52.154124-05	2025-12-05 15:55:52.154124-05
1664	200205	210205	PUNO	AZANGARO	CAMINACA	2025-12-05 15:55:52.154692-05	2025-12-05 15:55:52.154692-05
1665	200206	210206	PUNO	AZANGARO	CHUPA	2025-12-05 15:55:52.154692-05	2025-12-05 15:55:52.154692-05
1666	200207	210207	PUNO	AZANGARO	JOSE DOMINGO CHOQUEHUANCA	2025-12-05 15:55:52.155212-05	2025-12-05 15:55:52.155212-05
1667	200208	210208	PUNO	AZANGARO	MUÑANI	2025-12-05 15:55:52.155727-05	2025-12-05 15:55:52.155727-05
1668	200210	210209	PUNO	AZANGARO	POTONI	2025-12-05 15:55:52.155727-05	2025-12-05 15:55:52.155727-05
1669	200212	210210	PUNO	AZANGARO	SAMAN	2025-12-05 15:55:52.156244-05	2025-12-05 15:55:52.156244-05
1670	200213	210211	PUNO	AZANGARO	SAN ANTON	2025-12-05 15:55:52.156244-05	2025-12-05 15:55:52.156244-05
1671	200214	210212	PUNO	AZANGARO	SAN JOSE	2025-12-05 15:55:52.157259-05	2025-12-05 15:55:52.157259-05
1672	200215	210213	PUNO	AZANGARO	SAN JUAN DE SALINAS	2025-12-05 15:55:52.157259-05	2025-12-05 15:55:52.157259-05
1673	200216	210214	PUNO	AZANGARO	SANTIAGO DE PUPUJA	2025-12-05 15:55:52.15779-05	2025-12-05 15:55:52.15779-05
1674	200217	210215	PUNO	AZANGARO	TIRAPATA	2025-12-05 15:55:52.158312-05	2025-12-05 15:55:52.158312-05
1675	200301	210301	PUNO	CARABAYA	MACUSANI	2025-12-05 15:55:52.158312-05	2025-12-05 15:55:52.158312-05
1676	200302	210302	PUNO	CARABAYA	AJOYANI	2025-12-05 15:55:52.159435-05	2025-12-05 15:55:52.159435-05
1677	200303	210303	PUNO	CARABAYA	AYAPATA	2025-12-05 15:55:52.159953-05	2025-12-05 15:55:52.159953-05
1678	200304	210304	PUNO	CARABAYA	COASA	2025-12-05 15:55:52.160479-05	2025-12-05 15:55:52.160479-05
1679	200305	210305	PUNO	CARABAYA	CORANI	2025-12-05 15:55:52.161152-05	2025-12-05 15:55:52.161152-05
1680	200306	210306	PUNO	CARABAYA	CRUCERO	2025-12-05 15:55:52.161867-05	2025-12-05 15:55:52.161867-05
1681	200307	210307	PUNO	CARABAYA	ITUATA	2025-12-05 15:55:52.162462-05	2025-12-05 15:55:52.162462-05
1682	200308	210308	PUNO	CARABAYA	OLLACHEA	2025-12-05 15:55:52.163132-05	2025-12-05 15:55:52.163132-05
1683	200309	210309	PUNO	CARABAYA	SAN GABAN	2025-12-05 15:55:52.163758-05	2025-12-05 15:55:52.163758-05
1684	200310	210310	PUNO	CARABAYA	USICAYOS	2025-12-05 15:55:52.164498-05	2025-12-05 15:55:52.164498-05
1685	200401	210401	PUNO	CHUCUITO	JULI	2025-12-05 15:55:52.16526-05	2025-12-05 15:55:52.16526-05
1686	200402	210402	PUNO	CHUCUITO	DESAGUADERO	2025-12-05 15:55:52.166009-05	2025-12-05 15:55:52.166009-05
1687	200403	210403	PUNO	CHUCUITO	HUACULLANI	2025-12-05 15:55:52.166009-05	2025-12-05 15:55:52.166009-05
1688	200412	210404	PUNO	CHUCUITO	KELLUYO	2025-12-05 15:55:52.166534-05	2025-12-05 15:55:52.166534-05
1689	200406	210405	PUNO	CHUCUITO	PISACOMA	2025-12-05 15:55:52.167081-05	2025-12-05 15:55:52.167081-05
1690	200407	210406	PUNO	CHUCUITO	POMATA	2025-12-05 15:55:52.167633-05	2025-12-05 15:55:52.167633-05
1691	200410	210407	PUNO	CHUCUITO	ZEPITA	2025-12-05 15:55:52.168021-05	2025-12-05 15:55:52.168021-05
1692	201201	210501	PUNO	EL COLLAO	ILAVE	2025-12-05 15:55:52.168449-05	2025-12-05 15:55:52.168449-05
1693	201204	210502	PUNO	EL COLLAO	CAPAZO	2025-12-05 15:55:52.168449-05	2025-12-05 15:55:52.168449-05
1694	201202	210503	PUNO	EL COLLAO	PILCUYO	2025-12-05 15:55:52.169001-05	2025-12-05 15:55:52.169001-05
1695	201203	210504	PUNO	EL COLLAO	SANTA ROSA	2025-12-05 15:55:52.169534-05	2025-12-05 15:55:52.169534-05
1696	201205	210505	PUNO	EL COLLAO	CONDURIRI	2025-12-05 15:55:52.170156-05	2025-12-05 15:55:52.170156-05
1697	200501	210601	PUNO	HUANCANE	HUANCANE	2025-12-05 15:55:52.170156-05	2025-12-05 15:55:52.170156-05
1698	200502	210602	PUNO	HUANCANE	COJATA	2025-12-05 15:55:52.170683-05	2025-12-05 15:55:52.170683-05
1699	200511	210603	PUNO	HUANCANE	HUATASANI	2025-12-05 15:55:52.171215-05	2025-12-05 15:55:52.171215-05
1700	200504	210604	PUNO	HUANCANE	INCHUPALLA	2025-12-05 15:55:52.171215-05	2025-12-05 15:55:52.171215-05
1701	200506	210605	PUNO	HUANCANE	PUSI	2025-12-05 15:55:52.171798-05	2025-12-05 15:55:52.171798-05
1702	200507	210606	PUNO	HUANCANE	ROSASPATA	2025-12-05 15:55:52.172352-05	2025-12-05 15:55:52.172352-05
1703	200508	210607	PUNO	HUANCANE	TARACO	2025-12-05 15:55:52.172352-05	2025-12-05 15:55:52.172352-05
1704	200509	210608	PUNO	HUANCANE	VILQUE CHICO	2025-12-05 15:55:52.172897-05	2025-12-05 15:55:52.172897-05
1705	200601	210701	PUNO	LAMPA	LAMPA	2025-12-05 15:55:52.172897-05	2025-12-05 15:55:52.172897-05
1706	200602	210702	PUNO	LAMPA	CABANILLA	2025-12-05 15:55:52.17359-05	2025-12-05 15:55:52.17359-05
1707	200603	210703	PUNO	LAMPA	CALAPUJA	2025-12-05 15:55:52.174113-05	2025-12-05 15:55:52.174113-05
1708	200604	210704	PUNO	LAMPA	NICASIO	2025-12-05 15:55:52.174113-05	2025-12-05 15:55:52.174113-05
1709	200605	210705	PUNO	LAMPA	OCUVIRI	2025-12-05 15:55:52.174722-05	2025-12-05 15:55:52.174722-05
1710	200606	210706	PUNO	LAMPA	PALCA	2025-12-05 15:55:52.174722-05	2025-12-05 15:55:52.174722-05
1711	200607	210707	PUNO	LAMPA	PARATIA	2025-12-05 15:55:52.175275-05	2025-12-05 15:55:52.175275-05
1712	200608	210708	PUNO	LAMPA	PUCARA	2025-12-05 15:55:52.176373-05	2025-12-05 15:55:52.176373-05
1713	200609	210709	PUNO	LAMPA	SANTA LUCIA	2025-12-05 15:55:52.176897-05	2025-12-05 15:55:52.176897-05
1714	200610	210710	PUNO	LAMPA	VILAVILA	2025-12-05 15:55:52.177428-05	2025-12-05 15:55:52.177428-05
1715	200701	210801	PUNO	MELGAR	AYAVIRI	2025-12-05 15:55:52.177977-05	2025-12-05 15:55:52.177977-05
1716	200702	210802	PUNO	MELGAR	ANTAUTA	2025-12-05 15:55:52.1786-05	2025-12-05 15:55:52.1786-05
1717	200703	210803	PUNO	MELGAR	CUPI	2025-12-05 15:55:52.1793-05	2025-12-05 15:55:52.1793-05
1718	200704	210804	PUNO	MELGAR	LLALLI	2025-12-05 15:55:52.179923-05	2025-12-05 15:55:52.179923-05
1719	200705	210805	PUNO	MELGAR	MACARI	2025-12-05 15:55:52.180861-05	2025-12-05 15:55:52.180861-05
1720	200706	210806	PUNO	MELGAR	NUÑOA	2025-12-05 15:55:52.181215-05	2025-12-05 15:55:52.181215-05
1721	200707	210807	PUNO	MELGAR	ORURILLO	2025-12-05 15:55:52.181943-05	2025-12-05 15:55:52.181943-05
1722	200708	210808	PUNO	MELGAR	SANTA ROSA	2025-12-05 15:55:52.182471-05	2025-12-05 15:55:52.182471-05
1723	200709	210809	PUNO	MELGAR	UMACHIRI	2025-12-05 15:55:52.183015-05	2025-12-05 15:55:52.183015-05
1724	201301	210901	PUNO	MOHO	MOHO	2025-12-05 15:55:52.183745-05	2025-12-05 15:55:52.183745-05
1725	201302	210902	PUNO	MOHO	CONIMA	2025-12-05 15:55:52.184172-05	2025-12-05 15:55:52.184172-05
1726	201304	210903	PUNO	MOHO	HUAYRAPATA	2025-12-05 15:55:52.184749-05	2025-12-05 15:55:52.184749-05
1727	201303	210904	PUNO	MOHO	TILALI	2025-12-05 15:55:52.185329-05	2025-12-05 15:55:52.185329-05
1728	201101	211001	PUNO	SAN ANTONIO DE PUTINA	PUTINA	2025-12-05 15:55:52.185329-05	2025-12-05 15:55:52.185329-05
1729	201104	211002	PUNO	SAN ANTONIO DE PUTINA	ANANEA	2025-12-05 15:55:52.185851-05	2025-12-05 15:55:52.185851-05
1730	201102	211003	PUNO	SAN ANTONIO DE PUTINA	PEDRO VILCA APAZA	2025-12-05 15:55:52.185851-05	2025-12-05 15:55:52.185851-05
1731	201103	211004	PUNO	SAN ANTONIO DE PUTINA	QUILCAPUNCU	2025-12-05 15:55:52.186454-05	2025-12-05 15:55:52.186454-05
1732	201105	211005	PUNO	SAN ANTONIO DE PUTINA	SINA	2025-12-05 15:55:52.18699-05	2025-12-05 15:55:52.18699-05
1733	200901	211101	PUNO	SAN ROMAN	JULIACA	2025-12-05 15:55:52.187504-05	2025-12-05 15:55:52.187504-05
1734	200902	211102	PUNO	SAN ROMAN	CABANA	2025-12-05 15:55:52.188079-05	2025-12-05 15:55:52.188079-05
1735	200903	211103	PUNO	SAN ROMAN	CABANILLAS	2025-12-05 15:55:52.188079-05	2025-12-05 15:55:52.188079-05
1736	200904	211104	PUNO	SAN ROMAN	CARACOTO	2025-12-05 15:55:52.188844-05	2025-12-05 15:55:52.188844-05
1737	200905	211105	PUNO	SAN ROMAN	SAN MIGUEL	2025-12-05 15:55:52.189358-05	2025-12-05 15:55:52.189358-05
1738	200801	211201	PUNO	SANDIA	SANDIA	2025-12-05 15:55:52.189358-05	2025-12-05 15:55:52.189358-05
1739	200803	211202	PUNO	SANDIA	CUYOCUYO	2025-12-05 15:55:52.189965-05	2025-12-05 15:55:52.189965-05
1740	200804	211203	PUNO	SANDIA	LIMBANI	2025-12-05 15:55:52.190487-05	2025-12-05 15:55:52.190487-05
1741	200806	211204	PUNO	SANDIA	PATAMBUCO	2025-12-05 15:55:52.190487-05	2025-12-05 15:55:52.190487-05
1742	200805	211205	PUNO	SANDIA	PHARA	2025-12-05 15:55:52.191011-05	2025-12-05 15:55:52.191011-05
1743	200807	211206	PUNO	SANDIA	QUIACA	2025-12-05 15:55:52.19157-05	2025-12-05 15:55:52.19157-05
1744	200808	211207	PUNO	SANDIA	SAN JUAN DEL ORO	2025-12-05 15:55:52.19157-05	2025-12-05 15:55:52.19157-05
1745	200810	211208	PUNO	SANDIA	YANAHUAYA	2025-12-05 15:55:52.192092-05	2025-12-05 15:55:52.192092-05
1746	200811	211209	PUNO	SANDIA	ALTO INAMBARI	2025-12-05 15:55:52.192728-05	2025-12-05 15:55:52.192728-05
1747	200812	211210	PUNO	SANDIA	SAN PEDRO DE PUTINA PUNCO	2025-12-05 15:55:52.193927-05	2025-12-05 15:55:52.193927-05
1748	201001	211301	PUNO	YUNGUYO	YUNGUYO	2025-12-05 15:55:52.194502-05	2025-12-05 15:55:52.194502-05
1749	201003	211302	PUNO	YUNGUYO	ANAPIA	2025-12-05 15:55:52.195024-05	2025-12-05 15:55:52.195024-05
1750	201004	211303	PUNO	YUNGUYO	COPANI	2025-12-05 15:55:52.195645-05	2025-12-05 15:55:52.195645-05
1751	201005	211304	PUNO	YUNGUYO	CUTURAPI	2025-12-05 15:55:52.19623-05	2025-12-05 15:55:52.19623-05
1752	201006	211305	PUNO	YUNGUYO	OLLARAYA	2025-12-05 15:55:52.196753-05	2025-12-05 15:55:52.196753-05
1753	201007	211306	PUNO	YUNGUYO	TINICACHI	2025-12-05 15:55:52.197562-05	2025-12-05 15:55:52.197562-05
1754	201002	211307	PUNO	YUNGUYO	UNICACHI	2025-12-05 15:55:52.198314-05	2025-12-05 15:55:52.198314-05
1755	210101	220101	SAN MARTIN	MOYOBAMBA	MOYOBAMBA	2025-12-05 15:55:52.198899-05	2025-12-05 15:55:52.198899-05
1756	210102	220102	SAN MARTIN	MOYOBAMBA	CALZADA	2025-12-05 15:55:52.199361-05	2025-12-05 15:55:52.199361-05
1757	210103	220103	SAN MARTIN	MOYOBAMBA	HABANA	2025-12-05 15:55:52.200026-05	2025-12-05 15:55:52.200026-05
1758	210104	220104	SAN MARTIN	MOYOBAMBA	JEPELACIO	2025-12-05 15:55:52.200602-05	2025-12-05 15:55:52.200602-05
1759	210105	220105	SAN MARTIN	MOYOBAMBA	SORITOR	2025-12-05 15:55:52.200602-05	2025-12-05 15:55:52.200602-05
1760	210106	220106	SAN MARTIN	MOYOBAMBA	YANTALO	2025-12-05 15:55:52.201679-05	2025-12-05 15:55:52.201679-05
1761	210701	220201	SAN MARTIN	BELLAVISTA	BELLAVISTA	2025-12-05 15:55:52.202201-05	2025-12-05 15:55:52.202201-05
1762	210704	220202	SAN MARTIN	BELLAVISTA	ALTO BIAVO	2025-12-05 15:55:52.202201-05	2025-12-05 15:55:52.202201-05
1763	210706	220203	SAN MARTIN	BELLAVISTA	BAJO BIAVO	2025-12-05 15:55:52.20274-05	2025-12-05 15:55:52.20274-05
1764	210705	220204	SAN MARTIN	BELLAVISTA	HUALLAGA	2025-12-05 15:55:52.203375-05	2025-12-05 15:55:52.203375-05
1765	210703	220205	SAN MARTIN	BELLAVISTA	SAN PABLO	2025-12-05 15:55:52.203892-05	2025-12-05 15:55:52.203892-05
1766	210702	220206	SAN MARTIN	BELLAVISTA	SAN RAFAEL	2025-12-05 15:55:52.20441-05	2025-12-05 15:55:52.20441-05
1767	211001	220301	SAN MARTIN	EL DORADO	SAN JOSE DE SISA	2025-12-05 15:55:52.204821-05	2025-12-05 15:55:52.204821-05
1768	211002	220302	SAN MARTIN	EL DORADO	AGUA BLANCA	2025-12-05 15:55:52.205479-05	2025-12-05 15:55:52.205479-05
1769	211004	220303	SAN MARTIN	EL DORADO	SAN MARTIN	2025-12-05 15:55:52.205479-05	2025-12-05 15:55:52.205479-05
1770	211005	220304	SAN MARTIN	EL DORADO	SANTA ROSA	2025-12-05 15:55:52.206033-05	2025-12-05 15:55:52.206033-05
1771	211003	220305	SAN MARTIN	EL DORADO	SHATOJA	2025-12-05 15:55:52.206555-05	2025-12-05 15:55:52.206555-05
1772	210201	220401	SAN MARTIN	HUALLAGA	SAPOSOA	2025-12-05 15:55:52.207076-05	2025-12-05 15:55:52.207076-05
1773	210205	220402	SAN MARTIN	HUALLAGA	ALTO SAPOSOA	2025-12-05 15:55:52.207076-05	2025-12-05 15:55:52.207076-05
1774	210206	220403	SAN MARTIN	HUALLAGA	EL ESLABON	2025-12-05 15:55:52.207593-05	2025-12-05 15:55:52.207593-05
1775	210202	220404	SAN MARTIN	HUALLAGA	PISCOYACU	2025-12-05 15:55:52.208135-05	2025-12-05 15:55:52.208135-05
1776	210203	220405	SAN MARTIN	HUALLAGA	SACANCHE	2025-12-05 15:55:52.208135-05	2025-12-05 15:55:52.208135-05
1777	210204	220406	SAN MARTIN	HUALLAGA	TINGO DE SAPOSOA	2025-12-05 15:55:52.208733-05	2025-12-05 15:55:52.208733-05
1778	210301	220501	SAN MARTIN	LAMAS	LAMAS	2025-12-05 15:55:52.209299-05	2025-12-05 15:55:52.209299-05
1779	210315	220502	SAN MARTIN	LAMAS	ALONSO DE ALVARADO	2025-12-05 15:55:52.209828-05	2025-12-05 15:55:52.209828-05
1780	210303	220503	SAN MARTIN	LAMAS	BARRANQUITA	2025-12-05 15:55:52.210544-05	2025-12-05 15:55:52.210544-05
1781	210304	220504	SAN MARTIN	LAMAS	CAYNARACHI	2025-12-05 15:55:52.211067-05	2025-12-05 15:55:52.211067-05
1782	210305	220505	SAN MARTIN	LAMAS	CUÑUMBUQUI	2025-12-05 15:55:52.212122-05	2025-12-05 15:55:52.212122-05
1783	210306	220506	SAN MARTIN	LAMAS	PINTO RECODO	2025-12-05 15:55:52.212677-05	2025-12-05 15:55:52.212677-05
1784	210307	220507	SAN MARTIN	LAMAS	RUMISAPA	2025-12-05 15:55:52.212677-05	2025-12-05 15:55:52.212677-05
1785	210316	220508	SAN MARTIN	LAMAS	SAN ROQUE DE CUMBAZA	2025-12-05 15:55:52.213792-05	2025-12-05 15:55:52.213792-05
1786	210311	220509	SAN MARTIN	LAMAS	SHANAO	2025-12-05 15:55:52.214649-05	2025-12-05 15:55:52.214649-05
1787	210313	220510	SAN MARTIN	LAMAS	TABALOSOS	2025-12-05 15:55:52.215425-05	2025-12-05 15:55:52.215425-05
1788	210314	220511	SAN MARTIN	LAMAS	ZAPATERO	2025-12-05 15:55:52.215947-05	2025-12-05 15:55:52.215947-05
1789	210401	220601	SAN MARTIN	MARISCAL CACERES	JUANJUI	2025-12-05 15:55:52.216469-05	2025-12-05 15:55:52.216469-05
1790	210402	220602	SAN MARTIN	MARISCAL CACERES	CAMPANILLA	2025-12-05 15:55:52.217022-05	2025-12-05 15:55:52.217022-05
1791	210403	220603	SAN MARTIN	MARISCAL CACERES	HUICUNGO	2025-12-05 15:55:52.217622-05	2025-12-05 15:55:52.217622-05
1792	210404	220604	SAN MARTIN	MARISCAL CACERES	PACHIZA	2025-12-05 15:55:52.217622-05	2025-12-05 15:55:52.217622-05
1793	210405	220605	SAN MARTIN	MARISCAL CACERES	PAJARILLO	2025-12-05 15:55:52.218152-05	2025-12-05 15:55:52.218152-05
1794	210901	220701	SAN MARTIN	PICOTA	PICOTA	2025-12-05 15:55:52.218708-05	2025-12-05 15:55:52.218708-05
1795	210902	220702	SAN MARTIN	PICOTA	BUENOS AIRES	2025-12-05 15:55:52.218708-05	2025-12-05 15:55:52.218708-05
1796	210903	220703	SAN MARTIN	PICOTA	CASPISAPA	2025-12-05 15:55:52.219258-05	2025-12-05 15:55:52.219258-05
1797	210904	220704	SAN MARTIN	PICOTA	PILLUANA	2025-12-05 15:55:52.219843-05	2025-12-05 15:55:52.219843-05
1798	210905	220705	SAN MARTIN	PICOTA	PUCACACA	2025-12-05 15:55:52.220384-05	2025-12-05 15:55:52.220384-05
1799	210906	220706	SAN MARTIN	PICOTA	SAN CRISTOBAL	2025-12-05 15:55:52.220697-05	2025-12-05 15:55:52.220697-05
1800	210907	220707	SAN MARTIN	PICOTA	SAN HILARION	2025-12-05 15:55:52.221221-05	2025-12-05 15:55:52.221221-05
1801	210910	220708	SAN MARTIN	PICOTA	SHAMBOYACU	2025-12-05 15:55:52.221741-05	2025-12-05 15:55:52.221741-05
1802	210908	220709	SAN MARTIN	PICOTA	TINGO DE PONASA	2025-12-05 15:55:52.222264-05	2025-12-05 15:55:52.222264-05
1803	210909	220710	SAN MARTIN	PICOTA	TRES UNIDOS	2025-12-05 15:55:52.222264-05	2025-12-05 15:55:52.222264-05
1804	210501	220801	SAN MARTIN	RIOJA	RIOJA	2025-12-05 15:55:52.222782-05	2025-12-05 15:55:52.222782-05
1805	210509	220802	SAN MARTIN	RIOJA	AWAJUN	2025-12-05 15:55:52.223322-05	2025-12-05 15:55:52.223322-05
1806	210506	220803	SAN MARTIN	RIOJA	ELIAS SOPLIN VARGAS	2025-12-05 15:55:52.223841-05	2025-12-05 15:55:52.223841-05
1807	210505	220804	SAN MARTIN	RIOJA	NUEVA CAJAMARCA	2025-12-05 15:55:52.224363-05	2025-12-05 15:55:52.224363-05
1808	210508	220805	SAN MARTIN	RIOJA	PARDO MIGUEL	2025-12-05 15:55:52.225491-05	2025-12-05 15:55:52.225491-05
1809	210502	220806	SAN MARTIN	RIOJA	POSIC	2025-12-05 15:55:52.226431-05	2025-12-05 15:55:52.226431-05
1810	210507	220807	SAN MARTIN	RIOJA	SAN FERNANDO	2025-12-05 15:55:52.226963-05	2025-12-05 15:55:52.226963-05
1811	210503	220808	SAN MARTIN	RIOJA	YORONGOS	2025-12-05 15:55:52.227497-05	2025-12-05 15:55:52.227497-05
1812	210504	220809	SAN MARTIN	RIOJA	YURACYACU	2025-12-05 15:55:52.228092-05	2025-12-05 15:55:52.228092-05
1813	210601	220901	SAN MARTIN	SAN MARTIN	TARAPOTO	2025-12-05 15:55:52.228779-05	2025-12-05 15:55:52.228779-05
1814	210602	220902	SAN MARTIN	SAN MARTIN	ALBERTO LEVEAU	2025-12-05 15:55:52.229383-05	2025-12-05 15:55:52.229383-05
1815	210604	220903	SAN MARTIN	SAN MARTIN	CACATACHI	2025-12-05 15:55:52.229932-05	2025-12-05 15:55:52.229932-05
1816	210606	220904	SAN MARTIN	SAN MARTIN	CHAZUTA	2025-12-05 15:55:52.230985-05	2025-12-05 15:55:52.230985-05
1817	210607	220905	SAN MARTIN	SAN MARTIN	CHIPURANA	2025-12-05 15:55:52.231656-05	2025-12-05 15:55:52.231656-05
1818	210608	220906	SAN MARTIN	SAN MARTIN	EL PORVENIR	2025-12-05 15:55:52.232284-05	2025-12-05 15:55:52.232284-05
1819	210609	220907	SAN MARTIN	SAN MARTIN	HUIMBAYOC	2025-12-05 15:55:52.232875-05	2025-12-05 15:55:52.232875-05
1820	210610	220908	SAN MARTIN	SAN MARTIN	JUAN GUERRA	2025-12-05 15:55:52.233395-05	2025-12-05 15:55:52.233395-05
1821	210621	220909	SAN MARTIN	SAN MARTIN	LA BANDA DE SHILCAYO	2025-12-05 15:55:52.233395-05	2025-12-05 15:55:52.233395-05
1822	210611	220910	SAN MARTIN	SAN MARTIN	MORALES	2025-12-05 15:55:52.23398-05	2025-12-05 15:55:52.23398-05
1823	210612	220911	SAN MARTIN	SAN MARTIN	PAPAPLAYA	2025-12-05 15:55:52.234505-05	2025-12-05 15:55:52.234505-05
1824	210616	220912	SAN MARTIN	SAN MARTIN	SAN ANTONIO	2025-12-05 15:55:52.234909-05	2025-12-05 15:55:52.234909-05
1825	210619	220913	SAN MARTIN	SAN MARTIN	SAUCE	2025-12-05 15:55:52.235685-05	2025-12-05 15:55:52.235685-05
1826	210620	220914	SAN MARTIN	SAN MARTIN	SHAPAJA	2025-12-05 15:55:52.236197-05	2025-12-05 15:55:52.236197-05
1827	210801	221001	SAN MARTIN	TOCACHE	TOCACHE	2025-12-05 15:55:52.236893-05	2025-12-05 15:55:52.236893-05
1828	210802	221002	SAN MARTIN	TOCACHE	NUEVO PROGRESO	2025-12-05 15:55:52.23743-05	2025-12-05 15:55:52.23743-05
1829	210803	221003	SAN MARTIN	TOCACHE	POLVORA	2025-12-05 15:55:52.23743-05	2025-12-05 15:55:52.23743-05
1830	210804	221004	SAN MARTIN	TOCACHE	SHUNTE	2025-12-05 15:55:52.237951-05	2025-12-05 15:55:52.237951-05
1831	210805	221005	SAN MARTIN	TOCACHE	UCHIZA	2025-12-05 15:55:52.238479-05	2025-12-05 15:55:52.238479-05
1832	210806	221006	SAN MARTIN	TOCACHE	SANTA LUCIA	2025-12-05 15:55:52.238479-05	2025-12-05 15:55:52.238479-05
1833	220101	230101	TACNA	TACNA	TACNA	2025-12-05 15:55:52.239108-05	2025-12-05 15:55:52.239108-05
1834	220111	230102	TACNA	TACNA	ALTO DE LA ALIANZA	2025-12-05 15:55:52.239682-05	2025-12-05 15:55:52.239682-05
1835	220102	230103	TACNA	TACNA	CALANA	2025-12-05 15:55:52.240205-05	2025-12-05 15:55:52.240205-05
1836	220112	230104	TACNA	TACNA	CIUDAD NUEVA	2025-12-05 15:55:52.240205-05	2025-12-05 15:55:52.240205-05
1837	220104	230105	TACNA	TACNA	INCLAN	2025-12-05 15:55:52.240807-05	2025-12-05 15:55:52.240807-05
1838	220107	230106	TACNA	TACNA	PACHIA	2025-12-05 15:55:52.240807-05	2025-12-05 15:55:52.240807-05
1839	220108	230107	TACNA	TACNA	PALCA	2025-12-05 15:55:52.241356-05	2025-12-05 15:55:52.241356-05
1840	220109	230108	TACNA	TACNA	POCOLLAY	2025-12-05 15:55:52.241902-05	2025-12-05 15:55:52.241902-05
1841	220110	230109	TACNA	TACNA	SAMA	2025-12-05 15:55:52.242425-05	2025-12-05 15:55:52.242425-05
1842	220113	230110	TACNA	TACNA	CORONEL GREGORIO ALBARRACIN LA	2025-12-05 15:55:52.242985-05	2025-12-05 15:55:52.242985-05
1843	220114	230111	TACNA	TACNA	LA YARADA LOS PALOS	2025-12-05 15:55:52.243518-05	2025-12-05 15:55:52.243518-05
1844	220401	230201	TACNA	CANDARAVE	CANDARAVE	2025-12-05 15:55:52.24407-05	2025-12-05 15:55:52.24407-05
1845	220402	230202	TACNA	CANDARAVE	CAIRANI	2025-12-05 15:55:52.244738-05	2025-12-05 15:55:52.244738-05
1846	220406	230203	TACNA	CANDARAVE	CAMILACA	2025-12-05 15:55:52.245837-05	2025-12-05 15:55:52.245837-05
1847	220403	230204	TACNA	CANDARAVE	CURIBAYA	2025-12-05 15:55:52.246427-05	2025-12-05 15:55:52.246427-05
1848	220404	230205	TACNA	CANDARAVE	HUANUARA	2025-12-05 15:55:52.246958-05	2025-12-05 15:55:52.246958-05
1849	220405	230206	TACNA	CANDARAVE	QUILAHUANI	2025-12-05 15:55:52.247814-05	2025-12-05 15:55:52.247814-05
1850	220301	230301	TACNA	JORGE BASADRE	LOCUMBA	2025-12-05 15:55:52.248747-05	2025-12-05 15:55:52.248747-05
1851	220303	230302	TACNA	JORGE BASADRE	ILABAYA	2025-12-05 15:55:52.249279-05	2025-12-05 15:55:52.249279-05
1852	220302	230303	TACNA	JORGE BASADRE	ITE	2025-12-05 15:55:52.249859-05	2025-12-05 15:55:52.249859-05
1853	220201	230401	TACNA	TARATA	TARATA	2025-12-05 15:55:52.249859-05	2025-12-05 15:55:52.249859-05
1854	220205	230402	TACNA	TARATA	HEROES ALBARRACIN CHUCATAMANI	2025-12-05 15:55:52.250384-05	2025-12-05 15:55:52.250384-05
1855	220206	230403	TACNA	TARATA	ESTIQUE	2025-12-05 15:55:52.250995-05	2025-12-05 15:55:52.250995-05
1856	220207	230404	TACNA	TARATA	ESTIQUE-PAMPA	2025-12-05 15:55:52.251432-05	2025-12-05 15:55:52.251432-05
1857	220210	230405	TACNA	TARATA	SITAJARA	2025-12-05 15:55:52.251946-05	2025-12-05 15:55:52.251946-05
1858	220211	230406	TACNA	TARATA	SUSAPAYA	2025-12-05 15:55:52.252488-05	2025-12-05 15:55:52.252488-05
1859	220212	230407	TACNA	TARATA	TARUCACHI	2025-12-05 15:55:52.252488-05	2025-12-05 15:55:52.252488-05
1860	220213	230408	TACNA	TARATA	TICACO	2025-12-05 15:55:52.253021-05	2025-12-05 15:55:52.253021-05
1861	230101	240101	TUMBES	TUMBES	TUMBES	2025-12-05 15:55:52.25359-05	2025-12-05 15:55:52.25359-05
1862	230102	240102	TUMBES	TUMBES	CORRALES	2025-12-05 15:55:52.25359-05	2025-12-05 15:55:52.25359-05
1863	230103	240103	TUMBES	TUMBES	LA CRUZ	2025-12-05 15:55:52.254149-05	2025-12-05 15:55:52.254149-05
1864	230104	240104	TUMBES	TUMBES	PAMPAS DE HOSPITAL	2025-12-05 15:55:52.25468-05	2025-12-05 15:55:52.25468-05
1865	230105	240105	TUMBES	TUMBES	SAN JACINTO	2025-12-05 15:55:52.25468-05	2025-12-05 15:55:52.25468-05
1866	230106	240106	TUMBES	TUMBES	SAN JUAN DE LA VIRGEN	2025-12-05 15:55:52.255276-05	2025-12-05 15:55:52.255276-05
1867	230201	240201	TUMBES	CONTRALMIRANTE VILLAR	ZORRITOS	2025-12-05 15:55:52.255812-05	2025-12-05 15:55:52.255812-05
1868	230202	240202	TUMBES	CONTRALMIRANTE VILLAR	CASITAS	2025-12-05 15:55:52.256325-05	2025-12-05 15:55:52.256325-05
1869	230203	240203	TUMBES	CONTRALMIRANTE VILLAR	CANOAS DE PUNTA SAL	2025-12-05 15:55:52.256325-05	2025-12-05 15:55:52.256325-05
1870	230301	240301	TUMBES	ZARUMILLA	ZARUMILLA	2025-12-05 15:55:52.25698-05	2025-12-05 15:55:52.25698-05
1871	230304	240302	TUMBES	ZARUMILLA	AGUAS VERDES	2025-12-05 15:55:52.257525-05	2025-12-05 15:55:52.257525-05
1872	230302	240303	TUMBES	ZARUMILLA	MATAPALO	2025-12-05 15:55:52.257525-05	2025-12-05 15:55:52.257525-05
1873	230303	240304	TUMBES	ZARUMILLA	PAPAYAL	2025-12-05 15:55:52.258595-05	2025-12-05 15:55:52.258595-05
1874	250101	250101	UCAYALI	CORONEL PORTILLO	CALLERIA	2025-12-05 15:55:52.259208-05	2025-12-05 15:55:52.259208-05
1875	250104	250102	UCAYALI	CORONEL PORTILLO	CAMPOVERDE	2025-12-05 15:55:52.259725-05	2025-12-05 15:55:52.259725-05
1876	250105	250103	UCAYALI	CORONEL PORTILLO	IPARIA	2025-12-05 15:55:52.260443-05	2025-12-05 15:55:52.260443-05
1877	250103	250104	UCAYALI	CORONEL PORTILLO	MASISEA	2025-12-05 15:55:52.261512-05	2025-12-05 15:55:52.261512-05
1878	250102	250105	UCAYALI	CORONEL PORTILLO	YARINACOCHA	2025-12-05 15:55:52.262564-05	2025-12-05 15:55:52.262564-05
1879	250106	250106	UCAYALI	CORONEL PORTILLO	NUEVA REQUENA	2025-12-05 15:55:52.262564-05	2025-12-05 15:55:52.262564-05
1880	250107	250107	UCAYALI	CORONEL PORTILLO	MANANTAY	2025-12-05 15:55:52.263681-05	2025-12-05 15:55:52.263681-05
1881	250301	250201	UCAYALI	ATALAYA	RAYMONDI	2025-12-05 15:55:52.264485-05	2025-12-05 15:55:52.264485-05
1882	250304	250202	UCAYALI	ATALAYA	SEPAHUA	2025-12-05 15:55:52.265271-05	2025-12-05 15:55:52.265271-05
1883	250302	250203	UCAYALI	ATALAYA	TAHUANIA	2025-12-05 15:55:52.265938-05	2025-12-05 15:55:52.265938-05
1884	250303	250204	UCAYALI	ATALAYA	YURUA	2025-12-05 15:55:52.266531-05	2025-12-05 15:55:52.266531-05
1885	250201	250301	UCAYALI	PADRE ABAD	PADRE ABAD	2025-12-05 15:55:52.266973-05	2025-12-05 15:55:52.266973-05
1886	250202	250302	UCAYALI	PADRE ABAD	IRAZOLA	2025-12-05 15:55:52.267489-05	2025-12-05 15:55:52.267489-05
1887	250203	250303	UCAYALI	PADRE ABAD	CURIMANA	2025-12-05 15:55:52.267855-05	2025-12-05 15:55:52.267855-05
1888	250204	250304	UCAYALI	PADRE ABAD	NESHUYA	2025-12-05 15:55:52.268377-05	2025-12-05 15:55:52.268377-05
1889	250205	250305	UCAYALI	PADRE ABAD	ALEXANDER VON HUMBOLDT	2025-12-05 15:55:52.268939-05	2025-12-05 15:55:52.268939-05
1890	250207	250306	UCAYALI	PADRE ABAD	HUIPOCA	2025-12-05 15:55:52.268939-05	2025-12-05 15:55:52.268939-05
1891	250206	250307	UCAYALI	PADRE ABAD	BOQUERON	2025-12-05 15:55:52.269458-05	2025-12-05 15:55:52.269458-05
1892	250401	250401	UCAYALI	PURUS	PURUS	2025-12-05 15:55:52.269982-05	2025-12-05 15:55:52.269982-05
1893	170107	250402	MOQUEGUA	MARISCAL NIETO	SAN ANTONIO	2025-12-05 15:55:52.270544-05	2025-12-05 15:55:52.270544-05
\.


--
-- TOC entry 5036 (class 0 OID 21993)
-- Dependencies: 229
-- Data for Name: user_application_roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_application_roles (id, user_id, application_id, application_role_id, granted_at, granted_by, revoked_at, revoked_by, is_deleted, deleted_at, deleted_by, created_at, updated_at) FROM stdin;
9bface44-a84b-410b-aaa7-ede411d8d177	a54e954d-0c61-4861-ab94-d49eaf516672	d5a6d085-424d-45c8-8cd4-48f38981af85	434bed03-d869-4a7b-a9af-404d3c394dd7	2025-12-05 15:55:52.346821-05	a54e954d-0c61-4861-ab94-d49eaf516672	\N	\N	f	\N	\N	2025-12-05 15:55:52.346821-05	2025-12-05 15:55:52.346821-05
c844f1de-77aa-4ba8-9472-f342f603967a	a54e954d-0c61-4861-ab94-d49eaf516672	003ce6c1-b1fe-481d-854e-076fbc0882da	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	2025-12-05 15:55:52.350605-05	a54e954d-0c61-4861-ab94-d49eaf516672	\N	\N	f	\N	\N	2025-12-05 15:55:52.350605-05	2025-12-05 15:55:52.350605-05
21938445-87ff-4663-a112-f3a9111f8f0f	a54e954d-0c61-4861-ab94-d49eaf516672	799de4e5-f480-4f20-a9ba-d3e8174e4193	7dd3c008-6b79-4d5a-982f-2af0bbbe292d	2025-12-05 15:55:52.351839-05	a54e954d-0c61-4861-ab94-d49eaf516672	\N	\N	f	\N	\N	2025-12-05 15:55:52.351839-05	2025-12-05 15:55:52.351839-05
89e5f053-5645-4c60-9745-22804c1fce7f	a54e954d-0c61-4861-ab94-d49eaf516672	e0af8703-b3b0-4103-a0f6-64336111dc82	0c0e9a41-8e06-4798-8f83-a2e45d262807	2025-12-05 15:55:52.353615-05	a54e954d-0c61-4861-ab94-d49eaf516672	\N	\N	f	\N	\N	2025-12-05 15:55:52.353615-05	2025-12-05 15:55:52.353615-05
a3962a85-3474-4c95-b26d-182819f292ab	a54e954d-0c61-4861-ab94-d49eaf516672	5e18521f-8cf7-4570-b10e-26a432219cfa	416cac7b-8d7e-4e8a-bc65-bdeb0cb9eff1	2025-12-05 15:55:52.355192-05	a54e954d-0c61-4861-ab94-d49eaf516672	\N	\N	f	\N	\N	2025-12-05 15:55:52.355192-05	2025-12-05 15:55:52.355192-05
287e17ef-4d0c-44ea-84b6-6aeb4bdbd2fb	9a322e1b-c428-4ee8-9ac4-5b69168dcded	d5a6d085-424d-45c8-8cd4-48f38981af85	434bed03-d869-4a7b-a9af-404d3c394dd7	2025-12-05 15:55:52.35735-05	9a322e1b-c428-4ee8-9ac4-5b69168dcded	\N	\N	f	\N	\N	2025-12-05 15:55:52.35735-05	2025-12-05 15:55:52.35735-05
4e4537db-f3a8-4f70-855b-42c867883e2f	9a322e1b-c428-4ee8-9ac4-5b69168dcded	003ce6c1-b1fe-481d-854e-076fbc0882da	78908267-5474-4d26-9233-4ac1a847ff06	2025-12-05 15:55:52.359249-05	9a322e1b-c428-4ee8-9ac4-5b69168dcded	\N	\N	f	\N	\N	2025-12-05 15:55:52.359249-05	2025-12-05 15:55:52.359249-05
c985e1f2-2331-45c8-b588-63358c20f4fe	9a322e1b-c428-4ee8-9ac4-5b69168dcded	799de4e5-f480-4f20-a9ba-d3e8174e4193	d61a0a42-408c-4d0d-85ed-4e1c7003a735	2025-12-05 15:55:52.361565-05	9a322e1b-c428-4ee8-9ac4-5b69168dcded	\N	\N	f	\N	\N	2025-12-05 15:55:52.361565-05	2025-12-05 15:55:52.361565-05
18aa1f06-22b8-43b8-bb1c-4c61bb57e185	9a322e1b-c428-4ee8-9ac4-5b69168dcded	e0af8703-b3b0-4103-a0f6-64336111dc82	80f0567b-ed2a-4795-9b4b-55069a518964	2025-12-05 15:55:52.363173-05	9a322e1b-c428-4ee8-9ac4-5b69168dcded	\N	\N	f	\N	\N	2025-12-05 15:55:52.363173-05	2025-12-05 15:55:52.363173-05
5a67dea1-9ea9-4b4b-9ea7-803a69ba5221	9a322e1b-c428-4ee8-9ac4-5b69168dcded	5e18521f-8cf7-4570-b10e-26a432219cfa	416cac7b-8d7e-4e8a-bc65-bdeb0cb9eff1	2025-12-05 15:55:52.365224-05	9a322e1b-c428-4ee8-9ac4-5b69168dcded	\N	\N	f	\N	\N	2025-12-05 15:55:52.365224-05	2025-12-05 15:55:52.365224-05
fca62f94-16bb-4506-967c-f7d14a3e3691	980014a4-7574-4b2d-b061-0660868b120a	d5a6d085-424d-45c8-8cd4-48f38981af85	434bed03-d869-4a7b-a9af-404d3c394dd7	2025-12-05 15:55:52.368012-05	980014a4-7574-4b2d-b061-0660868b120a	\N	\N	f	\N	\N	2025-12-05 15:55:52.368012-05	2025-12-05 15:55:52.368012-05
11e90f62-cb64-4e76-b0f1-6a99e068a80e	980014a4-7574-4b2d-b061-0660868b120a	003ce6c1-b1fe-481d-854e-076fbc0882da	78908267-5474-4d26-9233-4ac1a847ff06	2025-12-05 15:55:52.36917-05	980014a4-7574-4b2d-b061-0660868b120a	\N	\N	f	\N	\N	2025-12-05 15:55:52.36917-05	2025-12-05 15:55:52.36917-05
9a9125de-36fb-44f7-8e2b-8488ea41370e	980014a4-7574-4b2d-b061-0660868b120a	799de4e5-f480-4f20-a9ba-d3e8174e4193	d73fa769-7ab7-43e6-9541-fbd80c0e5644	2025-12-05 15:55:52.370538-05	980014a4-7574-4b2d-b061-0660868b120a	\N	\N	f	\N	\N	2025-12-05 15:55:52.370538-05	2025-12-05 15:55:52.370538-05
73073d4c-85a6-47ba-867f-4c3f4e729be4	980014a4-7574-4b2d-b061-0660868b120a	e0af8703-b3b0-4103-a0f6-64336111dc82	80f0567b-ed2a-4795-9b4b-55069a518964	2025-12-05 15:55:52.37212-05	980014a4-7574-4b2d-b061-0660868b120a	\N	\N	f	\N	\N	2025-12-05 15:55:52.37212-05	2025-12-05 15:55:52.37212-05
a2d2e831-8668-427e-bb93-ea2f2f3618eb	980014a4-7574-4b2d-b061-0660868b120a	5e18521f-8cf7-4570-b10e-26a432219cfa	416cac7b-8d7e-4e8a-bc65-bdeb0cb9eff1	2025-12-05 15:55:52.37264-05	980014a4-7574-4b2d-b061-0660868b120a	\N	\N	f	\N	\N	2025-12-05 15:55:52.37264-05	2025-12-05 15:55:52.37264-05
76e92b8b-c025-45f0-8d27-28843d22f457	1d7899ba-7655-4a6a-ab8b-4e361a60115d	d5a6d085-424d-45c8-8cd4-48f38981af85	434bed03-d869-4a7b-a9af-404d3c394dd7	2025-12-05 15:55:52.37431-05	1d7899ba-7655-4a6a-ab8b-4e361a60115d	\N	\N	f	\N	\N	2025-12-05 15:55:52.37431-05	2025-12-05 15:55:52.37431-05
58fa250d-4aac-4278-ab93-4efc55077125	1d7899ba-7655-4a6a-ab8b-4e361a60115d	003ce6c1-b1fe-481d-854e-076fbc0882da	78908267-5474-4d26-9233-4ac1a847ff06	2025-12-05 15:55:52.375359-05	1d7899ba-7655-4a6a-ab8b-4e361a60115d	\N	\N	f	\N	\N	2025-12-05 15:55:52.375359-05	2025-12-05 15:55:52.375359-05
7d1eb3cc-3498-4dc1-b424-db89a0accdef	1d7899ba-7655-4a6a-ab8b-4e361a60115d	799de4e5-f480-4f20-a9ba-d3e8174e4193	b0b22ec5-73ff-4551-950c-4b5b8780abc2	2025-12-05 15:55:52.376678-05	1d7899ba-7655-4a6a-ab8b-4e361a60115d	\N	\N	f	\N	\N	2025-12-05 15:55:52.376678-05	2025-12-05 15:55:52.376678-05
64f37e02-878d-4d65-899a-343c2054592a	1d7899ba-7655-4a6a-ab8b-4e361a60115d	e0af8703-b3b0-4103-a0f6-64336111dc82	80f0567b-ed2a-4795-9b4b-55069a518964	2025-12-05 15:55:52.37779-05	1d7899ba-7655-4a6a-ab8b-4e361a60115d	\N	\N	f	\N	\N	2025-12-05 15:55:52.37779-05	2025-12-05 15:55:52.37779-05
9b342673-c155-4ea6-9a54-2609a71deba3	1d7899ba-7655-4a6a-ab8b-4e361a60115d	5e18521f-8cf7-4570-b10e-26a432219cfa	416cac7b-8d7e-4e8a-bc65-bdeb0cb9eff1	2025-12-05 15:55:52.378949-05	1d7899ba-7655-4a6a-ab8b-4e361a60115d	\N	\N	f	\N	\N	2025-12-05 15:55:52.378949-05	2025-12-05 15:55:52.378949-05
13d400c8-832a-4387-93ff-dca82eb7f806	f3629131-4d9a-40e7-bd2d-7e0d029e3c28	d5a6d085-424d-45c8-8cd4-48f38981af85	434bed03-d869-4a7b-a9af-404d3c394dd7	2025-12-05 15:55:52.381176-05	f3629131-4d9a-40e7-bd2d-7e0d029e3c28	\N	\N	f	\N	\N	2025-12-05 15:55:52.381176-05	2025-12-05 15:55:52.381176-05
60bb3b5a-aa21-4545-840a-9a477f8e1f5e	f3629131-4d9a-40e7-bd2d-7e0d029e3c28	003ce6c1-b1fe-481d-854e-076fbc0882da	78908267-5474-4d26-9233-4ac1a847ff06	2025-12-05 15:55:52.381734-05	f3629131-4d9a-40e7-bd2d-7e0d029e3c28	\N	\N	f	\N	\N	2025-12-05 15:55:52.381734-05	2025-12-05 15:55:52.381734-05
f8ded3c4-7187-43b2-8bc5-e04f9f89fc02	f3629131-4d9a-40e7-bd2d-7e0d029e3c28	799de4e5-f480-4f20-a9ba-d3e8174e4193	0771975d-4776-4603-a409-692070094283	2025-12-05 15:55:52.382779-05	f3629131-4d9a-40e7-bd2d-7e0d029e3c28	\N	\N	f	\N	\N	2025-12-05 15:55:52.382779-05	2025-12-05 15:55:52.382779-05
a7ca9d3f-fc96-4871-9807-28943cdee0dc	f3629131-4d9a-40e7-bd2d-7e0d029e3c28	e0af8703-b3b0-4103-a0f6-64336111dc82	80f0567b-ed2a-4795-9b4b-55069a518964	2025-12-05 15:55:52.383374-05	f3629131-4d9a-40e7-bd2d-7e0d029e3c28	\N	\N	f	\N	\N	2025-12-05 15:55:52.383374-05	2025-12-05 15:55:52.383374-05
0f68c53c-b58b-41bd-a773-e33a9e8ddfe8	f3629131-4d9a-40e7-bd2d-7e0d029e3c28	5e18521f-8cf7-4570-b10e-26a432219cfa	416cac7b-8d7e-4e8a-bc65-bdeb0cb9eff1	2025-12-05 15:55:52.38444-05	f3629131-4d9a-40e7-bd2d-7e0d029e3c28	\N	\N	f	\N	\N	2025-12-05 15:55:52.38444-05	2025-12-05 15:55:52.38444-05
adc0cd50-c924-4124-8f2f-5a70cad87e26	040d4285-a1e3-4ae2-af4f-9696efd7fb7e	d5a6d085-424d-45c8-8cd4-48f38981af85	434bed03-d869-4a7b-a9af-404d3c394dd7	2025-12-05 15:55:52.385878-05	040d4285-a1e3-4ae2-af4f-9696efd7fb7e	\N	\N	f	\N	\N	2025-12-05 15:55:52.385878-05	2025-12-05 15:55:52.385878-05
6dcdd836-4d3d-42f3-a39d-242f012e580f	040d4285-a1e3-4ae2-af4f-9696efd7fb7e	003ce6c1-b1fe-481d-854e-076fbc0882da	78908267-5474-4d26-9233-4ac1a847ff06	2025-12-05 15:55:52.386994-05	040d4285-a1e3-4ae2-af4f-9696efd7fb7e	\N	\N	f	\N	\N	2025-12-05 15:55:52.386994-05	2025-12-05 15:55:52.386994-05
218f439e-821c-449c-9531-78ba9a326a82	040d4285-a1e3-4ae2-af4f-9696efd7fb7e	799de4e5-f480-4f20-a9ba-d3e8174e4193	276cdc35-8604-415a-8eca-a22eb99c4074	2025-12-05 15:55:52.387516-05	040d4285-a1e3-4ae2-af4f-9696efd7fb7e	\N	\N	f	\N	\N	2025-12-05 15:55:52.387516-05	2025-12-05 15:55:52.387516-05
e6a4b31b-8476-4cc0-91d2-1444772db6f9	040d4285-a1e3-4ae2-af4f-9696efd7fb7e	e0af8703-b3b0-4103-a0f6-64336111dc82	80f0567b-ed2a-4795-9b4b-55069a518964	2025-12-05 15:55:52.388035-05	040d4285-a1e3-4ae2-af4f-9696efd7fb7e	\N	\N	f	\N	\N	2025-12-05 15:55:52.388035-05	2025-12-05 15:55:52.388035-05
f864df62-8d06-401a-9e02-83b937ca6869	040d4285-a1e3-4ae2-af4f-9696efd7fb7e	5e18521f-8cf7-4570-b10e-26a432219cfa	416cac7b-8d7e-4e8a-bc65-bdeb0cb9eff1	2025-12-05 15:55:52.389169-05	040d4285-a1e3-4ae2-af4f-9696efd7fb7e	\N	\N	f	\N	\N	2025-12-05 15:55:52.389169-05	2025-12-05 15:55:52.389169-05
1cf3e7aa-1c88-46a9-ab6e-ebad694b9e56	bf8b1a90-6f0d-42f9-afd4-b34fd1d15358	d5a6d085-424d-45c8-8cd4-48f38981af85	434bed03-d869-4a7b-a9af-404d3c394dd7	2025-12-05 15:55:52.390418-05	bf8b1a90-6f0d-42f9-afd4-b34fd1d15358	\N	\N	f	\N	\N	2025-12-05 15:55:52.390418-05	2025-12-05 15:55:52.390418-05
468d2edf-6b2e-4d2e-9cd8-a4cf3e79c0f4	bf8b1a90-6f0d-42f9-afd4-b34fd1d15358	003ce6c1-b1fe-481d-854e-076fbc0882da	ff9e0490-4dd8-4ea5-b706-71c92b5cc84a	2025-12-05 15:55:52.391631-05	bf8b1a90-6f0d-42f9-afd4-b34fd1d15358	\N	\N	f	\N	\N	2025-12-05 15:55:52.391631-05	2025-12-05 15:55:52.391631-05
2b9f145b-433c-4c35-8a72-3458907b8308	bf8b1a90-6f0d-42f9-afd4-b34fd1d15358	799de4e5-f480-4f20-a9ba-d3e8174e4193	7dd3c008-6b79-4d5a-982f-2af0bbbe292d	2025-12-05 15:55:52.392155-05	bf8b1a90-6f0d-42f9-afd4-b34fd1d15358	\N	\N	f	\N	\N	2025-12-05 15:55:52.392155-05	2025-12-05 15:55:52.392155-05
4367dfee-e159-4448-bf5b-3b32ee02aac3	bf8b1a90-6f0d-42f9-afd4-b34fd1d15358	e0af8703-b3b0-4103-a0f6-64336111dc82	0c0e9a41-8e06-4798-8f83-a2e45d262807	2025-12-05 15:55:52.39399-05	bf8b1a90-6f0d-42f9-afd4-b34fd1d15358	\N	\N	f	\N	\N	2025-12-05 15:55:52.39399-05	2025-12-05 15:55:52.39399-05
2bf4732b-1e2e-43ec-8234-6006cad5fafe	bf8b1a90-6f0d-42f9-afd4-b34fd1d15358	5e18521f-8cf7-4570-b10e-26a432219cfa	416cac7b-8d7e-4e8a-bc65-bdeb0cb9eff1	2025-12-05 15:55:52.395952-05	bf8b1a90-6f0d-42f9-afd4-b34fd1d15358	\N	\N	f	\N	\N	2025-12-05 15:55:52.395952-05	2025-12-05 15:55:52.395952-05
\.


--
-- TOC entry 5035 (class 0 OID 21965)
-- Dependencies: 228
-- Data for Name: user_details; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_details (id, user_id, cod_emp_sgd, first_name, last_name, phone, structural_position_id, organic_unit_id, ubigeo_id) FROM stdin;
\.


--
-- TOC entry 5038 (class 0 OID 22034)
-- Dependencies: 231
-- Data for Name: user_module_restrictions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_module_restrictions (id, user_id, module_id, application_id, restriction_type, max_permission_level, reason, expires_at, created_at, created_by, updated_at, updated_by, is_deleted, deleted_at, deleted_by) FROM stdin;
\.


--
-- TOC entry 5024 (class 0 OID 21852)
-- Dependencies: 217
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, email, dni, status, created_at, updated_at, is_deleted, deleted_at, deleted_by) FROM stdin;
a54e954d-0c61-4861-ab94-d49eaf516672	admin@regionayacucho.gob.pe	10101010	active	2025-12-05 15:55:52.342277-05	2025-12-05 15:55:52.342277-05	f	\N	\N
9a322e1b-c428-4ee8-9ac4-5b69168dcded	juan.perez@regionayacucho.gob.pe	11111111	active	2025-12-05 15:55:52.355722-05	2025-12-05 15:55:52.355722-05	f	\N	\N
980014a4-7574-4b2d-b061-0660868b120a	maria.lopez@regionayacucho.gob.pe	22222222	active	2025-12-05 15:55:52.366334-05	2025-12-05 15:55:52.366334-05	f	\N	\N
1d7899ba-7655-4a6a-ab8b-4e361a60115d	carlos.ramos@regionayacucho.gob.pe	33333333	active	2025-12-05 15:55:52.373733-05	2025-12-05 15:55:52.373733-05	f	\N	\N
f3629131-4d9a-40e7-bd2d-7e0d029e3c28	luisa.garcia@regionayacucho.gob.pe	44444444	active	2025-12-05 15:55:52.38-05	2025-12-05 15:55:52.38-05	f	\N	\N
040d4285-a1e3-4ae2-af4f-9696efd7fb7e	pedro.sanchez@regionayacucho.gob.pe	55555555	active	2025-12-05 15:55:52.384957-05	2025-12-05 15:55:52.384957-05	f	\N	\N
bf8b1a90-6f0d-42f9-afd4-b34fd1d15358	nicola.testa@regionayacucho.gob.pe	00000000	active	2025-12-05 15:55:52.389684-05	2025-12-05 15:55:52.389684-05	f	\N	\N
\.


--
-- TOC entry 5048 (class 0 OID 0)
-- Dependencies: 221
-- Name: organic_units_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.organic_units_id_seq', 40, true);


--
-- TOC entry 5049 (class 0 OID 0)
-- Dependencies: 223
-- Name: structural_positions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.structural_positions_id_seq', 142, true);


--
-- TOC entry 5050 (class 0 OID 0)
-- Dependencies: 225
-- Name: ubigeos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ubigeos_id_seq', 1893, true);


--
-- TOC entry 5051 (class 0 OID 0)
-- Dependencies: 227
-- Name: user_details_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_details_id_seq', 1, false);


--
-- TOC entry 4834 (class 2606 OID 21887)
-- Name: application_roles application_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application_roles
    ADD CONSTRAINT application_roles_pkey PRIMARY KEY (id);


--
-- TOC entry 4830 (class 2606 OID 21875)
-- Name: applications applications_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.applications
    ADD CONSTRAINT applications_pkey PRIMARY KEY (id);


--
-- TOC entry 4860 (class 2606 OID 22023)
-- Name: module_role_permissions module_role_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.module_role_permissions
    ADD CONSTRAINT module_role_permissions_pkey PRIMARY KEY (id);


--
-- TOC entry 4836 (class 2606 OID 21903)
-- Name: modules modules_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.modules
    ADD CONSTRAINT modules_pkey PRIMARY KEY (id);


--
-- TOC entry 4838 (class 2606 OID 21926)
-- Name: organic_units organic_units_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organic_units
    ADD CONSTRAINT organic_units_pkey PRIMARY KEY (id);


--
-- TOC entry 4844 (class 2606 OID 21948)
-- Name: structural_positions structural_positions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.structural_positions
    ADD CONSTRAINT structural_positions_pkey PRIMARY KEY (id);


--
-- TOC entry 4850 (class 2606 OID 21961)
-- Name: ubigeos ubigeos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ubigeos
    ADD CONSTRAINT ubigeos_pkey PRIMARY KEY (id);


--
-- TOC entry 4832 (class 2606 OID 21877)
-- Name: applications uni_applications_client_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.applications
    ADD CONSTRAINT uni_applications_client_id UNIQUE (client_id);


--
-- TOC entry 4840 (class 2606 OID 21928)
-- Name: organic_units uni_organic_units_acronym; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organic_units
    ADD CONSTRAINT uni_organic_units_acronym UNIQUE (acronym);


--
-- TOC entry 4842 (class 2606 OID 21930)
-- Name: organic_units uni_organic_units_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organic_units
    ADD CONSTRAINT uni_organic_units_name UNIQUE (name);


--
-- TOC entry 4846 (class 2606 OID 21950)
-- Name: structural_positions uni_structural_positions_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.structural_positions
    ADD CONSTRAINT uni_structural_positions_code UNIQUE (code);


--
-- TOC entry 4848 (class 2606 OID 21952)
-- Name: structural_positions uni_structural_positions_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.structural_positions
    ADD CONSTRAINT uni_structural_positions_name UNIQUE (name);


--
-- TOC entry 4852 (class 2606 OID 21963)
-- Name: ubigeos uni_ubigeos_ubigeo_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ubigeos
    ADD CONSTRAINT uni_ubigeos_ubigeo_code UNIQUE (ubigeo_code);


--
-- TOC entry 4854 (class 2606 OID 21972)
-- Name: user_details uni_user_details_user_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_details
    ADD CONSTRAINT uni_user_details_user_id UNIQUE (user_id);


--
-- TOC entry 4824 (class 2606 OID 21864)
-- Name: users uni_users_dni; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_dni UNIQUE (dni);


--
-- TOC entry 4826 (class 2606 OID 21862)
-- Name: users uni_users_email; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_email UNIQUE (email);


--
-- TOC entry 4858 (class 2606 OID 22001)
-- Name: user_application_roles user_application_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_application_roles
    ADD CONSTRAINT user_application_roles_pkey PRIMARY KEY (id);


--
-- TOC entry 4856 (class 2606 OID 21970)
-- Name: user_details user_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_details
    ADD CONSTRAINT user_details_pkey PRIMARY KEY (id);


--
-- TOC entry 4862 (class 2606 OID 22043)
-- Name: user_module_restrictions user_module_restrictions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_module_restrictions
    ADD CONSTRAINT user_module_restrictions_pkey PRIMARY KEY (id);


--
-- TOC entry 4828 (class 2606 OID 21860)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4863 (class 2606 OID 21888)
-- Name: application_roles fk_application_roles_application; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application_roles
    ADD CONSTRAINT fk_application_roles_application FOREIGN KEY (application_id) REFERENCES public.applications(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- TOC entry 4874 (class 2606 OID 22029)
-- Name: module_role_permissions fk_module_role_permissions_application_role; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.module_role_permissions
    ADD CONSTRAINT fk_module_role_permissions_application_role FOREIGN KEY (application_role_id) REFERENCES public.application_roles(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 4875 (class 2606 OID 22024)
-- Name: module_role_permissions fk_module_role_permissions_module; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.module_role_permissions
    ADD CONSTRAINT fk_module_role_permissions_module FOREIGN KEY (module_id) REFERENCES public.modules(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 4864 (class 2606 OID 21904)
-- Name: modules fk_modules_application; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.modules
    ADD CONSTRAINT fk_modules_application FOREIGN KEY (application_id) REFERENCES public.applications(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 4865 (class 2606 OID 21909)
-- Name: modules fk_modules_children; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.modules
    ADD CONSTRAINT fk_modules_children FOREIGN KEY (parent_id) REFERENCES public.modules(id);


--
-- TOC entry 4866 (class 2606 OID 21931)
-- Name: organic_units fk_organic_units_parent; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organic_units
    ADD CONSTRAINT fk_organic_units_parent FOREIGN KEY (parent_id) REFERENCES public.organic_units(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 4871 (class 2606 OID 22007)
-- Name: user_application_roles fk_user_application_roles_application; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_application_roles
    ADD CONSTRAINT fk_user_application_roles_application FOREIGN KEY (application_id) REFERENCES public.applications(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 4872 (class 2606 OID 22012)
-- Name: user_application_roles fk_user_application_roles_application_role; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_application_roles
    ADD CONSTRAINT fk_user_application_roles_application_role FOREIGN KEY (application_role_id) REFERENCES public.application_roles(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 4873 (class 2606 OID 22002)
-- Name: user_application_roles fk_user_application_roles_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_application_roles
    ADD CONSTRAINT fk_user_application_roles_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 4867 (class 2606 OID 21973)
-- Name: user_details fk_user_details_organic_unit; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_details
    ADD CONSTRAINT fk_user_details_organic_unit FOREIGN KEY (organic_unit_id) REFERENCES public.organic_units(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 4868 (class 2606 OID 21988)
-- Name: user_details fk_user_details_structural_position; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_details
    ADD CONSTRAINT fk_user_details_structural_position FOREIGN KEY (structural_position_id) REFERENCES public.structural_positions(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 4869 (class 2606 OID 21978)
-- Name: user_details fk_user_details_ubigeo; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_details
    ADD CONSTRAINT fk_user_details_ubigeo FOREIGN KEY (ubigeo_id) REFERENCES public.ubigeos(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 4870 (class 2606 OID 21983)
-- Name: user_details fk_user_details_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_details
    ADD CONSTRAINT fk_user_details_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 4876 (class 2606 OID 22044)
-- Name: user_module_restrictions fk_user_module_restrictions_application; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_module_restrictions
    ADD CONSTRAINT fk_user_module_restrictions_application FOREIGN KEY (application_id) REFERENCES public.applications(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 4877 (class 2606 OID 22054)
-- Name: user_module_restrictions fk_user_module_restrictions_module; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_module_restrictions
    ADD CONSTRAINT fk_user_module_restrictions_module FOREIGN KEY (module_id) REFERENCES public.modules(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 4878 (class 2606 OID 22049)
-- Name: user_module_restrictions fk_user_module_restrictions_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_module_restrictions
    ADD CONSTRAINT fk_user_module_restrictions_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


-- Completed on 2025-12-14 09:03:01

--
-- PostgreSQL database dump complete
--

\unrestrict kI76rdtxNdVGDNcKGXrqBO6wiyjEvYAnUoqEUtEJVWkGbKqhVnqzZsjWX3BkFwo

