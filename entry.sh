#!/bin/bash
if [ -z "$ACCOUNTSETUP" ]
then
    echo "Running BashDump cron ${CRONPERIOD:-30 2 * * *}"
    echo "${CRONPERIOD:-30 2 * * *} /backup.sh" >> /var/spool/cron/crontabs/root
    crond -f
else
    echo "Running BashDump auth setup"
    /dropbox_uploader.sh -f /config/.dropbox_uploader
fi