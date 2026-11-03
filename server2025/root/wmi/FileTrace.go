// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// FileTrace struct
type FileTrace struct {
	*MSNT_FileBaseTrace_Set1
}

func NewFileTraceEx1(instance *cim.WmiInstance) (newInstance *FileTrace, err error) {
	tmp, err := NewMSNT_FileBaseTrace_Set1Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &FileTrace{
		MSNT_FileBaseTrace_Set1: tmp,
	}
	return
}

func NewFileTraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FileTrace, err error) {
	tmp, err := NewMSNT_FileBaseTrace_Set1Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FileTrace{
		MSNT_FileBaseTrace_Set1: tmp,
	}
	return
}
