import type { Application, LogEntry, DetectResult, StackType } from "@/types";

export const MOCK_APPS: Application[] = [
  {
    id: "1",
    name: "api-service",
    namespace: "default",
    stack: "nodejs",
    framework: "Express 4.18",
    port: 3000,
    replicas: 2,
    readyReplicas: 2,
    status: "running",
    metrics: { cpu: "124m", memory: "256Mi", uptime: "99.9%" },
    createdAt: "2026-05-08T10:00:00Z",
    repoUrl: "https://github.com/codeurluce/api-service",
  },
  {
    id: "2",
    name: "frontend-app",
    namespace: "default",
    stack: "nextjs",
    framework: "Next.js 14",
    port: 3000,
    replicas: 1,
    readyReplicas: 1,
    status: "running",
    metrics: { cpu: "68m", memory: "128Mi", uptime: "100%" },
    createdAt: "2026-05-07T14:30:00Z",
    repoUrl: "https://github.com/codeurluce/frontend-app",
  },
  {
    id: "3",
    name: "worker-service",
    namespace: "prod",
    stack: "python",
    framework: "FastAPI",
    port: 8000,
    replicas: 3,
    readyReplicas: 3,
    status: "running",
    metrics: { cpu: "310m", memory: "512Mi", uptime: "98.2%" },
    createdAt: "2026-05-06T09:15:00Z",
  },
  {
    id: "4",
    name: "auth-service",
    namespace: "staging",
    stack: "nodejs",
    framework: "Express + JWT",
    port: 4000,
    replicas: 1,
    readyReplicas: 0,
    status: "pending",
    metrics: { cpu: "0m", memory: "0Mi", uptime: "—" },
    createdAt: "2026-05-09T16:00:00Z",
  },
  {
    id: "5",
    name: "go-gateway",
    namespace: "prod",
    stack: "go",
    framework: "Gin",
    port: 8080,
    replicas: 2,
    readyReplicas: 2,
    status: "running",
    metrics: { cpu: "42m", memory: "64Mi", uptime: "100%" },
    createdAt: "2026-05-05T11:00:00Z",
  },
];

export const MOCK_LOGS: LogEntry[] = [
  { time: "10:42:01.234", level: "INFO", message: "Server listening on port 3000" },
  { time: "10:42:03.891", level: "INFO", message: "GET /api/health → 200 (2ms)" },
  { time: "10:42:07.102", level: "OK", message: "Database connection established" },
  { time: "10:42:12.445", level: "INFO", message: "POST /api/users → 201 (14ms)" },
  { time: "10:42:18.003", level: "WARN", message: "Slow query detected: 320ms on users table" },
  { time: "10:42:22.771", level: "INFO", message: "GET /api/users/42 → 200 (5ms)" },
  { time: "10:42:29.118", level: "ERROR", message: "Missing env var DATABASE_URL — check your secrets" },
  { time: "10:42:30.002", level: "OK", message: "Pod restarted successfully (1 restart)" },
  { time: "10:42:35.550", level: "INFO", message: "GET /api/products → 200 (8ms)" },
  { time: "10:42:40.223", level: "DEBUG", message: "Cache miss for key: user:42:profile" },
];

export const STACK_CONFIG: Record<string, { label: string; color: string; bg: string; files: string[] }> = {
  // ✅ Détection réussie — toujours vert
  nodejs:       { label: "Node.js",         color: "#4ade80", bg: "#052e16", files: ["package.json"] },
  nextjs:       { label: "Next.js",         color: "#4ade80", bg: "#052e16", files: ["package.json", "next.config.js"] },
  "react-vite": { label: "React + Vite",    color: "#4ade80", bg: "#052e16", files: ["package.json", "vite.config.ts"] },
  python:       { label: "Python",          color: "#4ade80", bg: "#052e16", files: ["requirements.txt"] },
  go:           { label: "Go",              color: "#4ade80", bg: "#052e16", files: ["go.mod"] },
  docker:       { label: "Docker",          color: "#4ade80", bg: "#052e16", files: ["Dockerfile"] },
  static:       { label: "Static Site",     color: "#4ade80", bg: "#052e16", files: ["index.html"] },
  php:          { label: "PHP",             color: "#4ade80", bg: "#052e16", files: ["composer.json"] },
  ruby:         { label: "Ruby",            color: "#4ade80", bg: "#052e16", files: ["Gemfile"] },
  java:         { label: "Java",            color: "#4ade80", bg: "#052e16", files: ["pom.xml"] },
  rust:         { label: "Rust",            color: "#4ade80", bg: "#052e16", files: ["Cargo.toml"] },
  unknown:      { label: "Unknown",         color: "#94a3b8", bg: "#1e293b", files: [] },
};

export function detectStackFromUrl(url: string): DetectResult {
  if (url.includes("next") || url.includes("frontend")) {
    return { stack: "nextjs", port: 3000, confidence: "high", detected_files: ["package.json", "next.config.js"], framework: "Next.js 14" };
  }
  if (url.includes("python") || url.includes("fastapi") || url.includes("django")) {
    return { stack: "python", port: 8000, confidence: "high", detected_files: ["requirements.txt"], framework: "FastAPI" };
  }
  if (url.includes("go") || url.includes("gateway") || url.includes("grpc")) {
    return { stack: "go", port: 8080, confidence: "high", detected_files: ["go.mod"], framework: "Gin" };
  }
  return { stack: "nodejs", port: 3000, confidence: "high", detected_files: ["package.json"], framework: "Express 4.18" };
}

export function cn(...classes: (string | undefined | false | null)[]): string {
  return classes.filter(Boolean).join(" ");
}
