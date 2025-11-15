# go-trimmed-mean  
*A Go package for computing symmetric and asymmetric trimmed means with automated distribution diagnostics.*

---

## Overview

`go-trimmed-mean` is a robust statistical Go package designed to compute **trimmed means**—a method used to reduce the effect of extreme values by removing a percentage of data from the lower and/or upper tail of a distribution.

This version of the package supports:

### **Symmetric trimming**  
Remove the same proportion from both ends (e.g., 5% low and 5% high).

### **Asymmetric trimming**  
Remove different proportions (e.g., 0% low and 20% high) useful for skewed data.

### **Automatic trimming recommendations** via `AutoTrimmedMean`  
Analyzes the distribution and determines appropriate trimming levels based on:

- Sample **skewness**
- Presence/number of **outliers**
- Tail **imbalance**
- Sample **variation**

### **Comprehensive distribution diagnostics** (`EvaluateDistribution`)
Includes:
- Skewness  
- Empirical standard deviation  
- Outlier count via Tukey fences  
- Recommended trimming proportions  
- Tail direction (left/right skew)

### **Integer and float support**
All trimmed-mean functions have counterparts that accept `[]int`.

### **Extensive input validation & detailed error messages**
Prevents:
- Negative trimming proportions  
- LowTrim + HighTrim > 1  
- Trimming too much for sample size  
- Empty or nil slices  

### **Full unit test suite**
Validates:
- Symmetric trimming  
- Asymmetric trimming  
- Integer and float variants  
- Sorting behavior  
- Auto-trimming logic  
- Distribution evaluation  
- Extreme skew/outlier cases  
- Edge-case error conditions  

This package is intended for academic, analytical, and data-processing applications where robustness and reproducibility are required.

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
| **go.mod** | Module definition specifying Go version and repository module path. |
| **trimmedmean/trimmedmean.go** | Core implementation including trimmed mean logic, symmetric/asymmetric trimming, automatic trimming recommendation, skewness detection, input validation, and integer/float overloads. |
| **trimmedmean/trimmedmean_test.go** | Full unit test suite verifying symmetric/asymmetric trimming, AutoTrimmedMean, skewness analysis, integer handling, outlier logic, and boundary conditions. |
| **README.md** | Documentation for installation, usage, features, and AI tool disclosure. |


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

## Testing

```
go test ./...
```
The test suite validates:
- Symmetric trimming (TrimmedMean(data, trim))
- Asymmetric trimming (TrimmedMean(data, lowTrim, highTrim))
- Integer trimming (TrimmedMeanInts)
- Auto-trimming selection (AutoTrimmedMean, AutoTrimmedMeanInts)
- Distribution evaluation (skewness/outliers)
- Sorting correctness
- Error validation and boundary checks

## GenAI Tools

This project was developed with assistance from OpenAI’s ChatGPT and DeepSeek, which was used for:

- Helping design the trimmed mean algorithm structure
- Suggesting error-handling techniques and recommended validation checks
- Reviewing and refining Go code for idiomatic clarity (DeepSeek for Final Code Construction)
- Generating template boilerplate for test files 
- Creating the structure and text for this README.md

All final code was reviewed, tested, and validated by the developer before inclusion in this repository.