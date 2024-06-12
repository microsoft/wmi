// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// BcdStore struct
type BcdStore struct {
	*cim.WmiInstance

	// A BcdStore is uniquely identified by its file path. The system store is denoted via an empty file path.
	FilePath string
}

func NewBcdStoreEx1(instance *cim.WmiInstance) (newInstance *BcdStore, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &BcdStore{
		WmiInstance: tmp,
	}
	return
}

func NewBcdStoreEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BcdStore, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BcdStore{
		WmiInstance: tmp,
	}
	return
}

// SetFilePath sets the value of FilePath for the instance
func (instance *BcdStore) SetPropertyFilePath(value string) (err error) {
	return instance.SetProperty("FilePath", (value))
}

// GetFilePath gets the value of FilePath for the instance
func (instance *BcdStore) GetPropertyFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("FilePath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

//

// <param name="File" type="string "></param>

// <param name="ReturnValue" type="bool "></param>
// <param name="Store" type="BcdStore "></param>
func (instance *BcdStore) OpenStore( /* IN */ File string,
	/* OUT */ Store BcdStore) (result bool, err error) {
	retVal, err := instance.InvokeMethod("OpenStore", File)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = bool(retValue != 0)
	return

}

//

// <param name="File" type="string "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdStore) ImportStore( /* IN */ File string) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ImportStore", File)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="File" type="string "></param>
// <param name="Flags" type="uint32 "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdStore) ImportStoreWithFlags( /* IN */ File string,
	/* IN */ Flags uint32) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ImportStoreWithFlags", File, Flags)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="File" type="string "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdStore) ExportStore( /* IN */ File string) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ExportStore", File)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="File" type="string "></param>

// <param name="ReturnValue" type="bool "></param>
// <param name="Store" type="BcdStore "></param>
func (instance *BcdStore) CreateStore( /* IN */ File string,
	/* OUT */ Store BcdStore) (result bool, err error) {
	retVal, err := instance.InvokeMethod("CreateStore", File)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = bool(retValue != 0)
	return

}

//

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdStore) DeleteSystemStore() (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DeleteSystemStore")
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="Disk" type="string "></param>
// <param name="ReturnValue" type="bool "></param>
func (instance *BcdStore) GetSystemDisk( /* OUT */ Disk string) (result bool, err error) {
	retVal, err := instance.InvokeMethod("GetSystemDisk")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = bool(retValue != 0)
	return

}

//

// <param name="Partition" type="string "></param>
// <param name="ReturnValue" type="bool "></param>
func (instance *BcdStore) GetSystemPartition( /* OUT */ Partition string) (result bool, err error) {
	retVal, err := instance.InvokeMethod("GetSystemPartition")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = bool(retValue != 0)
	return

}

//

// <param name="Partition" type="string "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdStore) SetSystemStoreDevice( /* IN */ Partition string) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetSystemStoreDevice", Partition)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="Type" type="uint32 "></param>

// <param name="Objects" type="BcdObject []"></param>
// <param name="ReturnValue" type="bool "></param>
func (instance *BcdStore) EnumerateObjects( /* IN */ Type uint32,
	/* OUT */ Objects []BcdObject) (result bool, err error) {
	retVal, err := instance.InvokeMethod("EnumerateObjects", Type)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = bool(retValue != 0)
	return

}

//

// <param name="Id" type="string "></param>

// <param name="Object" type="BcdObject "></param>
// <param name="ReturnValue" type="bool "></param>
func (instance *BcdStore) OpenObject( /* IN */ Id string,
	/* OUT */ Object BcdObject) (result bool, err error) {
	retVal, err := instance.InvokeMethod("OpenObject", Id)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = bool(retValue != 0)
	return

}

//

// <param name="Id" type="string "></param>
// <param name="Type" type="uint32 "></param>

// <param name="Object" type="BcdObject "></param>
// <param name="ReturnValue" type="bool "></param>
func (instance *BcdStore) CreateObject( /* IN */ Id string,
	/* IN */ Type uint32,
	/* OUT */ Object BcdObject) (result bool, err error) {
	retVal, err := instance.InvokeMethod("CreateObject", Id, Type)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = bool(retValue != 0)
	return

}

//

// <param name="Id" type="string "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdStore) DeleteObject( /* IN */ Id string) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DeleteObject", Id)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="SourceId" type="string "></param>
// <param name="SourceStoreFile" type="string "></param>

// <param name="Object" type="BcdObject "></param>
// <param name="ReturnValue" type="bool "></param>
func (instance *BcdStore) CopyObject( /* IN */ SourceStoreFile string,
	/* IN */ SourceId string,
	/* IN */ Flags uint32,
	/* OUT */ Object BcdObject) (result bool, err error) {
	retVal, err := instance.InvokeMethod("CopyObject", SourceStoreFile, SourceId, Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = bool(retValue != 0)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="SourceStoreFile" type="string "></param>
// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *BcdStore) CopyObjects( /* IN */ SourceStoreFile string,
	/* IN */ Type uint32,
	/* IN */ Flags uint32) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CopyObjects", SourceStoreFile, Type, Flags)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}
