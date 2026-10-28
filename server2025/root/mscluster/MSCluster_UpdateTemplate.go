// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
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

// MSCluster_UpdateTemplate struct
type MSCluster_UpdateTemplate struct {
	*cim.WmiInstance

	//
	CoordinatorOptions string

	//
	DrainPluginOptions string

	//
	DrainPlugins []string

	//
	Name string

	//
	Reason string

	//
	UpdaterPluginOptions string

	//
	UpdaterPlugins []string

	//
	ValidatorPluginOptions string

	//
	ValidatorPlugins []string
}

func NewMSCluster_UpdateTemplateEx1(instance *cim.WmiInstance) (newInstance *MSCluster_UpdateTemplate, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_UpdateTemplate{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_UpdateTemplateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_UpdateTemplate, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_UpdateTemplate{
		WmiInstance: tmp,
	}
	return
}

// SetCoordinatorOptions sets the value of CoordinatorOptions for the instance
func (instance *MSCluster_UpdateTemplate) SetPropertyCoordinatorOptions(value string) (err error) {
	return instance.SetProperty("CoordinatorOptions", (value))
}

// GetCoordinatorOptions gets the value of CoordinatorOptions for the instance
func (instance *MSCluster_UpdateTemplate) GetPropertyCoordinatorOptions() (value string, err error) {
	retValue, err := instance.GetProperty("CoordinatorOptions")
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

// SetDrainPluginOptions sets the value of DrainPluginOptions for the instance
func (instance *MSCluster_UpdateTemplate) SetPropertyDrainPluginOptions(value string) (err error) {
	return instance.SetProperty("DrainPluginOptions", (value))
}

// GetDrainPluginOptions gets the value of DrainPluginOptions for the instance
func (instance *MSCluster_UpdateTemplate) GetPropertyDrainPluginOptions() (value string, err error) {
	retValue, err := instance.GetProperty("DrainPluginOptions")
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

// SetDrainPlugins sets the value of DrainPlugins for the instance
func (instance *MSCluster_UpdateTemplate) SetPropertyDrainPlugins(value []string) (err error) {
	return instance.SetProperty("DrainPlugins", (value))
}

// GetDrainPlugins gets the value of DrainPlugins for the instance
func (instance *MSCluster_UpdateTemplate) GetPropertyDrainPlugins() (value []string, err error) {
	retValue, err := instance.GetProperty("DrainPlugins")
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
func (instance *MSCluster_UpdateTemplate) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSCluster_UpdateTemplate) GetPropertyName() (value string, err error) {
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

// SetReason sets the value of Reason for the instance
func (instance *MSCluster_UpdateTemplate) SetPropertyReason(value string) (err error) {
	return instance.SetProperty("Reason", (value))
}

// GetReason gets the value of Reason for the instance
func (instance *MSCluster_UpdateTemplate) GetPropertyReason() (value string, err error) {
	retValue, err := instance.GetProperty("Reason")
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

// SetUpdaterPluginOptions sets the value of UpdaterPluginOptions for the instance
func (instance *MSCluster_UpdateTemplate) SetPropertyUpdaterPluginOptions(value string) (err error) {
	return instance.SetProperty("UpdaterPluginOptions", (value))
}

// GetUpdaterPluginOptions gets the value of UpdaterPluginOptions for the instance
func (instance *MSCluster_UpdateTemplate) GetPropertyUpdaterPluginOptions() (value string, err error) {
	retValue, err := instance.GetProperty("UpdaterPluginOptions")
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

// SetUpdaterPlugins sets the value of UpdaterPlugins for the instance
func (instance *MSCluster_UpdateTemplate) SetPropertyUpdaterPlugins(value []string) (err error) {
	return instance.SetProperty("UpdaterPlugins", (value))
}

// GetUpdaterPlugins gets the value of UpdaterPlugins for the instance
func (instance *MSCluster_UpdateTemplate) GetPropertyUpdaterPlugins() (value []string, err error) {
	retValue, err := instance.GetProperty("UpdaterPlugins")
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

// SetValidatorPluginOptions sets the value of ValidatorPluginOptions for the instance
func (instance *MSCluster_UpdateTemplate) SetPropertyValidatorPluginOptions(value string) (err error) {
	return instance.SetProperty("ValidatorPluginOptions", (value))
}

// GetValidatorPluginOptions gets the value of ValidatorPluginOptions for the instance
func (instance *MSCluster_UpdateTemplate) GetPropertyValidatorPluginOptions() (value string, err error) {
	retValue, err := instance.GetProperty("ValidatorPluginOptions")
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

// SetValidatorPlugins sets the value of ValidatorPlugins for the instance
func (instance *MSCluster_UpdateTemplate) SetPropertyValidatorPlugins(value []string) (err error) {
	return instance.SetProperty("ValidatorPlugins", (value))
}

// GetValidatorPlugins gets the value of ValidatorPlugins for the instance
func (instance *MSCluster_UpdateTemplate) GetPropertyValidatorPlugins() (value []string, err error) {
	retValue, err := instance.GetProperty("ValidatorPlugins")
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

//

// <param name="CoordinatorOptions" type="string "></param>
// <param name="DrainPluginOptions" type="string "></param>
// <param name="DrainPlugins" type="string []"></param>
// <param name="Name" type="string "></param>
// <param name="Reason" type="string "></param>
// <param name="UpdaterPluginOptions" type="string "></param>
// <param name="UpdaterPlugins" type="string []"></param>
// <param name="ValidatorPluginOptions" type="string "></param>
// <param name="ValidatorPlugins" type="string []"></param>

// <param name="CreatedUpdateTemplate" type="MSCluster_UpdateTemplate "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_UpdateTemplate) CreateUpdateTemplate( /* IN */ Name string,
	/* IN */ CoordinatorOptions string,
	/* IN */ UpdaterPlugins []string,
	/* IN */ UpdaterPluginOptions string,
	/* IN */ ValidatorPlugins []string,
	/* IN */ ValidatorPluginOptions string,
	/* IN */ DrainPlugins []string,
	/* IN */ DrainPluginOptions string,
	/* IN */ Reason string,
	/* OUT */ CreatedUpdateTemplate MSCluster_UpdateTemplate) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateUpdateTemplate", Name, CoordinatorOptions, UpdaterPlugins, UpdaterPluginOptions, ValidatorPlugins, ValidatorPluginOptions, DrainPlugins, DrainPluginOptions, Reason)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CoordinatorOptions" type="string "></param>
// <param name="DrainPluginOptions" type="string "></param>
// <param name="DrainPlugins" type="string []"></param>
// <param name="NewName" type="string "></param>
// <param name="Reason" type="string "></param>
// <param name="UpdaterPluginOptions" type="string "></param>
// <param name="UpdaterPlugins" type="string []"></param>
// <param name="ValidatorPluginOptions" type="string "></param>
// <param name="ValidatorPlugins" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_UpdateTemplate) SetUpdateTemplate( /* IN */ NewName string,
	/* IN */ CoordinatorOptions string,
	/* IN */ UpdaterPlugins []string,
	/* IN */ UpdaterPluginOptions string,
	/* IN */ ValidatorPlugins []string,
	/* IN */ ValidatorPluginOptions string,
	/* IN */ DrainPlugins []string,
	/* IN */ DrainPluginOptions string,
	/* IN */ Reason string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetUpdateTemplate", NewName, CoordinatorOptions, UpdaterPlugins, UpdaterPluginOptions, ValidatorPlugins, ValidatorPluginOptions, DrainPlugins, DrainPluginOptions, Reason)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_UpdateTemplate) RemoveUpdateTemplate() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveUpdateTemplate")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
