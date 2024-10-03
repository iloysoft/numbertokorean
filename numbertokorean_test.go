package numbertokorean_test

import (
	"fmt"
	"testing"

	"github.com/iloysoft/numbertokorean"
)

func printNumberKorean(n int64) {
	fmt.Printf("%d -> %#v -> %#v -> %#v\n",
		n,
		numbertokorean.Int64ToKoreanUnits(n),
		numbertokorean.Int64ToKoreanLanguage(n, false),
		numbertokorean.Int64ToKoreanLanguage(n, true),
	)
}

func TestInt64ToKorean(t *testing.T) {
	for i := -3; i <= 102; i++ {
		printNumberKorean(int64(i))
	}

	testValues := []int64{
		// -9223372036854775809
		-9223372036854775808,
		-9223372036854775807,
		10000,
		10001,
		11111,
		20000,
		20001,
		100000,
		100001,
		110000,
		110001,
		113560,
		120000,
		120001,
		200000,
		200001,
		210000,
		210001,
		220000,
		220001,
		1765000,
		2367295,
		100000000,
		100000001,
		100010000,
		100010001,
		100020000,
		100020001,
		100030000,
		100030001,
		100110001,
		101010001,
		110010001,
		300000000,
		300000001,
		300010001,
		300020000,
		300020001,
		1234567898,
		4567890123456789,
		73243786789276354,
		1230000000000000000,
		1230000890123456789,
		1234567000023456789,
		1234567890100006789,
		1234567890123450000,
		1234567890123456789,
		9223372036854775806,
		9223372036854775807,
		// 9223372036854775808
	}

	for _, v := range testValues {
		printNumberKorean(v)
	}
}
