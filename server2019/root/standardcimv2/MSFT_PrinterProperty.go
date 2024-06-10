// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_PrinterProperty struct
type MSFT_PrinterProperty struct {
	*CIM_ManagedElement

	//
	ComputerName string

	//
	PrinterName string

	//
	PropertyName string

	//
	Type uint32

	//
	Value string
}

func NewMSFT_PrinterPropertyEx1(instance *cim.WmiInstance) (newInstance *MSFT_PrinterProperty, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_PrinterProperty{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_PrinterPropertyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_PrinterProperty, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_PrinterProperty{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetComputerName sets the value of ComputerName for the instance
func (instance *MSFT_PrinterProperty) SetPropertyComputerName(value string) (err error) {
	return instance.SetProperty("ComputerName", (value))
}

// GetComputerName gets the value of ComputerName for the instance
func (instance *MSFT_PrinterProperty) GetPropertyComputerName() (value string, err error) {
	retValue, err := instance.GetProperty("ComputerName")
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

// SetPrinterName sets the value of PrinterName for the instance
func (instance *MSFT_PrinterProperty) SetPropertyPrinterName(value string) (err error) {
	return instance.SetProperty("PrinterName", (value))
}

// GetPrinterName gets the value of PrinterName for the instance
func (instance *MSFT_PrinterProperty) GetPropertyPrinterName() (value string, err error) {
	retValue, err := instance.GetProperty("PrinterName")
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

// SetPropertyName sets the value of PropertyName for the instance
func (instance *MSFT_PrinterProperty) SetPropertyPropertyName(value string) (err error) {
	return instance.SetProperty("PropertyName", (value))
}

// GetPropertyName gets the value of PropertyName for the instance
func (instance *MSFT_PrinterProperty) GetPropertyPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("PropertyName")
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

// SetType sets the value of Type for the instance
func (instance *MSFT_PrinterProperty) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSFT_PrinterProperty) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
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

// SetValue sets the value of Value for the instance
func (instance *MSFT_PrinterProperty) SetPropertyValue(value string) (err error) {
	return instance.SetProperty("Value", (value))
}

// GetValue gets the value of Value for the instance
func (instance *MSFT_PrinterProperty) GetPropertyValue() (value string, err error) {
	retValue, err := instance.GetProperty("Value")
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

//

// <param name="ComputerName" type="string "></param>
// <param name="PrinterName" type="string "></param>
// <param name="PropertyName" type="string "></param>
// <param name="Value" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrinterProperty) SetByPrinterName( /* IN */ ComputerName string,
	/* IN */ PrinterName string,
	/* IN */ PropertyName string,
	/* IN */ Value string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByPrinterName", ComputerName, PrinterName, PropertyName, Value)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="InputObject" type="MSFT_PrinterProperty "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrinterProperty) SetByPrinterPropertyObject( /* IN */ InputObject MSFT_PrinterProperty) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByPrinterPropertyObject", InputObject)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="PrinterObject" type="MSFT_Printer "></param>
// <param name="PropertyName" type="string "></param>
// <param name="Value" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrinterProperty) SetByPrinterObject( /* IN */ PrinterObject MSFT_Printer,
	/* IN */ PropertyName string,
	/* IN */ Value string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByPrinterObject", PrinterObject, PropertyName, Value)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
