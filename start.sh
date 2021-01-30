#!/bin/sh

/df-backend -log ${DF_LOG} &
/df-frontend # must be started in foreground