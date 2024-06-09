package libs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

type Rest struct {
	apiKey   string
	endpoint string
}

type accomplishmentPayload struct {
	HabitId int    `json:"habit_id"`
	APIKey  string `json:"POMODORO_API_KEY"`
}

func NewRest() *Rest {
	return &Rest{
		apiKey:   viper.GetString("POMODORO_API_KEY"),
		endpoint: viper.GetString("POMODORO_ENDPOINT"),
	}

}

func (r *Rest) SendAccomplishmentRequest(accomplishmentID int) error {
	payload := accomplishmentPayload{
		HabitId: accomplishmentID,
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
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create accomplishment, status code: %d", resp.StatusCode)
	}
	return nil
}
