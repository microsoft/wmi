// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.MSCluster
//
// ////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_ScaleoutVolume struct
type MSCluster_ScaleoutVolume struct {
	*cim.WmiInstance

	//
	AllocationSize uint64

	//
	FriendlyName string

	//
	numZones uint32

	//
	SizeInBytes uint64

	//
	VolumeGuid string

	//
	zoneArray []MSCluster_ScaleoutZone
}

func NewMSCluster_ScaleoutVolumeEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ScaleoutVolume, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_ScaleoutVolume{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_ScaleoutVolumeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ScaleoutVolume, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ScaleoutVolume{
		WmiInstance: tmp,
	}
	return
}

// SetAllocationSize sets the value of AllocationSize for the instance
func (instance *MSCluster_ScaleoutVolume) SetPropertyAllocationSize(value uint64) (err error) {
	return instance.SetProperty("AllocationSize", (value))
}

// GetAllocationSize gets the value of AllocationSize for the instance
func (instance *MSCluster_ScaleoutVolume) GetPropertyAllocationSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllocationSize")
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

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSCluster_ScaleoutVolume) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", (value))
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSCluster_ScaleoutVolume) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
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

// SetnumZones sets the value of numZones for the instance
func (instance *MSCluster_ScaleoutVolume) SetPropertynumZones(value uint32) (err error) {
	return instance.SetProperty("numZones", (value))
}

// GetnumZones gets the value of numZones for the instance
func (instance *MSCluster_ScaleoutVolume) GetPropertynumZones() (value uint32, err error) {
	retValue, err := instance.GetProperty("numZones")
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

// SetSizeInBytes sets the value of SizeInBytes for the instance
func (instance *MSCluster_ScaleoutVolume) SetPropertySizeInBytes(value uint64) (err error) {
	return instance.SetProperty("SizeInBytes", (value))
}

// GetSizeInBytes gets the value of SizeInBytes for the instance
func (instance *MSCluster_ScaleoutVolume) GetPropertySizeInBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("SizeInBytes")
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

// SetVolumeGuid sets the value of VolumeGuid for the instance
func (instance *MSCluster_ScaleoutVolume) SetPropertyVolumeGuid(value string) (err error) {
	return instance.SetProperty("VolumeGuid", (value))
}

// GetVolumeGuid gets the value of VolumeGuid for the instance
func (instance *MSCluster_ScaleoutVolume) GetPropertyVolumeGuid() (value string, err error) {
	retValue, err := instance.GetProperty("VolumeGuid")
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

// SetzoneArray sets the value of zoneArray for the instance
func (instance *MSCluster_ScaleoutVolume) SetPropertyzoneArray(value []MSCluster_ScaleoutZone) (err error) {
	return instance.SetProperty("zoneArray", (value))
}

// GetzoneArray gets the value of zoneArray for the instance
func (instance *MSCluster_ScaleoutVolume) GetPropertyzoneArray() (value []MSCluster_ScaleoutZone, err error) {
	retValue, err := instance.GetProperty("zoneArray")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSCluster_ScaleoutZone)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSCluster_ScaleoutZone is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSCluster_ScaleoutZone(valuetmp))
	}

	return
}

//

// <param name="MaxMetadataVolumeSize" type="uint64 "></param>
// <param name="VolumeName" type="string "></param>
// <param name="ZoneGroupId" type="string "></param>

// <param name="CreatedScaleoutVolume" type="MSCluster_ScaleoutVolume "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ScaleoutVolume) NewScaleoutVolume( /* IN */ VolumeName string,
	/* IN */ MaxMetadataVolumeSize uint64,
	/* IN */ ZoneGroupId string,
	/* OUT */ CreatedScaleoutVolume MSCluster_ScaleoutVolume) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewScaleoutVolume", VolumeName, MaxMetadataVolumeSize, ZoneGroupId)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="sizeInBytes" type="uint64 "></param>
// <param name="ZoneGroupId" type="string "></param>
// <param name="ZoneId" type="string "></param>
// <param name="ZoneResource" type="string "></param>
// <param name="ZoneTargetPath" type="string "></param>
// <param name="ZoneVolume" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Status" type="int32 "></param>
func (instance *MSCluster_ScaleoutVolume) AddDataZoneToSV( /* IN */ ZoneId string,
	/* IN */ ZoneGroupId string,
	/* IN */ sizeInBytes uint64,
	/* IN */ ZoneTargetPath string,
	/* IN */ ZoneVolume string,
	/* IN */ ZoneResource string,
	/* OUT */ Status int32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddDataZoneToSV", ZoneId, ZoneGroupId, sizeInBytes, ZoneTargetPath, ZoneVolume, ZoneResource)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ClusterName" type="string "></param>
// <param name="force" type="bool "></param>
// <param name="ResourceName" type="string "></param>
// <param name="VolumeName" type="string "></param>
// <param name="ZoneId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Status" type="int32 "></param>
func (instance *MSCluster_ScaleoutVolume) RemoveSVDataZone( /* IN */ VolumeName string,
	/* IN */ ZoneId string,
	/* IN */ ClusterName string,
	/* IN */ ResourceName string,
	/* IN */ force bool,
	/* OUT */ Status int32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveSVDataZone", VolumeName, ZoneId, ClusterName, ResourceName, force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="VolumeName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Status" type="int32 "></param>
func (instance *MSCluster_ScaleoutVolume) DeleteSV( /* IN */ VolumeName string,
	/* OUT */ Status int32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DeleteSV", VolumeName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="VolumeName" type="string "></param>

// <param name="RetrievedScaleoutVolume" type="MSCluster_ScaleoutVolume "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ScaleoutVolume) GetSVInformation( /* IN */ VolumeName string,
	/* OUT */ RetrievedScaleoutVolume MSCluster_ScaleoutVolume) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSVInformation", VolumeName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Guids" type="MSCluster_ScaleoutVolume []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ScaleoutVolume) GetAllSV( /* OUT */ Guids []MSCluster_ScaleoutVolume) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetAllSV")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="zoneId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Status" type="int32 "></param>
func (instance *MSCluster_ScaleoutVolume) SuspendSVDataZone( /* IN */ zoneId string,
	/* OUT */ Status int32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SuspendSVDataZone", zoneId)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="zoneId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Status" type="int32 "></param>
func (instance *MSCluster_ScaleoutVolume) ResumeSVDataZone( /* IN */ zoneId string,
	/* OUT */ Status int32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ResumeSVDataZone", zoneId)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="zoneId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Status" type="int32 "></param>
func (instance *MSCluster_ScaleoutVolume) RetireSVDataZone( /* IN */ zoneId string,
	/* OUT */ Status int32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RetireSVDataZone", zoneId)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="zoneGroupId" type="string "></param>
// <param name="zoneId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Status" type="int32 "></param>
func (instance *MSCluster_ScaleoutVolume) UpdateGroupIdForSVDataZone( /* IN */ zoneId string,
	/* IN */ zoneGroupId string,
	/* OUT */ Status int32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("UpdateGroupIdForSVDataZone", zoneId, zoneGroupId)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="zoneId" type="string "></param>

// <param name="RetrievedScaleoutZone" type="MSCluster_ScaleoutZone "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ScaleoutVolume) GetZoneInformation( /* IN */ zoneId string,
	/* OUT */ RetrievedScaleoutZone MSCluster_ScaleoutZone) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetZoneInformation", zoneId)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
