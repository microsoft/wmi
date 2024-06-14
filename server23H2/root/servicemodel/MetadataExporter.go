// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MetadataExporter struct
type MetadataExporter struct {
	*cim.WmiInstance

	// The policy version used in metadata retrieved from the service.
	PolicyVersion string
}

func NewMetadataExporterEx1(instance *cim.WmiInstance) (newInstance *MetadataExporter, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MetadataExporter{
		WmiInstance: tmp,
	}
	return
}

func NewMetadataExporterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MetadataExporter, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MetadataExporter{
		WmiInstance: tmp,
	}
	return
}

// SetPolicyVersion sets the value of PolicyVersion for the instance
func (instance *MetadataExporter) SetPropertyPolicyVersion(value string) (err error) {
	return instance.SetProperty("PolicyVersion", (value))
}

// GetPolicyVersion gets the value of PolicyVersion for the instance
func (instance *MetadataExporter) GetPropertyPolicyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyVersion")
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
