// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MintPipelineExecution struct
type MSFT_MintPipelineExecution struct {
	*MSFT_PipelineExecution
}

func NewMSFT_MintPipelineExecutionEx1(instance *cim.WmiInstance) (newInstance *MSFT_MintPipelineExecution, err error) {
	tmp, err := NewMSFT_PipelineExecutionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MintPipelineExecution{
		MSFT_PipelineExecution: tmp,
	}
	return
}

func NewMSFT_MintPipelineExecutionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MintPipelineExecution, err error) {
	tmp, err := NewMSFT_PipelineExecutionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MintPipelineExecution{
		MSFT_PipelineExecution: tmp,
	}
	return
}
