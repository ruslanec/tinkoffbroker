package domain

import "testing"

func TestQuotation_String(t *testing.T) {
	tests := []struct {
		name string
		s    *Quotation
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("Quotation.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuotation_Float64(t *testing.T) {
	tests := []struct {
		name string
		s    *Quotation
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Float64(); got != tt.want {
				t.Errorf("Quotation.Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoneyValue_String(t *testing.T) {
	tests := []struct {
		name string
		s    *MoneyValue
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("MoneyValue.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoneyValue_Float64(t *testing.T) {
	tests := []struct {
		name string
		s    *MoneyValue
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Float64(); got != tt.want {
				t.Errorf("MoneyValue.Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}
