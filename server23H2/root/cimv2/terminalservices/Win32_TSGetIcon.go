// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_TSGetIcon struct
type Win32_TSGetIcon struct {
	*CIM_LogicalElement
}

func NewWin32_TSGetIconEx1(instance *cim.WmiInstance) (newInstance *Win32_TSGetIcon, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSGetIcon{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewWin32_TSGetIconEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSGetIcon, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSGetIcon{
		CIM_LogicalElement: tmp,
	}
	return
}

// Returns the contents of the Icon in the filepath using the icon index

// <param name="FilePath" type="string ">String that contains the path to the file that contains the icon</param>
// <param name="Index" type="int32 ">Index for the Icon wanted</param>

// <param name="IconContents" type="uint8 []">Contents of the Icon</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSGetIcon) GetIcon( /* IN */ FilePath string,
	/* IN */ Index int32,
	/* OUT */ IconContents []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetIcon", FilePath, Index)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
