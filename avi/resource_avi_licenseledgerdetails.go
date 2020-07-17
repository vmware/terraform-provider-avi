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

func ResourceLicenseLedgerDetailsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"escrow_infos": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceLicenseInfoSchema(),
		},
		"se_infos": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceLicenseInfoSchema(),
		},
		"tier_usages": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceLicenseTierUsageSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviLicenseLedgerDetails() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviLicenseLedgerDetailsCreate,
		Read:   ResourceAviLicenseLedgerDetailsRead,
		Update: resourceAviLicenseLedgerDetailsUpdate,
		Delete: resourceAviLicenseLedgerDetailsDelete,
		Schema: ResourceLicenseLedgerDetailsSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceLicenseLedgerDetailsImporter,
		},
	}
}

func ResourceLicenseLedgerDetailsImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceLicenseLedgerDetailsSchema()
	return ResourceImporter(d, m, "licenseledgerdetails", s)
}

func ResourceAviLicenseLedgerDetailsRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLicenseLedgerDetailsSchema()
	err := ApiRead(d, meta, "licenseledgerdetails", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviLicenseLedgerDetailsCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLicenseLedgerDetailsSchema()
	err := ApiCreateOrUpdate(d, meta, "licenseledgerdetails", s)
	if err == nil {
		err = ResourceAviLicenseLedgerDetailsRead(d, meta)
	}
	return err
}

func resourceAviLicenseLedgerDetailsUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLicenseLedgerDetailsSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "licenseledgerdetails", s)
	if err == nil {
		err = ResourceAviLicenseLedgerDetailsRead(d, meta)
	}
	return err
}

func resourceAviLicenseLedgerDetailsDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "licenseledgerdetails"
	client := meta.(*clients.AviClient)
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviLicenseLedgerDetailsDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
