export interface User {
  id: number;
  githubId: string;
  githubUsername: string;
  avatarUrl: string;
  email: string | null;
  currentActiveGoalId: number | null;
  currentBalance: number;
  isBlocked: boolean;
  isAdmin: boolean;
  password: string;
  isDeleted: boolean;
  DeletedAt: Date | null;
  createdAt: Date;
  updatedAt: Date;
}
