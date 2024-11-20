// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package clustersharedvolume

import (
	"fmt"
	"path/filepath"
	"strings"

	"reflect"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	fcconstant "github.com/microsoft/wmi/pkg/cluster/constant"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"

	"github.com/microsoft/wmi/pkg/cluster/compute/resource"
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

// GetClusterSharedVolume gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetClusterSharedVolumebyVolumeName(whost *host.WmiHost, volumeName string) (cvolume *ClusterSharedVolume, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("MSCluster_ClusterSharedVolume", "VolumeName", volumeName)
	volinstance, err := fc.NewMSCluster_ClusterSharedVolumeEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	cvolume = &ClusterSharedVolume{volinstance}
	return
}

// GetClusterSharedVolumebyName gets an name of the cluster shared volume
// Make sure to call Close once done using this instance
// Sometimes the volume name is difficult to compare
func GetClusterSharedVolumebyName(whost *host.WmiHost, name string) (cvolume *ClusterSharedVolume, err error) {
	csvInstances, err := GetClusterSharedVolumes(whost)
	if err != nil {
		return
	}
	defer csvInstances.Close()
	inPath := strings.ToLower(filepath.Clean(name))
	for _, instance := range csvInstances {
		instanceName, err1 := instance.GetPropertyName()
		if err1 != nil {
			err = err1
			return
		}
		matchingPath := strings.ToLower(filepath.Clean(instanceName))
		// Append "\\" to eliminate false positive due to partial match
		// e.g. "C:\ClusterStorage\Volume10\image" and "C:\ClusterStorage\Volume1"
		if strings.HasPrefix(inPath+"\\", matchingPath+"\\") {
			tmpInstance, err1 := instance.Clone()
			if err1 != nil {
				err = err1
				return
			}
			return NewClusterSharedVolume(tmpInstance)
		}
	}

	return
}

// IsStatusOK get the cluster health status
func (c *ClusterSharedVolume) IsStatusOK() (status bool, err error) {
	statusValue, err := c.GetPropertyStatus()
	if err != nil {
		return
	}

	status = (len(statusValue) == 0 || (statusValue == fcconstant.STATUS_OK))
	return
}

// GetFaultState gets the value of FaultState for the instance
func (c *ClusterSharedVolume) GetFaultState() (value int32, err error) {
	retValue, err := c.GetProperty("FaultState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	value, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	return
}

// IsFaultStateOK get the cluster health status
func (c *ClusterSharedVolume) IsFaultStateOK() (status bool, err error) {
	faultstate, err := c.GetFaultState()
	if err != nil {
		return
	}

	status = (faultstate == (fcconstant.FAULT_STATE_NO_FAULT))
	return
}

// ContainsPath checks if the input path is part of this cluster shared volume
func (c *ClusterSharedVolume) ContainsPath(absolutePath string) (status bool, err error) {
	volName, err := c.GetPropertyName()
	if err != nil {
		return
	}

	if strings.Contains(strings.ToLower(absolutePath), strings.ToLower(volName)) {
		status = true
		return
	}

	return
}

func (c *ClusterSharedVolume) GetAssociatedClusterResource() (clusterResource *resource.Resource, err error) {
	volumeName, err := c.GetPropertyName()
	if err != nil {
		return
	}

	volumeAssocQuery := fmt.Sprintf("Associators of {MSCluster_ClusterSharedVolume='%s'} where assocclass=MSCluster_ClusterSharedVolumeToResource", volumeName)
	assocInstances, err := instance.GetWmiInstancesFromHostRawQuery(c.GetWmiHost(), string(constant.FailoverCluster), volumeAssocQuery)
	if err != nil {
		return
	}
	defer assocInstances.Close()
	if len(assocInstances) == 0 {
		err = errors.Wrapf(errors.NotFound, "No Cluster Resource found for Cluster Shared Volume %s", volumeName)
		return
	}

	clonedResource, err := assocInstances[0].Clone()
	if err != nil {
		return
	}
	clusterResource, err = resource.NewResource(clonedResource)
	return
}

// OwnerGroup checks if the input path is part of this cluster shared volume
func (c *ClusterSharedVolume) OwnerGroup() (ownerGroup string, err error) {
	cr, err := c.GetAssociatedClusterResource()
	if err != nil {
		return
	}
	defer cr.Close()

	ownerGroup, err = cr.GetPropertyOwnerGroup()
	if err != nil {
		return
	}
	return
}
