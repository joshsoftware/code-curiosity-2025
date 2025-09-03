import Leaderboard from "@/features/UserDashboard/components/Leaderboard";
import RecentActivities from "@/features/UserDashboard/components/RecentActivities";
import Overview from "@/features/UserDashboard/components/Overview";
import UserGoalSummaryChart from "./UserGoalSummary";

const UserDashboardComponent = () => {
  return (
    <div className="h-full">
      <div className="flex w-full gap-4 h-full">
        <Leaderboard className="w-[45%]" />
        <div className="flex w-[55%] flex-col gap-4">
          <Overview className="h-[45%]" />
          <RecentActivities className="h-[55%]" />
        </div>
      </div>
      <div className="h-[50%] p-5">
        <UserGoalSummaryChart />
      </div>
    </div>
  );
};

export default UserDashboardComponent;
