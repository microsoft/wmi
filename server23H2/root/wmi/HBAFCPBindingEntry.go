// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// HBAFCPBindingEntry struct
type HBAFCPBindingEntry struct {
	*cim.WmiInstance

	//
	FCPId HBAFCPID

	//
	ScsiId HBAScsiID

	//
	Type uint32
}

func NewHBAFCPBindingEntryEx1(instance *cim.WmiInstance) (newInstance *HBAFCPBindingEntry, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HBAFCPBindingEntry{
		WmiInstance: tmp,
	}
	return
}

func NewHBAFCPBindingEntryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HBAFCPBindingEntry, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HBAFCPBindingEntry{
		WmiInstance: tmp,
	}
	return
}

// SetFCPId sets the value of FCPId for the instance
func (instance *HBAFCPBindingEntry) SetPropertyFCPId(value HBAFCPID) (err error) {
	return instance.SetProperty("FCPId", (value))
}

// GetFCPId gets the value of FCPId for the instance
func (instance *HBAFCPBindingEntry) GetPropertyFCPId() (value HBAFCPID, err error) {
	retValue, err := instance.GetProperty("FCPId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(HBAFCPID)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " HBAFCPID is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = HBAFCPID(valuetmp)

	return
}

// SetScsiId sets the value of ScsiId for the instance
func (instance *HBAFCPBindingEntry) SetPropertyScsiId(value HBAScsiID) (err error) {
	return instance.SetProperty("ScsiId", (value))
}

// GetScsiId gets the value of ScsiId for the instance
func (instance *HBAFCPBindingEntry) GetPropertyScsiId() (value HBAScsiID, err error) {
	retValue, err := instance.GetProperty("ScsiId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(HBAScsiID)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " HBAScsiID is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = HBAScsiID(valuetmp)

	return
}

// SetType sets the value of Type for the instance
func (instance *HBAFCPBindingEntry) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *HBAFCPBindingEntry) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
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
