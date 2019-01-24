/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviSeProperties() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSePropertiesRead,
		Schema: map[string]*schema.Schema{
			"se_agent_properties": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentPropertiesSchema(),
			},
			"se_bootup_properties": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeBootupPropertiesSchema(),
			},
			"se_runtime_properties": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeRuntimePropertiesSchema(),
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
