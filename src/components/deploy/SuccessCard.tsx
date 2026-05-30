"use client";

import { CheckCircle2 } from "lucide-react";

type Props = {
  liveUrl: string;

  logs: string[];

  showLogs: boolean;

  setShowLogs: (value: boolean) => void;
};

export default function SuccessCard({
  liveUrl,
  logs,
  showLogs,
  setShowLogs,
}: Props) {
  return (
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
            Your application is now running on Kubernetes.
          </p>

          <div className="bg-[#022c22] border border-[#14532d] rounded-lg px-3 py-2 text-[11px] font-mono text-[#86efac] break-all mb-3">
            🌐 {liveUrl}
          </div>

          <div className="flex items-center gap-2 flex-wrap">
            <button
              onClick={() => window.open(liveUrl, "_blank")}
              className="px-3 py-2 rounded-lg bg-k-green text-black text-[11px] font-medium hover:opacity-90 transition-opacity"
            >
              Open application
            </button>

            <button
              onClick={() => setShowLogs(!showLogs)}
              className="px-3 py-2 rounded-lg bg-surface border border-border text-[11px] text-k-text hover:border-accent transition-colors"
            >
              {showLogs ? "Hide logs" : "View logs"}
            </button>
          </div>

          {showLogs && (
            <div className="mt-4 bg-black border border-border rounded-lg p-3 font-mono text-[11px] text-green-400 space-y-1 max-h-64 overflow-auto">
              {logs.length > 0 ? (
                logs.map((log, index) => (
                  <p key={index}>
                    {log}
                  </p>
                ))
              ) : (
                <p className="text-k-muted">
                  No logs available
                </p>
              )}
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
