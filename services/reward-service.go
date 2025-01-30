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

	// newIdInt := rand.Intn(10000)
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

// type Item struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// }

// var (
// 	items     = make(map[string]Item)
// 	itemsLock sync.Mutex
// )

// func GetItems(w http.ResponseWriter, r *http.Request) {
// 	itemsLock.Lock()
// 	defer itemsLock.Unlock()

// 	itemList := make([]Item, 0, len(items))
// 	for _, item := range items {
// 		itemList = append(itemList, item)
// 	}

// 	log.Println("Returning items")
// 	itemListJSON, err := json.Marshal(itemList)
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		return
// 	}
// 	log.Println(string(itemListJSON))

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(itemList)
// }

// func CreateItem(w http.ResponseWriter, r *http.Request) {
// 	itemsLock.Lock()
// 	defer itemsLock.Unlock()

// 	var item Item
// 	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}

// 	items[item.ID] = item

// 	log.Println("Created item")
// 	log.Println(item)

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(item)
// }

// func GetItem(w http.ResponseWriter, r *http.Request, id string) {
// 	itemsLock.Lock()
// 	defer itemsLock.Unlock()

// 	item, exists := items[id]
// 	if !exists {
// 		http.Error(w, "Not found", http.StatusNotFound)
// 		return
// 	}

// 	log.Println("Returning item")
// 	log.Println(item)

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(item)
// }

// func UpdateItem(w http.ResponseWriter, r *http.Request, id string) {
// 	itemsLock.Lock()
// 	defer itemsLock.Unlock()

// 	var item Item
// 	if err := json.NewDecoder(r.Body).Decode(&item); err != nil || item.ID != id {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}

// 	items[id] = item

// 	log.Println("Updated item")
// 	log.Println(item)

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(item)
// }

// func DeleteItem(w http.ResponseWriter, id string) {
// 	itemsLock.Lock()
// 	defer itemsLock.Unlock()

// 	if _, exists := items[id]; !exists {
// 		http.Error(w, "Not found", http.StatusNotFound)
// 		return
// 	}

// 	delete(items, id)

// 	log.Println("Deleted item")
// 	log.Println(id)

// 	w.WriteHeader(http.StatusNoContent)
// }
