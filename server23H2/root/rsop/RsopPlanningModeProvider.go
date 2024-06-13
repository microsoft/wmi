// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP
//////////////////////////////////////////////
package rsop

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RsopPlanningModeProvider struct
type RsopPlanningModeProvider struct {
	*cim.WmiInstance
}

func NewRsopPlanningModeProviderEx1(instance *cim.WmiInstance) (newInstance *RsopPlanningModeProvider, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RsopPlanningModeProvider{
		WmiInstance: tmp,
	}
	return
}

func NewRsopPlanningModeProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RsopPlanningModeProvider, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RsopPlanningModeProvider{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="computerGPOFilters" type="string []"></param>
// <param name="computerName" type="string "></param>
// <param name="computerSecurityGroups" type="string []"></param>
// <param name="computerSOM" type="string "></param>
// <param name="flags" type="uint32 "></param>
// <param name="site" type="string "></param>
// <param name="userGPOFilters" type="string []"></param>
// <param name="userName" type="string "></param>
// <param name="userSecurityGroups" type="string []"></param>
// <param name="userSOM" type="string "></param>

// <param name="ExtendedInfo" type="uint32 "></param>
// <param name="hResult" type="uint32 "></param>
// <param name="nameSpace" type="string "></param>
func (instance *RsopPlanningModeProvider) RsopCreateSession( /* IN */ flags uint32,
	/* IN */ computerName string,
	/* IN */ computerSOM string,
	/* IN */ computerSecurityGroups []string,
	/* IN */ computerGPOFilters []string,
	/* IN */ userName string,
	/* IN */ userSOM string,
	/* IN */ userSecurityGroups []string,
	/* IN */ userGPOFilters []string,
	/* IN */ site string,
	/* OUT */ nameSpace string,
	/* OUT */ hResult uint32,
	/* OUT */ ExtendedInfo uint32) (err error) {
	_, err = instance.InvokeMethod("RsopCreateSession", flags, computerName, computerSOM, computerSecurityGroups, computerGPOFilters, userName, userSOM, userSecurityGroups, userGPOFilters, site)
	if err != nil {
		return
	}
	return

}

//

// <param name="namespace" type="string "></param>

// <param name="hResult" type="uint32 "></param>
func (instance *RsopPlanningModeProvider) RsopDeleteSession( /* IN */ Namespace string,
	/* OUT */ hResult uint32) (err error) {
	_, err = instance.InvokeMethod("RsopDeleteSession", Namespace)
	if err != nil {
		return
	}
	return

}
