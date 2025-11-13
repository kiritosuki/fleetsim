# FleetSim
FleetSim — Vehicle Fleet Scheduling Optimization and Simulation System

[简体中文](./README.md) | [English](./README_EN.md)

## Table of Contents

- [Project Overview](#project-overview)  
- [Technology Stack](#technology-stack)  
- [Project Structure](#project-structure)  
- [Build Instructions](#build-instructions)  
  - [1. Environment Setup](#1-environment-setup)  
  - [2. Clone Repository](#2-clone-repository)  
  - [3. Configure Database](#3-configure-database)  
  - [4. Install Dependencies](#4-install-dependencies)  
  - [5. Run Service](#5-run-service)  
- [Documentation](#documentation)  
- [Contributing](#contributing)  
- [License](#license)  
- [Contact](#contact)

## Project Overview

FleetSim Name Explanation:

- **Fleet**: Refers to the fleet of vehicles managed  
- **Sim**: Short for Simulation, representing the simulation of business scenarios  

This project can serve as a teaching example and can be extended into a complete vehicle transport scheduling system.

Features:

- RESTful API based on **Gin**
- Simulate vehicle status (moving, parked, idle)
- Dynamic route planning and scheduling
- Support multiple fleets and vehicles
- Real-time vehicle location and status query
- Optional visualization frontend (Vue or other frameworks)
- Support import/export of vehicle data and route configurations

## Technology Stack

**Programming Language**: Go

**Backend Framework**: [Gin](https://github.com/gin-gonic/gin)

**ORM**: [GORM](https://gorm.io/)

**Database**: MySQL

**API Documentation**: [Swagger](https://github.com/swaggo/swag)

**Dependency Management**: Go Modules

**Deployment**: Docker / Kubernetes

## Project Structure

```
fleetsim/
├─ cmd/                   				# Executable entry
│   └─ main.go
├─ config/                				# Configuration files
│   └─ config.go
├─ internal/
│   ├─ api/               				# HTTP Controller layer
│   │   └─ vehicle.go
│   ├─ service/           				# Business logic layer
│   │   └─ vehicle_service.go
│   ├─ repository/        				# Database access layer
│   │   └─ vehicle_repo.go
│   ├─ model/             				# Database entities
│   │   └─ vehicle.go
│   └─ common/            				# Utilities, unified response
│       └─ response.go
├─ scripts/               				# DB init or migration scripts
├─ docs/                  				# Swagger auto-generated docs
├─ go.mod
└─ go.sum
```

## Build Instructions

### 1. Environment Setup

**OS**: macOS / Linux / Windows

**Go**: 1.21+ (in `PATH`)

**MySQL**: 5.7+ or 8.x

**Git**: Required for cloning the repo

> Project first uploaded on Nov 13, 2025. Tested with latest LTS versions of MySQL and Go without issues.

### 2. Clone Repository

```bash
# SSH
git clone git@github.com:kiritosuki/fleetsim.git
# HTTPS
git clone https://github.com/kiritosuki/fleetsim.git
```

- Enter project root `fleetsim`

### 3. Configure Database

- Create a database named `routing`
- Execute the database initialization script `routing.sql` under `/scripts`
- Modify `/config/config.go` for database configuration (see code comments)

### 4. Install Dependencies

```go
go mod tidy
```

### 5. Run Service

```go
go run ./cmd/main.go
```

Or launch via `Goland` / `VSCode`

## Documentation

See `/docs` folder

Contains Swagger API documentation and code samples, suitable for beginners

## Contributing

1. Fork the repository and clone your fork locally
2. Add the original repository as upstream

```
# Add upstream
git remote add upstream git@github.com/kiritosuki/fleetsim.git
# Fetch latest upstream content
git fetch upstream
```

3. Create a feature branch and make changes

```
git branch feature/xxx
git switch feature/xxx
# Modify/add code
git add .
git commit -m "Description"
git push -u origin feature/xxx
```

4. Open a Pull Request on GitHub, fill in title and description

> Please follow the Go and Gin community development guidelines, and make sure all code passes unit tests before submitting a PR.

## License

This project is licensed under the [MIT License](./LICENSE)

## Contact

- Email: 3322640054@qq.com
- GitHub: https://github.com/kiritosuki
