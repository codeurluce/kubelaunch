"use client";

import { Globe } from "lucide-react";

type Props = {
  liveUrl: string | null;
};

export default function ExposureCard({ liveUrl }: Props) {
  return (
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
        {liveUrl || "Waiting deployment..."}
      </div>
    </div>
  );
}
