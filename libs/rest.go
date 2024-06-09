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
	HabitId int `json:"habit_id"`
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
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", r.endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", r.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
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
