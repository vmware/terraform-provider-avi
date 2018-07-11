package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviFileService() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviFileServiceCreate,
		Schema: map[string]*schema.Schema{
			"uri": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"local_file": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			//upload flag to state current local file will be uploaded to remote server.
			"upload": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
