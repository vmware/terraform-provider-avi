/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviStringGroup() *schema.Resource {
	return &schema.Resource{
		Read:   ResourceAviStringGroupRead,
		Schema: ResourceStringGroupSchema(),
	}
}
