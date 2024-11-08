// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package clustersharedvolume

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	fcconstant "github.com/microsoft/wmi/pkg/cluster/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	fc "github.com/microsoft/wmi/server2019/root/mscluster"
)

type ClusterSharedVolume struct {
	*fc.MSCluster_ClusterSharedVolume
}

// NewClusterSharedVolume
func NewClusterSharedVolume(instance *wmi.WmiInstance) (*ClusterSharedVolume, error) {
	wmivm, err := fc.NewMSCluster_ClusterSharedVolumeEx1(instance)
	if err != nil {
		return nil, err
	}
	return &ClusterSharedVolume{wmivm}, nil
}

// GetClusterSharedVolume gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetClusterSharedVolumes(whost *host.WmiHost) (cvolumecollection ClusterSharedVolumeCollection, err error) {
	query := query.NewWmiQuery("MSCluster_ClusterSharedVolume")
	instances, err := instance.GetWmiInstancesFromHost(whost, string(constant.FailoverCluster), query)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			instances.Close()
		}
	}()

	cvolumecollection, err = NewClusterSharedVolumeCollection(instances)
	return
}

// GetClusterSharedVolume gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetClusterSharedVolume(whost *host.WmiHost, volumeName string) (cvolume *ClusterSharedVolume, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("MSCluster_ClusterSharedVolume", "Name", volumeName)
	volinstance, err := fc.NewMSCluster_ClusterSharedVolumeEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	cvolume = &ClusterSharedVolume{volinstance}
	return
}

// IsStatusOK get the cluster health status
func (c *ClusterSharedVolume) IsStatusOK() (status bool) {
	state, err := c.GetPropertyStatus()
	if err != nil {
		return
	}

	return (state == fcconstant.STATUS_OK)
}

// IsFaultStateOK get the cluster health status
func (c *ClusterSharedVolume) IsFaultStateOK()  (status bool) {
	state, err := c.GetPropertyFaultState()
	if err != nil {
		return
	}

	return (int32(state) == (fcconstant.FAULT_STATE_NO_FAULT))
}

// ContainsPath checks if the input path is part of this cluster shared volume
func (c *ClusterSharedVolume) ContainsPath(absolutePath string) (status bool) {
	// TBD
	return
}
