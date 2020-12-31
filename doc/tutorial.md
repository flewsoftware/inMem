# InMem Tutorial

### This tutorial contains

1. why InMem and the benefits of using it
2. how InMem works
3. using InMem

This tutorial is written for beginners, so It doesn't contain in-depth technical info about how InMem works.



### What is InMem

InMem(stands for ***In Memory***) is a command line utility that can be used to store and view files inside memory, like your operating systems file explorer. The only difference here is that InMem stores your files in your RAM(*R*andom *Access* *M*emory) but your operating system stores them in your hard disk or SSD(*S*olid *S*tate *D*rive).

#### But storing files in memory  means that they will be deleted on shutdown

Yes, the files will be deleted on shutdown. That's one of the perks of having your file system in memory. 

#### What are the other perks?

Some perks are,

* **Faster** than saving and deleting files in a hard drive or SSD (Fast read write speeds).
* **Safer** than storing to a hard drive or SSD because the OS only allows access   to the software's memory locations (Safe than other storage mediums).
* Files stored in memory are **Temporary**.

### How does InMem work

InMem works by creating a virtual temporary file system inside your RAM and then writes files to it.

## Using InMem
**Before we start please download the latest version of inMem**   

InMem is designed to be used by humans, but there are some prerequisites that are required to understand InMem.

### Prerequisites
* Little knowledge about how RAM.
* Comfortable with command lne interfaces.

#### Installing InMem
1. Download the latest version of InMem.
2. Add the path of the InMem binary to the paths variable.

#### Starting InMem
1. Open a command prompt and type ``inmem``.
2. Wait until the InMem shell info is displayed (It shouldn't take much time).

#### Using commands
Click [here](commands.md) to learn about commands
