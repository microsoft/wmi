// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Cluster.Validation
//////////////////////////////////////////////
package validation

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFTCluster_ValidationDiskInfo struct
type MSFTCluster_ValidationDiskInfo struct {
	cim.WmiInstance

	// 2
	AdapterDescription string

	// 2
	BusType ValidationDiskInfo_BusType

	// 2
	DevicePath string

	// 2
	DiskId string

	// 2
	DiskIdType ValidationDiskInfo_DiskIdType

	// 1
	DiskNumber uint32

	// 2
	ExcludeFromTests bool

	// 2
	ExtendedFlags uint32

	// 2
	Flags uint32

	// 2
	FriendlyName string

	// 2
	GptPartitionType []ValidationDiskInfo_GptPartitionType

	// 2
	IsClusterable bool

	// 2
	IsClustered bool

	// 2
	IsHiddenDisk bool

	// 2
	IsPage89Supported bool

	// 2
	IsPoolDisk bool

	// 2
	MbrPartitionType []ValidationDiskInfo_MbrPartitionType

	// 83
	MediaType ValidationDiskInfo_MediaType

	// 2
	MiniportDriver string

	// 2
	Page83Id string

	// 2
	ScsiAddress string

	// 2
	SerialNumber string

	// 2
	ServiceName string

	// 2
	StackType ValidationDiskInfo_StackType
}

// SetAdapterDescription sets the value of AdapterDescription for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyAdapterDescription(value string) (err error) {
	return instance.SetProperty("AdapterDescription", value)
}

// GetAdapterDescription gets the value of AdapterDescription for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyAdapterDescription() (value string, err error) {
	retValue, err := instance.GetProperty("AdapterDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBusType sets the value of BusType for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyBusType(value ValidationDiskInfo_BusType) (err error) {
	return instance.SetProperty("BusType", value)
}

// GetBusType gets the value of BusType for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyBusType() (value ValidationDiskInfo_BusType, err error) {
	retValue, err := instance.GetProperty("BusType")
	if err != nil {
		return
	}
	value, ok := retValue.(ValidationDiskInfo_BusType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDevicePath sets the value of DevicePath for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyDevicePath(value string) (err error) {
	return instance.SetProperty("DevicePath", value)
}

// GetDevicePath gets the value of DevicePath for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyDevicePath() (value string, err error) {
	retValue, err := instance.GetProperty("DevicePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDiskId sets the value of DiskId for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyDiskId(value string) (err error) {
	return instance.SetProperty("DiskId", value)
}

// GetDiskId gets the value of DiskId for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyDiskId() (value string, err error) {
	retValue, err := instance.GetProperty("DiskId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDiskIdType sets the value of DiskIdType for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyDiskIdType(value ValidationDiskInfo_DiskIdType) (err error) {
	return instance.SetProperty("DiskIdType", value)
}

// GetDiskIdType gets the value of DiskIdType for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyDiskIdType() (value ValidationDiskInfo_DiskIdType, err error) {
	retValue, err := instance.GetProperty("DiskIdType")
	if err != nil {
		return
	}
	value, ok := retValue.(ValidationDiskInfo_DiskIdType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDiskNumber sets the value of DiskNumber for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyDiskNumber(value uint32) (err error) {
	return instance.SetProperty("DiskNumber", value)
}

// GetDiskNumber gets the value of DiskNumber for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyDiskNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("DiskNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExcludeFromTests sets the value of ExcludeFromTests for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyExcludeFromTests(value bool) (err error) {
	return instance.SetProperty("ExcludeFromTests", value)
}

// GetExcludeFromTests gets the value of ExcludeFromTests for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyExcludeFromTests() (value bool, err error) {
	retValue, err := instance.GetProperty("ExcludeFromTests")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExtendedFlags sets the value of ExtendedFlags for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyExtendedFlags(value uint32) (err error) {
	return instance.SetProperty("ExtendedFlags", value)
}

// GetExtendedFlags gets the value of ExtendedFlags for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyExtendedFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExtendedFlags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", value)
}

// GetFlags gets the value of Flags for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", value)
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGptPartitionType sets the value of GptPartitionType for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyGptPartitionType(value []ValidationDiskInfo_GptPartitionType) (err error) {
	return instance.SetProperty("GptPartitionType", value)
}

// GetGptPartitionType gets the value of GptPartitionType for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyGptPartitionType() (value []ValidationDiskInfo_GptPartitionType, err error) {
	retValue, err := instance.GetProperty("GptPartitionType")
	if err != nil {
		return
	}
	value, ok := retValue.([]ValidationDiskInfo_GptPartitionType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsClusterable sets the value of IsClusterable for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyIsClusterable(value bool) (err error) {
	return instance.SetProperty("IsClusterable", value)
}

// GetIsClusterable gets the value of IsClusterable for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyIsClusterable() (value bool, err error) {
	retValue, err := instance.GetProperty("IsClusterable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsClustered sets the value of IsClustered for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyIsClustered(value bool) (err error) {
	return instance.SetProperty("IsClustered", value)
}

// GetIsClustered gets the value of IsClustered for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyIsClustered() (value bool, err error) {
	retValue, err := instance.GetProperty("IsClustered")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsHiddenDisk sets the value of IsHiddenDisk for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyIsHiddenDisk(value bool) (err error) {
	return instance.SetProperty("IsHiddenDisk", value)
}

// GetIsHiddenDisk gets the value of IsHiddenDisk for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyIsHiddenDisk() (value bool, err error) {
	retValue, err := instance.GetProperty("IsHiddenDisk")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsPage89Supported sets the value of IsPage89Supported for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyIsPage89Supported(value bool) (err error) {
	return instance.SetProperty("IsPage89Supported", value)
}

// GetIsPage89Supported gets the value of IsPage89Supported for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyIsPage89Supported() (value bool, err error) {
	retValue, err := instance.GetProperty("IsPage89Supported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsPoolDisk sets the value of IsPoolDisk for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyIsPoolDisk(value bool) (err error) {
	return instance.SetProperty("IsPoolDisk", value)
}

// GetIsPoolDisk gets the value of IsPoolDisk for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyIsPoolDisk() (value bool, err error) {
	retValue, err := instance.GetProperty("IsPoolDisk")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMbrPartitionType sets the value of MbrPartitionType for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyMbrPartitionType(value []ValidationDiskInfo_MbrPartitionType) (err error) {
	return instance.SetProperty("MbrPartitionType", value)
}

// GetMbrPartitionType gets the value of MbrPartitionType for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyMbrPartitionType() (value []ValidationDiskInfo_MbrPartitionType, err error) {
	retValue, err := instance.GetProperty("MbrPartitionType")
	if err != nil {
		return
	}
	value, ok := retValue.([]ValidationDiskInfo_MbrPartitionType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMediaType sets the value of MediaType for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyMediaType(value ValidationDiskInfo_MediaType) (err error) {
	return instance.SetProperty("MediaType", value)
}

// GetMediaType gets the value of MediaType for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyMediaType() (value ValidationDiskInfo_MediaType, err error) {
	retValue, err := instance.GetProperty("MediaType")
	if err != nil {
		return
	}
	value, ok := retValue.(ValidationDiskInfo_MediaType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMiniportDriver sets the value of MiniportDriver for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyMiniportDriver(value string) (err error) {
	return instance.SetProperty("MiniportDriver", value)
}

// GetMiniportDriver gets the value of MiniportDriver for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyMiniportDriver() (value string, err error) {
	retValue, err := instance.GetProperty("MiniportDriver")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPage83Id sets the value of Page83Id for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyPage83Id(value string) (err error) {
	return instance.SetProperty("Page83Id", value)
}

// GetPage83Id gets the value of Page83Id for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyPage83Id() (value string, err error) {
	retValue, err := instance.GetProperty("Page83Id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScsiAddress sets the value of ScsiAddress for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyScsiAddress(value string) (err error) {
	return instance.SetProperty("ScsiAddress", value)
}

// GetScsiAddress gets the value of ScsiAddress for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyScsiAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ScsiAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertySerialNumber(value string) (err error) {
	return instance.SetProperty("SerialNumber", value)
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertySerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("SerialNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServiceName sets the value of ServiceName for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyServiceName(value string) (err error) {
	return instance.SetProperty("ServiceName", value)
}

// GetServiceName gets the value of ServiceName for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyServiceName() (value string, err error) {
	retValue, err := instance.GetProperty("ServiceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStackType sets the value of StackType for the instance
func (instance *MSFTCluster_ValidationDiskInfo) SetPropertyStackType(value ValidationDiskInfo_StackType) (err error) {
	return instance.SetProperty("StackType", value)
}

// GetStackType gets the value of StackType for the instance
func (instance *MSFTCluster_ValidationDiskInfo) GetPropertyStackType() (value ValidationDiskInfo_StackType, err error) {
	retValue, err := instance.GetProperty("StackType")
	if err != nil {
		return
	}
	value, ok := retValue.(ValidationDiskInfo_StackType)
	if !ok {
		// TODO: Set an error
	}
	return
}
