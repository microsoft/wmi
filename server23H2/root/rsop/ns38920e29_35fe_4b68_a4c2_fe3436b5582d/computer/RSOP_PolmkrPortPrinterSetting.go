// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NS38920E29_35FE_4B68_A4C2_FE3436B5582D.Computer
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