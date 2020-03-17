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

// MDM_Policy_Result01_RemoteProcedureCall02 struct
type MDM_Policy_Result01_RemoteProcedureCall02 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string

	//
	RestrictUnauthenticatedRPCClients string

	//
	RPCEndpointMapperClientAuthentication string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_RemoteProcedureCall02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_RemoteProcedureCall02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_RemoteProcedureCall02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_RemoteProcedureCall02) GetPropertyParentID() (value string, err error) {
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

// SetRestrictUnauthenticatedRPCClients sets the value of RestrictUnauthenticatedRPCClients for the instance
func (instance *MDM_Policy_Result01_RemoteProcedureCall02) SetPropertyRestrictUnauthenticatedRPCClients(value string) (err error) {
	return instance.SetProperty("RestrictUnauthenticatedRPCClients", value)
}

// GetRestrictUnauthenticatedRPCClients gets the value of RestrictUnauthenticatedRPCClients for the instance
func (instance *MDM_Policy_Result01_RemoteProcedureCall02) GetPropertyRestrictUnauthenticatedRPCClients() (value string, err error) {
	retValue, err := instance.GetProperty("RestrictUnauthenticatedRPCClients")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRPCEndpointMapperClientAuthentication sets the value of RPCEndpointMapperClientAuthentication for the instance
func (instance *MDM_Policy_Result01_RemoteProcedureCall02) SetPropertyRPCEndpointMapperClientAuthentication(value string) (err error) {
	return instance.SetProperty("RPCEndpointMapperClientAuthentication", value)
}

// GetRPCEndpointMapperClientAuthentication gets the value of RPCEndpointMapperClientAuthentication for the instance
func (instance *MDM_Policy_Result01_RemoteProcedureCall02) GetPropertyRPCEndpointMapperClientAuthentication() (value string, err error) {
	retValue, err := instance.GetProperty("RPCEndpointMapperClientAuthentication")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
