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

// BcdObject struct
type BcdObject struct {
	*cim.WmiInstance

	// This is the guid id of this object, unique to this store.
	Id string

	// This is the file path of the store that this object belongs to.
	StoreFilePath string

	// The upper 4 bits (28-31) represent the object type. The meaning of the lower 28 bits (0-27) is dependent on the object type.
	Type uint32
}

func NewBcdObjectEx1(instance *cim.WmiInstance) (newInstance *BcdObject, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &BcdObject{
		WmiInstance: tmp,
	}
	return
}

func NewBcdObjectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BcdObject, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BcdObject{
		WmiInstance: tmp,
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *BcdObject) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *BcdObject) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetStoreFilePath sets the value of StoreFilePath for the instance
func (instance *BcdObject) SetPropertyStoreFilePath(value string) (err error) {
	return instance.SetProperty("StoreFilePath", (value))
}

// GetStoreFilePath gets the value of StoreFilePath for the instance
func (instance *BcdObject) GetPropertyStoreFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("StoreFilePath")
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

// SetType sets the value of Type for the instance
func (instance *BcdObject) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *BcdObject) GetPropertyType() (value uint32, err error) {
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

//

// <param name="ReturnValue" type="bool "></param>
// <param name="Types" type="uint32 []"></param>
func (instance *BcdObject) EnumerateElementTypes( /* OUT */ Types []uint32) (result bool, err error) {
	retVal, err := instance.InvokeMethod("EnumerateElementTypes")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = bool(retValue != 0)
	return

}

//

// <param name="Elements" type="BcdElement []"></param>
// <param name="ReturnValue" type="bool "></param>
func (instance *BcdObject) EnumerateElements( /* OUT */ Elements []BcdElement) (result bool, err error) {
	retVal, err := instance.InvokeMethod("EnumerateElements")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = bool(retValue != 0)
	return

}

//

// <param name="Type" type="uint32 "></param>

// <param name="Element" type="BcdElement "></param>
// <param name="ReturnValue" type="bool "></param>
func (instance *BcdObject) GetElement( /* IN */ Type uint32,
	/* OUT */ Element BcdElement) (result bool, err error) {
	retVal, err := instance.InvokeMethod("GetElement", Type)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = bool(retValue != 0)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Type" type="uint32 "></param>

// <param name="Element" type="BcdElement "></param>
// <param name="ReturnValue" type="bool "></param>
func (instance *BcdObject) GetElementWithFlags( /* IN */ Type uint32,
	/* IN */ Flags uint32,
	/* OUT */ Element BcdElement) (result bool, err error) {
	retVal, err := instance.InvokeMethod("GetElementWithFlags", Type, Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = bool(retValue != 0)
	return

}

//

// <param name="AdditionalOptions" type="string "></param>
// <param name="DeviceType" type="BcdObject_DeviceType "></param>
// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdObject) SetDeviceElement( /* IN */ Type uint32,
	/* IN */ DeviceType BcdObject_DeviceType,
	/* IN */ AdditionalOptions string) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetDeviceElement", Type, DeviceType, AdditionalOptions)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="AdditionalOptions" type="string "></param>
// <param name="DeviceType" type="BcdObject_DeviceType "></param>
// <param name="Path" type="string "></param>
// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdObject) SetPartitionDeviceElement( /* IN */ Type uint32,
	/* IN */ DeviceType BcdObject_DeviceType,
	/* IN */ AdditionalOptions string,
	/* IN */ Path string) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetPartitionDeviceElement", Type, DeviceType, AdditionalOptions, Path)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="AdditionalOptions" type="string "></param>
// <param name="DeviceType" type="BcdObject_DeviceType "></param>
// <param name="Flags" type="uint32 "></param>
// <param name="Path" type="string "></param>
// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdObject) SetPartitionDeviceElementWithFlags( /* IN */ Type uint32,
	/* IN */ DeviceType BcdObject_DeviceType,
	/* IN */ AdditionalOptions string,
	/* IN */ Path string,
	/* IN */ Flags uint32) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetPartitionDeviceElementWithFlags", Type, DeviceType, AdditionalOptions, Path, Flags)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="AdditionalOptions" type="string "></param>
// <param name="DeviceType" type="BcdObject_DeviceType "></param>
// <param name="ParentAdditionalOptions" type="string "></param>
// <param name="ParentDeviceType" type="uint32 "></param>
// <param name="ParentPath" type="string "></param>
// <param name="Path" type="string "></param>
// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdObject) SetFileDeviceElement( /* IN */ Type uint32,
	/* IN */ DeviceType BcdObject_DeviceType,
	/* IN */ AdditionalOptions string,
	/* IN */ Path string,
	/* IN */ ParentDeviceType uint32,
	/* IN */ ParentAdditionalOptions string,
	/* IN */ ParentPath string) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetFileDeviceElement", Type, DeviceType, AdditionalOptions, Path, ParentDeviceType, ParentAdditionalOptions, ParentPath)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="DiskSignature" type="string "></param>
// <param name="PartitionIdentifier" type="string "></param>
// <param name="PartitionStyle" type="BcdObject_PartitionStyle "></param>
// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdObject) SetQualifiedPartitionDeviceElement( /* IN */ Type uint32,
	/* IN */ PartitionStyle BcdObject_PartitionStyle,
	/* IN */ DiskSignature string,
	/* IN */ PartitionIdentifier string) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetQualifiedPartitionDeviceElement", Type, PartitionStyle, DiskSignature, PartitionIdentifier)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="CustomLocate" type="uint32 "></param>
// <param name="ParentAdditionalOptions" type="string "></param>
// <param name="ParentDeviceType" type="uint32 "></param>
// <param name="ParentPath" type="string "></param>
// <param name="Path" type="string "></param>
// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdObject) SetVhdDeviceElement( /* IN */ Type uint32,
	/* IN */ Path string,
	/* IN */ ParentDeviceType uint32,
	/* IN */ ParentAdditionalOptions string,
	/* IN */ ParentPath string,
	/* IN */ CustomLocate uint32) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetVhdDeviceElement", Type, Path, ParentDeviceType, ParentAdditionalOptions, ParentPath, CustomLocate)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="String" type="string "></param>
// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdObject) SetStringElement( /* IN */ Type uint32,
	/* IN */ String string) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetStringElement", Type, String)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="Id" type="string "></param>
// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdObject) SetObjectElement( /* IN */ Type uint32,
	/* IN */ Id string) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetObjectElement", Type, Id)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="Ids" type="string []"></param>
// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdObject) SetObjectListElement( /* IN */ Type uint32,
	/* IN */ Ids []string) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetObjectListElement", Type, Ids)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="Integer" type="uint64 "></param>
// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdObject) SetIntegerElement( /* IN */ Type uint32,
	/* IN */ Integer uint64) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetIntegerElement", Type, Integer)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="Integers" type="uint64 []"></param>
// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdObject) SetIntegerListElement( /* IN */ Type uint32,
	/* IN */ Integers []uint64) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetIntegerListElement", Type, Integers)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="Boolean" type="bool "></param>
// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdObject) SetBooleanElement( /* IN */ Type uint32,
	/* IN */ Boolean bool) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetBooleanElement", Type, Boolean)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdObject) DeleteElement( /* IN */ Type uint32) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DeleteElement", Type)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}
