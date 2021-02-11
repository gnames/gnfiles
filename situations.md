# possible states of remote and local

## Remote has metafile with some CIDs, local has no 'real' files.

1. local has no files, no metafile.

   Actions:

   - copy remote metaile to local dir
   - copy files from remote to local dir
   - check sha (just in case)

2. local has no files, metafile with some CIDs

   - complain that there is already metafile, ask to delete it


## Remote has no metadata. Local has files

1. Local has no metafile.

   - generate metafile
   - upload metafile
   - upload CIDS

2. Local has metafile

   - generate new metafile, sync it with exisiting metafile
   - upload metafile
   - upload CIDs

## Remote has metadata, local has files

1. Local has no metafile

   - generate new metafile, compare it with remote, ask to proceed or not
   - if no proceed, cancel
   - if yes procceed sync metafiles
   - download data
   - upload data




## sync metafiles

### path matches, sha matches

  - do nothing

### path matches, sha does not match

  - remote version goes to history with data, local gets uploaded

### only remote path exists

  - remote path goes to history with date, no file action

### only local path exists

  - meta gets new file, it is uploaded.

## Possible states of keys

### User has download key.ID, does not have sync key (does have self key)

We should not use `self`, because user can already have `ipns` connected to
it. If KeyName comes only with self key and upload is enabled, ask users if
they want to set a new key in config file.

Use key.ID for downloading files into given directory. If directory is
alredy not-empty, tell that for download it has to be totally empty.

### User has key.ID but does not have sync key

Make a warning if user uses `self` key

1. User tries to do download.

    Try to do download using key.Name, if it does not work say they have to
    upload first.

2. User sets upload

    if there are `_META` file, download from IPFS, upload to IPFS.
    Ask if it is what user wants.

### User has key.Name, but it does not exist, does not have key.ID

Show what keys do exist and ask to create a key.
