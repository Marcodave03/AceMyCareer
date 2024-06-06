package interviews

type interviewLevel struct {
    ID int `json:"id"`
    Name string `json:"name"`
}

type interviewPosition struct {
    Name string `json:"name"`
}

type interviewTag struct {
    Name string `json:"name"`
}


type interviewIndustry struct {
    Name string `json:"name"`
}

type interviewRequirements struct {
    Name string `json:"name"`
}

type interview struct {
    ID int `json:"id"`
    Title string `json:"title"`
    ImageURL string `json:"image_url"`
    Industry string `json:"industry"`
    Position string `json:"position"`
    Level string `json:"level"`
    Experience string `json:"experience"`
    Likes int `json:"likes"`
}

type interviewRequest struct {
    ID int `json:"id"`
    Title string `json:"title"`
    ImageURL string `json:"image_url"`
    Industry string `json:"industry"`
    Position string `json:"position"`
    Level string `json:"level"`
    Experience string `json:"experience"`
    Likes int `json:"likes"`
    Tags []interviewTag `json:"tags"`
    Requirements []interviewRequirements `json:"requirements"`
}

