terraform_deploy "system" {
  source {
    owner    = "terrafire-dev"
    repo     = "terraform"
    path     = "aws/system/"
    revision = "xxxx"
  }

  workspace   = "dev"

  vars = {
    "cluster_common_revision"          = "xxx"
  }
}

