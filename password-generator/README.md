# Password generator

[![Code Coverage](https://codecov.io/gh/nexylan/password-generator/coverage.svg)](https://codecov.io/gh/nexylan/password-generator)
[![Build Status](https://travis-ci.org/nexylan/password-generator.svg?branch=master)](https://travis-ci.org/nexylan/password-generator)

Simple Golang password generator function using Openfaas.

## Usage

### JSON request body sample

```json
{
  "length": 15,
  "upper_case_num": 5,
  "digit_num": 2,
  "special_char_num": 2
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
  "code": 200,
  "password": "9rj8ejRKYY+Cit9"
}
```
