// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCertificateManagementProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"script_params": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomParamsSchema(),
		},
		"script_path": {
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

func resourceAviCertificateManagementProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCertificateManagementProfileCreate,
		Read:   ResourceAviCertificateManagementProfileRead,
		Update: resourceAviCertificateManagementProfileUpdate,
		Delete: resourceAviCertificateManagementProfileDelete,
		Schema: ResourceCertificateManagementProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceCertificateManagementProfileImporter,
		},
	}
}

func ResourceCertificateManagementProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceCertificateManagementProfileSchema()
	return ResourceImporter(d, m, "certificatemanagementprofile", s)
}

func ResourceAviCertificateManagementProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCertificateManagementProfileSchema()
	err := APIRead(d, meta, "certificatemanagementprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviCertificateManagementProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCertificateManagementProfileSchema()
	err := APICreateOrUpdate(d, meta, "certificatemanagementprofile", s)
	if err == nil {
		err = ResourceAviCertificateManagementProfileRead(d, meta)
	}
	return err
}

func resourceAviCertificateManagementProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCertificateManagementProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "certificatemanagementprofile", s)
	if err == nil {
		err = ResourceAviCertificateManagementProfileRead(d, meta)
	}
	return err
}

func resourceAviCertificateManagementProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "certificatemanagementprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviCertificateManagementProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
