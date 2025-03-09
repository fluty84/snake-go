#!/bin/bash

# Clean previous build
rm -rf Snake.app

# Package application with Fyne
fyne package -os darwin -name Snake -icon icon.png -appID com.example.snake

# Launch the application
open Snake.app