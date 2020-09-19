// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_CollectionReplicationSettingData struct
type Msvm_CollectionReplicationSettingData struct {
	*Msvm_CollectionSettingData

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
	IncludedDisks []string

	//
	PrimaryConnectionPoint string

	//
	PrimaryReplicationEntityIDs []string

	//
	RecoveryConnectionPoint string

	//
	RecoveryHistory uint16

	//
	RecoveryServerHosts []string

	//
	RecoveryServerPortNumber uint16

	//
	ReplicateHostKvpItems bool

	//
	ReplicationInterval uint16

	//
	RootCertificateThumbPrint string
}

func NewMsvm_CollectionReplicationSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_CollectionReplicationSettingData, err error) {
	tmp, err := NewMsvm_CollectionSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectionReplicationSettingData{
		Msvm_CollectionSettingData: tmp,
	}
	return
}

func NewMsvm_CollectionReplicationSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_CollectionReplicationSettingData, err error) {
	tmp, err := NewMsvm_CollectionSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectionReplicationSettingData{
		Msvm_CollectionSettingData: tmp,
	}
	return
}

// SetApplicationConsistentSnapshotInterval sets the value of ApplicationConsistentSnapshotInterval for the instance
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyApplicationConsistentSnapshotInterval(value uint16) (err error) {
	return instance.SetProperty("ApplicationConsistentSnapshotInterval", (value))
}

// GetApplicationConsistentSnapshotInterval gets the value of ApplicationConsistentSnapshotInterval for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyApplicationConsistentSnapshotInterval() (value uint16, err error) {
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
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyAuthenticationType(value uint16) (err error) {
	return instance.SetProperty("AuthenticationType", (value))
}

// GetAuthenticationType gets the value of AuthenticationType for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyAuthenticationType() (value uint16, err error) {
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
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyAutoResynchronizeEnabled(value bool) (err error) {
	return instance.SetProperty("AutoResynchronizeEnabled", (value))
}

// GetAutoResynchronizeEnabled gets the value of AutoResynchronizeEnabled for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyAutoResynchronizeEnabled() (value bool, err error) {
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
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyAutoResynchronizeIntervalEnd(value string) (err error) {
	return instance.SetProperty("AutoResynchronizeIntervalEnd", (value))
}

// GetAutoResynchronizeIntervalEnd gets the value of AutoResynchronizeIntervalEnd for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyAutoResynchronizeIntervalEnd() (value string, err error) {
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
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyAutoResynchronizeIntervalStart(value string) (err error) {
	return instance.SetProperty("AutoResynchronizeIntervalStart", (value))
}

// GetAutoResynchronizeIntervalStart gets the value of AutoResynchronizeIntervalStart for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyAutoResynchronizeIntervalStart() (value string, err error) {
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
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyBypassProxyServer(value bool) (err error) {
	return instance.SetProperty("BypassProxyServer", (value))
}

// GetBypassProxyServer gets the value of BypassProxyServer for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyBypassProxyServer() (value bool, err error) {
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
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyCertificateThumbPrint(value string) (err error) {
	return instance.SetProperty("CertificateThumbPrint", (value))
}

// GetCertificateThumbPrint gets the value of CertificateThumbPrint for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyCertificateThumbPrint() (value string, err error) {
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
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyCompressionEnabled(value bool) (err error) {
	return instance.SetProperty("CompressionEnabled", (value))
}

// GetCompressionEnabled gets the value of CompressionEnabled for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyCompressionEnabled() (value bool, err error) {
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

// SetIncludedDisks sets the value of IncludedDisks for the instance
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyIncludedDisks(value []string) (err error) {
	return instance.SetProperty("IncludedDisks", (value))
}

// GetIncludedDisks gets the value of IncludedDisks for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyIncludedDisks() (value []string, err error) {
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
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyPrimaryConnectionPoint(value string) (err error) {
	return instance.SetProperty("PrimaryConnectionPoint", (value))
}

// GetPrimaryConnectionPoint gets the value of PrimaryConnectionPoint for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyPrimaryConnectionPoint() (value string, err error) {
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

// SetPrimaryReplicationEntityIDs sets the value of PrimaryReplicationEntityIDs for the instance
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyPrimaryReplicationEntityIDs(value []string) (err error) {
	return instance.SetProperty("PrimaryReplicationEntityIDs", (value))
}

// GetPrimaryReplicationEntityIDs gets the value of PrimaryReplicationEntityIDs for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyPrimaryReplicationEntityIDs() (value []string, err error) {
	retValue, err := instance.GetProperty("PrimaryReplicationEntityIDs")
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

// SetRecoveryConnectionPoint sets the value of RecoveryConnectionPoint for the instance
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyRecoveryConnectionPoint(value string) (err error) {
	return instance.SetProperty("RecoveryConnectionPoint", (value))
}

// GetRecoveryConnectionPoint gets the value of RecoveryConnectionPoint for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyRecoveryConnectionPoint() (value string, err error) {
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
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyRecoveryHistory(value uint16) (err error) {
	return instance.SetProperty("RecoveryHistory", (value))
}

// GetRecoveryHistory gets the value of RecoveryHistory for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyRecoveryHistory() (value uint16, err error) {
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

// SetRecoveryServerHosts sets the value of RecoveryServerHosts for the instance
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyRecoveryServerHosts(value []string) (err error) {
	return instance.SetProperty("RecoveryServerHosts", (value))
}

// GetRecoveryServerHosts gets the value of RecoveryServerHosts for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyRecoveryServerHosts() (value []string, err error) {
	retValue, err := instance.GetProperty("RecoveryServerHosts")
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

// SetRecoveryServerPortNumber sets the value of RecoveryServerPortNumber for the instance
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyRecoveryServerPortNumber(value uint16) (err error) {
	return instance.SetProperty("RecoveryServerPortNumber", (value))
}

// GetRecoveryServerPortNumber gets the value of RecoveryServerPortNumber for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyRecoveryServerPortNumber() (value uint16, err error) {
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
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyReplicateHostKvpItems(value bool) (err error) {
	return instance.SetProperty("ReplicateHostKvpItems", (value))
}

// GetReplicateHostKvpItems gets the value of ReplicateHostKvpItems for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyReplicateHostKvpItems() (value bool, err error) {
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
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyReplicationInterval(value uint16) (err error) {
	return instance.SetProperty("ReplicationInterval", (value))
}

// GetReplicationInterval gets the value of ReplicationInterval for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyReplicationInterval() (value uint16, err error) {
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

// SetRootCertificateThumbPrint sets the value of RootCertificateThumbPrint for the instance
func (instance *Msvm_CollectionReplicationSettingData) SetPropertyRootCertificateThumbPrint(value string) (err error) {
	return instance.SetProperty("RootCertificateThumbPrint", (value))
}

// GetRootCertificateThumbPrint gets the value of RootCertificateThumbPrint for the instance
func (instance *Msvm_CollectionReplicationSettingData) GetPropertyRootCertificateThumbPrint() (value string, err error) {
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
