// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Dns
//////////////////////////////////////////////
package dns

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_DnsClientNrptPolicy struct
type PS_DnsClientNrptPolicy struct {
	*cim.WmiInstance
}

func NewPS_DnsClientNrptPolicyEx1(instance *cim.WmiInstance) (newInstance *PS_DnsClientNrptPolicy, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_DnsClientNrptPolicy{
		WmiInstance: tmp,
	}
	return
}

func NewPS_DnsClientNrptPolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_DnsClientNrptPolicy, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_DnsClientNrptPolicy{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="Effective" type="bool "></param>
// <param name="Namespace" type="string "></param>

// <param name="cmdletOutput" type="DnsClientPolicyConfiguration []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DnsClientNrptPolicy) Get( /* IN */ Effective bool,
	/* IN */ Namespace string,
	/* OUT */ cmdletOutput []DnsClientPolicyConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", Effective, Namespace)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
