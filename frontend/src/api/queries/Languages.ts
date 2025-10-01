import type { ApiResponse } from "@/shared/types/api";
import type { Language } from "@/shared/types/types";
import { api } from "../axios";
import { BACKEND_URL } from "@/shared/constants/endpoints";
import { REPOSITORY_LANGUAGES_QUERY_KEY } from "@/shared/constants/query-keys";
import { useQuery } from "@tanstack/react-query";

const fetchRepositoryLanguages = async (
  id: number
): Promise<ApiResponse<Language[]>> => {
  const response = await api.get<{
    message: string;
    data: Language[];
  }>(`${BACKEND_URL}/api/v1/user/repositories/languages/${id}`);

  return response.data;
};

export const useRepositoryLanguages = (id: number) => {
  return useQuery({
    queryKey: [REPOSITORY_LANGUAGES_QUERY_KEY, id],
    queryFn: () => fetchRepositoryLanguages(id)
  });
};
