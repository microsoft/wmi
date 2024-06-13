// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore struct
type Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore struct {
	*Win32_PerfFormattedData

	//
	Cacheupdateoperationcount uint64

	//
	Cacheupdateoperationlatencymicroseconds uint64

	//
	Commitbytecount uint64

	//
	Commitbytelatencymicroseconds uint64

	//
	Commitcount uint64

	//
	Commitoperationcount uint64

	//
	Commitoperationlatencymicroseconds uint64

	//
	Compactoperationcount uint64

	//
	Compactoperationlatencymicroseconds uint64

	//
	CurrentreplaylogSize uint64

	//
	Dataalignment uint64

	//
	Dataend uint64

	//
	Disconnectcount uint64

	//
	Filedatasizeinbytes uint64

	//
	Filelockacquirelatencymicroseconds uint64

	//
	Filelockcount uint64

	//
	Filelockreleaselatencymicroseconds uint64

	//
	Fragmentationratio uint64

	//
	Getoperationcount uint64

	//
	Getoperationlatencymicroseconds uint64

	//
	Loadfileoperationcount uint64

	//
	Loadfileoperationlatencymicroseconds uint64

	//
	Namessizeinbytes uint64

	//
	Numberofavailableentriesinsideobjecttables uint64

	//
	Numberofemptyentriesinsideobjecttables uint64

	//
	Numberoffileobjects uint64

	//
	Numberoffreebytesinsidekeytables uint64

	//
	Numberofkeys uint64

	//
	Numberofkeytables uint64

	//
	Numberofobjecttables uint64

	//
	Querysizeoperationcount uint64

	//
	Querysizeoperationlatencymicroseconds uint64

	//
	Readfromfilebytecount uint64

	//
	Readfromfilebytelatencymicroseconds uint64

	//
	Readfromfilecount uint64

	//
	Readfromstoragebytecount uint64

	//
	Readfromstoragebytelatencymicroseconds uint64

	//
	Readfromstoragecount uint64

	//
	Reconnectlatencymicroseconds uint64

	//
	Removeoperationcount uint64

	//
	Removeoperationlatencymicroseconds uint64

	//
	Sectorsize uint64

	//
	Setoperationcount uint64

	//
	Setoperationlatencymicroseconds uint64

	//
	Storagelockacquirelatencymicroseconds uint64

	//
	Storagelockcount uint64

	//
	Storagelockreleaselatencymicroseconds uint64

	//
	Tabledatasizeinbytes uint64

	//
	Writetofilebytecount uint64

	//
	Writetofilebytelatencymicroseconds uint64

	//
	Writetofilecount uint64

	//
	Writetostoragebytecount uint64

	//
	Writetostoragebytelatencymicroseconds uint64

	//
	Writetostoragecount uint64
}

func NewWin32_PerfFormattedData_HyperVStorageStats_HyperVDataStoreEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_HyperVStorageStats_HyperVDataStoreEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetCacheupdateoperationcount sets the value of Cacheupdateoperationcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyCacheupdateoperationcount(value uint64) (err error) {
	return instance.SetProperty("Cacheupdateoperationcount", (value))
}

// GetCacheupdateoperationcount gets the value of Cacheupdateoperationcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyCacheupdateoperationcount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Cacheupdateoperationcount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCacheupdateoperationlatencymicroseconds sets the value of Cacheupdateoperationlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyCacheupdateoperationlatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Cacheupdateoperationlatencymicroseconds", (value))
}

// GetCacheupdateoperationlatencymicroseconds gets the value of Cacheupdateoperationlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyCacheupdateoperationlatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Cacheupdateoperationlatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCommitbytecount sets the value of Commitbytecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyCommitbytecount(value uint64) (err error) {
	return instance.SetProperty("Commitbytecount", (value))
}

// GetCommitbytecount gets the value of Commitbytecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyCommitbytecount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Commitbytecount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCommitbytelatencymicroseconds sets the value of Commitbytelatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyCommitbytelatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Commitbytelatencymicroseconds", (value))
}

// GetCommitbytelatencymicroseconds gets the value of Commitbytelatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyCommitbytelatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Commitbytelatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCommitcount sets the value of Commitcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyCommitcount(value uint64) (err error) {
	return instance.SetProperty("Commitcount", (value))
}

// GetCommitcount gets the value of Commitcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyCommitcount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Commitcount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCommitoperationcount sets the value of Commitoperationcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyCommitoperationcount(value uint64) (err error) {
	return instance.SetProperty("Commitoperationcount", (value))
}

// GetCommitoperationcount gets the value of Commitoperationcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyCommitoperationcount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Commitoperationcount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCommitoperationlatencymicroseconds sets the value of Commitoperationlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyCommitoperationlatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Commitoperationlatencymicroseconds", (value))
}

// GetCommitoperationlatencymicroseconds gets the value of Commitoperationlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyCommitoperationlatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Commitoperationlatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCompactoperationcount sets the value of Compactoperationcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyCompactoperationcount(value uint64) (err error) {
	return instance.SetProperty("Compactoperationcount", (value))
}

// GetCompactoperationcount gets the value of Compactoperationcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyCompactoperationcount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Compactoperationcount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCompactoperationlatencymicroseconds sets the value of Compactoperationlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyCompactoperationlatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Compactoperationlatencymicroseconds", (value))
}

// GetCompactoperationlatencymicroseconds gets the value of Compactoperationlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyCompactoperationlatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Compactoperationlatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCurrentreplaylogSize sets the value of CurrentreplaylogSize for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyCurrentreplaylogSize(value uint64) (err error) {
	return instance.SetProperty("CurrentreplaylogSize", (value))
}

// GetCurrentreplaylogSize gets the value of CurrentreplaylogSize for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyCurrentreplaylogSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentreplaylogSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetDataalignment sets the value of Dataalignment for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyDataalignment(value uint64) (err error) {
	return instance.SetProperty("Dataalignment", (value))
}

// GetDataalignment gets the value of Dataalignment for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyDataalignment() (value uint64, err error) {
	retValue, err := instance.GetProperty("Dataalignment")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetDataend sets the value of Dataend for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyDataend(value uint64) (err error) {
	return instance.SetProperty("Dataend", (value))
}

// GetDataend gets the value of Dataend for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyDataend() (value uint64, err error) {
	retValue, err := instance.GetProperty("Dataend")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetDisconnectcount sets the value of Disconnectcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyDisconnectcount(value uint64) (err error) {
	return instance.SetProperty("Disconnectcount", (value))
}

// GetDisconnectcount gets the value of Disconnectcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyDisconnectcount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Disconnectcount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFiledatasizeinbytes sets the value of Filedatasizeinbytes for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyFiledatasizeinbytes(value uint64) (err error) {
	return instance.SetProperty("Filedatasizeinbytes", (value))
}

// GetFiledatasizeinbytes gets the value of Filedatasizeinbytes for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyFiledatasizeinbytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("Filedatasizeinbytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFilelockacquirelatencymicroseconds sets the value of Filelockacquirelatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyFilelockacquirelatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Filelockacquirelatencymicroseconds", (value))
}

// GetFilelockacquirelatencymicroseconds gets the value of Filelockacquirelatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyFilelockacquirelatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Filelockacquirelatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFilelockcount sets the value of Filelockcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyFilelockcount(value uint64) (err error) {
	return instance.SetProperty("Filelockcount", (value))
}

// GetFilelockcount gets the value of Filelockcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyFilelockcount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Filelockcount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFilelockreleaselatencymicroseconds sets the value of Filelockreleaselatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyFilelockreleaselatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Filelockreleaselatencymicroseconds", (value))
}

// GetFilelockreleaselatencymicroseconds gets the value of Filelockreleaselatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyFilelockreleaselatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Filelockreleaselatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFragmentationratio sets the value of Fragmentationratio for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyFragmentationratio(value uint64) (err error) {
	return instance.SetProperty("Fragmentationratio", (value))
}

// GetFragmentationratio gets the value of Fragmentationratio for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyFragmentationratio() (value uint64, err error) {
	retValue, err := instance.GetProperty("Fragmentationratio")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetGetoperationcount sets the value of Getoperationcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyGetoperationcount(value uint64) (err error) {
	return instance.SetProperty("Getoperationcount", (value))
}

// GetGetoperationcount gets the value of Getoperationcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyGetoperationcount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Getoperationcount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetGetoperationlatencymicroseconds sets the value of Getoperationlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyGetoperationlatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Getoperationlatencymicroseconds", (value))
}

// GetGetoperationlatencymicroseconds gets the value of Getoperationlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyGetoperationlatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Getoperationlatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetLoadfileoperationcount sets the value of Loadfileoperationcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyLoadfileoperationcount(value uint64) (err error) {
	return instance.SetProperty("Loadfileoperationcount", (value))
}

// GetLoadfileoperationcount gets the value of Loadfileoperationcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyLoadfileoperationcount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Loadfileoperationcount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetLoadfileoperationlatencymicroseconds sets the value of Loadfileoperationlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyLoadfileoperationlatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Loadfileoperationlatencymicroseconds", (value))
}

// GetLoadfileoperationlatencymicroseconds gets the value of Loadfileoperationlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyLoadfileoperationlatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Loadfileoperationlatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetNamessizeinbytes sets the value of Namessizeinbytes for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyNamessizeinbytes(value uint64) (err error) {
	return instance.SetProperty("Namessizeinbytes", (value))
}

// GetNamessizeinbytes gets the value of Namessizeinbytes for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyNamessizeinbytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("Namessizeinbytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetNumberofavailableentriesinsideobjecttables sets the value of Numberofavailableentriesinsideobjecttables for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyNumberofavailableentriesinsideobjecttables(value uint64) (err error) {
	return instance.SetProperty("Numberofavailableentriesinsideobjecttables", (value))
}

// GetNumberofavailableentriesinsideobjecttables gets the value of Numberofavailableentriesinsideobjecttables for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyNumberofavailableentriesinsideobjecttables() (value uint64, err error) {
	retValue, err := instance.GetProperty("Numberofavailableentriesinsideobjecttables")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetNumberofemptyentriesinsideobjecttables sets the value of Numberofemptyentriesinsideobjecttables for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyNumberofemptyentriesinsideobjecttables(value uint64) (err error) {
	return instance.SetProperty("Numberofemptyentriesinsideobjecttables", (value))
}

// GetNumberofemptyentriesinsideobjecttables gets the value of Numberofemptyentriesinsideobjecttables for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyNumberofemptyentriesinsideobjecttables() (value uint64, err error) {
	retValue, err := instance.GetProperty("Numberofemptyentriesinsideobjecttables")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetNumberoffileobjects sets the value of Numberoffileobjects for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyNumberoffileobjects(value uint64) (err error) {
	return instance.SetProperty("Numberoffileobjects", (value))
}

// GetNumberoffileobjects gets the value of Numberoffileobjects for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyNumberoffileobjects() (value uint64, err error) {
	retValue, err := instance.GetProperty("Numberoffileobjects")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetNumberoffreebytesinsidekeytables sets the value of Numberoffreebytesinsidekeytables for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyNumberoffreebytesinsidekeytables(value uint64) (err error) {
	return instance.SetProperty("Numberoffreebytesinsidekeytables", (value))
}

// GetNumberoffreebytesinsidekeytables gets the value of Numberoffreebytesinsidekeytables for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyNumberoffreebytesinsidekeytables() (value uint64, err error) {
	retValue, err := instance.GetProperty("Numberoffreebytesinsidekeytables")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetNumberofkeys sets the value of Numberofkeys for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyNumberofkeys(value uint64) (err error) {
	return instance.SetProperty("Numberofkeys", (value))
}

// GetNumberofkeys gets the value of Numberofkeys for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyNumberofkeys() (value uint64, err error) {
	retValue, err := instance.GetProperty("Numberofkeys")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetNumberofkeytables sets the value of Numberofkeytables for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyNumberofkeytables(value uint64) (err error) {
	return instance.SetProperty("Numberofkeytables", (value))
}

// GetNumberofkeytables gets the value of Numberofkeytables for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyNumberofkeytables() (value uint64, err error) {
	retValue, err := instance.GetProperty("Numberofkeytables")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetNumberofobjecttables sets the value of Numberofobjecttables for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyNumberofobjecttables(value uint64) (err error) {
	return instance.SetProperty("Numberofobjecttables", (value))
}

// GetNumberofobjecttables gets the value of Numberofobjecttables for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyNumberofobjecttables() (value uint64, err error) {
	retValue, err := instance.GetProperty("Numberofobjecttables")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetQuerysizeoperationcount sets the value of Querysizeoperationcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyQuerysizeoperationcount(value uint64) (err error) {
	return instance.SetProperty("Querysizeoperationcount", (value))
}

// GetQuerysizeoperationcount gets the value of Querysizeoperationcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyQuerysizeoperationcount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Querysizeoperationcount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetQuerysizeoperationlatencymicroseconds sets the value of Querysizeoperationlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyQuerysizeoperationlatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Querysizeoperationlatencymicroseconds", (value))
}

// GetQuerysizeoperationlatencymicroseconds gets the value of Querysizeoperationlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyQuerysizeoperationlatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Querysizeoperationlatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetReadfromfilebytecount sets the value of Readfromfilebytecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyReadfromfilebytecount(value uint64) (err error) {
	return instance.SetProperty("Readfromfilebytecount", (value))
}

// GetReadfromfilebytecount gets the value of Readfromfilebytecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyReadfromfilebytecount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Readfromfilebytecount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetReadfromfilebytelatencymicroseconds sets the value of Readfromfilebytelatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyReadfromfilebytelatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Readfromfilebytelatencymicroseconds", (value))
}

// GetReadfromfilebytelatencymicroseconds gets the value of Readfromfilebytelatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyReadfromfilebytelatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Readfromfilebytelatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetReadfromfilecount sets the value of Readfromfilecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyReadfromfilecount(value uint64) (err error) {
	return instance.SetProperty("Readfromfilecount", (value))
}

// GetReadfromfilecount gets the value of Readfromfilecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyReadfromfilecount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Readfromfilecount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetReadfromstoragebytecount sets the value of Readfromstoragebytecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyReadfromstoragebytecount(value uint64) (err error) {
	return instance.SetProperty("Readfromstoragebytecount", (value))
}

// GetReadfromstoragebytecount gets the value of Readfromstoragebytecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyReadfromstoragebytecount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Readfromstoragebytecount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetReadfromstoragebytelatencymicroseconds sets the value of Readfromstoragebytelatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyReadfromstoragebytelatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Readfromstoragebytelatencymicroseconds", (value))
}

// GetReadfromstoragebytelatencymicroseconds gets the value of Readfromstoragebytelatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyReadfromstoragebytelatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Readfromstoragebytelatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetReadfromstoragecount sets the value of Readfromstoragecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyReadfromstoragecount(value uint64) (err error) {
	return instance.SetProperty("Readfromstoragecount", (value))
}

// GetReadfromstoragecount gets the value of Readfromstoragecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyReadfromstoragecount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Readfromstoragecount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetReconnectlatencymicroseconds sets the value of Reconnectlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyReconnectlatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Reconnectlatencymicroseconds", (value))
}

// GetReconnectlatencymicroseconds gets the value of Reconnectlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyReconnectlatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Reconnectlatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRemoveoperationcount sets the value of Removeoperationcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyRemoveoperationcount(value uint64) (err error) {
	return instance.SetProperty("Removeoperationcount", (value))
}

// GetRemoveoperationcount gets the value of Removeoperationcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyRemoveoperationcount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Removeoperationcount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRemoveoperationlatencymicroseconds sets the value of Removeoperationlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyRemoveoperationlatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Removeoperationlatencymicroseconds", (value))
}

// GetRemoveoperationlatencymicroseconds gets the value of Removeoperationlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyRemoveoperationlatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Removeoperationlatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSectorsize sets the value of Sectorsize for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertySectorsize(value uint64) (err error) {
	return instance.SetProperty("Sectorsize", (value))
}

// GetSectorsize gets the value of Sectorsize for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertySectorsize() (value uint64, err error) {
	retValue, err := instance.GetProperty("Sectorsize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSetoperationcount sets the value of Setoperationcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertySetoperationcount(value uint64) (err error) {
	return instance.SetProperty("Setoperationcount", (value))
}

// GetSetoperationcount gets the value of Setoperationcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertySetoperationcount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Setoperationcount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSetoperationlatencymicroseconds sets the value of Setoperationlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertySetoperationlatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Setoperationlatencymicroseconds", (value))
}

// GetSetoperationlatencymicroseconds gets the value of Setoperationlatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertySetoperationlatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Setoperationlatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetStoragelockacquirelatencymicroseconds sets the value of Storagelockacquirelatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyStoragelockacquirelatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Storagelockacquirelatencymicroseconds", (value))
}

// GetStoragelockacquirelatencymicroseconds gets the value of Storagelockacquirelatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyStoragelockacquirelatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Storagelockacquirelatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetStoragelockcount sets the value of Storagelockcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyStoragelockcount(value uint64) (err error) {
	return instance.SetProperty("Storagelockcount", (value))
}

// GetStoragelockcount gets the value of Storagelockcount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyStoragelockcount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Storagelockcount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetStoragelockreleaselatencymicroseconds sets the value of Storagelockreleaselatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyStoragelockreleaselatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Storagelockreleaselatencymicroseconds", (value))
}

// GetStoragelockreleaselatencymicroseconds gets the value of Storagelockreleaselatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyStoragelockreleaselatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Storagelockreleaselatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTabledatasizeinbytes sets the value of Tabledatasizeinbytes for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyTabledatasizeinbytes(value uint64) (err error) {
	return instance.SetProperty("Tabledatasizeinbytes", (value))
}

// GetTabledatasizeinbytes gets the value of Tabledatasizeinbytes for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyTabledatasizeinbytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("Tabledatasizeinbytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetWritetofilebytecount sets the value of Writetofilebytecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyWritetofilebytecount(value uint64) (err error) {
	return instance.SetProperty("Writetofilebytecount", (value))
}

// GetWritetofilebytecount gets the value of Writetofilebytecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyWritetofilebytecount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Writetofilebytecount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetWritetofilebytelatencymicroseconds sets the value of Writetofilebytelatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyWritetofilebytelatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Writetofilebytelatencymicroseconds", (value))
}

// GetWritetofilebytelatencymicroseconds gets the value of Writetofilebytelatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyWritetofilebytelatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Writetofilebytelatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetWritetofilecount sets the value of Writetofilecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyWritetofilecount(value uint64) (err error) {
	return instance.SetProperty("Writetofilecount", (value))
}

// GetWritetofilecount gets the value of Writetofilecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyWritetofilecount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Writetofilecount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetWritetostoragebytecount sets the value of Writetostoragebytecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyWritetostoragebytecount(value uint64) (err error) {
	return instance.SetProperty("Writetostoragebytecount", (value))
}

// GetWritetostoragebytecount gets the value of Writetostoragebytecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyWritetostoragebytecount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Writetostoragebytecount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetWritetostoragebytelatencymicroseconds sets the value of Writetostoragebytelatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyWritetostoragebytelatencymicroseconds(value uint64) (err error) {
	return instance.SetProperty("Writetostoragebytelatencymicroseconds", (value))
}

// GetWritetostoragebytelatencymicroseconds gets the value of Writetostoragebytelatencymicroseconds for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyWritetostoragebytelatencymicroseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("Writetostoragebytelatencymicroseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetWritetostoragecount sets the value of Writetostoragecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) SetPropertyWritetostoragecount(value uint64) (err error) {
	return instance.SetProperty("Writetostoragecount", (value))
}

// GetWritetostoragecount gets the value of Writetostoragecount for the instance
func (instance *Win32_PerfFormattedData_HyperVStorageStats_HyperVDataStore) GetPropertyWritetostoragecount() (value uint64, err error) {
	retValue, err := instance.GetProperty("Writetostoragecount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
