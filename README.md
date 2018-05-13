# BashDump
Tool for automatic backup of MongoDB to DropBox.

## Development

1. Don't forget to `git submodule update --recursive --remote`
2. Build with `docker build -t noxdew/bashdump .`

## Getting Started

NOTE: This repo has been restructured to work inside Docker. You can still use it independently. Run `dropbox_uploader.sh` to setup your account and then add `backup.sh` to cron to be run whenever you want to backup.

1. Run the container interactively with `ACCOUNTSETUP` set to `true`. This will walk you through setting up you account. Don't forget to mount `/config` to a folder on the host.
Command `docker run --rm -v "$(pwd)"/config:/config -e ACCOUNTSETUP='true' -i noxdew/bashdump`
2. Run the container normally with all of the following variables set:

`DROPBOXFOLDER`: The folder in dropbox the backups will be uploaded to. You must create the folder manually in advance. Defaults to `/`.

`BACKUPSTOKEEP`: The number of backups to be kept. If you have more backups than that the oldest will be deleted. Defaults to 365 - 1 year assuming daily backups.

`BACKUPSUFFIX`: a string which will be appended to each backup. This is to allow multiple instances to be backing up to the same folder and be able to identidy them.

`MONGODUMPPARAMS`: all parameters to be passed to `mongodump`. Empty by default, meaning it will try to connect to `localhost`.

`CRONPERIOD`: the string defining the cronjob period. Default is daily `30 2 * * *`.

### Guide for using cron to set up daily backup

1. Open crontab file `crontab -e`
2. Add a line for your job. The following like will make `backup.sh` run every night at 2:30 AM

  `30 2 * * * /path/to/backup.sh`

See [Dropbox Uploader](https://github.com/andreafabrizi/Dropbox-Uploader) for more information.

### Known problems

1. When identifying old backups to delete it doesn't consider the suffix. So if 2 instances are being backed up in the same folder, only half of the backus specified with `BACKUPSTOKEEP` will be kept.

## Ideas and Contributions
are always welcome :heart:
