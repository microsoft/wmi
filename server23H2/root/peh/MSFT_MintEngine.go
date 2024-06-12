// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MintEngine struct
type MSFT_MintEngine struct {
	*MSFT_Engine
}

func NewMSFT_MintEngineEx1(instance *cim.WmiInstance) (newInstance *MSFT_MintEngine, err error) {
	tmp, err := NewMSFT_EngineEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MintEngine{
		MSFT_Engine: tmp,
	}
	return
}

func NewMSFT_MintEngineEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MintEngine, err error) {
	tmp, err := NewMSFT_EngineEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MintEngine{
		MSFT_Engine: tmp,
	}
	return
}

//

// <param name="args" type="interface{} "></param>
// <param name="engine" type="MSFT_Engine "></param>
// <param name="expression" type="MSFT_Expression "></param>
// <param name="modules" type="string []"></param>
// <param name="pipelineExecutionName" type="string "></param>

// <param name="pipelineExecution" type="MSFT_PipelineExecution "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MintEngine) ExecuteExpression( /* IN */ expression MSFT_Expression,
	/* IN */ args interface{},
	/* IN */ modules []string,
	/* IN */ engine MSFT_Engine,
	/* IN */ pipelineExecutionName string,
	/* OUT */ pipelineExecution MSFT_PipelineExecution) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ExecuteExpression", expression, args, modules, engine, pipelineExecutionName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Flags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MintEngine) Wakeup( /* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Wakeup", Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
