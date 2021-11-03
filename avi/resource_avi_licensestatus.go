// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceLicenseStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"saas_status": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSaasLicensingStatusSchema(),
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

func resourceAviLicenseStatus() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviLicenseStatusCreate,
		Read:   ResourceAviLicenseStatusRead,
		Update: resourceAviLicenseStatusUpdate,
		Delete: resourceAviLicenseStatusDelete,
		Schema: ResourceLicenseStatusSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceLicenseStatusImporter,
		},
	}
}

func ResourceLicenseStatusImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceLicenseStatusSchema()
	return ResourceImporter(d, m, "licensestatus", s)
}

func ResourceAviLicenseStatusRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLicenseStatusSchema()
	err := APIRead(d, meta, "licensestatus", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviLicenseStatusCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLicenseStatusSchema()
	err := APICreateOrUpdate(d, meta, "licensestatus", s)
	if err == nil {
		err = ResourceAviLicenseStatusRead(d, meta)
	}
	return err
}

func resourceAviLicenseStatusUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLicenseStatusSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "licensestatus", s)
	if err == nil {
		err = ResourceAviLicenseStatusRead(d, meta)
	}
	return err
}

func resourceAviLicenseStatusDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "licensestatus"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviLicenseStatusDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
