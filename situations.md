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
