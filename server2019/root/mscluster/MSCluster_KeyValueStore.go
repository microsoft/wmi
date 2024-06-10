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
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_KeyValueStore struct
type MSCluster_KeyValueStore struct {
	*cim.WmiInstance

	//
	Manager string

	//
	Name string

	//
	Type uint32
}

func NewMSCluster_KeyValueStoreEx1(instance *cim.WmiInstance) (newInstance *MSCluster_KeyValueStore, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_KeyValueStore{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_KeyValueStoreEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_KeyValueStore, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_KeyValueStore{
		WmiInstance: tmp,
	}
	return
}

// SetManager sets the value of Manager for the instance
func (instance *MSCluster_KeyValueStore) SetPropertyManager(value string) (err error) {
	return instance.SetProperty("Manager", (value))
}

// GetManager gets the value of Manager for the instance
func (instance *MSCluster_KeyValueStore) GetPropertyManager() (value string, err error) {
	retValue, err := instance.GetProperty("Manager")
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

// SetName sets the value of Name for the instance
func (instance *MSCluster_KeyValueStore) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSCluster_KeyValueStore) GetPropertyName() (value string, err error) {
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

// SetType sets the value of Type for the instance
func (instance *MSCluster_KeyValueStore) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSCluster_KeyValueStore) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
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
