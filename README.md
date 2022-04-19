# launch-tui

Launch-tui is a small TUI app to manage launchD

## Install

You can either install it by donwloading the archive, uncompress it and add the folder to your path or by using homebrew :

`brew tap wecraftforfun/tools`

then

`brew install launch-tui`

## Features

### Available

- You can load and unload an Agent, list them (it's looking for agent in your `~/Library/LaunchAgents` folder).
- You can start or stop an agent
- You can delete an agent (it will erase the associated `.plist` file from your disk)

### Planned

- You can add a new agent using a user-friendly form that will create your `.plist` for you.

![launch-tui](https://s7.gifyu.com/images/launch-tui983a665fb166e1c7.gif)
