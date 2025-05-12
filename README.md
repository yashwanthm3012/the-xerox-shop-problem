# Xerox Shop Problem – Microservices Architecture

A modern, microservices-based system to streamline print request management in a Xerox shop. This system enables customers to upload files and request print jobs via a web client, while shop owners manage and view these jobs through a dedicated desktop application.

## Overview

The system is composed of the following components:

- **Client Frontend**: React-based UI for customers to upload print files and enter job metadata.
- **Client Backend**: Golang API that handles file uploads, generates reference IDs, and communicates with the shop's microservice.
- **Shop Storage Microservice**: Built using Rust and Warp, this service receives files and metadata, stores them locally, and exposes an API for the shop owner’s desktop application.
- **Shop Desktop App**: Tauri + React-based GUI that displays incoming print jobs with associated metadata and file preview.

## Technologies Used

- **Golang** – For the backend API that clients interact with.
- **Rust + Warp** – For the shop-side microservice that stores and manages uploaded files and job metadata.
- **React** – For both the client-side frontend and the shop owner UI.
- **Tauri** – To wrap the shop interface as a desktop application.
- **Tokio, Warp, Futures** – For async web server and multipart upload handling in Rust.

## Features

- Upload PDF documents with detailed print job metadata.
- Generate unique reference numbers for tracking.
- Desktop app interface for shop owners to view and manage print jobs.
- Local file storage for simplicity and control – no cloud dependency.

## License

This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.

## Contact

For questions or support, please contact [wepandas4@gmail.com](mailto:wepandas4@gmail.com)
