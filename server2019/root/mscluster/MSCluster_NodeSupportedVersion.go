// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_NodeSupportedVersion struct
type MSCluster_NodeSupportedVersion struct {
	*cim.WmiInstance

	//
	ClusterFunctionalLevel uint32

	//
	ClusterUpgradeVersion uint32
}

func NewMSCluster_NodeSupportedVersionEx1(instance *cim.WmiInstance) (newInstance *MSCluster_NodeSupportedVersion, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_NodeSupportedVersion{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_NodeSupportedVersionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_NodeSupportedVersion, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_NodeSupportedVersion{
		WmiInstance: tmp,
	}
	return
}

// SetClusterFunctionalLevel sets the value of ClusterFunctionalLevel for the instance
func (instance *MSCluster_NodeSupportedVersion) SetPropertyClusterFunctionalLevel(value uint32) (err error) {
	return instance.SetProperty("ClusterFunctionalLevel", value)
}

// GetClusterFunctionalLevel gets the value of ClusterFunctionalLevel for the instance
func (instance *MSCluster_NodeSupportedVersion) GetPropertyClusterFunctionalLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClusterFunctionalLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetClusterUpgradeVersion sets the value of ClusterUpgradeVersion for the instance
func (instance *MSCluster_NodeSupportedVersion) SetPropertyClusterUpgradeVersion(value uint32) (err error) {
	return instance.SetProperty("ClusterUpgradeVersion", value)
}

// GetClusterUpgradeVersion gets the value of ClusterUpgradeVersion for the instance
func (instance *MSCluster_NodeSupportedVersion) GetPropertyClusterUpgradeVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClusterUpgradeVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
