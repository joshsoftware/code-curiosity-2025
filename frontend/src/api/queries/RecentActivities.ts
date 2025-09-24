import type { ApiResponse } from "@/shared/types/api";
import type { RecentActivity } from "@/shared/types/types";
import { api } from "../axios";
import { BACKEND_URL } from "@/shared/constants/endpoints";
import { useQuery } from "@tanstack/react-query";
import {  RECENT_ACTIVITIES_QUERY_KEY } from "@/shared/constants/query-keys";

const fetchRecentActivities = async (): Promise<ApiResponse<RecentActivity[]>> => {
    const response = await api.get<{
        message: string;
        data: RecentActivity[];
    }>(`${BACKEND_URL}/api/v1/user/contributions/all`);

    return response.data;
};

export const useRecentActivities = () => {
    return useQuery({
        queryKey: [RECENT_ACTIVITIES_QUERY_KEY],
        queryFn: fetchRecentActivities,
    });
}