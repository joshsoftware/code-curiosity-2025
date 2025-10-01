import { BACKEND_URL } from "@/shared/constants/endpoints";
import type { ApiResponse } from "@/shared/types/api";
import { api } from "../axios";
import { useMutation, useQuery } from "@tanstack/react-query";
import type {
  Admin,
  AdminCredentials,
  ContributionScore,
  ContributionScoreUpdate
} from "@/shared/types/types";
import {
  CONTRIBUTION_TYPES_QUERY_KEY,
  GET_ALL_USERS_QUERY_KEY
} from "@/shared/constants/query-keys";

const LogInAdmin = async (
  adminCredentials: AdminCredentials
): Promise<ApiResponse<Admin>> => {
  const response = await api.post<{
    message: string;
    data: Admin;
  }>(`${BACKEND_URL}/api/v1/auth/admin`, {
    email: adminCredentials.email,
    password: adminCredentials.password
  });

  return response.data;
};

export const useLogInAdmin = () => {
  return useMutation({
    mutationFn: (adminCredentials: AdminCredentials) =>
      LogInAdmin(adminCredentials)
  });
};

const getAllUsers = async (): Promise<ApiResponse<Admin[]>> => {
  const response = await api.get<{
    message: string;
    data: Admin[];
  }>(`${BACKEND_URL}/api/v1/users`);

  return response.data;
};

export const useGetAllUsers = () => {
  return useQuery({
    queryKey: [GET_ALL_USERS_QUERY_KEY],
    queryFn: getAllUsers
  });
};

const updateUserBlockStatus = async (
  userId: number,
  block: boolean
): Promise<ApiResponse<null>> => {
  const response = await api.patch<{
    message: string;
    data: null;
  }>(`${BACKEND_URL}/api/v1/users/${userId}`, {
    block
  });

  return response.data;
};

export const useUpdateUserBlockStatus = () => {
  return useMutation({
    mutationFn: ({ userId, block }: { userId: number; block: boolean }) =>
      updateUserBlockStatus(userId, block)
  });
};

const fetchContributionTypes = async (): Promise<
  ApiResponse<ContributionScore[]>
> => {
  const response = await api.get<{
    message: string;
    data: ContributionScore[];
  }>(`${BACKEND_URL}/api/v1/contributions/types`);

  return response.data;
};

export const useFetchContributionTypes = () => {
  return useQuery({
    queryKey: [CONTRIBUTION_TYPES_QUERY_KEY],
    queryFn: fetchContributionTypes
  });
};

const configureContributionScore = async (
  contributionScore: ContributionScoreUpdate[]
): Promise<ApiResponse<ContributionScore[]>> => {
  const response = await api.patch<{
    message: string;
    data: ContributionScore[];
  }>(`${BACKEND_URL}/api/v1/contributions/scores/configure`, contributionScore);

  return response.data;
};

export const useConfigureContributionScore = () => {
  return useMutation({
    mutationFn: (contributionScore: ContributionScoreUpdate[]) =>
      configureContributionScore(contributionScore)
  });
};
