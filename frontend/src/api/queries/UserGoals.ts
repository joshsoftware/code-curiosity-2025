import type { ApiResponse } from "@/shared/types/api";
import { api } from "../axios";
import { BACKEND_URL } from "@/shared/constants/endpoints";
import { useMutation, useQuery } from "@tanstack/react-query";
import {
  CONTRIBUTION_TYPES_QUERY_KEY,
  GOAL_LEVELS_QUERY_KEY,
  USER_ACTIVE_GOAL_LEVEL_QUERY_KEY,
  USER_GOAL_LEVEL_PROGRESS_QUERY_KEY
} from "@/shared/constants/query-keys";
import type {
  ContributionTypeDetail,
  CustomGoalLevelTarget,
  CustomGoalLevelTargetResponse,
  GoalLevel,
  GoalLevelProgress
} from "@/shared/types/types";

const fetchUserActiveGoalLevel = async (): Promise<ApiResponse<string>> => {
  const response = await api.get<{
    message: string;
    data: string;
  }>(`${BACKEND_URL}/api/v1/user/goal/level`);

  return response.data;
};

export const useUserActiveGoalLevel = () => {
  return useQuery({
    queryKey: [USER_ACTIVE_GOAL_LEVEL_QUERY_KEY],
    queryFn: fetchUserActiveGoalLevel
  });
};

const setUserGoalLevel = async (
  selectedLevel: string
): Promise<ApiResponse<number>> => {
  const response = await api.patch<{
    message: string;
    data: number;
  }>(`${BACKEND_URL}/api/v1/user/goal/level`, {
    level: selectedLevel
  });

  return response.data;
};

export const useSetUserGoalLevel = () => {
  return useMutation({
    mutationFn: (selectedLevel: string) => setUserGoalLevel(selectedLevel)
  });
};

const fetchGoalLevels = async (): Promise<ApiResponse<GoalLevel[]>> => {
  const response = await api.get<{
    message: string;
    data: GoalLevel[];
  }>(`${BACKEND_URL}/api/v1/goal/level`);

  return response.data;
};

export const useGoalLevels = () => {
  return useQuery({
    queryKey: [GOAL_LEVELS_QUERY_KEY],
    queryFn: fetchGoalLevels
  });
};

const fetchUserGoalLevelProgress = async (): Promise<
  ApiResponse<GoalLevelProgress[]>
> => {
  const response = await api.get<{
    message: string;
    data: GoalLevelProgress[];
  }>(`${BACKEND_URL}/api/v1/user/goal/level/progress`);

  return response.data;
};

export const useUserGoalLevelProgress = () => {
  return useQuery({
    queryKey: [USER_GOAL_LEVEL_PROGRESS_QUERY_KEY],
    queryFn: fetchUserGoalLevelProgress
  });
};

const createCustomGoalLevelTarget = async (
  customGoalLevelTarget: CustomGoalLevelTarget[]
): Promise<ApiResponse<CustomGoalLevelTargetResponse[]>> => {
  const response = await api.post<{
    message: string;
    data: CustomGoalLevelTargetResponse[];
  }>(
    `${BACKEND_URL}/api/v1/user/goal/level/custom/targets`,
    customGoalLevelTarget
  );

  return response.data;
};

export const useCustomGoalLevelTarget = () => {
  return useMutation({
    mutationFn: (customGoalLevelTarget: CustomGoalLevelTarget[]) =>
      createCustomGoalLevelTarget(customGoalLevelTarget)
  });
};

const fetchAllContributionTypes = async (): Promise<
  ApiResponse<ContributionTypeDetail[]>
> => {
  const response = await api.get<{
    message: string;
    data: ContributionTypeDetail[];
  }>(`${BACKEND_URL}/api/v1/contributions/types`);

  return response.data;
};

export const useAllContributionTypes = () => {
  return useQuery({
    queryKey: [CONTRIBUTION_TYPES_QUERY_KEY],
    queryFn: fetchAllContributionTypes
  });
};
