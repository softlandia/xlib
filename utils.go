package xlib

//Epsilon - precission
const Epsilon float64 = 0.01

//
const (
	//CpEmpty not define code page
	CpEmpty      = 0
	CpEmptyAsStr = "undefine"
	//Windows1251 is the Windows 1251 code page
	CpWindows1251      = 1
	CpWindows1251AsStr = "1251"
	//CodePage866 is the IBM Code Page 866
	Cp866      = 2
	Cp866AsStr = "866"
	//CodePageUtf8
	CpUtf8      = 3
	CpUtf8AsStr = "UTF8"
)

//const CpNames [3]string = {CpEmptyAsStr, CpWindows1251AsStr, Cp866AsStr}
