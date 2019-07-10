
## What is Potato?

Potato is a command line tool which helps you to keep the development configs and settings of your Mac as a programmer/developer/coder safe in your GitHub that you
can later use to set up a new machine. Also, Potato is fully configurable so you can track anything from your `.vimrc` to all pip packages installed in your system.

## Why GO?

Because it seems exciting to develop something in Golang.

## Installing Potato

Make sure Go is installed in your system.

To install run - `go get github.com/horcrux2301/Potato`

## Using Potato

1. Add a setting

	Run `potato add` to add a new setting to potato.
	Each new setting takes in 4 arguments -
	1. Name of the setting  (eg - pip)
	2. Description of the setting (eg - Keeps track of all the pip3 packages installed)
	3. Command that needs to be executed for the setting (eg - `pip3 freeze`)
	4. Filename in which the settings will be saved (eg - pip3.txt)

	![](https://i.imgur.com/2yxuCPE.gif)
2. Delete a setting

	Run `potato delete` to delete a setting from potato.
	You need to enter the name of the setting in order to delete it.

![](https://i.imgur.com/jIdGOn8.gif)

3. Display all settings

	Run `potato display` to display all the current settings in Potato.

![](https://i.imgur.com/XhBAdu5.gif)

4. Update Git Directory

	Potato also needs to know the directory where all the files are to be saved. Run `potato git` to update the directory path. If the directory has git initialized and remote tracking enabled automatic git push to the master will also happen.

![](https://i.imgur.com/s4lvsQt.gif)

5. Run potato

	Run `potato run` to execute all the commands.

![](https://i.imgur.com/nW6MGEI.gif)

You can refer [here](https://github.com/horcrux2301/Potato/blob/master/SettingsDemo.json) to get an idea of settings that you might want to add to Potato.

Backup of my personal system - [Github Repo](https://github.com/horcrux2301/Potato-Backup)

## Suggestions/Issue ?

Submit a PR or open an Issue in the repo.
