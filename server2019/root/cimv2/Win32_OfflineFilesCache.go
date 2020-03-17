// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_OfflineFilesCache struct
type Win32_OfflineFilesCache struct {
	cim.WmiInstance

	//
	Active bool

	//
	Enabled bool

	//
	Location string
}

// SetActive sets the value of Active for the instance
func (instance *Win32_OfflineFilesCache) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", value)
}

// GetActive gets the value of Active for the instance
func (instance *Win32_OfflineFilesCache) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *Win32_OfflineFilesCache) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *Win32_OfflineFilesCache) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocation sets the value of Location for the instance
func (instance *Win32_OfflineFilesCache) SetPropertyLocation(value string) (err error) {
	return instance.SetProperty("Location", value)
}

// GetLocation gets the value of Location for the instance
func (instance *Win32_OfflineFilesCache) GetPropertyLocation() (value string, err error) {
	retValue, err := instance.GetProperty("Location")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Enable" type="bool "></param>

// <param name="RebootRequired" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_OfflineFilesCache) Enable( /* IN */ Enable bool,
	/* OUT */ RebootRequired bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable", Enable)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="NewPath" type="string "></param>
// <param name="OriginalPath" type="string "></param>
// <param name="ReplaceIfExists" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_OfflineFilesCache) RenameItem( /* IN */ OriginalPath string,
	/* IN */ NewPath string,
	/* IN */ ReplaceIfExists bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RenameItem", OriginalPath, NewPath, ReplaceIfExists)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewPath" type="string "></param>
// <param name="OriginalPath" type="string "></param>
// <param name="ReplaceIfExists" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_OfflineFilesCache) RenameItemEx( /* IN */ OriginalPath string,
	/* IN */ NewPath string,
	/* IN */ ReplaceIfExists bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RenameItemEx", OriginalPath, NewPath, ReplaceIfExists)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Paths" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_OfflineFilesCache) Synchronize( /* IN */ Paths []string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Synchronize", Paths, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Deep" type="bool "></param>
// <param name="Flags" type="uint32 "></param>
// <param name="Paths" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_OfflineFilesCache) Pin( /* IN */ Paths []string,
	/* IN */ Flags uint32,
	/* IN */ Deep bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Pin", Paths, Flags, Deep)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Deep" type="bool "></param>
// <param name="Flags" type="uint32 "></param>
// <param name="Paths" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_OfflineFilesCache) Unpin( /* IN */ Paths []string,
	/* IN */ Flags uint32,
	/* IN */ Deep bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Unpin", Paths, Flags, Deep)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Paths" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_OfflineFilesCache) DeleteItems( /* IN */ Paths []string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DeleteItems", Paths, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Encrypt" type="bool "></param>
// <param name="Flags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_OfflineFilesCache) Encrypt( /* IN */ Encrypt bool,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Encrypt", Encrypt, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Path" type="string "></param>
// <param name="Suspend" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_OfflineFilesCache) SuspendRoot( /* IN */ Path string,
	/* IN */ Suspend bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SuspendRoot", Path, Suspend)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Force" type="bool "></param>
// <param name="Path" type="string "></param>

// <param name="OpenFiles" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_OfflineFilesCache) TransitionOffline( /* IN */ Path string,
	/* IN */ Force bool,
	/* IN */ Flags uint32,
	/* OUT */ OpenFiles bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("TransitionOffline", Path, Force, Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Path" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_OfflineFilesCache) TransitionOnline( /* IN */ Path string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("TransitionOnline", Path, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
