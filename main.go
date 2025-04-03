/***************************************************************************
 * ========================================================================
 * Copyright (c) 2025 Broadcom Inc. and/or its subsidiaries. All Rights Reserved. Broadcom Confidential.
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
