package types

import "time"

type ShowtimeMovie struct {
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
		Province          string `json:"province"`
		RatingTitle       string `json:"ratingTitle"`
		RatingDescription string `json:"ratingDescription"`
		Warning           string `json:"warning"`
		ImageURL          string `json:"imageUrl"`
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

type Showtime struct {
	ShowStartDateTimeUtc          time.Time `json:"showStartDateTimeUtc"`
	ShowStartDateTime             string    `json:"showStartDateTime"`
	ShowStartDateTimeWithTimeZone time.Time `json:"showStartDateTimeWithTimeZone"`
	TimeZoneOffset                float64   `json:"timeZoneOffset"`
	SeatMapURL                    string    `json:"seatMapUrl"`
	TicketingURL                  string    `json:"ticketingUrl"`
	IsReservedSeating             bool      `json:"isReservedSeating"`
	IsShowtimeEnabled             bool      `json:"isShowtimeEnabled"`
	VistaSessionID                int       `json:"vistaSessionId"`
	IsSoldOut                     bool      `json:"isSoldOut"`
	InThePast                     bool      `json:"inThePast"`
	IsOnlineTicketingAvailable    bool      `json:"isOnlineTicketingAvailable"`
	ShowtimeShareKey              string    `json:"showtimeShareKey"`
	Auditorium                    string    `json:"auditorium"`
}

type ShowtimeDetail struct {
	MovieID                  int        `json:"movieId"`
	LocationID               int        `json:"locationId"`
	ExperienceType           string     `json:"experienceType"`
	AreaCode                 string     `json:"areaCode"`
	PresentationType         string     `json:"presentationType"`
	MovieExperienceImageName string     `json:"movieExperienceImageName"`
	ExperienceBannerAltText  string     `json:"experienceBannerAltText"`
	Order                    int        `json:"order"`
	PresentationTypeTitle    string     `json:"presentationTypeTitle"`
	IsPassesAllowed          bool       `json:"isPassesAllowed"`
	NoPassText               string     `json:"noPassText"`
	NoPassDescription        string     `json:"noPassDescription"`
	IsCcEnabled              bool       `json:"isCcEnabled"`
	IsDsEnabled              bool       `json:"isDsEnabled"`
	Showtimes                []Showtime `json:"showtimes"`
}

type ShowtimeData struct {
	Movie           ShowtimeMovie    `json:"movie"`
	ShowtimeDetails []ShowtimeDetail `json:"showtimeDetails"`
}

type ShowtimeOutput struct {
	Data             []ShowtimeData `json:"data"`
	Skip             int            `json:"skip"`
	Take             int            `json:"take"`
	TotalCount       int            `json:"totalCount"`
	Status           int            `json:"status"`
	Message          string         `json:"message"`
	MessageDetails   string         `json:"messageDetails"`
	APIVersionNumber string         `json:"apiVersionNumber"`
}
