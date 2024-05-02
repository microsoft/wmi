// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package host

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/pnp"
)

// GetPnpEntityByName
func GetPnpEntityByName(whost *host.WmiHost, pnpEntityName string) (entities pnp.PnpEntityCollection, err error) {
	query := query.NewWmiQuery("Win32_PnPEntity", "Name", pnpEntityName)
	rasdcollection, err := instance.GetWmiInstancesFromHost(whost, string(constant.CimV2), query)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get Win32_PnPEntity instances for name [%s]", pnpEntityName)
	}
	defer func() {
		if err != nil {
			rasdcollection.Close()
		}
	}()

	entities, err = pnp.NewPnpEntityCollection(rasdcollection)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get new PnpEntityCollection for rasdcollection [%+v]", rasdcollection)
	}

	if len(entities) == 0 {
		err = errors.Wrapf(errors.NotFound, "Unable to find PnP entity for name [%s]", pnpEntityName)
		return
	}

	return
}
