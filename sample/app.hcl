workspace "app" {
  source "github" {
    owner = "terrafire"
    repo  = "terraform"
    path  = "app/"
    ref   = "xxxx"
  }

  workspace = "dev"
  vars = {
    foo_revision = "xxx"
    ami_list = ["ami-abc123","ami-def456"]
    region_map = {"us-east-1":"ami-abc123","us-east-2":"ami-def456"}
    # "bar_revision" = resolve_github_revision("terrafire-dev", "bar", "branch-do-something") # user-defined function
  }
  var_files = [
    "./app/variables.tfvars",
    "./app/secrets.tfvars.enc" # 末尾が enc だったら sops で復号
  ]
}
