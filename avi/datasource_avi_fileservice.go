package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviFileService() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviFileServiceCreate,
		Schema: map[string]*schema.Schema{
			"uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"local_file": {
				Type:     schema.TypeString,
				Required: true,
			},
			//upload flag to state current local file will be uploaded to remote server.
			"upload": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
