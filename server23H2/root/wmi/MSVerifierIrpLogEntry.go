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

// MSVerifierIrpLogEntry struct
type MSVerifierIrpLogEntry struct {
	*cim.WmiInstance

	// Arg1
	Arg1 uint64

	// Arg2
	Arg2 uint64

	// Arg3
	Arg3 uint64

	// Arg4
	Arg4 uint64

	// Control
	Control uint8

	// Count
	Count uint32

	// Flags
	Flags uint8

	// Major Function
	Major uint8

	// Minor Function
	Minor uint8
}

func NewMSVerifierIrpLogEntryEx1(instance *cim.WmiInstance) (newInstance *MSVerifierIrpLogEntry, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSVerifierIrpLogEntry{
		WmiInstance: tmp,
	}
	return
}

func NewMSVerifierIrpLogEntryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSVerifierIrpLogEntry, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSVerifierIrpLogEntry{
		WmiInstance: tmp,
	}
	return
}

// SetArg1 sets the value of Arg1 for the instance
func (instance *MSVerifierIrpLogEntry) SetPropertyArg1(value uint64) (err error) {
	return instance.SetProperty("Arg1", (value))
}

// GetArg1 gets the value of Arg1 for the instance
func (instance *MSVerifierIrpLogEntry) GetPropertyArg1() (value uint64, err error) {
	retValue, err := instance.GetProperty("Arg1")
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

// SetArg2 sets the value of Arg2 for the instance
func (instance *MSVerifierIrpLogEntry) SetPropertyArg2(value uint64) (err error) {
	return instance.SetProperty("Arg2", (value))
}

// GetArg2 gets the value of Arg2 for the instance
func (instance *MSVerifierIrpLogEntry) GetPropertyArg2() (value uint64, err error) {
	retValue, err := instance.GetProperty("Arg2")
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

// SetArg3 sets the value of Arg3 for the instance
func (instance *MSVerifierIrpLogEntry) SetPropertyArg3(value uint64) (err error) {
	return instance.SetProperty("Arg3", (value))
}

// GetArg3 gets the value of Arg3 for the instance
func (instance *MSVerifierIrpLogEntry) GetPropertyArg3() (value uint64, err error) {
	retValue, err := instance.GetProperty("Arg3")
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

// SetArg4 sets the value of Arg4 for the instance
func (instance *MSVerifierIrpLogEntry) SetPropertyArg4(value uint64) (err error) {
	return instance.SetProperty("Arg4", (value))
}

// GetArg4 gets the value of Arg4 for the instance
func (instance *MSVerifierIrpLogEntry) GetPropertyArg4() (value uint64, err error) {
	retValue, err := instance.GetProperty("Arg4")
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

// SetControl sets the value of Control for the instance
func (instance *MSVerifierIrpLogEntry) SetPropertyControl(value uint8) (err error) {
	return instance.SetProperty("Control", (value))
}

// GetControl gets the value of Control for the instance
func (instance *MSVerifierIrpLogEntry) GetPropertyControl() (value uint8, err error) {
	retValue, err := instance.GetProperty("Control")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetCount sets the value of Count for the instance
func (instance *MSVerifierIrpLogEntry) SetPropertyCount(value uint32) (err error) {
	return instance.SetProperty("Count", (value))
}

// GetCount gets the value of Count for the instance
func (instance *MSVerifierIrpLogEntry) GetPropertyCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("Count")
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
func (instance *MSVerifierIrpLogEntry) SetPropertyFlags(value uint8) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSVerifierIrpLogEntry) GetPropertyFlags() (value uint8, err error) {
	retValue, err := instance.GetProperty("Flags")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetMajor sets the value of Major for the instance
func (instance *MSVerifierIrpLogEntry) SetPropertyMajor(value uint8) (err error) {
	return instance.SetProperty("Major", (value))
}

// GetMajor gets the value of Major for the instance
func (instance *MSVerifierIrpLogEntry) GetPropertyMajor() (value uint8, err error) {
	retValue, err := instance.GetProperty("Major")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetMinor sets the value of Minor for the instance
func (instance *MSVerifierIrpLogEntry) SetPropertyMinor(value uint8) (err error) {
	return instance.SetProperty("Minor", (value))
}

// GetMinor gets the value of Minor for the instance
func (instance *MSVerifierIrpLogEntry) GetPropertyMinor() (value uint8, err error) {
	retValue, err := instance.GetProperty("Minor")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}
