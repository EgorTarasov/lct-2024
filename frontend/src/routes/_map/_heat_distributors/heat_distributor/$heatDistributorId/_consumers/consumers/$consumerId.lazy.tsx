import { Section } from "@/components/cards/Section";
import { ConsumerCardReadonly } from "@/components/cards/consumer.card";
import { HeatDistributorCard } from "@/components/cards/heat-distributor.card";
import { IssueLink } from "@/components/cards/issue.card";
import { TitleInfo } from "@/components/cards/title-info";
import { Text } from "@/components/typography/Text";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger
} from "@/components/ui/accordion";
import { buttonVariants } from "@/components/ui/button";
import { Tooltip, TooltipContent, TooltipTrigger } from "@/components/ui/tooltip";
import { MapStore } from "@/stores/map.store";
import { cn } from "@/utils/cn";
import { SecondarySidebar } from "@/widgets/layoutMainSidebar/SecondarySidebar/secondary-sidebar.widget";
import { Link, createLazyFileRoute, useNavigate } from "@tanstack/react-router";
import { ChevronRightIcon } from "lucide-react";
import { observer } from "mobx-react-lite";

const Page = observer(() => {
  const navigate = useNavigate();
  const { consumerId, heatDistributorId } = Route.useParams();
  const vm = MapStore.consumersVm;
  const heatSource = MapStore.heatSourceVm.items.find((v) => v.id.toString() === heatDistributorId);

  if (!vm || !heatSource) {
    console.error("how ???");
    return null;
  }

  const consumer = vm.items.find((v) => v.id.toString() === consumerId);

  return (
    <SecondarySidebar
      closeAction={() => {
        navigate({
          to: "/heat_distributor/$heatDistributorId/consumers",
          params: { heatDistributorId }
        });
      }}>
      {consumer ? (
        <div className="flex flex-col">
          <ConsumerCardReadonly data={consumer} />
          <Section
            className="px-0 pb-2"
            title={<span className="px-4">Информация об объекте</span>}>
            <ul className="space-y-2 px-4">
              <TitleInfo title="Вид ТП" info="ЦТП" />
            </ul>
            <Accordion type="single" collapsible>
              <AccordionItem value="1">
                <AccordionTrigger className="px-4">Подробнее</AccordionTrigger>
                <AccordionContent className="px-4">tst</AccordionContent>
              </AccordionItem>
            </Accordion>
          </Section>
          {consumer.incidentCount > 0 && (
            <Section withoutSeparator className="pt-1">
              <IssueLink unom={consumer.unom} count={consumer.incidentCount} />
            </Section>
          )}
          <Section
            className="px-0"
            title={
              <div className="flex px-4 items-center justify-between w-full">
                <span>Источник тепла</span>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <Link
                      to="/heat_distributor/$heatDistributorId"
                      params={{ heatDistributorId }}
                      className={cn(
                        buttonVariants({ variant: "outline", size: "icon" }),
                        "shadow-none"
                      )}>
                      <ChevronRightIcon />
                    </Link>
                  </TooltipTrigger>
                  <TooltipContent>Перейти к источнику тепла</TooltipContent>
                </Tooltip>
              </div>
            }>
            <HeatDistributorCard className="mt-2" data={heatSource} />
          </Section>
        </div>
      ) : (
        <div className="flex flex-col gap-2 text-center pt-5">
          <Text.Large>Потребитель не найден</Text.Large>
          <span className="text-muted-foreground">Возможно он был удалён</span>
        </div>
      )}
    </SecondarySidebar>
  );
});

export const Route = createLazyFileRoute(
  "/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers/consumers/$consumerId"
)({
  component: Page
});
