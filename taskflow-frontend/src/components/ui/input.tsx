import { InputHTMLAttributes } from "react";

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  label: string;
}

export function Input({ label, id, className = "", ...props }: InputProps) {
  return (
    <div className="flex flex-col gap-1.5">
      <label htmlFor={id} className="font-mono text-xs uppercase tracking-wider text-ink-soft">
        {label}
      </label>
      <input
        id={id}
        className={`rounded-md border border-line bg-white px-3 py-2 text-sm text-ink outline-none focus:border-ink focus:ring-1 focus:ring-ink ${className}`}
        {...props}
      />
    </div>
  );
}
