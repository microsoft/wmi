// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package host

import (
	"log"

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
		log.Printf("[WMI] Error getting Win32_PnPEntity instances for name [%s] - Error details [%+v]\n", pnpEntityName, err)
		return
	}
	defer rasdcollection.Close()

	entities, err = pnp.NewPnpEntityCollection(rasdcollection)
	if err != nil {
		log.Printf("[WMI] Error getting new PnpEntityCollection for rasdcollection [%s] - Error details [%+v]\n", rasdcollection, err)
		return
	}

	if len(entities) == 0 {
		err = errors.Wrapf(errors.NotFound, "Unable to find PnP entity for name [%s]", pnpEntityName)
		log.Printf("[WMI] Found zero PnP entities for name [%s]", pnpEntityName)
		return
	}

	log.Printf("[WMI] Found PnP entities [%s] for name [%s]", entities, pnpEntityName)
	return
}
