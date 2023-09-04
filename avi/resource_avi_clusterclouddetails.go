// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"

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
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
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
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "clusterclouddetails")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
