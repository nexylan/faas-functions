# Password generator

[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE)
[![Code Coverage](https://codecov.io/gh/nexylan/password-generator/coverage.svg)](https://codecov.io/gh/nexylan/password-generator)
[![Build Status](https://travis-ci.org/nexylan/password-generator.svg?branch=master)](https://travis-ci.org/nexylan/password-generator)

Simple Golang password generator function using Openfaas.

## Usage

### JSON request body sample

```json
{
  "Length": 15,
  "UpperCaseNum": 5,
  "DigitNum": 2,
  "SpecialCharNum": 2
}
```

> Default values:
> * password: 8
> * uppercase: 1
> * digit: 1
> * special char: 1

### Response

```json
{
  "Code": 200,
  "Password": "9rj8ejRKYY+Cit9"
}
```
