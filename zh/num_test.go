package zh

import (
	"fmt"
	"testing"
)

var testdata = []struct {
	num Uint64
	str string
}{
	{0, "零"},
	{1, "一"},
	{2, "二"},
	{3, "三"},
	{4, "四"},
	{5, "五"},
	{6, "六"},
	{7, "七"},
	{8, "八"},
	{9, "九"},
	{10, "十"},
	{11, "十一"},
	{110, "一百一十"},
	{111, "一百一十一"},
	{100, "一百"},
	{102, "一百零二"},
	{1020, "一千零二十"},
	{1001, "一千零一"},
	{1015, "一千零一十五"},
	{1000, "一千"},
	{10000, "一万"},
	{20010, "二万零一十"},
	{20001, "二万零一"},
	{100000, "十万"},
	{1000000, "一百万"},
	{10000000, "一千万"},
	{100000000, "一亿"},
	{1000000000, "十亿"},
	{1000001000, "十亿一千"},
	{1000000100, "十亿零一百"},
	{200010, "二十万零一十"},
	{2000105, "二百万零一百零五"},
	{20001007, "二千万一千零七"},
	{2000100190, "二十亿零一十万零一百九十"},
	{1040010000, "十亿四千零一万"},
	{200012301, "二亿零一万二千三百零一"},
	{2005010010, "二十亿零五百零一万零一十"},
	{4009060200, "四十亿零九百零六万零二百"},
	{4294967295, "四十二亿九千四百九十六万七千二百九十五"},
}

func TestUint64_String(t *testing.T) {
	for _, v := range testdata {
		if str := v.num.String(); str != v.str {
			t.Errorf("spell %d error: return %q, want %q", v.num, str, v.str)
		}
	}
}

func TestUint64_Scan(t *testing.T) {
	for _, v := range testdata {
		var num Uint64

		_, err := fmt.Sscan(v.str, &num)
		if err != nil {
			t.Errorf("parse %q error: %v", v.str, err)
			continue
		}

		if num != v.num {
			t.Errorf("parse %q error: return %d, want %d", v.str, num, v.num)
		}
	}
}

func ExampleString() {
	fmt.Print(Uint64(1234))

	// Output: 一千二百三十四
}
