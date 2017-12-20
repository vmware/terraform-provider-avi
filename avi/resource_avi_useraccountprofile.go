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

func ResourceUserAccountProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_lock_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  30,
		},
		"credentials_timeout_threshold": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  180,
		},
		"max_concurrent_sessions": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"max_login_failure_count": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  3,
		},
		"max_password_history_count": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		"name": &schema.Schema{
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

func resourceAviUserAccountProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviUserAccountProfileCreate,
		Read:   ResourceAviUserAccountProfileRead,
		Update: resourceAviUserAccountProfileUpdate,
		Delete: resourceAviUserAccountProfileDelete,
		Schema: ResourceUserAccountProfileSchema(),
	}
}

func ResourceAviUserAccountProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserAccountProfileSchema()
	log.Printf("[INFO] ResourceAviUserAccountProfileRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/useraccountprofile/" + uuid.(string)
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
	log.Printf("ResourceAviUserAccountProfileRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviUserAccountProfileRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviUserAccountProfileRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviUserAccountProfileRead Updated %v\n", d)
	return nil
}

func resourceAviUserAccountProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserAccountProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "useraccountprofile", s)
	log.Printf("[DEBUG] created object %v: %v", "useraccountprofile", d)
	if err == nil {
		err = ResourceAviUserAccountProfileRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "useraccountprofile", d)
	return err
}

func resourceAviUserAccountProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserAccountProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "useraccountprofile", s)
	if err == nil {
		err = ResourceAviUserAccountProfileRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "useraccountprofile", d)
	return err
}

func resourceAviUserAccountProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "useraccountprofile"
	log.Println("[INFO] ResourceAviUserAccountProfileRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviUserAccountProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
