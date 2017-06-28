# qmpmeddler

## Building
`make`
binary will be at _exe/qmpmeddler on success

## Usage

Basically, just send raw JSON to QMP.

http://wiki.qemu.org/Documentation/QMP

Example usage:
```
[root@host ~]# ./qmpmeddler execute --sockfile /var/lib/libvirt/qemu/domain-3-supervm/monitor.sock --timeout 10 --json '{"execute": "query-spice"}'
Using QMP socketfile [/var/lib/libvirt/qemu/domain-3-supervm/monitor.sock]...
{"return": {"migrated": false, "enabled": true, "auth": "none", "port": 5902, "compiled-version": "0.12.4", "tls-port": 5903, "host": "10.30.76.81", "channels": [], "mouse-mode": "server"}}
[root@host ~]# 
```
