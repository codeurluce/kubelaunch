"use client";

import { useState } from "react";
import { Link2, ScanLine, Rocket, Plus, Trash2, CheckCircle } from "lucide-react";
import type { DetectResult, DeployFormData, EnvVar } from "@/types";
import { STACK_CONFIG, detectStackFromUrl } from "@/lib/data";

export function DeployForm() {
  const [repoUrl, setRepoUrl] = useState("");
  const [detecting, setDetecting] = useState(false);
  const [detected, setDetected] = useState<DetectResult | null>(null);
  const [deploying, setDeploying] = useState(false);
  const [deployed, setDeployed] = useState(false);
  const [envVars, setEnvVars] = useState<EnvVar[]>([{ key: "", value: "" }]);
  const [form, setForm] = useState({ appName: "", namespace: "default", replicas: 2 });

  const handleDetect = async () => {
    if (!repoUrl.trim()) return;
    setDetecting(true);
    setDetected(null);
    await new Promise((r) => setTimeout(r, 1200));
    const result = detectStackFromUrl(repoUrl);
    const repoName = repoUrl.split("/").pop() || "my-app";
    setForm((f) => ({ ...f, appName: repoName }));
    setDetected(result);
    setDetecting(false);
  };

  const handleDeploy = async () => {
    if (!detected) return;
    setDeploying(true);
    await new Promise((r) => setTimeout(r, 2000));
    setDeploying(false);
    setDeployed(true);
    setTimeout(() => setDeployed(false), 3000);
  };

  const addEnvVar = () => setEnvVars((v) => [...v, { key: "", value: "" }]);
  const removeEnvVar = (i: number) => setEnvVars((v) => v.filter((_, idx) => idx !== i));
  const updateEnvVar = (i: number, field: keyof EnvVar, val: string) => {
    setEnvVars((v) => v.map((ev, idx) => idx === i ? { ...ev, [field]: val } : ev));
  };

  const stackCfg = detected ? STACK_CONFIG[detected.stack] : null;

  return (
    <div className="bg-surface border border-border rounded-xl p-6">
      <p className="text-[13px] font-medium text-k-muted mb-4 flex items-center gap-2">
        <span className="text-base">⬡</span> Deploy from GitHub repo
      </p>

      {/* URL Input */}
      <div className="flex gap-2.5 mb-4">
        <div className="relative flex-1">
          <Link2 size={15} className="absolute left-3 top-1/2 -translate-y-1/2 text-k-muted pointer-events-none" />
          <input
            type="text"
            value={repoUrl}
            onChange={(e) => setRepoUrl(e.target.value)}
            onKeyDown={(e) => e.key === "Enter" && handleDetect()}
            placeholder="https://github.com/your-username/your-app"
            className="w-full bg-surface2 border border-border rounded-lg pl-9 pr-4 py-2.5 text-[12px] font-mono text-k-text placeholder:text-k-muted focus:outline-none focus:border-accent transition-colors"
          />
        </div>
        <button
          onClick={handleDetect}
          disabled={detecting || !repoUrl.trim()}
          className="flex items-center gap-2 px-4 py-2.5 rounded-lg text-[13px] bg-surface2 border border-border text-k-muted hover:bg-[#1e2d47] hover:text-k-text transition-all disabled:opacity-50 disabled:cursor-not-allowed"
        >
          <ScanLine size={15} className={detecting ? "animate-spin" : ""} />
          {detecting ? "Detecting..." : "Detect stack"}
        </button>
      </div>

      {/* Stack detected */}
      {detected && stackCfg && (
        <div
          className="rounded-lg p-3 flex items-center gap-3 mb-4 border animate-fadeIn"
          style={{ background: stackCfg.bg, borderColor: stackCfg.color + "33" }}
        >
          <div className="w-8 h-8 rounded-lg flex items-center justify-center text-sm" style={{ background: stackCfg.color + "22" }}>
            ⬡
          </div>
          <div className="flex-1 min-w-0">
            <p className="text-[12px] font-mono font-medium" style={{ color: stackCfg.color }}>
              {stackCfg.label} detected
            </p>
            <p className="text-[11px] mt-0.5" style={{ color: stackCfg.color + "99" }}>
              {detected.detected_files.join(", ")} · {detected.framework} · port {detected.port}
            </p>
          </div>
          <span
            className="text-[10px] font-mono px-2 py-1 rounded"
            style={{ background: stackCfg.color + "22", color: stackCfg.color }}
          >
            {detected.confidence} confidence
          </span>
        </div>
      )}

      {/* Form fields */}
      {detected && (
        <>
          <div className="grid grid-cols-3 gap-3 mb-4 animate-fadeIn">
            {[
              { label: "App name", key: "appName", val: form.appName },
              { label: "Namespace", key: "namespace", val: form.namespace },
              { label: "Replicas", key: "replicas", val: String(form.replicas) },
            ].map((f) => (
              <div key={f.key}>
                <label className="block text-[11px] font-mono text-k-muted mb-1.5">{f.label}</label>
                <input
                  type="text"
                  value={f.val}
                  onChange={(e) => setForm((prev) => ({ ...prev, [f.key]: e.target.value }))}
                  className="w-full bg-surface2 border border-border rounded-lg px-3 py-2 text-[12px] font-mono text-k-text focus:outline-none focus:border-accent transition-colors"
                />
              </div>
            ))}
          </div>

          {/* Env vars */}
          <div className="mb-4 animate-fadeIn">
            <p className="text-[11px] font-mono text-k-muted mb-2">Environment variables</p>
            <div className="space-y-2">
              {envVars.map((ev, i) => (
                <div key={i} className="flex gap-2 items-center">
                  <input
                    type="text"
                    value={ev.key}
                    onChange={(e) => updateEnvVar(i, "key", e.target.value)}
                    placeholder="KEY"
                    className="flex-1 bg-surface2 border border-border rounded-lg px-3 py-2 text-[11px] font-mono text-k-text placeholder:text-k-muted focus:outline-none focus:border-accent"
                  />
                  <input
                    type="text"
                    value={ev.value}
                    onChange={(e) => updateEnvVar(i, "value", e.target.value)}
                    placeholder="value"
                    className="flex-1 bg-surface2 border border-border rounded-lg px-3 py-2 text-[11px] font-mono text-k-text placeholder:text-k-muted focus:outline-none focus:border-accent"
                  />
                  <button onClick={() => removeEnvVar(i)} className="text-k-muted hover:text-k-red transition-colors p-1">
                    <Trash2 size={14} />
                  </button>
                </div>
              ))}
            </div>
            <button onClick={addEnvVar} className="mt-2 flex items-center gap-1.5 text-[11px] text-k-muted hover:text-k-text transition-colors font-mono">
              <Plus size={13} /> Add variable
            </button>
          </div>

          <button
            onClick={handleDeploy}
            disabled={deploying || deployed}
            className="flex items-center gap-2 px-5 py-2.5 rounded-lg text-[13px] bg-accent text-white hover:bg-blue-600 transition-all disabled:opacity-70 disabled:cursor-not-allowed font-medium"
          >
            {deployed ? (
              <><CheckCircle size={15} /> Deployed successfully!</>
            ) : deploying ? (
              <><Rocket size={15} className="animate-bounce" /> Deploying...</>
            ) : (
              <><Rocket size={15} /> Deploy to Kubernetes</>
            )}
          </button>
        </>
      )}
    </div>
  );
}
