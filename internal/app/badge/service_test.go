package badge

import (
	"context"
	"testing"
	"time"

	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func fromListOfRepositoryBadges(repoBadges []repository.Badge) []Badge {
	finalBadges := make([]Badge, len(repoBadges))

	for i, badge := range repoBadges {
		finalBadges[i] = FromRepositoryBadge(badge)
	}

	return finalBadges
}

func TestGetBadgeDetailsOfUser(t *testing.T) {
	mockRepo := &mocks.BadgeRepository{}
	var testBadges []repository.Badge = []repository.Badge{
		{
			BadgeId:   1,
			BadgeType: "Beginner",
			EarnedAt:  time.Now().Add(-48 * time.Hour),
		},
		{
			BadgeId:   2,
			BadgeType: "Intermediate",
			EarnedAt:  time.Now().Add(-24 * time.Hour),
		},
	}
	emptyBadges := []repository.Badge{}

	mockRepo.
	On("GetBadgeDetailsOfUser", mock.Anything, mock.Anything, 1).
	Return(testBadges, nil)

	mockRepo.
	On("GetBadgeDetailsOfUser", mock.Anything, mock.Anything, -1).
	Return(emptyBadges, nil)

	badgeService := &badgeService{
		badgeRepository: mockRepo,
	}
	
	userBadges, err := badgeService.GetBadgeDetailsOfUser(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, fromListOfRepositoryBadges(testBadges), userBadges)

	userBadges, err = badgeService.GetBadgeDetailsOfUser(context.Background(), -1)
	assert.NoError(t, err)
	assert.Equal(t, fromListOfRepositoryBadges(emptyBadges), userBadges)
}