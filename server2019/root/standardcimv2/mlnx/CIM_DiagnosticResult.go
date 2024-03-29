// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_DiagnosticResult struct
type CIM_DiagnosticResult struct {
	*cim.WmiInstance

	//
	DiagnosticCreationClassName string

	//
	DiagnosticName string

	//
	DiagSystemCreationClassName string

	//
	DiagSystemName string

	//
	ErrorCode []string

	//
	ErrorCount []uint32

	//
	EstimatedTimeOfPerforming uint32

	//
	ExecutionID string

	//
	HaltOnError bool

	//
	IsPackage bool

	//
	LoopsFailed uint32

	//
	LoopsPassed uint32

	//
	OtherStateDescription string

	//
	PercentComplete uint8

	//
	PercentOfTestCoverage uint8

	//
	QuickMode bool

	//
	ReportSoftErrors bool

	//
	ReportStatusMessages bool

	//
	TestCompletionTime string

	//
	TestResults []string

	//
	TestStartTime string

	//
	TestState DiagnosticResult_TestState

	//
	TestWarningLevel DiagnosticResult_TestWarningLevel

	//
	TimeStamp string
}

func NewCIM_DiagnosticResultEx1(instance *cim.WmiInstance) (newInstance *CIM_DiagnosticResult, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_DiagnosticResult{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_DiagnosticResultEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_DiagnosticResult, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_DiagnosticResult{
		WmiInstance: tmp,
	}
	return
}

// SetDiagnosticCreationClassName sets the value of DiagnosticCreationClassName for the instance
func (instance *CIM_DiagnosticResult) SetPropertyDiagnosticCreationClassName(value string) (err error) {
	return instance.SetProperty("DiagnosticCreationClassName", (value))
}

// GetDiagnosticCreationClassName gets the value of DiagnosticCreationClassName for the instance
func (instance *CIM_DiagnosticResult) GetPropertyDiagnosticCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("DiagnosticCreationClassName")
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

// SetDiagnosticName sets the value of DiagnosticName for the instance
func (instance *CIM_DiagnosticResult) SetPropertyDiagnosticName(value string) (err error) {
	return instance.SetProperty("DiagnosticName", (value))
}

// GetDiagnosticName gets the value of DiagnosticName for the instance
func (instance *CIM_DiagnosticResult) GetPropertyDiagnosticName() (value string, err error) {
	retValue, err := instance.GetProperty("DiagnosticName")
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

// SetDiagSystemCreationClassName sets the value of DiagSystemCreationClassName for the instance
func (instance *CIM_DiagnosticResult) SetPropertyDiagSystemCreationClassName(value string) (err error) {
	return instance.SetProperty("DiagSystemCreationClassName", (value))
}

// GetDiagSystemCreationClassName gets the value of DiagSystemCreationClassName for the instance
func (instance *CIM_DiagnosticResult) GetPropertyDiagSystemCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("DiagSystemCreationClassName")
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

// SetDiagSystemName sets the value of DiagSystemName for the instance
func (instance *CIM_DiagnosticResult) SetPropertyDiagSystemName(value string) (err error) {
	return instance.SetProperty("DiagSystemName", (value))
}

// GetDiagSystemName gets the value of DiagSystemName for the instance
func (instance *CIM_DiagnosticResult) GetPropertyDiagSystemName() (value string, err error) {
	retValue, err := instance.GetProperty("DiagSystemName")
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

// SetErrorCode sets the value of ErrorCode for the instance
func (instance *CIM_DiagnosticResult) SetPropertyErrorCode(value []string) (err error) {
	return instance.SetProperty("ErrorCode", (value))
}

// GetErrorCode gets the value of ErrorCode for the instance
func (instance *CIM_DiagnosticResult) GetPropertyErrorCode() (value []string, err error) {
	retValue, err := instance.GetProperty("ErrorCode")
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

// SetErrorCount sets the value of ErrorCount for the instance
func (instance *CIM_DiagnosticResult) SetPropertyErrorCount(value []uint32) (err error) {
	return instance.SetProperty("ErrorCount", (value))
}

// GetErrorCount gets the value of ErrorCount for the instance
func (instance *CIM_DiagnosticResult) GetPropertyErrorCount() (value []uint32, err error) {
	retValue, err := instance.GetProperty("ErrorCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetEstimatedTimeOfPerforming sets the value of EstimatedTimeOfPerforming for the instance
func (instance *CIM_DiagnosticResult) SetPropertyEstimatedTimeOfPerforming(value uint32) (err error) {
	return instance.SetProperty("EstimatedTimeOfPerforming", (value))
}

// GetEstimatedTimeOfPerforming gets the value of EstimatedTimeOfPerforming for the instance
func (instance *CIM_DiagnosticResult) GetPropertyEstimatedTimeOfPerforming() (value uint32, err error) {
	retValue, err := instance.GetProperty("EstimatedTimeOfPerforming")
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

// SetExecutionID sets the value of ExecutionID for the instance
func (instance *CIM_DiagnosticResult) SetPropertyExecutionID(value string) (err error) {
	return instance.SetProperty("ExecutionID", (value))
}

// GetExecutionID gets the value of ExecutionID for the instance
func (instance *CIM_DiagnosticResult) GetPropertyExecutionID() (value string, err error) {
	retValue, err := instance.GetProperty("ExecutionID")
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

// SetHaltOnError sets the value of HaltOnError for the instance
func (instance *CIM_DiagnosticResult) SetPropertyHaltOnError(value bool) (err error) {
	return instance.SetProperty("HaltOnError", (value))
}

// GetHaltOnError gets the value of HaltOnError for the instance
func (instance *CIM_DiagnosticResult) GetPropertyHaltOnError() (value bool, err error) {
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

// SetIsPackage sets the value of IsPackage for the instance
func (instance *CIM_DiagnosticResult) SetPropertyIsPackage(value bool) (err error) {
	return instance.SetProperty("IsPackage", (value))
}

// GetIsPackage gets the value of IsPackage for the instance
func (instance *CIM_DiagnosticResult) GetPropertyIsPackage() (value bool, err error) {
	retValue, err := instance.GetProperty("IsPackage")
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

// SetLoopsFailed sets the value of LoopsFailed for the instance
func (instance *CIM_DiagnosticResult) SetPropertyLoopsFailed(value uint32) (err error) {
	return instance.SetProperty("LoopsFailed", (value))
}

// GetLoopsFailed gets the value of LoopsFailed for the instance
func (instance *CIM_DiagnosticResult) GetPropertyLoopsFailed() (value uint32, err error) {
	retValue, err := instance.GetProperty("LoopsFailed")
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

// SetLoopsPassed sets the value of LoopsPassed for the instance
func (instance *CIM_DiagnosticResult) SetPropertyLoopsPassed(value uint32) (err error) {
	return instance.SetProperty("LoopsPassed", (value))
}

// GetLoopsPassed gets the value of LoopsPassed for the instance
func (instance *CIM_DiagnosticResult) GetPropertyLoopsPassed() (value uint32, err error) {
	retValue, err := instance.GetProperty("LoopsPassed")
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

// SetOtherStateDescription sets the value of OtherStateDescription for the instance
func (instance *CIM_DiagnosticResult) SetPropertyOtherStateDescription(value string) (err error) {
	return instance.SetProperty("OtherStateDescription", (value))
}

// GetOtherStateDescription gets the value of OtherStateDescription for the instance
func (instance *CIM_DiagnosticResult) GetPropertyOtherStateDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherStateDescription")
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

// SetPercentComplete sets the value of PercentComplete for the instance
func (instance *CIM_DiagnosticResult) SetPropertyPercentComplete(value uint8) (err error) {
	return instance.SetProperty("PercentComplete", (value))
}

// GetPercentComplete gets the value of PercentComplete for the instance
func (instance *CIM_DiagnosticResult) GetPropertyPercentComplete() (value uint8, err error) {
	retValue, err := instance.GetProperty("PercentComplete")
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

// SetPercentOfTestCoverage sets the value of PercentOfTestCoverage for the instance
func (instance *CIM_DiagnosticResult) SetPropertyPercentOfTestCoverage(value uint8) (err error) {
	return instance.SetProperty("PercentOfTestCoverage", (value))
}

// GetPercentOfTestCoverage gets the value of PercentOfTestCoverage for the instance
func (instance *CIM_DiagnosticResult) GetPropertyPercentOfTestCoverage() (value uint8, err error) {
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

// SetQuickMode sets the value of QuickMode for the instance
func (instance *CIM_DiagnosticResult) SetPropertyQuickMode(value bool) (err error) {
	return instance.SetProperty("QuickMode", (value))
}

// GetQuickMode gets the value of QuickMode for the instance
func (instance *CIM_DiagnosticResult) GetPropertyQuickMode() (value bool, err error) {
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

// SetReportSoftErrors sets the value of ReportSoftErrors for the instance
func (instance *CIM_DiagnosticResult) SetPropertyReportSoftErrors(value bool) (err error) {
	return instance.SetProperty("ReportSoftErrors", (value))
}

// GetReportSoftErrors gets the value of ReportSoftErrors for the instance
func (instance *CIM_DiagnosticResult) GetPropertyReportSoftErrors() (value bool, err error) {
	retValue, err := instance.GetProperty("ReportSoftErrors")
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

// SetReportStatusMessages sets the value of ReportStatusMessages for the instance
func (instance *CIM_DiagnosticResult) SetPropertyReportStatusMessages(value bool) (err error) {
	return instance.SetProperty("ReportStatusMessages", (value))
}

// GetReportStatusMessages gets the value of ReportStatusMessages for the instance
func (instance *CIM_DiagnosticResult) GetPropertyReportStatusMessages() (value bool, err error) {
	retValue, err := instance.GetProperty("ReportStatusMessages")
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

// SetTestCompletionTime sets the value of TestCompletionTime for the instance
func (instance *CIM_DiagnosticResult) SetPropertyTestCompletionTime(value string) (err error) {
	return instance.SetProperty("TestCompletionTime", (value))
}

// GetTestCompletionTime gets the value of TestCompletionTime for the instance
func (instance *CIM_DiagnosticResult) GetPropertyTestCompletionTime() (value string, err error) {
	retValue, err := instance.GetProperty("TestCompletionTime")
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

// SetTestResults sets the value of TestResults for the instance
func (instance *CIM_DiagnosticResult) SetPropertyTestResults(value []string) (err error) {
	return instance.SetProperty("TestResults", (value))
}

// GetTestResults gets the value of TestResults for the instance
func (instance *CIM_DiagnosticResult) GetPropertyTestResults() (value []string, err error) {
	retValue, err := instance.GetProperty("TestResults")
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

// SetTestStartTime sets the value of TestStartTime for the instance
func (instance *CIM_DiagnosticResult) SetPropertyTestStartTime(value string) (err error) {
	return instance.SetProperty("TestStartTime", (value))
}

// GetTestStartTime gets the value of TestStartTime for the instance
func (instance *CIM_DiagnosticResult) GetPropertyTestStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("TestStartTime")
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

// SetTestState sets the value of TestState for the instance
func (instance *CIM_DiagnosticResult) SetPropertyTestState(value DiagnosticResult_TestState) (err error) {
	return instance.SetProperty("TestState", (value))
}

// GetTestState gets the value of TestState for the instance
func (instance *CIM_DiagnosticResult) GetPropertyTestState() (value DiagnosticResult_TestState, err error) {
	retValue, err := instance.GetProperty("TestState")
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

	value = DiagnosticResult_TestState(valuetmp)

	return
}

// SetTestWarningLevel sets the value of TestWarningLevel for the instance
func (instance *CIM_DiagnosticResult) SetPropertyTestWarningLevel(value DiagnosticResult_TestWarningLevel) (err error) {
	return instance.SetProperty("TestWarningLevel", (value))
}

// GetTestWarningLevel gets the value of TestWarningLevel for the instance
func (instance *CIM_DiagnosticResult) GetPropertyTestWarningLevel() (value DiagnosticResult_TestWarningLevel, err error) {
	retValue, err := instance.GetProperty("TestWarningLevel")
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

	value = DiagnosticResult_TestWarningLevel(valuetmp)

	return
}

// SetTimeStamp sets the value of TimeStamp for the instance
func (instance *CIM_DiagnosticResult) SetPropertyTimeStamp(value string) (err error) {
	return instance.SetProperty("TimeStamp", (value))
}

// GetTimeStamp gets the value of TimeStamp for the instance
func (instance *CIM_DiagnosticResult) GetPropertyTimeStamp() (value string, err error) {
	retValue, err := instance.GetProperty("TimeStamp")
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
