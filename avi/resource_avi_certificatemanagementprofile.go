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
	}
}

func ResourceAviCertificateManagementProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCertificateManagementProfileSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/certificatemanagementprofile/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	if _, err := ApiDataToSchema(obj, d, s); err == nil {
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	return nil
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
	err := ApiCreateOrUpdate(d, meta, "certificatemanagementprofile", s)
	if err == nil {
		err = ResourceAviCertificateManagementProfileRead(d, meta)
	}
	return err
}

func resourceAviCertificateManagementProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "certificatemanagementprofile"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviCertificateManagementProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
