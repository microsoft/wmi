// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.DeliveryOptimization
//////////////////////////////////////////////
package deliveryoptimization

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DeliveryOptimizationFilePeerInfo struct
type MSFT_DeliveryOptimizationFilePeerInfo struct {
	*cim.WmiInstance

	//
	BytesReceived []uint64

	//
	BytesSent []uint64

	//
	ConnectionEstablished []bool

	//
	DownloadRates []uint32

	//
	FileId string

	//
	IPs []string

	//
	PeerTypes []uint8

	//
	UploadRates []uint32
}

func NewMSFT_DeliveryOptimizationFilePeerInfoEx1(instance *cim.WmiInstance) (newInstance *MSFT_DeliveryOptimizationFilePeerInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DeliveryOptimizationFilePeerInfo{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DeliveryOptimizationFilePeerInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DeliveryOptimizationFilePeerInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DeliveryOptimizationFilePeerInfo{
		WmiInstance: tmp,
	}
	return
}

// SetBytesReceived sets the value of BytesReceived for the instance
func (instance *MSFT_DeliveryOptimizationFilePeerInfo) SetPropertyBytesReceived(value []uint64) (err error) {
	return instance.SetProperty("BytesReceived", value)
}

// GetBytesReceived gets the value of BytesReceived for the instance
func (instance *MSFT_DeliveryOptimizationFilePeerInfo) GetPropertyBytesReceived() (value []uint64, err error) {
	retValue, err := instance.GetProperty("BytesReceived")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesSent sets the value of BytesSent for the instance
func (instance *MSFT_DeliveryOptimizationFilePeerInfo) SetPropertyBytesSent(value []uint64) (err error) {
	return instance.SetProperty("BytesSent", value)
}

// GetBytesSent gets the value of BytesSent for the instance
func (instance *MSFT_DeliveryOptimizationFilePeerInfo) GetPropertyBytesSent() (value []uint64, err error) {
	retValue, err := instance.GetProperty("BytesSent")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectionEstablished sets the value of ConnectionEstablished for the instance
func (instance *MSFT_DeliveryOptimizationFilePeerInfo) SetPropertyConnectionEstablished(value []bool) (err error) {
	return instance.SetProperty("ConnectionEstablished", value)
}

// GetConnectionEstablished gets the value of ConnectionEstablished for the instance
func (instance *MSFT_DeliveryOptimizationFilePeerInfo) GetPropertyConnectionEstablished() (value []bool, err error) {
	retValue, err := instance.GetProperty("ConnectionEstablished")
	if err != nil {
		return
	}
	value, ok := retValue.([]bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDownloadRates sets the value of DownloadRates for the instance
func (instance *MSFT_DeliveryOptimizationFilePeerInfo) SetPropertyDownloadRates(value []uint32) (err error) {
	return instance.SetProperty("DownloadRates", value)
}

// GetDownloadRates gets the value of DownloadRates for the instance
func (instance *MSFT_DeliveryOptimizationFilePeerInfo) GetPropertyDownloadRates() (value []uint32, err error) {
	retValue, err := instance.GetProperty("DownloadRates")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileId sets the value of FileId for the instance
func (instance *MSFT_DeliveryOptimizationFilePeerInfo) SetPropertyFileId(value string) (err error) {
	return instance.SetProperty("FileId", value)
}

// GetFileId gets the value of FileId for the instance
func (instance *MSFT_DeliveryOptimizationFilePeerInfo) GetPropertyFileId() (value string, err error) {
	retValue, err := instance.GetProperty("FileId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPs sets the value of IPs for the instance
func (instance *MSFT_DeliveryOptimizationFilePeerInfo) SetPropertyIPs(value []string) (err error) {
	return instance.SetProperty("IPs", value)
}

// GetIPs gets the value of IPs for the instance
func (instance *MSFT_DeliveryOptimizationFilePeerInfo) GetPropertyIPs() (value []string, err error) {
	retValue, err := instance.GetProperty("IPs")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPeerTypes sets the value of PeerTypes for the instance
func (instance *MSFT_DeliveryOptimizationFilePeerInfo) SetPropertyPeerTypes(value []uint8) (err error) {
	return instance.SetProperty("PeerTypes", value)
}

// GetPeerTypes gets the value of PeerTypes for the instance
func (instance *MSFT_DeliveryOptimizationFilePeerInfo) GetPropertyPeerTypes() (value []uint8, err error) {
	retValue, err := instance.GetProperty("PeerTypes")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUploadRates sets the value of UploadRates for the instance
func (instance *MSFT_DeliveryOptimizationFilePeerInfo) SetPropertyUploadRates(value []uint32) (err error) {
	return instance.SetProperty("UploadRates", value)
}

// GetUploadRates gets the value of UploadRates for the instance
func (instance *MSFT_DeliveryOptimizationFilePeerInfo) GetPropertyUploadRates() (value []uint32, err error) {
	retValue, err := instance.GetProperty("UploadRates")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
