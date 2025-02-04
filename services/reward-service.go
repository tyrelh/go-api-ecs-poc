package services

import (
	"go-api-poc/models"
	"log"
	"sync"
	"time"

	"go-api-poc/db"
)

var (
	rewardsLock sync.Mutex
)

func GetRewards() *[]models.Reward {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	startTime := time.Now()
	db := db.GetDbConnection()
	var rewards []models.Reward
	db.Find(&rewards)
	log.Printf("Time taken to get rewards: %v", time.Since(startTime))

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

	startTime := time.Now()
	db := db.GetDbConnection()
	result := db.Create(&reward)
	log.Println("Error during creation: ", result.Error)
	log.Println("Rows affected: ", result.RowsAffected)
	log.Printf("Time taken to create reward: %v", time.Since(startTime))

	return reward
}

func GetReward(rewardId int) *models.Reward {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	startTime := time.Now()
	db := db.GetDbConnection()
	var reward models.Reward
	result := db.First(&reward, rewardId)
	log.Printf("Time taken to get reward: %v", time.Since(startTime))

	if result.RowsAffected == 0 {
		return nil
	}

	return &reward
}

func DeleteReward(rewardId int) error {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	startTime := time.Now()
	db := db.GetDbConnection()
	db.Delete(&models.Reward{}, rewardId)
	log.Printf("Time taken to delete reward: %v", time.Since(startTime))

	return nil
}

func PutReward(rewardId int, brand *string, currency *string, denomination *float32) *models.Reward {
	rewardsLock.Lock()
	defer rewardsLock.Unlock()

	startTime := time.Now()
	db := db.GetDbConnection()
	var reward models.Reward
	result := db.First(&reward, rewardId)
	log.Printf("Time taken to get reward: %v", time.Since(startTime))

	if result.RowsAffected == 0 {
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

	startTime = time.Now()
	db.Save(&reward)
	log.Printf("Time taken to update reward: %v", time.Since(startTime))

	return &reward
}
