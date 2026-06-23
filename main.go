package main

import (
	"fmt"
 main
	"time"
)

func main() {
	fmt.Println("Quiz starting...")
	time.Sleep(2 * time.Second)
	fmt.Println("Goodbye")
=======
	"bufio"
	"os"
	"strings"
)

type Question struct {
	Text string
	Answer string
}

func main() {
	fmt.Println("Which album was the song from?")
	quiz := []Question{
		{Text: "Altogether ooky", Answer: "single"},
		{Text: "Dimes", Answer: "Hard-off"},
		{Text: "Fire Water Burn", Answer: "One Fierce Beer Coaster"},
		{Text: "Mope", Answer: "Hooray for Boobies"},
		{Text: "Foxtrot Uniform Charlie Kilo", Answer: "Hefty Fine"},
	}

	score := 0
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Start the pointlessness!!!\n\n", len(quiz))

	for i, q := range quiz {
		fmt.Println("Question %d: %s\n", i+1, q.Text)
		fmt.Println("Answer: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Skipping question.")
			continue
		}

		cleanedInput := strings.TrimSpace(input)
		if strings.EqualFold(cleanedInput, q.Answer) {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Printf("Wrong!!! The correct answer was: %s\n", q.Answer)
		}
       }
       fmt.Printf("Quiz is over. Final Score: %d out of %d\n", score, len(quiz))
       percent := (float64(score) / float64(len(quiz))) * 100
       fmt.Printf("Success Rate: %.1f%%\n", percent)
 dev
}
