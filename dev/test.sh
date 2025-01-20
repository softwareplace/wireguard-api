#!/bin/bash

mkdir .out

# Number of iterations
COUNT=1000000

# The curl command
URL='http://localhost:1080/api/private-network/v1/login'
API_KEY='eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcGlLZXkiOiJFQnFtMGgxS2o0M2RxSC8rSG05L2hEZjFBelpxdFhOeUpZRVhwMDZ0YzRjanZ0RFU2cU40eXc9PSIsImNsaWVudCI6IlNvZnR3YXJlIFBsYWNlIENPIiwiZXhwIjoyMDUyNzM3ODM1fQ.PaC_hYLbGMjv9ANJO1Ch09ul0nrMUkGnXM28Z1iLLr0'
USERNAME='my-username'
PASSWORD='ynT9558iiMga&ayTVGs3Gc6ug1'


response=$(curl --silent --location --write-out "%{http_code}" --output ./.out/response.json "$URL" \
    --header "X-Api-Key: $API_KEY" \
    --header "Content-Type: application/json" \
    --data "{
      \"username\": \"$USERNAME\",
      \"password\": \"$PASSWORD\"
    }")


makePeersRequest() {
  local AUTHORIZATION
  AUTHORIZATION=$1
  local PEERS_URL
  PEERS_URL='http://localhost:1080/api/private-network/v1/peers'

  for i in $(seq 1 $COUNT); do
    echo "Request #$i"
    curl --silent --output ./.out/response.json --location "$PEERS_URL?key=$i" \
      --header "X-Api-Key: $API_KEY" \
      --header "Authorization: $AUTHORIZATION" \
      --header "Content-Type: application/json"

    cat ./.out/response.json || jq || true
  done
}

if [ "$response" -eq 200 ]; then
  token=$(jq -r '.token' ./.out/response.json)
  makePeersRequest "$token"
else
  echo "Request failed with status code: $response"
fi
#
#for i in $(seq 1 $COUNT); do
#  echo "Request #$i"
#
#done

echo "Done sending $COUNT requests."
