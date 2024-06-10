import { NotificationFilters } from "@/constants/notification-filters";
import { Consumer } from "@/types/consumer.type";
import { Notification } from "@/types/notification.type";
import { Priority } from "@/types/priority.type";
import { makeAutoObservable } from "mobx";

class notificationStore {
  constructor() {
    makeAutoObservable(this);
  }

  _notifications: Notification.Item[] = [
    {
      id: 1,
      date: new Date(),
      description:
        "В зоне аварии находятся N потребителей. Предположительная дата устранения поломки по адресу ул. Московская, 11  — 01.08.2024.",
      title: "Произошла авария на источнике тепла",
      priority: Priority.Item.HIGH
    }
  ];

  sort: NotificationFilters.Sort = NotificationFilters.Sort.Priority;
  get notifications() {
    const notifications = this._notifications.slice();

    if (this.sort === NotificationFilters.Sort.Priority) {
      notifications.sort((a, b) => {
        return a.priority.localeCompare(b.priority);
      });
    }

    if (this.sort === NotificationFilters.Sort.Recent) {
      notifications.sort((a, b) => {
        return b.date.getTime() - a.date.getTime();
      });
    }

    return notifications;
  }

  resolve(notification: Notification.Item) {
    this._notifications = this._notifications.filter((v) => v !== notification);
  }
}

export const NotificationStore = new notificationStore();
