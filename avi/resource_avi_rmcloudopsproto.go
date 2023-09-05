// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceRmCloudOpsProtoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"last_queried_se_creation_limit": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"pending_se_creation_count": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"pending_vnic_op_count": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviRmCloudOpsProto() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviRmCloudOpsProtoCreate,
		Read:   ResourceAviRmCloudOpsProtoRead,
		Update: resourceAviRmCloudOpsProtoUpdate,
		Delete: resourceAviRmCloudOpsProtoDelete,
		Schema: ResourceRmCloudOpsProtoSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceRmCloudOpsProtoImporter,
		},
	}
}

func ResourceRmCloudOpsProtoImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceRmCloudOpsProtoSchema()
	return ResourceImporter(d, m, "rmcloudopsproto", s)
}

func ResourceAviRmCloudOpsProtoRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRmCloudOpsProtoSchema()
	err := APIRead(d, meta, "rmcloudopsproto", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviRmCloudOpsProtoCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRmCloudOpsProtoSchema()
	err := APICreateOrUpdate(d, meta, "rmcloudopsproto", s)
	if err == nil {
		err = ResourceAviRmCloudOpsProtoRead(d, meta)
	}
	return err
}

func resourceAviRmCloudOpsProtoUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRmCloudOpsProtoSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "rmcloudopsproto", s)
	if err == nil {
		err = ResourceAviRmCloudOpsProtoRead(d, meta)
	}
	return err
}

func resourceAviRmCloudOpsProtoDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "rmcloudopsproto")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
