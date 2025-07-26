import axios from "axios";

import { BACKEND_URL } from "@/shared/constants/endpoints";
import { getAccessToken } from "@/shared/utils/local-storage";

export const api = axios.create({
  baseURL: BACKEND_URL
});

api.interceptors.request.use(
  config => {
    const token = getAccessToken();
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  error => Promise.reject(error)
);


