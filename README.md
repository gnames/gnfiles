# GNfiles allows to download/restore/update files used in GlobalNames projects

Working with multiple large files is cumbersome. To make it easier to download,
upload, and update files using a standard approach (IPFS in this case) allows
to keep data decentralized and reused with more ease.

This program is an exploration how to utilize features availabe at IPFS for
GN purposes.

## Usage

### Get directory from IPFS

The program rebuilds directory structure getting a meta-data file with
information about paths and files. In general it works as the following:

```bash
gnfinder get [source] [destination]
```

Where the source is a place to get metafile, and destination is a local
directory, probably empty, where to download all the files. If files with
the same name exist already, they will be overwritten.

```bash
# get metadata from local file system
gnfinder get /some/dir/_META.json target_dir

# get metadata from a url
gnfinder get https://ipfs.io/ipfs/Qmhash target_dir

# get metadata from an IPFS path (requires local IPFS node)
gnfinder get /ipfs/Qmhash target_dir
gnfinder get /ipns/k5hash target_dir

# get metadata from a CID (requires local IPFS node)
gnfinder get Qmhash target_dir
```

