import { DisposableVm } from "@/utils/vm";
import { makeAutoObservable } from "mobx";
import L, { Map } from "leaflet";

class mapViewModel implements DisposableVm {
  private map: Map | null = null;
  constructor() {
    makeAutoObservable(this);
  }

  setMap(m: Map) {
    if (this.map) {
      this.dispose();
    }

    this.map = m;
  }

  dispose(): void {
    this.map = null;
  }
}

export const MapViewModel = new mapViewModel();
