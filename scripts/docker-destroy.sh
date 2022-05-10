#!/bin/bash
echo 'Destroy containers? Data in database will be lost. (y/n)'
read REPLY
if [[ ! $REPLY =~ ^[Yy]$ ]]
then
    echo 'No changes made.'
    exit 0
fi
docker-compose down -v
