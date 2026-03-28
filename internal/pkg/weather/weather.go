package weather

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
)

type CurrentWeather struct {
    Temperature float64 `json:"temperature_2m"`
}

type WeatherResponse struct {
    Current CurrentWeather `json:"current"`
}

type WeatherService struct {
    httpClient *http.Client
}

func NewWeatherService() *WeatherService {
    return &WeatherService{
        httpClient: &http.Client{},
    }
}

func (s *WeatherService) GetWeather(lat, lon float64) (*WeatherResponse, error) {
    params := fmt.Sprintf(
        "latitude=%f&longitude=%f&current=temperature_2m",
        lat, lon,
    )
    url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?%s", params)

    resp, err := s.httpClient.Get(url)
    if err != nil {
        return nil, errors.Join(errors.New("failed to fetch weather"), err)
    }
    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, errors.Join(errors.New("failed to read response"), err)
    }

    var result WeatherResponse
    if err := json.Unmarshal(data, &result); err != nil {
        return nil, errors.Join(errors.New("failed to parse response"), err)
    }

    return &result, nil
}
