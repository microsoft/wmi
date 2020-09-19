// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetLldpAgent struct
type MSFT_NetLldpAgent struct {
	*MSFT_NetLldpMsap
}

func NewMSFT_NetLldpAgentEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetLldpAgent, err error) {
	tmp, err := NewMSFT_NetLldpMsapEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetLldpAgent{
		MSFT_NetLldpMsap: tmp,
	}
	return
}

func NewMSFT_NetLldpAgentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetLldpAgent, err error) {
	tmp, err := NewMSFT_NetLldpMsapEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetLldpAgent{
		MSFT_NetLldpMsap: tmp,
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetLldpAgent) Enable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetLldpAgent) Disable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
