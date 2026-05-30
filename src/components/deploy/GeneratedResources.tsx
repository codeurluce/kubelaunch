"use client";

import { FileCode2, Box, Globe, Terminal, Package } from "lucide-react";

export default function GeneratedResources() {
  const resources = [
    {
      icon: <FileCode2 size={15} />,
      label: "Dockerfile",
    },
    {
      icon: <Box size={15} />,
      label: "Kubernetes Deployment",
    },
    {
      icon: <Globe size={15} />,
      label: "Kubernetes Service",
    },
    {
      icon: <Terminal size={15} />,
      label: "Kubernetes Secret",
    },
  ];

  return (
    <div className="bg-surface2 border border-border rounded-xl p-4">
      <div className="flex items-center gap-2 mb-3">
        <Package size={15} className="text-k-muted" />

        <p className="text-[12px] font-medium text-k-text">
          Generated resources preview
        </p>
      </div>

      <div className="grid grid-cols-2 gap-2">
        {resources.map((resource) => (
          <div
            key={resource.label}
            className="flex items-center gap-2 bg-surface rounded-lg border border-border px-3 py-2"
          >
            <span className="text-k-muted">{resource.icon}</span>

            <span className="text-[11px] font-mono text-k-text">
              {resource.label}
            </span>
          </div>
        ))}
      </div>
    </div>
  );
}
