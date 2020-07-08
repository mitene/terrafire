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
A Terrafire manifest includes *which* Terrafire module they want to apply of *which version*, with *what variables*.
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

``source`` block describes the Terraform module repository and the Terraform module version (`ref` can be git commit
hash, git branch or git tag).
``vars`` block declares variable inputs passed to Terraform module specified in ``source`` block.
``workspace`` keyword appeared twice. The inner ``workspace`` means `Terraform workspace <https://www.terraform.io/docs/state/workspaces.html>`_.
The top-level ``workspace`` means Terrafire workspace (similar to workspaces in Terraform Cloud).

What Terrafire do is very simple:

1. Terrafire first fetches a Terrafire manifest from a Git repository.
2. Next, it fetches the Terraform module from the Git repository specified in the Terrafire manifest.
3. Finally, it executes a Terraform command (plan/apply) with the parameters specified in the Terrafire manifest.

Projects and Workspaces
=======================

Terrafire deployments settings are organized by projects and workspaces.

A Terrafire **workspace** is a deployment settings of a single Terraform module coupled with input parameter values and
a Terraform workspace name.
It is similar to Terraform Cloud workspaces.
Each Terrafire workspace corresponds to a single Terraform state.
Terrafire workspaces are defined by ``workspace`` block in Terrafire manifests.

A Terrafire **project** is a set of workspaces.
A Terrafire project is corresponds to a directory in a Terrafire manifests repository.
Terrafire reads all HCL files in the directory and groups all workspaces appeared in the HCL files as a project.

When a new Terrafire workspace is added or existing Terrafire workspace is updated in manifests, Terrafire update the
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
The web server may be exposed to public internet, but the scheduler only needs to be accessible from a scheduler and
runners.

**Controller** subscribes to the scheduler and receives users request.
The controller launches runner processes.
Runner process can run in various platforms, for example on local sub process, AWS ECS Task, or maybe Kubernetes.

**Runner** fetches terrafire manifests and terraform modules from git repositories and executes terraform plan/apply.
Runners send a job execution status to scheduler.

Security
========

In many cases, Terraform command is executed under very strong permission (ex. AWS Admin role).
So, Terrafire runners will have the strong permission, and Terrafire controller also.

Terrafire isolates the controller and the runners to a private network.
Though Terrafire server may be public, the server never access to the controller nor the runners.
Only the controller and the runners access to the server.

***************
Execution Steps
***************

#. The client sends a "Plan" request to the server.
#. The server pushes the request to an internal job queue.
#. The controller fetches the request from the job queue via the scheduler.
#. The controller launches a runner process.
#. The runner fetches a Terrafire manifests from a Git repository.
#. The runner fetches a Terraform modules specified in the request from a Git repository.
#. The runner executes Terraform plan command.
#. The runner saves plan results in a blob store.
#. The runner notifies to server that it finishes a Plan phase.
#. The server waits for an user approve that plan.
#. When an user approve the plan result, the client send an "Apply" request to the server.
#. The server pushes the request to a job queue, the controller fetches the request from the scheduler and launches a
   runner.
#. The runner retrieves the saved plan result and execute Terraform apply command using the plan result.
