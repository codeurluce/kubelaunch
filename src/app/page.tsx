"use client";

import { useState } from "react";
import { Sidebar } from "@/components/Sidebar";
import { AppCard } from "@/components/AppCard";
import { DeployForm } from "@/components/DeployForm";
import { LogViewer } from "@/components/LogViewer";
import { StatCard } from "@/components/StatCard";
import { MOCK_APPS } from "@/lib/data";

const PAGE_TITLES: Record<string, string> = {
  overview: "Overview",
  deploy: "Deploy",
  apps: "Applications",
  logs: "Logs",
  // metrics: "Metrics",
  // envvars: "Env Vars",
  settings: "Settings",
};

export default function Home() {
  const [activeNav, setActiveNav] = useState("deploy");
  const [selectedApp, setSelectedApp] = useState<string | null>(null);

  const running = MOCK_APPS.filter((a) => a.status === "running").length;
  const pending = MOCK_APPS.filter((a) => a.status === "pending").length;
  const namespaces = [...new Set(MOCK_APPS.map((a) => a.namespace))].length;

  return (
    <div className="grid grid-cols-[220px_1fr] min-h-screen">
      <Sidebar activeNav={activeNav} onNavChange={setActiveNav} />

      <div className="flex flex-col min-h-screen overflow-auto">
        {/* Topbar */}
        <header className="bg-surface border-b border-border px-6 h-14 flex items-center gap-4 sticky top-0 z-10">
          <h1 className="text-[15px] font-medium">{PAGE_TITLES[activeNav]}</h1>
          <div className="w-px h-5 bg-border" />
          <span className="text-[12px] font-mono text-k-muted">cluster / kind-dev</span>
          <button
            onClick={() => setActiveNav("deploy")}
            className="ml-auto flex items-center gap-2 px-4 py-2 rounded-lg text-[13px] bg-accent text-white hover:bg-blue-600 transition-colors font-medium"
          >
            + Deploy app
          </button>
        </header>

        {/* Content */}
        <main className="flex-1 p-6">
          {/* OVERVIEW */}
          {activeNav === "overview" && (
            <div className="animate-fadeIn">
              {/* Stats */}
              <div className="grid grid-cols-4 gap-3 mb-6">
                <StatCard value={MOCK_APPS.length} label="Applications" change="↑ 2 this week" changeType="up" />
                <StatCard value={running} label="Running" change="↑ healthy" changeType="up" valueColor="text-k-green" />
                <StatCard value={pending} label="Pending" change="deploying..." changeType="neutral" valueColor="text-k-amber" />
                <StatCard value={namespaces} label="Namespaces" change="default · prod · staging" />
              </div>

              {/* Apps grid */}
              <div className="mb-2 flex items-center gap-2">
                <span className="text-[13px] font-medium text-k-text">Running applications</span>
              </div>
              <div className="grid grid-cols-2 gap-3 mb-6">
                {MOCK_APPS.map((app) => (
                  <AppCard
                    key={app.id}
                    app={app}
                    onClick={() => {
                      setSelectedApp(app.name);
                      setActiveNav("logs");
                    }}
                  />
                ))}
              </div>

              {/* Logs */}
              <div className="mb-2 flex items-center gap-2">
                <span className="text-[13px] font-medium text-k-text">Live logs — api-service</span>
              </div>
              <LogViewer appName="api-service" />
            </div>
          )}

          {/* DEPLOY */}
          {activeNav === "deploy" && (
            <div className="max-w-2xl animate-fadeIn">
              <p className="text-[13px] text-k-muted mb-4">
                Paste a GitHub URL and KubeLaunch will detect your stack, generate all Kubernetes manifests automatically, and deploy in one click.
              </p>
              <DeployForm />
            </div>
          )}

          {/* APPS */}
          {activeNav === "apps" && (
            <div className="animate-fadeIn">
              <div className="grid grid-cols-2 gap-3">
                {MOCK_APPS.map((app) => (
                  <AppCard
                    key={app.id}
                    app={app}
                    onClick={() => {
                      setSelectedApp(app.name);
                      setActiveNav("logs");
                    }}
                  />
                ))}
              </div>
            </div>
          )}

          {/* LOGS */}
          {activeNav === "logs" && (
            <div className="animate-fadeIn">
              <div className="flex items-center gap-3 mb-4">
                <span className="text-[13px] text-k-muted">App:</span>
                <select
                  value={selectedApp || "api-service"}
                  onChange={(e) => setSelectedApp(e.target.value)}
                  className="bg-surface2 border border-border rounded-lg px-3 py-1.5 text-[12px] font-mono text-k-text focus:outline-none focus:border-accent"
                >
                  {MOCK_APPS.map((a) => (
                    <option key={a.id} value={a.name}>{a.name}</option>
                  ))}
                </select>
              </div>
              <LogViewer appName={selectedApp || "api-service"} />
            </div>
          )}

          {/* METRICS placeholder */}
          {activeNav === "metrics" && (
            <div className="animate-fadeIn">
              <div className="grid grid-cols-4 gap-3 mb-6">
                <StatCard value="312m" label="Total CPU" change="across 5 apps" />
                <StatCard value="960Mi" label="Total Memory" change="of 2Gi limit" />
                <StatCard value="9" label="Total Pods" change="8 ready" changeType="up" valueColor="text-k-green" />
                <StatCard value="99.2%" label="Avg Uptime" change="last 7 days" changeType="up" />
              </div>
              <div className="bg-surface border border-border rounded-xl p-6 flex items-center justify-center h-48">
                <p className="text-k-muted text-[13px] font-mono">
                  📊 Charts coming in v0.2 — connect Prometheus to enable
                </p>
              </div>
            </div>
          )}

          {/* ENV VARS placeholder */}
          {activeNav === "envvars" && (
            <div className="max-w-2xl animate-fadeIn">
              <p className="text-[13px] text-k-muted mb-4">Manage Kubernetes Secrets for your applications.</p>
              <div className="bg-surface border border-border rounded-xl p-6 text-center">
                <p className="text-k-muted text-[13px] font-mono">
                  🔑 Select an application to manage its env vars
                </p>
                <div className="flex flex-col gap-2 mt-4">
                  {MOCK_APPS.map((a) => (
                    <button
                      key={a.id}
                      className="flex items-center justify-between px-4 py-2.5 bg-surface2 border border-border rounded-lg text-[13px] hover:border-accent transition-colors"
                    >
                      <span className="font-mono">{a.name}</span>
                      <span className="text-[10px] text-k-muted font-mono">{a.namespace}</span>
                    </button>
                  ))}
                </div>
              </div>
            </div>
          )}

          {/* SETTINGS placeholder */}
          {activeNav === "settings" && (
            <div className="max-w-lg animate-fadeIn">
              <div className="bg-surface border border-border rounded-xl p-6">
                <p className="text-[13px] font-medium mb-4">Cluster connection</p>
                <div className="space-y-3">
                  <div>
                    <label className="block text-[11px] font-mono text-k-muted mb-1.5">Kubeconfig path</label>
                    <input
                      type="text"
                      defaultValue="~/.kube/config"
                      className="w-full bg-surface2 border border-border rounded-lg px-3 py-2 text-[12px] font-mono text-k-text focus:outline-none focus:border-accent"
                    />
                  </div>
                  <div>
                    <label className="block text-[11px] font-mono text-k-muted mb-1.5">Active context</label>
                    <input
                      type="text"
                      defaultValue="kind-dev"
                      className="w-full bg-surface2 border border-border rounded-lg px-3 py-2 text-[12px] font-mono text-k-text focus:outline-none focus:border-accent"
                    />
                  </div>
                  <button className="flex items-center gap-2 px-4 py-2 rounded-lg text-[13px] bg-accent text-white hover:bg-blue-600 transition-colors mt-2">
                    Save settings
                  </button>
                </div>
              </div>
            </div>
          )}
        </main>
      </div>
    </div>
  );
}
