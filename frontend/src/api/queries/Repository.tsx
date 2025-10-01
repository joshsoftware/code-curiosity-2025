import { BACKEND_URL } from "@/shared/constants/endpoints";
import type { ApiResponse } from "@/shared/types/api";
import type { Repository } from "@/shared/types/types";
import { api } from "../axios";
import { REPOSITORY_KEY } from "@/shared/constants/query-keys";
import { useQuery } from "@tanstack/react-query";

const fetchRepository = async (
  repoId: number
): Promise<ApiResponse<Repository>> => {
  const response = await api.get<{
    message: string;
    data: Repository;
  }>(`${BACKEND_URL}/api/v1/user/repositories/${repoId}`);

  return response.data;
};

export const useRepository = (repoId: number) => {
  return useQuery({
    queryKey: [REPOSITORY_KEY, repoId],
    queryFn: () => fetchRepository(repoId)
  });
};
