#!/usr/bin/env bash

function launch_central {
    SWARM_DIR="$1"
    PREVENT_IMAGE="$2"
    NAMESPACE="$3"
    LOCAL_API_ENDPOINT="$4"
    PREVENT_DISABLE_REGISTRY_AUTH="$5"

    echo "Generating central config..."
    OLD_DOCKER_HOST="$DOCKER_HOST"
    OLD_DOCKER_CERT_PATH="$DOCKER_CERT_PATH"
    OLD_DOCKER_TLS_VERIFY="$DOCKER_TLS_VERIFY"
    unset DOCKER_HOST DOCKER_CERT_PATH DOCKER_TLS_VERIFY

    docker run "$PREVENT_IMAGE" -t swarm -i "$PREVENT_IMAGE" -p 8080 > "$SWARM_DIR/central.zip"

    export DOCKER_HOST="$OLD_DOCKER_HOST"
    export DOCKER_CERT_PATH="$OLD_DOCKER_CERT_PATH"
    export DOCKER_TLS_VERIFY="$OLD_DOCKER_TLS_VERIFY"

    UNZIP_DIR="$SWARM_DIR/central-deploy/"
    rm -rf "$UNZIP_DIR"
    unzip "$SWARM_DIR/central.zip" -d "$UNZIP_DIR"
    echo

    echo "Deploying Central..."
    if [ "$PREVENT_DISABLE_REGISTRY_AUTH" = "true" ]; then
        cp "$UNZIP_DIR/deploy.sh" "$UNZIP_DIR/tmp"
        cat "$UNZIP_DIR/tmp" | sed "s/--with-registry-auth//" > "$UNZIP_DIR/deploy.sh"
        rm "$UNZIP_DIR/tmp"
    fi
    $UNZIP_DIR/deploy.sh
    echo
    wait_for_central "$LOCAL_API_ENDPOINT"
    echo "Successfully launched central"
    echo "Access the UI at: https://$LOCAL_API_ENDPOINT"
}

function launch_sensor {
    SWARM_DIR="$1"
    PREVENT_IMAGE="$2"
    CLUSTER="$3"
    NAMESPACE="$4"
    CLUSTER_API_ENDPOINT="$5"
    LOCAL_API_ENDPOINT="$6"
    PREVENT_DISABLE_REGISTRY_AUTH="$7"

    EXTRA_CONFIG=""
    if [ "$DOCKER_CERT_PATH" = "" ]; then
        EXTRA_CONFIG="\"swarm\": { \"disableSwarmTls\":true } }"
    fi
    get_cluster_zip "$LOCAL_API_ENDPOINT" "$CLUSTER" SWARM_CLUSTER "$PREVENT_IMAGE" "$CLUSTER_API_ENDPOINT" "$SWARM_DIR" "$EXTRA_CONFIG"

    echo "Deploying Sensor..."
    UNZIP_DIR="$SWARM_DIR/sensor-deploy/"
    rm -rf "$UNZIP_DIR"
    unzip "$SWARM_DIR/sensor-deploy.zip" -d "$UNZIP_DIR"

    if [ "$PREVENT_DISABLE_REGISTRY_AUTH" = "true" ]; then
        cp "$UNZIP_DIR/sensor-deploy.sh" "$UNZIP_DIR/tmp"
        cat "$UNZIP_DIR/tmp" | sed "s/--with-registry-auth//" > "$UNZIP_DIR/sensor-deploy.sh"
        rm "$UNZIP_DIR/tmp"
    fi

    $UNZIP_DIR/sensor-deploy.sh
    echo

    echo "Successfully deployed!"
}