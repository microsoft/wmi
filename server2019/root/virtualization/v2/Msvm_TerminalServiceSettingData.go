// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_TerminalServiceSettingData struct
type Msvm_TerminalServiceSettingData struct {
	CIM_SettingData

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

// SetAllowedHashAlgorithms sets the value of AllowedHashAlgorithms for the instance
func (instance *Msvm_TerminalServiceSettingData) SetPropertyAllowedHashAlgorithms(value []string) (err error) {
	return instance.SetProperty("AllowedHashAlgorithms", value)
}

// GetAllowedHashAlgorithms gets the value of AllowedHashAlgorithms for the instance
func (instance *Msvm_TerminalServiceSettingData) GetPropertyAllowedHashAlgorithms() (value []string, err error) {
	retValue, err := instance.GetProperty("AllowedHashAlgorithms")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAuthCertificateHash sets the value of AuthCertificateHash for the instance
func (instance *Msvm_TerminalServiceSettingData) SetPropertyAuthCertificateHash(value []uint8) (err error) {
	return instance.SetProperty("AuthCertificateHash", value)
}

// GetAuthCertificateHash gets the value of AuthCertificateHash for the instance
func (instance *Msvm_TerminalServiceSettingData) GetPropertyAuthCertificateHash() (value []uint8, err error) {
	retValue, err := instance.GetProperty("AuthCertificateHash")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableSelfSignedCertificateGeneration sets the value of DisableSelfSignedCertificateGeneration for the instance
func (instance *Msvm_TerminalServiceSettingData) SetPropertyDisableSelfSignedCertificateGeneration(value bool) (err error) {
	return instance.SetProperty("DisableSelfSignedCertificateGeneration", value)
}

// GetDisableSelfSignedCertificateGeneration gets the value of DisableSelfSignedCertificateGeneration for the instance
func (instance *Msvm_TerminalServiceSettingData) GetPropertyDisableSelfSignedCertificateGeneration() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableSelfSignedCertificateGeneration")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetListenerPort sets the value of ListenerPort for the instance
func (instance *Msvm_TerminalServiceSettingData) SetPropertyListenerPort(value uint32) (err error) {
	return instance.SetProperty("ListenerPort", value)
}

// GetListenerPort gets the value of ListenerPort for the instance
func (instance *Msvm_TerminalServiceSettingData) GetPropertyListenerPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("ListenerPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedIssuerCertificateHashes sets the value of TrustedIssuerCertificateHashes for the instance
func (instance *Msvm_TerminalServiceSettingData) SetPropertyTrustedIssuerCertificateHashes(value []string) (err error) {
	return instance.SetProperty("TrustedIssuerCertificateHashes", value)
}

// GetTrustedIssuerCertificateHashes gets the value of TrustedIssuerCertificateHashes for the instance
func (instance *Msvm_TerminalServiceSettingData) GetPropertyTrustedIssuerCertificateHashes() (value []string, err error) {
	retValue, err := instance.GetProperty("TrustedIssuerCertificateHashes")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_TerminalServiceSettingData) GetRelatedTerminalService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_TerminalService")
}
