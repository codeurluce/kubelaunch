"use client";

import { CheckCircle2, Loader2, Rocket } from "lucide-react";

type Step = {
  id: string;
  label: string;
  status: "pending" | "running" | "success";
};

type Props = {
  steps: Step[];
};

export default function DeployPipeline({ steps }: Props) {
  return (
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
                <CheckCircle2 size={15} className="text-k-green" />
              ) : step.status === "running" ? (
                <Loader2 size={15} className="animate-spin text-accent" />
              ) : (
                <div className="w-2 h-2 rounded-full bg-k-muted/40" />
              )}
            </div>

            <span className="text-[11px] font-mono text-k-text">
              {step.label}
            </span>
          </div>
        ))}
      </div>
    </div>
  );
}
