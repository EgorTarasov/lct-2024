import { IssueCard } from "@/components/cards/issue.card";
import { useIncidentsContext } from "@/components/hoc/incidents-context";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger
} from "@/components/ui/accordion";
import { Button } from "@/components/ui/button";
import { IssueLocaleMap } from "@/types/issue.type";
import { IncidentComments } from "@/widgets/incidents/components/comments";
import { CoolingChart } from "@/widgets/incidents/components/cooling-chart";
import { IncidentCardWrapper } from "@/widgets/incidents/components/incident-card-wrapper";
import { IncidentsMap } from "@/widgets/incidents/components/incidents.map";
import { IncidentInfo } from "@/widgets/incidents/incident-info";
import { DependentObjectsSection } from "@/widgets/incidents/setions/DependentObjects.section";
import { createLazyFileRoute } from "@tanstack/react-router";
import { format } from "date-fns";
import { observer } from "mobx-react-lite";
import { useEffect } from "react";

const Page = observer(() => {
  const unom = Route.useParams().unom;
  const vm = useIncidentsContext();

  useEffect(() => {
    vm.select(unom);
  }, [unom]);

  if (!vm.selected) return null;

  const item = vm.selected;

  return (
    <div className="flex flex-col gap-6 px-4 appear">
      <Button variant="outline" className="w-fit">
        Задача решена
      </Button>
      <div className="grid grid-cols-1 gap-6 lg:grid-cols-2">
        <IncidentInfo data={item} />
        <div className="flex flex-col gap-2">
          <IncidentCardWrapper
            className="border-2 h-fit gap-2"
            style={{ borderColor: IssueLocaleMap[item.incidentIssue].borderColor }}>
            <IssueCard data={item.incidentIssue} />
            <p>
              <span className="font-medium">Дата фиксации события: </span>
              {format(item.date, "dd.MM.yyyy HH:mm")}
            </p>
          </IncidentCardWrapper>
          <IncidentComments data={item} />
        </div>
      </div>
      <Accordion type="single" defaultValue={"dependency"} collapsible>
        <AccordionItem value="dependency" defaultChecked>
          <AccordionTrigger>Зависимые объекты</AccordionTrigger>
          <AccordionContent className="grid grid-cols-1 gap-6 lg:grid-cols-2">
            <DependentObjectsSection data={item} selectedUnom={unom} />
            <IncidentsMap vm={vm} />
          </AccordionContent>
        </AccordionItem>
        <AccordionItem value="math" defaultChecked>
          <AccordionTrigger>Данные по счетчикам</AccordionTrigger>
          <AccordionContent className="grid grid-cols-1 gap-6 lg:grid-cols-[1.4fr_1fr]">
            <CoolingChart />
          </AccordionContent>
        </AccordionItem>
        <AccordionItem value="history" defaultChecked>
          <AccordionTrigger>История событий</AccordionTrigger>
          <AccordionContent className="grid grid-cols-1 gap-6 lg:grid-cols-[1.4fr_1fr]"></AccordionContent>
        </AccordionItem>
      </Accordion>
    </div>
  );
});

export const Route = createLazyFileRoute("/_incidents/incidents/$unom")({
  component: Page
});
