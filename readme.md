# Lightweight abstraction of Gopacket Lib for Packet Processing

Much in the vein of the venerable [Scapy](https://scapy.net/) library, I have decided to start a packet processing library that provides easy and useful abstractions of the [Gopacket](https://github.com/google/gopacket) library for the purposes of quickly adding network analysis and processing into Go software.

## Base API

The Base API is intended to export friendly and easy to use functions for many generic packet processing tasks.

## Primitives API

The Primitives API is intended to provide a slightly lower level abstraction of the Gopacket library so that you don't need to repeatedly do things like iterating over available network interfaces, applying different packet processing filters, and searching for default network interfaces on a device.

## TODOS

- Add test suite
- Ensure all handles to devices are closed (Perhaps already achieved with the timeout variable in pcap.OpenLive function?)
- Improved logging on functions that use pcap.Handle to see associated interface
- Add packet relay functionality
- Implement a hard limit on packets read into memory unless written to file.
- Implement true concurrency, currently running ReadHttpTraffic in goroutine leads to program main event loop closing immediately