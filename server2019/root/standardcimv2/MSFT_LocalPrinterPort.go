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

// MSFT_LocalPrinterPort struct
type MSFT_LocalPrinterPort struct {
	*MSFT_PrinterPort
}

func NewMSFT_LocalPrinterPortEx1(instance *cim.WmiInstance) (newInstance *MSFT_LocalPrinterPort, err error) {
	tmp, err := NewMSFT_PrinterPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_LocalPrinterPort{
		MSFT_PrinterPort: tmp,
	}
	return
}

func NewMSFT_LocalPrinterPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_LocalPrinterPort, err error) {
	tmp, err := NewMSFT_PrinterPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_LocalPrinterPort{
		MSFT_PrinterPort: tmp,
	}
	return
}
