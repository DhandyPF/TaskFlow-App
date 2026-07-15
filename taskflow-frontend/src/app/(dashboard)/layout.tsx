"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

import { authService } from "@/services/auth.service";
import { Button } from "@/components/ui/button";

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const router = useRouter();
  const [checked, setChecked] = useState(false);

  useEffect(() => {
    if (!authService.isAuthenticated()) {
      router.replace("/login");
      return;
    }
    setChecked(true);
  }, [router]);

  function handleLogout() {
    authService.logout();
    router.replace("/login");
  }

  if (!checked) {
    return (
      <div className="flex min-h-screen items-center justify-center font-mono text-sm text-ink-soft">
        Checking session...
      </div>
    );
  }

  return (
    <div className="min-h-screen">
      <header className="border-b border-line bg-white">
        <div className="mx-auto flex max-w-2xl items-center justify-between px-6 py-4">
          <span className="font-display text-lg font-semibold text-ink">
            TaskFlow
          </span>
          <Button variant="ghost" onClick={handleLogout}>
            Sign out
          </Button>
        </div>
      </header>
      <div className="px-6 py-10">{children}</div>
    </div>
  );
}
