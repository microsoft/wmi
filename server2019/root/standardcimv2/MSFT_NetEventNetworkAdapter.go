// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetEventNetworkAdapter struct
type MSFT_NetEventNetworkAdapter struct {
	*MSFT_NetEventPacketCaptureTarget

	//
	InterfaceDescription string

	//
	PromiscuousMode bool
}

func NewMSFT_NetEventNetworkAdapterEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetEventNetworkAdapter, err error) {
	tmp, err := NewMSFT_NetEventPacketCaptureTargetEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetEventNetworkAdapter{
		MSFT_NetEventPacketCaptureTarget: tmp,
	}
	return
}

func NewMSFT_NetEventNetworkAdapterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetEventNetworkAdapter, err error) {
	tmp, err := NewMSFT_NetEventPacketCaptureTargetEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetEventNetworkAdapter{
		MSFT_NetEventPacketCaptureTarget: tmp,
	}
	return
}

// SetInterfaceDescription sets the value of InterfaceDescription for the instance
func (instance *MSFT_NetEventNetworkAdapter) SetPropertyInterfaceDescription(value string) (err error) {
	return instance.SetProperty("InterfaceDescription", value)
}

// GetInterfaceDescription gets the value of InterfaceDescription for the instance
func (instance *MSFT_NetEventNetworkAdapter) GetPropertyInterfaceDescription() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPromiscuousMode sets the value of PromiscuousMode for the instance
func (instance *MSFT_NetEventNetworkAdapter) SetPropertyPromiscuousMode(value bool) (err error) {
	return instance.SetProperty("PromiscuousMode", value)
}

// GetPromiscuousMode gets the value of PromiscuousMode for the instance
func (instance *MSFT_NetEventNetworkAdapter) GetPropertyPromiscuousMode() (value bool, err error) {
	retValue, err := instance.GetProperty("PromiscuousMode")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
