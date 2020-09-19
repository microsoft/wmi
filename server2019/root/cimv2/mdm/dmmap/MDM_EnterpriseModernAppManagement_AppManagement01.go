// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_EnterpriseModernAppManagement_AppManagement01 struct
type MDM_EnterpriseModernAppManagement_AppManagement01 struct {
	*cim.WmiInstance

	//
	AppInventoryQuery string

	//
	AppInventoryResults string

	//
	InstanceID string

	//
	LastScanError int32

	//
	ParentID string

	//
	RemovePackage string
}

func NewMDM_EnterpriseModernAppManagement_AppManagement01Ex1(instance *cim.WmiInstance) (newInstance *MDM_EnterpriseModernAppManagement_AppManagement01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_EnterpriseModernAppManagement_AppManagement01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_EnterpriseModernAppManagement_AppManagement01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_EnterpriseModernAppManagement_AppManagement01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_EnterpriseModernAppManagement_AppManagement01{
		WmiInstance: tmp,
	}
	return
}

// SetAppInventoryQuery sets the value of AppInventoryQuery for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01) SetPropertyAppInventoryQuery(value string) (err error) {
	return instance.SetProperty("AppInventoryQuery", (value))
}

// GetAppInventoryQuery gets the value of AppInventoryQuery for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01) GetPropertyAppInventoryQuery() (value string, err error) {
	retValue, err := instance.GetProperty("AppInventoryQuery")
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

// SetAppInventoryResults sets the value of AppInventoryResults for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01) SetPropertyAppInventoryResults(value string) (err error) {
	return instance.SetProperty("AppInventoryResults", (value))
}

// GetAppInventoryResults gets the value of AppInventoryResults for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01) GetPropertyAppInventoryResults() (value string, err error) {
	retValue, err := instance.GetProperty("AppInventoryResults")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetLastScanError sets the value of LastScanError for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01) SetPropertyLastScanError(value int32) (err error) {
	return instance.SetProperty("LastScanError", (value))
}

// GetLastScanError gets the value of LastScanError for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01) GetPropertyLastScanError() (value int32, err error) {
	retValue, err := instance.GetProperty("LastScanError")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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

// SetRemovePackage sets the value of RemovePackage for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01) SetPropertyRemovePackage(value string) (err error) {
	return instance.SetProperty("RemovePackage", (value))
}

// GetRemovePackage gets the value of RemovePackage for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01) GetPropertyRemovePackage() (value string, err error) {
	retValue, err := instance.GetProperty("RemovePackage")
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

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01) UpdateScanMethod() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("UpdateScanMethod")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01) RemovePackageMethod( /* IN */ param string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemovePackageMethod", param)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
