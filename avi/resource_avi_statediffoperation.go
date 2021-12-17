// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceStatediffOperationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"events": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceStatediffEventSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"node_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"operation": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"phase": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"status": {
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
	}
}

func resourceAviStatediffOperation() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviStatediffOperationCreate,
		Read:   ResourceAviStatediffOperationRead,
		Update: resourceAviStatediffOperationUpdate,
		Delete: resourceAviStatediffOperationDelete,
		Schema: ResourceStatediffOperationSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceStatediffOperationImporter,
		},
	}
}

func ResourceStatediffOperationImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceStatediffOperationSchema()
	return ResourceImporter(d, m, "statediffoperation", s)
}

func ResourceAviStatediffOperationRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceStatediffOperationSchema()
	err := APIRead(d, meta, "statediffoperation", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviStatediffOperationCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceStatediffOperationSchema()
	err := APICreateOrUpdate(d, meta, "statediffoperation", s)
	if err == nil {
		err = ResourceAviStatediffOperationRead(d, meta)
	}
	return err
}

func resourceAviStatediffOperationUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceStatediffOperationSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "statediffoperation", s)
	if err == nil {
		err = ResourceAviStatediffOperationRead(d, meta)
	}
	return err
}

func resourceAviStatediffOperationDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "statediffoperation"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviStatediffOperationDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
