"use client";

import { Link2, ScanLine } from "lucide-react";

type Props = {
  repoUrl: string;
  setRepoUrl: (value: string) => void;
  detecting: boolean;
  onDetect: () => void;
};

export default function RepoInput({
  repoUrl,
  setRepoUrl,
  detecting,
  onDetect,
}: Props) {
  return (
    <div className="flex gap-2.5 mb-5">
      <div className="relative flex-1">
        <Link2
          size={15}
          className="absolute left-3 top-1/2 -translate-y-1/2 text-k-muted"
        />

        <input
          type="text"
          value={repoUrl}
          onChange={(e) => setRepoUrl(e.target.value)}
          onKeyDown={(e) => e.key === "Enter" && onDetect()}
          placeholder="https://github.com/username/project"
          className="w-full bg-surface2 border border-border rounded-xl pl-9 pr-4 py-3 text-[12px] font-mono text-k-text placeholder:text-k-muted focus:outline-none focus:border-accent transition-colors"
        />
      </div>

      <button
        onClick={onDetect}
        disabled={detecting || !repoUrl.trim()}
        className="flex items-center gap-2 px-4 py-3 rounded-xl text-[13px] bg-surface2 border border-border text-k-muted hover:bg-[#1e2d47] hover:text-k-text transition-all disabled:opacity-50"
      >
        <ScanLine size={15} className={detecting ? "animate-spin" : ""} />

        {detecting ? "Detecting..." : "Detect stack"}
      </button>
    </div>
  );
}
