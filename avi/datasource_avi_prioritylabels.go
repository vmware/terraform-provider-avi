/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviPriorityLabels() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPriorityLabelsRead,
		Schema: map[string]*schema.Schema{
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud",
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"equivalent_labels": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceEquivalentLabelsSchema(),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
