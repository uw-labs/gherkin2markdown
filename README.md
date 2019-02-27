# gherkin2markdown

A command to convert Gherkin files into Markdown.

## Installation

```
go get -u github.com/martinroddam/gherkin2markdown
```

## Usage

```
gherkin2markdown <file> [--ignoretags=<tags>]
```

or

```
gherkin2markdown <srcdir> <destdir> [--ignoretags=<tags>]
```

To suppress outputting Features or Scenarios with a particular tag, or tags, use the `--ignoretags` option and provide a comma separated list of tags, e.g.:

```
gherkin2markdown myfile.feature --ignoretags=@wip,@experimental
```

## Example

Given a file named `math.feature` with:

```gherkin
Feature: Python
  Scenario: Hello, world!
    Given a file named "main.py" with:
    """
    print("Hello, world!")
    """
    When I successfully run `python3 main.py`
    Then the stdout should contain exactly "Hello, world!"

  Scenario Outline: Add numbers
    Given a file named "main.py" with:
    """
    print(<x> + <y>)
    """
    When I successfully run `python3 main.py`
    Then the stdout should contain exactly "<z>"
    Examples:
      | x | y | z |
      | 1 | 2 | 3 |
      | 4 | 5 | 9 |
```

When I successfully run `gherkin2markdown math.feature`

Then the stdout should contain exactly:

````markdown
# Python

## Hello, world!

_Given_ a file named "main.py" with:

```
print("Hello, world!")
```

_When_ I successfully run `python3 main.py`

_Then_ the stdout should contain exactly "Hello, world!".

## Add numbers

_Given_ a file named "main.py" with:

```
print(<x> + <y>)
```

_When_ I successfully run `python3 main.py`

_Then_ the stdout should contain exactly "`<z>`".

### Examples

| x | y | z |
|---|---|---|
| 1 | 2 | 3 |
| 4 | 5 | 9 |
````

## License

[MIT](LICENSE)
