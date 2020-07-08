#########
Terrafire
#########

Terrafire is a GitOps continuous delivery tool for terraform.

***************
Getting Started
***************

Create docker-compose.yml as follows:

.. code-block:: yaml

    version: "3"
    services:
      server:
        image: "mitene/terrafire"
        command: ["server"]
        environment:
          TERRAFIRE_PROJECT_default: "https://github.com/mitene/terrafire"
          TERRAFIRE_PROJECT_default_PATH: "examples/manifest"
          TERRAFIRE_PROJECT_default_BRANCH: "master"
        ports:
          - "8080:8080"
          - "8081:8081"

      controller:
        image: "mitene/terrafire"
        command: ["controller"]
        environment:
          TERRAFIRE_SCHEDULER_ADDRESS: "server:8081"
          TERRAFIRE_PROJECT_default: "https://github.com/mitene/terrafire"
          TERRAFIRE_PROJECT_default_PATH: "examples/manifest"
          TERRAFIRE_PROJECT_default_BRANCH: "master"

Start services:

.. code-block::

    docker-compose up

Now, open http://localhost:8080

*************
Configuration
*************

You can customize all Terrafire configurations by environment variables.

server config
=============

:TERRAFIRE_SERVER_PORT: Terrafire server port.

    Web browser accesses to this port.

    *Default*: ``8080``

:TERRAFIRE_SCHEDULER_PORT: Terrafire scheduler port.

    Only terrafire controller and runner access to this port.

    *Default*: ``8081``

:TERRAFIRE_DB_DRIVER: Database driver type.

    *Available values*: ``sqlite3``

    *Default*: ``sqlite3``

:TERRAFIRE_DB_ADDRESS: Database address.

    *Default*: ``data/sqlite3.db``

:TERRAFIRE_PROJECT_<project name>: Terrafire manifest repository URL. (**required**)

    Supported repository is GitHub only currently.

:TERRAFIRE_PROJECT_<project name>_BRANCH: Branch name of manifest repository.

    *Default*: ``master``

:TERRAFIRE_PROJECT_<project name>_PATH: Relative path from manifest repository root. (optional)

:TERRAFIRE_GIT_CREDENTIAL_<credential id>_PROTOCOL: Git credentials protocol.

    Required if you have private repository.

    See https://git-scm.com/book/en/v2/Git-Tools-Credential-Storage

:TERRAFIRE_GIT_CREDENTIAL_<credential id>_HOST: Git credentials host.

    See https://git-scm.com/book/en/v2/Git-Tools-Credential-Storage

:TERRAFIRE_GIT_CREDENTIAL_<credential id>_USER: Git credentials user.

    See https://git-scm.com/book/en/v2/Git-Tools-Credential-Storage

:TERRAFIRE_GIT_CREDENTIAL_<credential id>_PASSWORD: Git credentials password.

    See https://git-scm.com/book/en/v2/Git-Tools-Credential-Storage

controller config
=================

:TERRAFIRE_SCHEDULER_ADDRESS: Terrafire scheduler address.

    *Default*: ``localhost:8081``

:TERRAFIRE_CONCURRENCY: Max concurrency that terrafire runners run in parallel.

    *Default*: ``1``

:TERRAFIRE_EXECUTOR_TYPE: Launch type that terrafire controller starts terrafire runners.

    *Available Values*: ``local``, ``ecs``

    *Default*: ``local``

:TERRAFIRE_EXECUTOR_ECS_CLUSTER: ECS custer name.

    Effects only when executor type is ``ecs``.

:TERRAFIRE_EXECUTOR_ECS_TASK_DEFINITION: ECS task definition ARN.

    **Required** if executor type is ``ecs``.

:TERRAFIRE_EXECUTOR_ECS_CONTAINER_NAME: ECS container name.

    *Default*: ``terrafire``

:TERRAFIRE_EXECUTOR_ECS_CAPACITY_PROVIDER: ECS capacit¥ provider

    *Default*: ``FARGATE``

:TERRAFIRE_EXECUTOR_ECS_SUBNETS: Comma separated values of subnet ids that ECS tasks use.

    **Required** if executor type is ``ecs``

:TERRAFIRE_EXECUTOR_ECS_SECURITY_GROUPS: Comma separated values of security groups ids that ECS tasks use.

:TERRAFIRE_EXECUTOR_ECS_ASSIGN_PUBLIC_IP: ECS assign public IP option.

    *Default*: ``true``

runner config
=============

:TERRAFIRE_SCHEDULER_ADDRESS: Terrafire scheduler address.

    *Default*: ``localhost:8081``

:TERRAFIRE_PROJECT_<project name>: Terrafire manifest repository URL. (**required**)

    Supported repository is GitHub only currently.

:TERRAFIRE_PROJECT_<project name>_BRANCH: Branch name of manifest repository.

    *Default*: ``master``

:TERRAFIRE_PROJECT_<project name>_PATH: Relative path from manifest repository root. (optional)

:TERRAFIRE_PROJECT_<project name>_ENV_<var name>: Environment values applied when terraform plan/apply are executed.

:TERRAFIRE_BLOB_TYPE: Blob type.

    Blob is an object storage where terraform plan results are stored.

    If the blob type is ``s3``, plan results are stored in local file system.

    *Available values*: ``local``, ``s3``

    *Default*: ``local``

:TERRAFIRE_BLOB_LOCAL_ROOT: Root directory of local blob.

    Effects only when the blob type is ``local``.

    *Default*: ``data/blob``

:TERRAFIRE_BLOB_S3_BUCKET: S3 bucket name for s3 blob.

    **Required** if the blob type is ``s3``.

:TERRAFIRE_BLOB_S3_PREFIX: S3 prefix under which plan results are stored.

    *Default*: ``""``

:TERRAFIRE_GIT_CREDENTIAL_<credential id>_PROTOCOL: Git credentials protocol.

    Required if you have private repository.

    See https://git-scm.com/book/en/v2/Git-Tools-Credential-Storage

:TERRAFIRE_GIT_CREDENTIAL_<credential id>_HOST: Git credentials host.

    See https://git-scm.com/book/en/v2/Git-Tools-Credential-Storage

:TERRAFIRE_GIT_CREDENTIAL_<credential id>_USER: Git credentials user.

    See https://git-scm.com/book/en/v2/Git-Tools-Credential-Storage

:TERRAFIRE_GIT_CREDENTIAL_<credential id>_PASSWORD: Git credentials password.

    See https://git-scm.com/book/en/v2/Git-Tools-Credential-Storage

***********
Development
***********

Build
=====

First, install Go and npm.

You can build the terrafire executable:

.. code-block:: bash

    make -C web setup # install npm packages
    make build

``make build`` command builds javascript assets, embeds them into Go project (using go.rice), and compiles Go sources.
See Makefile for details.

Update Protocol Buffer
======================

Terrafire uses gRPC for internal communication.
Proto files are placed in ``api`` directory.

Generated codes by protocol buffer are included in version control system.
So, if you update proto files, run the following command and commit auto generated codes.

.. code-block:: bash

    make proto

Build With
==========

* `go.rice <https://github.com/GeertJohan/go.rice>`_
* `gorm <https://github.com/go-gorm/gorm>`_
* gRPC
* `React <https://reactjs.org/>`_
