#!/bin/bash

# Wait for MongoDB to start
until mongo --eval "print('MongoDB is ready')" > /dev/null 2>&1
do
    echo "Waiting for MongoDB to start..."
    sleep 2
done

# Run initialization commands
mongo <<EOF
db = db.getSiblingDB("${PRIVATE_NETWORK_DATABASE}");

db.createUser({
    user: "${MONGO_INITDB_ROOT_USERNAME}",
    pwd: "${MONGO_INITDB_ROOT_PASSWORD},
    roles: [{
        role: 'root',
        db: "${PRIVATE_NETWORK_DATABASE}",
    }]
});
db.createCollection('users');
EOF
