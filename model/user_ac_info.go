package model

type UserSlug struct {
	UserSlug string `json:"userSlug"`
}

type PostData struct {
	OprationName string   `json:"oprationName"`
	Variables    UserSlug `json:"variables"`
	Query        string   `json:"query"`
}

type AcData struct {
	UserAcInfo UserAcInfo `json:"data"`
}
type SubmissionProgress struct {
	TotalSubmissions int `json:"totalSubmissions"`
	WaSubmissions int `json:"waSubmissions"`
	AcSubmissions int `json:"acSubmissions"`
	ReSubmissions int `json:"reSubmissions"`
	OtherSubmissions int `json:"otherSubmissions"`
	AcTotal int `json:"acTotal"`
	QuestionTotal int `json:"questionTotal"`
	Typename string `json:"__typename"`
}
type UserProfilePublicProfile struct {
	Username string `json:"username"`
	SubmissionProgress SubmissionProgress `json:"submissionProgress"`
	Typename string `json:"__typename"`
}
type UserAcInfo struct {
	UserProfilePublicProfile UserProfilePublicProfile `json:"userProfilePublicProfile"`
}