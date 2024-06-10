import { GenericAbortSignal } from "axios";
import { MapDto } from "../models/map.model";
import api from "../utils/api";

export namespace MapEndpoint {
  export const getProperty = (
    latitude: number,
    longitude: number,
    radius: number,
    signal: GenericAbortSignal
  ) =>
    api.get("/geo/property", MapDto.Property, {
      params: { latitude, longitude, radius },
      signal
    });
}
