"use client";

import { Todo } from "@/types/todo";
import { formatDueDate, isOverdue } from "@/utils/date";

const priorityColor: Record<string, string> = {
  HIGH: "bg-coral",
  MEDIUM: "bg-marigold",
  LOW: "bg-flow",
};

interface TodoItemProps {
  todo: Todo;
  onToggle: (id: string, completed: boolean) => void;
  onDelete: (id: string) => void;
}

export function TodoItem({ todo, onToggle, onDelete }: TodoItemProps) {
  const overdue = !todo.is_completed && isOverdue(todo.due_date);

  return (
    <div className="group flex items-start gap-3 border-b border-line py-4 last:border-b-0">
      <button
        aria-label={todo.is_completed ? "Mark as pending" : "Mark as done"}
        onClick={() => onToggle(todo.id, !todo.is_completed)}
        className={`mt-1 h-5 w-5 shrink-0 rounded-full border-2 transition-colors ${
          todo.is_completed ? "border-ink bg-ink" : "border-ink-soft bg-transparent"
        }`}
      />

      <div className="flex-1">
        <div className="flex items-center gap-2">
          <span className={`h-1.5 w-1.5 rounded-full ${priorityColor[todo.priority]}`} />
          <h3
            className={`font-display text-lg leading-snug ${
              todo.is_completed ? "text-ink-soft/60 line-through" : "text-ink"
            }`}
          >
            {todo.title}
          </h3>
        </div>
        {todo.description && (
          <p className="mt-1 text-sm text-ink-soft">{todo.description}</p>
        )}
        <p className={`mt-2 font-mono text-xs ${overdue ? "text-coral" : "text-ink-soft"}`}>
          {formatDueDate(todo.due_date)} {overdue && "· overdue"}
        </p>
      </div>

      <button
        onClick={() => onDelete(todo.id)}
        aria-label="Delete task"
        className="opacity-0 transition-opacity group-hover:opacity-100 text-ink-soft hover:text-coral"
      >
        ✕
      </button>
    </div>
  );
}
