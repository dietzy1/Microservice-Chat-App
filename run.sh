#! /bin/bash

# This script is used to run multiple GRPC services and the GRPC gateway

#add permissions to the script
#chmod u+x run.sh


# Function to open a new iTerm2 window and run the specified command


function open_window_and_run_command {
  path=$1
  command=$2
  name=$3
  echo "Opening new iTerm2 window for ${path} service"
  # Check if a Go program with the specified name is already running in the specified path
 if pgrep -f "${path}${name}" >/dev/null; then
    echo "A running instance of ${name} is already open in ${path}"
  else
    # Open a new iTerm2 window
    osascript -e 'tell application "iTerm" to activate'
    osascript -e 'tell application "iTerm" to create window with default profile'
    sleep 1
    # Run the command
    osascript -e "tell application \"iTerm\" to tell current session of current window to write text \"cd ${path}\""
    osascript -e "tell application \"iTerm\" to tell current session of current window to write text \"${command}\""
  fi
}

open_window_and_run_command "go/src/github.com/dietzy1/chatapp/cmd/user/" "go run main.go" "main"
open_window_and_run_command "go/src/github.com/dietzy1/chatapp/cmd/auth/" "go run main.go" "main"
open_window_and_run_command "go/src/github.com/dietzy1/chatapp/cmd/chatroom/" "go run main.go" "main"
open_window_and_run_command "go/src/github.com/dietzy1/chatapp/cmd/message/" "go run main.go" "main"
open_window_and_run_command "go/src/github.com/dietzy1/chatapp/cmd/account/" "go run main.go" "main"
open_window_and_run_command "go/src/github.com/dietzy1/chatapp/cmd/icon/" "go run main.go" "main"
open_window_and_run_command "go/src/github.com/dietzy1/chatapp/cmd/apigateway/" "go run main.go" "main"
