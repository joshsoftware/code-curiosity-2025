package repository

import (
	"context"
	"testing"
	"time"

	"github.com/joshsoftware/code-curiosity-2025/internal/repository/base"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository/base/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBadgeDetailsOfUser(t *testing.T) {
	mockDb := &mocks.QueryExecuter{}
	var testBadges []Badge = []Badge{
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

	emptyBadges := []Badge{}

	mockDb.
		On("SelectContext", mock.Anything, mock.AnythingOfType("*[]repository.Badge"), getBadgeDetailsOfUserQuery, 1).
		Run(func(args mock.Arguments) {
			arg := args.Get(1).(*[]Badge)
			*arg = testBadges
		}).
		Return(nil)

	mockDb.
		On("SelectContext", mock.Anything, mock.AnythingOfType("*[]repository.Badge"), getBadgeDetailsOfUserQuery, -1).
		Run(func(args mock.Arguments) {
			arg := args.Get(1).(*[]Badge)
			*arg = emptyBadges
		}).
		Return(nil)

	// if we pass other type's value to user id like "abcd"

	// our query is ok but what if 1 of the columns in the query is dropped from the table

	badgeRepo := &badgeRepository{
		BaseRepository: base.NewBaseRepository(mockDb),
	}

	badgesOfUser, err := badgeRepo.GetBadgeDetailsOfUser(context.Background(), nil, 1)
	assert.NoError(t, err)
	assert.Equal(t, testBadges, badgesOfUser)

	badgesOfUser, err = badgeRepo.GetBadgeDetailsOfUser(context.Background(), nil, -1)
	assert.NoError(t, err)
	assert.Equal(t, emptyBadges, badgesOfUser)

	mockDb.AssertExpectations(t)
}
