import type { Repositories } from "@/shared/types/types";
import { api } from "../axios";
import type { ApiResponse } from "@/shared/types/api";
import { BACKEND_URL } from "@/shared/constants/endpoints";
import { useQuery } from "@tanstack/react-query";
import { REPOSITORIES_KEY } from "@/shared/constants/query-keys";

const fetchRepositories = async (): Promise<ApiResponse<Repositories[]>> => {
  const response = await api.get<{
    message: string;
    data: Repositories[];
  }>(`${BACKEND_URL}/api/v1/user/repositories`);

  return response.data;
};

export const useRepositories = () => {
  return useQuery({
    queryKey: [REPOSITORIES_KEY],
    queryFn: fetchRepositories
  });
};
