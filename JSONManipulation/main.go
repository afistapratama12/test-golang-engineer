package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Inventory struct {
	InventoryID int       `json:"inventory_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Tags        []string  `json:"tags"`
	PurchasedAt int       `json:"purchased_at"`
	Placement   Placement `json:"placement"`
}

type Placement struct {
	RoomID int    `json:"room_id"`
	Name   string `json:"name"`
}

func main() {
	// get json file
	jsonFile, err := os.Open("data.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println("success open inventory-list.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var Inventories []Inventory

	json.Unmarshal(byteValue, &Inventories)

	// manipulation number 1
	fmt.Println("======================================")
	for _, inventory := range Inventories {
		if inventory.Placement.Name == "Meeting Room" {
			fmt.Println(inventory.Name)
		}
	}

	// manipulation number 2
	fmt.Println("======================================")
	for _, inventory := range Inventories {
		if inventory.Type == "electronic" {
			fmt.Println(inventory.Name)
		}
	}

	// manipulation number 3
	fmt.Println("======================================")
	for _, inventory := range Inventories {
		if inventory.Type == "furniture" {
			fmt.Println(inventory.Name)
		}
	}

	// manipulation number 4
	fmt.Println("======================================")
	for _, inventory := range Inventories {
		strDate := strconv.Itoa(inventory.PurchasedAt)
		i, _ := strconv.ParseInt(strDate, 10, 64)
		tm := time.Unix(i, 0)
		if tm.Day() == 16 && int(tm.Month()) == 1 && tm.Year() == 2020 {
			fmt.Println(inventory.Name)
		}
	}

	// manipulation number 5
	fmt.Println("======================================")
	for _, inventory := range Inventories {
		for _, tagsInv := range inventory.Tags {
			if tagsInv == "brown" {
				fmt.Println(inventory.Name)
			}
		}
	}
}
