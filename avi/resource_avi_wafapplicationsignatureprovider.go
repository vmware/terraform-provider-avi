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

func ResourceWafApplicationSignatureProviderSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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

func resourceAviWafApplicationSignatureProvider() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWafApplicationSignatureProviderCreate,
		Read:   ResourceAviWafApplicationSignatureProviderRead,
		Update: resourceAviWafApplicationSignatureProviderUpdate,
		Delete: resourceAviWafApplicationSignatureProviderDelete,
		Schema: ResourceWafApplicationSignatureProviderSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceWafApplicationSignatureProviderImporter,
		},
	}
}

func ResourceWafApplicationSignatureProviderImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceWafApplicationSignatureProviderSchema()
	return ResourceImporter(d, m, "wafapplicationsignatureprovider", s)
}

func ResourceAviWafApplicationSignatureProviderRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafApplicationSignatureProviderSchema()
	err := ApiRead(d, meta, "wafapplicationsignatureprovider", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWafApplicationSignatureProviderCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafApplicationSignatureProviderSchema()
	err := ApiCreateOrUpdate(d, meta, "wafapplicationsignatureprovider", s)
	if err == nil {
		err = ResourceAviWafApplicationSignatureProviderRead(d, meta)
	}
	return err
}

func resourceAviWafApplicationSignatureProviderUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafApplicationSignatureProviderSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "wafapplicationsignatureprovider", s)
	if err == nil {
		err = ResourceAviWafApplicationSignatureProviderRead(d, meta)
	}
	return err
}

func resourceAviWafApplicationSignatureProviderDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "wafapplicationsignatureprovider"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviWafApplicationSignatureProviderDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
