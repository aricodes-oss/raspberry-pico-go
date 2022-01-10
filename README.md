# pico-go

A quick and dirty example of using [`tinygo`](https://tinygo.org/) on your Raspberry Pi Pico, or other rp2040 compatible board.

This requires you to have `tinygo` installed. The `replace` directive inside of `go.mod` is for [gopls](https://pkg.go.dev/golang.org/x/tools/gopls) to be able to pick up the `tinygo` machine package and provide autocompletion, and may not be necessary in your setup.

## Why a Makefile?

`tinygo` has a flash command, but I was unable to get it working with a PicoProbe (detailed [here](https://datasheets.raspberrypi.com/pico/getting-started-with-pico.pdf) in appendix A, or page 64).

The Makefile instead uses [`tinygo`](https://tinygo.org/) to build the binary, and [openocd](https://github.com/raspberrypi/openocd) to flash it using the picoprobe.

## Why a goroutine/module setup for an LED blinker?

I wanted to see if cooperative multitasking worked in `tinygo`, as well as module resolution.

It does.
