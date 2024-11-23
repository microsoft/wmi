// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package affinityrule

import (
	"strings"
	"time"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	fc "github.com/microsoft/wmi/server2019/root/mscluster"
)

type AffinityRule struct {
	*fc.MSCluster_AffinityRule
}

type FailoverClusterAffinityRuleType int

const (
	SameFaultDomain      FailoverClusterAffinityRuleType = 1
	SameNode             FailoverClusterAffinityRuleType = 2
	DifferentFaultDomain FailoverClusterAffinityRuleType = 3
	DifferentNode        FailoverClusterAffinityRuleType = 4
)

// NewAffinityRule
func NewAffinityRule(instance *wmi.WmiInstance) (*AffinityRule, error) {
	wmiafRule, err := fc.NewMSCluster_AffinityRuleEx1(instance)
	if err != nil {
		return nil, err
	}
	return &AffinityRule{wmiafRule}, nil
}

// CreateAffinityRule
func CreateAffinityRule(whost *host.WmiHost, name string, ruleType int, strict bool) (affinityRule *AffinityRule, err error) {
	query := "SELECT * FROM meta_class WHERE __CLASS = 'MSCluster_AffinityRule'"
	classes, err := instance.GetWmiClasssesFromHostRawQuery(whost, string(constant.FailoverCluster), query)
	if err != nil {
		return
	}

	defer classes.Close()

	if len(classes) > 1 {
		err = errors.Wrapf(errors.Unknown, "More than one MSCluster_AffinityRule class found, unexpected error")
		return
	}
	arClass := classes[0]
	_, err = arClass.InvokeMethod("CreateAffinityRule", name, int(ruleType), nil)
	if err != nil {
		return nil, err
	}

	// added layer of protection, GetClusterAffinityRule can return not found immediately after creating the affinity rule
	// if this happens, we will retry the get operation a few times before returning an error
	maxAttempts := 5
	for i := 1; i <= maxAttempts; i++ {
		affinityRule, err = GetAffinityRule(whost, name)
		if err != nil {
			if errors.IsNotFound(err) && i < maxAttempts {
				time.Sleep(2 * time.Second)
				continue
			}
			return
		}
		break
	}
	defer func() {
		if err != nil {
			affinityRule.RemoveAffinityRule()
			affinityRule.Close()
		}
	}()

	isAntiAffinityRule := (ruleType == int(DifferentFaultDomain) || ruleType == int(DifferentNode))
	if !strict && isAntiAffinityRule && isSoftAntiAffinitySupported(affinityRule) {
		_, err = affinityRule.InvokeMethod("SetAffinityRule", int(ruleType), 1 /* Enabled */, 1 /* SoftAntiAffinity */)
		if err != nil {
			return nil, err
		}
	}

	return
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

func isSoftAntiAffinitySupported(affinityRule *AffinityRule) bool {
	propeties := affinityRule.GetClass().GetPropertiesNames()
	for _, property := range propeties {
		if strings.EqualFold(property, "SoftAntiAffinity") {
			return true
		}
	}
	return false
}
