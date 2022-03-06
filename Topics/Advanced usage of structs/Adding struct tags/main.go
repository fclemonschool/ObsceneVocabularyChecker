package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type bankAccount struct {
	// add the struct tags below!
	AccountId   int     `json:"accountId"`
	AccountType string  `json:"accountType"`
	Balance     float64 `json:"balance"`
}

func hi() {
	fmt.Sprint(json.Number(1))
	log.Print()
}
