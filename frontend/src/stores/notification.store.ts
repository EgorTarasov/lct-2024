import { makeAutoObservable } from "mobx";

class notificationStore {
  constructor() {
    makeAutoObservable(this);
  }

  get hasNotifications() {
    return true;
  }
}

export const NotificationStore = new notificationStore();
