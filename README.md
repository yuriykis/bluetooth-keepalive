[![Contributors][contributors-shield]][contributors-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/yuriykis/bluetooth-keepalive">
    <img src="logo.png" alt="Logo">
  </a>

<h3 align="center">Bluetooth KeepAlive</h3>

  <p align="center">
    Command-line tool that allows you to keep your Bluetooth devices from turning off automatically. 
    <br />
    <a href="https://github.com/github_username/repo_name/issues">Report Bug</a>
    ·
    <a href="https://github.com/github_username/repo_name/issues">Request Feature</a>
  </p>
</div>

## Bluetooth KeepAlive
Many Bluetooth devices, such as the JBL GO 2 Speaker, will turn off after a certain amount of time if there is no sound playing. This can be frustrating if you are using the device for an extended period of time, as you will need to turn it back on and reconnect it to your computer.

Bluetooth KeepAlive solves this problem by periodically checking if there is any sound currently playing on your computer. If there is no sound playing, it will set the volume to a low value, produce some sound, and then set the volume back to its previous value. This will keep your Bluetooth device from turning off, as it will think that there is still sound playing.

Currently, Bluetooth KeepAlive is only available for macOS and only supports Bluetooth speakers. However, there are TODOs in the code that make it easy to extend the tool to support other operating systems and devices, such as keyboards or mice.

## Installation
To install Bluetooth KeepAlive, you will need to have Go installed on your computer. Once you have Go installed, you can run the following command:
```sh
go install github.com/yuriykis/bluetooth-keepalive
```
This will download the source code and install the binary in your $GOPATH/bin directory.

Alternatively, you can download the source code and build the binary yourself:
```sh
make install
```
## Usage
To use Bluetooth KeepAlive, simply run the following command:
```sh
bluetooth-keepalive start
```
This will start the tool and periodically check if there is any sound playing on your computer. If there is no sound playing, it will keep your Bluetooth device from turning off. The default interval is 5 minutes, but you can change this by passing the -interval flag:
```sh
bluetooth-keepalive start --up-interval=10
```
## Contributing
If you would like to contribute to Bluetooth KeepAlive, please feel free to submit a pull request. There are TODOs in the code that make it easy to extend the tool to support other operating systems and devices.

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/yuriykis/bluetooth-keepalive.svg?style=for-the-badge
[contributors-url]: https://github.com/yuriykis/bluetooth-keepalive/graphs/contributors
[issues-shield]: https://img.shields.io/github/issues/yuriykis/bluetooth-keepalive.svg?style=for-the-badge
[issues-url]: https://github.com/yuriykis/bluetooth-keepalive/issues
[license-shield]: https://img.shields.io/github/license/yuriykis/bluetooth-keepalive.svg?style=for-the-badge
[license-url]: https://github.com/yuriykis/bluetooth-keepalive/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/yuriy-kis
