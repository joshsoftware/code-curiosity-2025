import { BACKEND_URL } from "@/shared/constants/endpoints";
import type { ApiResponse } from "@/shared/types/api";
import { api } from "../axios";
import { useMutation } from "@tanstack/react-query";
import type { AdminCredentials } from "@/shared/types/types";

const LogInAdmin = async (
  adminCredentials: AdminCredentials
): Promise<ApiResponse<null>> => {
  const response = await api.patch<{
    message: string;
    data: null;
  }>(`${BACKEND_URL}/api/v1/auth/admin`, {
    email: adminCredentials.email,
    password: adminCredentials.password
  });

  return response.data;
};

export const useLogInAdmin = () => {
  return useMutation({
    mutationFn: (adminCredentials: AdminCredentials) =>
      LogInAdmin(adminCredentials)
  });
};
