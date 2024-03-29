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

// CIM_DiagnosticTest struct
type CIM_DiagnosticTest struct {
	*CIM_DiagnosticService

	//
	Characteristics []DiagnosticTest_Characteristics

	//
	IsInUse bool

	//
	OtherCharacteristicDescription string

	//
	OtherCharacteristicsDescriptions []string

	//
	OtherTestTypesDescriptions []string

	//
	ResourcesUsed []DiagnosticTest_ResourcesUsed

	//
	TestTypes []DiagnosticTest_TestTypes
}

func NewCIM_DiagnosticTestEx1(instance *cim.WmiInstance) (newInstance *CIM_DiagnosticTest, err error) {
	tmp, err := NewCIM_DiagnosticServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_DiagnosticTest{
		CIM_DiagnosticService: tmp,
	}
	return
}

func NewCIM_DiagnosticTestEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_DiagnosticTest, err error) {
	tmp, err := NewCIM_DiagnosticServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_DiagnosticTest{
		CIM_DiagnosticService: tmp,
	}
	return
}

// SetCharacteristics sets the value of Characteristics for the instance
func (instance *CIM_DiagnosticTest) SetPropertyCharacteristics(value []DiagnosticTest_Characteristics) (err error) {
	return instance.SetProperty("Characteristics", (value))
}

// GetCharacteristics gets the value of Characteristics for the instance
func (instance *CIM_DiagnosticTest) GetPropertyCharacteristics() (value []DiagnosticTest_Characteristics, err error) {
	retValue, err := instance.GetProperty("Characteristics")
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
		value = append(value, DiagnosticTest_Characteristics(valuetmp))
	}

	return
}

// SetIsInUse sets the value of IsInUse for the instance
func (instance *CIM_DiagnosticTest) SetPropertyIsInUse(value bool) (err error) {
	return instance.SetProperty("IsInUse", (value))
}

// GetIsInUse gets the value of IsInUse for the instance
func (instance *CIM_DiagnosticTest) GetPropertyIsInUse() (value bool, err error) {
	retValue, err := instance.GetProperty("IsInUse")
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

// SetOtherCharacteristicDescription sets the value of OtherCharacteristicDescription for the instance
func (instance *CIM_DiagnosticTest) SetPropertyOtherCharacteristicDescription(value string) (err error) {
	return instance.SetProperty("OtherCharacteristicDescription", (value))
}

// GetOtherCharacteristicDescription gets the value of OtherCharacteristicDescription for the instance
func (instance *CIM_DiagnosticTest) GetPropertyOtherCharacteristicDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherCharacteristicDescription")
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

// SetOtherCharacteristicsDescriptions sets the value of OtherCharacteristicsDescriptions for the instance
func (instance *CIM_DiagnosticTest) SetPropertyOtherCharacteristicsDescriptions(value []string) (err error) {
	return instance.SetProperty("OtherCharacteristicsDescriptions", (value))
}

// GetOtherCharacteristicsDescriptions gets the value of OtherCharacteristicsDescriptions for the instance
func (instance *CIM_DiagnosticTest) GetPropertyOtherCharacteristicsDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherCharacteristicsDescriptions")
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

// SetOtherTestTypesDescriptions sets the value of OtherTestTypesDescriptions for the instance
func (instance *CIM_DiagnosticTest) SetPropertyOtherTestTypesDescriptions(value []string) (err error) {
	return instance.SetProperty("OtherTestTypesDescriptions", (value))
}

// GetOtherTestTypesDescriptions gets the value of OtherTestTypesDescriptions for the instance
func (instance *CIM_DiagnosticTest) GetPropertyOtherTestTypesDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherTestTypesDescriptions")
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

// SetResourcesUsed sets the value of ResourcesUsed for the instance
func (instance *CIM_DiagnosticTest) SetPropertyResourcesUsed(value []DiagnosticTest_ResourcesUsed) (err error) {
	return instance.SetProperty("ResourcesUsed", (value))
}

// GetResourcesUsed gets the value of ResourcesUsed for the instance
func (instance *CIM_DiagnosticTest) GetPropertyResourcesUsed() (value []DiagnosticTest_ResourcesUsed, err error) {
	retValue, err := instance.GetProperty("ResourcesUsed")
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
		value = append(value, DiagnosticTest_ResourcesUsed(valuetmp))
	}

	return
}

// SetTestTypes sets the value of TestTypes for the instance
func (instance *CIM_DiagnosticTest) SetPropertyTestTypes(value []DiagnosticTest_TestTypes) (err error) {
	return instance.SetProperty("TestTypes", (value))
}

// GetTestTypes gets the value of TestTypes for the instance
func (instance *CIM_DiagnosticTest) GetPropertyTestTypes() (value []DiagnosticTest_TestTypes, err error) {
	retValue, err := instance.GetProperty("TestTypes")
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
		value = append(value, DiagnosticTest_TestTypes(valuetmp))
	}

	return
}

//

// <param name="Setting" type="CIM_DiagnosticSetting "></param>
// <param name="SystemElement" type="CIM_ManagedSystemElement "></param>

// <param name="Result" type="CIM_DiagnosticResult "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_DiagnosticTest) RunTest( /* IN */ SystemElement CIM_ManagedSystemElement,
	/* IN */ Setting CIM_DiagnosticSetting,
	/* OUT */ Result CIM_DiagnosticResult) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RunTest", SystemElement, Setting)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SystemElement" type="CIM_ManagedSystemElement "></param>

// <param name="ResultsNotCleared" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_DiagnosticTest) ClearResults( /* IN */ SystemElement CIM_ManagedSystemElement,
	/* OUT */ ResultsNotCleared []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ClearResults", SystemElement)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Result" type="CIM_DiagnosticResult "></param>
// <param name="SystemElement" type="CIM_ManagedSystemElement "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="TestingStopped" type="bool "></param>
func (instance *CIM_DiagnosticTest) DiscontinueTest( /* IN */ SystemElement CIM_ManagedSystemElement,
	/* IN */ Result CIM_DiagnosticResult,
	/* OUT */ TestingStopped bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DiscontinueTest", SystemElement, Result)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
