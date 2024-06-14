// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NS4D70B888_0C83_4EC8_A075_63C360913417.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrLocalPrinterSetting struct
type RSOP_PolmkrLocalPrinterSetting struct {
	*RSOP_PolmkrPrinterSetting
}

func NewRSOP_PolmkrLocalPrinterSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrLocalPrinterSetting, err error) {
	tmp, err := NewRSOP_PolmkrPrinterSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrLocalPrinterSetting{
		RSOP_PolmkrPrinterSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrLocalPrinterSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrLocalPrinterSetting, err error) {
	tmp, err := NewRSOP_PolmkrPrinterSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrLocalPrinterSetting{
		RSOP_PolmkrPrinterSetting: tmp,
	}
	return
}
