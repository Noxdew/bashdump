# BashDump
Tool for automatic backup of MongoDB to DropBox.

## Getting Started

1. Run `dropbox_uploader` and go through the configuration wizard and set up a connection to your DropBox. It will store the configuration file in `~/.dropbox_uploader`
2. Change [`CONFIGFILE`](https://github.com/Noxdew/bashdump/blob/master/backup.sh#L4) to point to your dropbox_uploader config file.
3. Change the number of backups to be preserved [here](https://github.com/Noxdew/bashdump/blob/master/backup.sh#L19) (the default is the last 3, all other will be deleted)
4. Set up the `backup.sh` to be run daily (or any preferred frequency) using `cron` or similar tool.

### Guide for using cron to set up daily backup

1. Open crontab file `crontab -e`
2. Add a line for your job. The following like will make `backup.sh` run every night at 2:30 AM

  `30 2 * * * /path/to/backup.sh`

See [Dropbox Uploader](https://github.com/andreafabrizi/Dropbox-Uploader) for more information.

## Ideas and Contributions
are always welcome :heart:
