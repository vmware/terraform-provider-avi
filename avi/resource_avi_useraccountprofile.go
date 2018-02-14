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
		Importer: &schema.ResourceImporter{
			State: ResourceUserAccountProfileImporter,
		},
	}
}

func ResourceUserAccountProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceUserAccountProfileSchema()
	return ResourceImporter(d, m, "useraccountprofile", s)
}

func ResourceAviUserAccountProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserAccountProfileSchema()
	err := ApiRead(d, meta, "useraccountprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviUserAccountProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserAccountProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "useraccountprofile", s)
	if err == nil {
		err = ResourceAviUserAccountProfileRead(d, meta)
	}
	return err
}

func resourceAviUserAccountProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserAccountProfileSchema()
	var err error

	err = ApiCreateOrUpdate(d, meta, "useraccountprofile", s)
	if err == nil {
		err = ResourceAviUserAccountProfileRead(d, meta)
	}
	return err
}

func resourceAviUserAccountProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "useraccountprofile"
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
