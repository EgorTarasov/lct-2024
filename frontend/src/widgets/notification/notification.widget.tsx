import { Text } from "@/components/typography/Text";
import { Button } from "@/components/ui/button";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Sheet, SheetContent, SheetHeader, SheetTitle, SheetTrigger } from "@/components/ui/sheet";
import { Tooltip, TooltipContent, TooltipTrigger } from "@/components/ui/tooltip";
import { NotificationStore } from "@/stores/notification.store";
import { Consumer } from "@/types/consumer.type";
import { cn } from "@/utils/cn";
import { getTimeDifference } from "@/utils/time";
import { BellIcon, Clock4Icon, XIcon } from "lucide-react";
import { observer } from "mobx-react-lite";
import { LayerSelect } from "./NotificationSort";
import { useEffect, useState } from "react";
import { Priority } from "@/types/priority.type";
import { NotificationCard } from "@/components/cards/notification.card";

const vm = NotificationStore;

export const NotificationWidget = observer(() => {
  const [open, setOpen] = useState(false);

  useEffect(() => {
    if (vm.notifications.length === 0) {
      setOpen(false);
    }
  }, [vm.notifications]);

  return (
    <Sheet open={open} onOpenChange={(v) => setOpen(v)}>
      <SheetTrigger asChild>
        <Button size="icon" variant="ghost" className="relative">
          <BellIcon />
          {vm.notifications.length > 0 && (
            <div
              className={cn(
                "absolute top-2 border border-card right-2 w-2 h-2 rounded-full bg-primary"
              )}
            />
          )}
        </Button>
      </SheetTrigger>
      <SheetContent className="flex flex-col gap-3 h-full overflow-hidden">
        <SheetHeader>
          <SheetTitle>Уведомления</SheetTitle>
        </SheetHeader>
        {vm.notifications.length > 0 && <LayerSelect />}
        <ScrollArea className="flex-1 space-y-2">
          {vm.notifications.map((n) => (
            <NotificationCard
              key={n.id}
              data={n}
              onClose={() => {
                vm.resolve(n);
              }}
            />
          ))}
          {vm.notifications.map((n) => (
            <NotificationCard
              key={n.id}
              data={n}
              onClose={() => {
                vm.resolve(n);
              }}
            />
          ))}
          {vm.notifications.map((n) => (
            <NotificationCard
              key={n.id}
              data={n}
              onClose={() => {
                vm.resolve(n);
              }}
            />
          ))}
          {vm.notifications.map((n) => (
            <NotificationCard
              key={n.id}
              data={n}
              onClose={() => {
                vm.resolve(n);
              }}
            />
          ))}
          {vm.notifications.map((n) => (
            <NotificationCard
              key={n.id}
              data={n}
              onClose={() => {
                vm.resolve(n);
              }}
            />
          ))}
          {vm.notifications.map((n) => (
            <NotificationCard
              key={n.id}
              data={n}
              onClose={() => {
                vm.resolve(n);
              }}
            />
          ))}
          {vm.notifications.map((n) => (
            <NotificationCard
              key={n.id}
              data={n}
              onClose={() => {
                vm.resolve(n);
              }}
            />
          ))}
          {vm.notifications.map((n) => (
            <NotificationCard
              key={n.id}
              data={n}
              onClose={() => {
                vm.resolve(n);
              }}
            />
          ))}
          {vm.notifications.map((n) => (
            <NotificationCard
              key={n.id}
              data={n}
              onClose={() => {
                vm.resolve(n);
              }}
            />
          ))}
          {vm.notifications.map((n) => (
            <NotificationCard
              key={n.id}
              data={n}
              onClose={() => {
                vm.resolve(n);
              }}
            />
          ))}
          {vm.notifications.map((n) => (
            <NotificationCard
              key={n.id}
              data={n}
              onClose={() => {
                vm.resolve(n);
              }}
            />
          ))}
          {vm.notifications.length === 0 && (
            <div className="flex flex-col gap-3 pt-6 text-center">
              <Text.p>Нет уведомлений</Text.p>
              <Text.Small className="text-muted-foreground">
                На иконке будет индикатор, когда они появятся
              </Text.Small>
            </div>
          )}
        </ScrollArea>
      </SheetContent>
    </Sheet>
  );
});
