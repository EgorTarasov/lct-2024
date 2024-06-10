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
      <SheetContent className="flex flex-col gap-3">
        <SheetHeader>
          <SheetTitle>Уведомления</SheetTitle>
        </SheetHeader>
        {vm.notifications.length > 0 && <LayerSelect />}
        <ScrollArea className="h-full flex-1 space-y-2">
          {vm.notifications.map((n) => (
            <li
              key={n.id}
              className="p-4 pt-2 border space-y-2 list-none bg-card text-card-foreground rounded-md relative">
              <Tooltip>
                <TooltipTrigger asChild>
                  <Button
                    variant="ghost"
                    size="icon"
                    onClick={() => vm.resolve(n)}
                    className="ml-auto size-7 absolute p-1.5 right-0 top-0 hover:bg-transparent">
                    <XIcon />
                  </Button>
                </TooltipTrigger>
                <TooltipContent>Скрыть уведомление</TooltipContent>
              </Tooltip>
              <Text.UiMedium>{n.title}</Text.UiMedium>
              <Text.p>{n.description}</Text.p>
              <div className="flex items-center gap-2">
                <div
                  className="border px-2 rounded-full"
                  style={{
                    color: Priority.ItemMap[n.priority].color,
                    backgroundColor: Priority.ItemMap[n.priority].backgroundColor
                  }}>
                  <Text.Detail>{Priority.ItemMap[n.priority].alternateLocale} важность</Text.Detail>
                </div>
                <div className="gap-1 flex items-center text-muted-foreground">
                  <Clock4Icon className="size-4" />
                  <p className="text-xs">{getTimeDifference(n.date)}</p>
                </div>
              </div>
            </li>
          ))}
          {vm.notifications.length === 0 && (
            <Text.p className="text-center text-muted-foreground">Нет уведомлений</Text.p>
          )}
        </ScrollArea>
      </SheetContent>
    </Sheet>
  );
});
