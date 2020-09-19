// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_Cluster struct
type CIM_Cluster struct {
	*CIM_ComputerSystem

	//
	MaxNumberOfNodes uint32
}

func NewCIM_ClusterEx1(instance *cim.WmiInstance) (newInstance *CIM_Cluster, err error) {
	tmp, err := NewCIM_ComputerSystemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Cluster{
		CIM_ComputerSystem: tmp,
	}
	return
}

func NewCIM_ClusterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Cluster, err error) {
	tmp, err := NewCIM_ComputerSystemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Cluster{
		CIM_ComputerSystem: tmp,
	}
	return
}

// SetMaxNumberOfNodes sets the value of MaxNumberOfNodes for the instance
func (instance *CIM_Cluster) SetPropertyMaxNumberOfNodes(value uint32) (err error) {
	return instance.SetProperty("MaxNumberOfNodes", (value))
}

// GetMaxNumberOfNodes gets the value of MaxNumberOfNodes for the instance
func (instance *CIM_Cluster) GetPropertyMaxNumberOfNodes() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxNumberOfNodes")
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
