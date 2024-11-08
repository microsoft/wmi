// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package groupset

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	//fcconstant "github.com/microsoft/wmi/pkg/cluster/constant"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	fc "github.com/microsoft/wmi/server2019/root/mscluster"
)

type GroupSet struct {
	*fc.MSCluster_GroupSet
}

// NewGroupSet
func NewGroupSet(instance *wmi.WmiInstance) (*GroupSet, error) {
	wmigpset, err := fc.NewMSCluster_GroupSetEx1(instance)
	if err != nil {
		return nil, err
	}
	return &GroupSet{wmigpset}, nil
}

// GetGroupSet gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetGroupSets(whost *host.WmiHost) (cgpSetcollection GroupSetCollection, err error) {
	query := query.NewWmiQuery("MSCluster_GroupSet")
	instances, err := instance.GetWmiInstancesFromHost(whost, string(constant.FailoverCluster), query)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			instances.Close()
		}
	}()

	cgpSetcollection, err = NewGroupSetCollection(instances)
	return
}

// GetGroupSet gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetGroupSet(whost *host.WmiHost, grpSetName string) (cgpSet *GroupSet, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("MSCluster_GroupSet", "Name", grpSetName)
	wmigpset, err := fc.NewMSCluster_GroupSetEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	cgpSet = &GroupSet{wmigpset}
	return
}
