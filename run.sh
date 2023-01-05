#! /bin/bash

# This script is used to run multiple GRPC services and the GRPC gateway

#add permissions to the script
#chmod u+x run.sh

# Run the GRPC services
echo "Hello from run.bash"

#print current directory
echo "$PWD"

open -a iTerm cmd/auth/auth/
#open -a iTerm .
