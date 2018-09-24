# Gitlab Registry Cleaner

This Projects aims to create a Small Docker Container that can be run as part of a gitlab-ci Pipeline to trigger a Docker Registry cleanup.

## Development Setup

Install glide follwing the instructions here https://github.com/Masterminds/glide#install

    glide install

For the development setup, you have to connect your local copy to the Golang Dependency System.
In `$GOPATH/src/github.com/phaus` create a Softlink to your local copy (e.g. `ln -s /Users/philipp/GIT/go/gitlab-claner`).

### Optional

Install `direnv`. Copy and modify `.envrc.example` to `.envrc` matching to your gitlab project configuration.
Run `direnv allow .` in your current folder. Your ENV Variables should now be injected, everytime you enter that folder.

Install `jq`. Run the Script `scripts/get.sh` to test your setup.

## Use the Docker Container

Wer möchte kann unter gitlab-ci gerne mal etwas testen:

1. Create a new private access token: [https://gitlab.com/profile/personal_access_tokens](https://gitlab.com/profile/personal_access_tokens) (Access to `api`)
2. Add this Token as `PRIVATE_ACCESS_TOKEN` as a CI/CD Variable.

3. Add this snippet to your `.gitlab-ci.yml` file:

````yaml
cleaner:
  stage: test
  image: phaus/gitlab-cleaner
  script:
    - cleaner test
    - cleaner list
````

At the moment, it just parses the Registry API. No changes are made.
