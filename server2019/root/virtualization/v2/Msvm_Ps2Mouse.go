// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_Ps2Mouse struct
type Msvm_Ps2Mouse struct {
	*CIM_PointingDevice

	//
	AbsoluteCoordinates bool
}

func NewMsvm_Ps2MouseEx1(instance *cim.WmiInstance) (newInstance *Msvm_Ps2Mouse, err error) {
	tmp, err := NewCIM_PointingDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_Ps2Mouse{
		CIM_PointingDevice: tmp,
	}
	return
}

func NewMsvm_Ps2MouseEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_Ps2Mouse, err error) {
	tmp, err := NewCIM_PointingDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_Ps2Mouse{
		CIM_PointingDevice: tmp,
	}
	return
}

// SetAbsoluteCoordinates sets the value of AbsoluteCoordinates for the instance
func (instance *Msvm_Ps2Mouse) SetPropertyAbsoluteCoordinates(value bool) (err error) {
	return instance.SetProperty("AbsoluteCoordinates", (value))
}

// GetAbsoluteCoordinates gets the value of AbsoluteCoordinates for the instance
func (instance *Msvm_Ps2Mouse) GetPropertyAbsoluteCoordinates() (value bool, err error) {
	retValue, err := instance.GetProperty("AbsoluteCoordinates")
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

//

// <param name="ButtonIndex" type="uint32 "></param>

// <param name="ButtonState" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_Ps2Mouse) GetButtonState( /* IN */ ButtonIndex uint32,
	/* OUT */ ButtonState bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetButtonState", ButtonIndex)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ButtonIndex" type="uint32 "></param>
// <param name="ButtonState" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_Ps2Mouse) SetButtonState( /* IN */ ButtonIndex uint32,
	/* IN */ ButtonState bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetButtonState", ButtonIndex, ButtonState)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ButtonIndex" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_Ps2Mouse) ClickButton( /* IN */ ButtonIndex uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ClickButton", ButtonIndex)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="HorizontalDelta" type="int8 "></param>
// <param name="VerticalDelta" type="int8 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_Ps2Mouse) SetRelativePosition( /* IN */ HorizontalDelta int8,
	/* IN */ VerticalDelta int8) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetRelativePosition", HorizontalDelta, VerticalDelta)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ScrollPositionDelta" type="int8 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_Ps2Mouse) SetScrollPosition( /* IN */ ScrollPositionDelta int8) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetScrollPosition", ScrollPositionDelta)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
