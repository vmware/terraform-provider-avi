#!/bin/bash

sleep 360
response=0
retries=0
echo "checking controller service availability for $CONTROLLER_ADDRESS..."
# wait for the service to become available
while [ $response -ne 200 ] && [ $retries -lt 20 ]
do
    response=$(curl -kL \
        --write-out '%{http_code}' \
        --silent \
        --output /dev/null \
        https://$CONTROLLER_ADDRESS/api/cluster/runtime)
    retries=$(($retries+1))
    echo "cluster runtime endpoint for service status: $response"
    case $response in
        000)
        echo "CLUSTER ENDPOINT IS UNAVAILABLE"
        ;;
        200)
        echo "CLUSTER ENDPOINT IS AVAILABLE"
        break
        ;;
        204)
        echo "CLUSTER ENDPOINT IS AVAILABLE"
        response=200
        break
        ;;
    esac
    sleep $POLL_INTERVAL
done
if [ $response -ne 200 ]; then
    echo "ERROR: CLUSTER SERVICE IS UNAVAILABLE"
    exit 255
fi


retries=0
# wait for the service to return a progress = 100 response
while [ $response -ne 100 ] && [ $retries -lt 10 ]
do
    sleep $POLL_INTERVAL
    response=$(curl -kL \
        --silent \
        https://$CONTROLLER_ADDRESS/api/cluster/runtime | jq .cluster_state.progress)
    retries=$(($retries+1))
    echo "cluster progress status: $response"
    case $response in
        100)
        echo "CLUSTER COMPLETE"
        break
        ;;
        [0-99]*)
        echo "CLUSTER IN PROGRESS"
        ;;
        *)
        echo "UNEXPECTED RESPONSE TO PROGRESS REQUEST"
        ;;
    esac
done

