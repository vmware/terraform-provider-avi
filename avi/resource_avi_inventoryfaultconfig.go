// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceInventoryFaultConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"controller_faults": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceControllerFaultsSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"serviceengine_faults": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceServiceengineFaultsSchema(),
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
		"virtualservice_faults": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceVirtualserviceFaultsSchema(),
		},
	}
}

func resourceAviInventoryFaultConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviInventoryFaultConfigCreate,
		Read:   ResourceAviInventoryFaultConfigRead,
		Update: resourceAviInventoryFaultConfigUpdate,
		Delete: resourceAviInventoryFaultConfigDelete,
		Schema: ResourceInventoryFaultConfigSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceInventoryFaultConfigImporter,
		},
	}
}

func ResourceInventoryFaultConfigImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceInventoryFaultConfigSchema()
	return ResourceImporter(d, m, "inventoryfaultconfig", s)
}

func ResourceAviInventoryFaultConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceInventoryFaultConfigSchema()
	err := APIRead(d, meta, "inventoryfaultconfig", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviInventoryFaultConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceInventoryFaultConfigSchema()
	err := APICreateOrUpdate(d, meta, "inventoryfaultconfig", s)
	if err == nil {
		err = ResourceAviInventoryFaultConfigRead(d, meta)
	}
	return err
}

func resourceAviInventoryFaultConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceInventoryFaultConfigSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "inventoryfaultconfig", s)
	if err == nil {
		err = ResourceAviInventoryFaultConfigRead(d, meta)
	}
	return err
}

func resourceAviInventoryFaultConfigDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "inventoryfaultconfig"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviInventoryFaultConfigDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
