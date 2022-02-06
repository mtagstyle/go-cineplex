package types

type TheatreOutput struct {
	Data             []TheatreData `json:"data"`
	Skip             int           `json:"skip"`
	Take             int           `json:"take"`
	TotalCount       int           `json:"totalCount"`
	Status           int           `json:"status"`
	Message          string        `json:"message"`
	MessageDetails   string        `json:"messageDetails"`
	APIVersionNumber string        `json:"apiVersionNumber"`
}

type TheatreData struct {
	ID                       int                 `json:"id"`
	Name                     string              `json:"name"`
	Address1                 string              `json:"address1"`
	Address2                 string              `json:"address2"`
	City                     string              `json:"city"`
	ProvinceCode             string              `json:"provinceCode"`
	PostalCode               string              `json:"postalCode"`
	NearestIntersection      string              `json:"nearestIntersection"`
	Latitude                 float64             `json:"latitude"`
	Longitude                float64             `json:"longitude"`
	Distance                 float64             `json:"distance"`
	URLSlug                  string              `json:"urlSlug"`
	Experiences              []TheatreExperience `json:"experiences"`
	IsTicketingAvailable     bool                `json:"isTicketingAvailable"`
	TheatreMessages          []TheatreMessage    `json:"theatreMessages"`
	IsDriveIn                bool                `json:"isDriveIn"`
	MapImageURL              string              `json:"mapImageUrl"`
	MobileMapImageURL        string              `json:"mobileMapImageUrl"`
	MobileBackgroundImageURL string              `json:"mobileBackgroundImageUrl"`
	CalorieChartURL          string              `json:"calorieChartUrl"`
	IsFavourite              bool                `json:"isFavourite"`
}

type TheatreExperience struct {
	ExperienceID string `json:"experienceId"`
	Title        string `json:"title"`
	ImageName    string `json:"imageName"`
	Description  string `json:"description"`
}

type TheatreMessage struct {
	ID       int    `json:"id"`
	Body     string `json:"body"`
	Category string `json:"category"`
}
