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

func ResourceSnmpTrapProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"trap_servers": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceSnmpTrapServerSchema(),
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviSnmpTrapProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSnmpTrapProfileCreate,
		Read:   ResourceAviSnmpTrapProfileRead,
		Update: resourceAviSnmpTrapProfileUpdate,
		Delete: resourceAviSnmpTrapProfileDelete,
		Schema: ResourceSnmpTrapProfileSchema(),
	}
}

func ResourceAviSnmpTrapProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSnmpTrapProfileSchema()
	log.Printf("[INFO] ResourceAviSnmpTrapProfileRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/snmptrapprofile/" + uuid.(string)
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
	log.Printf("ResourceAviSnmpTrapProfileRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviSnmpTrapProfileRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviSnmpTrapProfileRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviSnmpTrapProfileRead Updated %v\n", d)
	return nil
}

func resourceAviSnmpTrapProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSnmpTrapProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "snmptrapprofile", s)
	log.Printf("[DEBUG] created object %v: %v", "snmptrapprofile", d)
	if err == nil {
		err = ResourceAviSnmpTrapProfileRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "snmptrapprofile", d)
	return err
}

func resourceAviSnmpTrapProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSnmpTrapProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "snmptrapprofile", s)
	if err == nil {
		err = ResourceAviSnmpTrapProfileRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "snmptrapprofile", d)
	return err
}

func resourceAviSnmpTrapProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "snmptrapprofile"
	log.Println("[INFO] ResourceAviSnmpTrapProfileRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviSnmpTrapProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
