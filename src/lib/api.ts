const API_URL = "http://localhost:8080/api";

export async function detectStack(repoUrl: string) {
  const res = await fetch(`${API_URL}/detect`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ repoUrl }),
  });
  const data = await res.json();
  console.log(data);

  if (!res.ok) {
    // Afficher le vrai message d'erreur du backend
    throw new Error(data.error || "Erreur de détection");
  }

  return data;
}

export async function listApps() {
  const res = await fetch(`${API_URL}/apps`);
  const data = await res.json();
  if (!res.ok) throw new Error(data.error || "Erreur liste apps");
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
}) {
  const res = await fetch(`${API_URL}/deploy`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  const result = await res.json();
  if (!res.ok) throw new Error(result.error || "Erreur déploiement");
  return result;
}