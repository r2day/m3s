#!/bin/bash

# Get the latest version tag
latest_tag=$(git describe --abbrev=0 --tags)

# Extract the major, minor, and patch version numbers
major=$(echo $latest_tag | cut -d. -f1)
minor=$(echo $latest_tag | cut -d. -f2)
patch=$(echo $latest_tag | cut -d. -f3)

# Increment the version number
patch=$((patch + 1))

# Create the new version tag
new_tag="v$major.$minor.$patch"

# Tag the commit with the new version number
git tag $new_tag
git push origin $new_tag

echo "Tagged commit with version $new_tag"
