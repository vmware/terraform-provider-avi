package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVISchedulerBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVISchedulerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVISchedulerConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVISchedulerExists("avi_scheduler.testScheduler"),
					resource.TestCheckResourceAttr(
						"avi_scheduler.testScheduler", "name", "test-Default-Scheduler-abc"),
					resource.TestCheckResourceAttr(
						"avi_scheduler.testScheduler", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"avi_scheduler.testScheduler", "frequency", "1"),
				),
			},
			{
				Config: testAccAVISchedulerupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVISchedulerExists("avi_scheduler.testScheduler"),
					resource.TestCheckResourceAttr(
						"avi_scheduler.testScheduler", "name", "test-Default-Scheduler-updated"),
					resource.TestCheckResourceAttr(
						"avi_scheduler.testScheduler", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"avi_scheduler.testScheduler", "frequency", "1"),
				),
			},
			{
				ResourceName:      "avi_scheduler.testScheduler",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVISchedulerConfig,
			},
		},
	})

}

func testAccCheckAVISchedulerExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI Scheduler ID is set")
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVISchedulerDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_scheduler" {
			continue
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI Scheduler still exists")
		}
	}
	return nil
}

const testAccAVISchedulerConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_backupconfiguration" "default_backupconfig" {
    name= "Backup-Configuration"
}
resource "avi_scheduler" "testScheduler" {
	"start_date_time" = "2018-04-06T07:05:32.242307"
	"name" = "test-Default-Scheduler-abc"
	"enabled" = true
	"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
	"frequency_unit" = "SCHEDULER_FREQUENCY_UNIT_DAY"
	"backup_config_ref" = "${data.avi_backupconfiguration.default_backupconfig.id}"
	"frequency" = "1"
	"scheduler_action" = "SCHEDULER_ACTION_BACKUP"
	"run_mode" = "RUN_MODE_PERIODIC"
}
`

const testAccAVISchedulerupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_backupconfiguration" "default_backupconfig" {
    name= "Backup-Configuration"
}
resource "avi_scheduler" "testScheduler" {
	"start_date_time" = "2018-04-06T07:05:32.242307"
	"name" = "test-Default-Scheduler-updated"
	"enabled" = true
	"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
	"frequency_unit" = "SCHEDULER_FREQUENCY_UNIT_DAY"
	"backup_config_ref" = "${data.avi_backupconfiguration.default_backupconfig.id}"
	"frequency" = "1"
	"scheduler_action" = "SCHEDULER_ACTION_BACKUP"
	"run_mode" = "RUN_MODE_PERIODIC"
}
`
