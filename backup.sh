#!/bin/bash
echo Dumping database
ROOTDIR=/app/bashdump
CONFIGFILE=$ROOTDIR/.dropbox_uploader
mongodump
NEWFILE=dump`date +%Y-%m-%d`.tar.gz
echo Compressing folder to $NEWFILE
tar -zcvf $NEWFILE dump
echo Uploading to DropBox
$ROOTDIR/dropbox_uploader.sh -f $CONFIGFILE upload $NEWFILE /
echo Removing temporary data
rm -rf dump
rm $NEWFILE
echo Removing old backups
TOTAL=0
DUMPS=`$ROOTDIR/dropbox_uploader.sh -f $CONFIGFILE list | awk '{print $3}' | grep dump | sort -r`
for i in $DUMPS
do
    if [[ "$TOTAL" == 3 ]]
    then
        echo Deleting $i
        $ROOTDIR/dropbox_uploader.sh -f $CONFIGFILE delete /$i
    else
        TOTAL=$((TOTAL + 1))
    fi
done
echo Done
