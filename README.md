> [!CAUTION]
> This library is intended solely for my personal use, so I may not fix bugs, actively support
> the project, and do ongoing maintenance.

# Onvif

Simple management of onvif IP-devices cameras. onvif is an implementation of ONVIF protocol for
managing onvif IP devices. The purpose of this library is convenient and easy management of IP
cameras and other devices that support ONVIF standard.

## Installation

To install the library, use `go get`:

```sh
go get github.com/av1ppp/onvif
```

## Supported services

The following services are implemented:

- Device
- Media
- PTZ
- Imaging
- Event
- Discovery
- Authentication
- SOAP
- Error handling

## Using

### General concept

1. Connecting to the device
2. Authentication (if necessary)
3. Carrying out the required method

#### Connecting to the device

If there is a device on the network at the address _192.168.12.34_, and its ONVIF services use the
_5678_ port, then you can connect to the device in the following way:

```go
dev, err := onvif.NewDevice(onvif.DeviceParams{
    Xaddr: "192.168.12.34:5678",
})
```

> The ONVIF port may differ depending on the device , to find out which port to use, you can go
> to the web interface of the device. **Usually this is 80 port.**

#### Authentication

If ONVIF requires authentication, you can pass the `username` and `password` in the `DeviceParams`
when creating a new device instance:

```go
device := onvif.NewDevice(onvif.DeviceParams{
    Xaddr: "192.168.12.34:5678",
    Username: username,
    Password: password,
})
```

### Example

You can find a simple example of using in the
[examples/device-service/main.go](./examples/device-service/main.go)
