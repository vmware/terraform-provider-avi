// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceClusterCloudDetailsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"azure_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAzureClusterInfoSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviClusterCloudDetails() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviClusterCloudDetailsCreate,
		Read:   ResourceAviClusterCloudDetailsRead,
		Update: resourceAviClusterCloudDetailsUpdate,
		Delete: resourceAviClusterCloudDetailsDelete,
		Schema: ResourceClusterCloudDetailsSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceClusterCloudDetailsImporter,
		},
	}
}

func ResourceClusterCloudDetailsImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceClusterCloudDetailsSchema()
	return ResourceImporter(d, m, "clusterclouddetails", s)
}

func ResourceAviClusterCloudDetailsRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClusterCloudDetailsSchema()
	err := APIRead(d, meta, "clusterclouddetails", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviClusterCloudDetailsCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClusterCloudDetailsSchema()
	err := APICreateOrUpdate(d, meta, "clusterclouddetails", s)
	if err == nil {
		err = ResourceAviClusterCloudDetailsRead(d, meta)
	}
	return err
}

func resourceAviClusterCloudDetailsUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClusterCloudDetailsSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "clusterclouddetails", s)
	if err == nil {
		err = ResourceAviClusterCloudDetailsRead(d, meta)
	}
	return err
}

func resourceAviClusterCloudDetailsDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "clusterclouddetails"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviClusterCloudDetailsDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
