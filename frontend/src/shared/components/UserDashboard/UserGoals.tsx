import { useState } from "react";
import { Progress } from "@/shared/components/ui/progress";
import { Button } from "@/shared/components/ui/button";
import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter
} from "@/shared/components/ui/dialog";
import { Loader2 } from "lucide-react";
import {
  useAllContributionTypes,
  useCustomGoalLevelTarget,
  useGoalLevels,
  useSetUserGoalLevel,
  useUserActiveGoalLevel,
  useUserGoalLevelProgress
} from "@/api/queries/UserGoals";
import { useQueryClient } from "@tanstack/react-query";
import {
  USER_ACTIVE_GOAL_LEVEL_QUERY_KEY,
  USER_GOAL_LEVEL_PROGRESS_QUERY_KEY
} from "@/shared/constants/query-keys";
import type {
  ContributionTypeDetail,
  CustomGoalLevelTarget
} from "@/shared/types/types";

const UserGoals = () => {
  const [dialogOpen, setDialogOpen] = useState(false);
  const [isSettingLevel, setIsSettingLevel] = useState(false);
  const [isCustomDialogOpen, setIsCustomDialogOpen] = useState(false);

  const [customGoals, setCustomGoals] = useState<CustomGoalLevelTarget[]>([]);
  const [selectedType, setSelectedType] = useState("");
  const [target, setTarget] = useState("");

  const { data: userGoalLevelRes, isLoading: isGoalLevelLoading } =
    useUserActiveGoalLevel();
  const { data: goalLevelsRes, isLoading: isGoalLevelsLoading } =
    useGoalLevels();
  const { data: userProgressRes, isLoading: isProgressLoading } =
    useUserGoalLevelProgress();
  const { mutate: setGoalLevel } = useSetUserGoalLevel();
  const { data: contributionTypesRes } = useAllContributionTypes();
  const { mutate: setCustomTarget, isPending: isSettingCustom } =
    useCustomGoalLevelTarget();

  const queryClient = useQueryClient();

  const userLevel = userGoalLevelRes?.data ?? "";
  const goalLevels = goalLevelsRes?.data ?? [];
  const userProgress = userProgressRes?.data ?? [];
  const allTypes: ContributionTypeDetail[] = contributionTypesRes?.data ?? [];

  const handleLevelSelect = (level: string) => {
    setIsSettingLevel(true);
    setGoalLevel(level, {
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: [USER_ACTIVE_GOAL_LEVEL_QUERY_KEY]
        });
        queryClient.invalidateQueries({
          queryKey: [USER_GOAL_LEVEL_PROGRESS_QUERY_KEY]
        });

        setIsSettingLevel(false);

        if (level.toLowerCase() === "custom") {
          setDialogOpen(false); // close default dialog
          setIsCustomDialogOpen(true); // open custom dialog
        } else {
          setDialogOpen(false);
        }
      },
      onError: () => {
        console.error("Failed to set user goal level");
        setIsSettingLevel(false);
      }
    });
  };

  const handleAddCustomGoal = () => {
    if (!selectedType || !target) return;
    if (customGoals.some(g => g.contributionType === selectedType)) return;

    setCustomGoals(prev => [
      ...prev,
      {
        contributionType: selectedType,
        target: Number(target)
      }
    ]);
    setSelectedType("");
    setTarget("");
  };

  const handleSubmitCustomGoals = () => {
    setCustomTarget(customGoals, {
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: [USER_ACTIVE_GOAL_LEVEL_QUERY_KEY]
        });
        queryClient.invalidateQueries({
          queryKey: [USER_GOAL_LEVEL_PROGRESS_QUERY_KEY]
        });
        setIsCustomDialogOpen(false);
        setCustomGoals([]);
      },
      onError: () => {
        console.error("Failed to set custom goal targets");
      }
    });
  };

  if (isGoalLevelLoading || isGoalLevelsLoading || isProgressLoading) {
    return (
      <div className="flex items-center gap-2 text-white">
        <Loader2 className="h-5 w-5 animate-spin" />
        Loading goals...
      </div>
    );
  }

  return (
    <div>
      <p className="text-cc-app-light-blue mb-4 text-left font-semibold">
        MY GOALS {userLevel && `(${userLevel.toUpperCase()})`}
      </p>

      {userLevel ? (
        <div className="space-y-6 text-white">
          {userProgress.map((goal, index) => {
            const percent = goal.targetCount
              ? Math.min((goal.achievedCount / goal.targetCount) * 100, 100)
              : 0;

            return (
              <div key={index + goal.targetCount}>
                <div className="mb-2 flex items-center justify-between">
                  <span className="text-sm capitalize">
                    {goal.contributionType.replace(/([A-Z])/g, " $1")}
                  </span>
                  <span className="text-sm">
                    {goal.achievedCount}/{goal.targetCount}
                  </span>
                </div>
                <Progress
                  value={percent}
                  className="bg-cc-app-mid-blue h-4"
                  indicatorClassName="from-cc-app-orange rounded-full bg-gradient-to-r to-yellow-400"
                />
              </div>
            );
          })}
        </div>
      ) : (
        <div className="bg-cc-app-mid-blue rounded-xl border border-gray-200 p-6 text-white shadow-md">
          <p className="mb-2 text-lg font-semibold">No Active Goal Set</p>
          <p className="mb-4 text-sm text-white">
            You haven't selected a goal level for this month yet. Choose a level
            to start tracking your contributions.
          </p>
          <Dialog open={dialogOpen} onOpenChange={setDialogOpen}>
            <DialogTrigger asChild>
              <Button variant="ccAppOutlineMidBlue">Set My Goal</Button>
            </DialogTrigger>

            <DialogContent className="bg-cc-app-mid-blue text-black">
              <DialogHeader>
                <DialogTitle>Select Goal Level for the month</DialogTitle>
              </DialogHeader>

              {!isSettingLevel ? (
                <div className="space-y-2">
                  {goalLevels.map(level => (
                    <Button
                      key={level.id}
                      variant="outline"
                      className="hover:bg-cc-app-sky-blue w-full capitalize hover:cursor-pointer"
                      onClick={() => handleLevelSelect(level.level)}
                    >
                      {level.level}
                    </Button>
                  ))}
                </div>
              ) : (
                <div className="text-cc-app-light-blue flex items-center justify-center gap-2 py-6">
                  <Loader2 className="h-5 w-5 animate-spin" />
                  Setting your goal...
                </div>
              )}

              <DialogFooter>
                <Button variant="ghost" onClick={() => setDialogOpen(false)}>
                  Cancel
                </Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        </div>
      )}

      <Dialog open={isCustomDialogOpen} onOpenChange={setIsCustomDialogOpen}>
        <DialogContent className="bg-cc-app-mid-blue text-black">
          <DialogHeader>
            <DialogTitle>Set Custom Contribution Goals</DialogTitle>
          </DialogHeader>

          <div className="space-y-4">
            <div className="flex gap-2">
              <select
                value={selectedType}
                onChange={e => setSelectedType(e.target.value)}
                className="w-full rounded border p-2"
              >
                <option value="">Select Type</option>
                {allTypes.map(type => (
                  <option key={type.id} value={type.contributionType}>
                    {type.contributionType}
                  </option>
                ))}
              </select>
              <input
                type="number"
                min={1}
                className="w-1/2 rounded border p-2"
                placeholder="Target"
                value={target}
                onChange={e => setTarget(e.target.value)}
              />
              <Button
                variant="ccAppOutlineMidBlue"
                onClick={handleAddCustomGoal}
              >
                Add
              </Button>
            </div>

            {customGoals.length > 0 && (
              <div className="space-y-2">
                {customGoals.map((goal, idx) => (
                  <div
                    key={goal.contributionType}
                    className="flex justify-between rounded bg-white/20 px-3 py-2 text-white"
                  >
                    <span className="capitalize">{goal.contributionType}</span>
                    <span>{goal.target}</span>
                    <Button
                      size="sm"
                      variant="ghost"
                      onClick={() =>
                        setCustomGoals(customGoals.filter((_, i) => i !== idx))
                      }
                    >
                      Remove
                    </Button>
                  </div>
                ))}
              </div>
            )}
          </div>

          <DialogFooter className="pt-4">
            <Button
              disabled={customGoals.length === 0 || isSettingCustom}
              onClick={handleSubmitCustomGoals}
            >
              {isSettingCustom ? (
                <div className="flex items-center gap-2">
                  <Loader2 className="h-4 w-4 animate-spin" /> Saving...
                </div>
              ) : (
                "Save Custom Goal"
              )}
            </Button>
            <Button
              variant="ghost"
              onClick={() => setIsCustomDialogOpen(false)}
            >
              Cancel
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default UserGoals;
