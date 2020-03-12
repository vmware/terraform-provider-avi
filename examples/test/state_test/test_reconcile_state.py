import subprocess
import os
import unittest
import pytest
from avi.sdk.avi_api import ApiSession

def create_sesionm():
    api = ApiSession.get_session(
        controller_ip=os.getenv("AVI_CONTROLLER"), username=os.getenv("AVI_USERNAME", "admin"),
        password=os.getenv("AVI_PASSWORD"), tenant="admin",
        api_version=os.getenv("AVI_VERSION"), verify=False)
    return api

def terraform_init():
  subprocess.check_output("terraform init", shell=True)

def terraform_plan():
  subprocess.check_output("terraform plan", shell=True)

def terraform_apply():
  subprocess.check_output("terraform apply -auto-approve", shell=True)

def terraform_destroy():
  subprocess.check_output("terraform destroy -auto-approve", shell=True)

class TestTerraform(unittest.TestCase):

  # Test terraform initialisation
  @pytest.mark.run(order=1)
  def test_terraform_init(self):
    terraform_init()

  # test terraform plan
  @pytest.mark.run(order=2)
  def test_terraform_plan(self):
    terraform_plan()

  # terraform apply will create configuration on the controller.
  @pytest.mark.run(order=3)
  def test_terraform_apply(self):
    terraform_apply()
    # Delete pool object from controller
    api = create_sesionm()
    resp = api.delete_by_name("pool", "terraform-simple-pool", api_version=os.getenv("AVI_VERSION"))
    assert resp.status_code in [200, 204]

  # try to recreate configuration on the controller
  @pytest.mark.run(order=4)
  def test_terraform_reapply(self):
    terraform_apply()

  # destroy the created configuration.
  @pytest.mark.run(order=5)
  def test_terraform_destroy(self):
    terraform_destroy()


if __name__ == "main":
  unittest.main()
