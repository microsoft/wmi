// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_WmiOffload struct
type MSNdis_WmiOffload struct {
	*MSNdis

	//
	Checksum MSNdis_WmiTcpIpChecksumOffload

	//
	Flags uint32

	//
	Header MSNdis_ObjectHeader

	//
	IPsecV1 MSNdis_WmiIPSecOffloadV1

	//
	LsoV1 MSNdis_WmiTcpLargeSendOffloadV1

	//
	LsoV2 MSNdis_WmiTcpLargeSendOffloadV2
}

func NewMSNdis_WmiOffloadEx1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiOffload, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiOffload{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_WmiOffloadEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_WmiOffload, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_WmiOffload{
		MSNdis: tmp,
	}
	return
}

// SetChecksum sets the value of Checksum for the instance
func (instance *MSNdis_WmiOffload) SetPropertyChecksum(value MSNdis_WmiTcpIpChecksumOffload) (err error) {
	return instance.SetProperty("Checksum", (value))
}

// GetChecksum gets the value of Checksum for the instance
func (instance *MSNdis_WmiOffload) GetPropertyChecksum() (value MSNdis_WmiTcpIpChecksumOffload, err error) {
	retValue, err := instance.GetProperty("Checksum")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_WmiTcpIpChecksumOffload)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_WmiTcpIpChecksumOffload is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_WmiTcpIpChecksumOffload(valuetmp)

	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MSNdis_WmiOffload) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSNdis_WmiOffload) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_WmiOffload) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_WmiOffload) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
	retValue, err := instance.GetProperty("Header")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_ObjectHeader)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_ObjectHeader is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_ObjectHeader(valuetmp)

	return
}

// SetIPsecV1 sets the value of IPsecV1 for the instance
func (instance *MSNdis_WmiOffload) SetPropertyIPsecV1(value MSNdis_WmiIPSecOffloadV1) (err error) {
	return instance.SetProperty("IPsecV1", (value))
}

// GetIPsecV1 gets the value of IPsecV1 for the instance
func (instance *MSNdis_WmiOffload) GetPropertyIPsecV1() (value MSNdis_WmiIPSecOffloadV1, err error) {
	retValue, err := instance.GetProperty("IPsecV1")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_WmiIPSecOffloadV1)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_WmiIPSecOffloadV1 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_WmiIPSecOffloadV1(valuetmp)

	return
}

// SetLsoV1 sets the value of LsoV1 for the instance
func (instance *MSNdis_WmiOffload) SetPropertyLsoV1(value MSNdis_WmiTcpLargeSendOffloadV1) (err error) {
	return instance.SetProperty("LsoV1", (value))
}

// GetLsoV1 gets the value of LsoV1 for the instance
func (instance *MSNdis_WmiOffload) GetPropertyLsoV1() (value MSNdis_WmiTcpLargeSendOffloadV1, err error) {
	retValue, err := instance.GetProperty("LsoV1")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_WmiTcpLargeSendOffloadV1)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_WmiTcpLargeSendOffloadV1 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_WmiTcpLargeSendOffloadV1(valuetmp)

	return
}

// SetLsoV2 sets the value of LsoV2 for the instance
func (instance *MSNdis_WmiOffload) SetPropertyLsoV2(value MSNdis_WmiTcpLargeSendOffloadV2) (err error) {
	return instance.SetProperty("LsoV2", (value))
}

// GetLsoV2 gets the value of LsoV2 for the instance
func (instance *MSNdis_WmiOffload) GetPropertyLsoV2() (value MSNdis_WmiTcpLargeSendOffloadV2, err error) {
	retValue, err := instance.GetProperty("LsoV2")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_WmiTcpLargeSendOffloadV2)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_WmiTcpLargeSendOffloadV2 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_WmiTcpLargeSendOffloadV2(valuetmp)

	return
}
