Google utiliza un sistema de **Single Sign-On (SSO)** centralizado que se basa en varios componentes técnicos clave:

## Arquitectura Central

**Google Identity Platform** es el núcleo del sistema. Cuando te autentificas en cualquier servicio de Google (Gmail, YouTube, Drive, etc.), no te estás autenticando directamente con ese servicio, sino con el sistema central de identidad de Google.

## Flujo de Autenticación

Cuando inicias sesión, Google genera un **token de autenticación** que se almacena en tu navegador (generalmente como cookies seguras). Este token contiene información encriptada sobre tu identidad y permisos. Todos los servicios de Google confían en este token central para validar tu identidad.

## Mecanismo de Cookies y Dominios

Google utiliza cookies que se comparten entre sus diferentes dominios. Aunque Gmail esté en `gmail.com` y YouTube en `youtube.com`, ambos pueden acceder a las cookies de autenticación almacenadas para el dominio `google.com` y `accounts.google.com`.

## Comunicación Entre Servicios

Los servicios de Google se comunican constantemente con el servidor central de autenticación para verificar el estado de tu sesión. Utilizan protocolos como **OAuth 2.0** y **OpenID Connect** para este intercambio de información.

## Cierre de Sesión Unificado

Cuando cierras sesión en cualquier servicio, este envía una señal al sistema central de Google Identity, que inmediatamente invalida tu token de autenticación. Todos los demás servicios reciben esta notificación y te desconectan automáticamente.

## Ventajas del Sistema

Este enfoque permite una experiencia fluida para el usuario, mayor seguridad centralizada, y facilita la gestión de permisos y accesos desde un punto único de control.

Es similar a cómo funcionan otros grandes proveedores como Microsoft con su sistema de cuentas unificadas o Apple con su Apple ID.
