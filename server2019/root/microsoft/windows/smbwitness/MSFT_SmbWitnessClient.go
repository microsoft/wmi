// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.SmbWitness
//////////////////////////////////////////////
package smbwitness

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_SmbWitnessClient struct
type MSFT_SmbWitnessClient struct {
	cim.WmiInstance

	//
	ClientName string

	//
	FileServerNodeName string

	//
	Flags uint32

	//
	IPAddress string

	//
	NetworkName string

	//
	NotificationsCancelled uint32

	//
	NotificationsSent uint32

	//
	QueuedNotifications uint32

	//
	ResourcesMonitored uint32

	//
	ShareName string

	//
	State SmbWitnessClient_State

	//
	WitnessNodeName string
}

// SetClientName sets the value of ClientName for the instance
func (instance *MSFT_SmbWitnessClient) SetPropertyClientName(value string) (err error) {
	return instance.SetProperty("ClientName", value)
}

// GetClientName gets the value of ClientName for the instance
func (instance *MSFT_SmbWitnessClient) GetPropertyClientName() (value string, err error) {
	retValue, err := instance.GetProperty("ClientName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileServerNodeName sets the value of FileServerNodeName for the instance
func (instance *MSFT_SmbWitnessClient) SetPropertyFileServerNodeName(value string) (err error) {
	return instance.SetProperty("FileServerNodeName", value)
}

// GetFileServerNodeName gets the value of FileServerNodeName for the instance
func (instance *MSFT_SmbWitnessClient) GetPropertyFileServerNodeName() (value string, err error) {
	retValue, err := instance.GetProperty("FileServerNodeName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MSFT_SmbWitnessClient) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", value)
}

// GetFlags gets the value of Flags for the instance
func (instance *MSFT_SmbWitnessClient) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPAddress sets the value of IPAddress for the instance
func (instance *MSFT_SmbWitnessClient) SetPropertyIPAddress(value string) (err error) {
	return instance.SetProperty("IPAddress", value)
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *MSFT_SmbWitnessClient) GetPropertyIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("IPAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkName sets the value of NetworkName for the instance
func (instance *MSFT_SmbWitnessClient) SetPropertyNetworkName(value string) (err error) {
	return instance.SetProperty("NetworkName", value)
}

// GetNetworkName gets the value of NetworkName for the instance
func (instance *MSFT_SmbWitnessClient) GetPropertyNetworkName() (value string, err error) {
	retValue, err := instance.GetProperty("NetworkName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNotificationsCancelled sets the value of NotificationsCancelled for the instance
func (instance *MSFT_SmbWitnessClient) SetPropertyNotificationsCancelled(value uint32) (err error) {
	return instance.SetProperty("NotificationsCancelled", value)
}

// GetNotificationsCancelled gets the value of NotificationsCancelled for the instance
func (instance *MSFT_SmbWitnessClient) GetPropertyNotificationsCancelled() (value uint32, err error) {
	retValue, err := instance.GetProperty("NotificationsCancelled")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNotificationsSent sets the value of NotificationsSent for the instance
func (instance *MSFT_SmbWitnessClient) SetPropertyNotificationsSent(value uint32) (err error) {
	return instance.SetProperty("NotificationsSent", value)
}

// GetNotificationsSent gets the value of NotificationsSent for the instance
func (instance *MSFT_SmbWitnessClient) GetPropertyNotificationsSent() (value uint32, err error) {
	retValue, err := instance.GetProperty("NotificationsSent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQueuedNotifications sets the value of QueuedNotifications for the instance
func (instance *MSFT_SmbWitnessClient) SetPropertyQueuedNotifications(value uint32) (err error) {
	return instance.SetProperty("QueuedNotifications", value)
}

// GetQueuedNotifications gets the value of QueuedNotifications for the instance
func (instance *MSFT_SmbWitnessClient) GetPropertyQueuedNotifications() (value uint32, err error) {
	retValue, err := instance.GetProperty("QueuedNotifications")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourcesMonitored sets the value of ResourcesMonitored for the instance
func (instance *MSFT_SmbWitnessClient) SetPropertyResourcesMonitored(value uint32) (err error) {
	return instance.SetProperty("ResourcesMonitored", value)
}

// GetResourcesMonitored gets the value of ResourcesMonitored for the instance
func (instance *MSFT_SmbWitnessClient) GetPropertyResourcesMonitored() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResourcesMonitored")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShareName sets the value of ShareName for the instance
func (instance *MSFT_SmbWitnessClient) SetPropertyShareName(value string) (err error) {
	return instance.SetProperty("ShareName", value)
}

// GetShareName gets the value of ShareName for the instance
func (instance *MSFT_SmbWitnessClient) GetPropertyShareName() (value string, err error) {
	retValue, err := instance.GetProperty("ShareName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_SmbWitnessClient) SetPropertyState(value SmbWitnessClient_State) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSFT_SmbWitnessClient) GetPropertyState() (value SmbWitnessClient_State, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(SmbWitnessClient_State)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWitnessNodeName sets the value of WitnessNodeName for the instance
func (instance *MSFT_SmbWitnessClient) SetPropertyWitnessNodeName(value string) (err error) {
	return instance.SetProperty("WitnessNodeName", value)
}

// GetWitnessNodeName gets the value of WitnessNodeName for the instance
func (instance *MSFT_SmbWitnessClient) GetPropertyWitnessNodeName() (value string, err error) {
	retValue, err := instance.GetProperty("WitnessNodeName")
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

// <param name="ClientName" type="string "></param>
// <param name="DestinationNode" type="string "></param>
// <param name="NetworkName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbWitnessClient) MoveClient( /* IN */ ClientName string,
	/* IN */ DestinationNode string,
	/* IN */ NetworkName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("MoveClient", ClientName, DestinationNode, NetworkName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
