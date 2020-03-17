// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Dns
//////////////////////////////////////////////
package dns

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_DnsClientNrptPolicy struct
type PS_DnsClientNrptPolicy struct {
	cim.WmiInstance
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
