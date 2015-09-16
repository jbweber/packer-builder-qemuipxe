# qemu ipxe builder for packer.io

This builder plugin adds functionality to packer.io allowing it to provision
qemu based images using ipxe to drive the build process.

This is based off of the in-tree qemu build and modified / updated to support
ipxe instead of cdrom.

XXXXXX DO NOT USE XXXXXX

this code is currently a very dirty hack and requires quite a bit of cleanup
while it works, it's not recommended for use

XXXXXX DO NOT USE XXXXXX

# example ipxe boot script

```
#!ipxe
kernel http://192.168.10.8/centos/7/os/x86_64/images/pxeboot/vmlinuz net.ifnames=0 inst.text inst.ks=http://192.168.10.8/centos7.ks ip=dhcp inst.geoloc=0 BOOTIF=01-${mac:hexhyp} initrd=initrd.img
initrd http://192.168.10.8/centos/7/os/x86_64/images/pxeboot/initrd.img
boot
```

# example builder configuration

```
{
  "builders":
  [
    {
      "type": "qemuipxe",
      "output_directory": "centos7_output",
      "ssh_wait_timeout": "30s",
      "shutdown_command": "shutdown -P now",
      "disk_size": 1000,
      "format": "raw",
      "headless": false,
      "accelerator": "kvm",
      "http_directory": "httpdir",
      "http_port_min": 10082,
      "http_port_max": 10089,
      "ssh_host_port_min": 2222,
      "ssh_host_port_max": 2229,
      "ssh_username": "root",
      "ssh_password": "vagrant",
      "ssh_port": 22,
      "ssh_wait_timeout": "90m",
      "vm_name": "tdhtest",
      "net_device": "virtio-net",
      "disk_interface": "virtio",
      "qemu_binary": "qemu-system-x86_64",
      "qemuargs": [
        [ "-m", "2048M" ]
      ]
    }
  ]
}
```

# notes

* found more ram was required for this configuration to load the initrd from pxe 2GB seemed to be enough
