import { NotificationFilters } from "@/constants/notification-filters";
import { Notification } from "@/types/notification.type";
import { Priority } from "@/types/priority.type";
import { makeAutoObservable } from "mobx";

class notificationStore {
  ws: WebSocket = new WebSocket("wss://push.larek.tech/stream?token=CqYaYW-WeaviCpP");

  constructor() {
    makeAutoObservable(this);

    this.ws.onopen = () => {
      console.log("Connected to Gotify WebSocket");
    };

    this.ws.onmessage = (event) => {
      const message = JSON.parse(event.data);
      console.log("New message:", message);
    };

    this.ws.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    this.ws.onclose = () => {
      console.log("WebSocket connection closed");
    };
  }

  _notifications: Notification[] = [
    {
      id: 1,
      date: new Date(),
      description:
        "В зоне аварии находятся N потребителей. Предположительная дата устранения поломки по адресу ул. Московская, 11  — 01.08.2024.",
      title: "Произошла авария на источнике тепла",
      priority: Priority.HIGH
    }
  ];

  sort: NotificationFilters.Sort = NotificationFilters.Sort.Priority;
  get notifications() {
    const notifications = this._notifications.slice();

    if (this.sort === NotificationFilters.Sort.Priority) {
      notifications.sort((a, b) => {
        return b.priority - a.priority;
      });
    }

    if (this.sort === NotificationFilters.Sort.Recent) {
      notifications.sort((a, b) => {
        return b.date.getTime() - a.date.getTime();
      });
    }

    return notifications;
  }

  resolve(notification: Notification) {
    this._notifications = this._notifications.filter((v) => v !== notification);
  }
}

export const NotificationStore = new notificationStore();
