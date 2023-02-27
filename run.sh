#! /bin/bash

# This script is used to run multiple GRPC services and the GRPC gateway

#add permissions to the script
#chmod u+x run.sh


# Function to open a new iTerm2 window and run the specified command
function open_window_and_run_command {
  path=$1
  command=$2
  echo "Opening new iTerm2 window for ${path} service"
  # Check if a window is already open with a running instance of main.go
  if pgrep -f "${command}" >/dev/null; then
    echo "A window with a running instance of ${path} is already open"
  else
    # Open a new iTerm2 window
    osascript -e 'tell application "iTerm" to activate'
    osascript -e 'tell application "iTerm" to create window with default profile'
    sleep 2
    # Run the command
    osascript -e "tell application \"iTerm\" to tell current session of current window to write text \"cd ${path}\""
    osascript -e "tell application \"iTerm\" to tell current session of current window to write text \"${command}\""
  fi
}

# Open windows for the user, auth, and apigateway services
open_window_and_run_command "go/src/github.com/dietzy1/chatapp/cmd/user/" "go run main.go"
open_window_and_run_command "go/src/github.com/dietzy1/chatapp/cmd/auth/" "go run main.go"
open_window_and_run_command "go/src/github.com/dietzy1/chatapp/cmd/apigateway/" "go run main.go"