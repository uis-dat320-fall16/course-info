## GitHub: SSH Authentication

The steps below are meant for more experienced GitHub users. They are necessary
when you want to use SSH instead HTTPS for GitHub authentication. 

1. Register your public SSH key (normally `~/.ssh/id_rsa.pub`) at
   [GitHub](https://github.com/settings/ssh). 

2. Run the following command: `git config --global
   url."git@github.com:".insteadOf https://github.com/`. This will make Git
   rewrite GitHub URLs to use SSH instead of HTTPS. This "hack" is necessary
   since it is not possible to specify the `go get` tool to use SSH
   authentication. 

3. Go through steps 1 to 3 in the section [Go
   Assignments](https://github.com/uis-dat320-fall16/labs/blob/master/lab1/README.md#go-assignments)
   in Lab 1, but in Step 3 replace the command with `git remote add labs
   git@github.com:uis-dat320-fall16/username-labs.git` (where `username` should be
   replaced with your own GitHub username).
   
