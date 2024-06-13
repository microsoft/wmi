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

// MSCluster_KeyValueStoreManager struct
type MSCluster_KeyValueStoreManager struct {
	*cim.WmiInstance

	//
	Name string

	//
	Path string
}

func NewMSCluster_KeyValueStoreManagerEx1(instance *cim.WmiInstance) (newInstance *MSCluster_KeyValueStoreManager, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_KeyValueStoreManager{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_KeyValueStoreManagerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_KeyValueStoreManager, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_KeyValueStoreManager{
		WmiInstance: tmp,
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSCluster_KeyValueStoreManager) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSCluster_KeyValueStoreManager) GetPropertyName() (value string, err error) {
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

// SetPath sets the value of Path for the instance
func (instance *MSCluster_KeyValueStoreManager) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *MSCluster_KeyValueStoreManager) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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

//

// <param name="Flags" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="Path" type="string "></param>

// <param name="CreatedKeyValueStoreManager" type="MSCluster_KeyValueStoreManager "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_KeyValueStoreManager) CreateKeyValueStoreManager( /* IN */ Name string,
	/* IN */ Path string,
	/* IN */ Flags uint32,
	/* OUT */ CreatedKeyValueStoreManager MSCluster_KeyValueStoreManager) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateKeyValueStoreManager", Name, Path, Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_KeyValueStoreManager) DeleteKeyValueStoreManager( /* IN */ Name string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DeleteKeyValueStoreManager", Name, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>

// <param name="KeyValueStoreManagers" type="MSCluster_KeyValueStoreManager []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_KeyValueStoreManager) GetKeyValueStoreManagers( /* IN */ Flags uint32,
	/* OUT */ KeyValueStoreManagers []MSCluster_KeyValueStoreManager) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetKeyValueStoreManagers", Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_KeyValueStoreManager) CreateKeyValueStore( /* IN */ Name string,
	/* IN */ Type uint32,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CreateKeyValueStore", Name, Type, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_KeyValueStoreManager) DeleteKeyValueStore( /* IN */ Name string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DeleteKeyValueStore", Name, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>

// <param name="KeyValueStores" type="MSCluster_KeyValueStore []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_KeyValueStoreManager) GetKeyValueStores( /* IN */ Flags uint32,
	/* OUT */ KeyValueStores []MSCluster_KeyValueStore) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetKeyValueStores", Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Flags" type="uint32 "></param>

// <param name="KeyValueStores" type="MSCluster_KeyValueStore []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_KeyValueStoreManager) GetAllKeyValueStores( /* IN */ Flags uint32,
	/* OUT */ KeyValueStores []MSCluster_KeyValueStore) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetAllKeyValueStores", Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
