// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Update_Rollback01 struct
type MDM_Update_Rollback01 struct {
	cim.WmiInstance

	//
	FeatureUpdateStatus string

	//
	InstanceID string

	//
	ParentID string

	//
	QualityUpdateStatus string
}

// SetFeatureUpdateStatus sets the value of FeatureUpdateStatus for the instance
func (instance *MDM_Update_Rollback01) SetPropertyFeatureUpdateStatus(value string) (err error) {
	return instance.SetProperty("FeatureUpdateStatus", value)
}

// GetFeatureUpdateStatus gets the value of FeatureUpdateStatus for the instance
func (instance *MDM_Update_Rollback01) GetPropertyFeatureUpdateStatus() (value string, err error) {
	retValue, err := instance.GetProperty("FeatureUpdateStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Update_Rollback01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Update_Rollback01) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Update_Rollback01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Update_Rollback01) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQualityUpdateStatus sets the value of QualityUpdateStatus for the instance
func (instance *MDM_Update_Rollback01) SetPropertyQualityUpdateStatus(value string) (err error) {
	return instance.SetProperty("QualityUpdateStatus", value)
}

// GetQualityUpdateStatus gets the value of QualityUpdateStatus for the instance
func (instance *MDM_Update_Rollback01) GetPropertyQualityUpdateStatus() (value string, err error) {
	retValue, err := instance.GetProperty("QualityUpdateStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
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
