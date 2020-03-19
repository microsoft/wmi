// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MintPipelineDefinition struct
type MSFT_MintPipelineDefinition struct {
	*MSFT_PipelineDefinition

	//
	expression MSFT_Expression
}

func NewMSFT_MintPipelineDefinitionEx1(instance *cim.WmiInstance) (newInstance *MSFT_MintPipelineDefinition, err error) {
	tmp, err := NewMSFT_PipelineDefinitionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MintPipelineDefinition{
		MSFT_PipelineDefinition: tmp,
	}
	return
}

func NewMSFT_MintPipelineDefinitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MintPipelineDefinition, err error) {
	tmp, err := NewMSFT_PipelineDefinitionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MintPipelineDefinition{
		MSFT_PipelineDefinition: tmp,
	}
	return
}

// Setexpression sets the value of expression for the instance
func (instance *MSFT_MintPipelineDefinition) SetPropertyexpression(value MSFT_Expression) (err error) {
	return instance.SetProperty("expression", value)
}

// Getexpression gets the value of expression for the instance
func (instance *MSFT_MintPipelineDefinition) GetPropertyexpression() (value MSFT_Expression, err error) {
	retValue, err := instance.GetProperty("expression")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Expression)
	if !ok {
		// TODO: Set an error
	}
	return
}
