// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_FileDownloadManager struct
type MSFT_FileDownloadManager struct {
	*OMI_ConfigurationDownloadManager

	//
	Credential MSFT_Credential

	//
	SourcePath string
}

func NewMSFT_FileDownloadManagerEx1(instance *cim.WmiInstance) (newInstance *MSFT_FileDownloadManager, err error) {
	tmp, err := NewOMI_ConfigurationDownloadManagerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_FileDownloadManager{
		OMI_ConfigurationDownloadManager: tmp,
	}
	return
}

func NewMSFT_FileDownloadManagerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_FileDownloadManager, err error) {
	tmp, err := NewOMI_ConfigurationDownloadManagerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_FileDownloadManager{
		OMI_ConfigurationDownloadManager: tmp,
	}
	return
}

// SetCredential sets the value of Credential for the instance
func (instance *MSFT_FileDownloadManager) SetPropertyCredential(value MSFT_Credential) (err error) {
	return instance.SetProperty("Credential", value)
}

// GetCredential gets the value of Credential for the instance
func (instance *MSFT_FileDownloadManager) GetPropertyCredential() (value MSFT_Credential, err error) {
	retValue, err := instance.GetProperty("Credential")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Credential)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourcePath sets the value of SourcePath for the instance
func (instance *MSFT_FileDownloadManager) SetPropertySourcePath(value string) (err error) {
	return instance.SetProperty("SourcePath", value)
}

// GetSourcePath gets the value of SourcePath for the instance
func (instance *MSFT_FileDownloadManager) GetPropertySourcePath() (value string, err error) {
	retValue, err := instance.GetProperty("SourcePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
