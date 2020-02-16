package xlib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	{"", false},
	{"contains no back slash", false},
}

func TestStrContainBackSlash(t *testing.T) {
	for _, tmp := range testsStrContainBackSlash {
		assert.Equal(t, tmp.r, StrContainBackSlash(tmp.s))
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
	for _, tmp := range testsStrIsPrintRune {
		assert.Equal(t, tmp.r, StrIsPrintRune(tmp.s))
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
	for _, tmp := range testsChangeFileExt {
		s := ChangeFileExt(tmp.iName, tmp.ext)
		assert.Equal(t, tmp.oName, s)
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
	{"ШАГ :. ", "ШАГ:"},
	{" ", " "},
	{"", ""},
}

func TestReplaceSeparators(t *testing.T) {
	s := ""
	for _, tmp := range dReplaceSeparators {
		s = ReplaceSeparators(tmp.iStr)
		assert.Equal(t, tmp.oStr, s)
	}

}

func TestReplaceAllSpace(t *testing.T) {
	s := ReplaceAllSpace("a  b  c")
	assert.Equal(t, s, "a b c", fmt.Sprintf("<ReplaceAllString> return '%s' expect 'a b c'\n", s))
	s = ReplaceAllSpace("a   b   c")
	assert.Equal(t, s, "a b c", fmt.Sprintf("<ReplaceAllString> return '%s' expect 'a b c'\n", s))
	s = ReplaceAllSpace("   a   b   c   ")
	assert.Equal(t, s, " a b c ", fmt.Sprintf("<ReplaceAllString> return '%s' expect ' a b c '\n", s))
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
	assert.False(t, res || (n > 0), fmt.Sprintf("<ContainsOtherRune> on 1 empty test return: '%v', and: '%d'\n", res, n))
	res, n = ContainsOtherRune("ts")
	assert.False(t, res || (n > 0), fmt.Sprintf("<ContainsOtherRune> on 2 empty test return: '%v', and: '%d'\n", res, n))
	res, n = ContainsOtherRune("", '.')
	assert.False(t, res || (n > 0), fmt.Sprintf("<ContainsOtherRune> on 3 empty test return: '%v', and: '%d'\n", res, n))
	for _, ts := range dContainsOtherRune {
		res, n := ContainsOtherRune(ts.s, ts.r)
		assert.Equal(t, ts.res, res)
		assert.Equal(t, ts.n, n)
	}
	res, n = ContainsOtherRune(".�  10 : ���", ' ', '.')
	assert.True(t, res)
	assert.Equal(t, 1, n)
	res, n = ContainsOtherRune(".  10.0 : ���", ' ', '.')
	assert.True(t, res)
	assert.Equal(t, 3, n)
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
	assert.Equal(t, "", s, fmt.Sprintf("<StrCopyStop> on special test return: '%s', expect: '%s'\n", s, "ts: s"))
	assert.Equal(t, 0, i, fmt.Sprintf("<StrCopyStop> on special test return index: '%d', expect: '%d'\n", i, 0))
	s, _ = StrCopyStop("ts: s")
	assert.Equal(t, "ts: s", s, fmt.Sprintf("<StrCopyStop> on special test return: '%s', expect: '%s'\n", s, "ts: s"))
	//space only
	for i, ts := range dStrCopyStop1 {
		s, n := StrCopyStop(ts.s, ' ')
		assert.Equal(t, ts.r, s, fmt.Sprintf("<StrCopyStop> on state 1 test %d return: '%s', expect: '%s'\n", i, s, ts.r))
		assert.Equal(t, ts.n, n, fmt.Sprintf("<StrCopyStop> on state 1 test %d return count: '%d', expect: '%d'\n", i, n, ts.n))
	}
	//space and comma
	for i, ts := range dStrCopyStop3 {
		s, n := StrCopyStop(ts.s, ':', ' ')
		assert.Equal(t, ts.r, s, fmt.Sprintf("<StrCopyStop> on state 3 test %d return: '%s', expect: '%s'\n", i, s, ts.r))
		assert.Equal(t, ts.n, n, fmt.Sprintf("<StrCopyStop> on state 3 test %d return count: '%d', expect: '%d'\n", i, n, ts.n))
	}
}
