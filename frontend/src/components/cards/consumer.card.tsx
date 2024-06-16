import { Consumer } from "@/types/consumer.type";
import { Link } from "@tanstack/react-router";
import { FC } from "react";
import { Text } from "../typography/Text";
import { PriorityCard, PriorityIcon } from "./priority-icon";
import { IssueCard } from "./issue.card";
import { TitleInfo } from "./title-info";
import { cn } from "@/utils/cn";
import { toJS } from "mobx";
import { observer } from "mobx-react-lite";

export const ConsumerCard: FC<{
  data: Consumer.Item;
  className?: string;
}> = (x) => {
  return (
    <div className={cn("flex flex-col px-4 py-3 hover:bg-muted/50 w-full", x.className)}>
      <div className="flex items-center justify-between w-full flex-wrap gap-3">
        <Text.SubtleMedium className="text-muted-foreground">{x.data.address}</Text.SubtleMedium>
        <PriorityIcon data={x.data.priority} />
      </div>
      <Text.Large className="pb-1">{x.data.name}</Text.Large>
      <TitleInfo title="Тип потребителя" info={x.data.info.type} className="pb-1" />
      <IssueCard data={x.data.issue} />
    </div>
  );
};

export const ConsumerCardReadonly: FC<{ data: Consumer.Item }> = observer((x) => {
  return (
    <div className="flex flex-col px-4 w-full pb-3">
      <Text.SubtleMedium className="text-muted-foreground">{x.data.address}</Text.SubtleMedium>
      <Text.Large className="pb-1">{x.data.name}</Text.Large>
      <TitleInfo title="Тип потребителя" info={x.data.info.type} />
      <div className="py-1.5">
        <IssueCard data={x.data.issue} />
      </div>
      <PriorityCard data={x.data.priority} />
    </div>
  );
});
