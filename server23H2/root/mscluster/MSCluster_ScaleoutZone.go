// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_ScaleoutZone struct
type MSCluster_ScaleoutZone struct {
	*cim.WmiInstance

	//
	ClusterSize uint32

	//
	Flags uint64

	//
	Size uint64

	//
	Type uint16

	//
	ZoneGroupId string

	//
	ZoneId string

	//
	ZoneNumber uint32
}

func NewMSCluster_ScaleoutZoneEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ScaleoutZone, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_ScaleoutZone{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_ScaleoutZoneEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ScaleoutZone, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ScaleoutZone{
		WmiInstance: tmp,
	}
	return
}

// SetClusterSize sets the value of ClusterSize for the instance
func (instance *MSCluster_ScaleoutZone) SetPropertyClusterSize(value uint32) (err error) {
	return instance.SetProperty("ClusterSize", (value))
}

// GetClusterSize gets the value of ClusterSize for the instance
func (instance *MSCluster_ScaleoutZone) GetPropertyClusterSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClusterSize")
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

// SetFlags sets the value of Flags for the instance
func (instance *MSCluster_ScaleoutZone) SetPropertyFlags(value uint64) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSCluster_ScaleoutZone) GetPropertyFlags() (value uint64, err error) {
	retValue, err := instance.GetProperty("Flags")
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

// SetSize sets the value of Size for the instance
func (instance *MSCluster_ScaleoutZone) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSCluster_ScaleoutZone) GetPropertySize() (value uint64, err error) {
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

// SetType sets the value of Type for the instance
func (instance *MSCluster_ScaleoutZone) SetPropertyType(value uint16) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSCluster_ScaleoutZone) GetPropertyType() (value uint16, err error) {
	retValue, err := instance.GetProperty("Type")
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

// SetZoneGroupId sets the value of ZoneGroupId for the instance
func (instance *MSCluster_ScaleoutZone) SetPropertyZoneGroupId(value string) (err error) {
	return instance.SetProperty("ZoneGroupId", (value))
}

// GetZoneGroupId gets the value of ZoneGroupId for the instance
func (instance *MSCluster_ScaleoutZone) GetPropertyZoneGroupId() (value string, err error) {
	retValue, err := instance.GetProperty("ZoneGroupId")
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

// SetZoneId sets the value of ZoneId for the instance
func (instance *MSCluster_ScaleoutZone) SetPropertyZoneId(value string) (err error) {
	return instance.SetProperty("ZoneId", (value))
}

// GetZoneId gets the value of ZoneId for the instance
func (instance *MSCluster_ScaleoutZone) GetPropertyZoneId() (value string, err error) {
	retValue, err := instance.GetProperty("ZoneId")
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

// SetZoneNumber sets the value of ZoneNumber for the instance
func (instance *MSCluster_ScaleoutZone) SetPropertyZoneNumber(value uint32) (err error) {
	return instance.SetProperty("ZoneNumber", (value))
}

// GetZoneNumber gets the value of ZoneNumber for the instance
func (instance *MSCluster_ScaleoutZone) GetPropertyZoneNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("ZoneNumber")
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
