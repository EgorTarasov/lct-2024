import { useIncidentsContext } from "@/components/hoc/incidents-context";
import { Text } from "@/components/typography/Text";
import { Button } from "@/components/ui/button";
import { createLazyFileRoute } from "@tanstack/react-router";

const Page = () => {
  const vm = useIncidentsContext();

  return (
    <div className="flex items-center justify-center size-full pt-40 appear flex-col gap-10">
      <Text.H3>Выберите объект из списка</Text.H3>
      <Button className="flex md:hidden" onClick={() => (vm.drawerOpen = true)}>
        Открыть список
      </Button>
    </div>
  );
};

export const Route = createLazyFileRoute("/_incidents/incidents/")({
  component: Page
});
