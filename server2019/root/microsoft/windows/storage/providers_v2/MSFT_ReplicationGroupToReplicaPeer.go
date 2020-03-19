// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ReplicationGroupToReplicaPeer struct
type MSFT_ReplicationGroupToReplicaPeer struct {
	*MSFT_Synchronized

	//
	ReplicaPeer MSFT_ReplicaPeer

	//
	ReplicationGroup MSFT_ReplicationGroup
}

func NewMSFT_ReplicationGroupToReplicaPeerEx1(instance *cim.WmiInstance) (newInstance *MSFT_ReplicationGroupToReplicaPeer, err error) {
	tmp, err := NewMSFT_SynchronizedEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ReplicationGroupToReplicaPeer{
		MSFT_Synchronized: tmp,
	}
	return
}

func NewMSFT_ReplicationGroupToReplicaPeerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ReplicationGroupToReplicaPeer, err error) {
	tmp, err := NewMSFT_SynchronizedEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ReplicationGroupToReplicaPeer{
		MSFT_Synchronized: tmp,
	}
	return
}

// SetReplicaPeer sets the value of ReplicaPeer for the instance
func (instance *MSFT_ReplicationGroupToReplicaPeer) SetPropertyReplicaPeer(value MSFT_ReplicaPeer) (err error) {
	return instance.SetProperty("ReplicaPeer", value)
}

// GetReplicaPeer gets the value of ReplicaPeer for the instance
func (instance *MSFT_ReplicationGroupToReplicaPeer) GetPropertyReplicaPeer() (value MSFT_ReplicaPeer, err error) {
	retValue, err := instance.GetProperty("ReplicaPeer")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_ReplicaPeer)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicationGroup sets the value of ReplicationGroup for the instance
func (instance *MSFT_ReplicationGroupToReplicaPeer) SetPropertyReplicationGroup(value MSFT_ReplicationGroup) (err error) {
	return instance.SetProperty("ReplicationGroup", value)
}

// GetReplicationGroup gets the value of ReplicationGroup for the instance
func (instance *MSFT_ReplicationGroupToReplicaPeer) GetPropertyReplicationGroup() (value MSFT_ReplicationGroup, err error) {
	retValue, err := instance.GetProperty("ReplicationGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_ReplicationGroup)
	if !ok {
		// TODO: Set an error
	}
	return
}
