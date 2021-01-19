#!/bin/sh

/deployments-watcher -log ${DF_LOG} &
/web-app &
/node-exporter