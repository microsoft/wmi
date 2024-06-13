// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
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

// HBAFCPScsiEntry struct
type HBAFCPScsiEntry struct {
	*cim.WmiInstance

	//
	FCPId HBAFCPID

	//
	Luid []uint8

	//
	ScsiId HBAScsiID
}

func NewHBAFCPScsiEntryEx1(instance *cim.WmiInstance) (newInstance *HBAFCPScsiEntry, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HBAFCPScsiEntry{
		WmiInstance: tmp,
	}
	return
}

func NewHBAFCPScsiEntryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HBAFCPScsiEntry, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HBAFCPScsiEntry{
		WmiInstance: tmp,
	}
	return
}

// SetFCPId sets the value of FCPId for the instance
func (instance *HBAFCPScsiEntry) SetPropertyFCPId(value HBAFCPID) (err error) {
	return instance.SetProperty("FCPId", (value))
}

// GetFCPId gets the value of FCPId for the instance
func (instance *HBAFCPScsiEntry) GetPropertyFCPId() (value HBAFCPID, err error) {
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

// SetLuid sets the value of Luid for the instance
func (instance *HBAFCPScsiEntry) SetPropertyLuid(value []uint8) (err error) {
	return instance.SetProperty("Luid", (value))
}

// GetLuid gets the value of Luid for the instance
func (instance *HBAFCPScsiEntry) GetPropertyLuid() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Luid")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetScsiId sets the value of ScsiId for the instance
func (instance *HBAFCPScsiEntry) SetPropertyScsiId(value HBAScsiID) (err error) {
	return instance.SetProperty("ScsiId", (value))
}

// GetScsiId gets the value of ScsiId for the instance
func (instance *HBAFCPScsiEntry) GetPropertyScsiId() (value HBAScsiID, err error) {
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
