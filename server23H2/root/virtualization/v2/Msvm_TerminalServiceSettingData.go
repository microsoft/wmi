// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_TerminalServiceSettingData struct
type Msvm_TerminalServiceSettingData struct {
	*CIM_SettingData

	//
	AllowedHashAlgorithms []string

	//
	AuthCertificateHash []uint8

	//
	DisableSelfSignedCertificateGeneration bool

	//
	ListenerPort uint32

	//
	TrustedIssuerCertificateHashes []string
}

func NewMsvm_TerminalServiceSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_TerminalServiceSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_TerminalServiceSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_TerminalServiceSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_TerminalServiceSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_TerminalServiceSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetAllowedHashAlgorithms sets the value of AllowedHashAlgorithms for the instance
func (instance *Msvm_TerminalServiceSettingData) SetPropertyAllowedHashAlgorithms(value []string) (err error) {
	return instance.SetProperty("AllowedHashAlgorithms", (value))
}

// GetAllowedHashAlgorithms gets the value of AllowedHashAlgorithms for the instance
func (instance *Msvm_TerminalServiceSettingData) GetPropertyAllowedHashAlgorithms() (value []string, err error) {
	retValue, err := instance.GetProperty("AllowedHashAlgorithms")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetAuthCertificateHash sets the value of AuthCertificateHash for the instance
func (instance *Msvm_TerminalServiceSettingData) SetPropertyAuthCertificateHash(value []uint8) (err error) {
	return instance.SetProperty("AuthCertificateHash", (value))
}

// GetAuthCertificateHash gets the value of AuthCertificateHash for the instance
func (instance *Msvm_TerminalServiceSettingData) GetPropertyAuthCertificateHash() (value []uint8, err error) {
	retValue, err := instance.GetProperty("AuthCertificateHash")
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

// SetDisableSelfSignedCertificateGeneration sets the value of DisableSelfSignedCertificateGeneration for the instance
func (instance *Msvm_TerminalServiceSettingData) SetPropertyDisableSelfSignedCertificateGeneration(value bool) (err error) {
	return instance.SetProperty("DisableSelfSignedCertificateGeneration", (value))
}

// GetDisableSelfSignedCertificateGeneration gets the value of DisableSelfSignedCertificateGeneration for the instance
func (instance *Msvm_TerminalServiceSettingData) GetPropertyDisableSelfSignedCertificateGeneration() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableSelfSignedCertificateGeneration")
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

// SetListenerPort sets the value of ListenerPort for the instance
func (instance *Msvm_TerminalServiceSettingData) SetPropertyListenerPort(value uint32) (err error) {
	return instance.SetProperty("ListenerPort", (value))
}

// GetListenerPort gets the value of ListenerPort for the instance
func (instance *Msvm_TerminalServiceSettingData) GetPropertyListenerPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("ListenerPort")
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

// SetTrustedIssuerCertificateHashes sets the value of TrustedIssuerCertificateHashes for the instance
func (instance *Msvm_TerminalServiceSettingData) SetPropertyTrustedIssuerCertificateHashes(value []string) (err error) {
	return instance.SetProperty("TrustedIssuerCertificateHashes", (value))
}

// GetTrustedIssuerCertificateHashes gets the value of TrustedIssuerCertificateHashes for the instance
func (instance *Msvm_TerminalServiceSettingData) GetPropertyTrustedIssuerCertificateHashes() (value []string, err error) {
	retValue, err := instance.GetProperty("TrustedIssuerCertificateHashes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
func (instance *Msvm_TerminalServiceSettingData) GetRelatedTerminalService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_TerminalService")
}
