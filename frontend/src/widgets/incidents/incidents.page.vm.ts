import { makeAutoObservable } from "mobx";

export class IncidentsPageViewModel {
  constructor() {
    makeAutoObservable(this);
  }
}
