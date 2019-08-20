package ops

import (
	"math"
	"reflect"
	"testing"
)

type unaryTest struct {
	val int
	res interface{}
}

type binaryTest struct {
	left  int
	right int
	res   interface{}
}

func TestAbs(t *testing.T) {
	samples := [...]unaryTest{
		{1, 1},
		{-1, 1},
		{0, 0},
		{-5, 5},
		{math.MaxInt32, math.MaxInt32},
		{-math.MaxInt32, math.MaxInt32},
	}
	for _, sample := range samples {
		if got := Abs(sample.val); got != sample.res.(int) {
			t.Errorf("Abs(%d) = %d; want %d", sample.val, got, sample.res)
		}
	}
}

func TestGCD(t *testing.T) {
	samples := [...]binaryTest{
		{1, 1, 1},
		{10, 5, 5},
		{5, 10, 5},
		{40, 16, 8},
		{16, 40, 8},
	}
	for _, sample := range samples {
		if got := GCD(sample.left, sample.right); got != sample.res.(int) {
			t.Errorf("GCD(%d, %d) = %d; want %d", sample.left, sample.right, got, sample.res)
		}
	}
}

func TestDivWithRemainder(t *testing.T) {
	var samples = [...]binaryTest{
		{4, 2, &DivResult{Negative: false, Quotient: 2, Remainder: 0, Denominator: 2}},
		{5, 2, &DivResult{Negative: false, Quotient: 2, Remainder: 1, Denominator: 2}},
		{10, 3, &DivResult{Negative: false, Quotient: 3, Remainder: 1, Denominator: 3}},
		{22, 7, &DivResult{Negative: false, Quotient: 3, Remainder: 1, Denominator: 7}},
		{1, 81, &DivResult{Negative: false, Quotient: 0, Remainder: 1, Denominator: 81}},
		{-6, -4, &DivResult{Negative: false, Quotient: 1, Remainder: 2, Denominator: 4}},
		{-1, 23, &DivResult{Negative: true, Quotient: 0, Remainder: 1, Denominator: 23}},
		{1, -23, &DivResult{Negative: true, Quotient: 0, Remainder: 1, Denominator: 23}},
		{0, 10, &DivResult{Negative: false, Quotient: 0, Remainder: 0, Denominator: 10}},
		{math.MaxInt32, math.MaxInt32 - 1, &DivResult{Negative: false, Quotient: 1, Remainder: 1, Denominator: math.MaxInt32 - 1}},
	}
	for _, sample := range samples {
		if got, _ := DivWithRemainder(sample.left, sample.right); !reflect.DeepEqual(got, sample.res) {
			t.Errorf("DivWithRemainder(%d, %d) = %v; want %v", sample.left, sample.right, got, sample.res)
		}
	}
}

func TestDivWithRemainderErrors(t *testing.T) {
	var samples = [...]binaryTest{
		{10, 0, divisionByZero},
		{math.MaxInt32 + 1, 0, divisionByZero},
		{math.MaxInt32 + 1, 10, intRangeMismatch},
		{-(math.MaxInt32 + 1), 10, intRangeMismatch},
		{10, math.MaxInt32 + 1, intRangeMismatch},
		{10, -(math.MaxInt32 + 1), intRangeMismatch},
	}
	for _, sample := range samples {
		if _, err := DivWithRemainder(sample.left, sample.right); err == nil || err.Error() != sample.res.(string) {
			t.Errorf("DivWithRemainder(%d, %d) error: %v; want %v", sample.left, sample.right, err, sample.res)
		}
	}
}

func TestDivResult_AsPeriodic(t *testing.T) {
	var samples = [...]binaryTest{
		{4, 2, "2"},
		{5, 2, "2.5"},
		{10, 3, "3.(3)"},
		{22, 7, "3.(142857)"},
		{1, 81, "0.(012345679)"},
		{7, 12, "0.58(3)"},
		{1, 23, "0.(0434782608695652173913)"},
		{-6, -4, "1.5"},
		{-1, 23, "-0.(0434782608695652173913)"},
		{1, -23, "-0.(0434782608695652173913)"},
		{0, 10, "0"},
		{math.MaxInt32, math.MaxInt32 - 1, "1.0(000000004656612877414201272105985574522973573322383289525642329385171019830900262883771455738480627293214786139516892041561093220078473184330829590867114803630034256381945923326486650227090949404138093259314161966847406669377728057445704990481683044211643789160664928295337528265395693728137476116546873111787133954267142297948843145695368895023473440691338331151137436880858090595210055443653888482278109045958248009866334507117359439951702430799326329305140608274508834140849145260498994272666977972413392711815790004856688906305179843963291350717927655873752809943401077709534277868954723578835580105740185925495052640787374825018807151372364863168788015068311258282802308241652611868132475640748139136236290574293854249914972344334211521161917151102719037889241276205742057604475000504846685104860630915370426061908179951746184333922513140293316114967052000544082373887377207993881039315723850685845921454807670279226890093858251435550163905648667277441050184370065279649631380708544869635761594060586387347976106543071667238205361401853488201139018126892874135536024473175429248414402127651872232232123829640563418753969873072551445171750565219438229891879698160923736319806181192217563457989658655589128523682363819035108963991616819046080875253380160083416998464071190416879197970851508873376537927777076072783280231788084126811552910890013827839879158735162726356799459417163822257112546132050963241654413977316034936640444152653630965066730012248018767915683582346591709504455057442612068208504531726710993542066769247936801284492762092959845525175189157179732953365643465300689884741501775329468562574562209262104899866604152942638986690565)"},
	}
	for _, sample := range samples {
		if got, _ := DivWithRemainder(sample.left, sample.right); got.AsPeriodic() != sample.res.(string) {
			t.Errorf("DivWithRemainder(%d, %d) = %v; want %v", sample.left, sample.right, got.AsPeriodic(), sample.res)
		}
	}
}

func TestDivResult_AsPlain(t *testing.T) {
	var samples = [...]binaryTest{
		{4, 2, "2"},
		{5, 2, "2(1/2)"},
		{10, 3, "3(1/3)"},
		{22, 7, "3(1/7)"},
		{1, 81, "1/81"},
		{7, 12, "7/12"},
		{16, 40, "2/5"},
		{40, 16, "2(1/2)"},
		{-6, -4, "1(1/2)"},
		{-1, 23, "-1/23"},
		{1, -23, "-1/23"},
		{-24, 23, "-1(1/23)"},
		{24, -23, "-1(1/23)"},
		{0, 10, "0"},
		{math.MaxInt32, math.MaxInt32 - 1, "1(1/2147483646)"},
	}
	for _, sample := range samples {
		if got, _ := DivWithRemainder(sample.left, sample.right); got.AsPlain() != sample.res.(string) {
			t.Errorf("DivWithRemainder(%d, %d) = %v; want %v", sample.left, sample.right, got.AsPlain(), sample.res)
		}
	}
}
