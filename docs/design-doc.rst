##########
Design Doc
##########

*************
Core Concepts
*************

Terrafire is inspired by `GitOps <https://www.weave.works/technologies/gitops/>`_ from Weaveworks and
`ArgoCD <https://www.weave.works/technologies/gitops/>`_.

Users declare the desired state of their infrastructure in **Terraform modules**.
And then they specify the deployment settings in **Terrafire manifests**.
A Terrafire manifest includes *which Terrafire module* they want to apply of *which version*, with *what variables*.
Both Terraform modules and Terrafire manifests must be saved in Git repositories (separate repositories are recommended).

A Terrafire manifest looks like this:

.. code-block:: hcl

    workspace "app" {
      source "github" {
        owner = "org"
        repo  = "terraform"
        ref   = "0da931"
      }

      workspace = "dev"
      vars = {
        data = {
          foo = "FOO"
          bar = "BAR"
        }
      }
    }

``source`` block describes the Terraform module repository path and its version (``ref`` can be git commit hash, git
branch or git tag).
``vars`` block declares input variables passed to the Terraform module.
``workspace`` keyword appears twice. The inner ``workspace`` means `Terraform workspace <https://www.terraform.io/docs/state/workspaces.html>`_.
The top-level ``workspace`` means Terrafire workspace (similar to workspaces in Terraform Cloud).

What Terrafire do is very simple:

1. Terrafire first fetches a Terrafire manifest from a Git repository.
2. Next, it fetches the Terraform module from the Git repository specified in the Terrafire manifest.
3. Finally, it executes a Terraform command (plan/apply) with the parameters specified in the Terrafire manifest.

Projects and Workspaces
=======================

Terrafire deployments settings are organized by projects and workspaces.

A Terrafire **workspace** is a deployment settings of a single Terraform module coupled with input variables and a
Terraform workspace name.
It is similar to Terraform Cloud workspaces.
Each Terrafire workspace corresponds to a single Terraform state.
A Terrafire workspace are defined by a ``workspace`` block in Terrafire manifests.

A Terrafire **project** is a set of workspaces.
A Terrafire project is corresponds to a directory in a Terrafire manifests repository.
Terrafire reads all HCL files in the directory and groups all workspaces appeared in the HCL files as a project.

When a new Terrafire workspace is added or an existing Terrafire workspace is updated in manifests, Terrafire update the
infrastructure resources by Terraform apply command.
When a Terrafire workspace is deleted from manifests, Terrafire destroys the infrastructure resources by Terraform
destroy command, and tries to delete Terraform state file.

************
Architecture
************

Components
==========

.. code-block::

                                                                                    +- git repo -+  +- git repo -+
                                                                                    | Terraform  |  | Terrafire  |
                                                                                    |  modules   |  | manifests  |
                                                                                    +------------+  +------------+
                                                                                                |    |
                                                                                          fetch |    |
                                     |                                                          v    v
    +--------+            +---------------------------+              +------------+           +--------+
    | client |            |         server            |              | controller |           | runner |
    |        |  request   | +-----+     +-----------+ |   subscribe  |            |  launch   +--------+
    |        | ---------> | | web | <-> | scheduler | | <----------- |            | -------->    ...
    |        |            | |     |     |           | |              +------------+              ...
    |        |            | |     |     |           | |   job status                          +--------+
    |        |            | +-----+     +-----------+ | <------------------------------------ | runner |
    +--------+            +---------------------------+                                       +--------+
                                     |
                   public network <- | -> private network
                                     |

Terrafire consists of 4 components: ``client``, ``server``, ``controller`` and ``runner``.

**Client** is a javascript application executed in users' web browser.

**Server** has two internal services, ``web`` and ``scheduler``.
**Web** is a HTTP server that handles requests from clients.
**Scheduler** is a gRPC server that communicates with the controller.
The scheduler acts as a job queue and manages jobs execution status.
The web server may be exposed to public internet, but the scheduler only needs to be accessible from a controller and
runners.

**Controller** subscribes to the scheduler and receives job requests.
When the controller receives a job request, it launches a runner processe.
Runner processes can run in various platforms, for example on local sub process, AWS ECS Task, or maybe Kubernetes.

**Runner** fetches Terrafire manifests and Terraform modules from git repositories and executes Terraform plan/apply
command.
Runners notify a job execution status to scheduler periodically.

Security
========

In many cases, Terraform commands are executed under very strong permission (ex. AWS Admin role).
So, Terrafire runners will have the strong permission, and a Terrafire controller also.

To protect the controller and the runners from attacks, Terrafire isolates them to a private network.
Though Terrafire server may be public, the server never access to the controller nor the runners.
Only the controller and the runners access to the server.

***************
Execution Steps
***************

#. The client sends a "Plan" request to the server.
#. The server pushes the request to an internal job queue.
#. The controller fetches the request from the job queue via the scheduler.
#. The controller launches a runner process.
#. The runner fetches the Terrafire manifests from a Git repository.
#. The runner fetches the Terraform modules from a Git repository specified in the job request.
#. The runner executes Terraform plan command.
#. The runner saves plan results in a blob store.
#. The runner notifies the server that a Plan job finishes.
#. The server waits for the user approves the plan result.
#. When the user approves the plan result, the client send an "Apply" request to the server.
#. The server pushes the request to a job queue, the controller fetches the request from the scheduler and launches a
   runner.
#. The runner retrieves the saved plan result from the blob store, and execute Terraform apply command using the plan
   result.
