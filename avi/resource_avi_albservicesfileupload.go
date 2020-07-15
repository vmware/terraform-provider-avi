/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
)

func ResourceALBServicesFileUploadSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"case_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"file_path": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"s3_directory": {
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

func resourceAviALBServicesFileUpload() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviALBServicesFileUploadCreate,
		Read:   ResourceAviALBServicesFileUploadRead,
		Update: resourceAviALBServicesFileUploadUpdate,
		Delete: resourceAviALBServicesFileUploadDelete,
		Schema: ResourceALBServicesFileUploadSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceALBServicesFileUploadImporter,
		},
	}
}

func ResourceALBServicesFileUploadImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceALBServicesFileUploadSchema()
	return ResourceImporter(d, m, "albservicesfileupload", s)
}

func ResourceAviALBServicesFileUploadRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesFileUploadSchema()
	err := ApiRead(d, meta, "albservicesfileupload", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviALBServicesFileUploadCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesFileUploadSchema()
	err := ApiCreateOrUpdate(d, meta, "albservicesfileupload", s)
	if err == nil {
		err = ResourceAviALBServicesFileUploadRead(d, meta)
	}
	return err
}

func resourceAviALBServicesFileUploadUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesFileUploadSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "albservicesfileupload", s)
	if err == nil {
		err = ResourceAviALBServicesFileUploadRead(d, meta)
	}
	return err
}

func resourceAviALBServicesFileUploadDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "albservicesfileupload"
	client := meta.(*clients.AviClient)
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviALBServicesFileUploadDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
