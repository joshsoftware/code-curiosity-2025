import { useUserBadges } from "@/api/queries/UserBadges";
import { Star } from "lucide-react";
import type { Badge } from "@/shared/types/types";

const badgeColorMap: Record<string, string> = {
  BEGINNER: "text-[#cd7f32]",
  INTERMEDIATE: "text-[#c0c0c0]",
  ADVANCED: "text-[#ffd700]",
  CUSTOM: "text-orange-400",
};

const UserBadges = () => {
  const { data } = useUserBadges();
  const badges = data?.data ?? [];

  const grouped = badges.reduce<Record<string, Badge[]>>((acc, badge) => {
    const type = badge.badgeType.toUpperCase();
    if (!acc[type]) acc[type] = [];
    acc[type].push(badge);
    return acc;
  }, {});

  return (
    <div>
      <p className="text-cc-app-light-blue mb-4 text-left font-semibold">
        BADGES
      </p>
      <div className="flex flex-wrap gap-4">
        {Object.entries(grouped).map(([type, badgeList]) => {
          const color = badgeColorMap[type] ?? "text-gray-400";
          return (
            <div key={type} className="relative group flex flex-col items-center">
              <Star
                className={`${color} h-6 w-6 fill-current`}
              />
              {badgeList.length > 1 && (
                <span className="text-xs mt-1">×{badgeList.length}</span>
              )}
              <div className="absolute bottom-8 z-10 hidden w-max rounded bg-gray-600 px-2 py-1 text-xs text-white group-hover:block">
                {type} <br />
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default UserBadges;
