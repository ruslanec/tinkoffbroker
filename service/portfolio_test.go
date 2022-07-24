package service

import (
	"testing"

	"github.com/ruslanec/tinkoffbroker/domain"
)

func TestPortfolio_RiskLevel(t *testing.T) {
	type args struct {
		figi string
	}
	tests := []struct {
		name    string
		p       *domain.Portfolio
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.RiskLevel(tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("Portfolio.RiskLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Portfolio.RiskLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPortfolioPosition_String(t *testing.T) {
	tests := []struct {
		name string
		s    *domain.PortfolioPosition
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("PortfolioPosition.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
