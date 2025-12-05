<h1 align="center">GoMovieAPI</h1>
<p align="center">
  <img src="https://socialify.git.ci/MostafaSensei106/GoPix/image?custom_language=Go&font=KoHo&language=1&logo=https%3A%2F%2Favatars.githubusercontent.com%2Fu%2F138288138%3Fv%3D4&name=1&owner=1&pattern=Floating+Cogs&theme=Light" alt="GoMovieAPI Banner">
</p>

<p align="center">
  <strong>A high-performance, feature-rich movie API built in Go.</strong><br>
  Fast. Scalable. Efficient. All your movie data at your fingertips.
</p>

<p align="center">
  <a href="#about">About</a> •
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage-examples">Usage Examples</a> •
  <a href="#technologies">Technologies</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#license">License</a>
</p>

---

## About

Welcome to **GoMovieAPI** — a blazing-fast movie API built with Go.
GoMovieAPI provides a robust and efficient backend for managing movie data, offering seamless CRUD operations and a well-structured set of endpoints. Whether you’re building a movie review application, a content management system, or integrating with other services, GoMovieAPI delivers speed and reliability.

---

## Features

### 🌟 Core Functionality

- CRUD operations for movie data (Create, Read, Update, Delete)
- RESTful API design
- Fast and efficient data handling
- Structured data models for movies

### 🛠️ Advanced Capabilities

- **Robust Routing**: Utilizes `gorilla/mux` for powerful URL matching and dispatching.
  - Clear separation of concerns with dedicated handlers and routes.
- Scalable architecture for handling multiple requests.
- Easy to extend for additional features.

---

## Installation

## 📦 Easy Install (Linux / Windows)

> [!IMPORTANT]
> GoMovieAPI is a backend API and does not currently provide official pre-built binaries for direct download.
> The instructions below are placeholders to match the template structure.
> To run the API, please follow the instructions in the "Build from Source" section.

Download the latest pre-built binary for your platform from the [Releases](https://github.com/MostafaSensei106/GoMovieAPI/releases) page. (Note: This link will lead to a page without releases for GoMovieAPI as they are not yet available.)

### 🐧 Linux
Extract the archive
```bash
# This command is a placeholder; binaries are not yet available.
tar -xzf gomovieapi-linux-amd64.vX.Y.Z.tar.gz
```

Move the binary to the local bin directory
```bash
# This command is a placeholder; binaries are not yet available.
sudo mv linux/amd64/gomovieapi /usr/local/bin
```

If you want to install for a specific user
```bash
# This command is a placeholder; binaries are not yet available.
mv linux/amd64/gomovieapi /home/$USER/.local/bin
```

Then you can test the API by running:

```bash
# This command is a placeholder; binaries are not yet available.
gomovieapi version # (or a similar command if implemented)
```
---

### 🪟 Windows

1. Download `gomovieapi-windows-amd64-vX.Y.Z.zip` from the [Releases](https://github.com/MostafaSensei106/GoMovieAPI/releases) page. (Note: This link will lead to a page without releases for GoMovieAPI as they are not yet available.)
2. Extract the archive to a folder of your choice.
3. Move the binary located at `windows/amd64/gomovieapi.exe` to any folder of your choice or `C:\Program Files\GoMovieAPI\bin`.
4. Add that folder to your **System PATH**:

   * Open *System Properties* → *Environment Variables* → *Path* → *Edit* → *Add new*.

Then you can test the API by running:
```powershell
# This command is a placeholder; binaries are not yet available.
gomovieapi -v # (or a similar command if implemented)
```
---

## 🏗️ Build from Source (Linux, Windows)

> ![📝 Note]
> GoMovieAPI uses standard Go build commands.
> Make sure you have `Go` and `git` installed on your system.

### ⚙️ Step 1: Clone and Build

```bash
git clone https://github.com/MostafaSensei106/GoMovieAPI.git
cd GoMovieAPI
go mod tidy
go build -o GoMovieAPI
```

### ✅ Result

- This will compile GoMovieAPI from source optimized for your OS and CPU architecture.
- The executable will be named `GoMovieAPI` in the project root directory.
- You can now run the API:

```bash
./GoMovieAPI
```
The API will start on `http://localhost:8080`.

---

### 🆙 Running the API

To start the API server after building:

```bash
./GoMovieAPI
```
The API will run on `http://localhost:8080`.

---

## 🚀 Quick Start

Assuming the API is running on `http://localhost:8080`:

### Get all movies

```bash
curl http://localhost:8080/movies
```

### Get a single movie by ID

```bash
curl http://localhost:8080/movies/123
```

### Create a new movie (example data)

```bash
curl -X POST -H "Content-Type: application/json" -d '{"title": "Inception", "director": "Christopher Nolan", "year": "2010"}' http://localhost:8080/movies
```

### Update an existing movie

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"title": "Inception (Revised)", "director": "Christopher Nolan", "year": "2010"}' http://localhost:8080/movies/123
```

### Delete a movie

```bash
curl -X DELETE http://localhost:8080/movies/123
```

---

## Technologies

| Technology         | Description                                                                                             |
| ------------------ | ------------------------------------------------------------------------------------------------------- |
| 🧠 **Golang**      | [go.dev](https://go.dev) — The core language powering GoMovieAPI: fast and efficient                    |
| 🌐 **Gorilla Mux** | [github.com/gorilla/mux](https://github.com/gorilla/mux) — A powerful URL router and dispatcher for Go. |

---

## Contributing

Contributions are welcome! Here’s how to get started:

1. Fork the repository
2. Create a new branch:
   `git checkout -b feature/YourFeature`
3. Commit your changes:
   `git commit -m "Add amazing feature"`
4. Push to your branch:
   `git push origin feature/YourFeature`
5. Open a pull request

> 💡 Please open an issue first for major feature ideas or changes.

---

## License

This project is licensed under the **MIT License**.
See the [LICENSE](LICENSE) file for full details.

<p align="center">
  Made with ❤️ by <a href="https://github.com/MostafaSensei106">MostafaSensei106</a>
</p>
