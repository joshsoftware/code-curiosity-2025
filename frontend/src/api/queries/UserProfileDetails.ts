import type { User } from "@/shared/types/types";
import { api } from "../axios";
import { useMutation, useQuery } from "@tanstack/react-query";
import type { ApiResponse } from "@/shared/types/api";
import { LOGGED_IN_USER_QUERY_KEY } from "@/shared/constants/query-keys";
import { BACKEND_URL } from "@/shared/constants/endpoints";

const fetchLoggedInUser = async (): Promise<ApiResponse<User>> => {
  const response = await api.get<{
    message: string;
    data: User;
  }>(`${BACKEND_URL}/api/v1/auth/user`);

  return response.data;
};

export const useLoggedInUser = () => {
  return useQuery({
    queryKey: [LOGGED_IN_USER_QUERY_KEY],
    queryFn: fetchLoggedInUser
  });
};

const updateUserEmail = async (email: string): Promise<ApiResponse<null>> => {
  const response = await api.patch<{
    message: string;
    data: null;
  }>(`${BACKEND_URL}/api/v1/user/email`, {
    email: email
  });

  return response.data;
};

export const useUpdateUserEmail = () => {
  return useMutation({
    mutationFn: (email: string) => updateUserEmail(email)
  });
};

const softDeleteUser = async (userId: number): Promise<ApiResponse<null>> => {
  const response = await api.delete<{ message: string; data: null }>(
    `${BACKEND_URL}/api/v1/user/delete/${userId}`
  );

  return response.data;
};

export const useSoftDeleteUser = () => {
  return useMutation({
    mutationFn: (userId: number) => softDeleteUser(userId)
  });
};
