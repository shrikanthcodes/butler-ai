package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	model "backend/pkg/models"
	llm "backend/pkg/services/ai/llm"
	//config "backend/internal/config"
	//"log"
)

func main() {
	//llm.SimpleTest()
	reader := bufio.NewReader(os.Stdin)

	ctx := context.Background()
	gs := llm.InitializeGeminiService()

	// Start a new chat session with a custom system prompt
	gs.StartNewChat("You are an expert in giving ideas for meals. You are an expert at explaining things like you are talking to a 5 year old (ELI5) (Answer in a single string)", []model.Dialogue{})

	fmt.Println("Chat session started, you are now talking to a go expert")

	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	// First user message
	response, err := gs.PredictChat(ctx, text+"(Answer in 50 words or less)", 150, 0.7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Assistant:", response)

	fmt.Print("Enter text: ")
	text, _ = reader.ReadString('\n')
	fmt.Println(text)

	// Next user message continues the conversation
	response, err = gs.PredictChat(ctx, text+"(Answer in 50 words or less)", 150, 0.7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Assistant:", response)

	// When you want to start over, explicitly start a new chat
	// gs.StartNewChat("You are a helpful assistant.", []model.Dialogue{})

	// Next user message continues the conversation
	// response, err = gs.Predict(ctx, "In 50 words, describe Indianapolis", 50, 0.5)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Assistant:", response)

	defer gs.Close()

}
