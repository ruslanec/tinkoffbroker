package tinkoffbroker

import (
	"reflect"
	"testing"
)

func TestTinkoffBrokerConfig_ToJSON(t *testing.T) {
	tests := []struct {
		name string
		s    *TinkoffBrokerConfig
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToJSON(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TinkoffBrokerConfig.ToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
