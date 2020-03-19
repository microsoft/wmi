// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_AvailableDisk struct
type MSCluster_AvailableDisk struct {
	*MSCluster_ClusterDisk

	//
	ConnectedNodes []string

	//
	Node string

	//
	ResourceName string
}

func NewMSCluster_AvailableDiskEx1(instance *cim.WmiInstance) (newInstance *MSCluster_AvailableDisk, err error) {
	tmp, err := NewMSCluster_ClusterDiskEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_AvailableDisk{
		MSCluster_ClusterDisk: tmp,
	}
	return
}

func NewMSCluster_AvailableDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_AvailableDisk, err error) {
	tmp, err := NewMSCluster_ClusterDiskEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_AvailableDisk{
		MSCluster_ClusterDisk: tmp,
	}
	return
}

// SetConnectedNodes sets the value of ConnectedNodes for the instance
func (instance *MSCluster_AvailableDisk) SetPropertyConnectedNodes(value []string) (err error) {
	return instance.SetProperty("ConnectedNodes", value)
}

// GetConnectedNodes gets the value of ConnectedNodes for the instance
func (instance *MSCluster_AvailableDisk) GetPropertyConnectedNodes() (value []string, err error) {
	retValue, err := instance.GetProperty("ConnectedNodes")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNode sets the value of Node for the instance
func (instance *MSCluster_AvailableDisk) SetPropertyNode(value string) (err error) {
	return instance.SetProperty("Node", value)
}

// GetNode gets the value of Node for the instance
func (instance *MSCluster_AvailableDisk) GetPropertyNode() (value string, err error) {
	retValue, err := instance.GetProperty("Node")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceName sets the value of ResourceName for the instance
func (instance *MSCluster_AvailableDisk) SetPropertyResourceName(value string) (err error) {
	return instance.SetProperty("ResourceName", value)
}

// GetResourceName gets the value of ResourceName for the instance
func (instance *MSCluster_AvailableDisk) GetPropertyResourceName() (value string, err error) {
	retValue, err := instance.GetProperty("ResourceName")
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

// <param name="ResourceName" type="string "></param>

// <param name="Path" type="string "></param>
func (instance *MSCluster_AvailableDisk) AddToCluster( /* IN */ ResourceName string,
	/* OUT */ Path string) (err error) {
	_, err = instance.InvokeMethod("AddToCluster", ResourceName)
	if err != nil {
		return
	}
	return

}

//

// <param name="ResourceName" type="string "></param>

// <param name="Path" type="string "></param>
func (instance *MSCluster_AvailableDisk) CreateStorageResource( /* IN */ ResourceName string,
	/* OUT */ Path string) (err error) {
	_, err = instance.InvokeMethod("CreateStorageResource", ResourceName)
	if err != nil {
		return
	}
	return

}
