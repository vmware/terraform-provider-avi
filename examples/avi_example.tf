provider "avi" {
  avi_username = "admin"
  avi_tenant = "admin"
  avi_password = "something"
  avi_controller= "10.10.25.42"
}

data "avi_applicationprofile" "system_app_profile" {
  name= "System-HTTP"
  type= "APPLICATION_PROFILE_TYPE_HTTP"
}


output "system app profile" {
  value = "${data.avi_applicationprofile.system_app_profile.name}"
}

resource "avi_networkprofile" "test_networkprofile" {
  name= "networkprofile-42"
  profile{
    type= "PROTOCOL_TYPE_TCP_PROXY"
  }
}

resource "avi_applicationpersistenceprofile" "test_applicationpersistenceprofile" {
  name = "applicationpersistence-42"
  persistence_type = "PERSISTENCE_TYPE_CLIENT_IP_ADDRESS"
}

resource "avi_vsvip" "test_vsvip" {
  name= "vip-42"
  vip {
    vip_id= "0"
    ip_address {
      type= "V4",
      addr= "10.90.64.88",
    }
  }
}

resource "avi_virtualservice" "test_vs" {
  name= "vs-42"
  pool_ref= "${avi_pool.testpool.id}"
  application_profile_ref= "${data.avi_applicationprofile.system_app_profile.id}"
  #application_profile_ref= "${avi_applicationprofile.test_applicationprofile.id}"
  network_profile_ref = "${avi_networkprofile.test_networkprofile.id}"
  vsvip_ref = "${avi_vsvip.test_vsvip.id}"
  vip {
    vip_id= "0"
    ip_address {
      type= "V4",
      addr= "10.90.64.88",
    }
  }
  services {
    port= 80
    enable_ssl= true
  }
}

resource "avi_applicationprofile" "test_applicationprofile" {
  name= "applicationprofile-42"
  type= "APPLICATION_PROFILE_TYPE_DNS"
}

resource "avi_healthmonitor" "test_hm_1" {
  name = "healthmonitor-42"
  type = "HEALTH_MONITOR_HTTP"
}


resource "avi_pool" "testpool" {
  name= "pool-42",
  health_monitor_refs= ["${avi_healthmonitor.test_hm_1.id}"]
  application_persistence_profile_ref= "${avi_applicationpersistenceprofile.test_applicationpersistenceprofile.id}"
  servers {
    ip= {
      type= "V4",
      addr= "10.90.64.66",
    }
    port= 8080
  }
  fail_action= {
    type= "FAIL_ACTION_CLOSE_CONN"
  }
}
