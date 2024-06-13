// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_AffinityRule struct
type MSCluster_AffinityRule struct {
	*cim.WmiInstance

	//
	Enabled uint32

	//
	Groups []string

	//
	Name string

	//
	RuleType uint32
}

func NewMSCluster_AffinityRuleEx1(instance *cim.WmiInstance) (newInstance *MSCluster_AffinityRule, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_AffinityRule{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_AffinityRuleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_AffinityRule, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_AffinityRule{
		WmiInstance: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MSCluster_AffinityRule) SetPropertyEnabled(value uint32) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSCluster_AffinityRule) GetPropertyEnabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetGroups sets the value of Groups for the instance
func (instance *MSCluster_AffinityRule) SetPropertyGroups(value []string) (err error) {
	return instance.SetProperty("Groups", (value))
}

// GetGroups gets the value of Groups for the instance
func (instance *MSCluster_AffinityRule) GetPropertyGroups() (value []string, err error) {
	retValue, err := instance.GetProperty("Groups")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetName sets the value of Name for the instance
func (instance *MSCluster_AffinityRule) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSCluster_AffinityRule) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetRuleType sets the value of RuleType for the instance
func (instance *MSCluster_AffinityRule) SetPropertyRuleType(value uint32) (err error) {
	return instance.SetProperty("RuleType", (value))
}

// GetRuleType gets the value of RuleType for the instance
func (instance *MSCluster_AffinityRule) GetPropertyRuleType() (value uint32, err error) {
	retValue, err := instance.GetProperty("RuleType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

//

// <param name="Name" type="string "></param>
// <param name="RuleType" type="uint32 "></param>

// <param name="CreatedAffinityRule" type="MSCluster_AffinityRule "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_AffinityRule) CreateAffinityRule( /* IN */ Name string,
	/* IN */ RuleType uint32,
	/* OUT */ CreatedAffinityRule MSCluster_AffinityRule) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateAffinityRule", Name, RuleType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Enabled" type="uint32 "></param>
// <param name="RuleType" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_AffinityRule) SetAffinityRule( /* IN */ RuleType uint32,
	/* IN */ Enabled uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetAffinityRule", RuleType, Enabled)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Groups" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_AffinityRule) AddGroupToAffinityRule( /* IN */ Groups []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddGroupToAffinityRule", Groups)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Groups" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_AffinityRule) RemoveGroupFromAffinityRule( /* IN */ Groups []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveGroupFromAffinityRule", Groups)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ClusterSharedVolumes" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_AffinityRule) AddClusterSharedVolumeToAffinityRule( /* IN */ ClusterSharedVolumes []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddClusterSharedVolumeToAffinityRule", ClusterSharedVolumes)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ClusterSharedVolumes" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_AffinityRule) RemoveClusterSharedVolumeFromAffinityRule( /* IN */ ClusterSharedVolumes []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveClusterSharedVolumeFromAffinityRule", ClusterSharedVolumes)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_AffinityRule) RemoveAffinityRule() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveAffinityRule")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
