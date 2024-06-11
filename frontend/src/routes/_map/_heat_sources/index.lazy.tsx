import { ConsumerCard } from "@/components/cards/consumer.card";
import { HeatSourceCard } from "@/components/cards/heat-source.card";
import { Text } from "@/components/typography/Text";
import { Breadcrumb, BreadcrumbItem } from "@/components/ui/breadcrumb";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import { MapStore } from "@/stores/map.store";
import { MainSidebar } from "@/widgets/layoutMainSidebar/main-sidebar.widget";
import { createLazyFileRoute } from "@tanstack/react-router";
import { observer } from "mobx-react-lite";

const vm = MapStore;

const Page = observer(() => <></>);

export const Route = createLazyFileRoute("/_map/_heat_sources/")({
  component: Page
});
