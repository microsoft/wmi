// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SDDC_Volume struct
type SDDC_Volume struct {
	*cim.WmiInstance

	//
	AdditionalStatusBitMap uint64

	//
	Alerts []SDDC_Alert

	//
	AverageLatency float64

	//
	DedupSavings uint64

	//
	DedupSavingsRate uint32

	//
	EncryptionPercentage uint16

	//
	EncryptionStatus uint16

	//
	FaultDomainAwareness uint16

	//
	FileSystem string

	//
	Footprint uint64

	//
	Id string

	//
	IsDedupEnabled bool

	//
	IsEncrypted bool

	//
	IsIntegrityEnabled bool

	//
	IsTiered bool

	//
	Jobs []SDDC_Job

	//
	Media uint16

	//
	Name string

	//
	Path string

	//
	ProvisioningType uint16

	//
	ReFSDedupCompressionSavingsSize uint64

	//
	ReFSDedupLastRunStatus uint64

	//
	ReFSDedupLastRunTime string

	//
	ReFSDedupMode uint32

	//
	ReFSDedupNextRunTime string

	//
	ReFSDedupSavingsSize uint64

	//
	ReplicatedDiskType uint16

	//
	ReplicationGroupName string

	//
	ReplicationMode uint32

	//
	ReplicationStatus uint32

	//
	Resiliency uint16

	//
	Server string

	//
	SiteName string

	//
	Size uint64

	//
	SizeRemaining uint64

	//
	Status []uint16

	//
	StatusCategory uint16

	//
	StoragePoolName string

	//
	TierFootprints []uint64

	//
	TierMedias []uint16

	//
	TierResiliencies []uint16

	//
	TierSizes []uint64

	//
	TotalIops float64

	//
	TotalThroughput float64
}

func NewSDDC_VolumeEx1(instance *cim.WmiInstance) (newInstance *SDDC_Volume, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &SDDC_Volume{
		WmiInstance: tmp,
	}
	return
}

func NewSDDC_VolumeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SDDC_Volume, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SDDC_Volume{
		WmiInstance: tmp,
	}
	return
}

// SetAdditionalStatusBitMap sets the value of AdditionalStatusBitMap for the instance
func (instance *SDDC_Volume) SetPropertyAdditionalStatusBitMap(value uint64) (err error) {
	return instance.SetProperty("AdditionalStatusBitMap", (value))
}

// GetAdditionalStatusBitMap gets the value of AdditionalStatusBitMap for the instance
func (instance *SDDC_Volume) GetPropertyAdditionalStatusBitMap() (value uint64, err error) {
	retValue, err := instance.GetProperty("AdditionalStatusBitMap")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAlerts sets the value of Alerts for the instance
func (instance *SDDC_Volume) SetPropertyAlerts(value []SDDC_Alert) (err error) {
	return instance.SetProperty("Alerts", (value))
}

// GetAlerts gets the value of Alerts for the instance
func (instance *SDDC_Volume) GetPropertyAlerts() (value []SDDC_Alert, err error) {
	retValue, err := instance.GetProperty("Alerts")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(SDDC_Alert)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " SDDC_Alert is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, SDDC_Alert(valuetmp))
	}

	return
}

// SetAverageLatency sets the value of AverageLatency for the instance
func (instance *SDDC_Volume) SetPropertyAverageLatency(value float64) (err error) {
	return instance.SetProperty("AverageLatency", (value))
}

// GetAverageLatency gets the value of AverageLatency for the instance
func (instance *SDDC_Volume) GetPropertyAverageLatency() (value float64, err error) {
	retValue, err := instance.GetProperty("AverageLatency")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float64(valuetmp)

	return
}

// SetDedupSavings sets the value of DedupSavings for the instance
func (instance *SDDC_Volume) SetPropertyDedupSavings(value uint64) (err error) {
	return instance.SetProperty("DedupSavings", (value))
}

// GetDedupSavings gets the value of DedupSavings for the instance
func (instance *SDDC_Volume) GetPropertyDedupSavings() (value uint64, err error) {
	retValue, err := instance.GetProperty("DedupSavings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetDedupSavingsRate sets the value of DedupSavingsRate for the instance
func (instance *SDDC_Volume) SetPropertyDedupSavingsRate(value uint32) (err error) {
	return instance.SetProperty("DedupSavingsRate", (value))
}

// GetDedupSavingsRate gets the value of DedupSavingsRate for the instance
func (instance *SDDC_Volume) GetPropertyDedupSavingsRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("DedupSavingsRate")
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

// SetEncryptionPercentage sets the value of EncryptionPercentage for the instance
func (instance *SDDC_Volume) SetPropertyEncryptionPercentage(value uint16) (err error) {
	return instance.SetProperty("EncryptionPercentage", (value))
}

// GetEncryptionPercentage gets the value of EncryptionPercentage for the instance
func (instance *SDDC_Volume) GetPropertyEncryptionPercentage() (value uint16, err error) {
	retValue, err := instance.GetProperty("EncryptionPercentage")
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

// SetEncryptionStatus sets the value of EncryptionStatus for the instance
func (instance *SDDC_Volume) SetPropertyEncryptionStatus(value uint16) (err error) {
	return instance.SetProperty("EncryptionStatus", (value))
}

// GetEncryptionStatus gets the value of EncryptionStatus for the instance
func (instance *SDDC_Volume) GetPropertyEncryptionStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("EncryptionStatus")
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

// SetFaultDomainAwareness sets the value of FaultDomainAwareness for the instance
func (instance *SDDC_Volume) SetPropertyFaultDomainAwareness(value uint16) (err error) {
	return instance.SetProperty("FaultDomainAwareness", (value))
}

// GetFaultDomainAwareness gets the value of FaultDomainAwareness for the instance
func (instance *SDDC_Volume) GetPropertyFaultDomainAwareness() (value uint16, err error) {
	retValue, err := instance.GetProperty("FaultDomainAwareness")
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

// SetFileSystem sets the value of FileSystem for the instance
func (instance *SDDC_Volume) SetPropertyFileSystem(value string) (err error) {
	return instance.SetProperty("FileSystem", (value))
}

// GetFileSystem gets the value of FileSystem for the instance
func (instance *SDDC_Volume) GetPropertyFileSystem() (value string, err error) {
	retValue, err := instance.GetProperty("FileSystem")
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

// SetFootprint sets the value of Footprint for the instance
func (instance *SDDC_Volume) SetPropertyFootprint(value uint64) (err error) {
	return instance.SetProperty("Footprint", (value))
}

// GetFootprint gets the value of Footprint for the instance
func (instance *SDDC_Volume) GetPropertyFootprint() (value uint64, err error) {
	retValue, err := instance.GetProperty("Footprint")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetId sets the value of Id for the instance
func (instance *SDDC_Volume) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *SDDC_Volume) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetIsDedupEnabled sets the value of IsDedupEnabled for the instance
func (instance *SDDC_Volume) SetPropertyIsDedupEnabled(value bool) (err error) {
	return instance.SetProperty("IsDedupEnabled", (value))
}

// GetIsDedupEnabled gets the value of IsDedupEnabled for the instance
func (instance *SDDC_Volume) GetPropertyIsDedupEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDedupEnabled")
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

// SetIsEncrypted sets the value of IsEncrypted for the instance
func (instance *SDDC_Volume) SetPropertyIsEncrypted(value bool) (err error) {
	return instance.SetProperty("IsEncrypted", (value))
}

// GetIsEncrypted gets the value of IsEncrypted for the instance
func (instance *SDDC_Volume) GetPropertyIsEncrypted() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEncrypted")
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

// SetIsIntegrityEnabled sets the value of IsIntegrityEnabled for the instance
func (instance *SDDC_Volume) SetPropertyIsIntegrityEnabled(value bool) (err error) {
	return instance.SetProperty("IsIntegrityEnabled", (value))
}

// GetIsIntegrityEnabled gets the value of IsIntegrityEnabled for the instance
func (instance *SDDC_Volume) GetPropertyIsIntegrityEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsIntegrityEnabled")
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

// SetIsTiered sets the value of IsTiered for the instance
func (instance *SDDC_Volume) SetPropertyIsTiered(value bool) (err error) {
	return instance.SetProperty("IsTiered", (value))
}

// GetIsTiered gets the value of IsTiered for the instance
func (instance *SDDC_Volume) GetPropertyIsTiered() (value bool, err error) {
	retValue, err := instance.GetProperty("IsTiered")
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

// SetJobs sets the value of Jobs for the instance
func (instance *SDDC_Volume) SetPropertyJobs(value []SDDC_Job) (err error) {
	return instance.SetProperty("Jobs", (value))
}

// GetJobs gets the value of Jobs for the instance
func (instance *SDDC_Volume) GetPropertyJobs() (value []SDDC_Job, err error) {
	retValue, err := instance.GetProperty("Jobs")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(SDDC_Job)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " SDDC_Job is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, SDDC_Job(valuetmp))
	}

	return
}

// SetMedia sets the value of Media for the instance
func (instance *SDDC_Volume) SetPropertyMedia(value uint16) (err error) {
	return instance.SetProperty("Media", (value))
}

// GetMedia gets the value of Media for the instance
func (instance *SDDC_Volume) GetPropertyMedia() (value uint16, err error) {
	retValue, err := instance.GetProperty("Media")
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

// SetName sets the value of Name for the instance
func (instance *SDDC_Volume) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *SDDC_Volume) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetPath sets the value of Path for the instance
func (instance *SDDC_Volume) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *SDDC_Volume) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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

// SetProvisioningType sets the value of ProvisioningType for the instance
func (instance *SDDC_Volume) SetPropertyProvisioningType(value uint16) (err error) {
	return instance.SetProperty("ProvisioningType", (value))
}

// GetProvisioningType gets the value of ProvisioningType for the instance
func (instance *SDDC_Volume) GetPropertyProvisioningType() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProvisioningType")
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

// SetReFSDedupCompressionSavingsSize sets the value of ReFSDedupCompressionSavingsSize for the instance
func (instance *SDDC_Volume) SetPropertyReFSDedupCompressionSavingsSize(value uint64) (err error) {
	return instance.SetProperty("ReFSDedupCompressionSavingsSize", (value))
}

// GetReFSDedupCompressionSavingsSize gets the value of ReFSDedupCompressionSavingsSize for the instance
func (instance *SDDC_Volume) GetPropertyReFSDedupCompressionSavingsSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReFSDedupCompressionSavingsSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetReFSDedupLastRunStatus sets the value of ReFSDedupLastRunStatus for the instance
func (instance *SDDC_Volume) SetPropertyReFSDedupLastRunStatus(value uint64) (err error) {
	return instance.SetProperty("ReFSDedupLastRunStatus", (value))
}

// GetReFSDedupLastRunStatus gets the value of ReFSDedupLastRunStatus for the instance
func (instance *SDDC_Volume) GetPropertyReFSDedupLastRunStatus() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReFSDedupLastRunStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetReFSDedupLastRunTime sets the value of ReFSDedupLastRunTime for the instance
func (instance *SDDC_Volume) SetPropertyReFSDedupLastRunTime(value string) (err error) {
	return instance.SetProperty("ReFSDedupLastRunTime", (value))
}

// GetReFSDedupLastRunTime gets the value of ReFSDedupLastRunTime for the instance
func (instance *SDDC_Volume) GetPropertyReFSDedupLastRunTime() (value string, err error) {
	retValue, err := instance.GetProperty("ReFSDedupLastRunTime")
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

// SetReFSDedupMode sets the value of ReFSDedupMode for the instance
func (instance *SDDC_Volume) SetPropertyReFSDedupMode(value uint32) (err error) {
	return instance.SetProperty("ReFSDedupMode", (value))
}

// GetReFSDedupMode gets the value of ReFSDedupMode for the instance
func (instance *SDDC_Volume) GetPropertyReFSDedupMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReFSDedupMode")
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

// SetReFSDedupNextRunTime sets the value of ReFSDedupNextRunTime for the instance
func (instance *SDDC_Volume) SetPropertyReFSDedupNextRunTime(value string) (err error) {
	return instance.SetProperty("ReFSDedupNextRunTime", (value))
}

// GetReFSDedupNextRunTime gets the value of ReFSDedupNextRunTime for the instance
func (instance *SDDC_Volume) GetPropertyReFSDedupNextRunTime() (value string, err error) {
	retValue, err := instance.GetProperty("ReFSDedupNextRunTime")
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

// SetReFSDedupSavingsSize sets the value of ReFSDedupSavingsSize for the instance
func (instance *SDDC_Volume) SetPropertyReFSDedupSavingsSize(value uint64) (err error) {
	return instance.SetProperty("ReFSDedupSavingsSize", (value))
}

// GetReFSDedupSavingsSize gets the value of ReFSDedupSavingsSize for the instance
func (instance *SDDC_Volume) GetPropertyReFSDedupSavingsSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReFSDedupSavingsSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetReplicatedDiskType sets the value of ReplicatedDiskType for the instance
func (instance *SDDC_Volume) SetPropertyReplicatedDiskType(value uint16) (err error) {
	return instance.SetProperty("ReplicatedDiskType", (value))
}

// GetReplicatedDiskType gets the value of ReplicatedDiskType for the instance
func (instance *SDDC_Volume) GetPropertyReplicatedDiskType() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicatedDiskType")
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

// SetReplicationGroupName sets the value of ReplicationGroupName for the instance
func (instance *SDDC_Volume) SetPropertyReplicationGroupName(value string) (err error) {
	return instance.SetProperty("ReplicationGroupName", (value))
}

// GetReplicationGroupName gets the value of ReplicationGroupName for the instance
func (instance *SDDC_Volume) GetPropertyReplicationGroupName() (value string, err error) {
	retValue, err := instance.GetProperty("ReplicationGroupName")
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

// SetReplicationMode sets the value of ReplicationMode for the instance
func (instance *SDDC_Volume) SetPropertyReplicationMode(value uint32) (err error) {
	return instance.SetProperty("ReplicationMode", (value))
}

// GetReplicationMode gets the value of ReplicationMode for the instance
func (instance *SDDC_Volume) GetPropertyReplicationMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReplicationMode")
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

// SetReplicationStatus sets the value of ReplicationStatus for the instance
func (instance *SDDC_Volume) SetPropertyReplicationStatus(value uint32) (err error) {
	return instance.SetProperty("ReplicationStatus", (value))
}

// GetReplicationStatus gets the value of ReplicationStatus for the instance
func (instance *SDDC_Volume) GetPropertyReplicationStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReplicationStatus")
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

// SetResiliency sets the value of Resiliency for the instance
func (instance *SDDC_Volume) SetPropertyResiliency(value uint16) (err error) {
	return instance.SetProperty("Resiliency", (value))
}

// GetResiliency gets the value of Resiliency for the instance
func (instance *SDDC_Volume) GetPropertyResiliency() (value uint16, err error) {
	retValue, err := instance.GetProperty("Resiliency")
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

// SetServer sets the value of Server for the instance
func (instance *SDDC_Volume) SetPropertyServer(value string) (err error) {
	return instance.SetProperty("Server", (value))
}

// GetServer gets the value of Server for the instance
func (instance *SDDC_Volume) GetPropertyServer() (value string, err error) {
	retValue, err := instance.GetProperty("Server")
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

// SetSiteName sets the value of SiteName for the instance
func (instance *SDDC_Volume) SetPropertySiteName(value string) (err error) {
	return instance.SetProperty("SiteName", (value))
}

// GetSiteName gets the value of SiteName for the instance
func (instance *SDDC_Volume) GetPropertySiteName() (value string, err error) {
	retValue, err := instance.GetProperty("SiteName")
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

// SetSize sets the value of Size for the instance
func (instance *SDDC_Volume) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *SDDC_Volume) GetPropertySize() (value uint64, err error) {
	retValue, err := instance.GetProperty("Size")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSizeRemaining sets the value of SizeRemaining for the instance
func (instance *SDDC_Volume) SetPropertySizeRemaining(value uint64) (err error) {
	return instance.SetProperty("SizeRemaining", (value))
}

// GetSizeRemaining gets the value of SizeRemaining for the instance
func (instance *SDDC_Volume) GetPropertySizeRemaining() (value uint64, err error) {
	retValue, err := instance.GetProperty("SizeRemaining")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetStatus sets the value of Status for the instance
func (instance *SDDC_Volume) SetPropertyStatus(value []uint16) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *SDDC_Volume) GetPropertyStatus() (value []uint16, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetStatusCategory sets the value of StatusCategory for the instance
func (instance *SDDC_Volume) SetPropertyStatusCategory(value uint16) (err error) {
	return instance.SetProperty("StatusCategory", (value))
}

// GetStatusCategory gets the value of StatusCategory for the instance
func (instance *SDDC_Volume) GetPropertyStatusCategory() (value uint16, err error) {
	retValue, err := instance.GetProperty("StatusCategory")
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

// SetStoragePoolName sets the value of StoragePoolName for the instance
func (instance *SDDC_Volume) SetPropertyStoragePoolName(value string) (err error) {
	return instance.SetProperty("StoragePoolName", (value))
}

// GetStoragePoolName gets the value of StoragePoolName for the instance
func (instance *SDDC_Volume) GetPropertyStoragePoolName() (value string, err error) {
	retValue, err := instance.GetProperty("StoragePoolName")
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

// SetTierFootprints sets the value of TierFootprints for the instance
func (instance *SDDC_Volume) SetPropertyTierFootprints(value []uint64) (err error) {
	return instance.SetProperty("TierFootprints", (value))
}

// GetTierFootprints gets the value of TierFootprints for the instance
func (instance *SDDC_Volume) GetPropertyTierFootprints() (value []uint64, err error) {
	retValue, err := instance.GetProperty("TierFootprints")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint64)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint64(valuetmp))
	}

	return
}

// SetTierMedias sets the value of TierMedias for the instance
func (instance *SDDC_Volume) SetPropertyTierMedias(value []uint16) (err error) {
	return instance.SetProperty("TierMedias", (value))
}

// GetTierMedias gets the value of TierMedias for the instance
func (instance *SDDC_Volume) GetPropertyTierMedias() (value []uint16, err error) {
	retValue, err := instance.GetProperty("TierMedias")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetTierResiliencies sets the value of TierResiliencies for the instance
func (instance *SDDC_Volume) SetPropertyTierResiliencies(value []uint16) (err error) {
	return instance.SetProperty("TierResiliencies", (value))
}

// GetTierResiliencies gets the value of TierResiliencies for the instance
func (instance *SDDC_Volume) GetPropertyTierResiliencies() (value []uint16, err error) {
	retValue, err := instance.GetProperty("TierResiliencies")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetTierSizes sets the value of TierSizes for the instance
func (instance *SDDC_Volume) SetPropertyTierSizes(value []uint64) (err error) {
	return instance.SetProperty("TierSizes", (value))
}

// GetTierSizes gets the value of TierSizes for the instance
func (instance *SDDC_Volume) GetPropertyTierSizes() (value []uint64, err error) {
	retValue, err := instance.GetProperty("TierSizes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint64)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint64(valuetmp))
	}

	return
}

// SetTotalIops sets the value of TotalIops for the instance
func (instance *SDDC_Volume) SetPropertyTotalIops(value float64) (err error) {
	return instance.SetProperty("TotalIops", (value))
}

// GetTotalIops gets the value of TotalIops for the instance
func (instance *SDDC_Volume) GetPropertyTotalIops() (value float64, err error) {
	retValue, err := instance.GetProperty("TotalIops")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float64(valuetmp)

	return
}

// SetTotalThroughput sets the value of TotalThroughput for the instance
func (instance *SDDC_Volume) SetPropertyTotalThroughput(value float64) (err error) {
	return instance.SetProperty("TotalThroughput", (value))
}

// GetTotalThroughput gets the value of TotalThroughput for the instance
func (instance *SDDC_Volume) GetPropertyTotalThroughput() (value float64, err error) {
	retValue, err := instance.GetProperty("TotalThroughput")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float64(valuetmp)

	return
}

//

// <param name="NewVolumeTemplate" type="SDDC_VolumeModificationTemplate "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Volume) GetNewVolumeTemplate( /* OUT */ NewVolumeTemplate SDDC_VolumeModificationTemplate) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetNewVolumeTemplate")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="BackupReoveryPasswordToAD" type="bool "></param>
// <param name="DedupMode" type="uint32 "></param>
// <param name="EnableBitLocker" type="bool "></param>
// <param name="IsTiered" type="bool "></param>
// <param name="PasswordProtector" type="string "></param>
// <param name="Resiliency" type="uint16 "></param>
// <param name="SetFileIntegrity" type="bool "></param>
// <param name="Sizes" type="uint64 []"></param>
// <param name="VolumeName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Volume) NewVolume( /* IN */ VolumeName string,
	/* IN */ Resiliency uint16,
	/* IN */ IsTiered bool,
	/* IN */ Sizes []uint64,
	/* IN */ DedupMode uint32,
	/* IN */ SetFileIntegrity bool,
	/* IN */ EnableBitLocker bool,
	/* IN */ BackupReoveryPasswordToAD bool,
	/* IN */ PasswordProtector string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("NewVolume", VolumeName, Resiliency, IsTiered, Sizes, DedupMode, SetFileIntegrity, EnableBitLocker, BackupReoveryPasswordToAD, PasswordProtector)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ResizeVolumeTemplate" type="SDDC_VolumeModificationTemplate "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Volume) GetResizeVolumeTemplate( /* OUT */ ResizeVolumeTemplate SDDC_VolumeModificationTemplate) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetResizeVolumeTemplate")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IsTiered" type="bool "></param>
// <param name="NewSize" type="uint64 []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Volume) ResizeVolume( /* IN */ IsTiered bool,
	/* IN */ NewSize []uint64) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ResizeVolume", IsTiered, NewSize)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Volume) DeleteVolume() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DeleteVolume")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Volume) OnlineVolume() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("OnlineVolume")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Volume) OfflineVolume() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("OfflineVolume")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="BackupRecoveryPasswordToAD" type="bool "></param>
// <param name="PasswordProtector" type="string "></param>

// <param name="Result" type="SDDC_BitlockerResult "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Volume) EncryptVolume( /* IN */ PasswordProtector string,
	/* IN */ BackupRecoveryPasswordToAD bool,
	/* OUT */ Result SDDC_BitlockerResult) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EncryptVolume", PasswordProtector, BackupRecoveryPasswordToAD)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Result" type="SDDC_BitlockerResult "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Volume) GetEncryptedVolumeRecoveryPassword( /* OUT */ Result SDDC_BitlockerResult) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetEncryptedVolumeRecoveryPassword")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Volume) DecryptVolume() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DecryptVolume")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Mode" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Volume) SetDedupMode( /* IN */ Mode uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetDedupMode", Mode)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="SeriesName" type="string "></param>
// <param name="TimeFrame" type="uint16 "></param>

// <param name="Metric" type="SDDC_Metric "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Volume) GetMetrics( /* IN */ SeriesName string,
	/* IN */ TimeFrame uint16,
	/* OUT */ Metric SDDC_Metric) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetMetrics", SeriesName, TimeFrame)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Volume) Refresh() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Refresh")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
