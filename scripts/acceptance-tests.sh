#!/usr/bin/env sh

if [ "$TRAVIS_PULL_REQUEST" = "true" ]; then
    echo "Not running acceptance tests for pull requests"
    exit 0
fi

# Using multiple Datadog accounts to minimise flakiness during parallel runs, as some tests rely
# an amount of items across a whole account.
case "$TRAVIS_GO_VERSION" in
    1.10*)
        export DATADOG_API_KEY=${DATADOG_API_KEY_A:?}
        export DATADOG_APP_KEY=${DATADOG_APP_KEY_A:?}
        echo "Using Datadog 'A' test account"
        ;;
    1.11*)
        export DATADOG_API_KEY=${DATADOG_API_KEY_B:?}
        export DATADOG_APP_KEY=${DATADOG_APP_KEY_B:?}
        echo "Using Datadog 'B' test account"
        ;;
    *)
        export DATADOG_API_KEY=${DATADOG_API_KEY_C:?}
        export DATADOG_APP_KEY=${DATADOG_APP_KEY_C:?}
        echo "Using Datadog 'C' test account"
        ;;
esac

echo "Running acceptance tests"
GOCACHE=off make testacc
