package services

import (
	"fmt"
	"go-api-poc/api"
	"math/rand"
	"sync"
)

var (
	rewardsDb   = make(map[string]api.Reward)
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
	newIdFloat := float32(newId)
	newIdStr := fmt.Sprintf("%d", newId)
	reward := api.Reward{
		Brand:        rewardCreation.Brand,
		Currency:     rewardCreation.Currency,
		Denomination: rewardCreation.Denomination,
		Id:           &newIdFloat,
	}

	rewardsDb[newIdStr] = reward

	return reward
}
