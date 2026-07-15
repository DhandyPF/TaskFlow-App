import Link from "next/link";

import { LoginForm } from "@/components/features/auth/login-form";

export default function LoginPage() {
  return (
    <main className="flex min-h-screen items-center justify-center px-6">
      <div className="w-full max-w-sm">
        <span className="font-mono text-xs uppercase tracking-[0.3em] text-ink-soft">
          TaskFlow
        </span>
        <h1 className="mt-2 font-display text-3xl font-semibold text-ink">Welcome back</h1>
        <p className="mt-1 text-sm text-ink-soft">Sign in to see your tasks.</p>

        <div className="mt-6">
          <LoginForm />
        </div>

        <p className="mt-6 text-center text-sm text-ink-soft">
          No account yet?{" "}
          <Link href="/register" className="font-semibold text-ink underline">
            Create one
          </Link>
        </p>
      </div>
    </main>
  );
}
