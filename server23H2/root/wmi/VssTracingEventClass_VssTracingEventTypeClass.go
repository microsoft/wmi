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

// VssTracingEventClass_VssTracingEventTypeClass struct
type VssTracingEventClass_VssTracingEventTypeClass struct {
	*VssTracingEventClass

	//
	FunctionName string

	//
	Indent uint32

	//
	LineNumber uint32

	//
	MessageDescription string

	//
	ModuleIndex uint32

	//
	SourceFileAlias string

	//
	SourceFileName string

	//
	Text string
}

func NewVssTracingEventClass_VssTracingEventTypeClassEx1(instance *cim.WmiInstance) (newInstance *VssTracingEventClass_VssTracingEventTypeClass, err error) {
	tmp, err := NewVssTracingEventClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &VssTracingEventClass_VssTracingEventTypeClass{
		VssTracingEventClass: tmp,
	}
	return
}

func NewVssTracingEventClass_VssTracingEventTypeClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VssTracingEventClass_VssTracingEventTypeClass, err error) {
	tmp, err := NewVssTracingEventClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VssTracingEventClass_VssTracingEventTypeClass{
		VssTracingEventClass: tmp,
	}
	return
}

// SetFunctionName sets the value of FunctionName for the instance
func (instance *VssTracingEventClass_VssTracingEventTypeClass) SetPropertyFunctionName(value string) (err error) {
	return instance.SetProperty("FunctionName", (value))
}

// GetFunctionName gets the value of FunctionName for the instance
func (instance *VssTracingEventClass_VssTracingEventTypeClass) GetPropertyFunctionName() (value string, err error) {
	retValue, err := instance.GetProperty("FunctionName")
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

// SetIndent sets the value of Indent for the instance
func (instance *VssTracingEventClass_VssTracingEventTypeClass) SetPropertyIndent(value uint32) (err error) {
	return instance.SetProperty("Indent", (value))
}

// GetIndent gets the value of Indent for the instance
func (instance *VssTracingEventClass_VssTracingEventTypeClass) GetPropertyIndent() (value uint32, err error) {
	retValue, err := instance.GetProperty("Indent")
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

// SetLineNumber sets the value of LineNumber for the instance
func (instance *VssTracingEventClass_VssTracingEventTypeClass) SetPropertyLineNumber(value uint32) (err error) {
	return instance.SetProperty("LineNumber", (value))
}

// GetLineNumber gets the value of LineNumber for the instance
func (instance *VssTracingEventClass_VssTracingEventTypeClass) GetPropertyLineNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("LineNumber")
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

// SetMessageDescription sets the value of MessageDescription for the instance
func (instance *VssTracingEventClass_VssTracingEventTypeClass) SetPropertyMessageDescription(value string) (err error) {
	return instance.SetProperty("MessageDescription", (value))
}

// GetMessageDescription gets the value of MessageDescription for the instance
func (instance *VssTracingEventClass_VssTracingEventTypeClass) GetPropertyMessageDescription() (value string, err error) {
	retValue, err := instance.GetProperty("MessageDescription")
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

// SetModuleIndex sets the value of ModuleIndex for the instance
func (instance *VssTracingEventClass_VssTracingEventTypeClass) SetPropertyModuleIndex(value uint32) (err error) {
	return instance.SetProperty("ModuleIndex", (value))
}

// GetModuleIndex gets the value of ModuleIndex for the instance
func (instance *VssTracingEventClass_VssTracingEventTypeClass) GetPropertyModuleIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("ModuleIndex")
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

// SetSourceFileAlias sets the value of SourceFileAlias for the instance
func (instance *VssTracingEventClass_VssTracingEventTypeClass) SetPropertySourceFileAlias(value string) (err error) {
	return instance.SetProperty("SourceFileAlias", (value))
}

// GetSourceFileAlias gets the value of SourceFileAlias for the instance
func (instance *VssTracingEventClass_VssTracingEventTypeClass) GetPropertySourceFileAlias() (value string, err error) {
	retValue, err := instance.GetProperty("SourceFileAlias")
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

// SetSourceFileName sets the value of SourceFileName for the instance
func (instance *VssTracingEventClass_VssTracingEventTypeClass) SetPropertySourceFileName(value string) (err error) {
	return instance.SetProperty("SourceFileName", (value))
}

// GetSourceFileName gets the value of SourceFileName for the instance
func (instance *VssTracingEventClass_VssTracingEventTypeClass) GetPropertySourceFileName() (value string, err error) {
	retValue, err := instance.GetProperty("SourceFileName")
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

// SetText sets the value of Text for the instance
func (instance *VssTracingEventClass_VssTracingEventTypeClass) SetPropertyText(value string) (err error) {
	return instance.SetProperty("Text", (value))
}

// GetText gets the value of Text for the instance
func (instance *VssTracingEventClass_VssTracingEventTypeClass) GetPropertyText() (value string, err error) {
	retValue, err := instance.GetProperty("Text")
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
