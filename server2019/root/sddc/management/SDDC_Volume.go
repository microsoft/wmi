// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	Resiliency uint16

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
	return instance.SetProperty("AdditionalStatusBitMap", value)
}

// GetAdditionalStatusBitMap gets the value of AdditionalStatusBitMap for the instance
func (instance *SDDC_Volume) GetPropertyAdditionalStatusBitMap() (value uint64, err error) {
	retValue, err := instance.GetProperty("AdditionalStatusBitMap")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAlerts sets the value of Alerts for the instance
func (instance *SDDC_Volume) SetPropertyAlerts(value []SDDC_Alert) (err error) {
	return instance.SetProperty("Alerts", value)
}

// GetAlerts gets the value of Alerts for the instance
func (instance *SDDC_Volume) GetPropertyAlerts() (value []SDDC_Alert, err error) {
	retValue, err := instance.GetProperty("Alerts")
	if err != nil {
		return
	}
	value, ok := retValue.([]SDDC_Alert)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageLatency sets the value of AverageLatency for the instance
func (instance *SDDC_Volume) SetPropertyAverageLatency(value float64) (err error) {
	return instance.SetProperty("AverageLatency", value)
}

// GetAverageLatency gets the value of AverageLatency for the instance
func (instance *SDDC_Volume) GetPropertyAverageLatency() (value float64, err error) {
	retValue, err := instance.GetProperty("AverageLatency")
	if err != nil {
		return
	}
	value, ok := retValue.(float64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDedupSavings sets the value of DedupSavings for the instance
func (instance *SDDC_Volume) SetPropertyDedupSavings(value uint64) (err error) {
	return instance.SetProperty("DedupSavings", value)
}

// GetDedupSavings gets the value of DedupSavings for the instance
func (instance *SDDC_Volume) GetPropertyDedupSavings() (value uint64, err error) {
	retValue, err := instance.GetProperty("DedupSavings")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDedupSavingsRate sets the value of DedupSavingsRate for the instance
func (instance *SDDC_Volume) SetPropertyDedupSavingsRate(value uint32) (err error) {
	return instance.SetProperty("DedupSavingsRate", value)
}

// GetDedupSavingsRate gets the value of DedupSavingsRate for the instance
func (instance *SDDC_Volume) GetPropertyDedupSavingsRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("DedupSavingsRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncryptionPercentage sets the value of EncryptionPercentage for the instance
func (instance *SDDC_Volume) SetPropertyEncryptionPercentage(value uint16) (err error) {
	return instance.SetProperty("EncryptionPercentage", value)
}

// GetEncryptionPercentage gets the value of EncryptionPercentage for the instance
func (instance *SDDC_Volume) GetPropertyEncryptionPercentage() (value uint16, err error) {
	retValue, err := instance.GetProperty("EncryptionPercentage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncryptionStatus sets the value of EncryptionStatus for the instance
func (instance *SDDC_Volume) SetPropertyEncryptionStatus(value uint16) (err error) {
	return instance.SetProperty("EncryptionStatus", value)
}

// GetEncryptionStatus gets the value of EncryptionStatus for the instance
func (instance *SDDC_Volume) GetPropertyEncryptionStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("EncryptionStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultDomainAwareness sets the value of FaultDomainAwareness for the instance
func (instance *SDDC_Volume) SetPropertyFaultDomainAwareness(value uint16) (err error) {
	return instance.SetProperty("FaultDomainAwareness", value)
}

// GetFaultDomainAwareness gets the value of FaultDomainAwareness for the instance
func (instance *SDDC_Volume) GetPropertyFaultDomainAwareness() (value uint16, err error) {
	retValue, err := instance.GetProperty("FaultDomainAwareness")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileSystem sets the value of FileSystem for the instance
func (instance *SDDC_Volume) SetPropertyFileSystem(value string) (err error) {
	return instance.SetProperty("FileSystem", value)
}

// GetFileSystem gets the value of FileSystem for the instance
func (instance *SDDC_Volume) GetPropertyFileSystem() (value string, err error) {
	retValue, err := instance.GetProperty("FileSystem")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFootprint sets the value of Footprint for the instance
func (instance *SDDC_Volume) SetPropertyFootprint(value uint64) (err error) {
	return instance.SetProperty("Footprint", value)
}

// GetFootprint gets the value of Footprint for the instance
func (instance *SDDC_Volume) GetPropertyFootprint() (value uint64, err error) {
	retValue, err := instance.GetProperty("Footprint")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *SDDC_Volume) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *SDDC_Volume) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsDedupEnabled sets the value of IsDedupEnabled for the instance
func (instance *SDDC_Volume) SetPropertyIsDedupEnabled(value bool) (err error) {
	return instance.SetProperty("IsDedupEnabled", value)
}

// GetIsDedupEnabled gets the value of IsDedupEnabled for the instance
func (instance *SDDC_Volume) GetPropertyIsDedupEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDedupEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsEncrypted sets the value of IsEncrypted for the instance
func (instance *SDDC_Volume) SetPropertyIsEncrypted(value bool) (err error) {
	return instance.SetProperty("IsEncrypted", value)
}

// GetIsEncrypted gets the value of IsEncrypted for the instance
func (instance *SDDC_Volume) GetPropertyIsEncrypted() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEncrypted")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsIntegrityEnabled sets the value of IsIntegrityEnabled for the instance
func (instance *SDDC_Volume) SetPropertyIsIntegrityEnabled(value bool) (err error) {
	return instance.SetProperty("IsIntegrityEnabled", value)
}

// GetIsIntegrityEnabled gets the value of IsIntegrityEnabled for the instance
func (instance *SDDC_Volume) GetPropertyIsIntegrityEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsIntegrityEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsTiered sets the value of IsTiered for the instance
func (instance *SDDC_Volume) SetPropertyIsTiered(value bool) (err error) {
	return instance.SetProperty("IsTiered", value)
}

// GetIsTiered gets the value of IsTiered for the instance
func (instance *SDDC_Volume) GetPropertyIsTiered() (value bool, err error) {
	retValue, err := instance.GetProperty("IsTiered")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobs sets the value of Jobs for the instance
func (instance *SDDC_Volume) SetPropertyJobs(value []SDDC_Job) (err error) {
	return instance.SetProperty("Jobs", value)
}

// GetJobs gets the value of Jobs for the instance
func (instance *SDDC_Volume) GetPropertyJobs() (value []SDDC_Job, err error) {
	retValue, err := instance.GetProperty("Jobs")
	if err != nil {
		return
	}
	value, ok := retValue.([]SDDC_Job)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMedia sets the value of Media for the instance
func (instance *SDDC_Volume) SetPropertyMedia(value uint16) (err error) {
	return instance.SetProperty("Media", value)
}

// GetMedia gets the value of Media for the instance
func (instance *SDDC_Volume) GetPropertyMedia() (value uint16, err error) {
	retValue, err := instance.GetProperty("Media")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *SDDC_Volume) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *SDDC_Volume) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPath sets the value of Path for the instance
func (instance *SDDC_Volume) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", value)
}

// GetPath gets the value of Path for the instance
func (instance *SDDC_Volume) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResiliency sets the value of Resiliency for the instance
func (instance *SDDC_Volume) SetPropertyResiliency(value uint16) (err error) {
	return instance.SetProperty("Resiliency", value)
}

// GetResiliency gets the value of Resiliency for the instance
func (instance *SDDC_Volume) GetPropertyResiliency() (value uint16, err error) {
	retValue, err := instance.GetProperty("Resiliency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSiteName sets the value of SiteName for the instance
func (instance *SDDC_Volume) SetPropertySiteName(value string) (err error) {
	return instance.SetProperty("SiteName", value)
}

// GetSiteName gets the value of SiteName for the instance
func (instance *SDDC_Volume) GetPropertySiteName() (value string, err error) {
	retValue, err := instance.GetProperty("SiteName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSize sets the value of Size for the instance
func (instance *SDDC_Volume) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", value)
}

// GetSize gets the value of Size for the instance
func (instance *SDDC_Volume) GetPropertySize() (value uint64, err error) {
	retValue, err := instance.GetProperty("Size")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSizeRemaining sets the value of SizeRemaining for the instance
func (instance *SDDC_Volume) SetPropertySizeRemaining(value uint64) (err error) {
	return instance.SetProperty("SizeRemaining", value)
}

// GetSizeRemaining gets the value of SizeRemaining for the instance
func (instance *SDDC_Volume) GetPropertySizeRemaining() (value uint64, err error) {
	retValue, err := instance.GetProperty("SizeRemaining")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *SDDC_Volume) SetPropertyStatus(value []uint16) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *SDDC_Volume) GetPropertyStatus() (value []uint16, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatusCategory sets the value of StatusCategory for the instance
func (instance *SDDC_Volume) SetPropertyStatusCategory(value uint16) (err error) {
	return instance.SetProperty("StatusCategory", value)
}

// GetStatusCategory gets the value of StatusCategory for the instance
func (instance *SDDC_Volume) GetPropertyStatusCategory() (value uint16, err error) {
	retValue, err := instance.GetProperty("StatusCategory")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStoragePoolName sets the value of StoragePoolName for the instance
func (instance *SDDC_Volume) SetPropertyStoragePoolName(value string) (err error) {
	return instance.SetProperty("StoragePoolName", value)
}

// GetStoragePoolName gets the value of StoragePoolName for the instance
func (instance *SDDC_Volume) GetPropertyStoragePoolName() (value string, err error) {
	retValue, err := instance.GetProperty("StoragePoolName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierFootprints sets the value of TierFootprints for the instance
func (instance *SDDC_Volume) SetPropertyTierFootprints(value []uint64) (err error) {
	return instance.SetProperty("TierFootprints", value)
}

// GetTierFootprints gets the value of TierFootprints for the instance
func (instance *SDDC_Volume) GetPropertyTierFootprints() (value []uint64, err error) {
	retValue, err := instance.GetProperty("TierFootprints")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierMedias sets the value of TierMedias for the instance
func (instance *SDDC_Volume) SetPropertyTierMedias(value []uint16) (err error) {
	return instance.SetProperty("TierMedias", value)
}

// GetTierMedias gets the value of TierMedias for the instance
func (instance *SDDC_Volume) GetPropertyTierMedias() (value []uint16, err error) {
	retValue, err := instance.GetProperty("TierMedias")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierResiliencies sets the value of TierResiliencies for the instance
func (instance *SDDC_Volume) SetPropertyTierResiliencies(value []uint16) (err error) {
	return instance.SetProperty("TierResiliencies", value)
}

// GetTierResiliencies gets the value of TierResiliencies for the instance
func (instance *SDDC_Volume) GetPropertyTierResiliencies() (value []uint16, err error) {
	retValue, err := instance.GetProperty("TierResiliencies")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierSizes sets the value of TierSizes for the instance
func (instance *SDDC_Volume) SetPropertyTierSizes(value []uint64) (err error) {
	return instance.SetProperty("TierSizes", value)
}

// GetTierSizes gets the value of TierSizes for the instance
func (instance *SDDC_Volume) GetPropertyTierSizes() (value []uint64, err error) {
	retValue, err := instance.GetProperty("TierSizes")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalIops sets the value of TotalIops for the instance
func (instance *SDDC_Volume) SetPropertyTotalIops(value float64) (err error) {
	return instance.SetProperty("TotalIops", value)
}

// GetTotalIops gets the value of TotalIops for the instance
func (instance *SDDC_Volume) GetPropertyTotalIops() (value float64, err error) {
	retValue, err := instance.GetProperty("TotalIops")
	if err != nil {
		return
	}
	value, ok := retValue.(float64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalThroughput sets the value of TotalThroughput for the instance
func (instance *SDDC_Volume) SetPropertyTotalThroughput(value float64) (err error) {
	return instance.SetProperty("TotalThroughput", value)
}

// GetTotalThroughput gets the value of TotalThroughput for the instance
func (instance *SDDC_Volume) GetPropertyTotalThroughput() (value float64, err error) {
	retValue, err := instance.GetProperty("TotalThroughput")
	if err != nil {
		return
	}
	value, ok := retValue.(float64)
	if !ok {
		// TODO: Set an error
	}
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
