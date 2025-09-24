import { BACKEND_URL } from "@/shared/constants/endpoints";
import { api } from "../axios";
import type { ApiResponse } from "@/shared/types/api";
import type { Badge } from "@/shared/types/types";
import { USER_BADGES_QUERY_KEY } from "@/shared/constants/query-keys";
import { useQuery } from "@tanstack/react-query";

const fetchUserBadges = async (): Promise<ApiResponse<Badge[]>> => {
    const response = await api.get<{
        message: string;
        data: Badge[];
    }>(`${BACKEND_URL}/api/v1/user/badges`);

    return response.data;
}

export const useUserBadges = () => {
    return useQuery({
        queryKey: [USER_BADGES_QUERY_KEY],
        queryFn: fetchUserBadges,
    });
}