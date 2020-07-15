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

func ResourceALBServicesConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"asset_contact": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceALBServicesUserSchema(),
		},
		"feature_opt_in_status": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcePortalFeatureOptInSchema(),
		},
		"ip_reputation_file_object_expiry_duration": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  3,
		},
		"ip_reputation_sync_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"polling_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"portal_url": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"proactive_support_defaults": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceProactiveSupportDefaultsSchema(),
		},
		"split_proxy_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceProxyConfigurationSchema(),
		},
		"use_system_proxy": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviALBServicesConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviALBServicesConfigCreate,
		Read:   ResourceAviALBServicesConfigRead,
		Update: resourceAviALBServicesConfigUpdate,
		Delete: resourceAviALBServicesConfigDelete,
		Schema: ResourceALBServicesConfigSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceALBServicesConfigImporter,
		},
	}
}

func ResourceALBServicesConfigImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceALBServicesConfigSchema()
	return ResourceImporter(d, m, "albservicesconfig", s)
}

func ResourceAviALBServicesConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesConfigSchema()
	err := ApiRead(d, meta, "albservicesconfig", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviALBServicesConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesConfigSchema()
	err := ApiCreateOrUpdate(d, meta, "albservicesconfig", s)
	if err == nil {
		err = ResourceAviALBServicesConfigRead(d, meta)
	}
	return err
}

func resourceAviALBServicesConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesConfigSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "albservicesconfig", s)
	if err == nil {
		err = ResourceAviALBServicesConfigRead(d, meta)
	}
	return err
}

func resourceAviALBServicesConfigDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "albservicesconfig"
	client := meta.(*clients.AviClient)
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviALBServicesConfigDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
