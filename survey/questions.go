package questions

import (
	"fmt"
	"math/rand"
	//"time"
)

type Question struct {
	Text        string
	AnswerType  string
	Options     []Answer
}

type Answer struct {
	Text string
}

func generateOptions(numOptions int) []Answer {
	//create a slice called options with type Answer with a capacity numOptions
	options := make([]Answer, numOptions)
	for i := 0; i < numOptions; i++ {
		options[i] = Answer{Text: fmt.Sprintf("Option %c", 'A'+i)}
	}
	return options
}

func generateQuestion(questionText string, answerType string, numOptions int) Question {
	// Generate options based on answer type
	var options []Answer
	if answerType == "multiple-choice" {
		options = generateOptions(numOptions)
	}

	// Create and return a Question object
	return Question{
		Text:       questionText,
		//AnswerType: answerType,
		Options:    options,
	}
}

func acceptsQuestionTexts(questionText string)string{
	return questionText
}

func acceptsAnswerType(answerType string, numOptions ...int)string{
	if answerType == "multiple-choice" {
		return "multiple-choice"
	}else {
		return ""
	}
}

func generateQuestionsBasedOnNumber(numQuestions int, numOptions int,answerType string, questionText string)[]Question{
	var questions []Question
	for i := 0; i < numQuestions; i++ {

		// For multiple-choice, prompt the user for the number of options
		numOptions = 0
		

		// Generate the question based on user input
		question := generateQuestion(acceptsQuestionTexts(questionText), acceptsAnswerType(answerType), numOptions)

		// Add the question to the list
		questions = append(questions, question)
	}
	return questions
}

func question() {
	var questions []Question
	// Seed the random number generator
	//rand.Seed(time.Now().UnixNano())

	var numQuestions int

	// Prompt the user for the number of questions
	fmt.Print("Enter the number of questions: ")
	fmt.Scanln(&numQuestions)

	// Display generated questions
	fmt.Println("\nGenerated Questions:")
	for i, question := range questions {
		fmt.Printf("%d. Question: %s\n", i+1, question.Text)
		fmt.Printf("   Answer Type: %s\n", question.AnswerType)
		if len(question.Options) > 0 {
			fmt.Println("   Options:")
			for _, option := range question.Options {
				fmt.Printf("      %s\n", option.Text)
			}
		}
		fmt.Println()
	}
}