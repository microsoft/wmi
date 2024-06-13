// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ContextRegistersARM64 struct
type ContextRegistersARM64 struct {
	*PerfInfo_V2

	//
	Cpsr uint32

	//
	Fp uint64

	//
	Lr uint64

	//
	Pc uint64

	//
	Sp uint64

	//
	X0 uint64

	//
	X1 uint64

	//
	X10 uint64

	//
	X11 uint64

	//
	X12 uint64

	//
	X13 uint64

	//
	X14 uint64

	//
	X15 uint64

	//
	X16 uint64

	//
	X17 uint64

	//
	X18 uint64

	//
	X19 uint64

	//
	X2 uint64

	//
	X20 uint64

	//
	X21 uint64

	//
	X22 uint64

	//
	X23 uint64

	//
	X24 uint64

	//
	X25 uint64

	//
	X26 uint64

	//
	X27 uint64

	//
	X28 uint64

	//
	X3 uint64

	//
	X4 uint64

	//
	X5 uint64

	//
	X6 uint64

	//
	X7 uint64

	//
	X8 uint64

	//
	X9 uint64
}

func NewContextRegistersARM64Ex1(instance *cim.WmiInstance) (newInstance *ContextRegistersARM64, err error) {
	tmp, err := NewPerfInfo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &ContextRegistersARM64{
		PerfInfo_V2: tmp,
	}
	return
}

func NewContextRegistersARM64Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ContextRegistersARM64, err error) {
	tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ContextRegistersARM64{
		PerfInfo_V2: tmp,
	}
	return
}

// SetCpsr sets the value of Cpsr for the instance
func (instance *ContextRegistersARM64) SetPropertyCpsr(value uint32) (err error) {
	return instance.SetProperty("Cpsr", (value))
}

// GetCpsr gets the value of Cpsr for the instance
func (instance *ContextRegistersARM64) GetPropertyCpsr() (value uint32, err error) {
	retValue, err := instance.GetProperty("Cpsr")
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

// SetFp sets the value of Fp for the instance
func (instance *ContextRegistersARM64) SetPropertyFp(value uint64) (err error) {
	return instance.SetProperty("Fp", (value))
}

// GetFp gets the value of Fp for the instance
func (instance *ContextRegistersARM64) GetPropertyFp() (value uint64, err error) {
	retValue, err := instance.GetProperty("Fp")
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

// SetLr sets the value of Lr for the instance
func (instance *ContextRegistersARM64) SetPropertyLr(value uint64) (err error) {
	return instance.SetProperty("Lr", (value))
}

// GetLr gets the value of Lr for the instance
func (instance *ContextRegistersARM64) GetPropertyLr() (value uint64, err error) {
	retValue, err := instance.GetProperty("Lr")
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

// SetPc sets the value of Pc for the instance
func (instance *ContextRegistersARM64) SetPropertyPc(value uint64) (err error) {
	return instance.SetProperty("Pc", (value))
}

// GetPc gets the value of Pc for the instance
func (instance *ContextRegistersARM64) GetPropertyPc() (value uint64, err error) {
	retValue, err := instance.GetProperty("Pc")
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

// SetSp sets the value of Sp for the instance
func (instance *ContextRegistersARM64) SetPropertySp(value uint64) (err error) {
	return instance.SetProperty("Sp", (value))
}

// GetSp gets the value of Sp for the instance
func (instance *ContextRegistersARM64) GetPropertySp() (value uint64, err error) {
	retValue, err := instance.GetProperty("Sp")
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

// SetX0 sets the value of X0 for the instance
func (instance *ContextRegistersARM64) SetPropertyX0(value uint64) (err error) {
	return instance.SetProperty("X0", (value))
}

// GetX0 gets the value of X0 for the instance
func (instance *ContextRegistersARM64) GetPropertyX0() (value uint64, err error) {
	retValue, err := instance.GetProperty("X0")
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

// SetX1 sets the value of X1 for the instance
func (instance *ContextRegistersARM64) SetPropertyX1(value uint64) (err error) {
	return instance.SetProperty("X1", (value))
}

// GetX1 gets the value of X1 for the instance
func (instance *ContextRegistersARM64) GetPropertyX1() (value uint64, err error) {
	retValue, err := instance.GetProperty("X1")
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

// SetX10 sets the value of X10 for the instance
func (instance *ContextRegistersARM64) SetPropertyX10(value uint64) (err error) {
	return instance.SetProperty("X10", (value))
}

// GetX10 gets the value of X10 for the instance
func (instance *ContextRegistersARM64) GetPropertyX10() (value uint64, err error) {
	retValue, err := instance.GetProperty("X10")
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

// SetX11 sets the value of X11 for the instance
func (instance *ContextRegistersARM64) SetPropertyX11(value uint64) (err error) {
	return instance.SetProperty("X11", (value))
}

// GetX11 gets the value of X11 for the instance
func (instance *ContextRegistersARM64) GetPropertyX11() (value uint64, err error) {
	retValue, err := instance.GetProperty("X11")
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

// SetX12 sets the value of X12 for the instance
func (instance *ContextRegistersARM64) SetPropertyX12(value uint64) (err error) {
	return instance.SetProperty("X12", (value))
}

// GetX12 gets the value of X12 for the instance
func (instance *ContextRegistersARM64) GetPropertyX12() (value uint64, err error) {
	retValue, err := instance.GetProperty("X12")
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

// SetX13 sets the value of X13 for the instance
func (instance *ContextRegistersARM64) SetPropertyX13(value uint64) (err error) {
	return instance.SetProperty("X13", (value))
}

// GetX13 gets the value of X13 for the instance
func (instance *ContextRegistersARM64) GetPropertyX13() (value uint64, err error) {
	retValue, err := instance.GetProperty("X13")
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

// SetX14 sets the value of X14 for the instance
func (instance *ContextRegistersARM64) SetPropertyX14(value uint64) (err error) {
	return instance.SetProperty("X14", (value))
}

// GetX14 gets the value of X14 for the instance
func (instance *ContextRegistersARM64) GetPropertyX14() (value uint64, err error) {
	retValue, err := instance.GetProperty("X14")
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

// SetX15 sets the value of X15 for the instance
func (instance *ContextRegistersARM64) SetPropertyX15(value uint64) (err error) {
	return instance.SetProperty("X15", (value))
}

// GetX15 gets the value of X15 for the instance
func (instance *ContextRegistersARM64) GetPropertyX15() (value uint64, err error) {
	retValue, err := instance.GetProperty("X15")
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

// SetX16 sets the value of X16 for the instance
func (instance *ContextRegistersARM64) SetPropertyX16(value uint64) (err error) {
	return instance.SetProperty("X16", (value))
}

// GetX16 gets the value of X16 for the instance
func (instance *ContextRegistersARM64) GetPropertyX16() (value uint64, err error) {
	retValue, err := instance.GetProperty("X16")
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

// SetX17 sets the value of X17 for the instance
func (instance *ContextRegistersARM64) SetPropertyX17(value uint64) (err error) {
	return instance.SetProperty("X17", (value))
}

// GetX17 gets the value of X17 for the instance
func (instance *ContextRegistersARM64) GetPropertyX17() (value uint64, err error) {
	retValue, err := instance.GetProperty("X17")
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

// SetX18 sets the value of X18 for the instance
func (instance *ContextRegistersARM64) SetPropertyX18(value uint64) (err error) {
	return instance.SetProperty("X18", (value))
}

// GetX18 gets the value of X18 for the instance
func (instance *ContextRegistersARM64) GetPropertyX18() (value uint64, err error) {
	retValue, err := instance.GetProperty("X18")
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

// SetX19 sets the value of X19 for the instance
func (instance *ContextRegistersARM64) SetPropertyX19(value uint64) (err error) {
	return instance.SetProperty("X19", (value))
}

// GetX19 gets the value of X19 for the instance
func (instance *ContextRegistersARM64) GetPropertyX19() (value uint64, err error) {
	retValue, err := instance.GetProperty("X19")
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

// SetX2 sets the value of X2 for the instance
func (instance *ContextRegistersARM64) SetPropertyX2(value uint64) (err error) {
	return instance.SetProperty("X2", (value))
}

// GetX2 gets the value of X2 for the instance
func (instance *ContextRegistersARM64) GetPropertyX2() (value uint64, err error) {
	retValue, err := instance.GetProperty("X2")
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

// SetX20 sets the value of X20 for the instance
func (instance *ContextRegistersARM64) SetPropertyX20(value uint64) (err error) {
	return instance.SetProperty("X20", (value))
}

// GetX20 gets the value of X20 for the instance
func (instance *ContextRegistersARM64) GetPropertyX20() (value uint64, err error) {
	retValue, err := instance.GetProperty("X20")
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

// SetX21 sets the value of X21 for the instance
func (instance *ContextRegistersARM64) SetPropertyX21(value uint64) (err error) {
	return instance.SetProperty("X21", (value))
}

// GetX21 gets the value of X21 for the instance
func (instance *ContextRegistersARM64) GetPropertyX21() (value uint64, err error) {
	retValue, err := instance.GetProperty("X21")
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

// SetX22 sets the value of X22 for the instance
func (instance *ContextRegistersARM64) SetPropertyX22(value uint64) (err error) {
	return instance.SetProperty("X22", (value))
}

// GetX22 gets the value of X22 for the instance
func (instance *ContextRegistersARM64) GetPropertyX22() (value uint64, err error) {
	retValue, err := instance.GetProperty("X22")
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

// SetX23 sets the value of X23 for the instance
func (instance *ContextRegistersARM64) SetPropertyX23(value uint64) (err error) {
	return instance.SetProperty("X23", (value))
}

// GetX23 gets the value of X23 for the instance
func (instance *ContextRegistersARM64) GetPropertyX23() (value uint64, err error) {
	retValue, err := instance.GetProperty("X23")
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

// SetX24 sets the value of X24 for the instance
func (instance *ContextRegistersARM64) SetPropertyX24(value uint64) (err error) {
	return instance.SetProperty("X24", (value))
}

// GetX24 gets the value of X24 for the instance
func (instance *ContextRegistersARM64) GetPropertyX24() (value uint64, err error) {
	retValue, err := instance.GetProperty("X24")
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

// SetX25 sets the value of X25 for the instance
func (instance *ContextRegistersARM64) SetPropertyX25(value uint64) (err error) {
	return instance.SetProperty("X25", (value))
}

// GetX25 gets the value of X25 for the instance
func (instance *ContextRegistersARM64) GetPropertyX25() (value uint64, err error) {
	retValue, err := instance.GetProperty("X25")
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

// SetX26 sets the value of X26 for the instance
func (instance *ContextRegistersARM64) SetPropertyX26(value uint64) (err error) {
	return instance.SetProperty("X26", (value))
}

// GetX26 gets the value of X26 for the instance
func (instance *ContextRegistersARM64) GetPropertyX26() (value uint64, err error) {
	retValue, err := instance.GetProperty("X26")
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

// SetX27 sets the value of X27 for the instance
func (instance *ContextRegistersARM64) SetPropertyX27(value uint64) (err error) {
	return instance.SetProperty("X27", (value))
}

// GetX27 gets the value of X27 for the instance
func (instance *ContextRegistersARM64) GetPropertyX27() (value uint64, err error) {
	retValue, err := instance.GetProperty("X27")
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

// SetX28 sets the value of X28 for the instance
func (instance *ContextRegistersARM64) SetPropertyX28(value uint64) (err error) {
	return instance.SetProperty("X28", (value))
}

// GetX28 gets the value of X28 for the instance
func (instance *ContextRegistersARM64) GetPropertyX28() (value uint64, err error) {
	retValue, err := instance.GetProperty("X28")
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

// SetX3 sets the value of X3 for the instance
func (instance *ContextRegistersARM64) SetPropertyX3(value uint64) (err error) {
	return instance.SetProperty("X3", (value))
}

// GetX3 gets the value of X3 for the instance
func (instance *ContextRegistersARM64) GetPropertyX3() (value uint64, err error) {
	retValue, err := instance.GetProperty("X3")
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

// SetX4 sets the value of X4 for the instance
func (instance *ContextRegistersARM64) SetPropertyX4(value uint64) (err error) {
	return instance.SetProperty("X4", (value))
}

// GetX4 gets the value of X4 for the instance
func (instance *ContextRegistersARM64) GetPropertyX4() (value uint64, err error) {
	retValue, err := instance.GetProperty("X4")
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

// SetX5 sets the value of X5 for the instance
func (instance *ContextRegistersARM64) SetPropertyX5(value uint64) (err error) {
	return instance.SetProperty("X5", (value))
}

// GetX5 gets the value of X5 for the instance
func (instance *ContextRegistersARM64) GetPropertyX5() (value uint64, err error) {
	retValue, err := instance.GetProperty("X5")
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

// SetX6 sets the value of X6 for the instance
func (instance *ContextRegistersARM64) SetPropertyX6(value uint64) (err error) {
	return instance.SetProperty("X6", (value))
}

// GetX6 gets the value of X6 for the instance
func (instance *ContextRegistersARM64) GetPropertyX6() (value uint64, err error) {
	retValue, err := instance.GetProperty("X6")
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

// SetX7 sets the value of X7 for the instance
func (instance *ContextRegistersARM64) SetPropertyX7(value uint64) (err error) {
	return instance.SetProperty("X7", (value))
}

// GetX7 gets the value of X7 for the instance
func (instance *ContextRegistersARM64) GetPropertyX7() (value uint64, err error) {
	retValue, err := instance.GetProperty("X7")
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

// SetX8 sets the value of X8 for the instance
func (instance *ContextRegistersARM64) SetPropertyX8(value uint64) (err error) {
	return instance.SetProperty("X8", (value))
}

// GetX8 gets the value of X8 for the instance
func (instance *ContextRegistersARM64) GetPropertyX8() (value uint64, err error) {
	retValue, err := instance.GetProperty("X8")
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

// SetX9 sets the value of X9 for the instance
func (instance *ContextRegistersARM64) SetPropertyX9(value uint64) (err error) {
	return instance.SetProperty("X9", (value))
}

// GetX9 gets the value of X9 for the instance
func (instance *ContextRegistersARM64) GetPropertyX9() (value uint64, err error) {
	retValue, err := instance.GetProperty("X9")
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
