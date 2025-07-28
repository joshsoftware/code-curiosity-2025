import type { ApiResponse } from "@/shared/types/api";
import type { Overview } from "@/shared/types/types";
import { api } from "../axios";
import { BACKEND_URL } from "@/shared/constants/endpoints";
import { useQuery } from "@tanstack/react-query";
import { OVERVIEW_QUERY_KEY } from "@/shared/constants/query-keys";

interface OverviewParams {
    year: number;
    month: number;
}

const fetchOverview = async ({ year, month }: OverviewParams): Promise<ApiResponse<Overview[]>> => {
    const response = await api.get<{
        message: string;
        data: Overview[];
    }>(`${BACKEND_URL}/api/v1/user/overview`, {
        params: {
            year,
            month
        }
    });

    return response.data;
}

export const useOverview = ({ year, month }: OverviewParams) => {
    return useQuery({
        queryKey: [OVERVIEW_QUERY_KEY, year, month],
        queryFn: () => fetchOverview({ year, month }),
    });
}