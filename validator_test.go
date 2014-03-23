package validator_test

import (
	"testing"

	. "github.com/alejandrodumas/validator"
)

type TestCase struct {
	f       func(string) bool
	valid   []string
	invalid []string
}

var testCases = map[string]TestCase{
	"IsEmail": TestCase{
		IsEmail,
		[]string{"foo@bar.com", "x@d.c", "foo@bar.com.au", "foo+bar@bar.com"},
		[]string{"invalidemail@", "invalid.com", "@invalid.com"},
	},
	"IsIP": TestCase{
		IsIP,
		[]string{"127.0.0.1", "0.0.0.0", "255.255.255.255", "1.2.3.4", "::1", "2001:db8:0000:1:1:1:1:1"},
		[]string{"abc", "256.0.0.0", "0.0.0.256"},
	},
	"IsAlpha": TestCase{
		IsAlpha,
		[]string{"abc", "ABC", "FoOBar"},
		[]string{"abc1", "  foo ", ""},
	},
	"IsAlphanumeric": TestCase{
		IsAlphanumeric,
		[]string{"abc123", "ABC11"},
		[]string{"abc ", "foo!!"},
	},
	"IsLowerCase": TestCase{
		IsLowerCase,
		[]string{"abc", "abc123", "this is lowercase", "très über"},
		[]string{"fooBar", "123A"},
	},
	"IsUpperCase": TestCase{
		IsUpperCase,
		[]string{"ABC", "ABC123", "ALL CAPS IS FUN"},
		[]string{"abc", "abc123", "this is lowercase", "très über"},
	},
	"IsHexColor": TestCase{
		IsHexColor,
		[]string{"#ff0034", "#CCCCCC", "fff", "#f00"},
		[]string{"#ff", "fff0", "#ff12FG"},
	},
	"IsHexadecimal": TestCase{
		IsHexadecimal,
		[]string{"deadBEEF", "ff0044"},
		[]string{"abcdefg", "", ".."},
	},
	"IsBool": TestCase{
		IsBool,
		[]string{"1", "t", "f", "0", "TRUE", "FALSE", "T", "F", "True", "False", "true", "false"},
		[]string{"a", "verdadero", "fv", "00", "01"},
	},
	"IsNull": TestCase{
		IsNull,
		[]string{""},
		[]string{" ", "foo"},
	},
	"IsDate": TestCase{
		IsDate,
		[]string{"2011-08-04"},
		[]string{""},
	},
	"IsISBN": TestCase{
		IsISBN,
		[]string{"340101319X", "9784873113685"},
		[]string{"3423214121", "9783836221190"},
	},
	"IsISBNv10": TestCase{
		IsISBNv10,
		[]string{
			"3836221195",
			"3-8362-2119-5",
			"3 8362 2119 5",
			"1617290858",
			"1-61729-085-8",
			"1 61729 085-8",
			"0007269706",
			"0-00-726970-6",
			"0 00 726970 6",
			"3423214120",
			"3-423-21412-0",
			"3 423 21412 0",
			"340101319X",
			"3-401-01319-X",
			"3 401 01319 X"},
		[]string{
			"3423214121",
			"3-423-21412-1",
			"3 423 21412 1",
			"978-3836221191",
			"9783836221191",
			"123456789a",
			"foo",
			""},
	},
	"IsISBNv13": TestCase{
		IsISBNv13,
		[]string{
			"9783836221191",
			"978-3-8362-2119-1",
			"978 3 8362 2119 1",
			"9783401013190",
			"978-3401013190",
			"978 3401013190",
			"9784873113685",
			"978-4-87311-368-5",
			"978 4 87311 368 5"},
		[]string{
			"9783836221190",
			"978-3-8362-2119-0",
			"978 3 8362 2119 0",
			"3836221195",
			"3-8362-2119-5",
			"3 8362 2119 5",
			"01234567890ab",
			"foo",
			""},
	},
	"IsInt": TestCase{
		IsInt,
		[]string{"12", "123", "0", "-0"},
		[]string{"  ", "foo", "123.123", ""},
	},
	"IsFloat": TestCase{
		IsFloat,
		[]string{
			"123",
			"123.",
			"123.123",
			"-123.123",
			"-0.123",
			"0.123",
			".0",
			"01.123",
			"-0.22250738585072011e-307"},
		[]string{
			"foo",
			"  ",
			""},
	},
	"IsUUID": TestCase{
		IsUUID,
		[]string{
			"A987FBC9-4BED-3078-CF07-9141BA07C9F3",
			"A987FBC9-4BED-4078-8F07-9141BA07C9F3",
			"A987FBC9-4BED-5078-AF07-9141BA07C9F3"},
		[]string{
			"",
			"xxxA987FBC9-4BED-3078-CF07-9141BA07C9F3",
			"A987FBC9-4BED-3078-CF07-9141BA07C9F3xxx",
			"A987FBC94BED3078CF079141BA07C9F3",
			"934859",
			"987FBC9-4BED-3078-CF07A-9141BA07C9F3",
			"AAAAAAAA-1111-1111-AAAG-111111111111"},
	},
	"IsUUIDv3": TestCase{
		IsUUIDv3,
		[]string{
			"A987FBC9-4BED-3078-CF07-9141BA07C9F3",
		},
		[]string{
			"",
			"xxxA987FBC9-4BED-3078-CF07-9141BA07C9F3",
			"934859",
			"AAAAAAAA-1111-1111-AAAG-111111111111",
			"A987FBC9-4BED-4078-8F07-9141BA07C9F3",
			"A987FBC9-4BED-5078-AF07-9141BA07C9F3"},
	},
	"IsUUIDv4": TestCase{
		IsUUIDv4,
		[]string{
			"713ae7e3-cb32-45f9-adcb-7c4fa86b90c1",
			"625e63f3-58f5-40b7-83a1-a72ad31acffb",
			"57b73598-8764-4ad0-a76a-679bb6640eb1",
			"9c858901-8a57-4791-81fe-4c455b099bc9"},
		[]string{
			"",
			"xxxA987FBC9-4BED-3078-CF07-9141BA07C9F3",
			"934859",
			"AAAAAAAA-1111-1111-AAAG-111111111111",
			"A987FBC9-4BED-5078-AF07-9141BA07C9F3",
			"A987FBC9-4BED-3078-CF07-9141BA07C9F3"},
	},
	"IsUUIDv5": TestCase{
		IsUUIDv5,
		[]string{
			"987FBC97-4BED-5078-AF07-9141BA07C9F3",
			"987FBC97-4BED-5078-BF07-9141BA07C9F3",
			"987FBC97-4BED-5078-8F07-9141BA07C9F3",
			"987FBC97-4BED-5078-9F07-9141BA07C9F3"},
		[]string{
			"",
			"xxxA987FBC9-4BED-3078-CF07-9141BA07C9F3",
			"934859",
			"AAAAAAAA-1111-1111-AAAG-111111111111",
			"9c858901-8a57-4791-81fe-4c455b099bc9",
			"A987FBC9-4BED-3078-CF07-9141BA07C9F3"},
	},
	"IsAscii": TestCase{
		IsAscii,
		[]string{"foo", "0987654321", "test@example.com", "1234abcDEF"},
		[]string{"ｆｏｏbar", "ｘｙｚ０９８", "１２３456", "ｶﾀｶﾅ"},
	},
}

func runTest(t *testing.T, funcName string) {
	f := testCases[funcName].f

	// assert examples
	for _, in := range testCases[funcName].valid {
		if out := f(in); out != true {
			t.Errorf("%s(%q) = %t want true", funcName, in, out)
		}
	}

	// refute non examples
	for _, in := range testCases[funcName].invalid {
		if out := f(in); out != false {
			t.Errorf("%s(%s) = %t want false", funcName, in, out)
		}
	}
}

func TestIsAlpha(t *testing.T) { runTest(t, "IsAlpha") }

func TestIsAlphanumeric(t *testing.T) { runTest(t, "IsAlphanumeric") }

func TestIsAscii(t *testing.T) { runTest(t, "IsAscii") }

func TestIsBool(t *testing.T) { runTest(t, "IsBool") }

func TestIsEmail(t *testing.T) { runTest(t, "IsEmail") }

func TestIsFloat(t *testing.T) { runTest(t, "IsFloat") }

func TestIsHexadecimal(t *testing.T) { runTest(t, "IsHexadecimal") }

func TestIsHexColor(t *testing.T) { runTest(t, "IsHexColor") }

func TestIsInt(t *testing.T) { runTest(t, "IsInt") }

func TestIsISBN(t *testing.T) { runTest(t, "IsISBN") }

func TestIsISBNv10(t *testing.T) { runTest(t, "IsISBNv10") }

func TestIsISBNv13(t *testing.T) { runTest(t, "IsISBNv13") }

func TestIsLowerCase(t *testing.T) { runTest(t, "IsLowerCase") }

func TestIsNull(t *testing.T) { runTest(t, "IsNull") }

func TestIsUpperCase(t *testing.T) { runTest(t, "IsUpperCase") }

func TestIsUUID(t *testing.T) { runTest(t, "IsUUID") }

func TestIsUUIDv3(t *testing.T) { runTest(t, "IsUUIDv3") }

func TestIsUUIDv4(t *testing.T) { runTest(t, "IsUUIDv4") }

func TestIsUUIDv5(t *testing.T) { runTest(t, "IsUUIDv5") }

// func TestIsDate(t *testing.T) { runTest(t, "IsDate") }

func TestIsIN(t *testing.T) {
	marx_brother := "groucho"
	moe := "moe"
	stooges := []string{"moe", "larry", "curly"}

	if v := IsIn(marx_brother, stooges); v == true {
		t.Errorf("IsIn(%s, %s) = %v want false", marx_brother, stooges, v)
	}

	if v := IsIn(moe, stooges); v == false {
		t.Errorf("IsIn(%v, %v) = %v want false", moe, stooges, v)
	}

}
