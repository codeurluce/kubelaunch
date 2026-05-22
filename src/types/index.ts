export type StackType =
  | "nodejs"
  | "nextjs"
  | "react-vite"
  | "python"
  | "go"
  | "docker"
  | "static"
  | "php"
  | "ruby"
  | "java"
  | "rust"
  | "unknown";

export type AppStatus = "running" | "pending" | "error" | "stopped";

export interface DetectResult {
  stack: StackType;
  port: number;
  confidence: "high" | "medium" | "low";
  detected_files: string[];
  framework?: string;
}

export interface EnvVar {
  key: string;
  value: string;
}

export interface AppMetrics {
  cpu: string;
  memory: string;
  uptime: string;
}

export interface Application {
  id: string;
  name: string;
  namespace: string;
  stack: StackType;
  framework?: string;
  port: number;
  replicas: number;
  readyReplicas: number;
  status: AppStatus;
  metrics: AppMetrics;
  createdAt: string;
  repoUrl?: string;
}

export interface LogEntry {
  time: string;
  level: "INFO" | "WARN" | "ERROR" | "OK" | "DEBUG";
  message: string;
}

export interface DeployFormData {
  repoUrl: string;
  appName: string;
  namespace: string;
  replicas: number;
  port: number;
  envVars: EnvVar[];
}
