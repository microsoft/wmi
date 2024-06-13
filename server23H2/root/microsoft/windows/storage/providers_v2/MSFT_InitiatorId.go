// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_InitiatorId struct
type MSFT_InitiatorId struct {
	*MSFT_StorageObject

	// This field specifies the operating system, version, driver, and other host environment factors that influence the behavior exposed by storage systems.
	HostType []InitiatorId_HostType

	// This field contains the address or unique identifier for the corresponding initiator port.
	InitiatorAddress string

	// When the corresponding array entry in HostType[] is "Other", this entry provides a string describing the manufacturer and OS/Environment. When the corresponding HostType[] entry is not "Other", this entry allows variations or qualifications of ClientTypes - for example, different versions of Solaris.
	OtherHostTypeDescription []string

	// This field specifies the type of the identifier used for initiator address.
	Type InitiatorId_Type
}

func NewMSFT_InitiatorIdEx1(instance *cim.WmiInstance) (newInstance *MSFT_InitiatorId, err error) {
	tmp, err := NewMSFT_StorageObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_InitiatorId{
		MSFT_StorageObject: tmp,
	}
	return
}

func NewMSFT_InitiatorIdEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_InitiatorId, err error) {
	tmp, err := NewMSFT_StorageObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_InitiatorId{
		MSFT_StorageObject: tmp,
	}
	return
}

// SetHostType sets the value of HostType for the instance
func (instance *MSFT_InitiatorId) SetPropertyHostType(value []InitiatorId_HostType) (err error) {
	return instance.SetProperty("HostType", (value))
}

// GetHostType gets the value of HostType for the instance
func (instance *MSFT_InitiatorId) GetPropertyHostType() (value []InitiatorId_HostType, err error) {
	retValue, err := instance.GetProperty("HostType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, InitiatorId_HostType(valuetmp))
	}

	return
}

// SetInitiatorAddress sets the value of InitiatorAddress for the instance
func (instance *MSFT_InitiatorId) SetPropertyInitiatorAddress(value string) (err error) {
	return instance.SetProperty("InitiatorAddress", (value))
}

// GetInitiatorAddress gets the value of InitiatorAddress for the instance
func (instance *MSFT_InitiatorId) GetPropertyInitiatorAddress() (value string, err error) {
	retValue, err := instance.GetProperty("InitiatorAddress")
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

// SetOtherHostTypeDescription sets the value of OtherHostTypeDescription for the instance
func (instance *MSFT_InitiatorId) SetPropertyOtherHostTypeDescription(value []string) (err error) {
	return instance.SetProperty("OtherHostTypeDescription", (value))
}

// GetOtherHostTypeDescription gets the value of OtherHostTypeDescription for the instance
func (instance *MSFT_InitiatorId) GetPropertyOtherHostTypeDescription() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherHostTypeDescription")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetType sets the value of Type for the instance
func (instance *MSFT_InitiatorId) SetPropertyType(value InitiatorId_Type) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSFT_InitiatorId) GetPropertyType() (value InitiatorId_Type, err error) {
	retValue, err := instance.GetProperty("Type")
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

	value = InitiatorId_Type(valuetmp)

	return
}

// Allows the user to delete an instance of an initiator id

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_InitiatorId) DeleteObject( /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DeleteObject")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
