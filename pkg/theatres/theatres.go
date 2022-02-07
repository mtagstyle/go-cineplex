package theatres

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mtagstyle/cineplex/pkg/types"
)

type GetTheatresInput struct {
	Range int
}

type GetTheatresOutput struct {
	Output types.TheatreOutput
}

type TheatresAPI interface {
	GetTheatres(input *GetTheatresInput) (*GetTheatresOutput, error)
}

type theatresAPIClient struct {
}

const theatresEndpoint = "https://www.cineplex.com/api/v1/theatres"

func NewTheatresAPIClient() *theatresAPIClient {
	return &theatresAPIClient{}
}

func (t *theatresAPIClient) GetTheatres(input *GetTheatresInput) (*GetTheatresOutput, error) {
	query, err := t.buildFullRequestPath(input)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	output := types.TheatreOutput{}
	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return &GetTheatresOutput{Output: output}, nil
}

func (t *theatresAPIClient) buildFullRequestPath(in *GetTheatresInput) (string, error) {
	req, err := http.NewRequest("GET", theatresEndpoint, nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("language", "en-us")
	q.Add("range", fmt.Sprintf("%d", in.Range))
	q.Add("skip", "0")
	q.Add("take", "1000")
	req.URL.RawQuery = q.Encode()

	return req.URL.String(), nil
}