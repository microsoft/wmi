// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NSE4751BA1_9751_4F55_96B3_16813F2826DE.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrPortPrinterSetting struct
type RSOP_PolmkrPortPrinterSetting struct {
	*RSOP_PolmkrPrinterSetting
}

func NewRSOP_PolmkrPortPrinterSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrPortPrinterSetting, err error) {
	tmp, err := NewRSOP_PolmkrPrinterSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrPortPrinterSetting{
		RSOP_PolmkrPrinterSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrPortPrinterSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrPortPrinterSetting, err error) {
	tmp, err := NewRSOP_PolmkrPrinterSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrPortPrinterSetting{
		RSOP_PolmkrPrinterSetting: tmp,
	}
	return
}
