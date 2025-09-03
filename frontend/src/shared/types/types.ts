export interface User {
  userId: number;
  githubId: number;
  githubUsername: string;
  email: string;
  avatarUrl: string;
  currentBalance: number;
  isBlocked: boolean;
  isAdmin: boolean;
  password: string;
  isDeleted: boolean;
  deletedAt: string;
  createdAt: string;
  updatedAt: string;
}

export interface Badge {
  id: number;
  userId: number;
  badgeType: string;
  earnedAt: string;
  createdAt: string;
}

export interface LeaderboardUser {
  id: number;
  githubUsername: string;
  avatarUrl: string;
  contributedReposCount: number;
  currentBalance: number;
  rank: number;
}

export interface RecentActivity {
  userId: number;
  repositoryId: number;
  contributionScoreId: number;
  contributionType: string;
  balanceChange: number;
  contributedAt: string;
  githubEventId: number;
  githubRepoId: number;
  repoName: string;
  description: string;
  languagesUrl: string;
  repoUrl: string;
  ownerName: string;
  updateDate: string;
  contributorsUrl: string;
}

export interface Overview {
  type: string;
  count: number;
  totalCoins: number;
  month: string;
}

export interface Repositories {
  id: number;
  githubRepoId: number;
  repoName: string;
  description: string;
  languagesUrl: string;
  repoUrl: string;
  ownerName: string;
  updateDate: string;
  contributorsUrl: string;
  createdAt: string;
  updatedAt: string;
  languages: string[];
  totalCoinsEarned: number;
}

export interface Language {
  name: string;
  bytes: number;
  percentage: number;
}

export interface Repository {
  id: number;
  githubRepoId: number;
  repoName: string;
  description: string;
  languagesUrl: string;
  repoUrl: string;
  ownerName: string;
  updateDate: string;
  contributorsUrl: string;
  createdAt: string;
  updatedAt: string;
  languages: string[];
}

export interface Contributor {
  id: number;
  name: string;
  avatar_url: string;
  github_url: string;
  contributions: number;
}

export interface RepositoryActivity {
  id: number;
  userId: number;
  repositoryId: number;
  contributionScoreId: number;
  contributionType: string;
  balanceChange: number;
  contributedAt: string;
  githubEventId: string;
  createdAt: string;
  updatedAt: string;
}

export interface GoalLevel {
  id: number;
  level: string;
  createdAt: string;
  updatedAt: string;
}

export interface UserGoalTargetProgress {
  contributionType: string;
  target: number;
  progress: number;
}

export interface UserCurrentGoalStatus {
  userGoalId: number;
  level: string;
  status: string;
  monthStartedAt: string;
  createdAt: string;
  updatedAt: string;
  goalTargetProgress: UserGoalTargetProgress[];
}

export interface UserGoal {
  id: number;
  userId: number;
  goalId: number;
  status: string;
  createdAt: string;
  updatedAt: string;
}

export interface SetUserGoalLevelRequest {
  level: string;
  customTargets: CustomGoalLevelTarget[];
}

export interface CustomGoalLevelTarget {
  contributionType: string;
  target: number;
}

export interface ContributionTypeDetail {
  id: number;
  adminId: number;
  contributionType: string;
  score: number;
  createdAt: string;
  updatedAt: string;
}

export interface UserGoalLevelStatus {
  id: number;
  userId: number;
  goalId: number;
  status: string;
  createdAt: string;
  updatedAt: string;
}

export interface Admin {
  userId: number;
  githubId: number;
  githubUsername: string;
  email: string;
  avatarUrl: string;
  currentBalance: number;
  currentActiveGoalId: {
    Int64: number;
    Valid: boolean;
  };
  isBlocked: boolean;
  isAdmin: boolean;
  password: string;
  isDeleted: boolean;
  deletedAt: {
    Time: string;
    Valid: boolean;
  };
  createdAt: string;
  updatedAt: string;
  jwtToken: string;
}

export interface AdminCredentials {
  email: string;
  password: string;
}

export interface ContributionScore {
  id: number;
  adminId: number;
  contributionType: string;
  score: number;
  createdAt: string;
  updatedAt: string;
}

export interface ContributionScoreUpdate {
  contributionType: string;
  score: number;
}

export interface GoalSummary {
  id: number;
  userId: number;
  snapshotDate: string;
  incompleteGoalsCount: number;
  targetSet: number;
  targetCompleted: number;
  createdAt: string;
  updatedAt: string;
}
