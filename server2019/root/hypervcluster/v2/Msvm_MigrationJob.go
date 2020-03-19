// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_MigrationJob struct
type Msvm_MigrationJob struct {
	*CIM_ConcreteJob

	//
	Cancellable bool

	//
	DestinationHost string

	//
	ErrorSummaryDescription string

	//
	JobType MigrationJob_JobType

	//
	MigrationType uint16

	//
	NewResourceSettingData []string

	//
	NewSystemSettingData string

	//
	VirtualSystemName string
}

func NewMsvm_MigrationJobEx1(instance *cim.WmiInstance) (newInstance *Msvm_MigrationJob, err error) {
	tmp, err := NewCIM_ConcreteJobEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_MigrationJob{
		CIM_ConcreteJob: tmp,
	}
	return
}

func NewMsvm_MigrationJobEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_MigrationJob, err error) {
	tmp, err := NewCIM_ConcreteJobEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_MigrationJob{
		CIM_ConcreteJob: tmp,
	}
	return
}

// SetCancellable sets the value of Cancellable for the instance
func (instance *Msvm_MigrationJob) SetPropertyCancellable(value bool) (err error) {
	return instance.SetProperty("Cancellable", value)
}

// GetCancellable gets the value of Cancellable for the instance
func (instance *Msvm_MigrationJob) GetPropertyCancellable() (value bool, err error) {
	retValue, err := instance.GetProperty("Cancellable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDestinationHost sets the value of DestinationHost for the instance
func (instance *Msvm_MigrationJob) SetPropertyDestinationHost(value string) (err error) {
	return instance.SetProperty("DestinationHost", value)
}

// GetDestinationHost gets the value of DestinationHost for the instance
func (instance *Msvm_MigrationJob) GetPropertyDestinationHost() (value string, err error) {
	retValue, err := instance.GetProperty("DestinationHost")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorSummaryDescription sets the value of ErrorSummaryDescription for the instance
func (instance *Msvm_MigrationJob) SetPropertyErrorSummaryDescription(value string) (err error) {
	return instance.SetProperty("ErrorSummaryDescription", value)
}

// GetErrorSummaryDescription gets the value of ErrorSummaryDescription for the instance
func (instance *Msvm_MigrationJob) GetPropertyErrorSummaryDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorSummaryDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobType sets the value of JobType for the instance
func (instance *Msvm_MigrationJob) SetPropertyJobType(value MigrationJob_JobType) (err error) {
	return instance.SetProperty("JobType", value)
}

// GetJobType gets the value of JobType for the instance
func (instance *Msvm_MigrationJob) GetPropertyJobType() (value MigrationJob_JobType, err error) {
	retValue, err := instance.GetProperty("JobType")
	if err != nil {
		return
	}
	value, ok := retValue.(MigrationJob_JobType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMigrationType sets the value of MigrationType for the instance
func (instance *Msvm_MigrationJob) SetPropertyMigrationType(value uint16) (err error) {
	return instance.SetProperty("MigrationType", value)
}

// GetMigrationType gets the value of MigrationType for the instance
func (instance *Msvm_MigrationJob) GetPropertyMigrationType() (value uint16, err error) {
	retValue, err := instance.GetProperty("MigrationType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNewResourceSettingData sets the value of NewResourceSettingData for the instance
func (instance *Msvm_MigrationJob) SetPropertyNewResourceSettingData(value []string) (err error) {
	return instance.SetProperty("NewResourceSettingData", value)
}

// GetNewResourceSettingData gets the value of NewResourceSettingData for the instance
func (instance *Msvm_MigrationJob) GetPropertyNewResourceSettingData() (value []string, err error) {
	retValue, err := instance.GetProperty("NewResourceSettingData")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNewSystemSettingData sets the value of NewSystemSettingData for the instance
func (instance *Msvm_MigrationJob) SetPropertyNewSystemSettingData(value string) (err error) {
	return instance.SetProperty("NewSystemSettingData", value)
}

// GetNewSystemSettingData gets the value of NewSystemSettingData for the instance
func (instance *Msvm_MigrationJob) GetPropertyNewSystemSettingData() (value string, err error) {
	retValue, err := instance.GetProperty("NewSystemSettingData")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualSystemName sets the value of VirtualSystemName for the instance
func (instance *Msvm_MigrationJob) SetPropertyVirtualSystemName(value string) (err error) {
	return instance.SetProperty("VirtualSystemName", value)
}

// GetVirtualSystemName gets the value of VirtualSystemName for the instance
func (instance *Msvm_MigrationJob) GetPropertyVirtualSystemName() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualSystemName")
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

// <param name="Errors" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_MigrationJob) GetErrorEx( /* OUT */ Errors []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetErrorEx")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
