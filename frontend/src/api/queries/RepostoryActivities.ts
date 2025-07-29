import type { ApiResponse } from "@/shared/types/api";
import type { RepositoryActivity } from "@/shared/types/types";
import { api } from "../axios";
import { BACKEND_URL } from "@/shared/constants/endpoints";
import { useQuery } from "@tanstack/react-query";
import { REPOSITORY_ACTIVITIES_QUERY_KEY } from "@/shared/constants/query-keys";

const fetchRepositoryActivivties = async (
  repoId: number
): Promise<ApiResponse<RepositoryActivity[]>> => {
  const response = await api.get<{
    message: string;
    data: RepositoryActivity[];
  }>(`${BACKEND_URL}/api/v1/user/repositories/contributions/recent/${repoId}`);

  return response.data;
};

export const useRepositoryActivities = (repoId: number) => {
  return useQuery({
    queryKey: [REPOSITORY_ACTIVITIES_QUERY_KEY, repoId],
    queryFn: () => fetchRepositoryActivivties(repoId),
    enabled: !!repoId
  });
};
