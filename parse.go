package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func makeLineReader(reader *bufio.Reader) func(string) string {
	return func(prompt string) string {
		fmt.Print(prompt)
		data, _ := reader.ReadString('\n')

		return data
	}
}

func parseReward(rewardStr string) ClientType {
	rewardStr = strings.Trim(rewardStr, "\n")
	rewardStr = strings.ToLower(rewardStr)

	if rewardStr == "true" {
		return RewardClientType
	}

	if rewardStr == "false" {
		return DefaultClientType
	}

	log.Fatal(rewardStr)

	return ""
}
