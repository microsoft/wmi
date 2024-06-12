// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_VirtualSystemMigrationSettingData struct
type CIM_VirtualSystemMigrationSettingData struct {
	*CIM_SettingData

	// Bandwidth indicates the bandwidth assigned to or requested for a virtual system migration operation. The special value 0 indicates:
	///- in migration requests the default bandwidth
	///- otherwise that bandwidths are not supported.
	///Bandwidth and Priority may be used in conjunction. Migration processes that have the highest equal priority value share the available bandwidth based on their requested bandwidth. If not all bandwidth is consumed by this set of processes, migration processes with the next lower equal priority share the remaining bandwidth. If still more bandwidth remains, migration processes with the next lower equal priority are considered, and so forth.
	///The unit applicable for the Bandwidth property is conveyed by the value of the BandwidthUnit property. If the value of the BandwidthUnit property matches "percent", the following restrictions apply:
	///The value of the Bandwidth property shall be between 0 and 100, with higher values indicating a higher bandwidth. A value of 100 indicates the total available bandwidth for performing virtual system migration operations. Values between 1 and 100 should linearly correlate with the available bandwidth range. For exampe, a value of 50 should request half of the available bandwidth, a value of 33 should request one third of the bandwidth, etc. .
	Bandwidth uint16

	// This property specifies the unit used by the Bandwidth property. The value of this property shall be a legal value of the Programmatic Units qualifier as defined in Appendix C.1 of DSP0004 V2.4 or later.
	///NOTE: Profiles like DMTF DSP1081 define means by that clients are enabled to discover the set of units supported by an implementation, along with ranges and increments for admissable values of the Bandwidth property.
	BandwidthUnit string

	// MigrationType describes a type of migration operation to be performed.
	///A value of 2 - Virtual System is to be migrated in a 'live' manner such that the running of the Virtual System is minimally impacted during the move.
	///A value of 3 - Virtual System will be temporarily paused prior to migration and then resume running after it is moved.
	///A value of 4 - The Virtual System will be quiesced to a stopped state prior to migration and then restarted after it is moved.
	MigrationType VirtualSystemMigrationSettingData_MigrationType

	// OtherTransportType indicates the type of transport to be applied if the value of TransportType is 1 (Other).
	OtherTransportType string

	// Priority specifies a relative migration importance which the virtual system migration implementation may use to order or otherwise give preference among multiple pending migration requests. The lower the value, the higher the priority. A value of 0 indicates:
	///- in migration requests the default priority
	///- otherwise that priorities are not supported
	Priority uint16

	// TransportType indicates the type of transport to be applied for a virtual system migration operation.
	///- 0(Unknown) indicates that the transport type is not exposed.- 1(Other) indicates that the transport type is specified as a textual value of the OtherTransportType property.
	///- 2(SSH) indicates the secure shell transport type.
	///- 3(TLS) indicats the transport layer security transport type.
	///- 4(TLS strict) indicats the transport layer security transport type with mutual authentication.
	///- 5(TCP) indicates the TCP transport type.
	///- 6(IPC) indicates the inter-process communication socket transport type. This transport type includes Unix domain sockets.
	TransportType VirtualSystemMigrationSettingData_TransportType
}

func NewCIM_VirtualSystemMigrationSettingDataEx1(instance *cim.WmiInstance) (newInstance *CIM_VirtualSystemMigrationSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_VirtualSystemMigrationSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewCIM_VirtualSystemMigrationSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_VirtualSystemMigrationSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_VirtualSystemMigrationSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetBandwidth sets the value of Bandwidth for the instance
func (instance *CIM_VirtualSystemMigrationSettingData) SetPropertyBandwidth(value uint16) (err error) {
	return instance.SetProperty("Bandwidth", (value))
}

// GetBandwidth gets the value of Bandwidth for the instance
func (instance *CIM_VirtualSystemMigrationSettingData) GetPropertyBandwidth() (value uint16, err error) {
	retValue, err := instance.GetProperty("Bandwidth")
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

// SetBandwidthUnit sets the value of BandwidthUnit for the instance
func (instance *CIM_VirtualSystemMigrationSettingData) SetPropertyBandwidthUnit(value string) (err error) {
	return instance.SetProperty("BandwidthUnit", (value))
}

// GetBandwidthUnit gets the value of BandwidthUnit for the instance
func (instance *CIM_VirtualSystemMigrationSettingData) GetPropertyBandwidthUnit() (value string, err error) {
	retValue, err := instance.GetProperty("BandwidthUnit")
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

// SetMigrationType sets the value of MigrationType for the instance
func (instance *CIM_VirtualSystemMigrationSettingData) SetPropertyMigrationType(value VirtualSystemMigrationSettingData_MigrationType) (err error) {
	return instance.SetProperty("MigrationType", (value))
}

// GetMigrationType gets the value of MigrationType for the instance
func (instance *CIM_VirtualSystemMigrationSettingData) GetPropertyMigrationType() (value VirtualSystemMigrationSettingData_MigrationType, err error) {
	retValue, err := instance.GetProperty("MigrationType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VirtualSystemMigrationSettingData_MigrationType(valuetmp)

	return
}

// SetOtherTransportType sets the value of OtherTransportType for the instance
func (instance *CIM_VirtualSystemMigrationSettingData) SetPropertyOtherTransportType(value string) (err error) {
	return instance.SetProperty("OtherTransportType", (value))
}

// GetOtherTransportType gets the value of OtherTransportType for the instance
func (instance *CIM_VirtualSystemMigrationSettingData) GetPropertyOtherTransportType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherTransportType")
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

// SetPriority sets the value of Priority for the instance
func (instance *CIM_VirtualSystemMigrationSettingData) SetPropertyPriority(value uint16) (err error) {
	return instance.SetProperty("Priority", (value))
}

// GetPriority gets the value of Priority for the instance
func (instance *CIM_VirtualSystemMigrationSettingData) GetPropertyPriority() (value uint16, err error) {
	retValue, err := instance.GetProperty("Priority")
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

// SetTransportType sets the value of TransportType for the instance
func (instance *CIM_VirtualSystemMigrationSettingData) SetPropertyTransportType(value VirtualSystemMigrationSettingData_TransportType) (err error) {
	return instance.SetProperty("TransportType", (value))
}

// GetTransportType gets the value of TransportType for the instance
func (instance *CIM_VirtualSystemMigrationSettingData) GetPropertyTransportType() (value VirtualSystemMigrationSettingData_TransportType, err error) {
	retValue, err := instance.GetProperty("TransportType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VirtualSystemMigrationSettingData_TransportType(valuetmp)

	return
}
