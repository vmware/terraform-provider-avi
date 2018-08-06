package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviServer() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviServerRead,
		Schema: map[string]*schema.Schema{
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "V4",
			},
			"autoscaling_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"external_orchestration_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGeoLocationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"nw_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prst_hdr_val": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rewrite_host_header": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vm_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
