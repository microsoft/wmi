// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_DiagnosticTest struct
type CIM_DiagnosticTest struct {
	CIM_DiagnosticService

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

// SetCharacteristics sets the value of Characteristics for the instance
func (instance *CIM_DiagnosticTest) SetPropertyCharacteristics(value []DiagnosticTest_Characteristics) (err error) {
	return instance.SetProperty("Characteristics", value)
}

// GetCharacteristics gets the value of Characteristics for the instance
func (instance *CIM_DiagnosticTest) GetPropertyCharacteristics() (value []DiagnosticTest_Characteristics, err error) {
	retValue, err := instance.GetProperty("Characteristics")
	if err != nil {
		return
	}
	value, ok := retValue.([]DiagnosticTest_Characteristics)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsInUse sets the value of IsInUse for the instance
func (instance *CIM_DiagnosticTest) SetPropertyIsInUse(value bool) (err error) {
	return instance.SetProperty("IsInUse", value)
}

// GetIsInUse gets the value of IsInUse for the instance
func (instance *CIM_DiagnosticTest) GetPropertyIsInUse() (value bool, err error) {
	retValue, err := instance.GetProperty("IsInUse")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherCharacteristicDescription sets the value of OtherCharacteristicDescription for the instance
func (instance *CIM_DiagnosticTest) SetPropertyOtherCharacteristicDescription(value string) (err error) {
	return instance.SetProperty("OtherCharacteristicDescription", value)
}

// GetOtherCharacteristicDescription gets the value of OtherCharacteristicDescription for the instance
func (instance *CIM_DiagnosticTest) GetPropertyOtherCharacteristicDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherCharacteristicDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherCharacteristicsDescriptions sets the value of OtherCharacteristicsDescriptions for the instance
func (instance *CIM_DiagnosticTest) SetPropertyOtherCharacteristicsDescriptions(value []string) (err error) {
	return instance.SetProperty("OtherCharacteristicsDescriptions", value)
}

// GetOtherCharacteristicsDescriptions gets the value of OtherCharacteristicsDescriptions for the instance
func (instance *CIM_DiagnosticTest) GetPropertyOtherCharacteristicsDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherCharacteristicsDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherTestTypesDescriptions sets the value of OtherTestTypesDescriptions for the instance
func (instance *CIM_DiagnosticTest) SetPropertyOtherTestTypesDescriptions(value []string) (err error) {
	return instance.SetProperty("OtherTestTypesDescriptions", value)
}

// GetOtherTestTypesDescriptions gets the value of OtherTestTypesDescriptions for the instance
func (instance *CIM_DiagnosticTest) GetPropertyOtherTestTypesDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherTestTypesDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourcesUsed sets the value of ResourcesUsed for the instance
func (instance *CIM_DiagnosticTest) SetPropertyResourcesUsed(value []DiagnosticTest_ResourcesUsed) (err error) {
	return instance.SetProperty("ResourcesUsed", value)
}

// GetResourcesUsed gets the value of ResourcesUsed for the instance
func (instance *CIM_DiagnosticTest) GetPropertyResourcesUsed() (value []DiagnosticTest_ResourcesUsed, err error) {
	retValue, err := instance.GetProperty("ResourcesUsed")
	if err != nil {
		return
	}
	value, ok := retValue.([]DiagnosticTest_ResourcesUsed)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTestTypes sets the value of TestTypes for the instance
func (instance *CIM_DiagnosticTest) SetPropertyTestTypes(value []DiagnosticTest_TestTypes) (err error) {
	return instance.SetProperty("TestTypes", value)
}

// GetTestTypes gets the value of TestTypes for the instance
func (instance *CIM_DiagnosticTest) GetPropertyTestTypes() (value []DiagnosticTest_TestTypes, err error) {
	retValue, err := instance.GetProperty("TestTypes")
	if err != nil {
		return
	}
	value, ok := retValue.([]DiagnosticTest_TestTypes)
	if !ok {
		// TODO: Set an error
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
