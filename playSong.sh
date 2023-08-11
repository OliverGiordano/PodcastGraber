#! /bin/bash
cd /home/oliver/Documents/podcastCatcher
go run retreivePodcasts.go
file=$(ls | grep .mp3)
echo $file
play -q "$file"
