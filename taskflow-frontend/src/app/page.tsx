"use client";

import Link from "next/link";
import { useEffect } from "react";
import { useRouter } from "next/navigation";

import { authService } from "@/services/auth.service";

export default function Home() {
  const router = useRouter();

  useEffect(() => {
    if (authService.isAuthenticated()) {
      router.replace("/dashboard");
    }
  }, [router]);

  return (
    <main className="flex min-h-screen flex-col items-center justify-center px-6 text-center">
      <span className="font-mono text-xs uppercase tracking-[0.3em] text-ink-soft">
        01 · Register 02 · Login 03 · Ship
      </span>
      <h1 className="mt-4 max-w-xl font-display text-5xl font-semibold leading-tight text-ink">
        Clear your list. Keep your flow.
      </h1>
      <p className="mt-4 max-w-md text-ink-soft">
        TaskFlow is a focused task manager. Sign in to see what is due, what is
        done, and what is next.
      </p>
      <div className="mt-8 flex gap-3">
        <Link
          href="/login"
          className="rounded-md bg-ink px-5 py-2.5 text-sm font-semibold text-paper hover:bg-ink-soft"
        >
          Sign in
        </Link>
        <Link
          href="/register"
          className="rounded-md border border-line px-5 py-2.5 text-sm font-semibold text-ink hover:bg-line/40"
        >
          Create account
        </Link>
      </div>
    </main>
  );
}
