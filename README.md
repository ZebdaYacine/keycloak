# Next.js + Keycloak demo

This demo follows the same architecture as the Medium article: Keycloak runs locally with PostgreSQL, and the Next.js app uses Auth.js / NextAuth with the Keycloak OIDC provider.

## 1. Start Keycloak

```bash
docker compose up -d
```

Open Keycloak:

- URL: http://localhost:8080
- Admin username: `admin`
- Admin password: `admin`

The realm is auto-imported from `realms/realm-export.json`.

Imported realm:

- Realm: `nextjs-kc`
- Client ID: `nextjs-web`
- Client secret: `change-me-nextjs-secret`
- Demo user: `demo`
- Demo password: `demo1234`

## 2. Install dependencies

```bash
pnpm install
```

## 3. Generate a real Auth.js secret

For production or serious local tests, replace the demo secret:

```bash
npx auth secret
```

Copy the generated value into `.env.local` as `AUTH_SECRET`.

## 4. Run Next.js

```bash
pnpm dev
```

Open:

```text
http://localhost:3000
```

Click **Sign in with Keycloak**.

## Important files

```text
.
├── docker-compose.yml
├── realms/realm-export.json
├── auth.ts
├── middleware.ts
├── app/api/auth/[...nextauth]/route.ts
├── app/api/auth/login/route.ts
├── app/api/auth/logout/route.ts
├── app/_lib/actions/auth.ts
├── app/_lib/components/Header.tsx
├── app/page.tsx
└── app/dashboard/page.tsx
```

## Notes

- `/dashboard` is protected.
- `/api/auth/[...nextauth]` handles Auth.js callbacks.
- Keycloak callback URL is `http://localhost:3000/api/auth/callback/keycloak`.
- Web origin is `http://localhost:3000`.
