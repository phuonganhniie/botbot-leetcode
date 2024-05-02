package model

type Challenge struct {
	Id           string `json:"questionFrontendId"`
	Name         string `json:"questionTitle"`
	QuestionLink string `json:"questionLink"`
	Difficulty   string `json:"difficulty"`
}
