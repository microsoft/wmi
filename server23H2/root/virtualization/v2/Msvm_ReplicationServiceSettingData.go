// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_ReplicationServiceSettingData struct
type Msvm_ReplicationServiceSettingData struct {
	*CIM_SettingData

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

func NewMsvm_ReplicationServiceSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_ReplicationServiceSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReplicationServiceSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_ReplicationServiceSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ReplicationServiceSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReplicationServiceSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetAllowedAuthenticationType sets the value of AllowedAuthenticationType for the instance
func (instance *Msvm_ReplicationServiceSettingData) SetPropertyAllowedAuthenticationType(value uint16) (err error) {
	return instance.SetProperty("AllowedAuthenticationType", (value))
}

// GetAllowedAuthenticationType gets the value of AllowedAuthenticationType for the instance
func (instance *Msvm_ReplicationServiceSettingData) GetPropertyAllowedAuthenticationType() (value uint16, err error) {
	retValue, err := instance.GetProperty("AllowedAuthenticationType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetCertificateThumbPrint sets the value of CertificateThumbPrint for the instance
func (instance *Msvm_ReplicationServiceSettingData) SetPropertyCertificateThumbPrint(value string) (err error) {
	return instance.SetProperty("CertificateThumbPrint", (value))
}

// GetCertificateThumbPrint gets the value of CertificateThumbPrint for the instance
func (instance *Msvm_ReplicationServiceSettingData) GetPropertyCertificateThumbPrint() (value string, err error) {
	retValue, err := instance.GetProperty("CertificateThumbPrint")
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

// SetHttpPort sets the value of HttpPort for the instance
func (instance *Msvm_ReplicationServiceSettingData) SetPropertyHttpPort(value uint16) (err error) {
	return instance.SetProperty("HttpPort", (value))
}

// GetHttpPort gets the value of HttpPort for the instance
func (instance *Msvm_ReplicationServiceSettingData) GetPropertyHttpPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("HttpPort")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetHttpsPort sets the value of HttpsPort for the instance
func (instance *Msvm_ReplicationServiceSettingData) SetPropertyHttpsPort(value uint16) (err error) {
	return instance.SetProperty("HttpsPort", (value))
}

// GetHttpsPort gets the value of HttpsPort for the instance
func (instance *Msvm_ReplicationServiceSettingData) GetPropertyHttpsPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("HttpsPort")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetMonitoringInterval sets the value of MonitoringInterval for the instance
func (instance *Msvm_ReplicationServiceSettingData) SetPropertyMonitoringInterval(value uint32) (err error) {
	return instance.SetProperty("MonitoringInterval", (value))
}

// GetMonitoringInterval gets the value of MonitoringInterval for the instance
func (instance *Msvm_ReplicationServiceSettingData) GetPropertyMonitoringInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("MonitoringInterval")
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

// SetMonitoringStartTime sets the value of MonitoringStartTime for the instance
func (instance *Msvm_ReplicationServiceSettingData) SetPropertyMonitoringStartTime(value string) (err error) {
	return instance.SetProperty("MonitoringStartTime", (value))
}

// GetMonitoringStartTime gets the value of MonitoringStartTime for the instance
func (instance *Msvm_ReplicationServiceSettingData) GetPropertyMonitoringStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("MonitoringStartTime")
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

// SetRecoveryServerEnabled sets the value of RecoveryServerEnabled for the instance
func (instance *Msvm_ReplicationServiceSettingData) SetPropertyRecoveryServerEnabled(value bool) (err error) {
	return instance.SetProperty("RecoveryServerEnabled", (value))
}

// GetRecoveryServerEnabled gets the value of RecoveryServerEnabled for the instance
func (instance *Msvm_ReplicationServiceSettingData) GetPropertyRecoveryServerEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("RecoveryServerEnabled")
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
func (instance *Msvm_ReplicationServiceSettingData) GetRelatedReplicationService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ReplicationService")
}
