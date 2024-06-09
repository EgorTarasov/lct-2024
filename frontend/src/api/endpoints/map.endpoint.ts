import { MapDto } from "../models/map.model";
import api from "../utils/api";

export namespace MapEndpoint {
  export const getMoek = (latitude: number, longitude: number, radius: number) =>
    api.get("/geo/moek", MapDto.Moek, {
      params: { latitude, longitude, radius }
    });
}
