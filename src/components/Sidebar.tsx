"use client";

import { useState } from "react";
import {
  LayoutDashboard, Rocket, AppWindow, Terminal,
  Activity, Key, Settings, Hexagon,
} from "lucide-react";
import { cn } from "@/lib/data";

const NAV_ITEMS = [
  { section: "Main", items: [
    { icon: LayoutDashboard, label: "Overview", id: "overview" },
    { icon: Rocket, label: "Deploy", id: "deploy", badge: "new" },
    { icon: AppWindow, label: "Applications", id: "apps", badge: "5" },
  ]},
  { section: "Observability", items: [
    { icon: Terminal, label: "Logs", id: "logs" },
    { icon: Activity, label: "Metrics", id: "metrics" },
  ]},
  { section: "Config", items: [
    { icon: Key, label: "Env Vars", id: "envvars" },
    { icon: Settings, label: "Settings", id: "settings" },
  ]},
];

interface SidebarProps {
  activeNav: string;
  onNavChange: (id: string) => void;
}

export function Sidebar({ activeNav, onNavChange }: SidebarProps) {
  return (
    <aside className="flex flex-col bg-surface border-r border-border h-screen sticky top-0">
      {/* Logo */}
      <div className="flex items-center gap-2.5 px-5 py-5 border-b border-border">
        <div className="w-8 h-8 rounded-lg flex items-center justify-center bg-gradient-to-br from-accent to-accent2 shrink-0">
          <Hexagon size={16} className="text-white" />
        </div>
        <span className="text-[15px] font-semibold tracking-tight">KubeLaunch</span>
        <span className="ml-auto text-[9px] bg-[#1e3a5f] text-[#60a5fa] px-1.5 py-0.5 rounded font-mono">
          v0.1
        </span>
      </div>

      {/* Nav */}
      <nav className="flex-1 overflow-y-auto px-2 py-3">
        {NAV_ITEMS.map((group) => (
          <div key={group.section} className="mb-4">
            <p className="text-[9px] tracking-[1.5px] text-k-muted uppercase font-mono px-3 mb-1">
              {group.section}
            </p>
            {group.items.map((item) => (
              <button
                key={item.id}
                onClick={() => onNavChange(item.id)}
                className={cn(
                  "w-full flex items-center gap-2.5 px-3 py-2 rounded-lg text-[13px] transition-all duration-150 mb-0.5 text-left",
                  activeNav === item.id
                    ? "bg-[#1e3a5f] text-[#60a5fa]"
                    : "text-k-muted hover:bg-surface2 hover:text-k-text"
                )}
              >
                <item.icon size={16} className="shrink-0" />
                <span>{item.label}</span>
                {item.badge && (
                  <span className="ml-auto text-[10px] font-mono bg-surface2 px-1.5 py-0.5 rounded-full">
                    {item.badge}
                  </span>
                )}
              </button>
            ))}
          </div>
        ))}
      </nav>

      {/* Cluster pill */}
      <div className="p-4 border-t border-border">
        <div className="bg-surface2 border border-border rounded-lg p-2.5 flex items-center gap-2.5">
          <div className="w-2 h-2 rounded-full bg-k-green animate-pulse2 shrink-0" />
          <div className="min-w-0">
            <p className="text-[12px] font-mono text-k-text truncate">kind-dev</p>
            <p className="text-[10px] text-k-muted mt-0.5">local · 3 nodes</p>
          </div>
        </div>
      </div>
    </aside>
  );
}
