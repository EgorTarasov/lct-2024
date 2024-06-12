import { z } from "zod";
import api from "../utils/api";

const anySchema = z.object({});

export namespace ConsumersEndpoint {
  export const getFilters = () => api.get("/consumers/filters", anySchema);
}
