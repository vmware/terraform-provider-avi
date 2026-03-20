// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceAkoAmkoClusterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"cluster_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"created_by": {
			Type:     schema.TypeString,
			Required: true,
		},
		"deployment_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAkoAmkoClusterDeploymentInfoSchema(),
		},
		"metadata": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAkoAmkoClusterMetadataSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"version_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAkoAmkoClusterVersionInfoSchema(),
		},
	}
}

func resourceAviAkoAmkoCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAkoAmkoClusterCreate,
		Read:   ResourceAviAkoAmkoClusterRead,
		Update: resourceAviAkoAmkoClusterUpdate,
		Delete: resourceAviAkoAmkoClusterDelete,
		Schema: ResourceAkoAmkoClusterSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceAkoAmkoClusterImporter,
		},
	}
}

func ResourceAkoAmkoClusterImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceAkoAmkoClusterSchema()
	return ResourceImporter(d, m, "akoamkocluster", s)
}

func ResourceAviAkoAmkoClusterRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAkoAmkoClusterSchema()
	err := APIRead(d, meta, "akoamkocluster", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviAkoAmkoClusterCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAkoAmkoClusterSchema()
	err := APICreate(d, meta, "akoamkocluster", s)
	if err == nil {
		err = ResourceAviAkoAmkoClusterRead(d, meta)
	}
	return err
}

func resourceAviAkoAmkoClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAkoAmkoClusterSchema()
	var err error
	err = APIUpdate(d, meta, "akoamkocluster", s)
	if err == nil {
		err = ResourceAviAkoAmkoClusterRead(d, meta)
	}
	return err
}

func resourceAviAkoAmkoClusterDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "akoamkocluster")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
