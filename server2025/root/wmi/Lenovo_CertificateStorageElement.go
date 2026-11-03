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

// Lenovo_CertificateStorageElement struct
type Lenovo_CertificateStorageElement struct {
	*cim.WmiInstance
}

func NewLenovo_CertificateStorageElementEx1(instance *cim.WmiInstance) (newInstance *Lenovo_CertificateStorageElement, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Lenovo_CertificateStorageElement{
		WmiInstance: tmp,
	}
	return
}

func NewLenovo_CertificateStorageElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Lenovo_CertificateStorageElement, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Lenovo_CertificateStorageElement{
		WmiInstance: tmp,
	}
	return
}
