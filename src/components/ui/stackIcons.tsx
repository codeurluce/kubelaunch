import type { ReactNode } from "react";

import {
  SiNextdotjs,
  SiReact,
  SiVite,
  SiNodedotjs,
  SiDocker,
  SiKubernetes,
  SiGo,
  SiPython,
  SiRust,
  SiPhp,
  SiOpenjdk,
} from "react-icons/si";

import { Globe } from "lucide-react";

export const stackIcons: Record<string, ReactNode> = {
  nodejs: <SiNodedotjs className="w-4 h-4 text-green-500" />,
  nextjs: <SiNextdotjs className="w-4 h-4" />,
  react: <SiReact className="w-4 h-4 text-cyan-400" />,
  python: <SiPython className="w-4 h-4 text-yellow-400" />,
  go: <SiGo className="w-4 h-4 text-cyan-400" />,
  rust: <SiRust className="w-4 h-4 text-orange-500" />,
  php: <SiPhp className="w-4 h-4 text-indigo-400" />,
  java: <SiOpenjdk className="w-4 h-4 text-red-500" />,
  docker: <SiDocker className="w-4 h-4 text-blue-500" />,
  kubernetes: <SiKubernetes className="w-4 h-4 text-blue-600" />,
  static: <Globe className="w-4 h-4 text-blue-400" />,
  "react-vite": (
    <div className="flex items-center gap-0.5">
      <SiReact className="w-4 h-4 text-cyan-400" />
      <SiVite className="w-3 h-3 text-purple-500" />
    </div>
  ),
};
