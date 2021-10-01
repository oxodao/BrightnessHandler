# BrightnessHandler
This tool was made to get and set the screen brightness on my Huawei Matebook 14, as it have a AMD CPU, xbacklight won't work out of the box. This now has found another use in my [metaprint](https://github.com/oxodao/metaprint) software.

## Get started

Go take a look in your folder `/sys/class/backlight` to see the name of your device. 

If it is called `amdgpu_bl0` you're already good to go, it's the default value in this tool.

If it is not, you need to pass it as a environment variable: `sudo BH_DEVICE=my_device_name bh set 100`, or even better depending on your use case, use it directly as a [metaprint](https://github.com/oxodao/metaprint) module.

Required build tools:
```
- go >= 1.16, maybe
```

Compile it and move it to your system tools
```bash
$ git clone https://github.com/oxodao/BrightnessHandler.git
$ make
$ sudo make install
```

## Usage
### Display brightness
Recommended way: Through metaprint as a module
```yaml
brightness:
    i3:
        prefix: ""
        device: amdgpu_bl0
```

Through i3blocks
```
# ~/.config/i3blocks/config
[brightness]
label=  
interval=1
markup=pango
color=#676E7D
command=bh get
```

### Changing the brightness
For now, changing the brightness requires root privileges. This will no longer be the case when I'll implement proper udev rules.

```
# ~/.config/i3/config
[...]
bindsym XF86MonBrightnessUp   exec sudo bh add 10
bindsym XF86MonBrightnessDown exec sudo bh sub 10
[...]
```

```
$ sudo bh set 100
```


### Device names
Most of the time those values will work:
- amdgpu_bl0 (tested on Huawei Matebook 14 2020 / AMD Edition)
- intel_backlight (tested on Lenovo Thinkpad x230)

## Futur work
- Maybe try to make a library that also works on Windows / OSX

## Licence

BrightnessHandler
Copyright (C) 2021 Nathan <Oxodao> JANCZEWSKI

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License along
with this program; if not, write to the Free Software Foundation, Inc.,
51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
