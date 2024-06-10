// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.MSCluster
//
// ////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_Service struct
type MSCluster_Service struct {
	*CIM_ClusteringService

	//
	NodeHighestVersion uint32

	//
	NodeLowestVersion uint32

	//
	State string
}

func NewMSCluster_ServiceEx1(instance *cim.WmiInstance) (newInstance *MSCluster_Service, err error) {
	tmp, err := NewCIM_ClusteringServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_Service{
		CIM_ClusteringService: tmp,
	}
	return
}

func NewMSCluster_ServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_Service, err error) {
	tmp, err := NewCIM_ClusteringServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_Service{
		CIM_ClusteringService: tmp,
	}
	return
}

// SetNodeHighestVersion sets the value of NodeHighestVersion for the instance
func (instance *MSCluster_Service) SetPropertyNodeHighestVersion(value uint32) (err error) {
	return instance.SetProperty("NodeHighestVersion", (value))
}

// GetNodeHighestVersion gets the value of NodeHighestVersion for the instance
func (instance *MSCluster_Service) GetPropertyNodeHighestVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("NodeHighestVersion")
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

// SetNodeLowestVersion sets the value of NodeLowestVersion for the instance
func (instance *MSCluster_Service) SetPropertyNodeLowestVersion(value uint32) (err error) {
	return instance.SetProperty("NodeLowestVersion", (value))
}

// GetNodeLowestVersion gets the value of NodeLowestVersion for the instance
func (instance *MSCluster_Service) GetPropertyNodeLowestVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("NodeLowestVersion")
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

// SetState sets the value of State for the instance
func (instance *MSCluster_Service) SetPropertyState(value string) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *MSCluster_Service) GetPropertyState() (value string, err error) {
	retValue, err := instance.GetProperty("State")
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

func (instance *MSCluster_Service) Start() (err error) {
	_, err = instance.InvokeMethodWithReturn("Start")
	if err != nil {
		return
	}
	return

}

func (instance *MSCluster_Service) Stop() (err error) {
	_, err = instance.InvokeMethodWithReturn("Stop")
	if err != nil {
		return
	}
	return

}
