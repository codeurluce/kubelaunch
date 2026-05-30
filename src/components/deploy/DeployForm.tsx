"use client";

import { useState } from "react";
import { Rocket, Loader2, CheckCircle2 } from "lucide-react";

import type { DetectResult, EnvVar } from "@/types";

import { detectStack, deployApp } from "@/lib/api";

import RepoInput from "./RepoInput";
import StackDetectedCard from "./StackDetectedCard";
import GeneratedResources from "./GeneratedResources";
import ExposureCard from "./ExposureCard";
import EnvVarsEditor from "./EnvVarsEditor";
import DeployPipeline from "./DeployPipeline";
import SuccessCard from "./SuccessCard";
import type { DeployStep } from "@/lib/api";


const INITIAL_STEPS: DeployStep[] = [
  {
    id: "clone",
    label: "Cloning repository",
    status: "pending",
  },
  {
    id: "dockerfile",
    label: "Generating Dockerfile",
    status: "pending",
  },
  {
    id: "manifests",
    label: "Generating Kubernetes manifests",
    status: "pending",
  },
  {
    id: "build",
    label: "Building Docker image",
    status: "pending",
  },
  {
    id: "push",
    label: "Pushing image",
    status: "pending",
  },
  {
    id: "deploy",
    label: "Deploying resources to Kubernetes",
    status: "pending",
  },
  {
    id: "ready",
    label: "Waiting for pod readiness",
    status: "pending",
  },
];

export function DeployForm() {
  const [repoUrl, setRepoUrl] = useState("");
  const [detecting, setDetecting] = useState(false);
  const [deploying, setDeploying] = useState(false);
  const [deployed, setDeployed] = useState(false);

  const [error, setError] = useState<string | null>(null);

  const [detected, setDetected] = useState<DetectResult | null>(null);

  const [envVars, setEnvVars] = useState<EnvVar[]>([
    {
      key: "",
      value: "",
    },
  ]);

  const [appName, setAppName] = useState("");

  const [steps, setSteps] = useState<DeployStep[]>(INITIAL_STEPS);

  const [liveUrl, setLiveUrl] = useState<string | null>(null);

  const [showLogs, setShowLogs] = useState(false);

  const [logs, setLogs] = useState<string[]>([]);

  const updateStep = (
    stepId: string,
    status: "pending" | "running" | "success",
  ) => {
    setSteps((prev) =>
      prev.map((s) =>
        s.id === stepId
          ? {
              ...s,
              status,
            }
          : s,
      ),
    );
  };

  const handleDetect = async () => {
    if (!repoUrl.trim()) return;

    try {
      setDetecting(true);

      const result = await detectStack(repoUrl);

      setDetected(result);

      const repoName =
        repoUrl.split("/").pop()?.replace(".git", "") || "my-app";

      setAppName(repoName);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Unknown error");
    } finally {
      setDetecting(false);
    }
  };

const handleDeploy = async () => {
  if (!detected) return;

  try {
    setDeploying(true);

    setError(null);

    setDeployed(false);

    setLogs([]);

    const envObject = envVars.reduce(
      (acc, env) => {
        if (env.key.trim()) {
          acc[env.key] = env.value;
        }

        return acc;
      },
      {} as Record<string, string>,
    );

    const result = await deployApp({
      repoUrl,
      appName,
      namespace: "default",
      replicas: 1,
      port: detected.port,
      stack: detected.stack,
      envVars: envObject,
    });

    
    setSteps(result.steps || []); /* VRAIES étapes backend */
    setLogs(result.logs || []); /* VRAIS logs backend */
    setLiveUrl(result.url || null); /* URL backend */
    setDeployed(true);

  } catch (err) {
    setError(
      err instanceof Error
        ? err.message
        : "Deploy failed",
    );
  } finally {
    setDeploying(false);
  }
};

  const addEnvVar = () => {
    setEnvVars((prev) => [...prev, { key: "", value: "" }]);
  };

  const removeEnvVar = (index: number) => {
    setEnvVars((prev) => prev.filter((_, i) => i !== index));
  };

  const updateEnvVar = (
    index: number,
    field: keyof EnvVar,
    value: string,
  ) => {
    setEnvVars((prev) =>
      prev.map((ev, i) =>
        i === index
          ? {
              ...ev,
              [field]: value,
            }
          : ev,
      ),
    );
  };

  return (
    <div className="bg-surface border border-border rounded-2xl p-6">
      <div className="flex items-center gap-2 mb-5">
        <div className="w-8 h-8 rounded-lg bg-accent/10 flex items-center justify-center text-accent">
          <Rocket size={16} />
        </div>

        <div>
          <p className="text-[13px] font-medium text-k-text">
            Deploy from GitHub
          </p>

          <p className="text-[11px] text-k-muted font-mono">
            GitHub repo → Kubernetes application
          </p>
        </div>
      </div>

      <RepoInput
        repoUrl={repoUrl}
        setRepoUrl={setRepoUrl}
        detecting={detecting}
        onDetect={handleDetect}
      />

      {error && (
        <div className="rounded-xl p-3 bg-[#450a0a] border border-[#7f1d1d] mb-5">
          <span className="text-[#f87171] text-[12px] font-mono">
            ⚠ {error}
          </span>
        </div>
      )}

      {detected && (
        <div className="space-y-5">
          <StackDetectedCard detected={detected} />

          <div>
            <label className="block text-[11px] font-mono text-k-muted mb-1.5">
              Application name
            </label>

            <input
              type="text"
              value={appName}
              onChange={(e) => setAppName(e.target.value)}
              className="w-full bg-surface2 border border-border rounded-xl px-3 py-3 text-[12px] font-mono text-k-text"
            />
          </div>

          <GeneratedResources />

          <ExposureCard liveUrl={liveUrl} />

          <EnvVarsEditor
            envVars={envVars}
            addEnvVar={addEnvVar}
            removeEnvVar={removeEnvVar}
            updateEnvVar={updateEnvVar}
          />

          {(deploying || deployed) && (
            <DeployPipeline steps={steps} />
          )}

          {deployed && liveUrl && (
            <SuccessCard
              liveUrl={liveUrl}
              logs={logs}
              showLogs={showLogs}
              setShowLogs={setShowLogs}
            />
          )}

          <button
            onClick={handleDeploy}
            disabled={deploying || deployed}
            className="w-full flex items-center justify-center gap-2 px-5 py-3 rounded-xl text-[13px] bg-accent text-white"
          >
            {deploying ? (
              <>
                <Loader2 size={15} className="animate-spin" />
                Deploying application...
              </>
            ) : deployed ? (
              <>
                <CheckCircle2 size={15} />
                Application deployed
              </>
            ) : (
              <>
                <Rocket size={15} />
                Deploy to Kubernetes
              </>
            )}
          </button>
        </div>
      )}
    </div>
  );
}
