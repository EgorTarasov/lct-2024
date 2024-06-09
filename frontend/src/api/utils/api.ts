/* eslint-disable @typescript-eslint/no-explicit-any */
import axios, { AxiosRequestConfig, AxiosResponse } from "axios";
import { ZodObject, z } from "zod";
import { authToken } from "./authToken";
import { toast } from "sonner";

axios.defaults.baseURL = "https://api.lct.larek.tech/";

const getConfig = (config?: AxiosRequestConfig<unknown>) => ({
  ...config,
  headers: {
    Authorization: `Bearer ${authToken.get()}`,
    "Access-Control-Allow-Origin": "*",
    ...config?.headers
  }
});

const handleRequest = async <T extends ZodObject<any> | z.ZodArray<ZodObject<any>>>(
  req: () => Promise<AxiosResponse>,
  schema: T
): Promise<z.infer<T>> => {
  try {
    const res = await req();

    if (schema) {
      const validatedData = schema.parse(res.data);

      return validatedData as any;
    }

    return res as any;
  } catch (error) {
    if (error instanceof z.ZodError) {
      toast.error("Ошибка валидации данных", {
        description: error.errors.map((e) => e.message).join("\n")
      });
      console.error("Validation error:", error.errors);
    } else if (axios.isAxiosError(error)) {
      if (error.response?.status === 401) {
        toast.error("Сессия истекла, пожалуйста, войдите снова");
        authToken.remove();
        window.location.replace("/login");
      } else if (error.response?.status === 500) {
        const errorData = error.response.data as any;
        toast.error("Ошибка", {
          description: errorData?.error
            ? errorData.error
            : "Мы все узнаем и починим, попробуйте позже"
        });
      } else {
        toast.error("Неизвестная ошибка при выполнении запроса");
      }
      console.error("Axios error:", error.message, error.status);
    } else {
      toast.error("Неизвестная ошибка при выполнении запроса");
      console.error("Unexpected error:", error);
    }
    throw error;
  }
};

const get = <T extends ZodObject<any> | z.ZodArray<ZodObject<any>>>(
  path: string,
  schema: T,
  config?: AxiosRequestConfig<unknown>
): Promise<z.infer<T>> => handleRequest(() => axios.get(path, getConfig(config)), schema) as any;

const post = <T extends ZodObject<any> | z.ZodArray<ZodObject<any>>>(
  path: string,
  schema: T,
  variables?: unknown,
  config?: AxiosRequestConfig<unknown>
): Promise<z.infer<T>> =>
  handleRequest(() => axios.post(path, variables, getConfig(config)), schema) as any;

const put = <T extends ZodObject<any> | z.ZodArray<ZodObject<any>>>(
  path: string,
  schema: T,
  variables?: unknown,
  config?: AxiosRequestConfig<unknown>
): Promise<z.infer<T>> =>
  handleRequest(() => axios.put(path, variables, getConfig(config)), schema);

const del = <T extends ZodObject<any> | z.ZodArray<ZodObject<any>>>(
  path: string,
  schema: T,
  config?: AxiosRequestConfig<unknown>
): Promise<z.infer<T>> => handleRequest(() => axios.delete(path, getConfig(config)), schema) as any;

export default {
  get,
  post,
  put,
  delete: del
};
