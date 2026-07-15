import Link from "next/link";

import { RegisterForm } from "@/components/features/auth/register-form";

export default function RegisterPage() {
  return (
    <main className="flex min-h-screen items-center justify-center px-6">
      <div className="w-full max-w-sm">
        <span className="font-mono text-xs uppercase tracking-[0.3em] text-ink-soft">
          TaskFlow
        </span>
        <h1 className="mt-2 font-display text-3xl font-semibold text-ink">Create your account</h1>
        <p className="mt-1 text-sm text-ink-soft">Start clearing your list today.</p>

        <div className="mt-6">
          <RegisterForm />
        </div>

        <p className="mt-6 text-center text-sm text-ink-soft">
          Already have an account?{" "}
          <Link href="/login" className="font-semibold text-ink underline">
            Sign in
          </Link>
        </p>
      </div>
    </main>
  );
}
