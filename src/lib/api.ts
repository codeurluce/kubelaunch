const API_URL = "http://localhost:8080/api";

export type DeployStep = {
  id: string;
  label: string;
  status: "pending" | "running" | "success" | "error";
  message?: string;
};

export type DeployResponse = {
  success: boolean;
  url?: string;
  error?: string;
  logs?: string[];
  steps?: DeployStep[];
};

export async function detectStack(repoUrl: string) {
  const res = await fetch(`${API_URL}/detect`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ repoUrl }),
  });

  const data = await res.json();

  if (!res.ok) {
    throw new Error(data.error || "Erreur de détection");
  }

  return data;
}

export async function listApps() {
  const res = await fetch(`${API_URL}/apps`);

  const data = await res.json();

  if (!res.ok) {
    throw new Error(data.error || "Erreur liste apps");
  }

  return data;
}

export async function deployApp(data: {
  repoUrl: string;
  appName: string;
  namespace: string;
  replicas: number;
  port: number;
  stack: string;
  envVars: Record<string, string>;
}): Promise<DeployResponse> {
  const res = await fetch(`${API_URL}/deploy`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  const result: DeployResponse = await res.json();

  /*
    Even if HTTP 200,
    we check for success backend
  */
  if (!res.ok || !result.success) {
    throw new Error(result.error || "Erreur déploiement");
  }

  return result;
}
