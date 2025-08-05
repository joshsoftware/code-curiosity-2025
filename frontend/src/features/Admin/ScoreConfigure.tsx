import { useState } from "react";

import { Button } from "@/shared/components/ui/button";
import { Input } from "@/shared/components/ui/input";
import { Card } from "@/shared/components/ui/card";
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

  const [scores, setScores] = useState<Record<string, number>>({});

  if (isLoading) return <div>Loading...</div>;
  if (isError || !data?.data)
    return <div>Failed to load contribution types.</div>;

  const handleScoreChange = (contributionType: string, newScore: number) => {
    setScores(prev => ({ ...prev, [contributionType]: newScore }));
  };

  const handleSaveAll = () => {
    const updates: ContributionScoreUpdate[] = Object.entries(scores).map(
      ([contributionType, score]) => ({
        contributionType,
        score
      })
    );

    if (updates.length === 0) {
      toast.info("No changes to save.");
      return;
    }

    configureScore(updates, {
      onSuccess: () => {
        toast.success("Contribution scores updated successfully.");
        setScores({});
      },
      onError: () => {
        toast.error("Failed to update contribution scores.");
        console.log(updates);
      }
    });
  };

  return (
    <Card className="mx-auto max-w-4xl p-6">
      <div className="mb-6">
        <h1 className="mb-2 text-2xl font-bold text-gray-900">
          Configure Contribution Scores
        </h1>
      </div>
      <div className="space-y-4">
        {data.data.map((item: ContributionScore) => (
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
              defaultValue={item.score}
              onChange={e =>
                handleScoreChange(item.contributionType, Number(e.target.value))
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
    </Card>
  );
};

export default ScoreConfigure;
