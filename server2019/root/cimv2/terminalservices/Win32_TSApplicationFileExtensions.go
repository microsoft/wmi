// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

// Win32_TSApplicationFileExtensions struct
type Win32_TSApplicationFileExtensions struct {
	CIM_LogicalElement
}

// Gives the list of file extensions owned by an application

// <param name="AppPath" type="string ">Path to the application</param>

// <param name="Extensions" type="string []">File extensions owned by the application</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSApplicationFileExtensions) FileExtensions( /* IN */ AppPath string,
	/* OUT */ Extensions []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("FileExtensions", AppPath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Scans the registry to get the current file associations for an application.

// <param name="AppPath" type="string ">Path to the application</param>

// <param name="FileAssociations" type="Win32_RDFileAssociation []">File associations for this application</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSApplicationFileExtensions) FileAssociations( /* IN */ AppPath string,
	/* OUT */ FileAssociations []Win32_RDFileAssociation) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("FileAssociations", AppPath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
