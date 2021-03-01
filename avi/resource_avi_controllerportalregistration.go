/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceControllerPortalRegistrationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"asset": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceControllerPortalAssetSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"portal_auth": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceControllerPortalAuthSchema(),
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

func resourceAviControllerPortalRegistration() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviControllerPortalRegistrationCreate,
		Read:   ResourceAviControllerPortalRegistrationRead,
		Update: resourceAviControllerPortalRegistrationUpdate,
		Delete: resourceAviControllerPortalRegistrationDelete,
		Schema: ResourceControllerPortalRegistrationSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceControllerPortalRegistrationImporter,
		},
	}
}

func ResourceControllerPortalRegistrationImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceControllerPortalRegistrationSchema()
	return ResourceImporter(d, m, "controllerportalregistration", s)
}

func ResourceAviControllerPortalRegistrationRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerPortalRegistrationSchema()
	err := APIRead(d, meta, "controllerportalregistration", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviControllerPortalRegistrationCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerPortalRegistrationSchema()
	err := APICreateOrUpdate(d, meta, "controllerportalregistration", s)
	if err == nil {
		err = ResourceAviControllerPortalRegistrationRead(d, meta)
	}
	return err
}

func resourceAviControllerPortalRegistrationUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerPortalRegistrationSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "controllerportalregistration", s)
	if err == nil {
		err = ResourceAviControllerPortalRegistrationRead(d, meta)
	}
	return err
}

func resourceAviControllerPortalRegistrationDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "controllerportalregistration"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviControllerPortalRegistrationDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
