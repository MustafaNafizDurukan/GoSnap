# GoSnap

GoSnap is a lightweight command-line tool for capturing process snapshots on Windows, macOS, and Linux systems. It allows users to quickly and easily generate detailed reports of the current state of their system, including information on running processes, resource utilization, and more.

[![GoSnap Usage Video](https://img.youtube.com/vi/9DEV7AAU8Hk/0.jpg)](https://www.youtube.com/watch?v=9DEV7AAU8Hk)


## **Installation**

To install GoSnap, simply download the latest release for your platform from the **[GitHub releases page](https://github.com/MustafaNafizDurukan/gosnap/releases)** and extract the binary to a directory in your PATH.

Alternatively, you can build GoSnap from source by cloning this repository and running:

```
go build
```

## **Usage**

GoSnap is designed to be simple and intuitive to use. To capture a snapshot of your system, simply run:

```
gosnap -d 10m
```

This will generate a detailed report of your system's current state and save it to a timestamped file in the current directory.

For more advanced usage, GoSnap supports a variety of flags and options. Run **`gosnap --help`** to see a full list of available commands and options.

## **Contributing**

Contributions to GoSnap are always welcome! If you find a bug or have a feature request, please open an issue on the GitHub repository.

If you'd like to contribute code, please fork this repository and submit a pull request. All contributions must adhere to the guidelines outlined in the **[CONTRIBUTING.md](https://github.com/MustafaNafizDurukan/GoSnap/CONTRIBUTING.md)** file.

## **License**

GoSnap is released under the **[MIT License](https://github.com/MustafaNafizDurukan/GoSnap/LICENSE.md)**.
