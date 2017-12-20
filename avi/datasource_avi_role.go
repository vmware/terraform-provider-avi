/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviRole() *schema.Resource {
	return &schema.Resource{
		Read:   ResourceAviRoleRead,
		Schema: ResourceRoleSchema(),
	}
}
