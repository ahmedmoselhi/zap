<p align="center">
    <img src="https://github.com/srevinsaju/zap/raw/main/assets/logo.svg" alt="zap Zap" width=128 height=128>

<h2 align="center">⚡️ Zap</h2>

  <p align="center">
    The delightful package manager for AppImages
    <br>
    <a href="https://github.com/srevinsaju/zap/issues/new">Report bug</a>
    ·
    <a href="https://github.com/srevinsaju/zap/issues/new">Request feature</a>
  </p>
</p>

<div align="center">

[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://golang.org/)<br/><br/>

[![Mentioned in Awesome AppImage](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/AppImage/awesome-appimage)

[![GitHub followers](https://img.shields.io/github/followers/srevinsaju?label=Follow%20me&style=social)](https://github.com/srevinsaju) [![GitHub stars](https://img.shields.io/github/stars/srevinsaju/zap?style=social)](https://github.com/srevinsaju/zap/stargazers)

<img src="assets/zap-firefox-install.gif" alt="Zap Installing Firefox"></img>
</div>

> Looking for the older Zap v1 (Python) implementation? Head over to [v1 branch](https://github.com/srevinsaju/zap/tree/v1).

## Getting Started ✨

### Automated Installation
For installing zap you can use our little [bash script](https://github.com/srevinsaju/zap/blob/main/install.sh).

The installer requires `curl`, `grep` `jq` and `wget` (optional). Please make sure you have it installed.

For system-wide installation (needs `sudo`)
```bash
curl https://raw.githubusercontent.com/srevinsaju/zap/main/install.sh | sudo sh
```

**Note:** Always Check bash scripts before running as sudo. You can feel free to check out install.sh, it's safe.

For local installation run
```bash
curl https://raw.githubusercontent.com/srevinsaju/zap/main/install.sh | sh
```

### Manual Installation

For system-wide installation (needs `sudo`)
```bash
sudo wget https://github.com/srevinsaju/zap/releases/download/continuous/zap-amd64 -O /usr/local/bin/zap
sudo chmod +x /usr/local/bin/zap
zap --help
```

For local installation, (requires `~/.local/bin` to be on `$PATH`)
```bash
mkdir -p ~/.local/bin
wget https://github.com/srevinsaju/zap/releases/download/continuous/zap-amd64 -O ~/.local/bin/zap
chmod +x ~/.local/bin/zap
zap --help
```

> NOTE: Replace `amd64` with your machine architecture. Supported architectures are listed in the [release](https://github.com/srevinsaju/zap/releases/tag/v2-continuous) page 

#### Installing AppImages
All AppImages from the [AppImage Catalog](https://appimage.github.io) and [AppImage catalog v2](https://g.srev.in/get-appimage) can be installed using zap with their registered name.

```bash
zap install element
```
will ask you the version of element you would like to install + download them and do all the hard work of integrating into your system, 
i.e creating desktop files, etc.

You can also install appimages from GitHub releases

```bash
zap install --github --from vscodium/vscodium
```

will put some options which will let you choose the best version for your system.

It is also possible to install AppImage from URLs

```bash
zap install --from https://f.sed.lol/wow.AppImage
```

To integrate a locally downloaded AppImage,
```bash
zap install libresprite ~/Downloads/Libresprite-x86_64.AppImage
```

... or using the `file://` protocol
```bash
zap install --from file:///home/username/Downloads/My_Super_Cool-x86_64.AppImage name_of_the_app_here
```
here, `name_of_the_app_here` specifies the name of the application. This name will be used 
as a unique identification of the AppImage, by zap, in its internal database.

 
#### Updating AppImages 🔄
AppImages can be optionally, [automatically updated using the `zapd`](#Daemon), but to achieve this manually, you need to 
```bash
zap update firefox
```
This will make use of the update information embedded within the appimage, which if exists, will be used to 'delta-update' the latest
version, by downloading "only" the parts which have changed.

For those AppImages not supporting delta updates, you can still do 
```bash
zap install zoom
```
to install the latest version of Zoom.


##### Upgrade 🚀
`zap` also supports updating all the apps together using `appimage-update`. 

```bash
zap upgrade
```


#### Configuration ⚙️
It is possible to interactively configure `zap`. All you need to do is 
```bash
zap init
```
And answer all the questions that would follow.


#### Daemon 🏃

`zapd` is a Zap AppImage daemon which periodically checks for updates.

```bash
zap daemon --install
```

This will install a `systemd` service in the local (user) level, which will spawn `zap daemon` which auto-updates 
the AppImages. 

To run the daemon (sync), do 
```bash
zap daemon
```
<br>

## Support 💸

All Pull Requests are welcome.

If you are a non-coder or was inspired by this small project, I would be glad if you would :star2: this repository, and spread the word with your friends and foes :smile:

## Credits 🙏

This project has been possible with the help and support provided by the AppImage community. Thanks to the detailed responses I received from mentors at AppImage's freenode channel.

## License ⚖️

```
MIT License

Copyright (c) 2020 Srevin Saju

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## Copyright ©️
(C) Srevin Saju 2020
