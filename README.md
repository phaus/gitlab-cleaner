# Gitlab Registry Cleaner

This Projects aims to create a Small Docker Container that can be run as part of a gitlab-ci Pipeline to trigger a Docker Registry cleanup.

## Development Setup

Install glide follwing the instructions here https://github.com/Masterminds/glide#install

    glide install

For the development setup, you have to connect your local copy to the Golang Dependency System.
In `$GOPATH/src/github.com/phaus` create a Softlink to your local copy (e.g. `ln -s /Users/philipp/GIT/go/registry-cleaner`).

### Optional

Install `direnv`. Copy and modify `.envrc.example` to `.envrc` matching to your gitlab project configuration.
Run `direnv allow .` in your current folder. Your ENV Variables should now be injected, everytime you enter that folder.