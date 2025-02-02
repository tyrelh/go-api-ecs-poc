package controllers

import (
	"context"

	"go-api-poc/api"
	"go-api-poc/services"
)

func (StrictServer) GetGoReward(ctx context.Context, request api.GetGoRewardRequestObject) (api.GetGoRewardResponseObject, error) {
	rewards := services.GetRewards()

	var rewardResponses []api.RewardResponse
	for _, reward := range *rewards {
		id := int(reward.ID)
		rewardResponses = append(rewardResponses, api.RewardResponse{
			Brand:        reward.Brand,
			Currency:     reward.Currency,
			Denomination: reward.Denomination,
			Id:           &id,
		})
	}

	return api.GetGoReward200JSONResponse{
		Rewards: &rewardResponses,
	}, nil
}

func (StrictServer) PostGoReward(ctx context.Context, request api.PostGoRewardRequestObject) (api.PostGoRewardResponseObject, error) {
	rewardRequestBody := *request.Body

	reward := services.CreateReward(
		rewardRequestBody.Brand,
		rewardRequestBody.Currency,
		rewardRequestBody.Denomination,
	)
	id := int(reward.ID)
	rewardResponse := api.RewardResponse{
		Brand:        reward.Brand,
		Currency:     reward.Currency,
		Denomination: reward.Denomination,
		Id:           &id,
	}
	return api.PostGoReward201JSONResponse(rewardResponse), nil
}

func (StrictServer) GetGoRewardId(ctx context.Context, request api.GetGoRewardIdRequestObject) (api.GetGoRewardIdResponseObject, error) {
	rewardId := request.Id
	reward := services.GetReward(rewardId)
	id := int(reward.ID)
	rewardResponse := api.RewardResponse{
		Brand:        reward.Brand,
		Currency:     reward.Currency,
		Denomination: reward.Denomination,
		Id:           &id,
	}
	return api.GetGoRewardId200JSONResponse(rewardResponse), nil
}

func (StrictServer) DeleteGoRewardId(ctx context.Context, request api.DeleteGoRewardIdRequestObject) (api.DeleteGoRewardIdResponseObject, error) {
	rewardId := request.Id
	err := services.DeleteReward(rewardId)
	if err != nil {
		return api.DeleteGoRewardId404Response{}, nil
	}
	return api.DeleteGoRewardId204Response{}, nil
}
