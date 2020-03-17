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

// Msvm_ReplicationServiceSettingData struct
type Msvm_ReplicationServiceSettingData struct {
	CIM_SettingData

	//
	AllowedAuthenticationType uint16

	//
	CertificateThumbPrint string

	//
	HttpPort uint16

	//
	HttpsPort uint16

	//
	MonitoringInterval uint32

	//
	MonitoringStartTime string

	//
	RecoveryServerEnabled bool
}

// SetAllowedAuthenticationType sets the value of AllowedAuthenticationType for the instance
func (instance *Msvm_ReplicationServiceSettingData) SetPropertyAllowedAuthenticationType(value uint16) (err error) {
	return instance.SetProperty("AllowedAuthenticationType", value)
}

// GetAllowedAuthenticationType gets the value of AllowedAuthenticationType for the instance
func (instance *Msvm_ReplicationServiceSettingData) GetPropertyAllowedAuthenticationType() (value uint16, err error) {
	retValue, err := instance.GetProperty("AllowedAuthenticationType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCertificateThumbPrint sets the value of CertificateThumbPrint for the instance
func (instance *Msvm_ReplicationServiceSettingData) SetPropertyCertificateThumbPrint(value string) (err error) {
	return instance.SetProperty("CertificateThumbPrint", value)
}

// GetCertificateThumbPrint gets the value of CertificateThumbPrint for the instance
func (instance *Msvm_ReplicationServiceSettingData) GetPropertyCertificateThumbPrint() (value string, err error) {
	retValue, err := instance.GetProperty("CertificateThumbPrint")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHttpPort sets the value of HttpPort for the instance
func (instance *Msvm_ReplicationServiceSettingData) SetPropertyHttpPort(value uint16) (err error) {
	return instance.SetProperty("HttpPort", value)
}

// GetHttpPort gets the value of HttpPort for the instance
func (instance *Msvm_ReplicationServiceSettingData) GetPropertyHttpPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("HttpPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHttpsPort sets the value of HttpsPort for the instance
func (instance *Msvm_ReplicationServiceSettingData) SetPropertyHttpsPort(value uint16) (err error) {
	return instance.SetProperty("HttpsPort", value)
}

// GetHttpsPort gets the value of HttpsPort for the instance
func (instance *Msvm_ReplicationServiceSettingData) GetPropertyHttpsPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("HttpsPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMonitoringInterval sets the value of MonitoringInterval for the instance
func (instance *Msvm_ReplicationServiceSettingData) SetPropertyMonitoringInterval(value uint32) (err error) {
	return instance.SetProperty("MonitoringInterval", value)
}

// GetMonitoringInterval gets the value of MonitoringInterval for the instance
func (instance *Msvm_ReplicationServiceSettingData) GetPropertyMonitoringInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("MonitoringInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMonitoringStartTime sets the value of MonitoringStartTime for the instance
func (instance *Msvm_ReplicationServiceSettingData) SetPropertyMonitoringStartTime(value string) (err error) {
	return instance.SetProperty("MonitoringStartTime", value)
}

// GetMonitoringStartTime gets the value of MonitoringStartTime for the instance
func (instance *Msvm_ReplicationServiceSettingData) GetPropertyMonitoringStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("MonitoringStartTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecoveryServerEnabled sets the value of RecoveryServerEnabled for the instance
func (instance *Msvm_ReplicationServiceSettingData) SetPropertyRecoveryServerEnabled(value bool) (err error) {
	return instance.SetProperty("RecoveryServerEnabled", value)
}

// GetRecoveryServerEnabled gets the value of RecoveryServerEnabled for the instance
func (instance *Msvm_ReplicationServiceSettingData) GetPropertyRecoveryServerEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("RecoveryServerEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_ReplicationServiceSettingData) GetRelatedReplicationService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ReplicationService")
}
