interface StatCardProps {
  value: string | number;
  label: string;
  change?: string;
  changeType?: "up" | "down" | "neutral";
  valueColor?: string;
}

export function StatCard({ value, label, change, changeType = "neutral", valueColor }: StatCardProps) {
  const changeColor =
    changeType === "up" ? "text-k-green" :
    changeType === "down" ? "text-k-red" :
    "text-k-muted";

  return (
    <div className="bg-surface border border-border rounded-xl p-4">
      <p className={`text-[22px] font-semibold font-mono mb-0.5 ${valueColor || "text-k-text"}`}>
        {value}
      </p>
      <p className="text-[11px] text-k-muted">{label}</p>
      {change && (
        <p className={`text-[10px] font-mono mt-1.5 ${changeColor}`}>{change}</p>
      )}
    </div>
  );
}
