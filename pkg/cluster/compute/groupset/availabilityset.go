// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package groupset

import (
	"time"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	//fcconstant "github.com/microsoft/wmi/pkg/cluster/constant"
	"github.com/microsoft/wmi/pkg/constant"
	fc "github.com/microsoft/wmi/server2019/root/mscluster"
)

// GetAvailabilitySet gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetAvailabilitySet(whost *host.WmiHost, grpSetName string) (cgpSet *GroupSet, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("MSCluster_GroupSet", "Name", grpSetName, "IsAvailabilitySet", "TRUE")
	wmigpset, err := fc.NewMSCluster_GroupSetEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	cgpSet = &GroupSet{wmigpset}
	return
}

// CreateAffinityRule
func CreateAvailabilitySet(whost *host.WmiHost, name string, updateDomainCount int) (availabilitySet *GroupSet, err error) {
	query := "SELECT * FROM meta_class WHERE __CLASS = 'MSCluster_GroupSet'"
	classes, err := instance.GetWmiClasssesFromHostRawQuery(whost, string(constant.FailoverCluster), query)
	if err != nil {
		return
	}
	defer classes.Close()

	if len(classes) > 1 {
		err = errors.Wrapf(errors.Unknown, "More than one MSCluster_GroupSet class found, unexpected error")
		return
	}
	gpClass := classes[0]
	_, err = gpClass.InvokeMethod("CreateAvailabilitySet", name, []string{} /* Groups */, updateDomainCount, 1 /* fault domain count */, false /* reserve spare node */)
	if err != nil {
		return nil, err
	}

	// added layer of protection, GetClusterAffinityRule can return not found immediately after creating the affinity rule
	// if this happens, we will retry the get operation a few times before returning an error
	maxAttempts := 5
	for i := 0; i < maxAttempts; i++ {
		availabilitySet, err := GetAvailabilitySet(whost, name)
		if err != nil {
			if errors.IsNotFound(err) {
				time.Sleep(2 * time.Second)
				continue
			}
			return nil, err
		}
		return availabilitySet, nil
	}

	return
}
