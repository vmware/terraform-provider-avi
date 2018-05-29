/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"fmt"
	"os"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/avinetworks/sdk/go/session"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"avi": testAccProvider,
	}
}

func TestProvider(t *testing.T) {

	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Validating Schema for Provider
	var configs = map[string]interface{}{"avi_username": "",
		"avi_controller": "", "avi_password": "", "avi_tenant": ""}

	_, errs := Provider().(*schema.Provider).Validate(
		&terraform.ResourceConfig{Config: configs})
	if errs != nil {
		t.Fatalf("err: %s", errs)
	}

	// Validating pool resource in avi provider and datasource for pool
	var poolconfigs = map[string]interface{}{"name": "", "uuid": "",
		"health_monitor_refs": make([]string, 0), "servers": make([]string, 0),
		"fail_action": make([]string, 0)}

	_, errs = Provider().(*schema.Provider).ValidateDataSource("avi_pool",
		&terraform.ResourceConfig{Config: poolconfigs})
	if errs != nil {
		t.Fatalf("err: %s", errs)
	}

	_, errs = Provider().(*schema.Provider).ValidateResource("avi_pool",
		&terraform.ResourceConfig{Config: poolconfigs})
	if errs != nil {
		t.Fatalf("err: %s", errs)
	}

	// Validating pool resource in avi provider and datasource for pool

	var hmconfigs = map[string]interface{}{"name": "", "uuid": "",
		"type": ""}

	_, errs = Provider().(*schema.Provider).ValidateDataSource(
		"avi_healthmonitor", &terraform.ResourceConfig{Config: hmconfigs})
	if errs != nil {
		t.Fatalf("err: %s", errs)
	}

	_, errs = Provider().(*schema.Provider).ValidateResource(
		"avi_healthmonitor", &terraform.ResourceConfig{Config: hmconfigs})
	if errs != nil {
		t.Fatalf("err: %s", errs)
	}
}

func testAccPreCheck(t *testing.T) {
	config := Credentials{
		Username:   os.Getenv("AVI_USERNAME"),
		Password:   os.Getenv("AVI_PASSWORD"),
		Controller: os.Getenv("AVI_CONTROLLER"),
		Tenant:     os.Getenv("AVI_TENANT"),
		Version:    os.Getenv("AVI_VERSION"),
	}
	required := []string{"AVI_PASSWORD", "AVI_CONTROLLER"}
	for _, property := range required {
		if os.Getenv(property) == "" {
			t.Fatalf("%s must be set for acceptance test", property)
		}
	}
	if config.Username == "" {
		config.Username = "admin"
	}
	if config.Tenant == "" {
		config.Tenant = "admin"
	}
	if config.Version == "" {
		config.Version = "18.1.2"
	}
	_, err := clients.NewAviClient(
		config.Controller, config.Username,
		session.SetPassword(config.Password),
		session.SetTenant(config.Tenant),
		session.SetVersion(config.Version),
		session.SetInsecure)

	if err != nil {
		t.Fatal(fmt.Sprintf("%+v", err))
	}
}
