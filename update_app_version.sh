#!/bin/bash

# Set the desired version number
app_version=$1
GITHUB_TOKEN=$2



# Clone the repository that contains the manifest
git clone https://github.com/samer955/argocd-config-login.git repo
cd repo

# Replace the old version with the new version in the manifest
sed -i "s/image: ghcr.io\/samer955\/go-login-cicd:.*/image: ghcr.io\/samer955\/go-login-cicd:$app_version/" dev/deployment.yaml

# Set the remote URL to use the GITHUB_TOKEN secret
git remote set-url origin https://x-access-token:$GITHUB_TOKEN@github.com/samer955/argocd-config-login.git

# Commit and push the changes to the repository
git config user.email "samer.osman95@hotmail.com"
git config user.name "samer955"

# Add changes and commit
git add dev/deployment.yaml
git commit -m "Update version in manifest to $app_version"

# Push changes to the repository
git push origin main
