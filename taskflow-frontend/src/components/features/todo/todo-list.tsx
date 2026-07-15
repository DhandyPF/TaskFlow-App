"use client";

import { FormEvent, useEffect, useMemo, useState } from "react";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { todoService } from "@/services/todo.service";
import { Priority, Todo } from "@/types/todo";
import { TodoItem } from "./todo-item";

type Filter = "all" | "pending" | "completed";

export function TodoList() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [filter, setFilter] = useState<Filter>("all");
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [priority, setPriority] = useState<Priority>("MEDIUM");
  const [dueDate, setDueDate] = useState("");
  const [formOpen, setFormOpen] = useState(false);

  async function loadTodos() {
    setLoading(true);
    setError(null);
    try {
      const params =
        filter === "all" ? undefined : { is_completed: filter === "completed" };
      const res = await todoService.list(params);
      setTodos(res.data || []);
    } catch {
      setError("Could not load your tasks. Check that the API is running.");
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    loadTodos();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [filter]);

  async function handleCreate(e: FormEvent) {
    e.preventDefault();
    if (!title.trim()) return;
    try {
      await todoService.create({
        title,
        description,
        priority,
        due_date: dueDate ? new Date(dueDate).toISOString() : null,
      });
      setTitle("");
      setDescription("");
      setPriority("MEDIUM");
      setDueDate("");
      setFormOpen(false);
      loadTodos();
    } catch {
      setError("Could not create the task. Try again.");
    }
  }

  async function handleToggle(id: string, completed: boolean) {
    setTodos((prev) =>
      prev.map((t) => (t.id === id ? { ...t, is_completed: completed } : t)),
    );
    try {
      await todoService.update(id, { is_completed: completed });
    } catch {
      loadTodos();
    }
  }

  async function handleDelete(id: string) {
    setTodos((prev) => prev.filter((t) => t.id !== id));
    try {
      await todoService.remove(id);
    } catch {
      loadTodos();
    }
  }

  const completedCount = useMemo(
    () => todos.filter((t) => t.is_completed).length,
    [todos],
  );
  const total = todos.length;
  const momentum = total === 0 ? 0 : Math.round((completedCount / total) * 100);

  return (
    <div className="mx-auto w-full max-w-2xl">
      {/* Signature element: momentum tally bar */}
      <div className="mb-8">
        <div className="flex items-baseline justify-between font-mono text-xs text-ink-soft">
          <span>TODAY&apos;S MOMENTUM</span>
          <span>
            {completedCount}/{total} done · {momentum}%
          </span>
        </div>
        <div className="mt-2 flex h-2 w-full gap-0.5 overflow-hidden rounded-full bg-line">
          {Array.from({ length: Math.max(total, 1) }).map((_, i) => (
            <div
              key={i}
              className={`h-full flex-1 ${
                i < completedCount ? "bg-ink" : "bg-transparent"
              }`}
            />
          ))}
        </div>
      </div>

      <div className="mb-6 flex items-center justify-between">
        <div className="flex gap-1 rounded-md border border-line bg-white p-1">
          {(["all", "pending", "completed"] as Filter[]).map((f) => (
            <button
              key={f}
              onClick={() => setFilter(f)}
              className={`rounded px-3 py-1 text-xs font-mono uppercase tracking-wide transition-colors ${
                filter === f
                  ? "bg-ink text-paper"
                  : "text-ink-soft hover:bg-line/50"
              }`}
            >
              {f}
            </button>
          ))}
        </div>
        <Button
          onClick={() => setFormOpen((v) => !v)}
          variant={formOpen ? "ghost" : "primary"}
        >
          {formOpen ? "Cancel" : "+ New task"}
        </Button>
      </div>

      {formOpen && (
        <form
          onSubmit={handleCreate}
          className="mb-6 flex flex-col gap-4 rounded-lg border border-line bg-white p-5"
        >
          <Input
            id="title"
            label="Title"
            required
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            placeholder="Fix auth middleware bug"
          />
          <Input
            id="description"
            label="Description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            placeholder="Optional detail"
          />
          <div className="grid grid-cols-2 gap-4">
            <div className="flex flex-col gap-1.5">
              <label className="font-mono text-xs uppercase tracking-wider text-ink-soft">
                Priority
              </label>
              <select
                value={priority}
                onChange={(e) => setPriority(e.target.value as Priority)}
                className="rounded-md border border-line bg-white px-3 py-2 text-sm text-ink outline-none focus:border-ink focus:ring-1 focus:ring-ink"
              >
                <option value="LOW">Low</option>
                <option value="MEDIUM">Medium</option>
                <option value="HIGH">High</option>
              </select>
            </div>
            <Input
              id="due_date"
              label="Due date"
              type="date"
              value={dueDate}
              onChange={(e) => setDueDate(e.target.value)}
            />
          </div>
          <Button type="submit">Add task</Button>
        </form>
      )}

      {error && <p className="mb-4 text-sm text-coral">{error}</p>}

      <div className="rounded-lg border border-line bg-white px-5">
        {loading ? (
          <p className="py-8 text-center font-mono text-sm text-ink-soft">
            Loading tasks...
          </p>
        ) : todos.length === 0 ? (
          <p className="py-8 text-center font-mono text-sm text-ink-soft">
            No tasks here. Add one to get moving.
          </p>
        ) : (
          todos.map((todo) => (
            <TodoItem
              key={todo.id}
              todo={todo}
              onToggle={handleToggle}
              onDelete={handleDelete}
            />
          ))
        )}
      </div>
    </div>
  );
}
