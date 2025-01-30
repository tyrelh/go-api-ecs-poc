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

func (StrictServer) GetGoRewardId(ctx context.Context, request api.GetGoRewardIdRequestObject) (api.GetGoRewardIdResponseObject, error) {
	rewardId := request.Id
	reward := services.GetReward(rewardId)
	return api.GetGoRewardId200JSONResponse(*reward), nil
}

func (StrictServer) DeleteGoRewardId(ctx context.Context, request api.DeleteGoRewardIdRequestObject) (api.DeleteGoRewardIdResponseObject, error) {
	rewardId := request.Id
	err := services.DeleteReward(rewardId)
	if err != nil {
		return api.DeleteGoRewardId404Response{}, nil
	}
	return api.DeleteGoRewardId204Response{}, nil
}
