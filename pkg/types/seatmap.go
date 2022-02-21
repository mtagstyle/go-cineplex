package types

type Seat struct {
	SeatNumber          string      `json:"SeatNumber"`
	AreaCategoryCode    string      `json:"AreaCategoryCode"`
	AccessibilityPrefix string      `json:"AccessibilityPrefix"`
	Status              interface{} `json:"Status"`
	HexBackgroundColour string      `json:"HexBackgroundColour"`
	HexTextColour       string      `json:"HexTextColour"`
	SeatType            int         `json:"SeatType"`
	SeatGridRowID       int64       `json:"SeatGridRowId"`
	AreaNumber          int64       `json:"AreaNumber"`
	GridSeatNum         string      `json:"GridSeatNum"`
}

type SeatMapRow struct {
	RowLabel string `json:"RowLabel"`
	Seats    []Seat `json:"Seats"`
}

type AreaCategory struct {
	VISTAAreaCategoryCode              string `json:"VISTAAreaCategoryCode"`
	VISTAHOAreaCategoryCode            string `json:"VISTAHOAreaCategoryCode"`
	VISTAAreaCategoryDescription       string `json:"VISTAAreaCategoryDescription"`
	VISTAAreaCategoryUseSeatAllocation bool   `json:"VISTAAreaCategoryUseSeatAllocation"`
	SeatingMapLabel                    string `json:"SeatingMapLabel"`
	SeatingMapLegendSeatText           string `json:"SeatingMapLegendSeatText"`
	HexBackgroundColour                string `json:"HexBackgroundColour"`
	HexTextColour                      string `json:"HexTextColour"`
	Image                              string `json:"Image"`
	MobileImage                        string `json:"MobileImage"`
	IsSoldOut                          bool   `json:"IsSoldOut"`
	NumberOfSeatsLeft                  int64  `json:"NumberOfSeatsLeft"`
	IsSessionReservedSeating           bool   `json:"IsSessionReservedSeating"`
	IsPremium                          bool   `json:"IsPremium"`
	CSSClassName                       string `json:"CssClassName"`
}

type SeatMapData struct {
	Rows               []SeatMapRow   `json:"Rows"`
	AreaCategories     []AreaCategory `json:"AreaCategories"`
	CSSWidth           string         `json:"CssWidth"`
	AllocatedSeatCount int            `json:"AllocatedSeatCount"`
	OccupiedSeatCount  int            `json:"OccupiedSeatCount"`
	AvailableSeatCount int            `json:"AvailableSeatCount"`
	IsSoldOut          bool           `json:"IsSoldOut"`
}

type ShowtimeInfo struct {
	SessionData struct {
		ID                int    `json:"Id"`
		VistaSessionID    int    `json:"VistaSessionId"`
		PriceGroupCode    string `json:"PriceGroupCode"`
		DateTime          string `json:"DateTime"`
		BusinessDateTime  string `json:"BusinessDateTime"`
		EndDateTime       string `json:"EndDateTime"`
		IsReservedSeating bool   `json:"IsReservedSeating"`
		IsSoldOut         bool   `json:"IsSoldOut"`
		ScreenName        string `json:"ScreenName"`
		Location          struct {
			ID                                     int64   `json:"Id"`
			Name                                   string  `json:"Name"`
			IsSeatSelectionMapEnabled              bool    `json:"IsSeatSelectionMapEnabled"`
			IsUsherpointEnabled                    bool    `json:"IsUsherpointEnabled"`
			VistaSalesServerWebServiceURL          string  `json:"VistaSalesServerWebServiceUrl"`
			TimeZoneOffset                         float64 `json:"TimeZoneOffset"`
			ServerPrefix                           string  `json:"ServerPrefix"`
			Address1                               string  `json:"Address1"`
			City                                   string  `json:"City"`
			ProvinceCode                           string  `json:"ProvinceCode"`
			MaxTicketsPerOrder                     int64   `json:"MaxTicketsPerOrder"`
			MaxTicketsPerTicketType                int64   `json:"MaxTicketsPerTicketType"`
			PaymentPageID                          int64   `json:"PaymentPageId"`
			MobilePaymentPageID                    int64   `json:"MobilePaymentPageId"`
			Longitude                              float64 `json:"Longitude"`
			Latitude                               float64 `json:"Latitude"`
			IsTimeplay                             bool    `json:"IsTimeplay"`
			SendRefundEmail                        bool    `json:"SendRefundEmail"`
			AllowAccessibilityEquipmentReservation bool    `json:"AllowAccessibilityEquipmentReservation"`
			AllowSellingAccessibilitySeats         bool    `json:"AllowSellingAccessibilitySeats"`
			IsRecRoom                              bool    `json:"IsRecRoom"`
			IsZeroDollarTicketsAvailableToSell     bool    `json:"IsZeroDollarTicketsAvailableToSell"`
			ForceConnectLogin                      bool    `json:"ForceConnectLogin"`
			MapImageURL                            string  `json:"MapImageUrl"`
			IsMakeSuperTicketSceneBurningCall      bool    `json:"IsMakeSuperTicketSceneBurningCall"`
			IsTaxExcluded                          bool    `json:"IsTaxExcluded"`
			IsSeatRestricted                       bool    `json:"IsSeatRestricted"`
			IPAddress                              string  `json:"IpAddress"`
			LocalDateTimeNow                       string  `json:"LocalDateTimeNow"`
		} `json:"Location"`
		Film struct {
			ID                      int         `json:"Id"`
			Title                   string      `json:"Title"`
			ImageURL                string      `json:"ImageUrl"`
			Rating                  string      `json:"Rating"`
			IsRestricted            bool        `json:"IsRestricted"`
			EnchancementCode        string      `json:"EnchancementCode"`
			EnhancementDescription  string      `json:"EnhancementDescription"`
			FirstInLineSceneEndDate interface{} `json:"FirstInLineSceneEndDate"`
			FirstInLinePromoEndDate interface{} `json:"FirstInLinePromoEndDate"`
			SuperTicketDescription  string      `json:"SuperTicketDescription"`
			TicketsMessage          string      `json:"TicketsMessage"`
			IsSuperTicketFilm       bool        `json:"IsSuperTicketFilm"`
			IsPremium               bool        `json:"IsPremium"`
			Is4DX                   bool        `json:"Is4DX"`
			IsVIP                   bool        `json:"IsVIP"`
			DistributorName         string      `json:"DistributorName"`
		} `json:"Film"`
		NumberOfSeatsLeftPerArea       int64         `json:"NumberOfSeatsLeftPerArea"`
		TitleID                        int64         `json:"TitleId"`
		SimultaneousGroupID            int64         `json:"SimultaneousGroupID"`
		TotalSeats                     int           `json:"TotalSeats"`
		Offers                         []interface{} `json:"Offers"`
		AreaCategories                 []interface{} `json:"AreaCategories"`
		BusinessDateTimeString         string        `json:"BusinessDateTimeString"`
		SelectedPerformancesIdentifier string        `json:"SelectedPerformancesIdentifier"`
		VISTAScreenLayoutID            int64         `json:"VISTAScreenLayoutId"`
		VISTAAreaCategoryDescription   string        `json:"VISTAAreaCategoryDescription"`
		VISTAHOAreaCategoryCode        string        `json:"VISTAHOAreaCategoryCode"`
		IsAreaCategorySoldOut          bool          `json:"IsAreaCategorySoldOut"`
		HasSuperTicket                 bool          `json:"HasSuperTicket"`
	} `json:"SessionData"`
	LocationData struct {
		ID                                     int     `json:"Id"`
		Name                                   string  `json:"Name"`
		IsSeatSelectionMapEnabled              bool    `json:"IsSeatSelectionMapEnabled"`
		IsUsherpointEnabled                    bool    `json:"IsUsherpointEnabled"`
		VistaSalesServerWebServiceURL          string  `json:"VistaSalesServerWebServiceUrl"`
		TimeZoneOffset                         int     `json:"TimeZoneOffset"`
		ServerPrefix                           string  `json:"ServerPrefix"`
		Address1                               string  `json:"Address1"`
		City                                   string  `json:"City"`
		ProvinceCode                           string  `json:"ProvinceCode"`
		MaxTicketsPerOrder                     int     `json:"MaxTicketsPerOrder"`
		MaxTicketsPerTicketType                int     `json:"MaxTicketsPerTicketType"`
		PaymentPageID                          int     `json:"PaymentPageId"`
		MobilePaymentPageID                    int     `json:"MobilePaymentPageId"`
		Longitude                              float64 `json:"Longitude"`
		Latitude                               float64 `json:"Latitude"`
		IsTimeplay                             bool    `json:"IsTimeplay"`
		SendRefundEmail                        bool    `json:"SendRefundEmail"`
		AllowAccessibilityEquipmentReservation bool    `json:"AllowAccessibilityEquipmentReservation"`
		AllowSellingAccessibilitySeats         bool    `json:"AllowSellingAccessibilitySeats"`
		IsRecRoom                              bool    `json:"IsRecRoom"`
		IsZeroDollarTicketsAvailableToSell     bool    `json:"IsZeroDollarTicketsAvailableToSell"`
		ForceConnectLogin                      bool    `json:"ForceConnectLogin"`
		MapImageURL                            string  `json:"MapImageUrl"`
		IsMakeSuperTicketSceneBurningCall      bool    `json:"IsMakeSuperTicketSceneBurningCall"`
		IsTaxExcluded                          bool    `json:"IsTaxExcluded"`
		IsSeatRestricted                       bool    `json:"IsSeatRestricted"`
		IPAddress                              string  `json:"IpAddress"`
		LocalDateTimeNow                       string  `json:"LocalDateTimeNow"`
	} `json:"LocationData"`
	Status         int           `json:"Status"`
	Message        string        `json:"Message"`
	MessageDetails string        `json:"MessageDetails"`
	Code           string        `json:"Code"`
	Arguments      []interface{} `json:"Arguments"`
	TransactionUID string        `json:"TransactionUid"`
}

type SeatMapOutput struct {
	SeatMapData      SeatMapData  `json:"SeatMapData"`
	ShowtimeInfo     ShowtimeInfo `json:"ShowtimeInfo"`
	IsMobileSeatMap  bool         `json:"IsMobileSeatMap"`
	ResourceProvider interface{}  `json:"ResourceProvider"`
	CommandResult    struct {
		Status         int           `json:"Status"`
		Message        string        `json:"Message"`
		MessageDetails string        `json:"MessageDetails"`
		Code           string        `json:"Code"`
		Arguments      []interface{} `json:"Arguments"`
		TransactionUID string        `json:"TransactionUid"`
	} `json:"CommandResult"`
	IsMobileTransaction   bool          `json:"IsMobileTransaction"`
	IsZeroDollar          bool          `json:"IsZeroDollar"`
	IsFastPayAvailable    bool          `json:"IsFastPayAvailable"`
	FlashScreens          []interface{} `json:"FlashScreens"`
	GenericErrorMessage   string        `json:"GenericErrorMessage"`
	DiscountAmount        float64       `json:"DiscountAmount"`
	AppliedGiftCardAmount float64       `json:"AppliedGiftCardAmount"`
	EmailAddress          string        `json:"EmailAddress"`
	IsSceneAllowed        bool          `json:"IsSceneAllowed"`
	SecondsToExpire       int64         `json:"SecondsToExpire"`
	MaxSecondsToExpire    int64         `json:"MaxSecondsToExpire"`
}
