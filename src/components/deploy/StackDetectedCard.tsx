"use client";

import { stackIcons } from "../ui/stackIcons";
import { STACK_CONFIG } from "@/lib/data";
import type { DetectResult } from "@/types";

type Props = {
  detected: DetectResult;
};

export default function StackDetectedCard({ detected }: Props) {
  const stackCfg = STACK_CONFIG[detected.stack];

  if (!stackCfg) return null;

  return (
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
  );
}
