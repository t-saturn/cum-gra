<div align="center">
  <h1>🌟 Central User Manager (CUM) - Client 🚀</h1>
  <p>🎨 Dashboard moderno para la gestión de usuarios, roles y permisos en la intranet</p>
  <a href="https://github.com/t-saturn/central-user-manager-client">
    <img src="https://img.shields.io/github/stars/t-saturn/central-user-manager-client?style=social" alt="GitHub stars" />
  </a>
  <a href="https://github.com/t-saturn/central-user-manager-client/LICENSE">
    <img src="https://img.shields.io/github/license/t-saturn/central-user-manager-client?color=blue" alt="License" />
  </a>
  <br /><br />
  <img src="https://img.shields.io/badge/Next.js-15.3.4-black?logo=next.js" alt="Next.js" />
  <img src="https://img.shields.io/badge/TypeScript-5.x-blue?logo=typescript" alt="TypeScript" />
  <img src="https://img.shields.io/badge/shadcn/ui-latest-rose?logo=shadcnui" alt="shadcn/ui" />
  <img src="https://img.shields.io/badge/Tailwind_CSS-4.x-38B2AC?logo=tailwindcss" alt="Tailwind CSS" />
  <img src="https://img.shields.io/badge/Bun-1.2.x-yellow?logo=bun" alt="Bun" />
  <img src="https://img.shields.io/badge/React_Query-5.x-red?logo=react-query" alt="React Query" />
</div>

---

## 📖 Sobre el Proyecto

**Central User Manager (CUM) - Client** es un dashboard elegante y dinámico para gestionar la intranet de tu organización. Desarrollado con **Next.js 15.3**, este proyecto proporciona una interfaz visual moderna para administrar usuarios, aplicaciones, roles, módulos y permisos, según el esquema de tu organización. 🌐

📍 **Repositorio Frontend**: [github.com/t-saturn/central-user-manager-client](https://github.com/t-saturn/central-user-manager-client)  
📍 **Repositorio Backend**: [github.com/t-saturn/central-user-manager](https://github.com/t-saturn/central-user-manager)  
📍 **Repositorio Auth Service**: [github.com/t-saturn/auth-service](https://github.com/t-saturn/auth-service)

> ⚠️ **Nota**: Este es **solo el frontend**. Necesita el backend (Go + PostgreSQL) y el servicio de autenticación SSO para funcionar completamente.

---

## ✨ Características Principales

- 🖥 **Dashboard Intuitivo**: Interfaz limpia para gestionar usuarios, roles y módulos.
- 👤 **Gestión de Usuarios**: Creación, edición, verificación de email/teléfono y borrado lógico.
- 🔒 **Roles y Permisos Granulares**: Asignación de roles por aplicación con restricciones específicas.
- 🧩 **Módulos Dinámicos**: Configuración de menús y rutas asociadas a aplicaciones.
- 🌍 **Soporte Multi-idioma**: Preparado para internacionalización (i18n).
- 📱 **Diseño Responsivo**: Optimizado para móviles, tablets y escritorios.
- 🔑 **Integración SSO**: Autenticación centralizada mediante OAuth 2.0.

---

## 🛠 Stack Tecnológico

| 🧰 Herramienta         | 📌 Versión | 📝 Descripción                              |
|------------------------|------------|---------------------------------------------|
| **Next.js**            | 15.3       | Framework React con App Router              |
| **TypeScript**         | 5.x        | Tipado estático para código robusto         |
| **shadcn/ui**          | Latest     | Componentes UI accesibles y personalizables |
| **Tailwind CSS**       | 4.x        | Estilizado utility-first                    |
| **Bun**                | 1.2.x        | Gestor de paquetes y runtime ultrarrápido   |
| **React Query**        | 5.x        | Gestión de datos asíncronos                 |
| **Axios**              | 1.x        | Cliente HTTP para consumir la API           |
| **ESLint + Prettier**  | Latest     | Linting y formateo de código                |

---

## 📋 Requisitos Previos

Asegúrate de tener instalados:

- 🟢 **Bun** (versión 1.2.x o superior) - [Instalar Bun](https://bun.sh/)
- 🟢 **Git** para clonar el repositorio
- 🔗 Acceso al **backend** ([central-user-manager](https://github.com/t-saturn/central-user-manager))
- 🔗 Servicio de **SSO** configurado ([auth-service](https://github.com/t-saturn/auth-service))

> 🔐 **Advertencia**: Este frontend depende del backend y del servicio SSO para su funcionamiento completo.

---

## ⚙️ Instalación

1. **Clona el repositorio**:
   ```bash
   git clone https://github.com/t-saturn/central-user-manager-client.git
   cd central-user-manager-client
   ```

2. **Instala las dependencias** con Bun:
   ```bash
   bun install
   ```

3. **Configura las variables de entorno**:
   Crea un archivo `.env.local` en la raíz del proyecto:
   ```env
   NEXT_PUBLIC_API_URL=http://localhost:8080/api
   NEXT_PUBLIC_SSO_AUTH_URL=https://sso.example.com
   NEXT_PUBLIC_CLIENT_ID=your-client-id
   ```

   > 🚨 **Seguridad**: Nunca compartas tus credenciales de SSO o `client_id` en repositorios públicos.

4. **Verifica el código**:
   ```bash
   bun lint
   bun format
   ```

---

## 🚀 Ejecución

Inicia el servidor de desarrollo:

```bash
bun dev
```

Abre [http://localhost:3000](http://localhost:3000) en tu navegador. 🌐

Para producción:

```bash
bun build
bun start
```

---

## 📂 Estructura del Proyecto

```plaintext
├── /src
│   ├── /app              # Rutas y layouts (Next.js App Router)
│   ├── /components       # Componentes React (incluye shadcn/ui)
│   ├── /hooks            # Hooks personalizados
│   ├── /lib              # Utilidades (API, clientes, etc.)
│   ├── /styles           # Configuración de Tailwind y estilos globales
│   ├── /types            # Tipos de TypeScript para la API
├── /public               # Archivos estáticos (imágenes, favicon)
├── .env.local            # Variables de entorno
├── next.config.mjs       # Configuración de Next.js
├── tsconfig.json         # Configuración de TypeScript
└── README.md             # ¡Este archivo!
```

---

## 🔐 Autenticación

El frontend se integra con un **servicio SSO externo** ([auth-service](https://github.com/t-saturn/auth-service)) mediante OAuth 2.0:

1. Redirige al usuario al SSO para autenticación.
2. Recibe un token JWT tras un login exitoso.
3. Utiliza el token para autenticar peticiones al backend.

> 🔑 **Nota**: Asegúrate de configurar `NEXT_PUBLIC_SSO_AUTH_URL` y `NEXT_PUBLIC_CLIENT_ID` correctamente.

---

## 📡 Integración con Backend

El frontend consume una API REST desarrollada en **Go** con **PostgreSQL** ([central-user-manager](https://github.com/t-saturn/central-user-manager)). Las entidades principales incluyen:

- **Usuarios**: Gestión con validación de DNI y correo.
- **Aplicaciones**: Configuración de client_id, secret y callbacks.
- **Roles y Permisos**: Permisos granulares por módulo.
- **Módulos**: Menús dinámicos y rutas.
- **Restricciones**: Limitaciones específicas por usuario.

Verifica que el backend esté accesible en `NEXT_PUBLIC_API_URL`.

---

## 🤝 Contribución

¡Contribuye al proyecto! Sigue estos pasos:

1. 🍴 Haz un fork del repositorio.
2. 🌿 Crea una rama (`git checkout -b feat/nueva-funcionalidad`).
3. 💻 Implementa tus cambios respetando ESLint y Prettier.
4. ✅ Prueba localmente (`bun test`).
5. 🚀 Envía un pull request con una descripción clara.

> 📝 Usa commits descriptivos: `feat: add user profile page`, `fix: resolve login error`.

---

## ⚠️ Notas Importantes

- 🚫 Este es **solo el frontend**. Requiere el backend y SSO para funcionar.
- 🛡️ Sanitiza entradas para prevenir XSS y otros ataques.
- 📈 Optimiza el rendimiento con herramientas como Vercel Analytics.
- 🌐 Prueba en múltiples navegadores para compatibilidad.
- 🎨 Usa **shadcn/ui** para consistencia en los componentes UI.

---

## 📄 Licencia

Licenciado bajo **MIT**. Consulta [LICENSE](https://github.com/t-saturn/central-user-manager-client/blob/mainLICENSE) para más detalles.
