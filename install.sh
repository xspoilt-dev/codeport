#!/bin/bash

GREEN="\033[0;32m"
BLUE="\033[0;34m"
YELLOW="\033[1;33m"
RED="\033[0;31m"
NC="\033[0m"

echo -e "${YELLOW}Building Executable....${NC}"
go build codeport.go

if [ $? -eq 0 ]; then
  echo -e "${BLUE}Checking environment...${NC}"

  if [ -d "$PREFIX" ]; then
    echo -e "${BLUE}Detected Termux environment.${NC}"
    mv codeport "$PREFIX/bin/codeport"
  else
    echo -e "${BLUE}Detected Linux environment.${NC}"
    sudo mv codeport /usr/local/bin/codeport
  fi

  if [ $? -eq 0 ]; then
    echo -e "${GREEN}Build and move completed successfully!${NC}"
  else
    echo -e "${RED}Failed to move executable. Please check permissions.${NC}"
  fi
else
  echo -e "${RED}Build failed. Please check for errors.${NC}"
fi
