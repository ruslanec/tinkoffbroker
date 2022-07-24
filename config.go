package tinkoffbroker

import (
	"encoding/json"
	"fmt"
	"time"
)

const REQ_TIMEOUT time.Duration = time.Second * 60 * 3

type TinkoffBrokerConfig struct {
	ApiUrl    string `json:"api_url,omitempty"`    // URL адрес API сервиса Тинькофф инвестиции
	Token     string `json:"token,omitempty"`      // Токен доступа к сервису Тинькофф инвестиции
	AccountID string `json:"account_id,omitempty"` // Номер счета в сервисе API Тинькофф инвестиции
}

func (s *TinkoffBrokerConfig) ToJSON() []byte {
	v, err := json.Marshal(&s)
	if err != nil {
		return []byte(fmt.Sprintf("{\"error\": \"%v\"}", err)) //FIXME: errors to hson
	}
	return v
}
