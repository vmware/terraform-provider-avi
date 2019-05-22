import shlex
import logging
import subprocess
import pytest
logging.basicConfig(filename='terraform_tests.log', level=logging.INFO)
LOG = logging.getLogger(__name__)



@pytest.mark.TCID1_48_1607_1_0
def test_terraform_acc():
    cmd = "make testacc"
    LOG.info("executing command %s ", cmd)
    out = subprocess.check_output(shlex.split(cmd))
    LOG.info("terraform testacc : out %s ", out)
