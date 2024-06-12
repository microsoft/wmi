// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_CollectionReferencePointExportJob struct
type Msvm_CollectionReferencePointExportJob struct {
	*CIM_ConcreteJob

	//
	BaseReferencePointGroupId string

	//
	Cancellable bool

	//
	CollectionId string

	//
	ErrorSummaryDescription string

	//
	ExportDirectory string

	//
	ExportedConfigFilePaths []string

	//
	ExportedDisks []string

	//
	ExportedGuestStateFilePaths []string

	//
	ExportedLogFilePaths []string

	//
	ExportedRuntimeFilePaths []string

	//
	ReferencePointGroupId string

	//
	VirtualMachineId []string
}

func NewMsvm_CollectionReferencePointExportJobEx1(instance *cim.WmiInstance) (newInstance *Msvm_CollectionReferencePointExportJob, err error) {
	tmp, err := NewCIM_ConcreteJobEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectionReferencePointExportJob{
		CIM_ConcreteJob: tmp,
	}
	return
}

func NewMsvm_CollectionReferencePointExportJobEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_CollectionReferencePointExportJob, err error) {
	tmp, err := NewCIM_ConcreteJobEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectionReferencePointExportJob{
		CIM_ConcreteJob: tmp,
	}
	return
}

// SetBaseReferencePointGroupId sets the value of BaseReferencePointGroupId for the instance
func (instance *Msvm_CollectionReferencePointExportJob) SetPropertyBaseReferencePointGroupId(value string) (err error) {
	return instance.SetProperty("BaseReferencePointGroupId", (value))
}

// GetBaseReferencePointGroupId gets the value of BaseReferencePointGroupId for the instance
func (instance *Msvm_CollectionReferencePointExportJob) GetPropertyBaseReferencePointGroupId() (value string, err error) {
	retValue, err := instance.GetProperty("BaseReferencePointGroupId")
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

// SetCancellable sets the value of Cancellable for the instance
func (instance *Msvm_CollectionReferencePointExportJob) SetPropertyCancellable(value bool) (err error) {
	return instance.SetProperty("Cancellable", (value))
}

// GetCancellable gets the value of Cancellable for the instance
func (instance *Msvm_CollectionReferencePointExportJob) GetPropertyCancellable() (value bool, err error) {
	retValue, err := instance.GetProperty("Cancellable")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetCollectionId sets the value of CollectionId for the instance
func (instance *Msvm_CollectionReferencePointExportJob) SetPropertyCollectionId(value string) (err error) {
	return instance.SetProperty("CollectionId", (value))
}

// GetCollectionId gets the value of CollectionId for the instance
func (instance *Msvm_CollectionReferencePointExportJob) GetPropertyCollectionId() (value string, err error) {
	retValue, err := instance.GetProperty("CollectionId")
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

// SetErrorSummaryDescription sets the value of ErrorSummaryDescription for the instance
func (instance *Msvm_CollectionReferencePointExportJob) SetPropertyErrorSummaryDescription(value string) (err error) {
	return instance.SetProperty("ErrorSummaryDescription", (value))
}

// GetErrorSummaryDescription gets the value of ErrorSummaryDescription for the instance
func (instance *Msvm_CollectionReferencePointExportJob) GetPropertyErrorSummaryDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorSummaryDescription")
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

// SetExportDirectory sets the value of ExportDirectory for the instance
func (instance *Msvm_CollectionReferencePointExportJob) SetPropertyExportDirectory(value string) (err error) {
	return instance.SetProperty("ExportDirectory", (value))
}

// GetExportDirectory gets the value of ExportDirectory for the instance
func (instance *Msvm_CollectionReferencePointExportJob) GetPropertyExportDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("ExportDirectory")
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

// SetExportedConfigFilePaths sets the value of ExportedConfigFilePaths for the instance
func (instance *Msvm_CollectionReferencePointExportJob) SetPropertyExportedConfigFilePaths(value []string) (err error) {
	return instance.SetProperty("ExportedConfigFilePaths", (value))
}

// GetExportedConfigFilePaths gets the value of ExportedConfigFilePaths for the instance
func (instance *Msvm_CollectionReferencePointExportJob) GetPropertyExportedConfigFilePaths() (value []string, err error) {
	retValue, err := instance.GetProperty("ExportedConfigFilePaths")
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

// SetExportedDisks sets the value of ExportedDisks for the instance
func (instance *Msvm_CollectionReferencePointExportJob) SetPropertyExportedDisks(value []string) (err error) {
	return instance.SetProperty("ExportedDisks", (value))
}

// GetExportedDisks gets the value of ExportedDisks for the instance
func (instance *Msvm_CollectionReferencePointExportJob) GetPropertyExportedDisks() (value []string, err error) {
	retValue, err := instance.GetProperty("ExportedDisks")
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

// SetExportedGuestStateFilePaths sets the value of ExportedGuestStateFilePaths for the instance
func (instance *Msvm_CollectionReferencePointExportJob) SetPropertyExportedGuestStateFilePaths(value []string) (err error) {
	return instance.SetProperty("ExportedGuestStateFilePaths", (value))
}

// GetExportedGuestStateFilePaths gets the value of ExportedGuestStateFilePaths for the instance
func (instance *Msvm_CollectionReferencePointExportJob) GetPropertyExportedGuestStateFilePaths() (value []string, err error) {
	retValue, err := instance.GetProperty("ExportedGuestStateFilePaths")
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

// SetExportedLogFilePaths sets the value of ExportedLogFilePaths for the instance
func (instance *Msvm_CollectionReferencePointExportJob) SetPropertyExportedLogFilePaths(value []string) (err error) {
	return instance.SetProperty("ExportedLogFilePaths", (value))
}

// GetExportedLogFilePaths gets the value of ExportedLogFilePaths for the instance
func (instance *Msvm_CollectionReferencePointExportJob) GetPropertyExportedLogFilePaths() (value []string, err error) {
	retValue, err := instance.GetProperty("ExportedLogFilePaths")
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

// SetExportedRuntimeFilePaths sets the value of ExportedRuntimeFilePaths for the instance
func (instance *Msvm_CollectionReferencePointExportJob) SetPropertyExportedRuntimeFilePaths(value []string) (err error) {
	return instance.SetProperty("ExportedRuntimeFilePaths", (value))
}

// GetExportedRuntimeFilePaths gets the value of ExportedRuntimeFilePaths for the instance
func (instance *Msvm_CollectionReferencePointExportJob) GetPropertyExportedRuntimeFilePaths() (value []string, err error) {
	retValue, err := instance.GetProperty("ExportedRuntimeFilePaths")
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

// SetReferencePointGroupId sets the value of ReferencePointGroupId for the instance
func (instance *Msvm_CollectionReferencePointExportJob) SetPropertyReferencePointGroupId(value string) (err error) {
	return instance.SetProperty("ReferencePointGroupId", (value))
}

// GetReferencePointGroupId gets the value of ReferencePointGroupId for the instance
func (instance *Msvm_CollectionReferencePointExportJob) GetPropertyReferencePointGroupId() (value string, err error) {
	retValue, err := instance.GetProperty("ReferencePointGroupId")
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

// SetVirtualMachineId sets the value of VirtualMachineId for the instance
func (instance *Msvm_CollectionReferencePointExportJob) SetPropertyVirtualMachineId(value []string) (err error) {
	return instance.SetProperty("VirtualMachineId", (value))
}

// GetVirtualMachineId gets the value of VirtualMachineId for the instance
func (instance *Msvm_CollectionReferencePointExportJob) GetPropertyVirtualMachineId() (value []string, err error) {
	retValue, err := instance.GetProperty("VirtualMachineId")
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

// <param name="Errors" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CollectionReferencePointExportJob) GetErrorEx( /* OUT */ Errors []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetErrorEx")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
