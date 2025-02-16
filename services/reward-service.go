package services

import (
	"go-api-poc/models"
	"log"
	"sync"

	"go-api-poc/db"
)

var (
	rewardsLock sync.Mutex
)

func GetRewards() *[]models.Reward {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	db := db.GetDbConnection()
	var rewards []models.Reward
	db.Find(&rewards)

	return &rewards
}

func CreateReward(brand *string, currency *string, denomination *float32) models.Reward {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	reward := models.Reward{
		Brand:        brand,
		Currency:     currency,
		Denomination: denomination,
	}

	db := db.GetDbConnection()
	result := db.Create(&reward)
	if result.Error != nil {
		log.Println("Error during creation: ", result.Error)
	}

	return reward
}

func GetReward(rewardId int) *models.Reward {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	db := db.GetDbConnection()
	var reward models.Reward
	result := db.First(&reward, rewardId)

	if result.RowsAffected == 0 {
		return nil
	}

	return &reward
}

func DeleteReward(rewardId int) error {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	db := db.GetDbConnection()
	db.Delete(&models.Reward{}, rewardId)

	return nil
}

func PutReward(rewardId int, brand *string, currency *string, denomination *float32) *models.Reward {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	db := db.GetDbConnection()
	var reward models.Reward
	result := db.First(&reward, rewardId)

	if result.RowsAffected == 0 {
		log.Printf("Reward with Id %v does not exist.", rewardId)
		return nil
	}

	if brand != nil {
		reward.Brand = brand
	}
	if currency != nil {
		reward.Currency = currency
	}
	if denomination != nil {
		reward.Denomination = denomination
	}

	db.Save(&reward)

	return &reward
}
