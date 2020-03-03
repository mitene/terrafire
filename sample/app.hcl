terraform_deploy "app_a" {
  source {
    owner    = "terrafire"
    repo     = "terraform"
    path     = "aws/app/"
    revision = "xxxx"
  }

  workspace     = "dev"
  allow_destroy = true

  vars = {
    "foo_revision"          = "xxx"
    # "bar_revision"      = resolve_github_revision("terrafire-dev", "bar", "branch-do-something") # user-defined function
  }

  var_files = [
    "./app_a/variables.tfvars",
    "./app_a/secrets.tfvars.enc" # 末尾が enc だったら sops で復号
  ]
}