package domain

import (
	"testing"

	"github.com/shopspring/decimal"
	"gotest.tools/assert"
)

func TestMoneyValue_Add(t *testing.T) {
	tests := []struct {
		src  *MoneyValue
		add  *MoneyValue
		want *MoneyValue
	}{
		{
			src: &MoneyValue{
				Currency: "RUB",
				Value:    decimal.RequireFromString("123.45"),
			},
			add: &MoneyValue{
				Currency: "RUB",
				Value:    decimal.RequireFromString("1.5"),
			},
			want: &MoneyValue{
				Currency: "RUB",
				Value:    decimal.RequireFromString("124.95"),
			},
		},
		{
			src: &MoneyValue{
				Currency: "RUB",
				Value:    decimal.RequireFromString("123.45"),
			},
			add: &MoneyValue{
				Currency: "RUB",
				Value:    decimal.RequireFromString("-1.4"),
			},
			want: &MoneyValue{
				Currency: "RUB",
				Value:    decimal.RequireFromString("122.05"),
			},
		},
		{
			src: &MoneyValue{
				Currency: "RUB",
				Value:    decimal.RequireFromString("123.45"),
			},
			add: &MoneyValue{
				Currency: "USD",
				Value:    decimal.RequireFromString("1.5"),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := tt.src.Add(tt.add)
			assert.DeepEqual(t, tt.want, got)
		})
	}
}

func TestMoneyValue_Mul(t *testing.T) {
	tests := []struct {
		src  *MoneyValue
		mul  *Quotation
		want *MoneyValue
	}{
		{
			src: &MoneyValue{
				Currency: "RUB",
				Value:    decimal.RequireFromString("10.5"),
			},
			mul: &Quotation{
				Value: decimal.RequireFromString("2"),
			},
			want: &MoneyValue{
				Currency: "RUB",
				Value:    decimal.RequireFromString("21"),
			},
		},
		{
			src: &MoneyValue{
				Currency: "RUB",
				Value:    decimal.RequireFromString("10.5"),
			},
			mul: &Quotation{
				Value: decimal.RequireFromString("0"),
			},
			want: &MoneyValue{
				Currency: "RUB",
				Value:    decimal.RequireFromString("0"),
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := tt.src.Mul(tt.mul)
			assert.DeepEqual(t, tt.want, got)
		})
	}
}

func TestMoneyValue_Div(t *testing.T) {
	tests := []struct {
		src  *MoneyValue
		mul  *Quotation
		want *MoneyValue
	}{
		{
			src: &MoneyValue{
				Currency: "RUB",
				Value:    decimal.RequireFromString("21"),
			},
			mul: &Quotation{
				Value: decimal.RequireFromString("2"),
			},
			want: &MoneyValue{
				Currency: "RUB",
				Value:    decimal.RequireFromString("10.5"),
			},
		},
		{
			src: &MoneyValue{
				Currency: "RUB",
				Value:    decimal.RequireFromString("21"),
			},
			mul: &Quotation{
				Value: decimal.RequireFromString("0"),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := tt.src.Div(tt.mul)
			assert.DeepEqual(t, tt.want, got)
		})
	}
}
