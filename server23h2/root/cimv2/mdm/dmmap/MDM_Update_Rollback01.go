// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// MDM_Update_Rollback01 struct
type MDM_Update_Rollback01 struct {
	*cim.WmiInstance

	//
	FeatureUpdateStatus string

	//
	InstanceID string

	//
	ParentID string

	//
	QualityUpdateStatus string
}

func NewMDM_Update_Rollback01Ex1(instance *cim.WmiInstance) (newInstance *MDM_Update_Rollback01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Update_Rollback01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Update_Rollback01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Update_Rollback01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Update_Rollback01{
		WmiInstance: tmp,
	}
	return
}

// SetFeatureUpdateStatus sets the value of FeatureUpdateStatus for the instance
func (instance *MDM_Update_Rollback01) SetPropertyFeatureUpdateStatus(value string) (err error) {
	return instance.SetProperty("FeatureUpdateStatus", (value))
}

// GetFeatureUpdateStatus gets the value of FeatureUpdateStatus for the instance
func (instance *MDM_Update_Rollback01) GetPropertyFeatureUpdateStatus() (value string, err error) {
	retValue, err := instance.GetProperty("FeatureUpdateStatus")
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
func (instance *MDM_Update_Rollback01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Update_Rollback01) GetPropertyInstanceID() (value string, err error) {
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Update_Rollback01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Update_Rollback01) GetPropertyParentID() (value string, err error) {
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

// SetQualityUpdateStatus sets the value of QualityUpdateStatus for the instance
func (instance *MDM_Update_Rollback01) SetPropertyQualityUpdateStatus(value string) (err error) {
	return instance.SetProperty("QualityUpdateStatus", (value))
}

// GetQualityUpdateStatus gets the value of QualityUpdateStatus for the instance
func (instance *MDM_Update_Rollback01) GetPropertyQualityUpdateStatus() (value string, err error) {
	retValue, err := instance.GetProperty("QualityUpdateStatus")
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
func (instance *MDM_Update_Rollback01) QualityUpdateMethod() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("QualityUpdateMethod")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_Update_Rollback01) FeatureUpdateMethod() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("FeatureUpdateMethod")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
