// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// NlSecChanlSetup_Start struct
type NlSecChanlSetup_Start struct {
	*NlSecChanlSetup
}

func NewNlSecChanlSetup_StartEx1(instance *cim.WmiInstance) (newInstance *NlSecChanlSetup_Start, err error) {
	tmp, err := NewNlSecChanlSetupEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NlSecChanlSetup_Start{
		NlSecChanlSetup: tmp,
	}
	return
}

func NewNlSecChanlSetup_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NlSecChanlSetup_Start, err error) {
	tmp, err := NewNlSecChanlSetupEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NlSecChanlSetup_Start{
		NlSecChanlSetup: tmp,
	}
	return
}
