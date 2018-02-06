# img2icesl

Hack to convert an image into IceSL 3D texture.

This project aims at converting image (png and jpeg) into a 3D texture object
in Lua script for [IceSL](http://www.loria.fr/~slefebvr/icesl/).

This is a work around until [IceSL](http://www.loria.fr/~slefebvr/icesl/) has native 2D texture and images loading capability.


It has been used in the [Lithophane 3D printing project](https://github.com/loic-fejoz/loic-fejoz-fabmoments/tree/master/lithophane).

## Usage

Download and rename depending on your platform:
* [MacOS](img2icesl-macos?raw=true)
* [Linux 64bits](img2icesl-linux64?raw=true)

```
./img2icesl -input=myimage.png > myimage.lua
```

## Source

Download or [fork on github](https://github.com/loic-fejoz/img2icesl) from https://github.com/loic-fejoz/img2icesl

## License

Go source code is released under the [MIT License](http://opensource.org/licenses/MIT).
