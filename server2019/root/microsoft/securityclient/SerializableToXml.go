// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.SecurityClient
//////////////////////////////////////////////
package securityclient

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// SerializableToXml struct
type SerializableToXml struct {
	cim.WmiInstance

	// This is an all-in-one state data that uses an XML format with a standard CIM DTD schema
	PackedXml string

	// Schema version (major, minor, build, revision)
	SchemaVersion string
}

// SetPackedXml sets the value of PackedXml for the instance
func (instance *SerializableToXml) SetPropertyPackedXml(value string) (err error) {
	return instance.SetProperty("PackedXml", value)
}

// GetPackedXml gets the value of PackedXml for the instance
func (instance *SerializableToXml) GetPropertyPackedXml() (value string, err error) {
	retValue, err := instance.GetProperty("PackedXml")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSchemaVersion sets the value of SchemaVersion for the instance
func (instance *SerializableToXml) SetPropertySchemaVersion(value string) (err error) {
	return instance.SetProperty("SchemaVersion", value)
}

// GetSchemaVersion gets the value of SchemaVersion for the instance
func (instance *SerializableToXml) GetPropertySchemaVersion() (value string, err error) {
	retValue, err := instance.GetProperty("SchemaVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
