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

// MSCluster_UpdateRun struct
type MSCluster_UpdateRun struct {
	*cim.WmiInstance

	//
	Name string

	//
	NodeResult []string

	//
	RunEndTime string

	//
	RunParameters MSCluster_UpdateTemplate

	//
	RunStartTime string

	//
	StartReason string

	//
	Status string

	//
	StopReason string
}

func NewMSCluster_UpdateRunEx1(instance *cim.WmiInstance) (newInstance *MSCluster_UpdateRun, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_UpdateRun{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_UpdateRunEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_UpdateRun, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_UpdateRun{
		WmiInstance: tmp,
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSCluster_UpdateRun) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSCluster_UpdateRun) GetPropertyName() (value string, err error) {
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

// SetNodeResult sets the value of NodeResult for the instance
func (instance *MSCluster_UpdateRun) SetPropertyNodeResult(value []string) (err error) {
	return instance.SetProperty("NodeResult", (value))
}

// GetNodeResult gets the value of NodeResult for the instance
func (instance *MSCluster_UpdateRun) GetPropertyNodeResult() (value []string, err error) {
	retValue, err := instance.GetProperty("NodeResult")
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

// SetRunEndTime sets the value of RunEndTime for the instance
func (instance *MSCluster_UpdateRun) SetPropertyRunEndTime(value string) (err error) {
	return instance.SetProperty("RunEndTime", (value))
}

// GetRunEndTime gets the value of RunEndTime for the instance
func (instance *MSCluster_UpdateRun) GetPropertyRunEndTime() (value string, err error) {
	retValue, err := instance.GetProperty("RunEndTime")
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

// SetRunParameters sets the value of RunParameters for the instance
func (instance *MSCluster_UpdateRun) SetPropertyRunParameters(value MSCluster_UpdateTemplate) (err error) {
	return instance.SetProperty("RunParameters", (value))
}

// GetRunParameters gets the value of RunParameters for the instance
func (instance *MSCluster_UpdateRun) GetPropertyRunParameters() (value MSCluster_UpdateTemplate, err error) {
	retValue, err := instance.GetProperty("RunParameters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSCluster_UpdateTemplate)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSCluster_UpdateTemplate is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSCluster_UpdateTemplate(valuetmp)

	return
}

// SetRunStartTime sets the value of RunStartTime for the instance
func (instance *MSCluster_UpdateRun) SetPropertyRunStartTime(value string) (err error) {
	return instance.SetProperty("RunStartTime", (value))
}

// GetRunStartTime gets the value of RunStartTime for the instance
func (instance *MSCluster_UpdateRun) GetPropertyRunStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("RunStartTime")
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

// SetStartReason sets the value of StartReason for the instance
func (instance *MSCluster_UpdateRun) SetPropertyStartReason(value string) (err error) {
	return instance.SetProperty("StartReason", (value))
}

// GetStartReason gets the value of StartReason for the instance
func (instance *MSCluster_UpdateRun) GetPropertyStartReason() (value string, err error) {
	retValue, err := instance.GetProperty("StartReason")
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

// SetStatus sets the value of Status for the instance
func (instance *MSCluster_UpdateRun) SetPropertyStatus(value string) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MSCluster_UpdateRun) GetPropertyStatus() (value string, err error) {
	retValue, err := instance.GetProperty("Status")
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

// SetStopReason sets the value of StopReason for the instance
func (instance *MSCluster_UpdateRun) SetPropertyStopReason(value string) (err error) {
	return instance.SetProperty("StopReason", (value))
}

// GetStopReason gets the value of StopReason for the instance
func (instance *MSCluster_UpdateRun) GetPropertyStopReason() (value string, err error) {
	retValue, err := instance.GetProperty("StopReason")
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

// <param name="CoordinatorOptions" type="string "></param>
// <param name="DrainPluginOptions" type="string "></param>
// <param name="DrainPlugins" type="string []"></param>
// <param name="Name" type="string "></param>
// <param name="Reason" type="string "></param>
// <param name="TemplateName" type="string "></param>
// <param name="UpdaterPluginOptions" type="string "></param>
// <param name="UpdaterPlugins" type="string []"></param>
// <param name="ValidatorPluginOptions" type="string "></param>
// <param name="ValidatorPlugins" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="StartedUpdateRun" type="MSCluster_UpdateRun "></param>
func (instance *MSCluster_UpdateRun) StartUpdateRun( /* IN */ Name string,
	/* IN */ TemplateName string,
	/* IN */ CoordinatorOptions string,
	/* IN */ UpdaterPlugins []string,
	/* IN */ UpdaterPluginOptions string,
	/* IN */ ValidatorPlugins []string,
	/* IN */ ValidatorPluginOptions string,
	/* IN */ DrainPlugins []string,
	/* IN */ DrainPluginOptions string,
	/* IN */ Reason string,
	/* OUT */ StartedUpdateRun MSCluster_UpdateRun) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("StartUpdateRun", Name, TemplateName, CoordinatorOptions, UpdaterPlugins, UpdaterPluginOptions, ValidatorPlugins, ValidatorPluginOptions, DrainPlugins, DrainPluginOptions, Reason)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Reason" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_UpdateRun) StopUpdateRun( /* IN */ Reason string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StopUpdateRun", Reason)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_UpdateRun) RemoveUpdateRun() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveUpdateRun")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
