// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Hgs
//////////////////////////////////////////////
package hgs

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_HgsKeyProtector struct
type MSFT_HgsKeyProtector struct {
	cim.WmiInstance

	//
	Guardians []MSFT_HgsGuardian

	//
	Owner MSFT_HgsGuardian

	//
	RawData []uint8
}

// SetGuardians sets the value of Guardians for the instance
func (instance *MSFT_HgsKeyProtector) SetPropertyGuardians(value []MSFT_HgsGuardian) (err error) {
	return instance.SetProperty("Guardians", value)
}

// GetGuardians gets the value of Guardians for the instance
func (instance *MSFT_HgsKeyProtector) GetPropertyGuardians() (value []MSFT_HgsGuardian, err error) {
	retValue, err := instance.GetProperty("Guardians")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_HgsGuardian)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOwner sets the value of Owner for the instance
func (instance *MSFT_HgsKeyProtector) SetPropertyOwner(value MSFT_HgsGuardian) (err error) {
	return instance.SetProperty("Owner", value)
}

// GetOwner gets the value of Owner for the instance
func (instance *MSFT_HgsKeyProtector) GetPropertyOwner() (value MSFT_HgsGuardian, err error) {
	retValue, err := instance.GetProperty("Owner")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_HgsGuardian)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRawData sets the value of RawData for the instance
func (instance *MSFT_HgsKeyProtector) SetPropertyRawData(value []uint8) (err error) {
	return instance.SetProperty("RawData", value)
}

// GetRawData gets the value of RawData for the instance
func (instance *MSFT_HgsKeyProtector) GetPropertyRawData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("RawData")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
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
