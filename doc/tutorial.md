# InMem Tutorial

## Abstract

This tutorial contains,

1. why InMem and the benefits of using it
2. how InMem works
3. using InMem

This document is written for beginners so It doesn't contain in-depth technical info about how InMem works.



## What is InMem

InMem(stands for ***In Memory***) is a command line utility that can be used to store and view files inside memory, like your operating systems file explorer. The only difference here is that InMem stores your files in your RAM(*R*andom *Access* *M*emory) but your operating system stores them in your hard disk or SSD(*S*olid *S*tate *D*rive).

### But storing files in memory  means that they will be deleted on shutdown

Yes, the files will be deleted on shutdown. That's one of the perks of having your file system in memory. 

### What are the other perks?

Some of the perks are,

* **Faster** than saving and deleting files in a hard drive or SSD (Fast read write speeds)
* **Safer** than storing to a hard drive or SSD because the OS only allows access   to the software's memory locations (Safe than other storage mediums)
* Files stored in memory are **Temporary**

## How does InMem work

InMem works by creating a virtual temporary file system inside your RAM and then   writes 