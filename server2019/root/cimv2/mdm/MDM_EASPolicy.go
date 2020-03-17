// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_EASPolicy struct
type MDM_EASPolicy struct {
	cim.WmiInstance

	//
	key uint32
}

// Setkey sets the value of key for the instance
func (instance *MDM_EASPolicy) SetPropertykey(value uint32) (err error) {
	return instance.SetProperty("key", value)
}

// Getkey gets the value of key for the instance
func (instance *MDM_EASPolicy) GetPropertykey() (value uint32, err error) {
	retValue, err := instance.GetProperty("key")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="NamedValuesList" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_EASPolicy) SetValues( /* IN */ NamedValuesList string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetValues", NamedValuesList)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
