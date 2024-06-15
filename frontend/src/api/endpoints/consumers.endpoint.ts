import { z } from "zod";
import api from "../utils/api";
import { ConsumersDto } from "../models/consumers.model";

export namespace ConsumersEndpoint {
  export const getFilters = () => api.get("/consumers/filters", ConsumersDto.Filters);
}
