# Nine Operating System

Nine is not a UNIX operating system but interface with it nicely.
The idea is to host a better designed operating system on top of a running POSIX environment.
Much like [Inferno](http://www.vitanuova.com/inferno/) but using all crap from current
industry to deliver a distributed system supporting the modern world needs (like GPU rendering,
web+browser stack and so on).

Did you dream about a Plan9 environment with a modern browser? If yes, Nine could be the answer.

## Nine

Nine runs as processes in your operating system but actually, it's a full-featured operating system on its all. We emphasize a good programming environment rather than performance at this moment as Nine could support a bare metal build in the future also.

Nine supports to run as a daemon/server or as graphical terminal. In the former case it could be deployed in the cloud running native or inside a container. In the latter, it could run hosted inside an electron app (chrome).

## Nine Server

Nine server runs compiled to WASM on top of a V8 javascript engine like NodeJS. Then it could run anywhere nodejs is supported. It has the same programming environment as the graphical interface but lacking some resources (graphical display, camera, mouse and so on). 
You could use the graphical terminal on your notebook but using resources from a server Nine.

## Linux/MacOS/Windows Host

Nine runs on Linux, OSX and Windows as a Electron App (Chrome + nodejs + webassembly).
This choice gives us the same interface on all platforms supported by electron stack. 
Nonix runs native and has access to all machine resources available to either NodeJS or Chrome (Filesystem, GPU, Keyboard, Mouse, Camera, Speaker and so on).

