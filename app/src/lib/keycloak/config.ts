export const keycloakConfig = {
  baseUrl: process.env.KEYCLOAK_BASE_URL || 'http://localhost:8080',
  realmName: process.env.KEYCLOAK_REALM || 'gore-ayacucho',
  
  // Solo para NextAuth (ya los tienes)
  clientId: process.env.KEYCLOAK_CLIENT_ID!,
  clientSecret: process.env.KEYCLOAK_CLIENT_SECRET,
  issuer: process.env.KEYCLOAK_ISSUER!,
} as const;