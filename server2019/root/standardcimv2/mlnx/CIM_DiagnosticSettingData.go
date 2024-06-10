// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_DiagnosticSettingData struct
type CIM_DiagnosticSettingData struct {
	*CIM_SettingData

	//
	HaltOnError bool

	//
	LogOptions []DiagnosticSettingData_LogOptions

	//
	LogStorage []DiagnosticSettingData_LogStorage

	//
	LoopControl []DiagnosticSettingData_LoopControl

	//
	LoopControlParameter []string

	//
	NonDestructive bool

	//
	OtherLogOptionsDescriptions []string

	//
	OtherLogStorageDescriptions []string

	//
	OtherLoopControlDescriptions []string

	//
	PercentOfTestCoverage uint8

	//
	QueryTimeout uint32

	//
	QuickMode bool

	//
	ResultPersistence uint32

	//
	VerbosityLevel []DiagnosticSettingData_VerbosityLevel
}

func NewCIM_DiagnosticSettingDataEx1(instance *cim.WmiInstance) (newInstance *CIM_DiagnosticSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_DiagnosticSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewCIM_DiagnosticSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_DiagnosticSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_DiagnosticSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetHaltOnError sets the value of HaltOnError for the instance
func (instance *CIM_DiagnosticSettingData) SetPropertyHaltOnError(value bool) (err error) {
	return instance.SetProperty("HaltOnError", (value))
}

// GetHaltOnError gets the value of HaltOnError for the instance
func (instance *CIM_DiagnosticSettingData) GetPropertyHaltOnError() (value bool, err error) {
	retValue, err := instance.GetProperty("HaltOnError")
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

// SetLogOptions sets the value of LogOptions for the instance
func (instance *CIM_DiagnosticSettingData) SetPropertyLogOptions(value []DiagnosticSettingData_LogOptions) (err error) {
	return instance.SetProperty("LogOptions", (value))
}

// GetLogOptions gets the value of LogOptions for the instance
func (instance *CIM_DiagnosticSettingData) GetPropertyLogOptions() (value []DiagnosticSettingData_LogOptions, err error) {
	retValue, err := instance.GetProperty("LogOptions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DiagnosticSettingData_LogOptions(valuetmp))
	}

	return
}

// SetLogStorage sets the value of LogStorage for the instance
func (instance *CIM_DiagnosticSettingData) SetPropertyLogStorage(value []DiagnosticSettingData_LogStorage) (err error) {
	return instance.SetProperty("LogStorage", (value))
}

// GetLogStorage gets the value of LogStorage for the instance
func (instance *CIM_DiagnosticSettingData) GetPropertyLogStorage() (value []DiagnosticSettingData_LogStorage, err error) {
	retValue, err := instance.GetProperty("LogStorage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DiagnosticSettingData_LogStorage(valuetmp))
	}

	return
}

// SetLoopControl sets the value of LoopControl for the instance
func (instance *CIM_DiagnosticSettingData) SetPropertyLoopControl(value []DiagnosticSettingData_LoopControl) (err error) {
	return instance.SetProperty("LoopControl", (value))
}

// GetLoopControl gets the value of LoopControl for the instance
func (instance *CIM_DiagnosticSettingData) GetPropertyLoopControl() (value []DiagnosticSettingData_LoopControl, err error) {
	retValue, err := instance.GetProperty("LoopControl")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DiagnosticSettingData_LoopControl(valuetmp))
	}

	return
}

// SetLoopControlParameter sets the value of LoopControlParameter for the instance
func (instance *CIM_DiagnosticSettingData) SetPropertyLoopControlParameter(value []string) (err error) {
	return instance.SetProperty("LoopControlParameter", (value))
}

// GetLoopControlParameter gets the value of LoopControlParameter for the instance
func (instance *CIM_DiagnosticSettingData) GetPropertyLoopControlParameter() (value []string, err error) {
	retValue, err := instance.GetProperty("LoopControlParameter")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetNonDestructive sets the value of NonDestructive for the instance
func (instance *CIM_DiagnosticSettingData) SetPropertyNonDestructive(value bool) (err error) {
	return instance.SetProperty("NonDestructive", (value))
}

// GetNonDestructive gets the value of NonDestructive for the instance
func (instance *CIM_DiagnosticSettingData) GetPropertyNonDestructive() (value bool, err error) {
	retValue, err := instance.GetProperty("NonDestructive")
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

// SetOtherLogOptionsDescriptions sets the value of OtherLogOptionsDescriptions for the instance
func (instance *CIM_DiagnosticSettingData) SetPropertyOtherLogOptionsDescriptions(value []string) (err error) {
	return instance.SetProperty("OtherLogOptionsDescriptions", (value))
}

// GetOtherLogOptionsDescriptions gets the value of OtherLogOptionsDescriptions for the instance
func (instance *CIM_DiagnosticSettingData) GetPropertyOtherLogOptionsDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherLogOptionsDescriptions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetOtherLogStorageDescriptions sets the value of OtherLogStorageDescriptions for the instance
func (instance *CIM_DiagnosticSettingData) SetPropertyOtherLogStorageDescriptions(value []string) (err error) {
	return instance.SetProperty("OtherLogStorageDescriptions", (value))
}

// GetOtherLogStorageDescriptions gets the value of OtherLogStorageDescriptions for the instance
func (instance *CIM_DiagnosticSettingData) GetPropertyOtherLogStorageDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherLogStorageDescriptions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetOtherLoopControlDescriptions sets the value of OtherLoopControlDescriptions for the instance
func (instance *CIM_DiagnosticSettingData) SetPropertyOtherLoopControlDescriptions(value []string) (err error) {
	return instance.SetProperty("OtherLoopControlDescriptions", (value))
}

// GetOtherLoopControlDescriptions gets the value of OtherLoopControlDescriptions for the instance
func (instance *CIM_DiagnosticSettingData) GetPropertyOtherLoopControlDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherLoopControlDescriptions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetPercentOfTestCoverage sets the value of PercentOfTestCoverage for the instance
func (instance *CIM_DiagnosticSettingData) SetPropertyPercentOfTestCoverage(value uint8) (err error) {
	return instance.SetProperty("PercentOfTestCoverage", (value))
}

// GetPercentOfTestCoverage gets the value of PercentOfTestCoverage for the instance
func (instance *CIM_DiagnosticSettingData) GetPropertyPercentOfTestCoverage() (value uint8, err error) {
	retValue, err := instance.GetProperty("PercentOfTestCoverage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetQueryTimeout sets the value of QueryTimeout for the instance
func (instance *CIM_DiagnosticSettingData) SetPropertyQueryTimeout(value uint32) (err error) {
	return instance.SetProperty("QueryTimeout", (value))
}

// GetQueryTimeout gets the value of QueryTimeout for the instance
func (instance *CIM_DiagnosticSettingData) GetPropertyQueryTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("QueryTimeout")
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

// SetQuickMode sets the value of QuickMode for the instance
func (instance *CIM_DiagnosticSettingData) SetPropertyQuickMode(value bool) (err error) {
	return instance.SetProperty("QuickMode", (value))
}

// GetQuickMode gets the value of QuickMode for the instance
func (instance *CIM_DiagnosticSettingData) GetPropertyQuickMode() (value bool, err error) {
	retValue, err := instance.GetProperty("QuickMode")
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

// SetResultPersistence sets the value of ResultPersistence for the instance
func (instance *CIM_DiagnosticSettingData) SetPropertyResultPersistence(value uint32) (err error) {
	return instance.SetProperty("ResultPersistence", (value))
}

// GetResultPersistence gets the value of ResultPersistence for the instance
func (instance *CIM_DiagnosticSettingData) GetPropertyResultPersistence() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResultPersistence")
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

// SetVerbosityLevel sets the value of VerbosityLevel for the instance
func (instance *CIM_DiagnosticSettingData) SetPropertyVerbosityLevel(value []DiagnosticSettingData_VerbosityLevel) (err error) {
	return instance.SetProperty("VerbosityLevel", (value))
}

// GetVerbosityLevel gets the value of VerbosityLevel for the instance
func (instance *CIM_DiagnosticSettingData) GetPropertyVerbosityLevel() (value []DiagnosticSettingData_VerbosityLevel, err error) {
	retValue, err := instance.GetProperty("VerbosityLevel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DiagnosticSettingData_VerbosityLevel(valuetmp))
	}

	return
}
