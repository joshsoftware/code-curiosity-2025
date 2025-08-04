export interface User {
  userId: number;
  githubId: number;
  githubUsername: string;
  email: string;
  avatarUrl: string;
  currentBalance: number;
  currentActiveGoalId: number;
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
  avatarUrl: string;
  githubUrl: string;
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
