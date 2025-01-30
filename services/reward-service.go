package services

import (
	"go-api-poc/api"
	"math/rand"
	"sync"
)

var (
	rewardsDb   = make(map[int]api.Reward)
	rewardsLock sync.Mutex
)

func GetRewards() *[]api.Reward {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	rewardList := make([]api.Reward, 0, len(rewardsDb))
	for _, reward := range rewardsDb {
		rewardList = append(rewardList, reward)
	}

	return &rewardList
}

func CreateReward(rewardCreation api.RewardCreation) api.Reward {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	newId := rand.Intn(10000)
	reward := api.Reward{
		Brand:        rewardCreation.Brand,
		Currency:     rewardCreation.Currency,
		Denomination: rewardCreation.Denomination,
		Id:           &newId,
	}

	rewardsDb[newId] = reward

	return reward
}

func GetReward(rewardId int) *api.Reward {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	reward, ok := rewardsDb[rewardId]
	if !ok {
		return nil
	}

	return &reward
}
