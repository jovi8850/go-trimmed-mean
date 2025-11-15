# go-trimmed-mean  
*A Go package for computing symmetric and asymmetric trimmed means with optional distribution diagnostics.*

---

## Overview

`go-trimmed-mean` is an open-source Go package designed to compute **trimmed means**, a robust estimator of central tendency that reduces the influence of extreme values. The package supports:

- **Symmetric trimming** (e.g., remove 5% from both tails)
- **Asymmetric trimming** (e.g., remove 2% low, 10% high)
- **Distribution diagnostics**, including:
  - Skewness
  - Outlier detection
  - Tail imbalance
- **Automatic trimming recommendations** based on skewness interpretation
- **Input validation & descriptive error handling**
- **Unit tests** validating correctness and edge cases

This package serves as a reusable statistical tool intended for data cleaning, robust statistical estimation, and academic coursework.

---

## Project Structure
```
C:.
│ go.mod
│ README.md
│
└───trimmedmean/
        trimmedmean.go
        trimmedmean_test.go
```


### File Descriptions

| File | Description |
|------|-------------|
| **go.mod** | Go module definition specifying module path and Go version. |
| **README.md** | Repository documentation describing installation, design, and package scope. |
| **trimmedmean/trimmedmean.go** | Core package logic containing trimmed-mean computation, symmetric/asymmetric trimming, skewness evaluation, outlier handling, and error validation. |
| **trimmedmean/trimmedmean_test.go** | Unit tests validating trimming behavior, skewness calculations, sorting correctness, and error-handling logic. |

---

## Installation

To install the package into another Go module:

### 1. Retrieve the trimmed mean package
```
go get github.com/jovi8850/go-trimmed-mean/trimmedmean
```
### 2. Import it into your Go program
```
import "github.com/jovi8850/go-trimmed-mean/trimmedmean"
```
You can now call the package’s exported functions such as:

- SymmetricTrimmedMean()
- AsymmetricTrimmedMean()
- EvaluateDistribution()

## Testing

```
go test ./...
```
The tests validate:
- Trimming correctness
- Symmetric vs. asymmetric logic
- Error handling
- Skewness & distribution diagnostics

## GenAI Tools

This project was developed with assistance from OpenAI’s ChatGPT, which was used for:

- Helping design the trimmed mean algorithm structure
- Suggesting error-handling techniques and recommended validation checks
- Reviewing and refining Go code for idiomatic clarity
- Generating template boilerplate for test files
- Creating the structure and text for this README.md

All final code was reviewed, tested, and validated by the developer before inclusion in this repository.