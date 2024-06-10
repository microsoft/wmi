// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_DiagnosticTest struct
type MLNX_DiagnosticTest struct {
	*CIM_DiagnosticTest
}

func NewMLNX_DiagnosticTestEx1(instance *cim.WmiInstance) (newInstance *MLNX_DiagnosticTest, err error) {
	tmp, err := NewCIM_DiagnosticTestEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticTest{
		CIM_DiagnosticTest: tmp,
	}
	return
}

func NewMLNX_DiagnosticTestEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DiagnosticTest, err error) {
	tmp, err := NewCIM_DiagnosticTestEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticTest{
		CIM_DiagnosticTest: tmp,
	}
	return
}
