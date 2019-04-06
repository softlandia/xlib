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

func TestConvertStrCodePage(t *testing.T) {
	_, err := ConvertStrCodePage("1234", Cp866, CpWindows1251)
	if err != nil {
		t.Errorf("<ConvertStrCodePage> on test 1 return unexpected err: %v", err)
	}
	_, err = ConvertStrCodePage("1234", CpWindows1251, Cp866)
	if err != nil {
		t.Errorf("<ConvertStrCodePage> on test 2 return unexpected err: %v", err)
	}
	_, err = ConvertStrCodePage("", Cp866, CpWindows1251)
	if err != nil {
		t.Errorf("<ConvertStrCodePage> with empty string must return ERROR, but retrurn: %v", err)
	}
	_, err = ConvertStrCodePage("1234", Cp866, Cp866)
	if err != nil {
		t.Errorf("<ConvertStrCodePage> with equal fromCP and toCp must return nil, but retrurn: %v", err)
	}
}
