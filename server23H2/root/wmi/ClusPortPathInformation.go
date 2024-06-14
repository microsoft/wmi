// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ClusPortPathInformation struct
type ClusPortPathInformation struct {
	*cim.WmiInstance

	// InActiveList.
	InActiveList bool

	// InReadOnlyList.
	InReadOnlyList bool

	// InStandbyList.
	InStandbyList bool

	// IsReadOnly.
	IsReadOnly bool

	// Node Id.
	NodeId uint32

	// NodeName .
	NodeName string

	// Path Id.
	PathId uint32

	// Path Key.
	PathKey uint64

	// ClusPort Device State.
	PathState uint32
}

func NewClusPortPathInformationEx1(instance *cim.WmiInstance) (newInstance *ClusPortPathInformation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ClusPortPathInformation{
		WmiInstance: tmp,
	}
	return
}

func NewClusPortPathInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ClusPortPathInformation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ClusPortPathInformation{
		WmiInstance: tmp,
	}
	return
}

// SetInActiveList sets the value of InActiveList for the instance
func (instance *ClusPortPathInformation) SetPropertyInActiveList(value bool) (err error) {
	return instance.SetProperty("InActiveList", (value))
}

// GetInActiveList gets the value of InActiveList for the instance
func (instance *ClusPortPathInformation) GetPropertyInActiveList() (value bool, err error) {
	retValue, err := instance.GetProperty("InActiveList")
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

// SetInReadOnlyList sets the value of InReadOnlyList for the instance
func (instance *ClusPortPathInformation) SetPropertyInReadOnlyList(value bool) (err error) {
	return instance.SetProperty("InReadOnlyList", (value))
}

// GetInReadOnlyList gets the value of InReadOnlyList for the instance
func (instance *ClusPortPathInformation) GetPropertyInReadOnlyList() (value bool, err error) {
	retValue, err := instance.GetProperty("InReadOnlyList")
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

// SetInStandbyList sets the value of InStandbyList for the instance
func (instance *ClusPortPathInformation) SetPropertyInStandbyList(value bool) (err error) {
	return instance.SetProperty("InStandbyList", (value))
}

// GetInStandbyList gets the value of InStandbyList for the instance
func (instance *ClusPortPathInformation) GetPropertyInStandbyList() (value bool, err error) {
	retValue, err := instance.GetProperty("InStandbyList")
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

// SetIsReadOnly sets the value of IsReadOnly for the instance
func (instance *ClusPortPathInformation) SetPropertyIsReadOnly(value bool) (err error) {
	return instance.SetProperty("IsReadOnly", (value))
}

// GetIsReadOnly gets the value of IsReadOnly for the instance
func (instance *ClusPortPathInformation) GetPropertyIsReadOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("IsReadOnly")
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

// SetNodeId sets the value of NodeId for the instance
func (instance *ClusPortPathInformation) SetPropertyNodeId(value uint32) (err error) {
	return instance.SetProperty("NodeId", (value))
}

// GetNodeId gets the value of NodeId for the instance
func (instance *ClusPortPathInformation) GetPropertyNodeId() (value uint32, err error) {
	retValue, err := instance.GetProperty("NodeId")
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

// SetNodeName sets the value of NodeName for the instance
func (instance *ClusPortPathInformation) SetPropertyNodeName(value string) (err error) {
	return instance.SetProperty("NodeName", (value))
}

// GetNodeName gets the value of NodeName for the instance
func (instance *ClusPortPathInformation) GetPropertyNodeName() (value string, err error) {
	retValue, err := instance.GetProperty("NodeName")
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

// SetPathId sets the value of PathId for the instance
func (instance *ClusPortPathInformation) SetPropertyPathId(value uint32) (err error) {
	return instance.SetProperty("PathId", (value))
}

// GetPathId gets the value of PathId for the instance
func (instance *ClusPortPathInformation) GetPropertyPathId() (value uint32, err error) {
	retValue, err := instance.GetProperty("PathId")
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

// SetPathKey sets the value of PathKey for the instance
func (instance *ClusPortPathInformation) SetPropertyPathKey(value uint64) (err error) {
	return instance.SetProperty("PathKey", (value))
}

// GetPathKey gets the value of PathKey for the instance
func (instance *ClusPortPathInformation) GetPropertyPathKey() (value uint64, err error) {
	retValue, err := instance.GetProperty("PathKey")
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

// SetPathState sets the value of PathState for the instance
func (instance *ClusPortPathInformation) SetPropertyPathState(value uint32) (err error) {
	return instance.SetProperty("PathState", (value))
}

// GetPathState gets the value of PathState for the instance
func (instance *ClusPortPathInformation) GetPropertyPathState() (value uint32, err error) {
	retValue, err := instance.GetProperty("PathState")
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
