###########
Development
###########

*****
Build
*****

First, install Go and npm.

You can build the terrafire executable:

.. code-block:: bash

    make -C web setup # install npm packages
    make build

``make build`` command builds javascript assets, embeds them into Go project (using go.rice), and compiles Go sources.
See Makefile for details.

**********************
Update Protocol Buffer
**********************

Terrafire uses gRPC for internal communication.
Proto files are placed in ``api`` directory.

Generated codes by protocol buffer are included in version control system.
So, if you update proto files, run the following command and commit auto generated codes.

.. code-block:: bash

    make proto

*******************
Development Servers
*******************

Start development servers:

.. code-block:: bash

    go run ./cmd/terrafire server      # start server

    go run ./cmd/terrafire controller  # start controller

    cd web && npm start                # start npm proxy server

These servers are useful for local development and debugging.

(It is recommended to configure environment variables using `direnv <https://github.com/direnv/direnv>`_.)

**********
Build With
**********

* `go.rice <https://github.com/GeertJohan/go.rice>`_
* `gorm <https://github.com/go-gorm/gorm>`_
* gRPC
* `React <https://reactjs.org/>`_

**************
How To Release
**************

We use `SemVer <http://semver.org/>`_ for versioning.

When you release a new version of Terrafire, just push a version tag to GitHub.

When a version tag is pushed:

- GitHub Action `goreleaser <https://github.com/mitene/terrafire/actions?query=workflow%3Agoreleaser>`_ starts, and
  publish a GitHub release.
- Docker Hub starts a `build <https://hub.docker.com/repository/docker/mitene/terrafire/builds>`_, and publish a docker
  image to Docker Hub.
