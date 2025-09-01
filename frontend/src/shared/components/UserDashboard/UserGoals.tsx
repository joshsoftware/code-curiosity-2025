import { Progress } from "@/shared/components/ui/progress";
const goals = [
  { name: "Issue Resolve", current: 2, total: 5, progress: 40 },
  { name: "PR Review", current: 6, total: 8, progress: 75 },
  { name: "PR Merge", current: 2, total: 2, progress: 100 },
  { name: "PR Close", current: 1, total: 5, progress: 20 },
  { name: "PR Close", current: 1, total: 5, progress: 20 },
  { name: "PR Close", current: 1, total: 5, progress: 20 },
  { name: "PR Close", current: 1, total: 5, progress: 20 },
  { name: "PR Close", current: 1, total: 5, progress: 20 },
  { name: "PR Close", current: 1, total: 5, progress: 20 },
  { name: "PR Close", current: 1, total: 5, progress: 20 }
];

const UserGoals = () => {
  return (
    <div>
      <p className="text-cc-app-light-blue mb-4 text-left font-semibold">
        MY GOALS (BEGINNER)
      </p>
      <div className="space-y-6 text-white">
        {goals.map((goal, index) => (
          <div key={index}>
            <div className="mb-2 flex items-center justify-between">
              <span className="text-sm">{goal.name}</span>
              <span className="text-sm">
                {goal.current}/{goal.total}
              </span>
            </div>
            <Progress
              value={goal.progress}
              className="bg-cc-app-mid-blue h-4"
              indicatorClassName="from-cc-app-orange rounded-full bg-gradient-to-r to-yellow-400"
            />
          </div>
        ))}
      </div>
    </div>
  );
};

export default UserGoals;
