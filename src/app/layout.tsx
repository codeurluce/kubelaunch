import type { Metadata } from "next";
// Ignore missing module/type declarations for CSS imports in this file
// (project may declare these types elsewhere or via a global d.ts)
// @ts-ignore
import "./globals.css";

export const metadata: Metadata = {
  title: "KubeLaunch — Deploy to Kubernetes without YAML",
  description: "Deploy your Node.js, React or Python app to Kubernetes in 3 minutes. No YAML. No PhD required.",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en" translate="no" suppressHydrationWarning>
      <body className="font-sans antialiased">{children}</body>
    </html>
  );
}
