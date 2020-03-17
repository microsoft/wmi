// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_LprPrinterPort struct
type MSFT_LprPrinterPort struct {
	MSFT_PrinterPort

	//
	HostName string

	//
	PrinterName string
}

// SetHostName sets the value of HostName for the instance
func (instance *MSFT_LprPrinterPort) SetPropertyHostName(value string) (err error) {
	return instance.SetProperty("HostName", value)
}

// GetHostName gets the value of HostName for the instance
func (instance *MSFT_LprPrinterPort) GetPropertyHostName() (value string, err error) {
	retValue, err := instance.GetProperty("HostName")
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
func (instance *MSFT_LprPrinterPort) SetPropertyPrinterName(value string) (err error) {
	return instance.SetProperty("PrinterName", value)
}

// GetPrinterName gets the value of PrinterName for the instance
func (instance *MSFT_LprPrinterPort) GetPropertyPrinterName() (value string, err error) {
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
