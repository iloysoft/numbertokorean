package numbertokorean

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

// int64 : -922_3372_0368_5477_5808 to 922_3372_0368_5477_5807
var (
	minus           = "마이너스"
	zero            = "영"
	unitsBig        = []string{"", "만", "억", "조", "경"} // in reverse order
	unitsSmall      = []string{"천", "백", "십"}
	numbersImplicit = []string{"", "", "이", "삼", "사", "오", "육", "칠", "팔", "구"}
	numbersExplicit = []string{"", "일", "이", "삼", "사", "오", "육", "칠", "팔", "구"}
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
	for i := prevPos - 4; ; i = i - 4 {
		if i < 0 {
			i = 0
		}

		word := strings.TrimLeft(s[i:prevPos], "0")
		ret = append(ret, word)

		if i == 0 {
			break
		} else {
			prevPos = i
		}
	}

	return ret
}

func Int64ToKoreanUnits(n int64, removeEmptyString bool) []string {
	if n == 0 {
		return []string{"0"}
	}

	ret := parseInt64(n)
	dst := 0

	for src := 0; src < len(ret); src++ {
		if !removeEmptyString {
			dst = src
		}

		if ret[src] != "" {
			if ret[src] != "-" {
				ret[dst] = ret[src] + unitsBig[src]
			} else {
				ret[dst] = "-"
			}
			dst += 1
		}
	}

	if removeEmptyString {
		for i := dst; i < len(ret); i++ {
			ret[i] = ""
		}
		ret = ret[:dst]
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

func readNumberInKorean(s string, isSecondPart bool, isMonetary bool) *strings.Builder {
	nums := splitNumberByDigits(s)

	var sb strings.Builder

	sb.Grow(8 * sizeofRune) // x천x백x십x + unitsBig

	for i := 0; i < 3; i++ {
		if nums[i] > 0 {
			if isMonetary {
				sb.WriteString(numbersExplicit[nums[i]])
			} else {
				sb.WriteString(numbersImplicit[nums[i]])
			}
			sb.WriteString(unitsSmall[i])
		}
	}

	if nums[3] > 0 {
		if isMonetary {
			sb.WriteString(numbersExplicit[nums[3]])
		} else {
			if isSecondPart && nums[0] == 0 && nums[1] == 0 && nums[2] == 0 {
				// NOTE
				// special case....
				sb.WriteString(numbersImplicit[nums[3]])
			} else {
				sb.WriteString(numbersExplicit[nums[3]])
			}
		}
	}

	return &sb
}

func Int64ToKoreanLanguage(n int64, isMonetary bool, removeEmptyString bool) []string {
	if n == 0 {
		return []string{zero}
	}

	ret := parseInt64(n)
	dst := 0

	for src := 0; src < len(ret); src++ {
		if !removeEmptyString {
			dst = src
		}

		if ret[src] != "" {
			if ret[src] != "-" {
				sb := readNumberInKorean(ret[src], src == 1, isMonetary)
				sb.WriteString(unitsBig[src])
				ret[dst] = sb.String()
			} else {
				ret[dst] = minus
			}
			dst += 1
		}
	}

	if removeEmptyString {
		for i := dst; i < len(ret); i++ {
			ret[i] = ""
		}
		ret = ret[:dst]
	}

	slices.Reverse(ret)

	return ret
}
