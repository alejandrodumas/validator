// Package validator provides multiple string validations
package validator

import (
	"net"
	"regexp"
	"strconv"
	"strings"
)

const (
	creditCard      = `^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`
	email           = `(?i)[A-Z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[A-Z0-9!#$%&'*+/=?^_{|}~-]+)*@(?:[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?\.)+[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?`
	hexcolor        = `^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`
	possibleISBNv10 = `^(?:[0-9]{9}X|[0-9]{10})$`
	possibleISBNv13 = `^(?:[0-9]{13})$`
	uuidv3          = `(?i)^[0-9A-F]{8}-[0-9A-F]{4}-3[0-9A-F]{3}-[0-9A-F]{4}-[0-9A-F]{12}$`
	uuidv4          = `(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$`
	uuidv5          = `(?i)^[0-9A-F]{8}-[0-9A-F]{4}-5[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$`
	uuidvAll        = `(?i)^[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}$`
)

var (
	reAlpha        = regexp.MustCompile("^[[:alpha:]]+$")
	reAlphanumeric = regexp.MustCompile("^[[:alnum:]]+$")
	reAscii        = regexp.MustCompile("^[[:ascii:]]+$")
	reCreditCard   = regexp.MustCompile(creditCard)
	reEmail        = regexp.MustCompile(email)
	reHexadecimal  = regexp.MustCompile("^[[:xdigit:]]+$")
	reHexColor     = regexp.MustCompile(hexcolor)
	reISBNv10      = regexp.MustCompile(possibleISBNv10)
	reISBNv13      = regexp.MustCompile(possibleISBNv13)
	reUUIDv3       = regexp.MustCompile(uuidv3)
	reUUIDv4       = regexp.MustCompile(uuidv4)
	reUUIDv5       = regexp.MustCompile(uuidv5)
	reUUIDvAll     = regexp.MustCompile(uuidvAll)
)

var (
	sanitizeISBN       = regexp.MustCompile(`[\s-]+`)
	sanitizeCreditCard = regexp.MustCompile(`[^0-9]+`)
)

func IsAlpha(s string) bool { return reAlpha.MatchString(s) }

func IsAlphanumeric(s string) bool { return reAlphanumeric.MatchString(s) }

func IsAscii(s string) bool { return reAscii.MatchString(s) }

func IsEmail(s string) bool { return reEmail.MatchString(s) }

func IsBool(s string) bool {
	if _, err := strconv.ParseBool(s); err != nil {
		return false
	}

	return true
}

// TODO IsCreditCard
func IsCreditCard(s string) bool {
	sanitized := sanitizeCreditCard.ReplaceAllString(s, "")

	if !reCreditCard.MatchString(sanitized) {
		return false
	}

	return false
}

func IsInt(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	return true
}

func IsISBN(s string) bool { return IsISBNv10(s) || IsISBNv13(s) }

func IsISBNv10(s string) bool {
	sanitized := sanitizeISBN.ReplaceAllString(s, "")

	if !reISBNv10.MatchString(sanitized) {
		return false
	}

	checksum := 0

	for i, d := range sanitized[:9] {
		digit, _ := strconv.Atoi(string(d))
		checksum += (i + 1) * digit
	}

	if last := string(sanitized[9]); last == "X" {
		checksum += 100
	} else {
		digit, _ := strconv.Atoi(last)
		checksum += 10 * digit
	}

	if (checksum % 11) == 0 {
		return true
	}

	return false
}

func IsISBNv13(s string) bool {
	sanitized := sanitizeISBN.ReplaceAllString(s, "")

	if !reISBNv13.MatchString(sanitized) {
		return false
	}

	factor := []int{1, 3}
	checksum := 0
	checkdigit, _ := strconv.Atoi(string(sanitized[12]))

	for i, d := range sanitized[:12] {
		digit, _ := strconv.Atoi(string(d))
		checksum += factor[i%2] * digit
	}

	if result := ((10 - (checksum % 10)) % 10); checkdigit-result == 0 {
		return true
	}

	return false
}

// IsFloat
func IsFloat(s string) bool {
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return false
	}

	return true
}

// IsUrl
// func IsUrl(s string) bool {
// 	return len(s) < maxURLlen && reUrl.MatchString(s)
// }

// func IsHttpsUrl() {

// }

func IsIP(s string) bool {
	if net.ParseIP(s) == nil {
		return false
	}

	return true
}

func IsIPv4(s string) bool {
	if ip := net.ParseIP(s); ip != nil && ip.To4() != nil {
		return true
	}

	return false
}

func IsLength(s string) bool {
	return false
}

// func IsIPv6(s string) bool {
// 	if ip := net.ParseIP(s); ip != nil && ip.To1() != nil {
// 		return true
// 	}

// 	return false
// }

func IsDate(s string) bool {
	return false
}

// IsIn
func IsIn(s string, w []string) bool {
	for _, t := range w {
		if t == s {
			return true
		}
	}
	return false
}

func IsLowerCase(s string) bool { return strings.ToLower(s) == s }

func IsNull(s string) bool { return s == "" }

func IsHexColor(s string) bool { return reHexColor.MatchString(s) }

func IsHexadecimal(s string) bool { return reHexadecimal.MatchString(s) }

func IsUpperCase(s string) bool { return strings.ToUpper(s) == s }

func IsUUID(s string) bool { return reUUIDvAll.MatchString(s) }

func IsUUIDv3(s string) bool { return reUUIDv3.MatchString(s) }

func IsUUIDv4(s string) bool { return reUUIDv4.MatchString(s) }

func IsUUIDv5(s string) bool { return reUUIDv5.MatchString(s) }
