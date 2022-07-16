#!/bin/bash
ssh sasami "mkdir -p ~/OnakaAPI"
rsync -av ./onaka-api sasami:~/OnakaAPI/
rsync -av ./.env sasami:~/OnakaAPI/
ssh sasami "nohup ~/OnakaAPI/onaka-api &"
