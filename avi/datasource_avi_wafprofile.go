/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviWafProfile() *schema.Resource {
	return &schema.Resource{
		Read:   ResourceAviWafProfileRead,
		Schema: ResourceWafProfileSchema(),
	}
}
