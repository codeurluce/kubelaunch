"use client";

import { useEffect, useRef, useState } from "react";
import { Terminal } from "lucide-react";
import type { LogEntry } from "@/types";
import { MOCK_LOGS } from "@/lib/data";

const LEVEL_STYLES: Record<string, string> = {
  INFO: "text-[#60a5fa]",
  OK: "text-[#4ade80]",
  WARN: "text-[#fbbf24]",
  ERROR: "text-[#f87171]",
  DEBUG: "text-[#94a3b8]",
};

interface LogViewerProps {
  appName?: string;
}

export function LogViewer({ appName = "api-service" }: LogViewerProps) {
  const [logs, setLogs] = useState<LogEntry[]>(MOCK_LOGS.slice(0, 5));
  const logsContainerRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    let idx = 5;

    const interval = setInterval(() => {
      if (idx < MOCK_LOGS.length) {
        setLogs((prev) => [...prev, MOCK_LOGS[idx]]);
        idx++;
      } else {
        clearInterval(interval);
      }
    }, 2000);

    return () => clearInterval(interval);
  }, []);

  useEffect(() => {
    if (logsContainerRef.current) {
      logsContainerRef.current.scrollTop =
        logsContainerRef.current.scrollHeight;
    }
  }, [logs]);

  return (
    <div className="bg-surface border border-border rounded-xl overflow-hidden">
      <div
        ref={logsContainerRef}
        className="bg-[#050a14] px-4 py-3.5 font-mono text-[11px] leading-loose max-h-64 overflow-y-auto"
      >
        <Terminal size={14} className="text-k-muted" />
        <span className="text-[12px] font-mono font-medium text-k-text">
          {appName}-7d9f8b-x4k2p
        </span>
        <div className="ml-auto flex items-center gap-1.5 text-[10px] text-k-green font-mono">
          <div className="w-1.5 h-1.5 rounded-full bg-k-green animate-pulse2" />
          live
        </div>
      </div>
      <div className="bg-[#050a14] px-4 py-3.5 font-mono text-[11px] leading-loose max-h-48 overflow-y-auto">
        {logs?.filter(Boolean).map((log, i) => (
          <div key={i} className="flex gap-3 animate-slideIn">
            <span className="text-[#334155] min-w-[80px] shrink-0">
              {log?.time || "--:--:--"}
            </span>
            <span
              className={`min-w-[40px] shrink-0 ${LEVEL_STYLES[log.level] || "text-k-muted"}`}
            >
              {log?.level || "INFO"}
            </span>
            <span className="text-[#94a3b8]">{log.message}</span>
          </div>
        ))}
        <div className="text-xs text-k-muted mt-2">
              Waiting for new logs...
        </div>
        {/* <div ref={bottomRef} /> */}
      </div>
    </div>
  );
}
