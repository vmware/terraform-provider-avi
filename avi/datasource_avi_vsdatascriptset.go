/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviVSDataScriptSet() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviVSDataScriptSetRead,
		Schema: map[string]*schema.Schema{
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"datascript": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVSDataScriptSchema(),
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipgroup_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_group_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pool_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"protocol_parser_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"string_group_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
