// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Policy_Result01_WindowsConnectionManager02 struct
type MDM_Policy_Result01_WindowsConnectionManager02 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string

	//
	ProhitConnectionToNonDomainNetworksWhenConnectedToDomainAuthenticatedNetwork string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_WindowsConnectionManager02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_WindowsConnectionManager02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_WindowsConnectionManager02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_WindowsConnectionManager02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProhitConnectionToNonDomainNetworksWhenConnectedToDomainAuthenticatedNetwork sets the value of ProhitConnectionToNonDomainNetworksWhenConnectedToDomainAuthenticatedNetwork for the instance
func (instance *MDM_Policy_Result01_WindowsConnectionManager02) SetPropertyProhitConnectionToNonDomainNetworksWhenConnectedToDomainAuthenticatedNetwork(value string) (err error) {
	return instance.SetProperty("ProhitConnectionToNonDomainNetworksWhenConnectedToDomainAuthenticatedNetwork", value)
}

// GetProhitConnectionToNonDomainNetworksWhenConnectedToDomainAuthenticatedNetwork gets the value of ProhitConnectionToNonDomainNetworksWhenConnectedToDomainAuthenticatedNetwork for the instance
func (instance *MDM_Policy_Result01_WindowsConnectionManager02) GetPropertyProhitConnectionToNonDomainNetworksWhenConnectedToDomainAuthenticatedNetwork() (value string, err error) {
	retValue, err := instance.GetProperty("ProhitConnectionToNonDomainNetworksWhenConnectedToDomainAuthenticatedNetwork")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
