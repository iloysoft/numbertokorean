
# numbertokorean

Library to convert `int64` into `[]string` in Sino-Korean system

## how to install

```
go get github.com/iloysoft/numbertokorean
```

## how to use

Import `github.com/iloysoft/numbertokorean` package into your code

There are 2 functions.

- `Int64ToKoreanUnits(n int64, removeEmptyString bool)`: split `int64` into `[]string` with Sino-Korean units (만, 억, 조, 경)
    - `Int64ToKoreanUnits(-9223372036854775808, false)` will be `[]string{"-", "922경", "3372조", "368억", "5477만", "5808"}`
    - `Int64ToKoreanUnits(9223372036854775807, false)` will be `[]string{"922경", "3372조", "368억", "5477만", "5807"}`
    - `Int64ToKoreanUnits(11000000010000, false)` will be `[]string{"11조", "", "1만", ""}`
    - `Int64ToKoreanUnits(11000000010000, true)` will be `[]string{"11조", "1만"}`

- `Int64ToKoreanLanguage(n int64, isMonetary bool, removeEmptyString bool)`: convert `int64` into `[]string` in Sino-Korean system
    - `Int64ToKoreanLanguage(-9223372036854775808, false, false)` will be `[]string{"마이너스", "구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백팔"}`
    - `Int64ToKoreanLanguage(9223372036854775807, false, false)` will be `[]string{"구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백칠"}`
    - `Int64ToKoreanLanguage(113560, false, false)` will be `[]string{"십일만", "삼천오백육십"}`
    - `Int64ToKoreanLanguage(113560, true, false)` will be `[]string{"일십일만", "삼천오백육십"}`
    - `Int64ToKoreanLanguage(1765000, false, false)` will be `[]string{"백칠십육만", "오천"}`
    - `Int64ToKoreanLanguage(1765000, true, false)` will be `[]string{"일백칠십육만", "오천"}`
    - `Int64ToKoreanLanguage(11000000010000, false, false)` will be `[]string{"십일조", "", "만", ""}`
    - `Int64ToKoreanLanguage(11000000010000, false, true)` will be `[]string{"십일조", "만"}`
    - `Int64ToKoreanLanguage(11000000010000, true, false)` will be `[]string{"일십일조", "", "일만", ""}`
    - `Int64ToKoreanLanguage(11000000010000, true, true)` will be `[]string{"일십일조", "일만"}`

## reference

한글 맞춤법 제5장 띄어쓰기 제44항 \
수를 적을 적에는 ‘만(萬)’ 단위로 띄어 쓴다. \
https://korean.go.kr/kornorms/regltn/regltnView.do?regltn_code=0001&regltn_no=182

지방자치단체 재무회계규칙 제7장 제124조(두서금액의 표시) \
문서 및 유가증권에 금액을 표시하는 때에는 아라비아숫자로 쓰되, 숫자 다음에 괄호를 하고 다음 예시와 같이 한글로 기재하여야 한다. ｛예시 : 금113,560원(금일십일만삼천오백육십원)｝\
https://law.go.kr/LSW/admRulLsInfoP.do?admRulSeq=2100000152149

