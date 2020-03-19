// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_ConcreteJob struct
type MLNX_ConcreteJob struct {
	*CIM_ConcreteJob

	//
	CommandLine string
}

func NewMLNX_ConcreteJobEx1(instance *cim.WmiInstance) (newInstance *MLNX_ConcreteJob, err error) {
	tmp, err := NewCIM_ConcreteJobEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_ConcreteJob{
		CIM_ConcreteJob: tmp,
	}
	return
}

func NewMLNX_ConcreteJobEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_ConcreteJob, err error) {
	tmp, err := NewCIM_ConcreteJobEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_ConcreteJob{
		CIM_ConcreteJob: tmp,
	}
	return
}

// SetCommandLine sets the value of CommandLine for the instance
func (instance *MLNX_ConcreteJob) SetPropertyCommandLine(value string) (err error) {
	return instance.SetProperty("CommandLine", value)
}

// GetCommandLine gets the value of CommandLine for the instance
func (instance *MLNX_ConcreteJob) GetPropertyCommandLine() (value string, err error) {
	retValue, err := instance.GetProperty("CommandLine")
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

// <param name="Priority" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_ConcreteJob) SetPriority( /* IN */ Priority uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetPriority", Priority)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="CommandLine" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_ConcreteJob) Create( /* IN */ CommandLine string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("Create", Action, PercentComplete, Timeout, CommandLine)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
