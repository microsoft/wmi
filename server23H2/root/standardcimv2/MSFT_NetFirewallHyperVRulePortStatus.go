// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetFirewallHyperVRulePortStatus struct
type MSFT_NetFirewallHyperVRulePortStatus struct {
	*cim.WmiInstance

	//
	Port MSFT_NetFirewallHyperVPort

	//
	Status uint16
}

func NewMSFT_NetFirewallHyperVRulePortStatusEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetFirewallHyperVRulePortStatus, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallHyperVRulePortStatus{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetFirewallHyperVRulePortStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetFirewallHyperVRulePortStatus, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallHyperVRulePortStatus{
		WmiInstance: tmp,
	}
	return
}

// SetPort sets the value of Port for the instance
func (instance *MSFT_NetFirewallHyperVRulePortStatus) SetPropertyPort(value MSFT_NetFirewallHyperVPort) (err error) {
	return instance.SetProperty("Port", (value))
}

// GetPort gets the value of Port for the instance
func (instance *MSFT_NetFirewallHyperVRulePortStatus) GetPropertyPort() (value MSFT_NetFirewallHyperVPort, err error) {
	retValue, err := instance.GetProperty("Port")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetFirewallHyperVPort)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetFirewallHyperVPort is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetFirewallHyperVPort(valuetmp)

	return
}

// SetStatus sets the value of Status for the instance
func (instance *MSFT_NetFirewallHyperVRulePortStatus) SetPropertyStatus(value uint16) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_NetFirewallHyperVRulePortStatus) GetPropertyStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
