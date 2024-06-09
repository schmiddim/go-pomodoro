package libs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

type Rest struct {
	apiKey   string
	endpoint string
}

type accomplishmentPayload struct {
	HabitId int    `json:"habit_id"`
	APIKey  string `json:"api_key"`
}

func NewRest() *Rest {
	return &Rest{
		apiKey:   viper.GetString("POMODORO_API_KEY"),
		endpoint: viper.GetString("POMODORO_ENDPOINT"),
	}

}

func (r *Rest) SendAccomplishmentRequest(habitId int) error {
	payload := accomplishmentPayload{
		HabitId: habitId,
		APIKey:  r.apiKey,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(r.endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create accomplishment, status code: %d", resp.StatusCode)
	}
	return nil
}
