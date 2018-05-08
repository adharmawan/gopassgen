package gopassgen

/*
Version		: 0.2.0
Author		: Arijit Basu <sayanarijit@gmail.com>
Docs		: https://github.com/sayanarijit/gopassgen#README

Usage:
	p := gopassgen.NewPolicy()

	p.MinDigits = 5
	p.MinCapsAlpha = 2
	p.MinSpclChars = 2

	password := gopassgen.Generate(p)
*/

import (
	"math/rand"
	"time"
)

// Policy of password to be passed in Generate() function
type Policy struct {
	MinLength     int // Minimum length of password
	MaxLength     int // Maximum length of password
	MinCapsAlpha  int // Minimum length of capital letters
	MinSmallAlpha int // Minimum length of small letters
	MinDigits     int // Minimum length of digits
	MinSpclChars  int // Minimum length of special characters
}

// NewPolicy returns a default password policy which can be modified
func NewPolicy() Policy {
	p := Policy{
		MinLength:     6,
		MaxLength:     16,
		MinCapsAlpha:  0,
		MinSmallAlpha: 0,
		MinDigits:     0,
		MinSpclChars:  0,
	}
	return p
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// CreateRandom returns a random byte string of given length from given byte string
func CreateRandom(bs []byte, length int) []byte {
	filled := make([]byte, length)
	max := len(bs)

	for i := 0; i < length; i++ {
		Shuffle(bs)
		filled[i] = bs[random(0, max)]
	}

	return filled
}

// Shuffle the given byte string
func Shuffle(s []byte) {
	rand.Seed(time.Now().UnixNano())
	n := len(s)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

// Generate a new password based on given policy
func Generate(p Policy) string {

	// Character length based policies should not be negative
	if p.MinLength*p.MaxLength*p.MinCapsAlpha*p.MinSmallAlpha*p.MinDigits*p.MinSpclChars < 0 {
		panic("Character length should not ne negative")
	}

	// Max length should be greater than minimun length
	if p.MinLength > p.MaxLength {
		panic("Minimum length cannot be greater than maximum length")
	}

	// Max length should be sufficient to hold all minimum length policies
	if p.MaxLength > 0 {
		if p.MinCapsAlpha+p.MinSmallAlpha+p.MinDigits+p.MinSpclChars > p.MaxLength {
			panic("Maximum length is not sufficient")
		}
	} else {
		return ""
	}

	capsAlpha := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	smallAlpha := []byte("abcdefghijklmnopqrstuvwxyz")
	digits := []byte("0123456789")
	spclChars := []byte("!@#$%^&*()-_=+,.?/:;{}[]`~")
	allChars := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")

	passwd := CreateRandom(capsAlpha, p.MinCapsAlpha)

	passwd = append(passwd, CreateRandom(smallAlpha, p.MinSmallAlpha)...)
	passwd = append(passwd, CreateRandom(digits, p.MinDigits)...)
	passwd = append(passwd, CreateRandom(spclChars, p.MinSpclChars)...)

	requiredMore := p.MinLength - len(passwd)
	if requiredMore > 0 {
		passwd = append(passwd, CreateRandom(allChars, requiredMore)...)
	}

	Shuffle(passwd)

	return string(passwd)
}
