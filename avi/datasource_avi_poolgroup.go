/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviPoolGroup() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPoolGroupRead,
		Schema: map[string]*schema.Schema{
			"cloud_config_cksum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"deployment_policy_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fail_action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceFailActionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"implicit_priority_labels": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"members": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePoolGroupMemberSchema(),
			},
			"min_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority_labels_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_metadata": &schema.Schema{
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
