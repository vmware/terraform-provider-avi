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

func ResourceCloudConnectorUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"azure_serviceprincipal": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAzureServicePrincipalCredentialsSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"azure_userpass": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAzureUserPassCredentialsSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"private_key": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"public_key": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
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

func resourceAviCloudConnectorUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCloudConnectorUserCreate,
		Read:   ResourceAviCloudConnectorUserRead,
		Update: resourceAviCloudConnectorUserUpdate,
		Delete: resourceAviCloudConnectorUserDelete,
		Schema: ResourceCloudConnectorUserSchema(),
	}
}

func ResourceAviCloudConnectorUserRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudConnectorUserSchema()
	log.Printf("[INFO] ResourceAviCloudConnectorUserRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/cloudconnectoruser/" + uuid.(string)
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
	log.Printf("ResourceAviCloudConnectorUserRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviCloudConnectorUserRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviCloudConnectorUserRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviCloudConnectorUserRead Updated %v\n", d)
	return nil
}

func resourceAviCloudConnectorUserCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudConnectorUserSchema()
	err := ApiCreateOrUpdate(d, meta, "cloudconnectoruser", s)
	log.Printf("[DEBUG] created object %v: %v", "cloudconnectoruser", d)
	if err == nil {
		err = ResourceAviCloudConnectorUserRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "cloudconnectoruser", d)
	return err
}

func resourceAviCloudConnectorUserUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudConnectorUserSchema()
	err := ApiCreateOrUpdate(d, meta, "cloudconnectoruser", s)
	if err == nil {
		err = ResourceAviCloudConnectorUserRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "cloudconnectoruser", d)
	return err
}

func resourceAviCloudConnectorUserDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "cloudconnectoruser"
	log.Println("[INFO] ResourceAviCloudConnectorUserRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviCloudConnectorUserDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
