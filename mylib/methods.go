package mylib

import (
	"errors"
	"github.com/ttacon/libphonenumber"
	"strconv"
)

type Arith int

type MsisdnFormat struct {
	Cc  string // country dialing code
	Ndc string // mno identifier (national destination code)
	Sn  string // sucriber code
	Ci  string // country indentifier
}

var (
	ErrInvalidMSISNDnumber = errors.New("invalid MSISDN nubmer")
)

func (t *Arith) ParseMsisdn(msisdn *string, reply *MsisdnFormat) error {

	phoneNumber, _ := libphonenumber.Parse(*msisdn, "")

	if !libphonenumber.IsValidNumber(phoneNumber) {
		return ErrInvalidMSISNDnumber
	}

	reply.Cc = strconv.Itoa(int(phoneNumber.GetCountryCode()))
	nationalSignificantNumber := libphonenumber.GetNationalSignificantNumber(phoneNumber)
	nationalDestinationCodeLength := libphonenumber.GetLengthOfNationalDestinationCode(phoneNumber)
	if nationalDestinationCodeLength > 0 {
		reply.Ndc = nationalSignificantNumber[0:nationalDestinationCodeLength]
		reply.Sn = nationalSignificantNumber[nationalDestinationCodeLength:]
	} else {
		reply.Sn = nationalSignificantNumber
	}
	reply.Ci = libphonenumber.GetRegionCodeForNumber(phoneNumber)

	return nil
}
