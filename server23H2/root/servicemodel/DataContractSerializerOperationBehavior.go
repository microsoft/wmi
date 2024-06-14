// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DataContractSerializerOperationBehavior struct
type DataContractSerializerOperationBehavior struct {
	*Behavior

	// When enabled the IExtensibleDataObject interface on data contract types will be ignored.
	IgnoreExtensionDataObject bool

	// Limits the maximum number of objects that may be deserialized by the data contract serializer in a single deserialization episode.
	MaxItemsInObjectGraph int32

	// Defines the style of the SOAP message.
	Style string
}

func NewDataContractSerializerOperationBehaviorEx1(instance *cim.WmiInstance) (newInstance *DataContractSerializerOperationBehavior, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DataContractSerializerOperationBehavior{
		Behavior: tmp,
	}
	return
}

func NewDataContractSerializerOperationBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DataContractSerializerOperationBehavior, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DataContractSerializerOperationBehavior{
		Behavior: tmp,
	}
	return
}

// SetIgnoreExtensionDataObject sets the value of IgnoreExtensionDataObject for the instance
func (instance *DataContractSerializerOperationBehavior) SetPropertyIgnoreExtensionDataObject(value bool) (err error) {
	return instance.SetProperty("IgnoreExtensionDataObject", (value))
}

// GetIgnoreExtensionDataObject gets the value of IgnoreExtensionDataObject for the instance
func (instance *DataContractSerializerOperationBehavior) GetPropertyIgnoreExtensionDataObject() (value bool, err error) {
	retValue, err := instance.GetProperty("IgnoreExtensionDataObject")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetMaxItemsInObjectGraph sets the value of MaxItemsInObjectGraph for the instance
func (instance *DataContractSerializerOperationBehavior) SetPropertyMaxItemsInObjectGraph(value int32) (err error) {
	return instance.SetProperty("MaxItemsInObjectGraph", (value))
}

// GetMaxItemsInObjectGraph gets the value of MaxItemsInObjectGraph for the instance
func (instance *DataContractSerializerOperationBehavior) GetPropertyMaxItemsInObjectGraph() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxItemsInObjectGraph")
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

// SetStyle sets the value of Style for the instance
func (instance *DataContractSerializerOperationBehavior) SetPropertyStyle(value string) (err error) {
	return instance.SetProperty("Style", (value))
}

// GetStyle gets the value of Style for the instance
func (instance *DataContractSerializerOperationBehavior) GetPropertyStyle() (value string, err error) {
	retValue, err := instance.GetProperty("Style")
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
