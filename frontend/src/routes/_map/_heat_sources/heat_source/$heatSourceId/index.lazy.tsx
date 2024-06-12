import { HeatSourceCardReadonly } from "@/components/cards/heat-source.card";
import { Text } from "@/components/typography/Text";
import { Button, buttonVariants } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { Tooltip, TooltipContent, TooltipTrigger } from "@/components/ui/tooltip";
import { MapStore } from "@/stores/map.store";
import { cn } from "@/utils/cn";
import { SecondarySidebar } from "@/widgets/layoutMainSidebar/SecondarySidebar/secondary-sidebar.widget";
import { Link, createLazyFileRoute, useNavigate } from "@tanstack/react-router";
import { ChevronRightIcon } from "lucide-react";
import { observer } from "mobx-react-lite";
import { FC, PropsWithChildren, ReactNode, useEffect } from "react";

function pluralizeIncident(count: number): string {
  const singular = "событии";
  const few = "событиях";

  if (count % 10 === 1 && count % 100 !== 11) {
    return singular;
  }
  return few;
}

const vm = MapStore;

const Section: FC<PropsWithChildren & { className?: string; title?: ReactNode }> = (x) => (
  <>
    <Separator />
    <section className={cn("px-4 py-3 pt-2 flex flex-col", x.className)}>
      {x.title && (
        <div className="h-9 flex items-center font-medium w-full justify-between">{x.title}</div>
      )}
      {x.children}
    </section>
  </>
);

const TitleInfo: FC<{ title: string; info: string }> = (x) => (
  <li className="text-sm space-x-1 flex">
    <h4 className="font-medium">{x.title}:</h4>
    <p>{x.info}</p>
  </li>
);

const Page = observer(() => {
  const navigate = useNavigate();
  const heatSourceId = Route.useParams().heatSourceId;

  const heatSource = vm.heatSourceVm.items.find((v) => v.id.toString() === heatSourceId);

  return (
    <SecondarySidebar
      closeAction={() => {
        navigate({
          to: "/"
        });
      }}>
      {heatSource ? (
        <div className="flex flex-col">
          <HeatSourceCardReadonly data={heatSource} />
          <Section title="Информация об объекте">
            <ul className="space-y-2">
              <TitleInfo title="Вид ТП" info="ЦТП" />
            </ul>
          </Section>
          {heatSource.incidentCount > 0 && (
            <Section>
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
                      to="/heat_source/$heatSourceId/consumers"
                      params={{ heatSourceId }}
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

export const Route = createLazyFileRoute("/_map/_heat_sources/heat_source/$heatSourceId/")({
  component: Page
});
