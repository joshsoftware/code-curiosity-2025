import { useState } from "react";

import { Button } from "@/shared/components/ui/button";
import { Input } from "@/shared/components/ui/input";
import {
  useFetchContributionTypes,
  useConfigureContributionScore
} from "@/api/queries/Admin";
import type {
  ContributionScore,
  ContributionScoreUpdate
} from "@/shared/types/types";
import { toast } from "sonner";

const ScoreConfigure = () => {
  const { data, isLoading, isError } = useFetchContributionTypes();
  const { mutate: configureScore, isPending } = useConfigureContributionScore();

  const [scores, setScores] = useState<Record<string, string>>({});
  const [version, setVersion] = useState(0); // force re-render after save

  if (isLoading) return <div>Loading...</div>;
  if (isError || !data?.data)
    return <div>Failed to load contribution types.</div>;

  // Initialize state with current scores (so they are controlled)
  if (Object.keys(scores).length === 0) {
    const init: Record<string, string> = {};
    data.data.forEach((item: ContributionScore) => {
      init[item.contributionType] = String(item.score);
    });
    setScores(init);
  }

  const handleScoreChange = (contributionType: string, newScore: string) => {
    setScores(prev => ({ ...prev, [contributionType]: newScore }));
  };

  const handleSaveAll = () => {
    const invalid = Object.entries(scores).some(
      ([, score]) => score === "" || isNaN(Number(score))
    );
    if (invalid) {
      toast.error("All scores must be valid numbers (no empty fields).");
      return;
    }

    const updates: ContributionScoreUpdate[] = Object.entries(scores).map(
      ([contributionType, score]) => ({
        contributionType,
        score: Number(score)
      })
    );

    configureScore(updates, {
      onSuccess: () => {
        toast.success("Contribution scores updated successfully.");
        // trigger re-render to update sorted order
        setVersion(prev => prev + 1);
      },
      onError: () => {
        toast.error("Failed to update contribution scores.");
        console.log(updates);
      }
    });
  };

  // Sort data based on current scores
  const sortedData = [...data.data].sort(
    (a, b) =>
      Number(scores[a.contributionType]) - Number(scores[b.contributionType])
  );

  return (
    <div className="mx-auto max-w-4xl p-6">
      <div className="mb-6">
        <h1 className="mb-2 text-2xl font-bold text-gray-900">
          Configure Contribution Scores
        </h1>
      </div>

      <div className="space-y-4">
        {sortedData.map((item: ContributionScore) => (
          <div
            key={item.id}
            className="flex items-center justify-between border-b pb-2"
          >
            <div className="flex-1 font-medium">
              {item.contributionType.replace(/([a-z])([A-Z])/g, "$1 $2")}
            </div>

            <Input
              type="number"
              className="mr-4 w-24"
              min={0}
              required
              value={scores[item.contributionType] ?? ""}
              onChange={e =>
                handleScoreChange(item.contributionType, e.target.value)
              }
            />
          </div>
        ))}
      </div>

      <div className="mt-6 text-right">
        <Button onClick={handleSaveAll} disabled={isPending}>
          {isPending ? "Saving..." : "Save All"}
        </Button>
      </div>
    </div>
  );
};

export default ScoreConfigure;
