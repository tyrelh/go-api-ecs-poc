package controllers

import (
	"context"

	"go-api-poc/api"
	"go-api-poc/services"
)

func (StrictServer) GetGoReward(ctx context.Context, request api.GetGoRewardRequestObject) (api.GetGoRewardResponseObject, error) {
	rewards := services.GetRewards()
	return api.GetGoReward200JSONResponse{
		Rewards: rewards,
	}, nil
}

func (StrictServer) PostGoReward(ctx context.Context, request api.PostGoRewardRequestObject) (api.PostGoRewardResponseObject, error) {
	rewardRequestBody := *request.Body
	reward := services.CreateReward(rewardRequestBody)
	return api.PostGoReward201JSONResponse(reward), nil
}
