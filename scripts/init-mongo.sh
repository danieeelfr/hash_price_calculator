#!/usr/bin/env bash


echo 'Creating application user and db';


mongo discountdb  --username user  --password pwd  --authenticationDatabase admin  --host localhost  --port 27017  --eval "db.createUser({user: 'user', pwd: 'pwd', roles:[{role:'dbOwner', db: 'discountdb'}]});"


echo 'User: user create to database discountdb';


