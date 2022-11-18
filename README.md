# any-connect

## Setup

```
mkdir -p ~/.cisco_vpn
```

Create `credentials` in `~/.cisco_vpn`.

```
connect HOST_NAME
y
GROUP_NAME
USER_NAME
PASSWORD
```

```
go build
sudo ln -fs /path-your-current-dir/.cisco_vpn/any-connect /usr/local/bin/any-connect
```

## Usage

```sh
# connect
any-connect connect
# disconnect
any-connect disconnect
# check state
any-connect state
```

## Completion

```sh
# ~/.zshrc
source <(any-connect completion zsh)
compdef _any-connect any-connect
```