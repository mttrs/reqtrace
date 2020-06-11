#!/bin/bash

psql $DATABASE_URL -c 'ALTER SYSTEM SET max_connections=500;' && pg_ctl restart >/dev/null
