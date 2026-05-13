import type { Application } from "@/types";
import { STACK_CONFIG } from "@/lib/data";

interface AppCardProps {
  app: Application;
  onClick?: () => void;
}

const STATUS_STYLES = {
  running: "bg-[#052e16] text-[#4ade80] border border-[#14532d]",
  pending: "bg-[#451a03] text-[#fbbf24] border border-[#78350f]",
  error: "bg-[#450a0a] text-[#f87171] border border-[#7f1d1d]",
  stopped: "bg-surface2 text-k-muted border border-border",
};

const STATUS_DOT = {
  running: "bg-[#4ade80]",
  pending: "bg-[#fbbf24]",
  error: "bg-[#f87171]",
  stopped: "bg-k-muted",
};

export function AppCard({ app, onClick }: AppCardProps) {
  const stackCfg = STACK_CONFIG[app.stack];

  return (
    <div
      onClick={onClick}
      className="bg-surface border border-border rounded-xl p-4 cursor-pointer hover:border-slate-600 transition-colors duration-150 animate-fadeIn"
    >
      {/* Header */}
      <div className="flex items-start justify-between mb-3">
        <div>
          <p className="text-[13px] font-medium text-k-text">{app.name}</p>
          <p className="text-[10px] font-mono text-k-muted mt-0.5">
            {app.namespace} · {app.readyReplicas}/{app.replicas} pods
          </p>
        </div>
        <span className={`text-[10px] font-mono px-2 py-1 rounded-md flex items-center gap-1.5 ${STATUS_STYLES[app.status]}`}>
          <span className={`w-1.5 h-1.5 rounded-full ${STATUS_DOT[app.status]}`} />
          {app.status}
        </span>
      </div>

      {/* Tags */}
      <div className="flex items-center gap-1.5 mb-3 flex-wrap">
        <span
          className="text-[10px] font-mono px-2 py-0.5 rounded"
          style={{ background: stackCfg.bg, color: stackCfg.color }}
        >
          {stackCfg.label}
        </span>
        {app.framework && (
          <span className="text-[10px] font-mono px-2 py-0.5 rounded bg-[#2e1065] text-[#a78bfa]">
            {app.framework}
          </span>
        )}
        <span className="text-[10px] font-mono px-2 py-0.5 rounded bg-[#083344] text-[#22d3ee]">
          :{app.port}
        </span>
      </div>

      {/* Metrics */}
      <div className="grid grid-cols-3 gap-1.5">
        <div className="bg-surface2 rounded-md p-2">
          <p className="text-[12px] font-mono font-medium" style={{
            color: app.status === "running" ? "#4ade80" : app.status === "pending" ? "#fbbf24" : "#94a3b8"
          }}>
            {app.metrics.uptime}
          </p>
          <p className="text-[9px] text-k-muted mt-0.5">uptime</p>
        </div>
        <div className="bg-surface2 rounded-md p-2">
          <p className="text-[12px] font-mono font-medium text-k-text">{app.metrics.cpu}</p>
          <p className="text-[9px] text-k-muted mt-0.5">CPU</p>
        </div>
        <div className="bg-surface2 rounded-md p-2">
          <p className="text-[12px] font-mono font-medium text-k-text">{app.metrics.memory}</p>
          <p className="text-[9px] text-k-muted mt-0.5">Memory</p>
        </div>
      </div>
    </div>
  );
}
