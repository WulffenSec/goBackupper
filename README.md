## goBackupper

Creates a backup (copy) of a given directory recursively (including all the files) to a target directory. Also files not longer present in the source directory will be deleted from the target directory.

***Example:***

- Source Directory
    - Folder 1 / File 1 (present in both source and target / don't differ) nothing is done.
    - Folder 2 / File 2 (present in source but not in target) is copied to the target directory.
    - File 4 (present in both but the content differs) is copied to the target directory.
- Target Directory
    - Folder 1 / File 1 (present in both source and target / don't differ) nothing is done.
    - Folder 3 / File 3 (present in target but not in source) is deleted from the target directory.
    - File 4 (present in both but the content differs) gets replace by the source directory version.

A test folder and files is provided to check functionality of this app. Also a test.sh script is provided to reset the test environment.

***Usage:***

```shell
goBackupper /path/to/source/dir/ /path/to/target/dir/
```

--------

***Disclaimer:*** This app has the potencial to delete data you may not intend (by misunderstanding of how the tools works or by code error), has been tested do whats intended to do. but not guarantee is given. ***Use at your own risk.***
