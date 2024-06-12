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

// MSFT_MintPipeline struct
type MSFT_MintPipeline struct {
	*MSFT_Pipeline
}

func NewMSFT_MintPipelineEx1(instance *cim.WmiInstance) (newInstance *MSFT_MintPipeline, err error) {
	tmp, err := NewMSFT_PipelineEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MintPipeline{
		MSFT_Pipeline: tmp,
	}
	return
}

func NewMSFT_MintPipelineEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MintPipeline, err error) {
	tmp, err := NewMSFT_PipelineEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MintPipeline{
		MSFT_Pipeline: tmp,
	}
	return
}
