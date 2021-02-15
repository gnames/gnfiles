# GNfiles allows to download/restore/update files used in GlobalNames projects

Working with multiple large files is cumbersome. To make it easier to download,
upload, and update files using a standard approach (IPFS in this case) allows
to keep data decentralized and reused with more ease.

This program is an exploration how to utilize features availabe at IPFS for
transferring files.

## Usage

Currently the app depends on running IPFS node on a local computer.
Install IPFS and run:

```bash
ipfs daemon &
```

### Upload directory to IPFS

A directory can by uploaded to IPFS if IPFS node is runnin locally.
Optionally, you can set a permalink using IPNS keys.

```bash
# To sync a directory to IPFS without permalink
gnfiles sync some_dir

# To sync a directory to IPNS with permalink (takes about a minute)
gnfiles sync somedir ipns-name

# You can find available key names with
ipfs key list
```

### Get directory from IPFS

The program rebuilds directory structure getting a meta-data file with
information about paths and files. In general it works as the following:

Files can be downloaded from anywhere, assuming that they are available.
If there are only available on one node, that node has to be running.

```bash
gnfiles get [source] [destination]
```

Where the source is a place to get metafile, and destination is a local
directory, probably empty, where to download all the files. If files with
the same name exist already, they will be overwritten.

```bash
# get metadata from local file system
gnfiles get /some/dir/_META.json target_dir

# get metadata from a url
gnfiles get https://ipfs.io/ipfs/Qmhash target_dir

# get metadata from an IPFS path (requires local IPFS node)
gnfiles get /ipfs/Qmhash target_dir
gnfiles get /ipns/k5hash target_dir

# get metadata from a CID (requires local IPFS node)
gnfiles get Qmhash target_dir
```

