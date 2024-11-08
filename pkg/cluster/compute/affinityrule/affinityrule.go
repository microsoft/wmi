// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package affinityrule

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	fc "github.com/microsoft/wmi/server2019/root/mscluster"
)

type AffinityRule struct {
	*fc.MSCluster_AffinityRule
}

// NewAffinityRule
func NewAffinityRule(instance *wmi.WmiInstance) (*AffinityRule, error) {
	wmiafRule, err := fc.NewMSCluster_AffinityRuleEx1(instance)
	if err != nil {
		return nil, err
	}
	return &AffinityRule{wmiafRule}, nil
}

// GetAffinityRule gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetAffinityRules(whost *host.WmiHost) (caffinityRulecollection AffinityRuleCollection, err error) {
	query := query.NewWmiQuery("MSCluster_AffinityRule")
	instances, err := instance.GetWmiInstancesFromHost(whost, string(constant.FailoverCluster), query)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			instances.Close()
		}
	}()

	caffinityRulecollection, err = NewAffinityRuleCollection(instances)
	return
}

// GetAffinityRule gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetAffinityRule(whost *host.WmiHost, affinityRuleName string) (caffinityRule *AffinityRule, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("MSCluster_AffinityRule", "Name", affinityRuleName)
	wmiafRule, err := fc.NewMSCluster_AffinityRuleEx6(whost.HostName, string(constant.FailoverCluster), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	caffinityRule = &AffinityRule{wmiafRule}
	return
}
