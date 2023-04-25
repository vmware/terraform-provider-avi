/***************************************************************************
 * ========================================================================
 * Copyright 2023 VMware, Inc. All rights reserved. VMware Confidential
 * ========================================================================
 */

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/vmware/terraform-provider-avi/avi"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: avi.Provider})
}
