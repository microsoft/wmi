// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Lenovo_BIOSElement struct
type Lenovo_BIOSElement struct {
	*cim.WmiInstance
}

func NewLenovo_BIOSElementEx1(instance *cim.WmiInstance) (newInstance *Lenovo_BIOSElement, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Lenovo_BIOSElement{
		WmiInstance: tmp,
	}
	return
}

func NewLenovo_BIOSElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Lenovo_BIOSElement, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Lenovo_BIOSElement{
		WmiInstance: tmp,
	}
	return
}
