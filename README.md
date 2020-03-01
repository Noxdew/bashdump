# BashDump
Tool for automatic backup of MongoDB to S3/DigitalOcean Spaces.

## Development

1. Build the executable with `go build`
2. Build the docker image with `docker build -t noxdew/bashdump .`

## Getting Started

NOTE: This repo has been restructured to work inside Kubernetes Cron Jobs and to upload to a cloud provider bucket. You can still use the older version by checking out to `01f22810647b2e2987853b775f6011bfe4403f1c` and read the README for that version.

1. To backup a DB run `bashdump dump` with the environmental variables listed below. It will create a dump and upload it to `{prefix}/{year}/{month}/{day}/dump{date}.tar.gz` in your bucket
2. To restore a DB run `bashdump restore` with the environmental variables listed below. It will find the latest backup in the prefix and restore only databases that are not present in the instance

- `DO_ACCESS_KEY`: DO or S3 access key
- `DO_SECRET_ACCESS_KEY`: DO or S3 secret key
- `DO_SPACES_ENDPOINT`: the DO or S3 endpoint
- `DO_SPACES_BUCKET`: the name of the DO Space or S3 bucket bashdump should upload to
- `DO_SPACES_PREFIX`: what is the starting path to the backups
- `MONGO_URI`: the URI to connect to the mongo instance

## Ideas and Contributions
are always welcome :heart:
