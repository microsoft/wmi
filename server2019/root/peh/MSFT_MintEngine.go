// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

// MSFT_MintEngine struct
type MSFT_MintEngine struct {
	MSFT_Engine
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
