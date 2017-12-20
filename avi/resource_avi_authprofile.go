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

func ResourceAuthProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"http": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAuthProfileHTTPClientParamsSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"ldap": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceLdapAuthSettingsSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"saml": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceSamlSettingsSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"tacacs_plus": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceTacacsPlusAuthSettingsSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"type": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviAuthProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAuthProfileCreate,
		Read:   ResourceAviAuthProfileRead,
		Update: resourceAviAuthProfileUpdate,
		Delete: resourceAviAuthProfileDelete,
		Schema: ResourceAuthProfileSchema(),
	}
}

func ResourceAviAuthProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAuthProfileSchema()
	log.Printf("[INFO] ResourceAviAuthProfileRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/authprofile/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	// no need to set the ID
	log.Printf("ResourceAviAuthProfileRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviAuthProfileRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviAuthProfileRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviAuthProfileRead Updated %v\n", d)
	return nil
}

func resourceAviAuthProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAuthProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "authprofile", s)
	log.Printf("[DEBUG] created object %v: %v", "authprofile", d)
	if err == nil {
		err = ResourceAviAuthProfileRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "authprofile", d)
	return err
}

func resourceAviAuthProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAuthProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "authprofile", s)
	if err == nil {
		err = ResourceAviAuthProfileRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "authprofile", d)
	return err
}

func resourceAviAuthProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "authprofile"
	log.Println("[INFO] ResourceAviAuthProfileRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviAuthProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
