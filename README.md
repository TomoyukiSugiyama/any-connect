# AnyConnect cli

## Setup

```
go build
sudo ln -fs /path-your-current-dir/.cisco_vpn/any-connect /usr/local/bin/any-connect
```

> **Note**
> Please install AnyConnect.

## Usage
For general use, the `any-connect config` command is the fastest way to set up your AnyConnect cli.
Configuration file is stored in `~/.cisco_vpn/credentials`.

```sh
$ any-connect connect 
✔ YOUR_VPN_HOST
Group YOUR_VPN_GROUP
User USER_NAME
Password: **********
```

After configuration, you can use `any-connect connect / disconnect` command.

```sh
# connect
any-connect connect
# disconnect
any-connect disconnect
# check state
any-connect state
```

> **Note**
> If you want to use other command, please see help.

## Completion

```sh
# ~/.zshrc
source <(any-connect completion zsh)
compdef _any-connect any-connect
```