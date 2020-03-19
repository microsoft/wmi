// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_StorageSpacesDirect struct
type MSCluster_StorageSpacesDirect struct {
	*cim.WmiInstance

	//
	CacheDeviceModel []string

	//
	CacheMetadataReserveBytes uint64

	//
	CacheModeHDD uint32

	//
	CacheModeSSD uint32

	//
	CachePageSizeKBytes uint32

	//
	CacheState uint32

	//
	EnableReportName string

	//
	Name string

	//
	Node string

	//
	ScmUse uint32

	//
	SedProtectionState uint32

	//
	State uint32

	//
	UseSedExclusively bool
}

func NewMSCluster_StorageSpacesDirectEx1(instance *cim.WmiInstance) (newInstance *MSCluster_StorageSpacesDirect, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_StorageSpacesDirect{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_StorageSpacesDirectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_StorageSpacesDirect, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_StorageSpacesDirect{
		WmiInstance: tmp,
	}
	return
}

// SetCacheDeviceModel sets the value of CacheDeviceModel for the instance
func (instance *MSCluster_StorageSpacesDirect) SetPropertyCacheDeviceModel(value []string) (err error) {
	return instance.SetProperty("CacheDeviceModel", value)
}

// GetCacheDeviceModel gets the value of CacheDeviceModel for the instance
func (instance *MSCluster_StorageSpacesDirect) GetPropertyCacheDeviceModel() (value []string, err error) {
	retValue, err := instance.GetProperty("CacheDeviceModel")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheMetadataReserveBytes sets the value of CacheMetadataReserveBytes for the instance
func (instance *MSCluster_StorageSpacesDirect) SetPropertyCacheMetadataReserveBytes(value uint64) (err error) {
	return instance.SetProperty("CacheMetadataReserveBytes", value)
}

// GetCacheMetadataReserveBytes gets the value of CacheMetadataReserveBytes for the instance
func (instance *MSCluster_StorageSpacesDirect) GetPropertyCacheMetadataReserveBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheMetadataReserveBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheModeHDD sets the value of CacheModeHDD for the instance
func (instance *MSCluster_StorageSpacesDirect) SetPropertyCacheModeHDD(value uint32) (err error) {
	return instance.SetProperty("CacheModeHDD", value)
}

// GetCacheModeHDD gets the value of CacheModeHDD for the instance
func (instance *MSCluster_StorageSpacesDirect) GetPropertyCacheModeHDD() (value uint32, err error) {
	retValue, err := instance.GetProperty("CacheModeHDD")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheModeSSD sets the value of CacheModeSSD for the instance
func (instance *MSCluster_StorageSpacesDirect) SetPropertyCacheModeSSD(value uint32) (err error) {
	return instance.SetProperty("CacheModeSSD", value)
}

// GetCacheModeSSD gets the value of CacheModeSSD for the instance
func (instance *MSCluster_StorageSpacesDirect) GetPropertyCacheModeSSD() (value uint32, err error) {
	retValue, err := instance.GetProperty("CacheModeSSD")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCachePageSizeKBytes sets the value of CachePageSizeKBytes for the instance
func (instance *MSCluster_StorageSpacesDirect) SetPropertyCachePageSizeKBytes(value uint32) (err error) {
	return instance.SetProperty("CachePageSizeKBytes", value)
}

// GetCachePageSizeKBytes gets the value of CachePageSizeKBytes for the instance
func (instance *MSCluster_StorageSpacesDirect) GetPropertyCachePageSizeKBytes() (value uint32, err error) {
	retValue, err := instance.GetProperty("CachePageSizeKBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheState sets the value of CacheState for the instance
func (instance *MSCluster_StorageSpacesDirect) SetPropertyCacheState(value uint32) (err error) {
	return instance.SetProperty("CacheState", value)
}

// GetCacheState gets the value of CacheState for the instance
func (instance *MSCluster_StorageSpacesDirect) GetPropertyCacheState() (value uint32, err error) {
	retValue, err := instance.GetProperty("CacheState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableReportName sets the value of EnableReportName for the instance
func (instance *MSCluster_StorageSpacesDirect) SetPropertyEnableReportName(value string) (err error) {
	return instance.SetProperty("EnableReportName", value)
}

// GetEnableReportName gets the value of EnableReportName for the instance
func (instance *MSCluster_StorageSpacesDirect) GetPropertyEnableReportName() (value string, err error) {
	retValue, err := instance.GetProperty("EnableReportName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSCluster_StorageSpacesDirect) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSCluster_StorageSpacesDirect) GetPropertyName() (value string, err error) {
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

// SetNode sets the value of Node for the instance
func (instance *MSCluster_StorageSpacesDirect) SetPropertyNode(value string) (err error) {
	return instance.SetProperty("Node", value)
}

// GetNode gets the value of Node for the instance
func (instance *MSCluster_StorageSpacesDirect) GetPropertyNode() (value string, err error) {
	retValue, err := instance.GetProperty("Node")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScmUse sets the value of ScmUse for the instance
func (instance *MSCluster_StorageSpacesDirect) SetPropertyScmUse(value uint32) (err error) {
	return instance.SetProperty("ScmUse", value)
}

// GetScmUse gets the value of ScmUse for the instance
func (instance *MSCluster_StorageSpacesDirect) GetPropertyScmUse() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScmUse")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSedProtectionState sets the value of SedProtectionState for the instance
func (instance *MSCluster_StorageSpacesDirect) SetPropertySedProtectionState(value uint32) (err error) {
	return instance.SetProperty("SedProtectionState", value)
}

// GetSedProtectionState gets the value of SedProtectionState for the instance
func (instance *MSCluster_StorageSpacesDirect) GetPropertySedProtectionState() (value uint32, err error) {
	retValue, err := instance.GetProperty("SedProtectionState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSCluster_StorageSpacesDirect) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSCluster_StorageSpacesDirect) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUseSedExclusively sets the value of UseSedExclusively for the instance
func (instance *MSCluster_StorageSpacesDirect) SetPropertyUseSedExclusively(value bool) (err error) {
	return instance.SetProperty("UseSedExclusively", value)
}

// GetUseSedExclusively gets the value of UseSedExclusively for the instance
func (instance *MSCluster_StorageSpacesDirect) GetPropertyUseSedExclusively() (value bool, err error) {
	retValue, err := instance.GetProperty("UseSedExclusively")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="AutoConfig" type="bool "></param>
// <param name="BusTypesToUse" type="uint16 []"></param>
// <param name="CacheDeviceModel" type="string []"></param>
// <param name="CacheMetadataReserveBytes" type="uint64 "></param>
// <param name="CachePageSizeKBytes" type="uint32 "></param>
// <param name="CacheState" type="uint32 "></param>
// <param name="CollectPerformanceHistory" type="bool "></param>
// <param name="PoolFriendlyName" type="string "></param>
// <param name="SedProtectionState" type="uint32 "></param>
// <param name="SkipEligibilityChecks" type="bool "></param>
// <param name="UseSedExclusively" type="bool "></param>
// <param name="XML" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="StorageSpacesDirect" type="MSCluster_StorageSpacesDirect "></param>
func (instance *MSCluster_StorageSpacesDirect) EnableStorageSpacesDirect( /* IN */ CacheState uint32,
	/* IN */ CacheMetadataReserveBytes uint64,
	/* IN */ XML string,
	/* IN */ CacheDeviceModel []string,
	/* IN */ AutoConfig bool,
	/* IN */ CachePageSizeKBytes uint32,
	/* IN */ PoolFriendlyName string,
	/* IN */ SkipEligibilityChecks bool,
	/* IN */ CollectPerformanceHistory bool,
	/* IN */ BusTypesToUse []uint16,
	/* IN */ UseSedExclusively bool,
	/* IN */ SedProtectionState uint32,
	/* OUT */ StorageSpacesDirect MSCluster_StorageSpacesDirect) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnableStorageSpacesDirect", CacheState, CacheMetadataReserveBytes, XML, CacheDeviceModel, AutoConfig, CachePageSizeKBytes, PoolFriendlyName, SkipEligibilityChecks, CollectPerformanceHistory, BusTypesToUse, UseSedExclusively, SedProtectionState)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CleanupCache" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_StorageSpacesDirect) DisableStorageSpacesDirect( /* IN */ CleanupCache bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DisableStorageSpacesDirect", CleanupCache)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="CacheModeHdd" type="uint32 "></param>
// <param name="CacheModeSsd" type="uint32 "></param>
// <param name="CacheState" type="uint32 "></param>
// <param name="Nodes" type="string []"></param>
// <param name="ScmUse" type="uint32 "></param>
// <param name="SedProtectionState" type="uint32 "></param>
// <param name="SkipEligibilityChecks" type="bool "></param>
// <param name="UseSedExclusively" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_StorageSpacesDirect) SetStorageSpacesDirect( /* IN */ CacheState uint32,
	/* IN */ CacheModeHdd uint32,
	/* IN */ CacheModeSsd uint32,
	/* IN */ SkipEligibilityChecks bool,
	/* IN */ ScmUse uint32,
	/* IN */ Nodes []string,
	/* IN */ UseSedExclusively bool,
	/* IN */ SedProtectionState uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetStorageSpacesDirect", CacheState, CacheModeHdd, CacheModeSsd, SkipEligibilityChecks, ScmUse, Nodes, UseSedExclusively, SedProtectionState)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Node" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="TheStorageSpacesDirect" type="MSCluster_StorageSpacesDirect "></param>
func (instance *MSCluster_StorageSpacesDirect) GetStorageSpacesDirect( /* IN */ Node string,
	/* OUT */ TheStorageSpacesDirect MSCluster_StorageSpacesDirect) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetStorageSpacesDirect", Node)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="nNodesInSite" type="uint32 "></param>
// <param name="PoolHasNonCacheHdd" type="bool "></param>
// <param name="PoolHasNonCacheScm" type="bool "></param>
// <param name="PoolHasNonCacheSsd" type="bool "></param>
// <param name="PoolUniqueId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_StorageSpacesDirect) PostCreatePool( /* IN */ PoolUniqueId string,
	/* IN */ nNodesInSite uint32,
	/* IN */ PoolHasNonCacheHdd bool,
	/* IN */ PoolHasNonCacheSsd bool,
	/* IN */ PoolHasNonCacheScm bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("PostCreatePool", PoolUniqueId, nNodesInSite, PoolHasNonCacheHdd, PoolHasNonCacheSsd, PoolHasNonCacheScm)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="DisableStorageMaintenanceMode" type="bool "></param>
// <param name="Node" type="string "></param>
// <param name="RecoverUnboundDrives" type="bool "></param>
// <param name="SkipDiskRecovery" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_StorageSpacesDirect) RepairStorageSpacesDirect( /* IN */ SkipDiskRecovery bool,
	/* IN */ DisableStorageMaintenanceMode bool,
	/* IN */ RecoverUnboundDrives bool,
	/* IN */ Node string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RepairStorageSpacesDirect", SkipDiskRecovery, DisableStorageMaintenanceMode, RecoverUnboundDrives, Node)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="CacheUsage" type="uint32 "></param>
// <param name="CanBeClaimed" type="bool "></param>
// <param name="PhysicalDiskIds" type="string []"></param>
// <param name="Reset" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_StorageSpacesDirect) SetStorageSpacesDirectDisk( /* IN */ CanBeClaimed bool,
	/* IN */ Reset bool,
	/* IN */ CacheUsage uint32,
	/* IN */ PhysicalDiskIds []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetStorageSpacesDirectDisk", CanBeClaimed, Reset, CacheUsage, PhysicalDiskIds)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="CanBeClaimed" type="bool "></param>

// <param name="PhysicalDiskIds" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_StorageSpacesDirect) GetStorageSpacesDirectDisk( /* IN */ CanBeClaimed bool,
	/* OUT */ PhysicalDiskIds []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetStorageSpacesDirectDisk", CanBeClaimed)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
