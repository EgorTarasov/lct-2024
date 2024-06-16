import { Incident } from "@/types/incident.type";
import { observer } from "mobx-react-lite";
import { FC } from "react";
import { IncidentCardWrapper } from "./components/incident-card-wrapper";
import { PriorityCard } from "@/components/cards/priority-icon";
import { Text } from "@/components/typography/Text";
import { Separator } from "@/components/ui/separator";
import { TitleInfo } from "@/components/cards/title-info";

export const IncidentInfo: FC<{ data: Incident.Item }> = observer((x) => (
  <IncidentCardWrapper className="gap-2">
    <div className="flex gap-4">
      <Text.SubtleMedium className="text-muted-foreground">{x.data.number}</Text.SubtleMedium>
      <PriorityCard data={x.data.priority} />
    </div>
    <Text.H3>{x.data.address}</Text.H3>
    <Separator />
    <Text.H4>Информация об объекте</Text.H4>
    <ul className="space-y-1">
      {x.data.info.map(([title, info], i) => (
        <TitleInfo key={i} title={title} info={info} />
      ))}
    </ul>
  </IncidentCardWrapper>
));