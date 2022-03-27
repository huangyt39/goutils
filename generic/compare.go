package generic

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/constraints"
)

type Str interface {
	~string
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func SumAsInt[T Str](a, b T) (T, error) {
	ai, err := strconv.ParseInt(string(a), 10, 64)
	if err != nil {
		return T(""), fmt.Errorf("ParseInt failed, str: %s", a)
	}
	bi, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return T(""), fmt.Errorf("ParseInt failed, str: %s", b)
	}
	return T(strconv.FormatInt(ai+bi, 10)), nil
}

func SumAsFloat[T Str](a, b T, prec int) (T, error) {
	af, err := strconv.ParseFloat(string(a), 64)
	if err != nil {
		return T(""), fmt.Errorf("ParseFloat failed, str: %s", a)
	}
	bf, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return T(""), fmt.Errorf("ParseFloat failed, str: %s", a)
	}
	return T(strconv.FormatFloat(af+bf, 'f', prec, 64)), nil
}
