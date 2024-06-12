// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// StackWalk_Event struct
type StackWalk_Event struct {
	*StackWalk

	//
	EventTimeStamp uint64

	//
	Stack1 uint32

	//
	Stack10 uint32

	//
	Stack11 uint32

	//
	Stack12 uint32

	//
	Stack13 uint32

	//
	Stack14 uint32

	//
	Stack15 uint32

	//
	Stack16 uint32

	//
	Stack17 uint32

	//
	Stack18 uint32

	//
	Stack19 uint32

	//
	Stack2 uint32

	//
	Stack20 uint32

	//
	Stack21 uint32

	//
	Stack22 uint32

	//
	Stack23 uint32

	//
	Stack24 uint32

	//
	Stack25 uint32

	//
	Stack26 uint32

	//
	Stack27 uint32

	//
	Stack28 uint32

	//
	Stack29 uint32

	//
	Stack3 uint32

	//
	Stack30 uint32

	//
	Stack31 uint32

	//
	Stack32 uint32

	//
	Stack4 uint32

	//
	Stack5 uint32

	//
	Stack6 uint32

	//
	Stack7 uint32

	//
	Stack8 uint32

	//
	Stack9 uint32

	//
	StackProcess uint32

	//
	StackThread uint32
}

func NewStackWalk_EventEx1(instance *cim.WmiInstance) (newInstance *StackWalk_Event, err error) {
	tmp, err := NewStackWalkEx1(instance)

	if err != nil {
		return
	}
	newInstance = &StackWalk_Event{
		StackWalk: tmp,
	}
	return
}

func NewStackWalk_EventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *StackWalk_Event, err error) {
	tmp, err := NewStackWalkEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &StackWalk_Event{
		StackWalk: tmp,
	}
	return
}

// SetEventTimeStamp sets the value of EventTimeStamp for the instance
func (instance *StackWalk_Event) SetPropertyEventTimeStamp(value uint64) (err error) {
	return instance.SetProperty("EventTimeStamp", (value))
}

// GetEventTimeStamp gets the value of EventTimeStamp for the instance
func (instance *StackWalk_Event) GetPropertyEventTimeStamp() (value uint64, err error) {
	retValue, err := instance.GetProperty("EventTimeStamp")
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

// SetStack1 sets the value of Stack1 for the instance
func (instance *StackWalk_Event) SetPropertyStack1(value uint32) (err error) {
	return instance.SetProperty("Stack1", (value))
}

// GetStack1 gets the value of Stack1 for the instance
func (instance *StackWalk_Event) GetPropertyStack1() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack1")
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

// SetStack10 sets the value of Stack10 for the instance
func (instance *StackWalk_Event) SetPropertyStack10(value uint32) (err error) {
	return instance.SetProperty("Stack10", (value))
}

// GetStack10 gets the value of Stack10 for the instance
func (instance *StackWalk_Event) GetPropertyStack10() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack10")
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

// SetStack11 sets the value of Stack11 for the instance
func (instance *StackWalk_Event) SetPropertyStack11(value uint32) (err error) {
	return instance.SetProperty("Stack11", (value))
}

// GetStack11 gets the value of Stack11 for the instance
func (instance *StackWalk_Event) GetPropertyStack11() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack11")
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

// SetStack12 sets the value of Stack12 for the instance
func (instance *StackWalk_Event) SetPropertyStack12(value uint32) (err error) {
	return instance.SetProperty("Stack12", (value))
}

// GetStack12 gets the value of Stack12 for the instance
func (instance *StackWalk_Event) GetPropertyStack12() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack12")
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

// SetStack13 sets the value of Stack13 for the instance
func (instance *StackWalk_Event) SetPropertyStack13(value uint32) (err error) {
	return instance.SetProperty("Stack13", (value))
}

// GetStack13 gets the value of Stack13 for the instance
func (instance *StackWalk_Event) GetPropertyStack13() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack13")
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

// SetStack14 sets the value of Stack14 for the instance
func (instance *StackWalk_Event) SetPropertyStack14(value uint32) (err error) {
	return instance.SetProperty("Stack14", (value))
}

// GetStack14 gets the value of Stack14 for the instance
func (instance *StackWalk_Event) GetPropertyStack14() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack14")
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

// SetStack15 sets the value of Stack15 for the instance
func (instance *StackWalk_Event) SetPropertyStack15(value uint32) (err error) {
	return instance.SetProperty("Stack15", (value))
}

// GetStack15 gets the value of Stack15 for the instance
func (instance *StackWalk_Event) GetPropertyStack15() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack15")
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

// SetStack16 sets the value of Stack16 for the instance
func (instance *StackWalk_Event) SetPropertyStack16(value uint32) (err error) {
	return instance.SetProperty("Stack16", (value))
}

// GetStack16 gets the value of Stack16 for the instance
func (instance *StackWalk_Event) GetPropertyStack16() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack16")
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

// SetStack17 sets the value of Stack17 for the instance
func (instance *StackWalk_Event) SetPropertyStack17(value uint32) (err error) {
	return instance.SetProperty("Stack17", (value))
}

// GetStack17 gets the value of Stack17 for the instance
func (instance *StackWalk_Event) GetPropertyStack17() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack17")
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

// SetStack18 sets the value of Stack18 for the instance
func (instance *StackWalk_Event) SetPropertyStack18(value uint32) (err error) {
	return instance.SetProperty("Stack18", (value))
}

// GetStack18 gets the value of Stack18 for the instance
func (instance *StackWalk_Event) GetPropertyStack18() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack18")
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

// SetStack19 sets the value of Stack19 for the instance
func (instance *StackWalk_Event) SetPropertyStack19(value uint32) (err error) {
	return instance.SetProperty("Stack19", (value))
}

// GetStack19 gets the value of Stack19 for the instance
func (instance *StackWalk_Event) GetPropertyStack19() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack19")
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

// SetStack2 sets the value of Stack2 for the instance
func (instance *StackWalk_Event) SetPropertyStack2(value uint32) (err error) {
	return instance.SetProperty("Stack2", (value))
}

// GetStack2 gets the value of Stack2 for the instance
func (instance *StackWalk_Event) GetPropertyStack2() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack2")
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

// SetStack20 sets the value of Stack20 for the instance
func (instance *StackWalk_Event) SetPropertyStack20(value uint32) (err error) {
	return instance.SetProperty("Stack20", (value))
}

// GetStack20 gets the value of Stack20 for the instance
func (instance *StackWalk_Event) GetPropertyStack20() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack20")
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

// SetStack21 sets the value of Stack21 for the instance
func (instance *StackWalk_Event) SetPropertyStack21(value uint32) (err error) {
	return instance.SetProperty("Stack21", (value))
}

// GetStack21 gets the value of Stack21 for the instance
func (instance *StackWalk_Event) GetPropertyStack21() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack21")
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

// SetStack22 sets the value of Stack22 for the instance
func (instance *StackWalk_Event) SetPropertyStack22(value uint32) (err error) {
	return instance.SetProperty("Stack22", (value))
}

// GetStack22 gets the value of Stack22 for the instance
func (instance *StackWalk_Event) GetPropertyStack22() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack22")
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

// SetStack23 sets the value of Stack23 for the instance
func (instance *StackWalk_Event) SetPropertyStack23(value uint32) (err error) {
	return instance.SetProperty("Stack23", (value))
}

// GetStack23 gets the value of Stack23 for the instance
func (instance *StackWalk_Event) GetPropertyStack23() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack23")
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

// SetStack24 sets the value of Stack24 for the instance
func (instance *StackWalk_Event) SetPropertyStack24(value uint32) (err error) {
	return instance.SetProperty("Stack24", (value))
}

// GetStack24 gets the value of Stack24 for the instance
func (instance *StackWalk_Event) GetPropertyStack24() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack24")
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

// SetStack25 sets the value of Stack25 for the instance
func (instance *StackWalk_Event) SetPropertyStack25(value uint32) (err error) {
	return instance.SetProperty("Stack25", (value))
}

// GetStack25 gets the value of Stack25 for the instance
func (instance *StackWalk_Event) GetPropertyStack25() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack25")
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

// SetStack26 sets the value of Stack26 for the instance
func (instance *StackWalk_Event) SetPropertyStack26(value uint32) (err error) {
	return instance.SetProperty("Stack26", (value))
}

// GetStack26 gets the value of Stack26 for the instance
func (instance *StackWalk_Event) GetPropertyStack26() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack26")
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

// SetStack27 sets the value of Stack27 for the instance
func (instance *StackWalk_Event) SetPropertyStack27(value uint32) (err error) {
	return instance.SetProperty("Stack27", (value))
}

// GetStack27 gets the value of Stack27 for the instance
func (instance *StackWalk_Event) GetPropertyStack27() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack27")
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

// SetStack28 sets the value of Stack28 for the instance
func (instance *StackWalk_Event) SetPropertyStack28(value uint32) (err error) {
	return instance.SetProperty("Stack28", (value))
}

// GetStack28 gets the value of Stack28 for the instance
func (instance *StackWalk_Event) GetPropertyStack28() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack28")
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

// SetStack29 sets the value of Stack29 for the instance
func (instance *StackWalk_Event) SetPropertyStack29(value uint32) (err error) {
	return instance.SetProperty("Stack29", (value))
}

// GetStack29 gets the value of Stack29 for the instance
func (instance *StackWalk_Event) GetPropertyStack29() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack29")
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

// SetStack3 sets the value of Stack3 for the instance
func (instance *StackWalk_Event) SetPropertyStack3(value uint32) (err error) {
	return instance.SetProperty("Stack3", (value))
}

// GetStack3 gets the value of Stack3 for the instance
func (instance *StackWalk_Event) GetPropertyStack3() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack3")
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

// SetStack30 sets the value of Stack30 for the instance
func (instance *StackWalk_Event) SetPropertyStack30(value uint32) (err error) {
	return instance.SetProperty("Stack30", (value))
}

// GetStack30 gets the value of Stack30 for the instance
func (instance *StackWalk_Event) GetPropertyStack30() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack30")
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

// SetStack31 sets the value of Stack31 for the instance
func (instance *StackWalk_Event) SetPropertyStack31(value uint32) (err error) {
	return instance.SetProperty("Stack31", (value))
}

// GetStack31 gets the value of Stack31 for the instance
func (instance *StackWalk_Event) GetPropertyStack31() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack31")
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

// SetStack32 sets the value of Stack32 for the instance
func (instance *StackWalk_Event) SetPropertyStack32(value uint32) (err error) {
	return instance.SetProperty("Stack32", (value))
}

// GetStack32 gets the value of Stack32 for the instance
func (instance *StackWalk_Event) GetPropertyStack32() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack32")
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

// SetStack4 sets the value of Stack4 for the instance
func (instance *StackWalk_Event) SetPropertyStack4(value uint32) (err error) {
	return instance.SetProperty("Stack4", (value))
}

// GetStack4 gets the value of Stack4 for the instance
func (instance *StackWalk_Event) GetPropertyStack4() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack4")
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

// SetStack5 sets the value of Stack5 for the instance
func (instance *StackWalk_Event) SetPropertyStack5(value uint32) (err error) {
	return instance.SetProperty("Stack5", (value))
}

// GetStack5 gets the value of Stack5 for the instance
func (instance *StackWalk_Event) GetPropertyStack5() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack5")
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

// SetStack6 sets the value of Stack6 for the instance
func (instance *StackWalk_Event) SetPropertyStack6(value uint32) (err error) {
	return instance.SetProperty("Stack6", (value))
}

// GetStack6 gets the value of Stack6 for the instance
func (instance *StackWalk_Event) GetPropertyStack6() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack6")
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

// SetStack7 sets the value of Stack7 for the instance
func (instance *StackWalk_Event) SetPropertyStack7(value uint32) (err error) {
	return instance.SetProperty("Stack7", (value))
}

// GetStack7 gets the value of Stack7 for the instance
func (instance *StackWalk_Event) GetPropertyStack7() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack7")
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

// SetStack8 sets the value of Stack8 for the instance
func (instance *StackWalk_Event) SetPropertyStack8(value uint32) (err error) {
	return instance.SetProperty("Stack8", (value))
}

// GetStack8 gets the value of Stack8 for the instance
func (instance *StackWalk_Event) GetPropertyStack8() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack8")
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

// SetStack9 sets the value of Stack9 for the instance
func (instance *StackWalk_Event) SetPropertyStack9(value uint32) (err error) {
	return instance.SetProperty("Stack9", (value))
}

// GetStack9 gets the value of Stack9 for the instance
func (instance *StackWalk_Event) GetPropertyStack9() (value uint32, err error) {
	retValue, err := instance.GetProperty("Stack9")
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

// SetStackProcess sets the value of StackProcess for the instance
func (instance *StackWalk_Event) SetPropertyStackProcess(value uint32) (err error) {
	return instance.SetProperty("StackProcess", (value))
}

// GetStackProcess gets the value of StackProcess for the instance
func (instance *StackWalk_Event) GetPropertyStackProcess() (value uint32, err error) {
	retValue, err := instance.GetProperty("StackProcess")
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

// SetStackThread sets the value of StackThread for the instance
func (instance *StackWalk_Event) SetPropertyStackThread(value uint32) (err error) {
	return instance.SetProperty("StackThread", (value))
}

// GetStackThread gets the value of StackThread for the instance
func (instance *StackWalk_Event) GetPropertyStackThread() (value uint32, err error) {
	retValue, err := instance.GetProperty("StackThread")
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
