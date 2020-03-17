// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root
//////////////////////////////////////////////
package root

// __ListOfEventActiveNamespaces struct
type __ListOfEventActiveNamespaces struct {
	__SystemClass

	//
	Namespaces []string
}

// SetNamespaces sets the value of Namespaces for the instance
func (instance *__ListOfEventActiveNamespaces) SetPropertyNamespaces(value []string) (err error) {
	return instance.SetProperty("Namespaces", value)
}

// GetNamespaces gets the value of Namespaces for the instance
func (instance *__ListOfEventActiveNamespaces) GetPropertyNamespaces() (value []string, err error) {
	retValue, err := instance.GetProperty("Namespaces")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
