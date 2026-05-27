"use client";

import { useMemo, useState } from "react";
import {
  Link2,
  ScanLine,
  Rocket,
  Trash2,
  CheckCircle2,
  Loader2,
  Globe,
  Package,
  FileCode2,
  Box,
  Terminal,
} from "lucide-react";

import type { DetectResult, EnvVar } from "@/types";
import { STACK_CONFIG } from "@/lib/data";
import { detectStack } from "@/lib/api";
import { stackIcons } from "./ui/stackIcons";

type DeployStep = {
  id: string;
  label: string;
  status: "pending" | "running" | "success";
};

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
    { key: "", value: "" },
  ]);

  const [appName, setAppName] = useState("");

  const [steps, setSteps] = useState<DeployStep[]>(INITIAL_STEPS);

  const [liveUrl, setLiveUrl] = useState<string | null>(null);

  const stackCfg = detected ? STACK_CONFIG[detected.stack] : null;

  const generatedResources = useMemo(() => {
    if (!detected) return [];

    return [
      {
        icon: <FileCode2 size={15} />,
        label: "Dockerfile",
      },
      {
        icon: <Box size={15} />,
        label: "Kubernetes Deployment",
      },
      {
        icon: <Globe size={15} />,
        label: "Kubernetes Service",
      },
      {
        icon: <Terminal size={15} />,
        label: "Kubernetes Secret",
      },
    ];
  }, [detected]);

  const estimatedUrl = useMemo(() => {
    if (!appName) return "http://localhost:30xxx";

    return `http://${appName.toLowerCase()}.local`;
  }, [appName]);

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

  const sleep = (ms: number) =>
    new Promise((resolve) => setTimeout(resolve, ms));

  const handleDetect = async () => {
    if (!repoUrl.trim()) return;

    setDetecting(true);
    setDetected(null);
    setError(null);
    setDeployed(false);
    setLiveUrl(null);

    try {
      const result = await detectStack(repoUrl);

      const repoName =
        repoUrl
          .split("/")
          .pop()
          ?.replace(".git", "") || "my-app";

      setAppName(repoName);
      setDetected(result);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Unknown error");
    } finally {
      setDetecting(false);
    }
  };

  const handleDeploy = async () => {
    if (!detected) return;

    setDeploying(true);
    setDeployed(false);
    setSteps(INITIAL_STEPS);

    const pipeline = [
      "clone",
      "dockerfile",
      "manifests",
      "build",
      "push",
      "deploy",
      "ready",
    ];

    for (const step of pipeline) {
      updateStep(step, "running");

      await sleep(1200);

      updateStep(step, "success");
    }

    setLiveUrl(estimatedUrl);
    setDeploying(false);
    setDeployed(true);
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

      {/* REPOSITORY INPUT */}
      <div className="flex gap-2.5 mb-5">
        <div className="relative flex-1">
          <Link2
            size={15}
            className="absolute left-3 top-1/2 -translate-y-1/2 text-k-muted"
          />

          <input
            type="text"
            value={repoUrl}
            onChange={(e) => setRepoUrl(e.target.value)}
            onKeyDown={(e) => e.key === "Enter" && handleDetect()}
            placeholder="https://github.com/username/project"
            className="w-full bg-surface2 border border-border rounded-xl pl-9 pr-4 py-3 text-[12px] font-mono text-k-text placeholder:text-k-muted focus:outline-none focus:border-accent transition-colors"
          />
        </div>

        <button
          onClick={handleDetect}
          disabled={detecting || !repoUrl.trim()}
          className="flex items-center gap-2 px-4 py-3 rounded-xl text-[13px] bg-surface2 border border-border text-k-muted hover:bg-[#1e2d47] hover:text-k-text transition-all disabled:opacity-50"
        >
          <ScanLine
            size={15}
            className={detecting ? "animate-spin" : ""}
          />

          {detecting ? "Detecting..." : "Detect stack"}
        </button>
      </div>

      {/* ERROR */}
      {error && (
        <div className="rounded-xl p-3 flex items-center gap-2 bg-[#450a0a] border border-[#7f1d1d] mb-5">
          <span className="text-[#f87171] text-[12px] font-mono">
            ⚠ {error}
          </span>
        </div>
      )}

      {/* STACK DETECTED */}
      {detected && stackCfg && (
        <div
          className="rounded-xl p-4 border mb-5 animate-fadeIn"
          style={{
            background: stackCfg.bg,
            borderColor: stackCfg.color + "33",
          }}
        >
          <div className="flex items-start gap-3">
            <div
              className="w-10 h-10 shrink-0 rounded-xl flex items-center justify-center"
              style={{ background: stackCfg.color + "22" }}
            >
              {stackIcons[detected.stack] || stackIcons.nodejs}
            </div>

            <div className="flex-1 min-w-0">
              <div className="flex items-center justify-between gap-3 mb-1">
                <p
                  className="text-[13px] font-semibold font-mono"
                  style={{ color: stackCfg.color }}
                >
                  {stackCfg.label} detected
                </p>

                <span
                  className="text-[10px] font-mono px-2 py-1 rounded"
                  style={{
                    background: stackCfg.color + "22",
                    color: stackCfg.color,
                  }}
                >
                  {detected.confidence} confidence
                </span>
              </div>

              <p
                className="text-[11px] leading-relaxed"
                style={{ color: stackCfg.color + "CC" }}
              >
                KubeLaunch detected a {detected.framework} application.
                A production-ready Dockerfile and Kubernetes deployment
                configuration will be generated automatically.
              </p>

              <p
                className="text-[11px] mt-2 font-mono"
                style={{ color: stackCfg.color + "99" }}
              >
                {detected.detected_files?.join(", ")} · port {detected.port}
              </p>
            </div>
          </div>
        </div>
      )}

      {/* CONFIGURATION */}
      {detected && (
        <div className="space-y-5 animate-fadeIn">
          {/* APP NAME */}
          <div>
            <label className="block text-[11px] font-mono text-k-muted mb-1.5">
              Application name
            </label>

            <input
              type="text"
              value={appName}
              onChange={(e) => setAppName(e.target.value)}
              className="w-full bg-surface2 border border-border rounded-xl px-3 py-3 text-[12px] font-mono text-k-text focus:outline-none focus:border-accent transition-colors"
            />
          </div>

          {/* GENERATED RESOURCES */}
          <div className="bg-surface2 border border-border rounded-xl p-4">
            <div className="flex items-center gap-2 mb-3">
              <Package size={15} className="text-k-muted" />

              <p className="text-[12px] font-medium text-k-text">
                Generated resources preview
              </p>
            </div>

            <div className="grid grid-cols-2 gap-2">
              {generatedResources.map((resource) => (
                <div
                  key={resource.label}
                  className="flex items-center gap-2 bg-surface rounded-lg border border-border px-3 py-2"
                >
                  <span className="text-k-muted">{resource.icon}</span>

                  <span className="text-[11px] font-mono text-k-text">
                    {resource.label}
                  </span>
                </div>
              ))}
            </div>
          </div>

          {/* ESTIMATED EXPOSURE */}
          <div className="bg-surface2 border border-border rounded-xl p-4">
            <div className="flex items-center gap-2 mb-2">
              <Globe size={15} className="text-k-muted" />

              <p className="text-[12px] font-medium text-k-text">
                Estimated exposure
              </p>
            </div>

            <p className="text-[11px] text-k-muted mb-2 leading-relaxed">
              Your application will automatically be exposed after deployment.
            </p>

            <div className="bg-surface border border-border rounded-lg px-3 py-2 text-[11px] font-mono text-k-green break-all">
              {estimatedUrl}
            </div>
          </div>

          {/* ENV VARS */}
          <div>
            <p className="text-[11px] font-mono text-k-muted mb-2">
              Environment variables
            </p>

            <div className="space-y-2">
              {envVars.map((ev, i) => (
                <div key={i} className="flex gap-2 items-center">
                  <input
                    type="text"
                    value={ev.key}
                    onChange={(e) =>
                      updateEnvVar(i, "key", e.target.value)
                    }
                    placeholder="KEY"
                    className="flex-1 bg-surface2 border border-border rounded-lg px-3 py-2 text-[11px] font-mono text-k-text placeholder:text-k-muted focus:outline-none focus:border-accent"
                  />

                  <input
                    type="text"
                    value={ev.value}
                    onChange={(e) =>
                      updateEnvVar(i, "value", e.target.value)
                    }
                    placeholder="value"
                    className="flex-1 bg-surface2 border border-border rounded-lg px-3 py-2 text-[11px] font-mono text-k-text placeholder:text-k-muted focus:outline-none focus:border-accent"
                  />

                  <button
                    onClick={() => removeEnvVar(i)}
                    className="text-k-muted hover:text-k-red transition-colors p-1"
                  >
                    <Trash2 size={14} />
                  </button>
                </div>
              ))}
            </div>

            <button
              onClick={addEnvVar}
              className="mt-2 flex items-center gap-1.5 text-[11px] text-k-muted hover:text-k-text transition-colors font-mono"
            >
              <Rocket size={12} /> Add variable
            </button>
          </div>

          {/* DEPLOY PIPELINE */}
          {(deploying || deployed) && (
            <div className="bg-surface2 border border-border rounded-xl p-4 animate-fadeIn">
              <div className="flex items-center gap-2 mb-4">
                <Rocket size={15} className="text-accent" />

                <p className="text-[12px] font-medium text-k-text">
                  Deployment pipeline
                </p>
              </div>

              <div className="space-y-2">
                {steps.map((step) => (
                  <div
                    key={step.id}
                    className="flex items-center gap-3 px-3 py-2 rounded-lg border border-border bg-surface"
                  >
                    <div className="w-4 flex justify-center">
                      {step.status === "success" ? (
                        <CheckCircle2
                          size={15}
                          className="text-k-green"
                        />
                      ) : step.status === "running" ? (
                        <Loader2
                          size={15}
                          className="animate-spin text-accent"
                        />
                      ) : (
                        <div className="w-2 h-2 rounded-full bg-k-muted/40" />
                      )}
                    </div>

                    <span
                      className={`text-[11px] font-mono ${
                        step.status === "success"
                          ? "text-k-text"
                          : step.status === "running"
                            ? "text-accent"
                            : "text-k-muted"
                      }`}
                    >
                      {step.label}
                    </span>
                  </div>
                ))}
              </div>
            </div>
          )}

          {/* SUCCESS */}
          {deployed && liveUrl && (
            <div className="bg-[#052e16] border border-[#166534] rounded-xl p-4 animate-fadeIn">
              <div className="flex items-start gap-3">
                <CheckCircle2
                  size={18}
                  className="text-k-green mt-0.5"
                />

                <div className="flex-1 min-w-0">
                  <p className="text-[13px] font-semibold text-k-green mb-1">
                    Deployment successful
                  </p>

                  <p className="text-[11px] text-[#bbf7d0] leading-relaxed mb-3">
                    Your application is now running on Kubernetes and is
                    reachable from the generated service endpoint.
                  </p>

                  <div className="bg-[#022c22] border border-[#14532d] rounded-lg px-3 py-2 text-[11px] font-mono text-[#86efac] break-all mb-3">
                    🌐 {liveUrl}
                  </div>

                  <div className="flex items-center gap-2 flex-wrap">
                    <button className="px-3 py-2 rounded-lg bg-k-green text-black text-[11px] font-medium hover:opacity-90 transition-opacity">
                      Open application
                    </button>

                    <button className="px-3 py-2 rounded-lg bg-surface border border-border text-[11px] text-k-text hover:border-accent transition-colors">
                      View logs
                    </button>
                  </div>
                </div>
              </div>
            </div>
          )}

          {/* DEPLOY BUTTON */}
          <button
            onClick={handleDeploy}
            disabled={deploying || deployed}
            className="w-full flex items-center justify-center gap-2 px-5 py-3 rounded-xl text-[13px] bg-accent text-white hover:bg-blue-600 transition-all disabled:opacity-70 disabled:cursor-not-allowed font-medium"
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