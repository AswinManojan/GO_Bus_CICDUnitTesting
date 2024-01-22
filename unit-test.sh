#!/bin/bash

# Assuming you are in the ~/Desktop/GoBus-CICD directory
mockgen -source=repository/UserRepositoryImpl.go \
  -destination=repository/user/repo_mock_test.go \
  -package=user \
  -self_package=github.com/AswinManojan/GoBus-CICD/repository/internal/user
