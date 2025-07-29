import type { ApiResponse } from "@/shared/types/api";
import type { Contributor } from "@/shared/types/types";
import { api } from "../axios";
import { BACKEND_URL } from "@/shared/constants/endpoints";
import { REPOSITORY_CONTRIBUTORS_QUERY_KEY } from "@/shared/constants/query-keys";
import { useQuery } from "@tanstack/react-query";

const fetchRepositoryContributors = async (
  repoId: number
): Promise<ApiResponse<Contributor[]>> => {
  const response = await api.get<{
    message: string;
    data: Contributor[];
  }>(`${BACKEND_URL}/api/v1/user/repositories/contributors/${repoId}`);

  return response.data;
};

export const useRepositoryContributors = (repoId: number) => {
  return useQuery({
    queryKey: [REPOSITORY_CONTRIBUTORS_QUERY_KEY, repoId],
    queryFn: () => fetchRepositoryContributors(repoId),
    enabled: !!repoId
  });
};
