#!/bin/bash

/run/import_config.sh &
exec /usr/bin/supervisord -c /etc/supervisor/supervisord.conf