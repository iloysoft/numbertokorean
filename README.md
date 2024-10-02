
# numbertokorean

library to convert `int64` into `[]string` in Sino-Korean system

## how to install

```
go get github.com/iloysoft/numbertokorean
```

## how to use

Import `github.com/iloysoft/numbertokorean` package into your code

There are 2 functions.

 - `Int64ToKoreanUnits()`: split `int64` into `[]string` with Sino-Korean units (만, 억, 조, 경)
    - `9223372036854775807` will be `[]string{"922경", "3372조", "368억", "5477만", "5807"}`

     - `Int64ToKoreanLanguage()`: convert `int64` into `[]string` in Sino-Korean system
        - `9223372036854775807` will be `[]string{"구백이십이경", "삼천삼백칠십이조", "삼백육십팔억", "오천사백칠십칠만", "오천팔백칠"}`

