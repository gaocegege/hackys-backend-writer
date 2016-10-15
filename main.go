// test project main.go
package main

import (
	"github.com/gaocegege/hackys-backend-writer/cognitive"
	"github.com/gaocegege/hackys-backend-writer/pkg/log"
)

func main() {
	result, err := cognitive.RecognizeText("gaocegege is awesome.")
	if err != nil {
		panic(err)
	}
	log.Infof("Result: %v", result)
	// cognitive.RecognizeEmotion("https://portalstoragewuprod.azureedge.net/emotion/recognition1.jpg")
}
