import type { ApiResponse } from "@/shared/types/api";
import { api } from "../axios";
import { BACKEND_URL } from "@/shared/constants/endpoints";
import { useMutation, useQuery } from "@tanstack/react-query";
import {
  CONTRIBUTION_TYPES_QUERY_KEY,
  GOAL_LEVEL_TARGETS_QUERY_KEY,
  GOAL_LEVELS_QUERY_KEY,
  USER_ACTIVE_GOAL_LEVEL_QUERY_KEY,
  USER_GOAL_LEVEL_SUMMARY_QUERY_KEY
} from "@/shared/constants/query-keys";
import type {
  ContributionTypeDetail,
  GoalLevel,
  GoalLevelTarget,
  GoalSummary,
  SetUserGoalLevelRequest,
  UserCurrentGoalStatus,
  UserGoal,
  UserGoalLevelStatus
} from "@/shared/types/types";

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

const fetchGoalLevelTargets = async (
  goalLevel: GoalLevel
): Promise<ApiResponse<GoalLevelTarget[]>> => {
  const response = await api.post<{
    message: string;
    data: GoalLevelTarget[];
  }>(`${BACKEND_URL}/api/v1/goal/level/targets`, goalLevel);

  return response.data;
};

export const useGoalLevelTargets = () => {
  return useMutation({
    mutationFn: (goalLevel: GoalLevel) => fetchGoalLevelTargets(goalLevel)
  });
};

const fetchUserCurrentGoalStatus = async (): Promise<
  ApiResponse<UserCurrentGoalStatus>
> => {
  const response = await api.get<{
    message: string;
    data: UserCurrentGoalStatus;
  }>(`${BACKEND_URL}/api/v1/user/goal/level`);

  return response.data;
};

export const useUserCurrentGoalStatus = () => {
  return useQuery({
    queryKey: [USER_ACTIVE_GOAL_LEVEL_QUERY_KEY],
    queryFn: fetchUserCurrentGoalStatus
  });
};

const setUserGoalLevel = async (
  userGoalLevelRequest: SetUserGoalLevelRequest
): Promise<ApiResponse<UserGoalLevelStatus>> => {
  const response = await api.post<{
    message: string;
    data: UserGoalLevelStatus;
  }>(`${BACKEND_URL}/api/v1/user/goal/level`, userGoalLevelRequest);

  return response.data;
};

export const useSetUserGoalLevel = () => {
  return useMutation({
    mutationFn: (userGoalLevelRequest: SetUserGoalLevelRequest) =>
      setUserGoalLevel(userGoalLevelRequest)
  });
};

const resetUserGoalStatus = async (): Promise<ApiResponse<UserGoal>> => {
  const response = await api.post<{
    message: string;
    data: UserGoal;
  }>(`${BACKEND_URL}/api/v1/user/goal/level/reset`);

  return response.data;
};

export const useResetUserGoalStatus = () => {
  return useMutation({
    mutationFn: () => resetUserGoalStatus()
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

const fetchUserGoalSummary = async (): Promise<ApiResponse<GoalSummary[]>> => {
  const response = await api.get<{
    message: string;
    data: GoalSummary[];
  }>(`${BACKEND_URL}/api/v1/user/goal/summary`);

  return response.data;
};

export const useUserGoalSummary = () => {
  return useQuery({
    queryKey: [USER_GOAL_LEVEL_SUMMARY_QUERY_KEY],
    queryFn: fetchUserGoalSummary
  });
};
