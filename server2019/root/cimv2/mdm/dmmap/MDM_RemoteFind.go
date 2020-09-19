// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_RemoteFind struct
type MDM_RemoteFind struct {
	*cim.WmiInstance

	//
	DesiredAccuracy int32

	//
	InstanceID string

	//
	MaximumAge int32

	//
	ParentID string

	//
	Timeout int32
}

func NewMDM_RemoteFindEx1(instance *cim.WmiInstance) (newInstance *MDM_RemoteFind, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_RemoteFind{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_RemoteFindEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_RemoteFind, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_RemoteFind{
		WmiInstance: tmp,
	}
	return
}

// SetDesiredAccuracy sets the value of DesiredAccuracy for the instance
func (instance *MDM_RemoteFind) SetPropertyDesiredAccuracy(value int32) (err error) {
	return instance.SetProperty("DesiredAccuracy", (value))
}

// GetDesiredAccuracy gets the value of DesiredAccuracy for the instance
func (instance *MDM_RemoteFind) GetPropertyDesiredAccuracy() (value int32, err error) {
	retValue, err := instance.GetProperty("DesiredAccuracy")
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

	value = int32(valuetmp)

	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_RemoteFind) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_RemoteFind) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetMaximumAge sets the value of MaximumAge for the instance
func (instance *MDM_RemoteFind) SetPropertyMaximumAge(value int32) (err error) {
	return instance.SetProperty("MaximumAge", (value))
}

// GetMaximumAge gets the value of MaximumAge for the instance
func (instance *MDM_RemoteFind) GetPropertyMaximumAge() (value int32, err error) {
	retValue, err := instance.GetProperty("MaximumAge")
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

	value = int32(valuetmp)

	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_RemoteFind) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_RemoteFind) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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

// SetTimeout sets the value of Timeout for the instance
func (instance *MDM_RemoteFind) SetPropertyTimeout(value int32) (err error) {
	return instance.SetProperty("Timeout", (value))
}

// GetTimeout gets the value of Timeout for the instance
func (instance *MDM_RemoteFind) GetPropertyTimeout() (value int32, err error) {
	retValue, err := instance.GetProperty("Timeout")
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

	value = int32(valuetmp)

	return
}
