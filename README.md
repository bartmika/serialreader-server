# serialreader-server

## Overview

The purpose of this application is to provide a remote procedure call (gRPC) interface over an external Arduino device with a [SparkFun Weather Shield (DEV-13956)](https://github.com/sparkfun/Weather_Shield).

Supports collection of multiple time-series data from 6 different sensors and the following features:

* Temperature
* Humidity
* Pressure
* Altitude
* Illuminance
* Supports `JSON` formatted outputs
* Powered by open-source hardware and software!

## Prerequisites

You must have the following installed before proceeding. If you are missing any one of these then you cannot begin.

* ``Go 1.16.3``

## Installation

1. Please visit the [sparkfunweathershield-arduino](https://github.com/bartmika/sparkfunweathershield-arduino) repository and setup the external device and connect it to your development machine.

2. Please find out what USB port your external device is connected on. Note: please replace ``/dev/cu.usbmodem14201`` with the value on your machine, a Raspberry Pi would most likely have the value ``/dev/ttyACM0``.

3. Download the source code, build and install the application.

    ```
    GO111MODULE=on go get -u github.com/bartmika/serialreader-server
    ```

4. Run our application.

    ```
    serialreader-server --port=50052 --arduinoDevicePath="/dev/cu.usbmodem14201"
    ```

5. If you see a message saying ``gRPC server is running.`` then the application has been successfully started.

## License

This application is licensed under the **BSD 3-Clause License**. See [LICENSE](LICENSE) for more information.
