package partition

import (
	"aoc/test"
	"testing"
)

func genSolver(startIdx int) string {
	str, err := Gen(startIdx, 12)
	if err != nil {
		return ""
	}
	return str
}

func TestPartition(t *testing.T) {

	cases := []test.Case[int, string]{
		{0, "PARTITION ( archive, June_2019, July_2019, August_2019, September_2019, October_2019, November_2019, December_2019, January_2020, February_2020, March_2020, April_2020 )"},
		{1, "PARTITION ( June_2019, July_2019, August_2019, September_2019, October_2019, November_2019, December_2019, January_2020, February_2020, March_2020, April_2020, May_2020 )"},
		{43, "PARTITION ( December_2022, April_2023, May_2023, June_2023, July_2023, August_2023, September_2023, October_2023, November_2023 )"},
		{44, "PARTITION ( April_2023, May_2023, June_2023, July_2023, August_2023, September_2023, October_2023, November_2023, December_2023 )"},
		{45, "PARTITION ( April_2023, May_2023, June_2023, July_2023, August_2023, September_2023, October_2023, November_2023, December_2023, January_2024 )"},
		{46, "PARTITION ( April_2023, May_2023, June_2023, July_2023, August_2023, September_2023, October_2023, November_2023, December_2023, January_2024, February_2024 )"},
		{47, "PARTITION ( April_2023, May_2023, June_2023, July_2023, August_2023, September_2023, October_2023, November_2023, December_2023, January_2024, February_2024, March_2024 )"},
		{48, "PARTITION ( May_2023, June_2023, July_2023, August_2023, September_2023, October_2023, November_2023, December_2023, January_2024, February_2024, March_2024, April_2024 )"},
		{250, "PARTITION ( March_2040, April_2040, May_2040, future )"},
		{252, "PARTITION ( May_2040, future )"},
		{253, "PARTITION ( future )"},
		{254, ""},
	}

	err := test.Execute(cases, genSolver)
	if err != nil {
		t.Error(err)
	}

}
