// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ApplicationPoolProcessorSettings struct
type ApplicationPoolProcessorSettings struct {
	*EmbeddedObject

	//
	Action int32

	//
	Limit uint32

	//
	NumaNodeAffinityMode int32

	//
	NumaNodeAssignment int32

	//
	processorGroup int32

	//
	ResetInterval string

	//
	SmpAffinitized bool

	//
	SmpProcessorAffinityMask uint32

	//
	SmpProcessorAffinityMask2 uint32
}

func NewApplicationPoolProcessorSettingsEx1(instance *cim.WmiInstance) (newInstance *ApplicationPoolProcessorSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ApplicationPoolProcessorSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewApplicationPoolProcessorSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ApplicationPoolProcessorSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ApplicationPoolProcessorSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetAction sets the value of Action for the instance
func (instance *ApplicationPoolProcessorSettings) SetPropertyAction(value int32) (err error) {
	return instance.SetProperty("Action", (value))
}

// GetAction gets the value of Action for the instance
func (instance *ApplicationPoolProcessorSettings) GetPropertyAction() (value int32, err error) {
	retValue, err := instance.GetProperty("Action")
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

// SetLimit sets the value of Limit for the instance
func (instance *ApplicationPoolProcessorSettings) SetPropertyLimit(value uint32) (err error) {
	return instance.SetProperty("Limit", (value))
}

// GetLimit gets the value of Limit for the instance
func (instance *ApplicationPoolProcessorSettings) GetPropertyLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("Limit")
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

// SetNumaNodeAffinityMode sets the value of NumaNodeAffinityMode for the instance
func (instance *ApplicationPoolProcessorSettings) SetPropertyNumaNodeAffinityMode(value int32) (err error) {
	return instance.SetProperty("NumaNodeAffinityMode", (value))
}

// GetNumaNodeAffinityMode gets the value of NumaNodeAffinityMode for the instance
func (instance *ApplicationPoolProcessorSettings) GetPropertyNumaNodeAffinityMode() (value int32, err error) {
	retValue, err := instance.GetProperty("NumaNodeAffinityMode")
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

// SetNumaNodeAssignment sets the value of NumaNodeAssignment for the instance
func (instance *ApplicationPoolProcessorSettings) SetPropertyNumaNodeAssignment(value int32) (err error) {
	return instance.SetProperty("NumaNodeAssignment", (value))
}

// GetNumaNodeAssignment gets the value of NumaNodeAssignment for the instance
func (instance *ApplicationPoolProcessorSettings) GetPropertyNumaNodeAssignment() (value int32, err error) {
	retValue, err := instance.GetProperty("NumaNodeAssignment")
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

// SetprocessorGroup sets the value of processorGroup for the instance
func (instance *ApplicationPoolProcessorSettings) SetPropertyprocessorGroup(value int32) (err error) {
	return instance.SetProperty("processorGroup", (value))
}

// GetprocessorGroup gets the value of processorGroup for the instance
func (instance *ApplicationPoolProcessorSettings) GetPropertyprocessorGroup() (value int32, err error) {
	retValue, err := instance.GetProperty("processorGroup")
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

// SetResetInterval sets the value of ResetInterval for the instance
func (instance *ApplicationPoolProcessorSettings) SetPropertyResetInterval(value string) (err error) {
	return instance.SetProperty("ResetInterval", (value))
}

// GetResetInterval gets the value of ResetInterval for the instance
func (instance *ApplicationPoolProcessorSettings) GetPropertyResetInterval() (value string, err error) {
	retValue, err := instance.GetProperty("ResetInterval")
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

// SetSmpAffinitized sets the value of SmpAffinitized for the instance
func (instance *ApplicationPoolProcessorSettings) SetPropertySmpAffinitized(value bool) (err error) {
	return instance.SetProperty("SmpAffinitized", (value))
}

// GetSmpAffinitized gets the value of SmpAffinitized for the instance
func (instance *ApplicationPoolProcessorSettings) GetPropertySmpAffinitized() (value bool, err error) {
	retValue, err := instance.GetProperty("SmpAffinitized")
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

// SetSmpProcessorAffinityMask sets the value of SmpProcessorAffinityMask for the instance
func (instance *ApplicationPoolProcessorSettings) SetPropertySmpProcessorAffinityMask(value uint32) (err error) {
	return instance.SetProperty("SmpProcessorAffinityMask", (value))
}

// GetSmpProcessorAffinityMask gets the value of SmpProcessorAffinityMask for the instance
func (instance *ApplicationPoolProcessorSettings) GetPropertySmpProcessorAffinityMask() (value uint32, err error) {
	retValue, err := instance.GetProperty("SmpProcessorAffinityMask")
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

// SetSmpProcessorAffinityMask2 sets the value of SmpProcessorAffinityMask2 for the instance
func (instance *ApplicationPoolProcessorSettings) SetPropertySmpProcessorAffinityMask2(value uint32) (err error) {
	return instance.SetProperty("SmpProcessorAffinityMask2", (value))
}

// GetSmpProcessorAffinityMask2 gets the value of SmpProcessorAffinityMask2 for the instance
func (instance *ApplicationPoolProcessorSettings) GetPropertySmpProcessorAffinityMask2() (value uint32, err error) {
	retValue, err := instance.GetProperty("SmpProcessorAffinityMask2")
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
