package users

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Firstname string `json:"firstname"`
    Lastname string `json:"lastname"`
    ProfilePictureUrl string `json:"profile_picture_url"`
    Score_1 int `json:"score_1"`
    Score_2 int `json:"score_2"`
    Score_3 int `json:"score_3"`
    Score_4 int `json:"score_4"`

}



