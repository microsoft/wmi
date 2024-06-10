// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ContextRegistersAMD64 struct
type ContextRegistersAMD64 struct {
	*PerfInfo_V2

	//
	R10 uint64

	//
	R11 uint64

	//
	R12 uint64

	//
	R13 uint64

	//
	R14 uint64

	//
	R15 uint64

	//
	R8 uint64

	//
	R9 uint64

	//
	Rax uint64

	//
	Rbp uint64

	//
	Rbx uint64

	//
	Rcx uint64

	//
	Rdi uint64

	//
	Rdx uint64

	//
	Rip uint64

	//
	Rsi uint64

	//
	Rsp uint64
}

func NewContextRegistersAMD64Ex1(instance *cim.WmiInstance) (newInstance *ContextRegistersAMD64, err error) {
	tmp, err := NewPerfInfo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &ContextRegistersAMD64{
		PerfInfo_V2: tmp,
	}
	return
}

func NewContextRegistersAMD64Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ContextRegistersAMD64, err error) {
	tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ContextRegistersAMD64{
		PerfInfo_V2: tmp,
	}
	return
}

// SetR10 sets the value of R10 for the instance
func (instance *ContextRegistersAMD64) SetPropertyR10(value uint64) (err error) {
	return instance.SetProperty("R10", (value))
}

// GetR10 gets the value of R10 for the instance
func (instance *ContextRegistersAMD64) GetPropertyR10() (value uint64, err error) {
	retValue, err := instance.GetProperty("R10")
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

// SetR11 sets the value of R11 for the instance
func (instance *ContextRegistersAMD64) SetPropertyR11(value uint64) (err error) {
	return instance.SetProperty("R11", (value))
}

// GetR11 gets the value of R11 for the instance
func (instance *ContextRegistersAMD64) GetPropertyR11() (value uint64, err error) {
	retValue, err := instance.GetProperty("R11")
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

// SetR12 sets the value of R12 for the instance
func (instance *ContextRegistersAMD64) SetPropertyR12(value uint64) (err error) {
	return instance.SetProperty("R12", (value))
}

// GetR12 gets the value of R12 for the instance
func (instance *ContextRegistersAMD64) GetPropertyR12() (value uint64, err error) {
	retValue, err := instance.GetProperty("R12")
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

// SetR13 sets the value of R13 for the instance
func (instance *ContextRegistersAMD64) SetPropertyR13(value uint64) (err error) {
	return instance.SetProperty("R13", (value))
}

// GetR13 gets the value of R13 for the instance
func (instance *ContextRegistersAMD64) GetPropertyR13() (value uint64, err error) {
	retValue, err := instance.GetProperty("R13")
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

// SetR14 sets the value of R14 for the instance
func (instance *ContextRegistersAMD64) SetPropertyR14(value uint64) (err error) {
	return instance.SetProperty("R14", (value))
}

// GetR14 gets the value of R14 for the instance
func (instance *ContextRegistersAMD64) GetPropertyR14() (value uint64, err error) {
	retValue, err := instance.GetProperty("R14")
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

// SetR15 sets the value of R15 for the instance
func (instance *ContextRegistersAMD64) SetPropertyR15(value uint64) (err error) {
	return instance.SetProperty("R15", (value))
}

// GetR15 gets the value of R15 for the instance
func (instance *ContextRegistersAMD64) GetPropertyR15() (value uint64, err error) {
	retValue, err := instance.GetProperty("R15")
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

// SetR8 sets the value of R8 for the instance
func (instance *ContextRegistersAMD64) SetPropertyR8(value uint64) (err error) {
	return instance.SetProperty("R8", (value))
}

// GetR8 gets the value of R8 for the instance
func (instance *ContextRegistersAMD64) GetPropertyR8() (value uint64, err error) {
	retValue, err := instance.GetProperty("R8")
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

// SetR9 sets the value of R9 for the instance
func (instance *ContextRegistersAMD64) SetPropertyR9(value uint64) (err error) {
	return instance.SetProperty("R9", (value))
}

// GetR9 gets the value of R9 for the instance
func (instance *ContextRegistersAMD64) GetPropertyR9() (value uint64, err error) {
	retValue, err := instance.GetProperty("R9")
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

// SetRax sets the value of Rax for the instance
func (instance *ContextRegistersAMD64) SetPropertyRax(value uint64) (err error) {
	return instance.SetProperty("Rax", (value))
}

// GetRax gets the value of Rax for the instance
func (instance *ContextRegistersAMD64) GetPropertyRax() (value uint64, err error) {
	retValue, err := instance.GetProperty("Rax")
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

// SetRbp sets the value of Rbp for the instance
func (instance *ContextRegistersAMD64) SetPropertyRbp(value uint64) (err error) {
	return instance.SetProperty("Rbp", (value))
}

// GetRbp gets the value of Rbp for the instance
func (instance *ContextRegistersAMD64) GetPropertyRbp() (value uint64, err error) {
	retValue, err := instance.GetProperty("Rbp")
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

// SetRbx sets the value of Rbx for the instance
func (instance *ContextRegistersAMD64) SetPropertyRbx(value uint64) (err error) {
	return instance.SetProperty("Rbx", (value))
}

// GetRbx gets the value of Rbx for the instance
func (instance *ContextRegistersAMD64) GetPropertyRbx() (value uint64, err error) {
	retValue, err := instance.GetProperty("Rbx")
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

// SetRcx sets the value of Rcx for the instance
func (instance *ContextRegistersAMD64) SetPropertyRcx(value uint64) (err error) {
	return instance.SetProperty("Rcx", (value))
}

// GetRcx gets the value of Rcx for the instance
func (instance *ContextRegistersAMD64) GetPropertyRcx() (value uint64, err error) {
	retValue, err := instance.GetProperty("Rcx")
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

// SetRdi sets the value of Rdi for the instance
func (instance *ContextRegistersAMD64) SetPropertyRdi(value uint64) (err error) {
	return instance.SetProperty("Rdi", (value))
}

// GetRdi gets the value of Rdi for the instance
func (instance *ContextRegistersAMD64) GetPropertyRdi() (value uint64, err error) {
	retValue, err := instance.GetProperty("Rdi")
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

// SetRdx sets the value of Rdx for the instance
func (instance *ContextRegistersAMD64) SetPropertyRdx(value uint64) (err error) {
	return instance.SetProperty("Rdx", (value))
}

// GetRdx gets the value of Rdx for the instance
func (instance *ContextRegistersAMD64) GetPropertyRdx() (value uint64, err error) {
	retValue, err := instance.GetProperty("Rdx")
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

// SetRip sets the value of Rip for the instance
func (instance *ContextRegistersAMD64) SetPropertyRip(value uint64) (err error) {
	return instance.SetProperty("Rip", (value))
}

// GetRip gets the value of Rip for the instance
func (instance *ContextRegistersAMD64) GetPropertyRip() (value uint64, err error) {
	retValue, err := instance.GetProperty("Rip")
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

// SetRsi sets the value of Rsi for the instance
func (instance *ContextRegistersAMD64) SetPropertyRsi(value uint64) (err error) {
	return instance.SetProperty("Rsi", (value))
}

// GetRsi gets the value of Rsi for the instance
func (instance *ContextRegistersAMD64) GetPropertyRsi() (value uint64, err error) {
	retValue, err := instance.GetProperty("Rsi")
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

// SetRsp sets the value of Rsp for the instance
func (instance *ContextRegistersAMD64) SetPropertyRsp(value uint64) (err error) {
	return instance.SetProperty("Rsp", (value))
}

// GetRsp gets the value of Rsp for the instance
func (instance *ContextRegistersAMD64) GetPropertyRsp() (value uint64, err error) {
	retValue, err := instance.GetProperty("Rsp")
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
