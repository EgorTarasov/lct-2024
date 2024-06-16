import { z } from "zod";
import { MapDto } from "./map.model";
import { ConsumersDto } from "./consumers.model";

export namespace IncidentDto {
  export const Item = z.object({
    id: z.number(),
    openedAt: z.string(),
    closedAt: z.string().nullable(),
    title: z.string(),
    status: z.string(),
    priority: z.string(),
    unom: z.number(),
    createdAt: z.string(),
    updatedAt: z.string(),
    mkdConsumer: z.unknown().nullable(),
    dispatchServices: z.unknown().nullable(),
    graph: z.object({
      name: z.string(),
      values: z
        .object({
          temp: z.number(),
          time: z.number()
        })
        .array()
        .nullable()
    }),
    measurements: z.unknown().nullable(),
    heatingPoint: MapDto.Property.optional(),
    relatedObjects: z.object({
      unom: z.number(),
      consumers: ConsumersDto.Item.array().optional(),
      heatingPoint: MapDto.Property.optional()
    })
  });
  export type Item = z.infer<typeof Item>;
}
