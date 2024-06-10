import { Issue, IssueLocaleMap } from "@/types/issue.type";
import { FC } from "react";

export const IssueCard: FC<{ data: Issue }> = (x) => {
  const v = IssueLocaleMap[x.data];
  return (
    <div
      className="p-1 w-fit space-x-2 rounded-sm flex gap-1"
      style={{
        backgroundColor: v.backgroundColor,
        color: v.color
      }}>
      {v.icon}
      {v.locale}
    </div>
  );
};

export const IssueIcon = ({ data }: { data: Issue }) => {
  const v = IssueLocaleMap[data];
  return v.icon;
};
