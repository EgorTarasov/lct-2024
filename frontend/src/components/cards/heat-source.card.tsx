import { HeatSource } from "@/types/heat.type";
import { FC } from "react";
import { Text } from "../typography/Text";
import { IssueCard, IssueIcon } from "./issue.card";
import { Link } from "@tanstack/react-router";
import { PriorityCard, PriorityIcon } from "./priority-icon";

function pluralizeConsumer(count: number): string {
  const singular = "Потребитель";
  const genitiveSingular = "Потребителя";
  const genitivePlural = "Потребителей";

  if (count % 10 === 1 && count % 100 !== 11) {
    return singular;
  } else if ([2, 3, 4].includes(count % 10) && ![12, 13, 14].includes(count % 100)) {
    return genitiveSingular;
  } else {
    return genitivePlural;
  }
}

export const HeatSourceCard: FC<{ data: HeatSource.Item }> = (x) => {
  return (
    <Link
      to="/heat_source/$heatSourceId"
      params={{
        heatSourceId: x.data.id.toString()
      }}
      className="flex flex-col px-4 py-2 hover:bg-muted/50 w-full">
      <div className="flex items-center justify-between w-full">
        <Text.SubtleMedium className="text-muted-foreground">{x.data.number}</Text.SubtleMedium>
        <PriorityIcon data={x.data.priority} />
      </div>
      <Text.Large className="pb-1">{x.data.address}</Text.Large>
      <IssueCard data={x.data.issue} />
      <div className="flex items-center gap-1 pt-2">
        <p className="text-sm text-muted-foreground">
          {x.data.consumerCount} {pluralizeConsumer(x.data.consumerCount)} тепла
        </p>
        {x.data.issues.map((v) => (
          <IssueIcon className="*:size-4" data={v} />
        ))}
      </div>
    </Link>
  );
};

export const HeatSourceCardReadonly: FC<{ data: HeatSource.Item }> = (x) => (
  <div className="flex flex-col px-4 py-2 w-full">
    <div className="flex items-center justify-between w-full">
      <Text.SubtleMedium className="text-muted-foreground">{x.data.number}</Text.SubtleMedium>
    </div>
    <Text.Large className="pb-1">{x.data.address}</Text.Large>
    <div className="flex gap-2 items-center">
      <IssueCard data={x.data.issue} />
      <PriorityCard data={x.data.priority} />
    </div>
  </div>
);
