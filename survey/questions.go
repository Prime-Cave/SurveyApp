package questions

import (
)

type Survey struct{
	Question 
	Options 
}

type Question struct {
	Text       string
}


type Options struct {
	Text string
	AnswerType string
	Options    []string
}



func CreateQuestionnaire(questionnaire ...string) map[int]string {
	questions := make(map[int]string, 0)

	for i, question := range questionnaire{
		questions[i+1] = question
	}

	return questions
}

func question() {

}
