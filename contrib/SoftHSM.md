# SoftHSM Configuration Example

_This file contains instructions about how to start working with SoftHSM._

## MacOS

### Install softhsm from source code

#### Install the pre-requisites

```bash
brew install \
  automake \
  pkg-config \
  openssl \
  sqlite \
  cppunit \
  libtool
```

#### Clone the source code

```bash
git clone https://github.com/softhsm/SoftHSMv2.git
```

_Make sure you end up on the branch called "develop" ..._

#### Compile the source code

Set the compiler first, whatever comes with x-code does not seem to work. The following exports do depend on the
version of software which was installed by homebrew!

```bash
export CC=/opt/homebrew/bin/gcc-15
export CXX=/opt/homebrew/bin/g++-15
export CPP=/opt/homebrew/bin/c++-15
```

Generate the config for the compilation. Notice that I point it to the openssl installation what homebrew installed
so it will probably needs to get adjusted time to time, also I added the `--prefix` parameter to get this installed
into my home folder.

```bash
./autogen.sh
./configure --with-openssl=/opt/homebrew/Cellar/openssl\@3/3.6.0 --prefix=~/softhsm
```

Get it compiled and installed.

```bash
make -j4
make install
```

### Install softhsm with homebrew (DOES NOT WORK)

```bash
brew install softhsm
```

## Working with softhsm

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