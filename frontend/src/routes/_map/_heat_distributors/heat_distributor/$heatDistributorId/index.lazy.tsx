import { Section } from "@/components/cards/Section";
import { HeatDistributorCardReadonly } from "@/components/cards/heat-distributor.card";
import { TitleInfo } from "@/components/cards/title-info";
import { Text } from "@/components/typography/Text";
import {
  Accordion,
  AccordionItem,
  AccordionTrigger,
  AccordionContent
} from "@/components/ui/accordion";
import { buttonVariants } from "@/components/ui/button";
import { Tooltip, TooltipContent, TooltipTrigger } from "@/components/ui/tooltip";
import { MapStore } from "@/stores/map.store";
import { cn } from "@/utils/cn";
import { pluralizeIncident } from "@/utils/pluralize/incident";
import { SecondarySidebar } from "@/widgets/layoutMainSidebar/SecondarySidebar/secondary-sidebar.widget";
import { Link, createLazyFileRoute, useNavigate } from "@tanstack/react-router";
import { ChevronRightIcon } from "lucide-react";
import { observer } from "mobx-react-lite";

const vm = MapStore;

const Page = observer(() => {
  const navigate = useNavigate();
  const heatDistributorId = Route.useParams().heatDistributorId;

  const heatSource = vm.heatSourceVm.items.find((v) => v.id.toString() === heatDistributorId);

  return (
    <SecondarySidebar
      closeAction={() => {
        navigate({
          to: "/"
        });
      }}>
      {heatSource ? (
        <div className="flex flex-col">
          <HeatDistributorCardReadonly data={heatSource} />
          <Section className="px-0" title={<span className="px-4">Информация об объекте</span>}>
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
          {heatSource.incidentCount > 0 && (
            <Section withoutSeparator className="pt-0.5">
              <Link className="border px-4 py-3 rounded-xl hover:bg-background text-sm space-y-2 group transition-colors">
                <p>
                  Имеется информация о {heatSource.incidentCount}{" "}
                  {pluralizeIncident(heatSource.incidentCount)}, связанных с этим объектом
                </p>
                <div className="flex items-center gap-2 text-muted-foreground">
                  Перейти к инцидентам <ChevronRightIcon className="size-5" />
                </div>
              </Link>
            </Section>
          )}
          <Section
            title={
              <>
                <span>Потребители</span>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <Link
                      to="/heat_distributor/$heatDistributorId/consumers"
                      params={{ heatDistributorId }}
                      className={cn(
                        buttonVariants({ variant: "outline", size: "icon" }),
                        "shadow-none"
                      )}>
                      <ChevronRightIcon />
                    </Link>
                  </TooltipTrigger>
                  <TooltipContent>Перейти к потребителям</TooltipContent>
                </Tooltip>
              </>
            }></Section>
        </div>
      ) : (
        <div className="flex flex-col gap-2 text-center">
          <Text.Large>Источник тепла не найден</Text.Large>
          <span className="text-muted-foreground">Возможно он был удалён</span>
        </div>
      )}
    </SecondarySidebar>
  );
});

export const Route = createLazyFileRoute(
  "/_map/_heat_distributors/heat_distributor/$heatDistributorId/"
)({
  component: Page
});
