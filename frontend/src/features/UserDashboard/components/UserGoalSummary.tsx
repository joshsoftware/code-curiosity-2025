import { useUserGoalSummary } from "@/api/queries/UserGoals";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer
} from "recharts";

const UserGoalSummaryChart = () => {
  const { data, isLoading, isError } = useUserGoalSummary();

  if (isLoading) return <div>Loading...</div>;
  if (isError || !data) return <div>Error loading goal summary</div>;

  const summaryData = data.data;
  const chartData = summaryData.map(item => ({
    day: new Date(item.snapshotDate).toLocaleDateString("en-GB", {
      day: "2-digit",
      month: "short"
    }),
    incomplete: item.incompleteGoalsCount,
    targetSet: item.targetSet,
    targetCompleted: item.targetCompleted
  }));

  return (
    <ResponsiveContainer width="100%" height={400}>
      <LineChart data={chartData}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="day" />
        <YAxis allowDecimals={false} />
        <Tooltip />
        <Legend />

        <Line
          type="monotone"
          dataKey="targetSet"
          stroke="#007bff"
          name="Target Set"
        />
        <Line
          type="monotone"
          dataKey="targetCompleted"
          stroke="#28a745"
          name="Target Completed"
        />
        <Line
          type="monotone"
          dataKey="incomplete"
          stroke="#dc3545"
          name="Incomplete Goals"
        />
      </LineChart>
    </ResponsiveContainer>
  );
};

export default UserGoalSummaryChart;
