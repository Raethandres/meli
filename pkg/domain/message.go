package domain

import (
	"errors"
	"strings"
)

func GetMessage(messages ...[]string) (msg string, err error) {

	length, err := verifySymmetric(messages)
	var message string = ""
	if err != nil {
		messages, length = completeAsymmetricArray(messages)
	}
	for i := 0; i < length; i++ {
		res, err := getSecret(messages, 0, i, "")
		if err != nil {
			return "", err
		} else {
			message += res
			message += " "
		}
	}

	return strings.TrimSpace(message), nil
}

func getSecret(messages [][]string, i int, j int, m string) (string, error) {
	if i == len(messages) {
		return m, nil
	}
	if m != "" && messages[i][j] != "" && m != messages[i][j] {
		return "", errors.New("not found")
	}
	if messages[i][j] == "" {
		return getSecret(messages, i+1, j, m)
	}
	return getSecret(messages, i+1, j, messages[i][j])
}

func verifySymmetric(messages [][]string) (int, error) {
	if len(messages[0]) == 0 {
		return 0, errors.New("not found")
	}
	for i := 0; i < len(messages)-1; i++ {
		if len(messages[i]) != len(messages[i+1]) {
			return 0, errors.New("not found")
		}
	}
	return len(messages[0]), nil
}

func completeAsymmetricArray(messages [][]string) (msg [][]string, qty int) {
	biggestArraySize := 0

	for _, m := range messages {
		if len(m) > biggestArraySize {
			biggestArraySize = len(m)
		}
	}

	var balancedMessages [][]string

	for _, message := range messages {
		if len(message) != biggestArraySize {
			diff := biggestArraySize - len(message)
			words := make([]string, diff)
			words = append(words, message...)
			balancedMessages = append(balancedMessages, words)
		} else {
			balancedMessages = append(balancedMessages, message)
		}
	}

	return balancedMessages, biggestArraySize
}
