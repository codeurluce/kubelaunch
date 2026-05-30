"use client";

import { Rocket, Trash2 } from "lucide-react";
import type { EnvVar } from "@/types";

type Props = {
  envVars: EnvVar[];
  addEnvVar: () => void;
  removeEnvVar: (index: number) => void;
  updateEnvVar: (
    index: number,
    field: keyof EnvVar,
    value: string,
  ) => void;
};

export default function EnvVarsEditor({
  envVars,
  addEnvVar,
  removeEnvVar,
  updateEnvVar,
}: Props) {
  return (
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
              onChange={(e) => updateEnvVar(i, "key", e.target.value)}
              placeholder="KEY"
              className="flex-1 bg-surface2 border border-border rounded-lg px-3 py-2 text-[11px] font-mono text-k-text"
            />

            <input
              type="text"
              value={ev.value}
              onChange={(e) => updateEnvVar(i, "value", e.target.value)}
              placeholder="value"
              className="flex-1 bg-surface2 border border-border rounded-lg px-3 py-2 text-[11px] font-mono text-k-text"
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
  );
}
