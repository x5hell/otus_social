#!/bin/sh
processors=$(cat /proc/cpuinfo | grep 'cpu cores' | awk '{print $4}' | head -1)
la1min=$(uptime | awk '{print $8 $9 $10}' | cut -d',' -f1)
la5min=$(uptime | awk '{print $8 $9 $10}' | cut -d',' -f2)
la15min=$(uptime | awk '{print $8 $9 $10}' | cut -d',' -f3)
awk "BEGIN {print $la1min/$processors, $la5min/$processors, $la15min/$processors}"