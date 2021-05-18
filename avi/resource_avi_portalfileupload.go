// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourcePortalFileUploadSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"case_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"error": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"file_path": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"s3_directory": {
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

func resourceAviPortalFileUpload() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPortalFileUploadCreate,
		Read:   ResourceAviPortalFileUploadRead,
		Update: resourceAviPortalFileUploadUpdate,
		Delete: resourceAviPortalFileUploadDelete,
		Schema: ResourcePortalFileUploadSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourcePortalFileUploadImporter,
		},
	}
}

func ResourcePortalFileUploadImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourcePortalFileUploadSchema()
	return ResourceImporter(d, m, "portalfileupload", s)
}

func ResourceAviPortalFileUploadRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePortalFileUploadSchema()
	err := APIRead(d, meta, "portalfileupload", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviPortalFileUploadCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePortalFileUploadSchema()
	err := APICreateOrUpdate(d, meta, "portalfileupload", s)
	if err == nil {
		err = ResourceAviPortalFileUploadRead(d, meta)
	}
	return err
}

func resourceAviPortalFileUploadUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePortalFileUploadSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "portalfileupload", s)
	if err == nil {
		err = ResourceAviPortalFileUploadRead(d, meta)
	}
	return err
}

func resourceAviPortalFileUploadDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "portalfileupload"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviPortalFileUploadDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
