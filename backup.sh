#!/bin/bash
echo Dumping database
ROOTDIR=/config
CONFIGFILE=$ROOTDIR/.dropbox_uploader
DROPBOXFOLDER="${DROPBOXFOLDER:-/}"
BACKUPSTOKEEP="${BACKUPSTOKEEP:-365}"

mongodump $MONGODUMPPARAMS

NEWFILE=dump`date +%Y-%m-%d-%H-%M-%S`_${BACKUPSUFFIX:-mongo}.tar.gz
echo Compressing folder to $NEWFILE
tar -zcvf $NEWFILE dump

echo Uploading to DropBox
/dropbox_uploader.sh -f $CONFIGFILE upload $NEWFILE $DROPBOXFOLDER

echo Removing temporary data
rm -rf dump
rm $NEWFILE

echo Removing old backups
TOTAL=0
DUMPS=`/dropbox_uploader.sh -f $CONFIGFILE list $DROPBOXFOLDER | awk '{print $3}' | grep dump | sort -r`

case "$DROPBOXFOLDER" in
*/)
    DROPBOXFOLDERWITHSLASH=$DROPBOXFOLDER
    ;;
*)
    DROPBOXFOLDERWITHSLASH=$DROPBOXFOLDER/
    ;;
esac

for i in $DUMPS
do
    if [[ "$TOTAL" == $BACKUPSTOKEEP ]]
    then
        echo Deleting $i
        /dropbox_uploader.sh -f $CONFIGFILE delete $DROPBOXFOLDERWITHSLASH$i
    else
        TOTAL=$((TOTAL + 1))
    fi
done
echo Done
