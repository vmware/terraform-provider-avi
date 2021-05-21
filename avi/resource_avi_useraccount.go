/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceUserAccountSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"username": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"old_password": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"password": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"local": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"full_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"email": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func resourceAviUserAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviUserAccountCreate,
		Read:   ResourceAviUserAccountRead,
		Update: resourceAviUserAccountUpdate,
		Delete: resourceAviUserAccountDelete,
		Schema: ResourceUserAccountSchema(),
	}
}

func ResourceAviUserAccountRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAviUserAccountCreate(d *schema.ResourceData, meta interface{}) error {
	if strings.Compare(d.Get("old_password").(string), d.Get("password").(string)) == 0 {
		return nil
	}
	err := resourceAviUserAccountUpdate(d, meta)
	return err
}

func resourceAviUserAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserAccountSchema()
	var err error
	client := meta.(*clients.AviClient)
	var robj interface{}
	obj := d
	username := d.Get("username")
	name := d.Get("name")
	fullName := d.Get("full_name")
	email := d.Get("email")
	local := d.Get("local")

	if data, err := SchemaToAviData(obj, s); err == nil {
		path := "api/useraccount"
		err = client.AviSession.Put(path, data, &robj)
		if err != nil {
			log.Printf("[ERROR] in updating the object %v\n", err)
		} else {
			// we dont get UUID because of nil response and username is unique in useraccount
			d.SetId(username.(string))
			d.Set("username", username)
			d.Set("name", name)
			d.Set("full_name", fullName)
			d.Set("email", email)
			d.Set("local", local)
			if err = client.AviSession.ResetPassword(d.Get("password").(string)); err != nil {
				log.Printf("[ERROR] while resetting password %v\n", err)
			}
		}
	}
	return err
}

func resourceAviUserAccountDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
