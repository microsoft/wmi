// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cim

import (
	"github.com/go-ole/go-ole"
	"strings"
)

func GetVariantValue(rawValue *ole.VARIANT) (interface{}, error) {
	values, err := GetVariantValues(rawValue)
	if err != nil {
		panic("Couldn't retreive a variant value.")
	}

	if len(values) != 1 {
		panic("Returned an unexpected number of variants")
	}

	return values[0], nil
}

func GetVariantValues(rawValue *ole.VARIANT) ([]interface{}, error) {

	var values []interface{}
	array := rawValue.ToArray()

	if array == nil {
		// Not an array
		values = append(values, rawValue.Value())
	} else {
		values = array.ToValueArray()
	}

	return values, nil
}

func EscapeQueryValue(rawString string) string {
	// Double the backslash character as per required by the "WHERE" WMI clause
	// Documentation: https://docs.microsoft.com/en-us/windows/win32/wmisdk/where-clause
	// Interestingly, double quotes don't seem to need escaping.
	return strings.ReplaceAll(strings.ReplaceAll(rawString, "\\", "\\\\"), "'", "\\'")
}

func FindStringInSlice(stringList []string, value string) (int, bool) {
	for i, item := range stringList {
		if item == value {
			return i, true
		}
	}
	return -1, false
}
