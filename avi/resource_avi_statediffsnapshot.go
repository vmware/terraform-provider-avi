// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceStatediffSnapshotSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"gslb_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"gslb_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"pool_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"pool_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"post_snapshot": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcepostsnapshotSchema(),
		},
		"pre_snapshot": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcepresnapshotSchema(),
		},
		"se_group_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"se_group_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"se_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"se_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"snapshot_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"statediff_operation_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
		"vs_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vs_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviStatediffSnapshot() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviStatediffSnapshotCreate,
		Read:   ResourceAviStatediffSnapshotRead,
		Update: resourceAviStatediffSnapshotUpdate,
		Delete: resourceAviStatediffSnapshotDelete,
		Schema: ResourceStatediffSnapshotSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceStatediffSnapshotImporter,
		},
	}
}

func ResourceStatediffSnapshotImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceStatediffSnapshotSchema()
	return ResourceImporter(d, m, "statediffsnapshot", s)
}

func ResourceAviStatediffSnapshotRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceStatediffSnapshotSchema()
	err := APIRead(d, meta, "statediffsnapshot", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviStatediffSnapshotCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceStatediffSnapshotSchema()
	err := APICreateOrUpdate(d, meta, "statediffsnapshot", s)
	if err == nil {
		err = ResourceAviStatediffSnapshotRead(d, meta)
	}
	return err
}

func resourceAviStatediffSnapshotUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceStatediffSnapshotSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "statediffsnapshot", s)
	if err == nil {
		err = ResourceAviStatediffSnapshotRead(d, meta)
	}
	return err
}

func resourceAviStatediffSnapshotDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "statediffsnapshot"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviStatediffSnapshotDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
