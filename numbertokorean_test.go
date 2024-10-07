package numbertokorean_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iloysoft/numbertokorean"
)

func printNumberKorean(n int64) {
	fmt.Printf("%d\n{false, %#v},\n{true, %#v},\n{false, false, %#v},\n{false, true, %#v},\n{true, false, %#v},\n{true, true, %#v},\n",
		n,
		numbertokorean.Int64ToKoreanUnits(n, false),
		numbertokorean.Int64ToKoreanUnits(n, true),
		numbertokorean.Int64ToKoreanLanguage(n, false, false),
		numbertokorean.Int64ToKoreanLanguage(n, false, true),
		numbertokorean.Int64ToKoreanLanguage(n, true, false),
		numbertokorean.Int64ToKoreanLanguage(n, true, true),
	)
}

func TestInt64ToKorean(t *testing.T) {
	for i := -3; i <= 102; i++ {
		//printNumberKorean(int64(i))
	}

	testValues := []int64{
		// -9223372036854775809
		/*
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
			11000000010000,
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
		*/
		// 9223372036854775808
	}

	for _, v := range testValues {
		printNumberKorean(v)
	}
}

func TestInt64ToKoreanUnits(t *testing.T) {
	type inout struct {
		removeEmptyString bool
		expected          []string
	}

	testValues := map[int64][]inout{
		-9223372036854775808: {
			{false, []string{"-", "922경", "3372조", "368억", "5477만", "5808"}},
			{true, []string{"-", "922경", "3372조", "368억", "5477만", "5808"}},
		},
		-9223372036854775807: {
			{false, []string{"-", "922경", "3372조", "368억", "5477만", "5807"}},
			{true, []string{"-", "922경", "3372조", "368억", "5477만", "5807"}},
		},
		-3: {
			{false, []string{"-", "3"}},
			{true, []string{"-", "3"}},
		},
		-2: {
			{false, []string{"-", "2"}},
			{true, []string{"-", "2"}},
		},
		-1: {
			{false, []string{"-", "1"}},
			{true, []string{"-", "1"}},
		},
		0: {
			{false, []string{"0"}},
			{true, []string{"0"}},
		},
		1: {
			{false, []string{"1"}},
			{true, []string{"1"}},
		},
		2: {
			{false, []string{"2"}},
			{true, []string{"2"}},
		},
		3: {
			{false, []string{"3"}},
			{true, []string{"3"}},
		},
		10: {
			{false, []string{"10"}},
			{true, []string{"10"}},
		},
		10000: {
			{false, []string{"1만", ""}},
			{true, []string{"1만"}},
		},
		10001: {
			{false, []string{"1만", "1"}},
			{true, []string{"1만", "1"}},
		},
		20000: {
			{false, []string{"2만", ""}},
			{true, []string{"2만"}},
		},
		100000000: {
			{false, []string{"1억", "", ""}},
			{true, []string{"1억"}},
		},
		100000001: {
			{false, []string{"1억", "", "1"}},
			{true, []string{"1억", "1"}},
		},
		100010000: {
			{false, []string{"1억", "1만", ""}},
			{true, []string{"1억", "1만"}},
		},
		100010001: {
			{false, []string{"1억", "1만", "1"}},
			{true, []string{"1억", "1만", "1"}},
		},
		9223372036854775806: {
			{false, []string{"922경", "3372조", "368억", "5477만", "5806"}},
			{true, []string{"922경", "3372조", "368억", "5477만", "5806"}},
		},
		9223372036854775807: {
			{false, []string{"922경", "3372조", "368억", "5477만", "5807"}},
			{true, []string{"922경", "3372조", "368억", "5477만", "5807"}},
		},
	}

	for k, v := range testValues {
		for _, v2 := range v {
			assert.Equal(t, numbertokorean.Int64ToKoreanUnits(k, v2.removeEmptyString), v2.expected, []any{k, v2.removeEmptyString})
		}
	}
}

func TestInt64ToKoreanLanguage(t *testing.T) {
	type inout struct {
		isMonetary        bool
		removeEmptyString bool
		expected          []string
	}

	testValues := map[int64][]inout{
		-9223372036854775808: {
			{false, false, []string{"마이너스", "구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백팔"}},
			{false, true, []string{"마이너스", "구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백팔"}},
			{true, false, []string{"마이너스", "구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백팔"}},
			{true, true, []string{"마이너스", "구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백팔"}},
		},
		-9223372036854775807: {
			{false, false, []string{"마이너스", "구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백칠"}},
			{false, true, []string{"마이너스", "구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백칠"}},
			{true, false, []string{"마이너스", "구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백칠"}},
			{true, true, []string{"마이너스", "구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백칠"}},
		},
		-3: {
			{false, false, []string{"마이너스", "삼"}},
			{false, true, []string{"마이너스", "삼"}},
			{true, false, []string{"마이너스", "삼"}},
			{true, true, []string{"마이너스", "삼"}},
		},
		-2: {
			{false, false, []string{"마이너스", "이"}},
			{false, true, []string{"마이너스", "이"}},
			{true, false, []string{"마이너스", "이"}},
			{true, true, []string{"마이너스", "이"}},
		},
		-1: {
			{false, false, []string{"마이너스", "일"}},
			{false, true, []string{"마이너스", "일"}},
			{true, false, []string{"마이너스", "일"}},
			{true, true, []string{"마이너스", "일"}},
		},
		0: {
			{false, false, []string{"영"}},
			{false, true, []string{"영"}},
			{true, false, []string{"영"}},
			{true, true, []string{"영"}},
		},
		1: {
			{false, false, []string{"일"}},
			{false, true, []string{"일"}},
			{true, false, []string{"일"}},
			{true, true, []string{"일"}},
		},
		2: {
			{false, false, []string{"이"}},
			{false, true, []string{"이"}},
			{true, false, []string{"이"}},
			{true, true, []string{"이"}},
		},
		3: {
			{false, false, []string{"삼"}},
			{false, true, []string{"삼"}},
			{true, false, []string{"삼"}},
			{true, true, []string{"삼"}},
		},
		10: {
			{false, false, []string{"십"}},
			{false, true, []string{"십"}},
			{true, false, []string{"일십"}},
			{true, true, []string{"일십"}},
		},
		11: {
			{false, false, []string{"십일"}},
			{false, true, []string{"십일"}},
			{true, false, []string{"일십일"}},
			{true, true, []string{"일십일"}},
		},
		12: {
			{false, false, []string{"십이"}},
			{false, true, []string{"십이"}},
			{true, false, []string{"일십이"}},
			{true, true, []string{"일십이"}},
		},
		99: {
			{false, false, []string{"구십구"}},
			{false, true, []string{"구십구"}},
			{true, false, []string{"구십구"}},
			{true, true, []string{"구십구"}},
		},
		100: {
			{false, false, []string{"백"}},
			{false, true, []string{"백"}},
			{true, false, []string{"일백"}},
			{true, true, []string{"일백"}},
		},
		101: {
			{false, false, []string{"백일"}},
			{false, true, []string{"백일"}},
			{true, false, []string{"일백일"}},
			{true, true, []string{"일백일"}},
		},
		102: {
			{false, false, []string{"백이"}},
			{false, true, []string{"백이"}},
			{true, false, []string{"일백이"}},
			{true, true, []string{"일백이"}},
		},
		10000: {
			{false, false, []string{"만", ""}},
			{false, true, []string{"만"}},
			{true, false, []string{"일만", ""}},
			{true, true, []string{"일만"}},
		},
		10001: {
			{false, false, []string{"만", "일"}},
			{false, true, []string{"만", "일"}},
			{true, false, []string{"일만", "일"}},
			{true, true, []string{"일만", "일"}},
		},
		20000: {
			{false, false, []string{"이만", ""}},
			{false, true, []string{"이만"}},
			{true, false, []string{"이만", ""}},
			{true, true, []string{"이만"}},
		},
		110000: {
			{false, false, []string{"십일만", ""}},
			{false, true, []string{"십일만"}},
			{true, false, []string{"일십일만", ""}},
			{true, true, []string{"일십일만"}},
		},
		110001: {
			{false, false, []string{"십일만", "일"}},
			{false, true, []string{"십일만", "일"}},
			{true, false, []string{"일십일만", "일"}},
			{true, true, []string{"일십일만", "일"}},
		},
		113560: {
			{false, false, []string{"십일만", "삼천오백육십"}},
			{false, true, []string{"십일만", "삼천오백육십"}},
			{true, false, []string{"일십일만", "삼천오백육십"}},
			{true, true, []string{"일십일만", "삼천오백육십"}},
		},
		100000000: {
			{false, false, []string{"일억", "", ""}},
			{false, true, []string{"일억"}},
			{true, false, []string{"일억", "", ""}},
			{true, true, []string{"일억"}},
		},
		100010000: {
			{false, false, []string{"일억", "만", ""}},
			{false, true, []string{"일억", "만"}},
			{true, false, []string{"일억", "일만", ""}},
			{true, true, []string{"일억", "일만"}},
		},
		11000000010000: {
			{false, false, []string{"십일조", "", "만", ""}},
			{false, true, []string{"십일조", "만"}},
			{true, false, []string{"일십일조", "", "일만", ""}},
			{true, true, []string{"일십일조", "일만"}},
		},
		9223372036854775806: {
			{false, false, []string{"구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백육"}},
			{false, true, []string{"구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백육"}},
			{true, false, []string{"구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백육"}},
			{true, true, []string{"구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백육"}},
		},
		9223372036854775807: {
			{false, false, []string{"구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백칠"}},
			{false, true, []string{"구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백칠"}},
			{true, false, []string{"구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백칠"}},
			{true, true, []string{"구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백칠"}},
		},
	}

	for k, v := range testValues {
		for _, v2 := range v {
			assert.Equal(t, numbertokorean.Int64ToKoreanLanguage(k, v2.isMonetary, v2.removeEmptyString), v2.expected, []any{k, v2.isMonetary, v2.removeEmptyString})
		}
	}
}
