package users

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Firstname string `json:"firstname"`
    Lastname string `json:"lastname"`
    Email string `json:"email"`
    ProfilePictureUrl string `json:"profile_picture_url"`
    ScoreTechnical int `json:"score_technical"`
    ScoreLeadership int `json:"score_leadership"`
    ScoreTeamwork int `json:"score_teamwork"`
    ScoreOrganization int `json:"score_organization"`
}

type UserCredentialRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}



