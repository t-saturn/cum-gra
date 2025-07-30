# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| v1.x    | :white_check_mark: |
| v0.x    | :x:                |

## Reporting a Vulnerability

If you discover a security vulnerability in this project, please report it responsibly by following these steps:

1. **Do not create a public issue** for the vulnerability.
2. Send an email to the maintainers at: `security@example.com` (replace with actual contact).
3. Include the following information:
   - A detailed description of the vulnerability.
   - Steps to reproduce the issue.
   - Potential impact and severity.
   - Any suggested fixes or mitigations.

We will respond promptly and work with you to resolve the issue.

## Security Best Practices

This project follows these security best practices:

- Passwords are hashed securely using strong algorithms.
- Access and refresh tokens are generated and validated securely.
- Sessions and tokens have proper lifecycle management including revocation.
- Audit logs are maintained for authentication attempts and token activities.
- Rate limiting and CAPTCHA are used to prevent brute force attacks.
- All sensitive operations require proper authentication and authorization.

## Dependencies

We regularly update dependencies to patch known vulnerabilities. Please ensure you keep your environment up to date.

## Contribution Guidelines

When contributing code, please:

- Avoid introducing security vulnerabilities.
- Follow secure coding standards.
- Include tests for security-related features.
- Review dependencies for security issues.
