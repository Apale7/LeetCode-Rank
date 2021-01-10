package model

type QuestionLevelInfo struct {
	Data QuestionLevelData `json:"data"`
}

type QuestionLevel struct {
	Difficulty string `json:"difficulty"`
}
type QuestionLevelData struct {
	Question QuestionLevel `json:"question"`
}

type Info struct {
	Data Data `json:"data"`
}
type Question struct {
	QuestionFrontendID string `json:"questionFrontendId"`
	Title              string `json:"title"`
	TranslatedTitle    string `json:"translatedTitle"`
	TitleSlug          string `json:"titleSlug"`
	Typename           string `json:"__typename"`
}
type RecentSubmissions struct {
	Status     string   `json:"status"`
	Lang       string   `json:"lang"`
	Question   Question `json:"question"`
	SubmitTime int64    `json:"submitTime"`
	Typename   string   `json:"__typename"`
}
type Data struct {
	RecentSubmissions []RecentSubmissions `json:"recentSubmissions"`
}
