import { IncidentDto } from "../models/incident.model";
import api from "../utils/api";

export namespace IncidentsEndpoint {
  export const getRecents = () => api.get("/issue/recent", IncidentDto.Item.array());
}
