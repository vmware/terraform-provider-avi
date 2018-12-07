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
	"time"
)

func ResourceCertificateManagementProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"script_params": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomParamsSchema(),
		},
		"script_path": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"uuid": &schema.Schema{
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
	err := ApiRead(d, meta, "certificatemanagementprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviCertificateManagementProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCertificateManagementProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "certificatemanagementprofile", s)
	if err == nil {
		err = ResourceAviCertificateManagementProfileRead(d, meta)
	}
	return err
}

func resourceAviCertificateManagementProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCertificateManagementProfileSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "certificatemanagementprofile", s)
	if err == nil {
		err = ResourceAviCertificateManagementProfileRead(d, meta)
	}
	return err
}

func resourceAviCertificateManagementProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "certificatemanagementprofile"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
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
