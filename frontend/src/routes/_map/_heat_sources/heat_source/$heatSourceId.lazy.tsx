import { HeatSourceCardReadonly } from "@/components/cards/heat-source.card";
import { Text } from "@/components/typography/Text";
import { Separator } from "@/components/ui/separator";
import { MapStore } from "@/stores/map.store";
import { cn } from "@/utils/cn";
import { SecondarySidebar } from "@/widgets/layoutMainSidebar/SecondarySidebar/secondary-sidebar.widget";
import { createLazyFileRoute, useNavigate } from "@tanstack/react-router";
import { observer } from "mobx-react-lite";
import { FC, PropsWithChildren, useEffect } from "react";

const vm = MapStore;

const Section: FC<PropsWithChildren & { className?: string; title?: string }> = (x) => (
  <>
    <Separator />
    <section className={cn("px-4 py-3 flex flex-col", x.className)}>
      {x.title && <Text.UiMedium className="pb-3">{x.title}</Text.UiMedium>}
      {x.children}
    </section>
    <Separator />
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

  const heatSource = vm.heatSources.find((v) => v.id.toString() === heatSourceId);

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
          <Section title="Подробная информация">
            <ul className="space-y-2">
              <TitleInfo title="Вид ТП" info="ЦТП" />
            </ul>
          </Section>
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

export const Route = createLazyFileRoute("/_map/_heat_sources/heat_source/$heatSourceId")({
  component: Page
});
