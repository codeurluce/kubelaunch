"use client";

import {
  CheckCircle2,
  Loader2,
  Rocket,
  XCircle,
} from "lucide-react";

type Step = {
  id: string;
  label: string;
  status: "pending" | "running" | "success" | "error";
  message?: string;
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
            className="flex items-start gap-3 px-3 py-2 rounded-lg border border-border bg-surface"
          >
            <div className="w-4 flex justify-center mt-[2px]">
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
              ) : step.status === "error" ? (
                <XCircle
                  size={15}
                  className="text-red-500"
                />
              ) : (
                <div className="w-2 h-2 rounded-full bg-k-muted/40" />
              )}
            </div>

            <div className="flex-1 min-w-0">
              <p
                className={`text-[11px] font-mono ${
                  step.status === "error"
                    ? "text-red-400"
                    : step.status === "running"
                      ? "text-accent"
                      : "text-k-text"
                }`}
              >
                {step.label}
              </p>

              {step.message && (
                <p className="text-[10px] text-k-muted mt-1 font-mono">
                  {step.message}
                </p>
              )}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
