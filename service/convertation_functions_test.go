package service

import (
	"testing"

	tkf "github.com/ruslanec/tinkoffbroker/proto"
	"github.com/shopspring/decimal"
	"gotest.tools/assert"
)

func TestConvQuotation(t *testing.T) {
	tests := []struct {
		quotation *tkf.Quotation
		want      string
	}{
		{quotation: nil, want: "0.00"},
		{quotation: &tkf.Quotation{}, want: "0.00"},
		{quotation: &tkf.Quotation{Units: 12, Nano: 250000000}, want: "12.25"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := ConvQuotationFromTkf(tt.quotation)
			assert.Equal(t, tt.want, got.Value.StringFixed(2))
		})
	}
}

func Test_toDecimal(t *testing.T) {
	tests := []struct {
		units int64
		nano  int32
		want  string
	}{
		{units: 114, nano: 250000000, want: "114.25"},
		{units: -214, nano: -700000000, want: "-214.70"},
		{units: 0, nano: 10000000, want: "0.01"},
		{units: -114, nano: 890000000, want: "-114.89"},
		{units: 114, nano: -780000000, want: "114.78"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := toDecimal(tt.units, tt.nano)
			assert.Equal(t, tt.want, got.StringFixed(2))
		})
	}
}

func Test_fromDecimal(t *testing.T) {
	tests := []struct {
		value     string
		wantUnits int64
		wantNano  int32
	}{
		{value: "0.00", wantUnits: 0, wantNano: 0},
		{value: "114.25", wantUnits: 114, wantNano: 250000000},
		{value: "114.250000000", wantUnits: 114, wantNano: 250000000},
		{value: "-200.2", wantUnits: -200, wantNano: -200000000},
		{value: "-200.20", wantUnits: -200, wantNano: -200000000},
		{value: "-200.200000000", wantUnits: -200, wantNano: -200000000},
		{value: "-0.01", wantUnits: 0, wantNano: -10000000},
		{value: "-0.010000000", wantUnits: 0, wantNano: -10000000},
		{value: "123.32", wantUnits: 123, wantNano: 320000000},
		{value: "123.320000000", wantUnits: 123, wantNano: 320000000},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			gotUnits, gotNano := fromDecimal(decimal.RequireFromString(tt.value))
			assert.Equal(t, tt.wantUnits, gotUnits)
			assert.Equal(t, tt.wantNano, gotNano)
		})
	}
}
