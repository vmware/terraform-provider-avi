/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/avinetworks/sdk/go/session"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"strconv"
	"testing"
	"time"
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
		"avi_controller": "", "avi_password": "", "avi_tenant": "", "avi_api_timeout": 0}

	_, errs := Provider().(*schema.Provider).Validate(
		&terraform.ResourceConfig{Config: configs})
	if errs != nil {
		t.Fatalf("err: %s", errs)
	}

	// Validating pool resource in avi provider and datasource for pool
	var poolconfigs_data = map[string]interface{}{"name": ""}

	_, errs = Provider().(*schema.Provider).ValidateDataSource("avi_pool",
		&terraform.ResourceConfig{Config: poolconfigs_data})
	if errs != nil {
		t.Fatalf("err: %s", errs)
	}

	var poolconfigs_res = map[string]interface{}{"name": "", "uuid": "",
		"health_monitor_refs": make([]string, 0), "servers": make([]string, 0),
		"fail_action": make([]string, 0)}

	_, errs = Provider().(*schema.Provider).ValidateResource("avi_pool",
		&terraform.ResourceConfig{Config: poolconfigs_res})
	if errs != nil {
		t.Fatalf("err: %s", errs)
	}

	// Validating pool resource in avi provider and datasource for pool

	var hmconfigs_data = map[string]interface{}{"name": ""}

	_, errs = Provider().(*schema.Provider).ValidateDataSource(
		"avi_healthmonitor", &terraform.ResourceConfig{Config: hmconfigs_data})
	if errs != nil {
		t.Fatalf("err: %s", errs)
	}

	var hmconfigs_res = map[string]interface{}{"name": "", "uuid": "",
		"type": ""}

	_, errs = Provider().(*schema.Provider).ValidateResource(
		"avi_healthmonitor", &terraform.ResourceConfig{Config: hmconfigs_res})
	if errs != nil {
		t.Fatalf("err: %s", errs)
	}
}

func testAccPreCheck(t *testing.T) {
	var timeout time.Duration
	if tm, err := strconv.Atoi(os.Getenv("AVI_API_TIMEOUT")); err == nil {
		timeout = time.Duration(time.Duration(int(tm)) * time.Second)
	} else {
		t.Fatalf("AVI_API_TIMEOUT must be numeric value to set timeout for acceptance test")
	}

	config := Credentials{
		Username:   os.Getenv("AVI_USERNAME"),
		Password:   os.Getenv("AVI_PASSWORD"),
		Controller: os.Getenv("AVI_CONTROLLER"),
		Tenant:     os.Getenv("AVI_TENANT"),
		Version:    os.Getenv("AVI_VERSION"),
		AuthToken:  os.Getenv("AVI_AUTHTOKEN"),
		Timeout:    timeout,
	}

	if config.Controller == "" {
		t.Fatalf("AVI_CONTROLLER must be set for acceptance test")
	}

	if config.Password == "" && config.AuthToken == "" {
		t.Fatalf("AVI_PASSWORD or AVI_AUTHTOKEN must be set for acceptance test")
	}

	if config.Username == "" {
		config.Username = "admin"
	}
	if config.Tenant == "" {
		config.Tenant = "admin"
	}
	if config.Version == "" {
		config.Version = "18.2.1"
	}

	_, err := clients.NewAviClient(
		config.Controller, config.Username,
		session.SetPassword(config.Password),
		session.SetTenant(config.Tenant),
		session.SetVersion(config.Version),
		session.SetAuthToken(config.AuthToken),
		session.SetInsecure, session.SetTimeout(config.Timeout))

	if err != nil {
		t.Fatal(fmt.Sprintf("%+v", err))
	}
}
