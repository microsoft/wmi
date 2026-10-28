// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// TeeDriverAmtAdtRecWmiInfo struct
type TeeDriverAmtAdtRecWmiInfo struct {
	*cim.WmiInstance

	//
	Active bool

	//
	AmtNvarIsValid bool

	// AMT certificate Serials
	CertificateSerials []TeeDriverAmtCertificateSerialsWmiInfo

	// AMT Host initiated manual or PKI
	HostInitiated uint32

	//
	InstanceName string

	// AMT IP address
	IPAddress string

	// AMT Is default root CA hash
	IsDefault bool

	// AMT Leaf cert Fingerprint
	LeafCertFingerprint []uint8

	// AMT Leaf expiration Date
	LeafExpirationDate []byte

	// AMT Leaf Hash Algorithm
	LeafHashAlgorithm uint32

	// AMT Leaf IssuerDN
	LeafIssuerDN []byte

	// AMT Leaf Subject DN
	LeafSubjectDN []byte

	// AMT Secure DNS
	SecureDNS uint32

	// AMT Provisoning Timestamp
	Timestamp uint32
}

func NewTeeDriverAmtAdtRecWmiInfoEx1(instance *cim.WmiInstance) (newInstance *TeeDriverAmtAdtRecWmiInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &TeeDriverAmtAdtRecWmiInfo{
		WmiInstance: tmp,
	}
	return
}

func NewTeeDriverAmtAdtRecWmiInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TeeDriverAmtAdtRecWmiInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TeeDriverAmtAdtRecWmiInfo{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetAmtNvarIsValid sets the value of AmtNvarIsValid for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) SetPropertyAmtNvarIsValid(value bool) (err error) {
	return instance.SetProperty("AmtNvarIsValid", (value))
}

// GetAmtNvarIsValid gets the value of AmtNvarIsValid for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) GetPropertyAmtNvarIsValid() (value bool, err error) {
	retValue, err := instance.GetProperty("AmtNvarIsValid")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetCertificateSerials sets the value of CertificateSerials for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) SetPropertyCertificateSerials(value []TeeDriverAmtCertificateSerialsWmiInfo) (err error) {
	return instance.SetProperty("CertificateSerials", (value))
}

// GetCertificateSerials gets the value of CertificateSerials for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) GetPropertyCertificateSerials() (value []TeeDriverAmtCertificateSerialsWmiInfo, err error) {
	retValue, err := instance.GetProperty("CertificateSerials")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(TeeDriverAmtCertificateSerialsWmiInfo)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " TeeDriverAmtCertificateSerialsWmiInfo is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, TeeDriverAmtCertificateSerialsWmiInfo(valuetmp))
	}

	return
}

// SetHostInitiated sets the value of HostInitiated for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) SetPropertyHostInitiated(value uint32) (err error) {
	return instance.SetProperty("HostInitiated", (value))
}

// GetHostInitiated gets the value of HostInitiated for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) GetPropertyHostInitiated() (value uint32, err error) {
	retValue, err := instance.GetProperty("HostInitiated")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetIPAddress sets the value of IPAddress for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) SetPropertyIPAddress(value string) (err error) {
	return instance.SetProperty("IPAddress", (value))
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) GetPropertyIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("IPAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetIsDefault sets the value of IsDefault for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) SetPropertyIsDefault(value bool) (err error) {
	return instance.SetProperty("IsDefault", (value))
}

// GetIsDefault gets the value of IsDefault for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) GetPropertyIsDefault() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDefault")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetLeafCertFingerprint sets the value of LeafCertFingerprint for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) SetPropertyLeafCertFingerprint(value []uint8) (err error) {
	return instance.SetProperty("LeafCertFingerprint", (value))
}

// GetLeafCertFingerprint gets the value of LeafCertFingerprint for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) GetPropertyLeafCertFingerprint() (value []uint8, err error) {
	retValue, err := instance.GetProperty("LeafCertFingerprint")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetLeafExpirationDate sets the value of LeafExpirationDate for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) SetPropertyLeafExpirationDate(value []byte) (err error) {
	return instance.SetProperty("LeafExpirationDate", (value))
}

// GetLeafExpirationDate gets the value of LeafExpirationDate for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) GetPropertyLeafExpirationDate() (value []byte, err error) {
	retValue, err := instance.GetProperty("LeafExpirationDate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(byte)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " byte is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, byte(valuetmp))
	}

	return
}

// SetLeafHashAlgorithm sets the value of LeafHashAlgorithm for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) SetPropertyLeafHashAlgorithm(value uint32) (err error) {
	return instance.SetProperty("LeafHashAlgorithm", (value))
}

// GetLeafHashAlgorithm gets the value of LeafHashAlgorithm for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) GetPropertyLeafHashAlgorithm() (value uint32, err error) {
	retValue, err := instance.GetProperty("LeafHashAlgorithm")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetLeafIssuerDN sets the value of LeafIssuerDN for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) SetPropertyLeafIssuerDN(value []byte) (err error) {
	return instance.SetProperty("LeafIssuerDN", (value))
}

// GetLeafIssuerDN gets the value of LeafIssuerDN for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) GetPropertyLeafIssuerDN() (value []byte, err error) {
	retValue, err := instance.GetProperty("LeafIssuerDN")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(byte)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " byte is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, byte(valuetmp))
	}

	return
}

// SetLeafSubjectDN sets the value of LeafSubjectDN for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) SetPropertyLeafSubjectDN(value []byte) (err error) {
	return instance.SetProperty("LeafSubjectDN", (value))
}

// GetLeafSubjectDN gets the value of LeafSubjectDN for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) GetPropertyLeafSubjectDN() (value []byte, err error) {
	retValue, err := instance.GetProperty("LeafSubjectDN")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(byte)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " byte is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, byte(valuetmp))
	}

	return
}

// SetSecureDNS sets the value of SecureDNS for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) SetPropertySecureDNS(value uint32) (err error) {
	return instance.SetProperty("SecureDNS", (value))
}

// GetSecureDNS gets the value of SecureDNS for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) GetPropertySecureDNS() (value uint32, err error) {
	retValue, err := instance.GetProperty("SecureDNS")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetTimestamp sets the value of Timestamp for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) SetPropertyTimestamp(value uint32) (err error) {
	return instance.SetProperty("Timestamp", (value))
}

// GetTimestamp gets the value of Timestamp for the instance
func (instance *TeeDriverAmtAdtRecWmiInfo) GetPropertyTimestamp() (value uint32, err error) {
	retValue, err := instance.GetProperty("Timestamp")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
