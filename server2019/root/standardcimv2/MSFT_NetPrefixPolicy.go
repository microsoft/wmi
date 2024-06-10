// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetPrefixPolicy struct
type MSFT_NetPrefixPolicy struct {
	*CIM_ManagedElement

	//
	Label uint32

	//
	Precedence uint32

	//
	Prefix string

	//
	Store uint8
}

func NewMSFT_NetPrefixPolicyEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetPrefixPolicy, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetPrefixPolicy{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_NetPrefixPolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetPrefixPolicy, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetPrefixPolicy{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetLabel sets the value of Label for the instance
func (instance *MSFT_NetPrefixPolicy) SetPropertyLabel(value uint32) (err error) {
	return instance.SetProperty("Label", (value))
}

// GetLabel gets the value of Label for the instance
func (instance *MSFT_NetPrefixPolicy) GetPropertyLabel() (value uint32, err error) {
	retValue, err := instance.GetProperty("Label")
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

// SetPrecedence sets the value of Precedence for the instance
func (instance *MSFT_NetPrefixPolicy) SetPropertyPrecedence(value uint32) (err error) {
	return instance.SetProperty("Precedence", (value))
}

// GetPrecedence gets the value of Precedence for the instance
func (instance *MSFT_NetPrefixPolicy) GetPropertyPrecedence() (value uint32, err error) {
	retValue, err := instance.GetProperty("Precedence")
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

// SetPrefix sets the value of Prefix for the instance
func (instance *MSFT_NetPrefixPolicy) SetPropertyPrefix(value string) (err error) {
	return instance.SetProperty("Prefix", (value))
}

// GetPrefix gets the value of Prefix for the instance
func (instance *MSFT_NetPrefixPolicy) GetPropertyPrefix() (value string, err error) {
	retValue, err := instance.GetProperty("Prefix")
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

// SetStore sets the value of Store for the instance
func (instance *MSFT_NetPrefixPolicy) SetPropertyStore(value uint8) (err error) {
	return instance.SetProperty("Store", (value))
}

// GetStore gets the value of Store for the instance
func (instance *MSFT_NetPrefixPolicy) GetPropertyStore() (value uint8, err error) {
	retValue, err := instance.GetProperty("Store")
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

//

// <param name="Label" type="uint32 "></param>
// <param name="PassThru" type="bool "></param>
// <param name="PolicyStore" type="string "></param>
// <param name="Precedence" type="uint32 "></param>
// <param name="Prefix" type="string "></param>

// <param name="CmdletOutput" type="MSFT_NetPrefixPolicy []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetPrefixPolicy) Create( /* IN */ Prefix string,
	/* IN */ Precedence uint32,
	/* IN */ Label uint32,
	/* IN */ PolicyStore string,
	/* IN */ PassThru bool,
	/* OUT */ CmdletOutput []MSFT_NetPrefixPolicy) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Create", Prefix, Precedence, Label, PolicyStore, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
