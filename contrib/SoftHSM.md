# SoftHSM Configuration Example

_This file contains instructions about how to start working with SoftHSM._

## MacOS

### Install softhsm with homebrew

```bash
brew install softhsm
```

### Create configuration and token directory

```bash
mkdir -p ~/softhsm/tokens
mkdir -p ~/.config/softhsm2
cp -v /opt/homebrew/etc/softhsm/softhsm2.conf ~/.config/softhsm2/
```

Edit the configuration file and update the value of `directories.tokendir` to point to your newly created tokens folder. Don't forget the trailing slash.

```bash
vim ~/.config/softhsm2/softhsm2.conf
```

### Create a new slot

```bash
softhsm2-util --init-token --slot 0 --label p11tool --so-pin 1234 --pin 5678
```

### Delete an existing slot

You need a fresh start? Just delete and re-create the slot...

```bash
softhsm2-util --delete-token --token p11tool
```