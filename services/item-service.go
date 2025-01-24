package services

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	items     = make(map[string]Item)
	itemsLock sync.Mutex
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	itemsLock.Lock()
	defer itemsLock.Unlock()

	itemList := make([]Item, 0, len(items))
	for _, item := range items {
		itemList = append(itemList, item)
	}

	log.Println("Returning items")
	itemListJSON, err := json.Marshal(itemList)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Println(string(itemListJSON))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemList)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	itemsLock.Lock()
	defer itemsLock.Unlock()

	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	items[item.ID] = item

	log.Println("Created item")
	log.Println(item)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func GetItem(w http.ResponseWriter, r *http.Request, id string) {
	itemsLock.Lock()
	defer itemsLock.Unlock()

	item, exists := items[id]
	if !exists {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	log.Println("Returning item")
	log.Println(item)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request, id string) {
	itemsLock.Lock()
	defer itemsLock.Unlock()

	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil || item.ID != id {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	items[id] = item

	log.Println("Updated item")
	log.Println(item)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, id string) {
	itemsLock.Lock()
	defer itemsLock.Unlock()

	if _, exists := items[id]; !exists {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	delete(items, id)

	log.Println("Deleted item")
	log.Println(id)

	w.WriteHeader(http.StatusNoContent)
}
