import { HeatSource } from "@/types/heat.type";
import { FC } from "react";
import { Text } from "../typography/Text";
import { Priority } from "@/types/priority.type";
import { IssueCard, IssueIcon } from "./issue.card";

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
    <button className="flex flex-col px-4 py-2 hover:bg-muted w-full">
      <div className="flex items-center justify-between w-full">
        <Text.SubtleMedium className="text-muted-foreground">{x.data.number}</Text.SubtleMedium>
        {Priority.ItemMap[x.data.priority].icon}
      </div>
      <Text.Large className="pb-1">{x.data.address}</Text.Large>
      <IssueCard data={x.data.issue} />
      <div className="flex items-center gap-1 pt-2">
        <p>
          {x.data.consumerCount} {pluralizeConsumer(x.data.consumerCount)} тепла
        </p>
        {x.data.issues.map((v) => (
          <IssueIcon data={v} />
        ))}
      </div>
    </button>
  );
};
