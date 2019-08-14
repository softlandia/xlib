package xlib

import "testing"

//StrContainBackSlash
type testpairStrContainBackSlash struct {
	s string
	r bool
}

var testsStrContainBackSlash = []testpairStrContainBackSlash{
	{"consist back slash at end \\", true},
	{"consist back slash \\ in middle", true},
	{"\\consist back slash at start", true},
	{"\\", true},
	{"non consist back slash", false},
}

func TestStrContainBackSlash(t *testing.T) {
	for i, tmp := range testsStrContainBackSlash {
		if StrContainBackSlash(tmp.s) != tmp.r {
			t.Errorf("<StrContainBackSlash> on %d test expected %v", i, tmp.r)
		}
	}
}

//StrIsPrintRune
type testpairStrIsPrintRune struct {
	s string
	r bool
}

var testsStrIsPrintRune = []testpairStrIsPrintRune{
	{"consist printable rune \\", true},
	{"consist non" + string(0x00) + "printable rune", false},
}

func TestStrIsPrintRune(t *testing.T) {
	for i, tmp := range testsStrIsPrintRune {
		if StrIsPrintRune(tmp.s) != tmp.r {
			t.Errorf("<StrIsPrintRune> on %d test expected %v", i, tmp.r)
		}
	}
}

//ChangeFileExt
type testpairChangeFileExt struct {
	iName string
	ext   string
	oName string
}

var testsChangeFileExt = []testpairChangeFileExt{
	{"fn.txt", "..", "fn.."},
	{"fn.txt", ".xyz", "fn.xyz"},
}

func TestChangeFileExt(t *testing.T) {
	for i, tmp := range testsChangeFileExt {
		s := ChangeFileExt(tmp.iName, tmp.ext)
		if s != tmp.oName {
			t.Errorf("<ChangeFileExt> on test %d for input: '%s' expected: '%s' got: '%s'", i, tmp.iName, tmp.oName, s)
		}
	}
}

//ConvertCodePage
func TestStrConvertCodePage(t *testing.T) {
	_, err := StrConvertCodePage("1234", Cp866, CpWindows1251)
	if err != nil {
		t.Errorf("<StrConvertCodePage> on test 1 return unexpected err: %v", err)
	}
	_, err = StrConvertCodePage("1234", CpWindows1251, Cp866)
	if err != nil {
		t.Errorf("<StrConvertCodePage> on test 2 return unexpected err: %v", err)
	}
	_, err = StrConvertCodePage("", Cp866, CpWindows1251)
	if err != nil {
		t.Errorf("<StrConvertCodePage> with empty string must return ERROR, but retrurn: %v", err)
	}
	_, err = StrConvertCodePage("1234", Cp866, Cp866)
	if err != nil {
		t.Errorf("<StrConvertCodePage> with equal fromCP and toCp must return nil, but retrurn: %v", err)
	}
}

type TTestReplaceSeparators = struct {
	iStr string
	oStr string
}

var dReplaceSeparators = []TTestReplaceSeparators{
	{"Ш .л", "Ш.л"},
	{"Ш. л", "Ш.л"},
	{"Ш: л", "Ш:л"},
	{"Ш :л", "Ш:л"},
	{"Ш : л", "Ш:л"},
	{"Шаг . ал : фтор", "Шаг.ал:фтор"},
	{" ", " "},
	{"", ""},
}

func TestReplaceSeparators(t *testing.T) {
	s := ""
	for i, tmp := range dReplaceSeparators {
		s = ReplaceSeparators(tmp.iStr)
		if s != tmp.oStr {
			t.Errorf("<ReplaceSeparators> on test %d return '%s' expect '%s'\n", i, s, tmp.oStr)
		}
	}

}

func TestReplaceAllSpace(t *testing.T) {
	s := ReplaceAllSpace("a  b  c")
	if s != "a b c" {
		t.Errorf("<ReplaceAllString> return '%s' expect 'a b c'\n", s)
	}
	s = ReplaceAllSpace("a   b   c")
	if s != "a b c" {
		t.Errorf("<ReplaceAllString> return '%s' expect 'a b c'\n", s)
	}
	s = ReplaceAllSpace("   a   b   c   ")
	if s != " a b c " {
		t.Errorf("<ReplaceAllString> return '%s' expect ' a b c '\n", s)
	}
}

//ContainsOtherRune
type tContainsOtherRune struct {
	s   string
	r   rune
	res bool
	n   int
}

var dContainsOtherRune = []tContainsOtherRune{
	{"ёжз", '.', true, 0},         //0
	{"ё Ё", ' ', true, 0},         //1
	{"  ё  Ё", ' ', true, 2},      //2
	{".... ё .  Ё", '.', true, 4}, //3
	{":::Ё Ж", ':', true, 3},      //4
	{":::", ':', false, 0},        //5
}

func TestContainsOtherRune(t *testing.T) {
	res, n := ContainsOtherRune("")
	if res || (n > 0) {
		t.Errorf("<ContainsOtherRune> on 1 empty test return: '%v', and: '%d'\n", res, n)
	}
	res, n = ContainsOtherRune("ts")
	if res || (n > 0) {
		t.Errorf("<ContainsOtherRune> on 2 empty test return: '%v', and: '%d'\n", res, n)
	}
	res, n = ContainsOtherRune("", '.')
	if res || (n > 0) {
		t.Errorf("<ContainsOtherRune> on 3 empty test return: '%v', and: '%d'\n", res, n)
	}

	for i, ts := range dContainsOtherRune {
		res, n := ContainsOtherRune(ts.s, ts.r)
		if res != ts.res {
			t.Errorf("<ContainsOtherRune> on state 1 test %d return: '%v', expect: '%v'\n", i, res, ts.res)
		}
		if n != ts.n {
			t.Errorf("<ContainsOtherRune> on state 1 test %d return count: '%d', expect: '%d'\n", i, n, ts.n)
		}
	}
	res, n = ContainsOtherRune(".�  10 : ���", ' ', '.')
	if !res {
		t.Errorf("<ContainsOtherRune> on state 2 test 1 return: '%v', expect: '%v'\n", res, true)
	}
	if n != 1 {
		t.Errorf("<ContainsOtherRune> on state 2 test 1 return count: '%d', expect: '%d'\n", n, 1)
	}
	res, n = ContainsOtherRune(".  10.0 : ���", ' ', '.')
	if !res {
		t.Errorf("<ContainsOtherRune> on state 2 test 2 return: '%v', expect: '%v'\n", res, true)
	}
	if n != 3 {
		t.Errorf("<ContainsOtherRune> on state 2 test 2 return count: '%d', expect: '%d'\n", n, 3)
	}
}

type tStrCopyStop struct {
	s string
	r string
	n int
}

//only space
var dStrCopyStop1 = []tStrCopyStop{
	{"ёжз", "ёжз", 6},
	{"ё Ё", "ё", 2},
	{"ё  Ё", "ё", 2},
	{" ё  Ё", "", 0},
	{"ёж:Ё Ж", "ёж:Ё", 7},
	{"ёж :Ё Ж", "ёж", 4},
}

//space and comma
var dStrCopyStop3 = []tStrCopyStop{
	{"ёжз", "ёжз", 6},
	{"ё Ё", "ё", 2},
	{"ё  Ё", "ё", 2},
	{" ё  Ё", "", 0},
	{"ёж:Ё Ж", "ёж", 4},
	{"ёж :Ё Ж", "ёж", 4},
}

func TestStrCopyStop(t *testing.T) {
	s, i := StrCopyStop("")
	if s != "" {
		t.Errorf("<StrCopyStop> on special test return: '%s', expect: '%s'\n", s, "ts: s")
	}
	if i != 0 {
		t.Errorf("<StrCopyStop> on special test return index: '%d', expect: '%d'\n", i, 0)
	}
	s, _ = StrCopyStop("ts: s")
	if s != "ts: s" {
		t.Errorf("<StrCopyStop> on special test return: '%s', expect: '%s'\n", s, "ts: s")
	}
	//space only
	for i, ts := range dStrCopyStop1 {
		s, n := StrCopyStop(ts.s, ' ')
		if s != ts.r {
			t.Errorf("<StrCopyStop> on state 1 test %d return: '%s', expect: '%s'\n", i, s, ts.r)
		}
		if n != ts.n {
			t.Errorf("<StrCopyStop> on state 1 test %d return count: '%d', expect: '%d'\n", i, n, ts.n)
		}
	}
	//space and comma
	for i, ts := range dStrCopyStop3 {
		s, n := StrCopyStop(ts.s, ':', ' ')
		if s != ts.r {
			t.Errorf("<StrCopyStop> on state 3 test %d return: '%s', expect: '%s'\n", i, s, ts.r)
		}
		if n != ts.n {
			t.Errorf("<StrCopyStop> on state 3 test %d return count: '%d', expect: '%d'\n", i, n, ts.n)
		}
	}
}
