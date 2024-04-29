package otp

import (
	"github.com/stretchr/testify/assert"
	"github.com/xlzd/gotp"
	"testing"
	"time"
)

// 功能类似库 github.com/pquerna/otp

var secret = "4S62BZNFXXSZLCRO"

// TOTP Time-based OTPs
func TestTOTP(t *testing.T) {
	totp := gotp.NewDefaultTOTP(secret)
	totpURL := totp.ProvisioningUri("tiptok", "gopp")
	t.Logf("%v\n%v\n%v", totpURL, totp.At(time.Now().Unix()), totp.Verify("065688", time.Now().Unix())) // otpauth://hotp/gopp:tiptok?counter=0&issuer=gopp&secret=4S62BZNFXXSZLCRO 密钥
}

// HOTP  Counter-based OTPs
func TestHOTP(t *testing.T) {
	hotp := gotp.NewDefaultHOTP("4S62BZNFXXSZLCRO")
	t.Log(hotp.At(0)) // '944181'
	t.Log(hotp.At(1)) // '770975'

	// OTP verified for a given timestamp
	t.Log(hotp.Verify("944181", 0)) // true
	t.Log(hotp.Verify("944181", 1)) // false

	// generate a provisioning uri
	t.Log(hotp.ProvisioningUri("demoAccountName", "issuerName", 1))
}

func TestValidCode(t *testing.T) {
	code := "944181"
	assert.True(t, ValidCode(code, secret))
}
