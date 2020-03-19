// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_PrinterConfiguration struct
type MSFT_PrinterConfiguration struct {
	*cim.WmiInstance

	//
	Collate bool

	//
	Color bool

	//
	ComputerName string

	//
	DuplexingMode uint32

	//
	PaperSize uint32

	//
	PrintCapabilitiesXML string

	//
	PrinterName string

	//
	PrintTicketXML string
}

func NewMSFT_PrinterConfigurationEx1(instance *cim.WmiInstance) (newInstance *MSFT_PrinterConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_PrinterConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_PrinterConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_PrinterConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_PrinterConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetCollate sets the value of Collate for the instance
func (instance *MSFT_PrinterConfiguration) SetPropertyCollate(value bool) (err error) {
	return instance.SetProperty("Collate", value)
}

// GetCollate gets the value of Collate for the instance
func (instance *MSFT_PrinterConfiguration) GetPropertyCollate() (value bool, err error) {
	retValue, err := instance.GetProperty("Collate")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetColor sets the value of Color for the instance
func (instance *MSFT_PrinterConfiguration) SetPropertyColor(value bool) (err error) {
	return instance.SetProperty("Color", value)
}

// GetColor gets the value of Color for the instance
func (instance *MSFT_PrinterConfiguration) GetPropertyColor() (value bool, err error) {
	retValue, err := instance.GetProperty("Color")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetComputerName sets the value of ComputerName for the instance
func (instance *MSFT_PrinterConfiguration) SetPropertyComputerName(value string) (err error) {
	return instance.SetProperty("ComputerName", value)
}

// GetComputerName gets the value of ComputerName for the instance
func (instance *MSFT_PrinterConfiguration) GetPropertyComputerName() (value string, err error) {
	retValue, err := instance.GetProperty("ComputerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDuplexingMode sets the value of DuplexingMode for the instance
func (instance *MSFT_PrinterConfiguration) SetPropertyDuplexingMode(value uint32) (err error) {
	return instance.SetProperty("DuplexingMode", value)
}

// GetDuplexingMode gets the value of DuplexingMode for the instance
func (instance *MSFT_PrinterConfiguration) GetPropertyDuplexingMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("DuplexingMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPaperSize sets the value of PaperSize for the instance
func (instance *MSFT_PrinterConfiguration) SetPropertyPaperSize(value uint32) (err error) {
	return instance.SetProperty("PaperSize", value)
}

// GetPaperSize gets the value of PaperSize for the instance
func (instance *MSFT_PrinterConfiguration) GetPropertyPaperSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PaperSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrintCapabilitiesXML sets the value of PrintCapabilitiesXML for the instance
func (instance *MSFT_PrinterConfiguration) SetPropertyPrintCapabilitiesXML(value string) (err error) {
	return instance.SetProperty("PrintCapabilitiesXML", value)
}

// GetPrintCapabilitiesXML gets the value of PrintCapabilitiesXML for the instance
func (instance *MSFT_PrinterConfiguration) GetPropertyPrintCapabilitiesXML() (value string, err error) {
	retValue, err := instance.GetProperty("PrintCapabilitiesXML")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrinterName sets the value of PrinterName for the instance
func (instance *MSFT_PrinterConfiguration) SetPropertyPrinterName(value string) (err error) {
	return instance.SetProperty("PrinterName", value)
}

// GetPrinterName gets the value of PrinterName for the instance
func (instance *MSFT_PrinterConfiguration) GetPropertyPrinterName() (value string, err error) {
	retValue, err := instance.GetProperty("PrinterName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrintTicketXML sets the value of PrintTicketXML for the instance
func (instance *MSFT_PrinterConfiguration) SetPropertyPrintTicketXML(value string) (err error) {
	return instance.SetProperty("PrintTicketXML", value)
}

// GetPrintTicketXML gets the value of PrintTicketXML for the instance
func (instance *MSFT_PrinterConfiguration) GetPropertyPrintTicketXML() (value string, err error) {
	retValue, err := instance.GetProperty("PrintTicketXML")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="PrinterName" type="string "></param>

// <param name="cmdletOutput" type="MSFT_PrinterConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrinterConfiguration) GetByPrinterName( /* IN */ ComputerName string,
	/* IN */ PrinterName string,
	/* OUT */ cmdletOutput MSFT_PrinterConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByPrinterName", ComputerName, PrinterName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PrinterObject" type="MSFT_Printer "></param>

// <param name="cmdletOutput" type="MSFT_PrinterConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrinterConfiguration) GetByPrinterObject( /* IN */ PrinterObject MSFT_Printer,
	/* OUT */ cmdletOutput MSFT_PrinterConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByPrinterObject", PrinterObject)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Collate" type="bool "></param>
// <param name="Color" type="bool "></param>
// <param name="ComputerName" type="string "></param>
// <param name="DuplexingMode" type="uint32 "></param>
// <param name="PaperSize" type="uint32 "></param>
// <param name="PrinterName" type="string "></param>
// <param name="PrintTicketXML" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrinterConfiguration) SetByPrinterName( /* IN */ Collate bool,
	/* IN */ Color bool,
	/* IN */ DuplexingMode uint32,
	/* IN */ PaperSize uint32,
	/* IN */ PrintTicketXML string,
	/* IN */ ComputerName string,
	/* IN */ PrinterName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByPrinterName", Collate, Color, DuplexingMode, PaperSize, PrintTicketXML, ComputerName, PrinterName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Collate" type="bool "></param>
// <param name="Color" type="bool "></param>
// <param name="DuplexingMode" type="uint32 "></param>
// <param name="PaperSize" type="uint32 "></param>
// <param name="PrinterObject" type="MSFT_Printer "></param>
// <param name="PrintTicketXML" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrinterConfiguration) SetByPrinterObject( /* IN */ Collate bool,
	/* IN */ Color bool,
	/* IN */ DuplexingMode uint32,
	/* IN */ PaperSize uint32,
	/* IN */ PrintTicketXML string,
	/* IN */ PrinterObject MSFT_Printer) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByPrinterObject", Collate, Color, DuplexingMode, PaperSize, PrintTicketXML, PrinterObject)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="InputObject" type="MSFT_PrinterConfiguration "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrinterConfiguration) SetByPrintConfigObject( /* IN */ InputObject MSFT_PrinterConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByPrintConfigObject", InputObject)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
