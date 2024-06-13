// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Hgs
//////////////////////////////////////////////
package hgs

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_HgsKeyProtector struct
type MSFT_HgsKeyProtector struct {
	*cim.WmiInstance

	//
	Guardians []MSFT_HgsGuardian

	//
	Owner MSFT_HgsGuardian

	//
	RawData []uint8
}

func NewMSFT_HgsKeyProtectorEx1(instance *cim.WmiInstance) (newInstance *MSFT_HgsKeyProtector, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_HgsKeyProtector{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_HgsKeyProtectorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_HgsKeyProtector, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_HgsKeyProtector{
		WmiInstance: tmp,
	}
	return
}

// SetGuardians sets the value of Guardians for the instance
func (instance *MSFT_HgsKeyProtector) SetPropertyGuardians(value []MSFT_HgsGuardian) (err error) {
	return instance.SetProperty("Guardians", (value))
}

// GetGuardians gets the value of Guardians for the instance
func (instance *MSFT_HgsKeyProtector) GetPropertyGuardians() (value []MSFT_HgsGuardian, err error) {
	retValue, err := instance.GetProperty("Guardians")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_HgsGuardian)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_HgsGuardian is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_HgsGuardian(valuetmp))
	}

	return
}

// SetOwner sets the value of Owner for the instance
func (instance *MSFT_HgsKeyProtector) SetPropertyOwner(value MSFT_HgsGuardian) (err error) {
	return instance.SetProperty("Owner", (value))
}

// GetOwner gets the value of Owner for the instance
func (instance *MSFT_HgsKeyProtector) GetPropertyOwner() (value MSFT_HgsGuardian, err error) {
	retValue, err := instance.GetProperty("Owner")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_HgsGuardian)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_HgsGuardian is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_HgsGuardian(valuetmp)

	return
}

// SetRawData sets the value of RawData for the instance
func (instance *MSFT_HgsKeyProtector) SetPropertyRawData(value []uint8) (err error) {
	return instance.SetProperty("RawData", (value))
}

// GetRawData gets the value of RawData for the instance
func (instance *MSFT_HgsKeyProtector) GetPropertyRawData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("RawData")
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

//

// <param name="AllowExpired" type="bool "></param>
// <param name="AllowUntrustedRoot" type="bool "></param>
// <param name="Guardian" type="MSFT_HgsGuardian []"></param>
// <param name="Owner" type="MSFT_HgsGuardian "></param>

// <param name="cmdletOutput" type="MSFT_HgsKeyProtector "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsKeyProtector) NewByGuardians( /* IN */ AllowUntrustedRoot bool,
	/* IN */ AllowExpired bool,
	/* IN */ Owner MSFT_HgsGuardian,
	/* IN */ Guardian []MSFT_HgsGuardian,
	/* OUT */ cmdletOutput MSFT_HgsKeyProtector) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewByGuardians", AllowUntrustedRoot, AllowExpired, Owner, Guardian)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AllowExpired" type="bool "></param>
// <param name="AllowUntrustedRoot" type="bool "></param>
// <param name="Guardian" type="MSFT_HgsGuardian "></param>
// <param name="KeyProtector" type="MSFT_HgsKeyProtector "></param>

// <param name="cmdletOutput" type="MSFT_HgsKeyProtector "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsKeyProtector) Grant( /* IN */ KeyProtector MSFT_HgsKeyProtector,
	/* IN */ Guardian MSFT_HgsGuardian,
	/* IN */ AllowExpired bool,
	/* IN */ AllowUntrustedRoot bool,
	/* OUT */ cmdletOutput MSFT_HgsKeyProtector) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Grant", KeyProtector, Guardian, AllowExpired, AllowUntrustedRoot)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Guardian" type="MSFT_HgsGuardian "></param>
// <param name="KeyProtector" type="MSFT_HgsKeyProtector "></param>

// <param name="cmdletOutput" type="MSFT_HgsKeyProtector "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsKeyProtector) Revoke( /* IN */ KeyProtector MSFT_HgsKeyProtector,
	/* IN */ Guardian MSFT_HgsGuardian,
	/* OUT */ cmdletOutput MSFT_HgsKeyProtector) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Revoke", KeyProtector, Guardian)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Bytes" type="uint8 []"></param>

// <param name="cmdletOutput" type="MSFT_HgsKeyProtector "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsKeyProtector) ConvertToByRawBytes( /* IN */ Bytes []uint8,
	/* OUT */ cmdletOutput MSFT_HgsKeyProtector) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ConvertToByRawBytes", Bytes)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
