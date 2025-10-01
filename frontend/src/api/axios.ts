import axios from "axios";

import { BACKEND_URL } from "@/shared/constants/endpoints";
import { clearAccessToken, getAccessToken } from "@/shared/utils/local-storage";
import { LOGIN_PATH } from "@/shared/constants/routes";

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

api.interceptors.response.use(
  response => response,
  error => {
    if (error.response?.status === 401) {
      clearAccessToken();
      window.location.href = LOGIN_PATH; 
    }
    return Promise.reject(error);
  }
);


