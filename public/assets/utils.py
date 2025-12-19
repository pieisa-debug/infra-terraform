import logging
import os
import re
from typing import List, Tuple

# Set up logging
logging.basicConfig(level=logging.INFO)

def get_tfvars(args: List[str]) -> Tuple[str, bool]:
    """Find the first Terraform variable file in the current directory and its subdirectories."""
    for root, dirs, files in os.walk('.'):
        if 'terraform.tfvars' in files or 'terraform.tfvar' in files:
            return os.path.join(root, 'terraform.tfvars'), True
    return '', False

def get_tf_version() -> str:
    """Find the current Terraform version."""
    version = os.environ.get('TF_VERSION')
    if version:
        return version
    return 'unknown'

def get_terraform_executable_path() -> str:
    """Find the path to the Terraform executable."""
    for key in os.environ:
        if re.match(r'TF_(EXECUTABLE|PATH|BIN)', key):
            return os.environ[key]
    return 'terraform'

def get_terraform_executable() -> str:
    """Find the name of the Terraform executable."""
    return os.path.basename(get_terraform_executable_path())