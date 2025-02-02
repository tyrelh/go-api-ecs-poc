package services

import (
	"fmt"
	"go-api-poc/models"
	"math/rand"
	"sync"
)

var (
	rewardsDb   = make(map[int]models.Reward)
	rewardsLock sync.Mutex
)

func GetRewards() *[]models.Reward {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	rewardList := make([]models.Reward, 0, len(rewardsDb))
	for _, reward := range rewardsDb {
		rewardList = append(rewardList, reward)
	}

	return &rewardList
}

func CreateReward(brand *string, currency *string, denomination *float32) models.Reward {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	newId := rand.Intn(10000)
	reward := models.Reward{
		Brand:        brand,
		Currency:     currency,
		Denomination: denomination,
		Id:           &newId,
	}

	rewardsDb[newId] = reward

	return reward
}

func GetReward(rewardId int) *models.Reward {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	reward, ok := rewardsDb[rewardId]
	if !ok {
		return nil
	}

	return &reward
}

func DeleteReward(rewardId int) error {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	_, ok := rewardsDb[rewardId]
	if !ok {
		return fmt.Errorf("reward with id %d not found", rewardId)
	}

	delete(rewardsDb, rewardId)

	return nil
}
