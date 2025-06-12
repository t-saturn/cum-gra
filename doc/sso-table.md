Explicación del Diseño
Tablas Principales:

users - Información básica de usuarios
user_sessions - Sesiones activas con tokens de autenticación
applications - Servicios registrados (Gmail, YouTube, Drive, etc.)
oauth_tokens - Tokens de acceso específicos por aplicación
user_permissions - Permisos otorgados por usuario/aplicación

Características Clave:

Tokens centralizados: Una sesión principal que autoriza acceso a múltiples servicios
Revocación global: Al cerrar sesión, se invalidan todos los tokens
Auditoría completa: Registro de todas las acciones de autenticación
Seguridad: Tokens con expiración, intentos de login monitoreados

Flujo de Funcionamiento:

Login único: El usuario se autentica una vez en el servidor central
Propagación automática: Los otros servicios reconocen la sesión existente
Verificación continua: Los servicios validan periódicamente con el servidor central
Logout global: Al cerrar sesión, se notifica a todos los servicios

Este diseño te permite replicar la experiencia de Google donde "una sesión = acceso a todo" y "un logout = salida de todo".

```mermaid
sequenceDiagram
participant U as Usuario
participant B as Navegador
participant A1 as Gmail
participant A2 as YouTube
participant AS as Auth Server<br/>(accounts.google.com)
participant DB as Base de Datos

    Note over U,DB: FLUJO DE LOGIN INICIAL

    U->>B: Accede a Gmail
    B->>A1: GET /login
    A1->>B: Redirect a Auth Server
    B->>AS: GET /auth?app=gmail
    AS->>B: Muestra formulario login
    U->>B: Ingresa credenciales
    B->>AS: POST /login {email, password}

    AS->>DB: Valida credenciales
    DB-->>AS: Usuario válido

    AS->>DB: Crea session_token y oauth_token
    DB-->>AS: Tokens creados

    AS->>B: Set cookies + Redirect
    Note right of AS: Cookies: session_token,<br/>refresh_token (httpOnly, secure)

    B->>A1: GET /callback?code=auth_code
    A1->>AS: Valida auth_code
    AS-->>A1: access_token válido
    A1->>B: Usuario logueado en Gmail

    Note over U,DB: ACCESO A SEGUNDO SERVICIO (SSO)

    U->>B: Accede a YouTube
    B->>A2: GET /
    A2->>B: Verifica cookies de auth

    alt Cookies válidas
        B->>AS: Valida session_token
        AS->>DB: Verifica sesión activa
        DB-->>AS: Sesión válida
        AS-->>B: Usuario autorizado
        B->>A2: access_token para YouTube
        A2->>AS: Valida token
        AS-->>A2: Token válido + permisos
        A2->>B: Usuario logueado en YouTube
    else Sin cookies o inválidas
        A2->>B: Redirect a Auth Server
        Note right of A2: Mismo flujo de login
    end

    Note over U,DB: FLUJO DE LOGOUT GLOBAL

    U->>B: Click "Cerrar sesión" en Gmail
    B->>A1: POST /logout
    A1->>AS: POST /logout/global

    AS->>DB: Invalida todas las sesiones del usuario
    Note right of DB: UPDATE user_sessions<br/>SET is_active = false<br/>WHERE user_id = ?

    AS->>DB: Revoca todos los tokens OAuth
    Note right of DB: UPDATE oauth_tokens<br/>SET revoked_at = now()<br/>WHERE user_id = ?

    AS->>B: Clear cookies + confirmación
    B->>A1: Usuario deslogueado

    par Notificación a otros servicios
        AS->>A2: Notifica logout global
        A2->>B: Invalida sesión local
    end

    Note over U,DB: VERIFICACIÓN DE SESIÓN CONTINUA

    loop Cada 5-15 minutos
        A1->>AS: Verifica session_token
        A2->>AS: Verifica session_token
        AS->>DB: Check sesión activa
        alt Sesión válida
            DB-->>AS: Sesión OK
            AS-->>A1: Continúa logueado
            AS-->>A2: Continúa logueado
        else Sesión expirada/inválida
            DB-->>AS: Sesión inválida
            AS-->>A1: Redirect a login
            AS-->>A2: Redirect a login
        end
    end
```
