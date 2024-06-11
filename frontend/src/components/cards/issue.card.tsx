import { Issue, IssueLocaleMap } from "@/types/issue.type";
import { FC } from "react";

export const IssueCard: FC<{ data: Issue }> = (x) => {
  const v = IssueLocaleMap[x.data];
  return (
    <div
      className="p-1 w-fit space-x-2 rounded-sm flex gap-1 text-sm items-center"
      style={{
        backgroundColor: v.backgroundColor,
        color: v.color
      }}>
      <span className="*:size-4">{v.icon}</span>
      {v.locale}
    </div>
  );
};

export const IssueIcon = ({ data, className }: { data: Issue; className?: string }) => {
  const v = IssueLocaleMap[data];
  return (
    <span className={className} style={{ color: v.color }}>
      {v.icon}
    </span>
  );
};
