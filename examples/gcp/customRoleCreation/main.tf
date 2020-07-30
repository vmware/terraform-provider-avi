########################################################################### 
### NOTE: You sould have the following permissions to run this script:- ###
###### 1: resourcemanager.projects.setIamPolicy  ########################## 
###### 2: resourcemanager.projects.getIamPolicy  ########################## 
###### 3: resourcemanager.projects.get  ################################### 
###### 4: iam.roles.create  ############################################### 
###### 5: iam.roles.delete  ############################################### 
###### 6: iam.roles.list  ################################################# 
###### 7: iam.roles.undelete  ############################################# 
###### 8: iam.roles.update  ############################################### 
###### 9: iam.roles.get  ################################################## 
########################################################################### 

// Configure the Google Cloud provider
provider "google" {
  credentials = file(var.credential_file)
}

// Creating network project role
resource "google_project_iam_custom_role" "network_project_role" {
  role_id     = var.network_role_id
  project     = var.network_project_id
  title       = "AVI Network Project Role"
  description = "Access to resources required for operations in Network Project"
  stage       = "ALPHA"
  permissions = ["compute.networks.get", "compute.networks.list", "compute.networks.updatePolicy", "compute.regions.get", "compute.routes.create", "compute.routes.delete", "compute.routes.list", "compute.subnetworks.get", "compute.subnetworks.list", "compute.subnetworks.use"]
}

// Creating service engine project role
resource "google_project_iam_custom_role" "se_project_role" {
  role_id     = var.se_role_id
  project     = var.se_project_id
  title       = "AVI Service Engine Project Role"
  description = "Access to resources required for operations on Service Engines and Virtual Services"
  stage       = "ALPHA"
  permissions = ["compute.addresses.create", "compute.addresses.delete", "compute.addresses.get", "compute.addresses.list", "compute.addresses.use", "compute.disks.create", "compute.forwardingRules.create", "compute.forwardingRules.delete", "compute.forwardingRules.list", "compute.globalOperations.get", "compute.images.create", "compute.images.delete", "compute.images.list", "compute.images.useReadOnly", "compute.instances.create", "compute.instances.delete", "compute.instances.get", "compute.instances.list", "compute.instances.setLabels", "compute.instances.setMetadata", "compute.instances.setTags", "compute.instances.use", "compute.regionOperations.get", "compute.regions.get", "compute.regions.list", "compute.targetPools.addInstance", "compute.targetPools.create", "compute.targetPools.delete", "compute.targetPools.get", "compute.targetPools.list", "compute.targetPools.removeInstance", "compute.targetPools.use", "compute.zoneOperations.get", "compute.zones.list"]
}

// Creating storage project role
resource "google_project_iam_custom_role" "storage_project_role" {
  role_id     = var.storage_role_id
  project     = var.storage_project_id
  title       = "AVI Storage Project Role"
  description = "Access to resources required for operations on GCS Buckets and Objects"
  stage       = "ALPHA"
  permissions = ["storage.buckets.create", "storage.buckets.delete", "storage.objects.create", "storage.objects.delete", "storage.objects.list"]
}

// Getting our service account (sa)
data "google_service_account" "sa" {
  account_id = var.sa_account_id
  project    = var.sa_account_project
}

// Adding it to our new project(s) with respective custom role
resource "google_project_iam_member" "network" {
  project = var.network_project_id
  role    = "projects/${var.network_project_id}/roles/${var.network_role_id}"
  member  = "serviceAccount:${data.google_service_account.sa.email}"
}

resource "google_project_iam_member" "se" {
  project = var.se_project_id
  role    = "projects/${var.se_project_id}/roles/${var.se_role_id}"
  member  = "serviceAccount:${data.google_service_account.sa.email}"
}

resource "google_project_iam_member" "storage" {
  project = var.storage_project_id
  role    = "projects/${var.storage_project_id}/roles/${var.storage_role_id}"
  member  = "serviceAccount:${data.google_service_account.sa.email}"
}

