import { ButtonHTMLAttributes } from "react";

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "primary" | "ghost" | "danger";
}

export function Button({ variant = "primary", className = "", ...props }: ButtonProps) {
  const base =
    "inline-flex items-center justify-center rounded-md px-4 py-2 text-sm font-semibold transition-colors disabled:opacity-50 disabled:cursor-not-allowed";

  const variants: Record<string, string> = {
    primary: "bg-ink text-paper hover:bg-ink-soft",
    ghost: "bg-transparent text-ink border border-line hover:bg-line/40",
    danger: "bg-transparent text-coral border border-coral/40 hover:bg-coral/10",
  };

  return <button className={`${base} ${variants[variant]} ${className}`} {...props} />;
}
