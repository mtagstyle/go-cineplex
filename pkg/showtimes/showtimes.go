package showtimes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/mtagstyle/go-cineplex/pkg/types"
)

type ShowtimesAPI interface {
	GetShowtimes(input *GetShowtimesInput) (*GetShowtimesOutput, error)
	GetSeatMapData(input *GetSeatMapDataInput) (*GetSeatMapDataOutput, error)
}

type showtimesAPIClient struct {
	client *retryablehttp.Client
}

type GetShowtimesInput struct {
	TheatreID string
	ShowDate  string
}

type GetShowtimesOutput struct {
	Output types.ShowtimeOutput
}

type GetSeatMapDataInput struct {
	VistaSessionID string
	TheatreID      string
}

type GetSeatMapDataOutput struct {
	Output types.SeatMapOutput
}

const showtimesBaseEndpoint = "https://www.cineplex.com/api/v1/theatres"

func NewShowtimesAPIClient() *showtimesAPIClient {
	client := &showtimesAPIClient{
		client: retryablehttp.NewClient(),
	}

	client.client.RetryMax = 10
	return client
}

//===================
// GetShowtimes
//===================
func (t *showtimesAPIClient) GetShowtimes(input *GetShowtimesInput) (*GetShowtimesOutput, error) {
	if err := t.validateGetShowtimesInputArgs(input); err != nil {
		return nil, err
	}

	query, err := t.buildFullRequestPath(input)
	if err != nil {
		return nil, err
	}

	resp, err := t.client.Get(query)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	output := types.ShowtimeOutput{}
	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return &GetShowtimesOutput{Output: output}, nil
}

func (t *showtimesAPIClient) validateGetShowtimesInputArgs(in *GetShowtimesInput) error {
	if len(in.TheatreID) == 0 {
		return fmt.Errorf("invalid TheatreID '%s'", in.TheatreID)
	}

	if len(in.ShowDate) == 0 {
		return fmt.Errorf("invalid ShowDate '%s'", in.ShowDate)
	}

	return nil
}

func (t *showtimesAPIClient) buildFullRequestPath(in *GetShowtimesInput) (string, error) {
	fullEp := fmt.Sprintf("%s/%s/availablemovies/showtimesoneposter?showDate=%s", showtimesBaseEndpoint, in.TheatreID, in.ShowDate)
	req, err := http.NewRequest("GET", fullEp, nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	return req.URL.String(), nil
}

//===================
// GetSeatMap
//===================
func (t *showtimesAPIClient) GetSeatMapData(input *GetSeatMapDataInput) (*GetSeatMapDataOutput, error) {
	data := url.Values{
		"VistaSessionId": {fmt.Sprint(input.VistaSessionID)},
		"LocationId":     {fmt.Sprint(input.TheatreID)},
	}
	resp, err := t.client.PostForm("https://onlineticketing.cineplex.com/SeatMap/GetSeatMapData", data)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	output := types.SeatMapOutput{}
	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return &GetSeatMapDataOutput{Output: output}, nil
}
