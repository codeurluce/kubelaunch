"use client";

import { CheckCircle2 } from "lucide-react";

type Props = {
  liveUrl: string;
  showLogs: boolean;
  setShowLogs: (value: boolean) => void;
};

export default function SuccessCard({
  liveUrl,
  showLogs,
  setShowLogs,
}: Props) {
  return (
    <div className="bg-[#052e16] border border-[#166534] rounded-xl p-4 animate-fadeIn">
      <div className="flex items-start gap-3">
        <CheckCircle2 size={18} className="text-k-green mt-0.5" />

        <div className="flex-1 min-w-0">
          <p className="text-[13px] font-semibold text-k-green mb-1">
            Deployment successful
          </p>

          <div className="bg-[#022c22] border border-[#14532d] rounded-lg px-3 py-2 text-[11px] font-mono text-[#86efac] break-all mb-3">
            🌐 {liveUrl}
          </div>

          <div className="flex items-center gap-2 flex-wrap">
            <button
              onClick={() => window.open(liveUrl, "_blank")}
              className="px-3 py-2 rounded-lg bg-k-green text-black text-[11px] font-medium"
            >
              Open application
            </button>

            <button
              onClick={() => setShowLogs(!showLogs)}
              className="px-3 py-2 rounded-lg bg-surface border border-border text-[11px] text-k-text"
            >
              View logs
            </button>
          </div>

          {showLogs && (
            <div className="mt-4 bg-black border border-border rounded-lg p-3 font-mono text-[11px] text-green-400 space-y-1">
              <p>✓ Repository cloned</p>
              <p>✓ Docker image built</p>
              <p>✓ Image loaded into Kind</p>
              <p>✓ Kubernetes deployment created</p>
              <p>✓ Service exposed</p>
              <p>✓ Pod ready</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
