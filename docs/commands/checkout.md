# Checkout

Checkout checks a commit or branch out to the current workspace, it overwrites all changes.
Currently it doesn't remove files:

`svcs checkout <branch|tag|hash>`

Arguments:

  `--no-head`: This option switches if it moves the head to the newly checked out branch, tag or commit.
