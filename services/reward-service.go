package services

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	rewardsDb   = make(map[int]Reward)
	rewardsLock sync.Mutex
)

type Reward struct {
	Brand        *string  `json:"brand,omitempty"`
	Currency     *string  `json:"currency,omitempty"`
	Denomination *float32 `json:"denomination,omitempty"`
	Id           *int     `json:"id,omitempty"`
}

func GetRewards() *[]Reward {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	rewardList := make([]Reward, 0, len(rewardsDb))
	for _, reward := range rewardsDb {
		rewardList = append(rewardList, reward)
	}

	return &rewardList
}

func CreateReward(brand *string, currency *string, denomination *float32) Reward {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	newId := rand.Intn(10000)
	reward := Reward{
		Brand:        brand,
		Currency:     currency,
		Denomination: denomination,
		Id:           &newId,
	}

	rewardsDb[newId] = reward

	return reward
}

func GetReward(rewardId int) *Reward {
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
