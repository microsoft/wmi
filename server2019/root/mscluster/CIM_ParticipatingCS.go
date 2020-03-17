// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// CIM_ParticipatingCS struct
type CIM_ParticipatingCS struct {
	CIM_Dependency

	//
	RoleOfNode uint16

	//
	StateOfNode uint16
}

// SetRoleOfNode sets the value of RoleOfNode for the instance
func (instance *CIM_ParticipatingCS) SetPropertyRoleOfNode(value uint16) (err error) {
	return instance.SetProperty("RoleOfNode", value)
}

// GetRoleOfNode gets the value of RoleOfNode for the instance
func (instance *CIM_ParticipatingCS) GetPropertyRoleOfNode() (value uint16, err error) {
	retValue, err := instance.GetProperty("RoleOfNode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStateOfNode sets the value of StateOfNode for the instance
func (instance *CIM_ParticipatingCS) SetPropertyStateOfNode(value uint16) (err error) {
	return instance.SetProperty("StateOfNode", value)
}

// GetStateOfNode gets the value of StateOfNode for the instance
func (instance *CIM_ParticipatingCS) GetPropertyStateOfNode() (value uint16, err error) {
	retValue, err := instance.GetProperty("StateOfNode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
