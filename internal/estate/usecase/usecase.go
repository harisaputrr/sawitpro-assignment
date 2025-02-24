package usecase

import (
	"context"
	"errors"
	"math"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/internal/estate/repository"
	treeRepository "github.com/SawitProRecruitment/UserService/internal/tree/repository"
	"github.com/google/uuid"
)

type Usecase struct {
	repository     repository.EstateRepository
	treeRepository treeRepository.TreeRepository
}

func NewUsecase(repository repository.EstateRepository, treeRepository treeRepository.TreeRepository) EstateUsecase {
	return &Usecase{
		repository:     repository,
		treeRepository: treeRepository,
	}
}

func (u *Usecase) GetEstateStats(ctx context.Context, estateID uuid.UUID) (result *generated.EstateStatsResponse, err error) {
	result, err = u.repository.GetEstateStats(ctx, estateID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *Usecase) GetEstateDronePlan(ctx context.Context, estateID uuid.UUID) (result *generated.GetEstateDronePlanResponse, err error) {
	estate, err := u.repository.GetEstateByID(ctx, estateID)
	if err != nil {
		return nil, errors.New("estate not found")
	}

	trees, _ := u.treeRepository.FindAllTrees(estateID)
	if len(trees) == 0 {
		return nil, errors.New("no trees found in the estate")
	}

	// Build a quick lookup for tree heights: (x,y) -> height
	treeHeights := make(map[[2]int]int)
	for _, t := range trees {
		treeHeights[[2]int{t.X, t.Y}] = t.Height
	}

	// Ascend to first plot's required altitude
	totalDistance := float32(0.0)
	currentAltitude := float32(0.0)
	firstHeight := treeHeights[[2]int{1, 1}]
	neededAltitude := float32(firstHeight + 1) // 1 if no tree
	totalDistance += neededAltitude - currentAltitude
	currentAltitude = neededAltitude

	// Calculate total distance using snake path
	direction := 1 // 1 means left-to-right, -1 means right-to-left
	for y := 1; y <= estate.Width; y++ {
		startX := 1
		endX := estate.Length
		if direction < 0 {
			startX = estate.Length
			endX = 1
		}

		// If we haven't reached the final plot in this row, move horizontally to the next plot
		for x := startX; x != (endX + direction); x += direction {
			if x+direction != endX+direction {
				totalDistance += 10 // 10m horizontally
				// Ascend/descend to next plot's altitude
				nextHeight := treeHeights[[2]int{x + direction, y}]
				nextAltitude := float32(nextHeight + 1)
				distDelta := float32(math.Abs(float64(nextAltitude - currentAltitude)))
				totalDistance += distDelta
				currentAltitude = nextAltitude
			}
		}

		// If not on the last row, move 1 plot north
		if y < estate.Width {
			totalDistance += 10 // 10m north
			// Ascend/descend to next plot's altitude
			nextHeight := treeHeights[[2]int{endX, y + 1}]
			nextAltitude := float32(nextHeight + 1)
			distDelta := float32(math.Abs(float64(nextAltitude - currentAltitude)))
			totalDistance += distDelta
			currentAltitude = nextAltitude
		}

		// Switch direction each row
		direction *= -1
	}

	totalDistance += currentAltitude
	currentAltitude = 0

	return &generated.GetEstateDronePlanResponse{
		Distance: &totalDistance,
	}, nil
}

func (u *Usecase) CreateEstate(ctx context.Context, payload generated.CreateEstateRequest) (result *generated.CreateEstateResponse, err error) {
	result, err = u.repository.CreateEstate(ctx, payload)
	if err != nil {
		return nil, err
	}

	return result, nil
}
