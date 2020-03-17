// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

// Win32_TSGetIcon struct
type Win32_TSGetIcon struct {
	CIM_LogicalElement
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
