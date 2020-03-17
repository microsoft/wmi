// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_iSCSITarget struct
type MSFT_iSCSITarget struct {
	cim.WmiInstance

	//
	IsConnected bool

	//
	NodeAddress string
}

// SetIsConnected sets the value of IsConnected for the instance
func (instance *MSFT_iSCSITarget) SetPropertyIsConnected(value bool) (err error) {
	return instance.SetProperty("IsConnected", value)
}

// GetIsConnected gets the value of IsConnected for the instance
func (instance *MSFT_iSCSITarget) GetPropertyIsConnected() (value bool, err error) {
	retValue, err := instance.GetProperty("IsConnected")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNodeAddress sets the value of NodeAddress for the instance
func (instance *MSFT_iSCSITarget) SetPropertyNodeAddress(value string) (err error) {
	return instance.SetProperty("NodeAddress", value)
}

// GetNodeAddress gets the value of NodeAddress for the instance
func (instance *MSFT_iSCSITarget) GetPropertyNodeAddress() (value string, err error) {
	retValue, err := instance.GetProperty("NodeAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="SessionIdentifier" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_iSCSITarget) Disconnect( /* IN */ SessionIdentifier string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disconnect", SessionIdentifier)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_iSCSITarget) Update() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Update")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="AuthenticationType" type="string "></param>
// <param name="ChapSecret" type="string "></param>
// <param name="ChapUsername" type="string "></param>
// <param name="InitiatorInstanceName" type="string "></param>
// <param name="InitiatorPortalAddress" type="string "></param>
// <param name="IsDataDigest" type="bool "></param>
// <param name="IsHeaderDigest" type="bool "></param>
// <param name="IsMultipathEnabled" type="bool "></param>
// <param name="IsPersistent" type="bool "></param>
// <param name="NodeAddress" type="string "></param>
// <param name="ReportToPnP" type="bool "></param>
// <param name="TargetPortalAddress" type="string "></param>
// <param name="TargetPortalPortNumber" type="uint16 "></param>

// <param name="CreatediSCSISession" type="MSFT_iSCSISession "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_iSCSITarget) Connect( /* IN */ NodeAddress string,
	/* IN */ TargetPortalAddress string,
	/* IN */ TargetPortalPortNumber uint16,
	/* IN */ InitiatorPortalAddress string,
	/* IN */ IsDataDigest bool,
	/* IN */ IsHeaderDigest bool,
	/* IN */ ReportToPnP bool,
	/* IN */ AuthenticationType string,
	/* IN */ ChapUsername string,
	/* IN */ ChapSecret string,
	/* IN */ IsMultipathEnabled bool,
	/* IN */ IsPersistent bool,
	/* IN */ InitiatorInstanceName string,
	/* OUT */ CreatediSCSISession MSFT_iSCSISession) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Connect", NodeAddress, TargetPortalAddress, TargetPortalPortNumber, InitiatorPortalAddress, IsDataDigest, IsHeaderDigest, ReportToPnP, AuthenticationType, ChapUsername, ChapSecret, IsMultipathEnabled, IsPersistent, InitiatorInstanceName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
