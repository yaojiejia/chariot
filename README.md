# Go Data Pipeline

![Go](https://img.shields.io/badge/Language-Go-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Build](https://github.com/yourusername/go-data-pipeline/actions/workflows/main.yml/badge.svg)

## Introduction

**Chariot** is a robust and efficient data processing tool built with Go. It is designed to extract data from various sources, perform cleaning and transformation, and store the processed data into a PostgreSQL database. This pipeline is ideal for ETL (Extract, Transform, Load) processes, enabling seamless data integration and management.

## Features

- **Data Extraction:** Connects to multiple data sources including APIs, CSV files, and more.
- **Data Cleaning:** Implements comprehensive data validation and cleaning mechanisms.
- **Data Transformation:** Transforms raw data into structured formats suitable for analysis.
- **Data Storage:** Efficiently stores cleaned and transformed data into PostgreSQL.
- **Logging & Monitoring:** Provides detailed logs for monitoring pipeline operations.
- **Scalable & Performant:** Built with Go's concurrency features for high performance.

## Architecture

![Architecture Diagram](docs/architecture.png)

1. **Extractor:** Fetches raw data from configured sources.
2. **Cleaner:** Validates and cleans the extracted data.
3. **Transformer:** Transforms data into the desired format.
4. **Loader:** Inserts the processed data into the PostgreSQL database.
5. **Logger:** Logs each step for monitoring and debugging purposes.

## Technologies

- **Language:** Go (Golang)
- **Database:** PostgreSQL, MySQL (under dev)
- **Libraries:**
  - `database/sql` for database interactions
  - `logrus` for logging
  - `viper` for configuration management

## Prerequisites

- **Go:** Version 1.16 or higher
- **PostgreSQL:** Version 12 or higher
- **Git:** For version control

## Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/yourusername/go-data-pipeline.git
   cd go-data-pipeline
