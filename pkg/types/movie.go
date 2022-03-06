package types

type MoviesData struct {
	IsWatchListed            bool   `json:"isWatchListed"`
	ID                       int    `json:"id"`
	PresentationType         string `json:"presentationType"`
	Name                     string `json:"name"`
	ReleaseDate              string `json:"releaseDate"`
	FirstShowStartDate       string `json:"firstShowStartDate"`
	MarketLanguageCode       string `json:"marketLanguageCode"`
	IsNowPlaying             bool   `json:"isNowPlaying"`
	IsComingSoon             bool   `json:"isComingSoon"`
	IsReleasingSoon          bool   `json:"isReleasingSoon"`
	IsEvent                  bool   `json:"isEvent"`
	HasShowtimes             bool   `json:"hasShowtimes"`
	HasPosterImage           bool   `json:"hasPosterImage"`
	SmallPosterImageURL      string `json:"smallPosterImageUrl"`
	MediumPosterImageURL     string `json:"mediumPosterImageUrl"`
	LargePosterImageURL      string `json:"largePosterImageUrl"`
	MobileBackgroundImageURL string `json:"mobileBackgroundImageUrl"`
	MpaaRating               struct {
		Province          interface{} `json:"province"`
		RatingTitle       interface{} `json:"ratingTitle"`
		RatingDescription interface{} `json:"ratingDescription"`
		Warning           interface{} `json:"warning"`
		ImageURL          interface{} `json:"imageUrl"`
	} `json:"mpaaRating"`
	Formats               []string `json:"formats"`
	BrightcoveVideoID     int64    `json:"brightcoveVideoId"`
	URLSlug               string   `json:"urlSlug"`
	Duration              string   `json:"duration"`
	VistaEventCode        string   `json:"vistaEventCode"`
	VistaEventCodeVIP     string   `json:"vistaEventCodeVIP"`
	SeriesTicketingURL    string   `json:"seriesTicketingUrl"`
	SeriesTicketingVIPURL string   `json:"seriesTicketingVIPUrl"`
	Runtime               int      `json:"runtime"`
}

type MoviesOutput struct {
	Data             []MoviesData `json:"data"`
	Skip             int          `json:"skip"`
	Take             int          `json:"take"`
	TotalCount       int          `json:"totalCount"`
	Status           int          `json:"status"`
	Message          string       `json:"message"`
	MessageDetails   string       `json:"messageDetails"`
	APIVersionNumber string       `json:"apiVersionNumber"`
}
