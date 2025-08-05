import { BACKEND_URL } from "@/shared/constants/endpoints";
import type { ApiResponse } from "@/shared/types/api";
import { api } from "../axios";
import { useQuery } from "@tanstack/react-query";
import { GITHUB_OAUTH_LOGIN_QUERY_KEY } from "@/shared/constants/query-keys";

const githubOauthLogin = async (code: string): Promise<ApiResponse<string>> => {
  const response = await api.get<{
    message: string;
    data: string;
  }>(`${BACKEND_URL}/api/v1/auth/github/callback?code=${code}`);

  return response.data;
};

export const useGithubOauthLogin = (code: string | null) => {
  return useQuery({
    queryKey: [GITHUB_OAUTH_LOGIN_QUERY_KEY, code],
    queryFn: () => githubOauthLogin(code!),
    enabled: !!code
  });
};
