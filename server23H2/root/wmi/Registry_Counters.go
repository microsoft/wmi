// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Registry_Counters struct
type Registry_Counters struct {
	*Registry

	//
	Counter1 uint64

	//
	Counter10 uint64

	//
	Counter11 uint64

	//
	Counter2 uint64

	//
	Counter3 uint64

	//
	Counter4 uint64

	//
	Counter5 uint64

	//
	Counter6 uint64

	//
	Counter7 uint64

	//
	Counter8 uint64

	//
	Counter9 uint64
}

func NewRegistry_CountersEx1(instance *cim.WmiInstance) (newInstance *Registry_Counters, err error) {
	tmp, err := NewRegistryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Registry_Counters{
		Registry: tmp,
	}
	return
}

func NewRegistry_CountersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Registry_Counters, err error) {
	tmp, err := NewRegistryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Registry_Counters{
		Registry: tmp,
	}
	return
}

// SetCounter1 sets the value of Counter1 for the instance
func (instance *Registry_Counters) SetPropertyCounter1(value uint64) (err error) {
	return instance.SetProperty("Counter1", (value))
}

// GetCounter1 gets the value of Counter1 for the instance
func (instance *Registry_Counters) GetPropertyCounter1() (value uint64, err error) {
	retValue, err := instance.GetProperty("Counter1")
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

// SetCounter10 sets the value of Counter10 for the instance
func (instance *Registry_Counters) SetPropertyCounter10(value uint64) (err error) {
	return instance.SetProperty("Counter10", (value))
}

// GetCounter10 gets the value of Counter10 for the instance
func (instance *Registry_Counters) GetPropertyCounter10() (value uint64, err error) {
	retValue, err := instance.GetProperty("Counter10")
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

// SetCounter11 sets the value of Counter11 for the instance
func (instance *Registry_Counters) SetPropertyCounter11(value uint64) (err error) {
	return instance.SetProperty("Counter11", (value))
}

// GetCounter11 gets the value of Counter11 for the instance
func (instance *Registry_Counters) GetPropertyCounter11() (value uint64, err error) {
	retValue, err := instance.GetProperty("Counter11")
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

// SetCounter2 sets the value of Counter2 for the instance
func (instance *Registry_Counters) SetPropertyCounter2(value uint64) (err error) {
	return instance.SetProperty("Counter2", (value))
}

// GetCounter2 gets the value of Counter2 for the instance
func (instance *Registry_Counters) GetPropertyCounter2() (value uint64, err error) {
	retValue, err := instance.GetProperty("Counter2")
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

// SetCounter3 sets the value of Counter3 for the instance
func (instance *Registry_Counters) SetPropertyCounter3(value uint64) (err error) {
	return instance.SetProperty("Counter3", (value))
}

// GetCounter3 gets the value of Counter3 for the instance
func (instance *Registry_Counters) GetPropertyCounter3() (value uint64, err error) {
	retValue, err := instance.GetProperty("Counter3")
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

// SetCounter4 sets the value of Counter4 for the instance
func (instance *Registry_Counters) SetPropertyCounter4(value uint64) (err error) {
	return instance.SetProperty("Counter4", (value))
}

// GetCounter4 gets the value of Counter4 for the instance
func (instance *Registry_Counters) GetPropertyCounter4() (value uint64, err error) {
	retValue, err := instance.GetProperty("Counter4")
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

// SetCounter5 sets the value of Counter5 for the instance
func (instance *Registry_Counters) SetPropertyCounter5(value uint64) (err error) {
	return instance.SetProperty("Counter5", (value))
}

// GetCounter5 gets the value of Counter5 for the instance
func (instance *Registry_Counters) GetPropertyCounter5() (value uint64, err error) {
	retValue, err := instance.GetProperty("Counter5")
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

// SetCounter6 sets the value of Counter6 for the instance
func (instance *Registry_Counters) SetPropertyCounter6(value uint64) (err error) {
	return instance.SetProperty("Counter6", (value))
}

// GetCounter6 gets the value of Counter6 for the instance
func (instance *Registry_Counters) GetPropertyCounter6() (value uint64, err error) {
	retValue, err := instance.GetProperty("Counter6")
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

// SetCounter7 sets the value of Counter7 for the instance
func (instance *Registry_Counters) SetPropertyCounter7(value uint64) (err error) {
	return instance.SetProperty("Counter7", (value))
}

// GetCounter7 gets the value of Counter7 for the instance
func (instance *Registry_Counters) GetPropertyCounter7() (value uint64, err error) {
	retValue, err := instance.GetProperty("Counter7")
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

// SetCounter8 sets the value of Counter8 for the instance
func (instance *Registry_Counters) SetPropertyCounter8(value uint64) (err error) {
	return instance.SetProperty("Counter8", (value))
}

// GetCounter8 gets the value of Counter8 for the instance
func (instance *Registry_Counters) GetPropertyCounter8() (value uint64, err error) {
	retValue, err := instance.GetProperty("Counter8")
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

// SetCounter9 sets the value of Counter9 for the instance
func (instance *Registry_Counters) SetPropertyCounter9(value uint64) (err error) {
	return instance.SetProperty("Counter9", (value))
}

// GetCounter9 gets the value of Counter9 for the instance
func (instance *Registry_Counters) GetPropertyCounter9() (value uint64, err error) {
	retValue, err := instance.GetProperty("Counter9")
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
