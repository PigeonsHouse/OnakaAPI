#!/bin/bash
rsync -av ./ sasami:~/OnakaAPI/
ssh sasami "cd ~/OnakaAPI; docker-compose build; docker-compose down; docker-compose up -d"
