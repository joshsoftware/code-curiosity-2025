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
  useGoalLevels,
  useResetUserGoalStatus,
  useSetUserGoalLevel,
  useUserCurrentGoalStatus
} from "@/api/queries/UserGoals";
import { useQueryClient } from "@tanstack/react-query";
import { USER_ACTIVE_GOAL_LEVEL_QUERY_KEY } from "@/shared/constants/query-keys";
import type {
  ContributionTypeDetail,
  CustomGoalLevelTarget
} from "@/shared/types/types";
import { toast } from "sonner";

const UserGoals = () => {
  const [dialogOpen, setDialogOpen] = useState(false);
  const [resetDialogOpen, setResetDialogOpen] = useState(false);
  const [isSettingLevel, setIsSettingLevel] = useState(false);
  const [isCustomDialogOpen, setIsCustomDialogOpen] = useState(false);

  const [customGoals, setCustomGoals] = useState<CustomGoalLevelTarget[]>([]);
  const [selectedType, setSelectedType] = useState("");
  const [target, setTarget] = useState("");

  const { data: userGoalLevelRes, isLoading: isGoalLevelLoading } =
    useUserCurrentGoalStatus();
  const { data: goalLevelsRes, isLoading: isGoalLevelsLoading } =
    useGoalLevels();
  const { mutate: setGoalLevel } = useSetUserGoalLevel();
  const { mutate: resetGoalStatus } = useResetUserGoalStatus();
  const { data: contributionTypesRes } = useAllContributionTypes();

  const queryClient = useQueryClient();

  const userLevel = userGoalLevelRes?.data ?? null;
  const goalLevels = goalLevelsRes?.data ?? [];
  const allTypes: ContributionTypeDetail[] = contributionTypesRes?.data ?? [];

  const createdAt = userLevel?.createdAt
    ? new Date(userLevel?.createdAt)
    : null;

  const isWithin48HoursOrGreaterThan30Days = (createdAt?: Date | null) => {
    if (!createdAt) return false;
    const diff = Date.now() - createdAt.getTime();
    return diff < 48 * 60 * 60 * 1000 || diff > 30 * 24 * 60 * 60 * 1000;
  };

  const handleLevelSelect = (level: string) => {
    setIsSettingLevel(true);

    if (level.toLowerCase() === "custom") {
      setDialogOpen(false);
      setIsCustomDialogOpen(true);
      setIsSettingLevel(false);
      return;
    }

    setGoalLevel(
      { level, customTargets: [] },
      {
        onSuccess: () => {
          queryClient.invalidateQueries({
            queryKey: [USER_ACTIVE_GOAL_LEVEL_QUERY_KEY]
          });
          toast.success("goal set successfully");
          setIsSettingLevel(false);
          setDialogOpen(false);
        },
        onError: () => {
          toast.error("Failed to set user goal level");
          setIsSettingLevel(false);
        }
      }
    );
  };

  const handleGoalReset = () => {
    resetGoalStatus(undefined, {
      onSuccess: () => {
        queryClient.removeQueries({
          queryKey: [USER_ACTIVE_GOAL_LEVEL_QUERY_KEY]
        });
        queryClient.refetchQueries({
          queryKey: [USER_ACTIVE_GOAL_LEVEL_QUERY_KEY]
        });
        toast.success("goal reset successfully");
        setResetDialogOpen(false);
      },
      onError: (err: any) => {
        const message =
          err?.response?.data?.message || "Failed to reset goal status";
        toast.error(message);
      }
    });
  };

  const handleAddCustomGoal = () => {
    if (!selectedType || !target) return;
    if (customGoals.some(g => g.contributionType === selectedType)) return;
    setCustomGoals(prev => [
      ...prev,
      { contributionType: selectedType, target: Number(target) }
    ]);
    setSelectedType("");
    setTarget("");
  };

  const handleSubmitCustomGoals = () => {
    setGoalLevel(
      { level: "Custom", customTargets: customGoals },
      {
        onSuccess: () => {
          queryClient.invalidateQueries({
            queryKey: [USER_ACTIVE_GOAL_LEVEL_QUERY_KEY]
          });
          toast.success("goal set successfully");
          setIsCustomDialogOpen(false);
          setCustomGoals([]);
        },
        onError: () => toast.error("Failed to set custom goal targets")
      }
    );
  };

  if (isGoalLevelLoading || isGoalLevelsLoading) {
    return (
      <div className="flex items-center gap-2 text-white">
        <Loader2 className="h-5 w-5 animate-spin" />
        Loading goals...
      </div>
    );
  }

  return (
    <div>
      <div className="mb-4 flex items-center justify-between">
        <p className="text-cc-app-light-blue mb-4 text-left font-semibold">
          MY GOALS {userLevel?.level && `(${userLevel.level.toUpperCase()})`}
        </p>
        {isWithin48HoursOrGreaterThan30Days(createdAt) && (
          <Dialog
            open={resetDialogOpen}
            onOpenChange={open => setResetDialogOpen(open)}
          >
            <DialogTrigger asChild>
              <Button variant="ccAppOutlineMidBlue">Reset</Button>
            </DialogTrigger>
            <DialogContent className="bg-cc-app-gray-background text-black">
              <DialogHeader>
                <DialogTitle>Confirm Goal Reset</DialogTitle>
              </DialogHeader>
              <p className="mb-4 text-sm">
                - Reset is available within 48 hours of setting a goal.
                <br />- After 48 hours, goals reset automatically after 30 days.
              </p>
              <DialogFooter>
                <Button
                  variant="ghost"
                  onClick={() => setResetDialogOpen(false)}
                >
                  Cancel
                </Button>
                <Button variant="ccAppOutlineMidBlue" onClick={handleGoalReset}>
                  Confirm Reset
                </Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        )}
      </div>

      {userLevel ? (
        <div className="space-y-6 text-white">
          {userLevel.goalTargetProgress?.map((goal, idx) => {
            const percent = goal.target
              ? Math.min((goal.progress / goal.target) * 100, 100)
              : 0;
            return (
              <div key={idx + goal.contributionType}>
                <div className="mb-2 flex items-center justify-between">
                  <span className="text-sm capitalize">
                    {goal.contributionType.replace(/([A-Z])/g, " $1")}
                  </span>
                  <span className="text-sm">
                    {goal.progress}/{goal.target}
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
        <div className="bg-cc-app-mid-blue rounded-xl border p-6 text-white shadow-md">
          <p className="mb-2 text-lg font-semibold">No Active Goal Set</p>
          <p className="mb-4 text-sm">
            You haven't selected a goal level yet. Choose a level to start
            tracking contributions.
          </p>
          <Dialog open={dialogOpen} onOpenChange={setDialogOpen} >
            <DialogTrigger asChild >
              <Button variant="ccAppOutlineMidBlue">Set My Goal</Button>
            </DialogTrigger>
            <DialogContent className="flex w-[20%] flex-col rounded-md bg-white p-4 text-black shadow-sm">
              <DialogHeader className="pb-3">
                <DialogTitle className=" text-lg">
                  Select Goal Level
                </DialogTitle>
              </DialogHeader>
              {!isSettingLevel ? (
                <div className="flex flex-col items-center gap-3 px-2">
                  {goalLevels.map(level => (
                    <Button
                      key={level.id}
                      variant="outline"
                      className="hover:bg-cc-app-blue bg-cc-app-mid-blue w-full rounded-lg px-5 py-2.5 text-white capitalize transition-colors duration-200 hover:cursor-pointer"
                      onClick={() => handleLevelSelect(level.level)}
                    >
                      {level.level}
                    </Button>
                  ))}
                </div>
              ) : (
                <div className="text-cc-app-light-blue flex items-center justify-center gap-2 py-5">
                  <Loader2 className="h-4 w-4 animate-spin" />
                  <span className="text-sm">Setting your goal...</span>
                </div>
              )}
              <DialogFooter className="pt-3">
                <Button
                  variant="ghost"
                  size="sm"
                  onClick={() => setDialogOpen(false)}
                >
                  Cancel
                </Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        </div>
      )}

      {/* Custom Goal Dialog */}
      <Dialog open={isCustomDialogOpen} onOpenChange={setIsCustomDialogOpen}>
        <DialogContent className="bg-white text-black">
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
                    className="flex justify-between rounded bg-gray-100 px-3 py-2"
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
              disabled={customGoals.length === 0}
              onClick={handleSubmitCustomGoals}
            >
              Save Custom Goal
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
