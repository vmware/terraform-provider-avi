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

func ResourceApplicationPersistenceProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_cookie_persistence_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAppCookiePersistenceProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"hdr_persistence_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHdrPersistenceProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"http_cookie_persistence_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHttpCookiePersistenceProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"ip_persistence_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceIPPersistenceProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"is_federated": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"persistence_type": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"server_hm_down_recovery": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "HM_DOWN_PICK_NEW_SERVER",
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

func resourceAviApplicationPersistenceProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviApplicationPersistenceProfileCreate,
		Read:   ResourceAviApplicationPersistenceProfileRead,
		Update: resourceAviApplicationPersistenceProfileUpdate,
		Delete: resourceAviApplicationPersistenceProfileDelete,
		Schema: ResourceApplicationPersistenceProfileSchema(),
	}
}

func ResourceAviApplicationPersistenceProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationPersistenceProfileSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/applicationpersistenceprofile/" + uuid.(string)
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

func resourceAviApplicationPersistenceProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationPersistenceProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "applicationpersistenceprofile", s)
	if err == nil {
		err = ResourceAviApplicationPersistenceProfileRead(d, meta)
	}
	return err
}

func resourceAviApplicationPersistenceProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationPersistenceProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "applicationpersistenceprofile", s)
	if err == nil {
		err = ResourceAviApplicationPersistenceProfileRead(d, meta)
	}
	return err
}

func resourceAviApplicationPersistenceProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "applicationpersistenceprofile"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviApplicationPersistenceProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
