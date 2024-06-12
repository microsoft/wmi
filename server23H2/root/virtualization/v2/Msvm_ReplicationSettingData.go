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

// Msvm_ReplicationSettingData struct
type Msvm_ReplicationSettingData struct {
	*CIM_VirtualSystemSettingData

	//
	AdditionalSettings string

	//
	ApplicationConsistentSnapshotInterval uint16

	//
	AuthenticationType uint16

	//
	AutoResynchronizeEnabled bool

	//
	AutoResynchronizeIntervalEnd string

	//
	AutoResynchronizeIntervalStart string

	//
	BypassProxyServer bool

	//
	CertificateThumbPrint string

	//
	CompressionEnabled bool

	//
	EnableWriteOrderPreservationAcrossDisks bool

	// The list of VHD attached to the ComputerSystem that will be replicated by the Failover Replication Engine. This is an array of strings each containing an the InstanceID of the resource allocation setting data (RASD) of the VHD.
	IncludedDisks []string

	//
	PrimaryConnectionPoint string

	//
	PrimaryHostSystem string

	//
	RecoveryConnectionPoint string

	//
	RecoveryHistory uint16

	//
	RecoveryHostSystem string

	//
	RecoveryServerPortNumber uint16

	//
	ReplicateHostKvpItems bool

	//
	ReplicationInterval uint16

	//
	ReplicationProvider string

	//
	RootCertificateThumbPrint string
}

func NewMsvm_ReplicationSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_ReplicationSettingData, err error) {
	tmp, err := NewCIM_VirtualSystemSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReplicationSettingData{
		CIM_VirtualSystemSettingData: tmp,
	}
	return
}

func NewMsvm_ReplicationSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ReplicationSettingData, err error) {
	tmp, err := NewCIM_VirtualSystemSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReplicationSettingData{
		CIM_VirtualSystemSettingData: tmp,
	}
	return
}

// SetAdditionalSettings sets the value of AdditionalSettings for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyAdditionalSettings(value string) (err error) {
	return instance.SetProperty("AdditionalSettings", (value))
}

// GetAdditionalSettings gets the value of AdditionalSettings for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyAdditionalSettings() (value string, err error) {
	retValue, err := instance.GetProperty("AdditionalSettings")
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

// SetApplicationConsistentSnapshotInterval sets the value of ApplicationConsistentSnapshotInterval for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyApplicationConsistentSnapshotInterval(value uint16) (err error) {
	return instance.SetProperty("ApplicationConsistentSnapshotInterval", (value))
}

// GetApplicationConsistentSnapshotInterval gets the value of ApplicationConsistentSnapshotInterval for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyApplicationConsistentSnapshotInterval() (value uint16, err error) {
	retValue, err := instance.GetProperty("ApplicationConsistentSnapshotInterval")
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

// SetAuthenticationType sets the value of AuthenticationType for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyAuthenticationType(value uint16) (err error) {
	return instance.SetProperty("AuthenticationType", (value))
}

// GetAuthenticationType gets the value of AuthenticationType for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyAuthenticationType() (value uint16, err error) {
	retValue, err := instance.GetProperty("AuthenticationType")
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

// SetAutoResynchronizeEnabled sets the value of AutoResynchronizeEnabled for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyAutoResynchronizeEnabled(value bool) (err error) {
	return instance.SetProperty("AutoResynchronizeEnabled", (value))
}

// GetAutoResynchronizeEnabled gets the value of AutoResynchronizeEnabled for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyAutoResynchronizeEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoResynchronizeEnabled")
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

// SetAutoResynchronizeIntervalEnd sets the value of AutoResynchronizeIntervalEnd for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyAutoResynchronizeIntervalEnd(value string) (err error) {
	return instance.SetProperty("AutoResynchronizeIntervalEnd", (value))
}

// GetAutoResynchronizeIntervalEnd gets the value of AutoResynchronizeIntervalEnd for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyAutoResynchronizeIntervalEnd() (value string, err error) {
	retValue, err := instance.GetProperty("AutoResynchronizeIntervalEnd")
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

// SetAutoResynchronizeIntervalStart sets the value of AutoResynchronizeIntervalStart for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyAutoResynchronizeIntervalStart(value string) (err error) {
	return instance.SetProperty("AutoResynchronizeIntervalStart", (value))
}

// GetAutoResynchronizeIntervalStart gets the value of AutoResynchronizeIntervalStart for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyAutoResynchronizeIntervalStart() (value string, err error) {
	retValue, err := instance.GetProperty("AutoResynchronizeIntervalStart")
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

// SetBypassProxyServer sets the value of BypassProxyServer for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyBypassProxyServer(value bool) (err error) {
	return instance.SetProperty("BypassProxyServer", (value))
}

// GetBypassProxyServer gets the value of BypassProxyServer for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyBypassProxyServer() (value bool, err error) {
	retValue, err := instance.GetProperty("BypassProxyServer")
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

// SetCertificateThumbPrint sets the value of CertificateThumbPrint for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyCertificateThumbPrint(value string) (err error) {
	return instance.SetProperty("CertificateThumbPrint", (value))
}

// GetCertificateThumbPrint gets the value of CertificateThumbPrint for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyCertificateThumbPrint() (value string, err error) {
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

// SetCompressionEnabled sets the value of CompressionEnabled for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyCompressionEnabled(value bool) (err error) {
	return instance.SetProperty("CompressionEnabled", (value))
}

// GetCompressionEnabled gets the value of CompressionEnabled for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyCompressionEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("CompressionEnabled")
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

// SetEnableWriteOrderPreservationAcrossDisks sets the value of EnableWriteOrderPreservationAcrossDisks for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyEnableWriteOrderPreservationAcrossDisks(value bool) (err error) {
	return instance.SetProperty("EnableWriteOrderPreservationAcrossDisks", (value))
}

// GetEnableWriteOrderPreservationAcrossDisks gets the value of EnableWriteOrderPreservationAcrossDisks for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyEnableWriteOrderPreservationAcrossDisks() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableWriteOrderPreservationAcrossDisks")
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

// SetIncludedDisks sets the value of IncludedDisks for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyIncludedDisks(value []string) (err error) {
	return instance.SetProperty("IncludedDisks", (value))
}

// GetIncludedDisks gets the value of IncludedDisks for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyIncludedDisks() (value []string, err error) {
	retValue, err := instance.GetProperty("IncludedDisks")
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

// SetPrimaryConnectionPoint sets the value of PrimaryConnectionPoint for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyPrimaryConnectionPoint(value string) (err error) {
	return instance.SetProperty("PrimaryConnectionPoint", (value))
}

// GetPrimaryConnectionPoint gets the value of PrimaryConnectionPoint for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyPrimaryConnectionPoint() (value string, err error) {
	retValue, err := instance.GetProperty("PrimaryConnectionPoint")
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

// SetPrimaryHostSystem sets the value of PrimaryHostSystem for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyPrimaryHostSystem(value string) (err error) {
	return instance.SetProperty("PrimaryHostSystem", (value))
}

// GetPrimaryHostSystem gets the value of PrimaryHostSystem for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyPrimaryHostSystem() (value string, err error) {
	retValue, err := instance.GetProperty("PrimaryHostSystem")
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

// SetRecoveryConnectionPoint sets the value of RecoveryConnectionPoint for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyRecoveryConnectionPoint(value string) (err error) {
	return instance.SetProperty("RecoveryConnectionPoint", (value))
}

// GetRecoveryConnectionPoint gets the value of RecoveryConnectionPoint for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyRecoveryConnectionPoint() (value string, err error) {
	retValue, err := instance.GetProperty("RecoveryConnectionPoint")
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

// SetRecoveryHistory sets the value of RecoveryHistory for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyRecoveryHistory(value uint16) (err error) {
	return instance.SetProperty("RecoveryHistory", (value))
}

// GetRecoveryHistory gets the value of RecoveryHistory for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyRecoveryHistory() (value uint16, err error) {
	retValue, err := instance.GetProperty("RecoveryHistory")
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

// SetRecoveryHostSystem sets the value of RecoveryHostSystem for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyRecoveryHostSystem(value string) (err error) {
	return instance.SetProperty("RecoveryHostSystem", (value))
}

// GetRecoveryHostSystem gets the value of RecoveryHostSystem for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyRecoveryHostSystem() (value string, err error) {
	retValue, err := instance.GetProperty("RecoveryHostSystem")
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

// SetRecoveryServerPortNumber sets the value of RecoveryServerPortNumber for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyRecoveryServerPortNumber(value uint16) (err error) {
	return instance.SetProperty("RecoveryServerPortNumber", (value))
}

// GetRecoveryServerPortNumber gets the value of RecoveryServerPortNumber for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyRecoveryServerPortNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("RecoveryServerPortNumber")
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

// SetReplicateHostKvpItems sets the value of ReplicateHostKvpItems for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyReplicateHostKvpItems(value bool) (err error) {
	return instance.SetProperty("ReplicateHostKvpItems", (value))
}

// GetReplicateHostKvpItems gets the value of ReplicateHostKvpItems for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyReplicateHostKvpItems() (value bool, err error) {
	retValue, err := instance.GetProperty("ReplicateHostKvpItems")
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

// SetReplicationInterval sets the value of ReplicationInterval for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyReplicationInterval(value uint16) (err error) {
	return instance.SetProperty("ReplicationInterval", (value))
}

// GetReplicationInterval gets the value of ReplicationInterval for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyReplicationInterval() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationInterval")
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

// SetReplicationProvider sets the value of ReplicationProvider for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyReplicationProvider(value string) (err error) {
	return instance.SetProperty("ReplicationProvider", (value))
}

// GetReplicationProvider gets the value of ReplicationProvider for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyReplicationProvider() (value string, err error) {
	retValue, err := instance.GetProperty("ReplicationProvider")
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

// SetRootCertificateThumbPrint sets the value of RootCertificateThumbPrint for the instance
func (instance *Msvm_ReplicationSettingData) SetPropertyRootCertificateThumbPrint(value string) (err error) {
	return instance.SetProperty("RootCertificateThumbPrint", (value))
}

// GetRootCertificateThumbPrint gets the value of RootCertificateThumbPrint for the instance
func (instance *Msvm_ReplicationSettingData) GetPropertyRootCertificateThumbPrint() (value string, err error) {
	retValue, err := instance.GetProperty("RootCertificateThumbPrint")
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
