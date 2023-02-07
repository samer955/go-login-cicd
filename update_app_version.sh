#!/bin/bash

#Goal of this script is to update after every push the image:tag of the kubernetes Manifest in order to trigger argoCD and update the running app on kubernetes (Continuous Delivery).

# Env variables passed from the CI Pipeline (Github Workflow).
APP_VERSION=$1
GITHUB_TOKEN=$2
GITHUB_ACTOR=$3


# Clone the repository that contains the manifest
echo "Cloning the repository containing the manifest"
git clone https://github.com/samer955/argocd-config-login.git repo
cd repo

# Replace the old version with the new version in the manifest
echo "Updating the Manifest to version $APP_VERSION"
sed -i "s/image: ghcr.io\/samer955\/go-login-cicd:.*/image: ghcr.io\/samer955\/go-login-cicd:$APP_VERSION/" dev/deployment.yaml

# Set the remote URL to use the GITHUB_TOKEN secret
git remote set-url origin https://x-access-token:$GITHUB_TOKEN@github.com/samer955/argocd-config-login.git

# Commit and push the changes to the repository
git config user.email "$GITHUB_ACTOR@github.com"
git config user.name "$GITHUB_ACTOR"

# Add changes and commit
git add dev/deployment.yaml
git commit -m "Update version in manifest to $APP_VERSION"

# Push changes to the repository
git push origin main
echo "Pushed changes to the repo"

# Clean up
cd ..
rm -rf repo

echo "Done"
