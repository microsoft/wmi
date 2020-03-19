// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_3DPrinter struct
type MSFT_3DPrinter struct {
	*MSFT_Printer
}

func NewMSFT_3DPrinterEx1(instance *cim.WmiInstance) (newInstance *MSFT_3DPrinter, err error) {
	tmp, err := NewMSFT_PrinterEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_3DPrinter{
		MSFT_Printer: tmp,
	}
	return
}

func NewMSFT_3DPrinterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_3DPrinter, err error) {
	tmp, err := NewMSFT_PrinterEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_3DPrinter{
		MSFT_Printer: tmp,
	}
	return
}
