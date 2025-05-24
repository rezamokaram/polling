# Polling Service

## Project Structure & Design Notes

This project implements a polling application consisting of:

- `Polling Service:` Manages questions and options.
- `Vote Service:` Handles user answers (votes) tied to specific options.
- `Stats Aggregator:` Provides summarized reports (e.g., number of votes per option per question).

## Implementation Note

Due to time constraints, the focus was intentionally placed on:

- Writing clean and idiomatic Go code
- Ensuring core functionality works reliably

Some aspects like advanced error handling, full test coverage, and deployment optimization may be limited but the project is structured to be extended easily.

## Ideal System Design

I also designed a high-scale system architecture to efficiently handle a large volume of write operations. You can find the high-level architecture diagram below.

![architecture](https://raw.githubusercontent.com/rezamokaram/polling/refs/heads/main/docs/ideal-system-design.png)
