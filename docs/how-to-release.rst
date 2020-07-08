##############
How To Release
##############

We use `SemVer <http://semver.org/>`_ for versioning.

When you release a new version of Terrafire, just push a version tag to GitHub.

When a version tag is pushed:

- GitHub Action `goreleaser <https://github.com/mitene/terrafire/actions?query=workflow%3Agoreleaser>`_ starts, and
  publish a GitHub release.
- Docker Hub starts a `build <https://hub.docker.com/repository/docker/mitene/terrafire/builds>`_, and publish a docker
  image to Docker Hub.
