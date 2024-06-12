// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_SyntheticMouse struct
type Msvm_SyntheticMouse struct {
	*CIM_PointingDevice

	//
	AbsoluteCoordinates bool

	//
	HorizontalPosition int32

	//
	ScrollPosition int32

	//
	VerticalPosition int32
}

func NewMsvm_SyntheticMouseEx1(instance *cim.WmiInstance) (newInstance *Msvm_SyntheticMouse, err error) {
	tmp, err := NewCIM_PointingDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SyntheticMouse{
		CIM_PointingDevice: tmp,
	}
	return
}

func NewMsvm_SyntheticMouseEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SyntheticMouse, err error) {
	tmp, err := NewCIM_PointingDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SyntheticMouse{
		CIM_PointingDevice: tmp,
	}
	return
}

// SetAbsoluteCoordinates sets the value of AbsoluteCoordinates for the instance
func (instance *Msvm_SyntheticMouse) SetPropertyAbsoluteCoordinates(value bool) (err error) {
	return instance.SetProperty("AbsoluteCoordinates", (value))
}

// GetAbsoluteCoordinates gets the value of AbsoluteCoordinates for the instance
func (instance *Msvm_SyntheticMouse) GetPropertyAbsoluteCoordinates() (value bool, err error) {
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

// SetHorizontalPosition sets the value of HorizontalPosition for the instance
func (instance *Msvm_SyntheticMouse) SetPropertyHorizontalPosition(value int32) (err error) {
	return instance.SetProperty("HorizontalPosition", (value))
}

// GetHorizontalPosition gets the value of HorizontalPosition for the instance
func (instance *Msvm_SyntheticMouse) GetPropertyHorizontalPosition() (value int32, err error) {
	retValue, err := instance.GetProperty("HorizontalPosition")
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

// SetScrollPosition sets the value of ScrollPosition for the instance
func (instance *Msvm_SyntheticMouse) SetPropertyScrollPosition(value int32) (err error) {
	return instance.SetProperty("ScrollPosition", (value))
}

// GetScrollPosition gets the value of ScrollPosition for the instance
func (instance *Msvm_SyntheticMouse) GetPropertyScrollPosition() (value int32, err error) {
	retValue, err := instance.GetProperty("ScrollPosition")
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

// SetVerticalPosition sets the value of VerticalPosition for the instance
func (instance *Msvm_SyntheticMouse) SetPropertyVerticalPosition(value int32) (err error) {
	return instance.SetProperty("VerticalPosition", (value))
}

// GetVerticalPosition gets the value of VerticalPosition for the instance
func (instance *Msvm_SyntheticMouse) GetPropertyVerticalPosition() (value int32, err error) {
	retValue, err := instance.GetProperty("VerticalPosition")
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

//

// <param name="ButtonIndex" type="uint32 "></param>

// <param name="IsDown" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_SyntheticMouse) GetButtonState( /* IN */ ButtonIndex uint32,
	/* OUT */ IsDown bool) (result uint32, err error) {
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
// <param name="IsDown" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_SyntheticMouse) SetButtonState( /* IN */ ButtonIndex uint32,
	/* IN */ IsDown bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetButtonState", ButtonIndex, IsDown)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ButtonIndex" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_SyntheticMouse) ClickButton( /* IN */ ButtonIndex uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ClickButton", ButtonIndex)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="HorizontalPosition" type="int32 "></param>
// <param name="VerticalPosition" type="int32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_SyntheticMouse) SetAbsolutePosition( /* IN */ HorizontalPosition int32,
	/* IN */ VerticalPosition int32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetAbsolutePosition", HorizontalPosition, VerticalPosition)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ScrollPositionDelta" type="int32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_SyntheticMouse) SetScrollPosition( /* IN */ ScrollPositionDelta int32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetScrollPosition", ScrollPositionDelta)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
