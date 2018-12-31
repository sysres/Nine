# Nonix Operating System

Nonix (Spelled Nonics) is not a UNIX operating system but interface with it nicely.
The idea is to host a better designed operating system on top of a running POSIX environment.
Much like [Inferno](http://www.vitanuova.com/inferno/) but using all crap from current
industry to deliver a distributed system supporting the modern world needs (like GPU rendering,
web+browser stack and so on).

Did you dream about a Plan9 environment with a modern browser? If yes, Nonix could be the answer.

## Nonix

Nonix runs as processes in your operating system but actually, it's a full-featured operating system on its all. We emphasize a good programming environment rather than performance at this moment as Nonix could support a bare metal build in the future also.

## Linux/MacOS/Windows Host

Nonix runs on Linux, OSX and Windows as a Electron App (Chrome + nodejs + webassembly).
This choice gives us the same interface on all platforms supported by electron stack. 
Nonix runs native and has access to all machine resources available to either NodeJS or Chrome (Filesystem, GPU, Keyboard, Mouse, Camera, Speaker and so on).

