package movies

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/mtagstyle/go-cineplex/pkg/types"
)

type MoviesAPI interface {
	GetMovies(input *GetMoviesInput) (*GetMoviesOutput, error)
}

type moviesAPIClient struct{}

type GetMoviesInput struct {
	Language                 string
	MarketLanguageCodeFilter bool
	MovieType                int
	ShowtimeType             int
	ShowtimeStatus           int
	Skip                     int
	Take                     int
}

const moviesBaseEndpoint = "https://www.cineplex.com/api/v1/movies"

type GetMoviesOutput struct {
	Output types.MoviesOutput
}

func (m *moviesAPIClient) GetMovies(input *GetMoviesInput) (*GetMoviesOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *moviesAPIClient) GetShowtimes(input *GetMoviesInput) (*GetMoviesOutput, error) {
	query, err := m.buildFullRequestPath(input)
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

	output := types.MoviesOutput{}
	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return &GetMoviesOutput{Output: output}, nil
}

func (m *moviesAPIClient) buildFullRequestPath(in *GetMoviesInput) (string, error) {
	fullEp := fmt.Sprintf("%s", moviesBaseEndpoint)
	return fullEp, nil
}
