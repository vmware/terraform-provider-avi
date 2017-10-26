/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviPool() *schema.Resource {
	return &schema.Resource{
		Read:   ResourceAviPoolRead,
		Schema: ResourcePoolSchema(),
	}
}
