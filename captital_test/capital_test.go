package captital_test

import (
	"../capital"
	"fmt"
	"math"
	"testing"
)

func compareWithTolerance(a, b float32) int {
	tolerance := 0.000001
	fmt.Println(a, b, math.Abs(float64(a-b)))
	if math.Abs(float64(a-b)) < tolerance {
		return 0
	}
	return 1
}

func Test_CapitalStateShareValueCalculation(t *testing.T) {
	entity := capital.State{
		Month:      1,
		Year:       2020,
		TotalShare: 950000,
		Balance:    11358.10,
	}
	//if entity.ShareValue() != 0.011956 {
	if compareWithTolerance(entity.ShareValue(), 0.011956) != 0 {
		t.Errorf("Expected: %f, received %f for %d shares with %.2f$ in balance.",
			0.011956, entity.ShareValue(), entity.TotalShare, entity.Balance)
	}
}
