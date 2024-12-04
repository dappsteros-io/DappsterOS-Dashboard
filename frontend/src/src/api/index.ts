"use client";
import { notification } from "antd";
import axios, { AxiosError, AxiosRequestConfig } from "axios";
const host = process.env.NEXT_PUBLIC_API_ADDR;
const port = process.env.NEXT_PUBLIC_API_PORT;

export const baseURL = `http://${host}:${port}/api/v1/`;
export const wsURL = `ws://${host}:${port}/`;

export default function api(url: string, config?: AxiosRequestConfig) {
  return axios(baseURL + url, {
    ...config,
    headers: {
      Authorization: "Bearer " + localStorage.getItem("token"),
      "Content-Type": "application/json",
    },
  })
    .then((res) => {
      if (res.config.method?.toLowerCase() != "get") {
        notification.success({
          message: res?.data?.title,
          description: res?.data?.message || res?.data?.detail,
        });
      }
      return res.data;
    })
    .catch((err: AxiosError) => {
      notification.error({
        message: (err?.response?.data as any)?.title,
        description: (err?.response?.data as any)?.message,
      });
      throw new Error((err?.response?.data as any)?.message);
    });
}
