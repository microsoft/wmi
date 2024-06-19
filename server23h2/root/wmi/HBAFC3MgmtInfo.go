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

// HBAFC3MgmtInfo struct
type HBAFC3MgmtInfo struct {
	*cim.WmiInstance

	//
	IPAddress []uint8

	//
	IPVersion uint16

	//
	NumberOfAttachedNodes uint32

	//
	PortId uint32

	//
	reserved uint16

	//
	reserved1 uint32

	//
	TopologyDiscoveryFlags uint16

	//
	UDPPort uint16

	//
	UniqueAdapterId uint64

	//
	unittype uint32

	//
	wwn []uint8
}

func NewHBAFC3MgmtInfoEx1(instance *cim.WmiInstance) (newInstance *HBAFC3MgmtInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HBAFC3MgmtInfo{
		WmiInstance: tmp,
	}
	return
}

func NewHBAFC3MgmtInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HBAFC3MgmtInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HBAFC3MgmtInfo{
		WmiInstance: tmp,
	}
	return
}

// SetIPAddress sets the value of IPAddress for the instance
func (instance *HBAFC3MgmtInfo) SetPropertyIPAddress(value []uint8) (err error) {
	return instance.SetProperty("IPAddress", (value))
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *HBAFC3MgmtInfo) GetPropertyIPAddress() (value []uint8, err error) {
	retValue, err := instance.GetProperty("IPAddress")
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

// SetIPVersion sets the value of IPVersion for the instance
func (instance *HBAFC3MgmtInfo) SetPropertyIPVersion(value uint16) (err error) {
	return instance.SetProperty("IPVersion", (value))
}

// GetIPVersion gets the value of IPVersion for the instance
func (instance *HBAFC3MgmtInfo) GetPropertyIPVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("IPVersion")
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

// SetNumberOfAttachedNodes sets the value of NumberOfAttachedNodes for the instance
func (instance *HBAFC3MgmtInfo) SetPropertyNumberOfAttachedNodes(value uint32) (err error) {
	return instance.SetProperty("NumberOfAttachedNodes", (value))
}

// GetNumberOfAttachedNodes gets the value of NumberOfAttachedNodes for the instance
func (instance *HBAFC3MgmtInfo) GetPropertyNumberOfAttachedNodes() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfAttachedNodes")
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

// SetPortId sets the value of PortId for the instance
func (instance *HBAFC3MgmtInfo) SetPropertyPortId(value uint32) (err error) {
	return instance.SetProperty("PortId", (value))
}

// GetPortId gets the value of PortId for the instance
func (instance *HBAFC3MgmtInfo) GetPropertyPortId() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortId")
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

// Setreserved sets the value of reserved for the instance
func (instance *HBAFC3MgmtInfo) SetPropertyreserved(value uint16) (err error) {
	return instance.SetProperty("reserved", (value))
}

// Getreserved gets the value of reserved for the instance
func (instance *HBAFC3MgmtInfo) GetPropertyreserved() (value uint16, err error) {
	retValue, err := instance.GetProperty("reserved")
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

// Setreserved1 sets the value of reserved1 for the instance
func (instance *HBAFC3MgmtInfo) SetPropertyreserved1(value uint32) (err error) {
	return instance.SetProperty("reserved1", (value))
}

// Getreserved1 gets the value of reserved1 for the instance
func (instance *HBAFC3MgmtInfo) GetPropertyreserved1() (value uint32, err error) {
	retValue, err := instance.GetProperty("reserved1")
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

// SetTopologyDiscoveryFlags sets the value of TopologyDiscoveryFlags for the instance
func (instance *HBAFC3MgmtInfo) SetPropertyTopologyDiscoveryFlags(value uint16) (err error) {
	return instance.SetProperty("TopologyDiscoveryFlags", (value))
}

// GetTopologyDiscoveryFlags gets the value of TopologyDiscoveryFlags for the instance
func (instance *HBAFC3MgmtInfo) GetPropertyTopologyDiscoveryFlags() (value uint16, err error) {
	retValue, err := instance.GetProperty("TopologyDiscoveryFlags")
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

// SetUDPPort sets the value of UDPPort for the instance
func (instance *HBAFC3MgmtInfo) SetPropertyUDPPort(value uint16) (err error) {
	return instance.SetProperty("UDPPort", (value))
}

// GetUDPPort gets the value of UDPPort for the instance
func (instance *HBAFC3MgmtInfo) GetPropertyUDPPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("UDPPort")
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

// SetUniqueAdapterId sets the value of UniqueAdapterId for the instance
func (instance *HBAFC3MgmtInfo) SetPropertyUniqueAdapterId(value uint64) (err error) {
	return instance.SetProperty("UniqueAdapterId", (value))
}

// GetUniqueAdapterId gets the value of UniqueAdapterId for the instance
func (instance *HBAFC3MgmtInfo) GetPropertyUniqueAdapterId() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniqueAdapterId")
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

// Setunittype sets the value of unittype for the instance
func (instance *HBAFC3MgmtInfo) SetPropertyunittype(value uint32) (err error) {
	return instance.SetProperty("unittype", (value))
}

// Getunittype gets the value of unittype for the instance
func (instance *HBAFC3MgmtInfo) GetPropertyunittype() (value uint32, err error) {
	retValue, err := instance.GetProperty("unittype")
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

// Setwwn sets the value of wwn for the instance
func (instance *HBAFC3MgmtInfo) SetPropertywwn(value []uint8) (err error) {
	return instance.SetProperty("wwn", (value))
}

// Getwwn gets the value of wwn for the instance
func (instance *HBAFC3MgmtInfo) GetPropertywwn() (value []uint8, err error) {
	retValue, err := instance.GetProperty("wwn")
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
