// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VideoHeadOnController struct
type Msvm_VideoHeadOnController struct {
	*CIM_VideoHeadOnController
}

func NewMsvm_VideoHeadOnControllerEx1(instance *cim.WmiInstance) (newInstance *Msvm_VideoHeadOnController, err error) {
	tmp, err := NewCIM_VideoHeadOnControllerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VideoHeadOnController{
		CIM_VideoHeadOnController: tmp,
	}
	return
}

func NewMsvm_VideoHeadOnControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VideoHeadOnController, err error) {
	tmp, err := NewCIM_VideoHeadOnControllerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VideoHeadOnController{
		CIM_VideoHeadOnController: tmp,
	}
	return
}
