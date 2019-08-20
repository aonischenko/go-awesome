package ops

import (
	"fmt"
	"math"
	"strings"
)

const (
	divisionByZero   = "division by zero"
	intRangeMismatch = "only int32 values are supported: [-2147483647, 2147483647]"
)

type DivResult struct {
	Negative    bool
	Quotient    int
	Remainder   int
	Denominator int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GCD(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func DivWithRemainder(numerator, denominator int) (*DivResult, error) {
	if denominator == 0 {
		return nil, fmt.Errorf(divisionByZero)
	}
	if Abs(numerator) > math.MaxInt32 || Abs(denominator) > math.MaxInt32 {
		return nil, fmt.Errorf(intRangeMismatch)
	}
	return &DivResult{
		Negative:    (numerator < 0) != (denominator < 0),
		Quotient:    Abs(numerator / denominator),
		Remainder:   Abs(numerator % denominator),
		Denominator: Abs(denominator),
	}, nil
}

func (r *DivResult) AsPeriodic() string {
	var sb strings.Builder
	if r.Negative {
		sb.WriteRune('-')
	}
	sb.WriteString(fmt.Sprint(r.Quotient))
	if r.Remainder > 0 {
		sb.WriteRune('.')
		dict := make(map[int]int)
		res := make([]rune, 0) // seems it's a common practice to initialize zero length slice to prevent panics
		pos := 0
		for rem := r.Remainder; rem > 0; rem = rem * 10 % r.Denominator {
			if val, exists := dict[rem]; !exists {
				dict[rem] = pos
				next := rem * 10 / r.Denominator
				code := 48 + int(next)
				res = append(res, rune(code))
				pos++
			} else {
				res = append(res[:val], append([]rune{'('}, append(res[val:], ')')...)...)
				break
			}
		}
		for k := 0; k < len(res); k++ {
			sb.WriteRune(res[k])
		}
	}
	return sb.String()
}

func (r *DivResult) AsPlain() string {
	var sb strings.Builder
	if r.Negative {
		sb.WriteRune('-')
	}
	hasQuotient := r.Quotient > 0 || r.Remainder == 0
	if hasQuotient {
		sb.WriteString(fmt.Sprint(r.Quotient))
	}
	if r.Remainder > 0 {
		gcd := GCD(r.Remainder, r.Denominator)
		if hasQuotient {
			sb.WriteRune('(')
		}
		sb.WriteString(fmt.Sprint(r.Remainder / gcd))
		sb.WriteRune('/')
		sb.WriteString(fmt.Sprint(r.Denominator / gcd))
		if hasQuotient {
			sb.WriteRune(')')
		}
	}
	return sb.String()
}
