# Make generate will always update datadog-accessors.go
cp datadog-accessors.go datadog-accessors.go.bak

# Regenerate code and compare results
make generate && cmp -s datadog-accessors.go.bak datadog-accessors.go
changed=$?

# Clean up after ourselves
rm datadog-accessors.go.bak

# See if contents have changed and error if they have
if [[ $changed != 0 ]] ; then
    echo "Did you run 'make generate' before committing?"
    exit 1
fi
