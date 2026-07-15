# TaskFlow Frontend

Next.js 14 (App Router) client for the TaskFlow API. TypeScript, Tailwind CSS.

## Setup

```bash
npm install
cp .env.local.example .env.local
npm run dev
```

Set `NEXT_PUBLIC_API_URL` in `.env.local` to point at the running backend
(default `http://localhost:8080`).

## Pages

- `/` — landing page, redirects to `/dashboard` if already signed in
- `/login`, `/register` — auth forms
- `/dashboard` — protected task list. Redirects to `/login` if no token is stored

## How auth works

On login, the JWT returned by the API is stored in `localStorage` under
`taskflow_token`. `src/lib/api.ts` attaches it to every request as
`Authorization: Bearer <token>`. The `(dashboard)` route group checks for
the token on mount and redirects unauthenticated visitors to `/login`.

## Build

```bash
npm run build
npm run start
```
