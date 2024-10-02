package numbertokorean

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

// int64 : -922_3372_0368_5477_5808 to 922_3372_0368_5477_5807
var (
	minus         = "마이너스"
	zero          = "영"
	unitsBig      = []string{"", "만", "억", "조", "경"} // in reverse order
	unitsSmall    = []string{"천", "백", "십"}
	numbersFor012 = []string{"", "", "이", "삼", "사", "오", "육", "칠", "팔", "구"}
	numbersFor3   = []string{"", "일", "이", "삼", "사", "오", "육", "칠", "팔", "구", "십"}
)

const (
	sizeofRune = 4 // rune is int32 in golang
)

func parseInt64(n int64) []string {
	if n == 0 {
		return []string{"0"}
	}

	if n == math.MinInt64 {
		return []string{"5808", "5477", "368", "3372", "922", "-"}
	}

	neg := false
	if n < 0 {
		neg = true
		n = -n
	}

	ret := parseUint64(uint64(n))

	if neg {
		ret = append(ret, "-")
	}

	return ret
}

func parseUint64(n uint64) []string {
	if n == 0 {
		return []string{"0"}
	}

	ret := make([]string, 0, len(unitsBig)+1) // +1 for minus sign

	s := strconv.FormatUint(n, 10)

	prevPos := len(s)
	curUnit := 0
	for i := prevPos - 4; ; i = i - 4 {
		if i < 0 {
			i = 0
		}

		word := strings.TrimLeft(s[i:prevPos], "0")
		ret = append(ret, word)
		curUnit += 1

		if i == 0 {
			break
		} else {
			prevPos = i
		}
	}

	return ret
}

func Int64ToKoreanUnits(n int64) []string {
	if n == 0 {
		return []string{"0"}
	}

	ret := parseInt64(n)

	for i := 0; i < len(ret); i++ {
		if len(ret[i]) > 0 {
			if ret[i] != "-" {
				ret[i] += unitsBig[i]
			}
		}
	}

	slices.Reverse(ret)

	return ret
}

func splitNumberByDigits(s string) [4]byte {
	s2 := []byte(s)

	switch len(s2) {
	case 4:
		return [4]byte{s2[0] - '0', s2[1] - '0', s2[2] - '0', s2[3] - '0'}
	case 3:
		return [4]byte{0, s2[0] - '0', s2[1] - '0', s2[2] - '0'}
	case 2:
		return [4]byte{0, 0, s2[0] - '0', s2[1] - '0'}
	case 1:
		return [4]byte{0, 0, 0, s2[0] - '0'}
	default:
		panic("invalid argument")
	}
}

func readNumberInKorean(s string, isSecondPart bool) *strings.Builder {
	nums := splitNumberByDigits(s)

	var sb strings.Builder

	sb.Grow(8 * sizeofRune) // x천x백x십x + unitsBig

	for i := 0; i < 3; i++ {
		if nums[i] > 0 {
			sb.WriteString(numbersFor012[nums[i]])
			sb.WriteString(unitsSmall[i])
		}
	}

	if nums[3] > 0 {
		if isSecondPart && nums[0] == 0 && nums[1] == 0 && nums[2] == 0 {
			// NOTE
			// special case....
			sb.WriteString(numbersFor012[nums[3]])
		} else {
			sb.WriteString(numbersFor3[nums[3]])
		}
	}

	return &sb
}

func Int64ToKoreanLanguage(n int64) []string {
	if n == 0 {
		return []string{zero}
	}

	ret := parseInt64(n)

	for i := 0; i < len(ret); i++ {
		if len(ret[i]) > 0 {
			if ret[i] != "-" {
				sb := readNumberInKorean(ret[i], i == 1)
				sb.WriteString(unitsBig[i])
				ret[i] = sb.String()
			} else {
				ret[i] = minus
			}
		}
	}

	slices.Reverse(ret)

	return ret
}
