// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NS3BC654AA_392A_4C67_BCCB_7765E6837951
//////////////////////////////////////////////
package ns3bc654aa_392a_4c67_bccb_7765e6837951

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __PARAMETERS struct
type __PARAMETERS struct {
	*cim.WmiInstance
}

func New__PARAMETERSEx1(instance *cim.WmiInstance) (newInstance *__PARAMETERS, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &__PARAMETERS{
		WmiInstance: tmp,
	}
	return
}

func New__PARAMETERSEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__PARAMETERS, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__PARAMETERS{
		WmiInstance: tmp,
	}
	return
}
