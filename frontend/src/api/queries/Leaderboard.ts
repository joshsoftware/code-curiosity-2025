import type { LeaderboardUser } from "@/shared/types/types";
import { api } from "../axios";
import { BACKEND_URL } from "@/shared/constants/endpoints";
import type { ApiResponse } from "@/shared/types/api";
import { CURRENT_USER_RANK_QUERY_KEY, LEADERBOARD_QUERY_KEY } from "@/shared/constants/query-keys";
import { useQuery } from "@tanstack/react-query";

const fetchLeaderboard = async (): Promise<ApiResponse<LeaderboardUser[]>> => {
    const response = await api.get<{
        message: string;
        data: LeaderboardUser[];
    }>(`${BACKEND_URL}/api/v1/leaderboard`);

    return response.data;
}

export const useLeaderboard = () => {
    return useQuery({
        queryKey: [LEADERBOARD_QUERY_KEY],
        queryFn: fetchLeaderboard,
    });
}

const fetchCurrentUserRank = async (): Promise<ApiResponse<LeaderboardUser>> => {
    const response = await api.get<{
        message: string;
        data: LeaderboardUser;
    }>(`${BACKEND_URL}/api/v1/user/leaderboard`);

    return response.data;
}

export const useCurrentUserRank = () => {
    return useQuery({
        queryKey: [CURRENT_USER_RANK_QUERY_KEY],
        queryFn: fetchCurrentUserRank,
    });
}