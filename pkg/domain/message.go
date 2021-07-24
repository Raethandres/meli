package domain

import "math"

func GetMessage(messages ...[]string) (msg string) {
	totalArrayItems := 0
	
	for _, m := range messages {
		totalArrayItems += len(m)
	}
	
	itemsPerArray := float64(totalArrayItems) / float64(len(messages))
	isAsymmetricArray := math.Mod(itemsPerArray, 1.0) != 0

	if isAsymmetricArray {
		messages, itemsPerArray = completeAsymmetricArray(messages)
	}

	return decodeMessage(messages, itemsPerArray)
}

func decodeMessage(messages [][]string, itemsPerArray float64) string {
	var decodedMessage string = ""

	for i := 0; i < int(itemsPerArray); i++ {
		var word string = ""

		for j := 0; j < len(messages); j++ {
			if word == "" && messages[j][i] != "" {
				word = messages[j][i]
			}
		}

		decodedMessage += spaceOnNonBlank(decodedMessage) + word
	}

	return decodedMessage
}

func completeAsymmetricArray(messages [][]string)  (msg [][]string, qty float64) {
	biggestArraySize := 0

	for _, m := range messages {
		if len(m) > biggestArraySize  {
			biggestArraySize = len(m)
		}
	}

	var balancedMessages [][]string 

	for _, message := range messages {
		if len(message) != biggestArraySize  {
			diff := biggestArraySize - len(message)
			words := make([]string, diff)
			words = append(words, message...)
			
			balancedMessages = append(balancedMessages, words)
		} else {
			balancedMessages = append(balancedMessages, message)
		}
	}

	return balancedMessages, float64(biggestArraySize)
}

func spaceOnNonBlank(msj string) string {
	space := " "
	
	if msj == "" {
		space = ""
	}

	return space
}