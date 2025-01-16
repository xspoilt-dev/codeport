#!/bin/bash

GREEN="\033[0;32m"
BLUE="\033[0;34m"
YELLOW="\033[1;33m"
NC="\033[0m"

echo -e "${YELLOW}Building Executable....${NC}"
go build codeport.go


if [ $? -eq 0 ]; then
  echo -e "${BLUE}Moving executable to the bin folder...${NC}"
  mv codeport ./bin/
  echo -e "${GREEN}Build and move completed successfully!${NC}"
else
  echo -e "${RED}Build failed. Please check for errors.${NC}"
fi
